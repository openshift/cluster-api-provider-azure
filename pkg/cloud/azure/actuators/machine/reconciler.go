/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package machine

import (
	"context"
	"encoding/base64"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/azure-sdk-for-go/services/network/mgmt/2018-12-01/network"
	"github.com/Azure/go-autorest/autorest/to"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/pkg/errors"
	apicorev1 "k8s.io/api/core/v1"
	"k8s.io/klog"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/converters"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/availabilityzones"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/disks"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/networkinterfaces"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/publicips"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/virtualmachineextensions"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/virtualmachines"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	// DefaultBootstrapTokenTTL default ttl for bootstrap token
	DefaultBootstrapTokenTTL = 10 * time.Minute

	// MachineRegionLabelName as annotation name for a machine region
	MachineRegionLabelName = "machine.openshift.io/region"

	// MachineAZLabelName as annotation name for a machine AZ
	MachineAZLabelName = "machine.openshift.io/zone"

	// MachineInstanceStateAnnotationName as annotation name for a machine instance state
	MachineInstanceStateAnnotationName = "machine.openshift.io/instance-state"

	// MachineInstanceTypeLabelName as annotation name for a machine instance type
	MachineInstanceTypeLabelName = "machine.openshift.io/instance-type"
)

// Reconciler are list of services required by cluster actuator, easy to create a fake
type Reconciler struct {
	machine       *machinev1.Machine
	machineConfig *v1beta1.AzureMachineProviderSpec
	client        client.Client

	scope                 *actuators.Scope
	availabilityZonesSvc  azure.Service
	networkInterfacesSvc  azure.Service
	publicIPSvc           azure.Service
	virtualMachinesSvc    azure.Service
	virtualMachinesExtSvc azure.Service
	disksSvc              azure.Service
}

// NewReconciler populates all the services based on input scope
func NewReconciler(scope *actuators.Scope, client client.Client, machine *machinev1.Machine, machineConfig *v1beta1.AzureMachineProviderSpec) *Reconciler {
	return &Reconciler{
		scope:                 scope,
		client:                client,
		machine:               machine,
		machineConfig:         machineConfig,
		availabilityZonesSvc:  availabilityzones.NewService(scope),
		networkInterfacesSvc:  networkinterfaces.NewService(scope),
		virtualMachinesSvc:    virtualmachines.NewService(scope),
		virtualMachinesExtSvc: virtualmachineextensions.NewService(scope),
		publicIPSvc:           publicips.NewService(scope),
		disksSvc:              disks.NewService(scope),
	}
}

// Create creates machine if and only if machine exists, handled by cluster-api
func (s *Reconciler) Create(ctx context.Context) (*compute.VirtualMachine, error) {
	// TODO: update once machine controllers have a way to indicate a machine has been provisoned. https://github.com/kubernetes-sigs/cluster-api/issues/253
	// Seeing a node cannot be purely relied upon because the provisioned control plane will not be registering with
	// the stack that provisions it.
	nicName := azure.GenerateNetworkInterfaceName(s.machine.Name)
	if err := s.createNetworkInterface(ctx, nicName); err != nil {
		return nil, errors.Wrapf(err, "failed to create nic %s for machine %s", nicName, s.machine.Name)
	}

	vm, err := s.createVirtualMachine(ctx, nicName)
	if err != nil {
		return vm, fmt.Errorf("failed to create vm %s: %v", s.machine.Name, err)
	}

	return vm, nil
}

// Update updates machine if and only if machine exists, handled by cluster-api
func (s *Reconciler) Update(ctx context.Context) (*compute.VirtualMachine, error) {
	vmSpec := &virtualmachines.Spec{
		Name: s.machine.Name,
	}
	vmInterface, err := s.virtualMachinesSvc.Get(ctx, vmSpec)
	if err != nil {
		return nil, errors.Errorf("failed to get vm: %+v", err)
	}

	vm, ok := vmInterface.(compute.VirtualMachine)
	if !ok {
		return nil, errors.New("returned incorrect vm interface")
	}

	// We can now compare the various Azure state to the state we were passed.
	// We will check immutable state first, in order to fail quickly before
	// moving on to state that we can mutate.
	if isMachineOutdated(s.machineConfig, converters.SDKToVM(vm)) {
		return nil, errors.Errorf("found attempt to change immutable state")
	}

	// TODO: Uncomment after implementing tagging.
	// Ensure that the tags are correct.
	/*
		_, err = a.ensureTags(computeSvc, machine, scope.MachineStatus.VMID, scope.MachineConfig.AdditionalTags)
		if err != nil {
			return errors.Errorf("failed to ensure tags: %+v", err)
		}
	*/

	return &vm, nil
}

