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

package virtualmachines

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/compute/mgmt/compute"
	"github.com/Azure/azure-sdk-for-go/profiles/2019-03-01/network/mgmt/network"
	"github.com/Azure/go-autorest/autorest/to"
	"golang.org/x/crypto/ssh"
	"k8s.io/klog/v2"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/networkinterfaces"
)

// StackHubSpec input specification for Get/CreateOrUpdate/Delete calls
type StackHubSpec struct {
	Name            string
	NICName         string
	SSHKeyData      string
	Size            string
	Zone            string
	Image           v1beta1.Image
	OSDisk          v1beta1.OSDisk
	CustomData      string
	ManagedIdentity string
	Tags            map[string]string
	SecurityProfile *v1beta1.SecurityProfile
}

// Get provides information about a virtual network.
func (s *StackHubService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	vmSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return compute.VirtualMachine{}, errors.New("invalid vm specification")
	}
	vm, err := s.Client.Get(ctx, s.Scope.MachineConfig.ResourceGroup, vmSpec.Name, compute.InstanceView)
	if err != nil && azure.ResourceNotFound(err) {
		return nil, fmt.Errorf("vm %s not found: %w", vmSpec.Name, err)
	} else if err != nil {
		return vm, err
	}
	return vm, nil
}

// CreateOrUpdate creates or updates a virtual network.
func (s *StackHubService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	vmSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("invalid vm specification")
	}

	klog.V(2).Infof("getting nic %s", vmSpec.NICName)
	nicInterface, err := networkinterfaces.NewService(s.Scope).Get(ctx, &networkinterfaces.Spec{Name: vmSpec.NICName})
	if err != nil {
		return err
	}
	nic, ok := nicInterface.(network.Interface)
	if !ok {
		return errors.New("error getting network security group")
	}
	klog.V(2).Infof("got nic %s", vmSpec.NICName)

	klog.V(2).Infof("creating vm %s ", vmSpec.Name)

	virtualMachine, err := deriveVirtualMachineParametersStackHub(vmSpec, s.Scope.MachineConfig.Location, s.Scope.SubscriptionID, nic)
	if err != nil {
		return err
	}

	future, err := s.Client.CreateOrUpdate(
		ctx,
		s.Scope.MachineConfig.ResourceGroup,
		vmSpec.Name,
		*virtualMachine)
	if err != nil {
		return fmt.Errorf("cannot create vm: %w", err)
	}

	// Do not wait until the operation completes. Just check the result
	// so the call to Create actuator operation is async.
	_, err = future.Result(s.Client)
	if err != nil {
		return err
	}

	klog.V(2).Infof("successfully created vm %s ", vmSpec.Name)
	return err
}

func generateOSProfileStackHub(vmSpec *StackHubSpec) (*compute.OSProfile, error) {
	sshKeyData := vmSpec.SSHKeyData
	if sshKeyData == "" && compute.OperatingSystemTypes(vmSpec.OSDisk.OSType) != compute.Windows {
		privateKey, perr := rsa.GenerateKey(rand.Reader, 2048)
		if perr != nil {
			return nil, fmt.Errorf("Failed to generate private key: %w", perr)
		}

		publicRsaKey, perr := ssh.NewPublicKey(&privateKey.PublicKey)
		if perr != nil {
			return nil, fmt.Errorf("Failed to generate public key: %w", perr)
		}
		sshKeyData = string(ssh.MarshalAuthorizedKey(publicRsaKey))
	}

	randomPassword, err := GenerateRandomString(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate random string: %w", err)
	}

	osProfile := &compute.OSProfile{
		ComputerName:  to.StringPtr(vmSpec.Name),
		AdminUsername: to.StringPtr(azure.DefaultUserName),
		AdminPassword: to.StringPtr(randomPassword),
	}

	if compute.OperatingSystemTypes(vmSpec.OSDisk.OSType) == compute.Windows {
		osProfile.WindowsConfiguration = &compute.WindowsConfiguration{
			EnableAutomaticUpdates: to.BoolPtr(false),
			AdditionalUnattendContent: &[]compute.AdditionalUnattendContent{
				{
					PassName:      "OobeSystem",
					ComponentName: "Microsoft-Windows-Shell-Setup",
					SettingName:   "AutoLogon",
					Content:       to.StringPtr(fmt.Sprintf(winAutoLogonFormatString, *osProfile.AdminUsername, *osProfile.AdminPassword)),
				},
				{
					PassName:      "OobeSystem",
					ComponentName: "Microsoft-Windows-Shell-Setup",
					SettingName:   "FirstLogonCommands",
					Content:       to.StringPtr(winFirstLogonCommandsString),
				},
			},
		}
	} else if sshKeyData != "" {
		osProfile.LinuxConfiguration = &compute.LinuxConfiguration{
			SSH: &compute.SSHConfiguration{
				PublicKeys: &[]compute.SSHPublicKey{
					{
						Path:    to.StringPtr(fmt.Sprintf("/home/%s/.ssh/authorized_keys", azure.DefaultUserName)),
						KeyData: to.StringPtr(sshKeyData),
					},
				},
			},
		}
	}

	if vmSpec.CustomData != "" {
		osProfile.CustomData = to.StringPtr(vmSpec.CustomData)
	}

	return osProfile, nil
}