func getVMState(vm compute.VirtualMachine) v1beta1.VMState {
	if vm.ProvisioningState == nil {
		return ""
	}

	if *vm.ProvisioningState != "Succeeded" {
		return v1beta1.VMState(*vm.ProvisioningState)
	}

	if vm.InstanceView == nil || vm.InstanceView.Statuses == nil {
		return ""
	}

	for _, status := range *vm.InstanceView.Statuses {
		if status.Code == nil {
			continue
		}
		switch *status.Code {
		case "ProvisioningState/succeeded":
			continue
		case "PowerState/starting":
			return v1beta1.VMStateStarting
		case "PowerState/running":
			return v1beta1.VMStateRunning
		case "PowerState/stopping":
			return v1beta1.VMStateStopping
		case "PowerState/stopped":
			return v1beta1.VMStateStopped
		case "PowerState/deallocating":
			return v1beta1.VMStateDeallocating
		case "PowerState/deallocated":
			return v1beta1.VMStateDeallocated
		default:
			return v1beta1.VMStateUnknown
		}
	}

	return ""
}

// Exists checks if machine exists
func (s *Reconciler) Exists(ctx context.Context) (bool, error) {
	vmSpec := &virtualmachines.Spec{
		Name: s.machine.Name,
	}
	vmInterface, err := s.virtualMachinesSvc.Get(ctx, vmSpec)

	if err != nil && vmInterface == nil {
		return false, nil
	}

	if err != nil {
		return false, errors.Wrap(err, "Failed to get vm")
	}

	vm, ok := vmInterface.(compute.VirtualMachine)
	if !ok {
		return false, errors.Errorf("returned incorrect vm interface: %T", vmInterface)
	}

	klog.Infof("Found vm for machine %s", s.machine.Name)

	if s.machineConfig.UserDataSecret == nil {
		vmExtSpec := &virtualmachineextensions.Spec{
			Name:   "startupScript",
			VMName: s.machine.Name,
		}

		vmExt, err := s.virtualMachinesExtSvc.Get(ctx, vmExtSpec)
		if err != nil && vmExt == nil {
			return false, nil
		}

		if err != nil {
			return false, errors.Wrapf(err, "failed to get vm extension")
		}
	}

	switch v1beta1.VMState(*vm.ProvisioningState) {
	case v1beta1.VMStateSucceeded:
		klog.Infof("Machine %v is running", to.String(vm.VMID))
	case v1beta1.VMStateUpdating:
		klog.Infof("Machine %v is updating", to.String(vm.VMID))
	default:
		return false, nil
	}

	return true, nil
}

// Delete reconciles all the services in pre determined order
func (s *Reconciler) Delete(ctx context.Context) error {
	vmSpec := &virtualmachines.Spec{
		Name: s.machine.Name,
	}

	err := s.virtualMachinesSvc.Delete(ctx, vmSpec)
	if err != nil {
		return errors.Wrapf(err, "failed to delete machine")
	}

	osDiskSpec := &disks.Spec{
		Name: azure.GenerateOSDiskName(s.machine.Name),
	}
	err = s.disksSvc.Delete(ctx, osDiskSpec)
	if err != nil {
		return errors.Wrapf(err, "failed to delete OS disk")
	}

	if s.machineConfig.Vnet == "" {
		return errors.Errorf("MachineConfig vnet is missing on machine %s", s.machine.Name)
	}

	networkInterfaceSpec := &networkinterfaces.Spec{
		Name:     azure.GenerateNetworkInterfaceName(s.machine.Name),
		VnetName: s.machineConfig.Vnet,
	}

	err = s.networkInterfacesSvc.Delete(ctx, networkInterfaceSpec)
	if err != nil {
		return errors.Wrapf(err, "Unable to delete network interface")
	}

	if s.machineConfig.PublicIP {
		publicIPName, err := azure.GenerateMachinePublicIPName(s.scope.Cluster.Name, s.machine.Name)
		if err != nil {
			return errors.Wrap(err, "unable to create Public IP")
		}

		err = s.publicIPSvc.Delete(ctx, &publicips.Spec{
			Name: publicIPName,
		})
		if err != nil {
			return errors.Wrap(err, "unable to delete Public IP")
		}
	}

	return nil
}

// isMachineOutdated checks that no immutable fields have been updated in an
// Update request.
// Returns a bool indicating if an attempt to change immutable state occurred.
//  - true:  An attempt to change immutable state occurred.
//  - false: Immutable state was untouched.
func isMachineOutdated(machineSpec *v1beta1.AzureMachineProviderSpec, vm *v1beta1.VM) bool {
	// VM Size
	if !strings.EqualFold(machineSpec.VMSize, vm.VMSize) {
		return true
	}

	// TODO: Add additional checks for immutable fields

	// No immutable state changes found.
	return false
}

func (s *Reconciler) getZone(ctx context.Context) (string, error) {
	return to.String(s.machineConfig.Zone), nil
}

func (s *Reconciler) createNetworkInterface(ctx context.Context, nicName string) error {
	if s.machineConfig.Vnet == "" {
		return errors.Errorf("MachineConfig vnet is missing on machine %s", s.machine.Name)
	}

	networkInterfaceSpec := &networkinterfaces.Spec{
		Name:     nicName,
		VnetName: s.machineConfig.Vnet,
	}

	if s.machineConfig.Subnet == "" {
		return errors.Errorf("MachineConfig subnet is missing on machine %s, skipping machine creation", s.machine.Name)
	}

	networkInterfaceSpec.SubnetName = s.machineConfig.Subnet

	if s.machineConfig.PublicLoadBalancer != "" {
		networkInterfaceSpec.PublicLoadBalancerName = s.machineConfig.PublicLoadBalancer
		if s.machineConfig.NatRule != nil {
			networkInterfaceSpec.NatRule = s.machineConfig.NatRule
		}
	}
	if s.machineConfig.InternalLoadBalancer != "" {
		networkInterfaceSpec.InternalLoadBalancerName = s.machineConfig.InternalLoadBalancer
	}

	if s.machineConfig.PublicIP {
		publicIPName, err := azure.GenerateMachinePublicIPName(s.scope.Cluster.Name, s.machine.Name)
		if err != nil {
			return errors.Wrap(err, "unable to create Public IP")
		}
		err = s.publicIPSvc.CreateOrUpdate(ctx, &publicips.Spec{Name: publicIPName})
		if err != nil {
			return errors.Wrap(err, "unable to create Public IP")
		}
		networkInterfaceSpec.PublicIP = publicIPName
	}

	err := s.networkInterfacesSvc.CreateOrUpdate(ctx, networkInterfaceSpec)
	if err != nil {
		return errors.Wrap(err, "unable to create VM network interface")
	}

	return err
}

func (s *Reconciler) createVirtualMachine(ctx context.Context, nicName string) (*compute.VirtualMachine, error) {
	decoded, err := base64.StdEncoding.DecodeString(s.machineConfig.SSHPublicKey)
	if err != nil {
		errors.Wrapf(err, "failed to decode ssh public key")
	}

	vmSpec := &virtualmachines.Spec{
		Name: s.machine.Name,
	}

	vmInterface, err := s.virtualMachinesSvc.Get(ctx, vmSpec)
	if err != nil && vmInterface == nil {
		zone, err := s.getZone(ctx)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to get zone")
		}

		if s.machineConfig.ManagedIdentity == "" {
			return nil, errors.Errorf("MachineConfig managedIdentity is missing on machine %s", s.machine.Name)
		}

		vmSpec = &virtualmachines.Spec{
			Name:            s.machine.Name,
			NICName:         nicName,
			SSHKeyData:      string(decoded),
			Size:            s.machineConfig.VMSize,
			OSDisk:          s.machineConfig.OSDisk,
			Image:           s.machineConfig.Image,
			Zone:            zone,
			Tags:            s.machineConfig.Tags,
			ManagedIdentity: azure.GenerateManagedIdentityName(s.scope.SubscriptionID, s.scope.ClusterConfig.ResourceGroup, s.machineConfig.ManagedIdentity),
		}

		userData, userDataErr := s.getCustomUserData()
		if userDataErr != nil {
			return nil, errors.Wrapf(userDataErr, "failed to get custom script data")
		}

		if userData != "" {
			vmSpec.CustomData = userData
		}

		// This always returns and it's always nil
		err = s.virtualMachinesSvc.CreateOrUpdate(ctx, vmSpec)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to create or get machine")
		}
		return nil, nil
	} else if err != nil {
		return nil, errors.Wrap(err, "failed to get vm")
	} else {
		vm, ok := vmInterface.(compute.VirtualMachine)
		if !ok {
			return nil, errors.New("returned incorrect vm interface")
		}
		if vm.ProvisioningState == nil {
			return nil, errors.Errorf("vm %s is nil provisioning state, reconcile", s.machine.Name)
		}

		if *vm.ProvisioningState == "Failed" {
			// If VM failed provisioning, delete it so it can be recreated
			err = s.Delete(ctx)
			if err != nil {
				return nil, errors.Wrapf(err, "failed to delete machine")
			}
			return nil, errors.Errorf("vm %s is deleted, retry creating in next reconcile", s.machine.Name)
		}

		return &vm, nil
	}
}