// Derive virtual machine parameters for CreateOrUpdate API call based
// on the provided virtual machine specification, resource location,
// subscription ID, and the network interface.
func deriveVirtualMachineParametersStackHub(vmSpec *StackHubSpec, location string, subscription string, nic network.Interface) (*compute.VirtualMachine, error) {
	osProfile, err := generateOSProfileStackHub(vmSpec)
	if err != nil {
		return nil, err
	}

	imageReference := &compute.ImageReference{
		Publisher: to.StringPtr(vmSpec.Image.Publisher),
		Offer:     to.StringPtr(vmSpec.Image.Offer),
		Sku:       to.StringPtr(vmSpec.Image.SKU),
		Version:   to.StringPtr(vmSpec.Image.Version),
	}
	if vmSpec.Image.ResourceID != "" {
		imageReference = &compute.ImageReference{
			ID: to.StringPtr(fmt.Sprintf("/subscriptions/%s%s", subscription, vmSpec.Image.ResourceID)),
		}
	}

	virtualMachine := &compute.VirtualMachine{
		Location: to.StringPtr(location),
		Tags:     getTagListFromSpecStackHub(vmSpec),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.VirtualMachineSizeTypes(vmSpec.Size),
			},
			StorageProfile: &compute.StorageProfile{
				ImageReference: imageReference,
				OsDisk: &compute.OSDisk{
					Name:         to.StringPtr(fmt.Sprintf("%s_OSDisk", vmSpec.Name)),
					OsType:       compute.OperatingSystemTypes(vmSpec.OSDisk.OSType),
					CreateOption: compute.DiskCreateOptionTypesFromImage,
					DiskSizeGB:   to.Int32Ptr(vmSpec.OSDisk.DiskSizeGB),
					ManagedDisk: &compute.ManagedDiskParameters{
						StorageAccountType: compute.StorageAccountTypes(vmSpec.OSDisk.ManagedDisk.StorageAccountType),
					},
				},
			},
			OsProfile: osProfile,
			NetworkProfile: &compute.NetworkProfile{
				NetworkInterfaces: &[]compute.NetworkInterfaceReference{
					{
						ID: nic.ID,
						NetworkInterfaceReferenceProperties: &compute.NetworkInterfaceReferenceProperties{
							Primary: to.BoolPtr(true),
						},
					},
				},
			},
		},
	}

	if vmSpec.Zone != "" {
		zones := []string{vmSpec.Zone}
		virtualMachine.Zones = &zones
	}

	return virtualMachine, nil
}

// Delete deletes the virtual network with the provided name.
func (s *StackHubService) Delete(ctx context.Context, spec azure.Spec) error {
	vmSpec, ok := spec.(*StackHubSpec)
	if !ok {
		return errors.New("invalid vm Specification")
	}
	klog.V(2).Infof("deleting vm %s ", vmSpec.Name)
	future, err := s.Client.Delete(ctx, s.Scope.MachineConfig.ResourceGroup, vmSpec.Name)
	if err != nil && azure.ResourceNotFound(err) {
		// already deleted
		return nil
	}
	if err != nil {
		return fmt.Errorf("failed to delete vm %s in resource group %s: %w", vmSpec.Name, s.Scope.MachineConfig.ResourceGroup, err)
	}

	// Do not wait until the operation completes. Just check the result
	// so the call to Delete actuator operation is async.
	_, err = future.Result(s.Client)
	if err != nil {
		return err
	}

	klog.V(2).Infof("successfully deleted vm %s ", vmSpec.Name)
	return err
}

func getTagListFromSpecStackHub(spec *StackHubSpec) map[string]*string {
	if len(spec.Tags) < 1 {
		return nil
	}

	tagList := map[string]*string{}
	for key, element := range spec.Tags {
		tagList[key] = to.StringPtr(element)
	}
	return tagList
}