func (s *Reconciler) getCustomUserData() (string, error) {
	if s.machineConfig.UserDataSecret == nil {
		return "", nil
	}
	var userDataSecret apicorev1.Secret
	if err := s.client.Get(context.Background(), client.ObjectKey{Namespace: s.machine.Namespace, Name: s.machineConfig.UserDataSecret.Name}, &userDataSecret); err != nil {
		return "", errors.Wrapf(err, "error getting user data secret %s in namespace %s", s.machineConfig.UserDataSecret.Name, s.machine.Namespace)
	}
	data, exists := userDataSecret.Data["userData"]
	if !exists {
		return "", errors.Errorf("Secret %v/%v does not have userData field set. Thus, no user data applied when creating an instance.", s.machine.Namespace, s.machineConfig.UserDataSecret.Name)
	}

	return base64.StdEncoding.EncodeToString(data), nil
}

func (s *Reconciler) getNetworkAddresses(ctx context.Context, vm compute.VirtualMachine) ([]apicorev1.NodeAddress, error) {
	networkAddresses := []apicorev1.NodeAddress{}

	// The computer name for a VM instance is the hostname of the VM
	// TODO(jchaloup): find a way how to propagete the hostname change in case
	// someone/something changes the hostname inside the VM
	if vm.OsProfile != nil && vm.OsProfile.ComputerName != nil {
		networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
			Type:    apicorev1.NodeHostName,
			Address: *vm.OsProfile.ComputerName,
		})

		// csr approved requires node internal dns name to be equal to a node name
		networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
			Type:    apicorev1.NodeInternalDNS,
			Address: *vm.OsProfile.ComputerName,
		})
	}

	if vm.NetworkProfile != nil && vm.NetworkProfile.NetworkInterfaces != nil {
		if s.machineConfig.Vnet == "" {
			return nil, errors.Errorf("MachineConfig vnet is missing on machine %s", s.machine.Name)
		}

		for _, iface := range *vm.NetworkProfile.NetworkInterfaces {
			// Get iface name from the ID
			ifaceName := path.Base(*iface.ID)
			networkIface, err := s.networkInterfacesSvc.Get(ctx, &networkinterfaces.Spec{
				Name:     ifaceName,
				VnetName: s.machineConfig.Vnet,
			})
			if err != nil {
				klog.Errorf("Unable to get %q network interface: %v", ifaceName, err)
				continue
			}

			niface, ok := networkIface.(network.Interface)
			if !ok {
				klog.Errorf("Network interfaces get returned invalid network interface, getting %T instead", networkIface)
				continue
			}

			// Internal dns name consists of a hostname and internal dns suffix
			if niface.InterfacePropertiesFormat.DNSSettings != nil && niface.InterfacePropertiesFormat.DNSSettings.InternalDomainNameSuffix != nil && vm.OsProfile != nil && vm.OsProfile.ComputerName != nil {
				networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
					Type:    apicorev1.NodeInternalDNS,
					Address: fmt.Sprintf("%s.%s", *vm.OsProfile.ComputerName, *niface.InterfacePropertiesFormat.DNSSettings.InternalDomainNameSuffix),
				})
			}

			if niface.InterfacePropertiesFormat.IPConfigurations == nil {
				continue
			}

			for _, ipConfig := range *niface.InterfacePropertiesFormat.IPConfigurations {
				if ipConfig.PrivateIPAddress != nil {
					networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
						Type:    apicorev1.NodeInternalIP,
						Address: *ipConfig.PrivateIPAddress,
					})
				}

				if ipConfig.PublicIPAddress != nil && ipConfig.PublicIPAddress.ID != nil {
					publicIPInterface, publicIPErr := s.publicIPSvc.Get(ctx, &publicips.Spec{Name: path.Base(*ipConfig.PublicIPAddress.ID)})
					if publicIPErr != nil {
						klog.Errorf("Unable to get %q public IP: %v", path.Base(*ipConfig.PublicIPAddress.ID), publicIPErr)
						continue
					}

					ip, ok := publicIPInterface.(network.PublicIPAddress)
					if !ok {
						klog.Errorf("Public ip get returned invalid network interface, getting %T instead", publicIPInterface)
						continue
					}

					if ip.IPAddress != nil {
						networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
							Type:    apicorev1.NodeExternalIP,
							Address: *ip.IPAddress,
						})
					}

					if ip.DNSSettings != nil && ip.DNSSettings.Fqdn != nil {
						networkAddresses = append(networkAddresses, apicorev1.NodeAddress{
							Type:    apicorev1.NodeExternalDNS,
							Address: *ip.DNSSettings.Fqdn,
						})
					}
				}
			}
		}
	}

	return networkAddresses, nil
}
