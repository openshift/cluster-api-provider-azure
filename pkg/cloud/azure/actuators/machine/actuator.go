/*
Copyright 2018 The Kubernetes Authors.

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
	"fmt"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	clusterv1 "github.com/openshift/cluster-api/pkg/apis/cluster/v1alpha1"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	client "github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/typed/machine/v1beta1"
	controllerError "github.com/openshift/cluster-api/pkg/controller/error"
	apierrors "github.com/openshift/cluster-api/pkg/errors"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/record"
	"k8s.io/klog"
	providerconfig "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	controllerclient "sigs.k8s.io/controller-runtime/pkg/client"
)

const (
	createEventAction = "Create"
	updateEventAction = "Update"
	deleteEventAction = "Delete"
	noEventAction     = ""

	requeueAfterSeconds = 20
)

//+kubebuilder:rbac:groups=azureprovider.k8s.io,resources=azuremachineproviderconfigs;azuremachineproviderstatuses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=machines;machines/status;machinedeployments;machinedeployments/status;machinesets;machinesets/status;machineclasses,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=cluster.k8s.io,resources=clusters;clusters/status,verbs=get;list;watch
//+kubebuilder:rbac:groups="",resources=nodes;events,verbs=get;list;watch;create;update;patch;delete

// Actuator is responsible for performing machine reconciliation.
type Actuator struct {
	client            client.MachineV1beta1Interface
	coreClient        controllerclient.Client
	eventRecorder     record.EventRecorder
	codec             *providerconfig.AzureProviderConfigCodec
	reconcilerBuilder func(scope *actuators.MachineScope) *Reconciler
}

// ActuatorParams holds parameter information for Actuator.
type ActuatorParams struct {
	Client            client.MachineV1beta1Interface
	CoreClient        controllerclient.Client
	EventRecorder     record.EventRecorder
	Codec             *providerconfig.AzureProviderConfigCodec
	ReconcilerBuilder func(scope *actuators.MachineScope) *Reconciler
}

// NewActuator returns an actuator.
func NewActuator(params ActuatorParams) *Actuator {
	return &Actuator{
		client:            params.Client,
		coreClient:        params.CoreClient,
		eventRecorder:     params.EventRecorder,
		codec:             params.Codec,
		reconcilerBuilder: params.ReconcilerBuilder,
	}
}

// Set corresponding event based on error. It also returns the original error
// for convenience, so callers can do "return handleMachineError(...)".
func (a *Actuator) handleMachineError(machine *machinev1.Machine, err *apierrors.MachineError, eventAction string) error {
	if eventAction != noEventAction {
		a.eventRecorder.Eventf(machine, corev1.EventTypeWarning, "Failed"+eventAction, "%v: %v", err.Reason, err.Message)
	}

	klog.Errorf("Machine error: %v", err.Message)
	return err
}

// Create creates a machine and is invoked by the machine controller.
func (a *Actuator) Create(ctx context.Context, cluster *clusterv1.Cluster, machine *machinev1.Machine) error {
	klog.Infof("Creating machine %v", machine.Name)

	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{
		Machine:    machine,
		Cluster:    nil,
		Client:     a.client,
		CoreClient: a.coreClient,
	})
	if err != nil {
		return a.handleMachineError(machine, apierrors.CreateMachine("failed to create machine %q scope: %v", machine.Name, err), createEventAction)
	}

	vm, err := a.reconcilerBuilder(scope).Create(context.Background())
	if vm != nil {
		modMachine, err := a.setMachineCloudProviderSpecifics(machine, *vm)
		if err != nil {
			klog.Errorf("%s: error updating machine cloud provider specifics: %v", machine.Name, err)
		} else {
			machine = modMachine
		}

		modMachine, err = a.updateStatus(ctx, machine, a.reconcilerBuilder(scope), *vm)
		if err != nil {
			klog.Infof("%s: failed to update status: %v", machine.Name, err)
			return &controllerError.RequeueAfterError{RequeueAfter: requeueAfterSeconds * time.Second}
		}
		machine = modMachine

		modMachine, err = a.updateProviderID(machine, *vm, scope.SubscriptionID, scope.ClusterConfig.ResourceGroup)
		if err != nil {
			klog.Infof("%s: failed to set provider ID: %v", machine.Name, err)
		} else {
			machine = modMachine
		}

		// If machine state is still provisioning state, we will return an error to keep the controllers
		// attempting to update status until it hits a more permanent state.
		if *vm.ProvisioningState != "Succeeded" {
			klog.Infof("%s: vm state not yet succeeded, returning an error to requeue", machine.Name)
			return &controllerError.RequeueAfterError{RequeueAfter: requeueAfterSeconds * time.Second}
		}
	}

	if err != nil {
		modMachine, updateConditionError := a.updateMachineProviderConditions(machine, providerconfig.MachineCreated, machineCreationFailedReason, err.Error())
		if updateConditionError != nil {
			klog.Errorf("%s: error updating machine conditions: %v", machine.Name, updateConditionError)
		} else {
			machine = modMachine
		}

		a.handleMachineError(modMachine, apierrors.CreateMachine("failed to reconcile machine %qs: %v", machine.Name, err), createEventAction)
		return &controllerError.RequeueAfterError{
			RequeueAfter: time.Minute,
		}
	}

	modMachine, updateConditionError := a.updateMachineProviderConditions(machine, providerconfig.MachineCreated, machineCreationSucceedReason, machineCreationSucceedMessage)
	if updateConditionError != nil {
		klog.Errorf("%s: error updating machine conditions: %v", machine.Name, updateConditionError)
	} else {
		machine = modMachine
	}

	a.eventRecorder.Eventf(machine, corev1.EventTypeNormal, "Created", "Created machine %q", machine.Name)

	return nil
}

// Delete deletes a machine and is invoked by the Machine Controller.
func (a *Actuator) Delete(ctx context.Context, cluster *clusterv1.Cluster, machine *machinev1.Machine) error {
	klog.Infof("Deleting machine %v", machine.Name)

	modMachine, err := a.setDeletingState(ctx, machine)
	if err != nil {
		klog.Errorf("unable to set machine deleting state: %v", err)
	} else {
		machine = modMachine
	}

	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{
		Machine:    machine,
		Cluster:    nil,
		Client:     a.client,
		CoreClient: a.coreClient,
	})
	if err != nil {
		return a.handleMachineError(machine, apierrors.DeleteMachine("failed to create machine %q scope: %v", machine.Name, err), deleteEventAction)
	}

	err = a.reconcilerBuilder(scope).Delete(context.Background())
	if err != nil {
		a.handleMachineError(machine, apierrors.DeleteMachine("failed to delete machine %q: %v", machine.Name, err), deleteEventAction)
		return &controllerError.RequeueAfterError{
			RequeueAfter: time.Minute,
		}
	}

	a.eventRecorder.Eventf(machine, corev1.EventTypeNormal, "Deleted", "Deleted machine %q", machine.Name)

	return nil
}

// Update updates a machine and is invoked by the Machine Controller.
// If the Update attempts to mutate any immutable state, the method will error
// and no updates will be performed.
func (a *Actuator) Update(ctx context.Context, cluster *clusterv1.Cluster, machine *machinev1.Machine) error {
	klog.Infof("Updating machine %v", machine.Name)

	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{
		Machine:    machine,
		Cluster:    nil,
		Client:     a.client,
		CoreClient: a.coreClient,
	})
	if err != nil {
		return a.handleMachineError(machine, apierrors.UpdateMachine("failed to create machine %q scope: %v", machine.Name, err), updateEventAction)
	}

	vm, err := a.reconcilerBuilder(scope).Update(context.Background())
	if vm != nil {
		modMachine, err := a.setMachineCloudProviderSpecifics(machine, *vm)
		if err != nil {
			klog.Errorf("%s: error updating machine cloud provider specifics: %v", machine.Name, err)
		} else {
			machine = modMachine
		}

		modMachine, err = a.updateStatus(ctx, machine, a.reconcilerBuilder(scope), *vm)
		if err != nil {
			klog.Infof("%s: failed to update status: %v", machine.Name, err)
			return &controllerError.RequeueAfterError{RequeueAfter: requeueAfterSeconds * time.Second}
		}
		machine = modMachine

		modMachine, err = a.updateProviderID(machine, *vm, scope.SubscriptionID, scope.ClusterConfig.ResourceGroup)
		if err != nil {
			klog.Infof("%s: failed to set provider ID: %v", machine.Name, err)
		} else {
			machine = modMachine
		}

		modMachine, updateConditionError := a.updateMachineProviderConditions(machine, providerconfig.MachineCreated, machineCreationSucceedReason, machineCreationSucceedMessage)
		if updateConditionError != nil {
			klog.Errorf("%s: error updating machine conditions: %v", machine.Name, updateConditionError)
		} else {
			machine = modMachine
		}
	}

	if err != nil {
		a.handleMachineError(machine, apierrors.UpdateMachine("failed to update machine %q: %v", machine.Name, err), updateEventAction)
		return &controllerError.RequeueAfterError{
			RequeueAfter: time.Minute,
		}
	}

	a.eventRecorder.Eventf(machine, corev1.EventTypeNormal, "Updated", "Updated machine %q", machine.Name)

	return nil
}

// Exists test for the existence of a machine and is invoked by the Machine Controller
func (a *Actuator) Exists(ctx context.Context, cluster *clusterv1.Cluster, machine *machinev1.Machine) (bool, error) {
	klog.Infof("Checking if machine %v exists", machine.Name)

	scope, err := actuators.NewMachineScope(actuators.MachineScopeParams{
		Machine:    machine,
		Cluster:    nil,
		Client:     a.client,
		CoreClient: a.coreClient,
	})
	if err != nil {
		return false, errors.Errorf("failed to create scope: %+v", err)
	}

	isExists, err := a.reconcilerBuilder(scope).Exists(context.Background())
	if err != nil {
		klog.Errorf("failed to check machine %s exists: %v", machine.Name, err)
	}

	return isExists, err
}

func (a *Actuator) updateMachineProviderConditions(machine *machinev1.Machine, conditionType providerconfig.AzureMachineProviderConditionType, reason string, msg string) (*machinev1.Machine, error) {
	klog.Infof("%s: updating machine conditions", machine.Name)

	azureStatus := &providerconfig.AzureMachineProviderStatus{}
	if err := a.codec.DecodeProviderStatus(machine.Status.ProviderStatus, azureStatus); err != nil {
		klog.Errorf("%s: error decoding machine provider status: %v", machine.Name, err)
		return nil, err
	}

	azureStatus.Conditions = setMachineProviderCondition(azureStatus.Conditions, providerconfig.AzureMachineProviderCondition{
		Type:    conditionType,
		Status:  corev1.ConditionTrue,
		Reason:  reason,
		Message: msg,
	})

	modMachine, err := a.updateMachineStatus(machine, azureStatus, nil)
	if err != nil {
		return nil, err
	}

	return modMachine, nil
}

func (a *Actuator) updateMachineStatus(machine *machinev1.Machine, azureStatus *providerconfig.AzureMachineProviderStatus, networkAddresses []corev1.NodeAddress) (*machinev1.Machine, error) {
	azureStatusRaw, err := a.codec.EncodeProviderStatus(azureStatus)
	if err != nil {
		klog.Errorf("%s: error encoding AWS provider status: %v", machine.Name, err)
		return nil, err
	}

	machineCopy := machine.DeepCopy()
	machineCopy.Status.ProviderStatus = azureStatusRaw
	if networkAddresses != nil {
		machineCopy.Status.Addresses = networkAddresses
	}

	oldAWSStatus := &providerconfig.AzureMachineProviderStatus{}
	if err := a.codec.DecodeProviderStatus(machine.Status.ProviderStatus, oldAWSStatus); err != nil {
		klog.Errorf("%s: error updating machine status: %v", machine.Name, err)
		return nil, err
	}

	if !equality.Semantic.DeepEqual(azureStatus, oldAWSStatus) || !equality.Semantic.DeepEqual(machine.Status.Addresses, machineCopy.Status.Addresses) {
		klog.Infof("%s: machine status has changed, updating", machine.Name)
		time := metav1.Now()
		machineCopy.Status.LastUpdated = &time

		if err := a.coreClient.Status().Update(context.Background(), machineCopy); err != nil {
			klog.Errorf("%s: error updating machine status: %v", machine.Name, err)
			return nil, err
		}
		return machineCopy, nil
	}

	klog.Infof("%s: status unchanged", machine.Name)
	return machine, nil
}

// providerConfigFromMachine gets the machine provider config MachineSetSpec from the
// specified cluster-api MachineSpec.
func providerConfigFromMachine(machine *machinev1.Machine, codec *providerconfig.AzureProviderConfigCodec) (*providerconfig.AzureMachineProviderSpec, error) {
	if machine.Spec.ProviderSpec.Value == nil {
		return nil, fmt.Errorf("unable to find machine provider config: Spec.ProviderSpec.Value is not set")
	}

	var config providerconfig.AzureMachineProviderSpec
	if err := codec.DecodeProviderSpec(&machine.Spec.ProviderSpec, &config); err != nil {
		return nil, err
	}
	return &config, nil
}

func (a *Actuator) setMachineCloudProviderSpecifics(machine *machinev1.Machine, vm compute.VirtualMachine) (*machinev1.Machine, error) {
	machineCopy := machine.DeepCopy()

	if machineCopy.Labels == nil {
		machineCopy.Labels = make(map[string]string)
	}

	if machineCopy.Annotations == nil {
		machineCopy.Annotations = make(map[string]string)
	}

	machineCopy.Annotations[MachineInstanceStateAnnotationName] = string(getVMState(vm))

	if vm.HardwareProfile != nil {
		machineCopy.Labels[MachineInstanceTypeLabelName] = string(vm.HardwareProfile.VMSize)
	}
	if vm.Location != nil {
		machineCopy.Labels[MachineRegionLabelName] = *vm.Location
	}
	if vm.Zones != nil {
		machineCopy.Labels[MachineAZLabelName] = strings.Join(*vm.Zones, ",")
	}

	if err := a.coreClient.Update(context.Background(), machineCopy); err != nil {
		return nil, fmt.Errorf("%s: error updating machine spec: %v", machine.Name, err)
	}

	return machineCopy, nil
}

// updateStatus calculates the new machine status, checks if anything has changed, and updates if so.
func (a *Actuator) updateStatus(ctx context.Context, machine *machinev1.Machine, reconciler *Reconciler, vm compute.VirtualMachine) (*machinev1.Machine, error) {
	klog.Infof("%s: Updating status", machine.Name)

	azureStatus := &providerconfig.AzureMachineProviderStatus{}
	if err := a.codec.DecodeProviderStatus(machine.Status.ProviderStatus, azureStatus); err != nil {
		klog.Errorf("%s: Error decoding machine provider status: %v", machine.Name, err)
		return nil, err
	}

	networkAddresses, err := reconciler.getNetworkAddresses(ctx, vm)
	if err != nil {
		return nil, err
	}

	vmState := getVMState(vm)
	azureStatus.VMID = vm.ID
	azureStatus.VMState = &vmState

	klog.Infof("%s: finished calculating Azure status", machine.Name)

	modMachine, err := a.updateMachineStatus(machine, azureStatus, networkAddresses)
	if err != nil {
		return nil, err
	}

	return modMachine, nil
}

func (a *Actuator) setDeletingState(ctx context.Context, machine *machinev1.Machine) (*machinev1.Machine, error) {
	// Getting a vm object does not work here so let's assume
	// an instance is really being deleted
	azureStatus := &providerconfig.AzureMachineProviderStatus{}
	if err := a.codec.DecodeProviderStatus(machine.Status.ProviderStatus, azureStatus); err != nil {
		klog.Errorf("%s: Error decoding machine provider status: %v", machine.Name, err)
		return nil, err
	}

	azureStatus.VMState = &providerconfig.VMStateDeleting

	modMachine, err := a.updateMachineStatus(machine, azureStatus, nil)
	if err != nil {
		return nil, fmt.Errorf("%s: error updating machine status: %v", machine.Name, err)
	}

	if modMachine.Annotations == nil {
		modMachine.Annotations = make(map[string]string)
	}
	modMachine.Annotations[MachineInstanceStateAnnotationName] = string(providerconfig.VMStateDeleting)

	if err := a.coreClient.Update(ctx, modMachine); err != nil {
		return nil, fmt.Errorf("%s: error updating machine spec: %v", modMachine.Name, err)
	}

	return modMachine, nil
}

// updateProviderID adds providerID in the machine spec
func (a *Actuator) updateProviderID(machine *machinev1.Machine, vm compute.VirtualMachine, subscriptionID, resourceGroup string) (*machinev1.Machine, error) {
	machineCopy := machine.DeepCopy()

	// Set provider ID
	if vm.OsProfile != nil && vm.OsProfile.ComputerName != nil {
		providerID := azure.GenerateMachineProviderID(subscriptionID, resourceGroup, *vm.OsProfile.ComputerName)
		klog.Infof("%s: setting ProviderID %s", machine.Name, providerID)
		machineCopy.Spec.ProviderID = &providerID
	} else {
		klog.Warningf("Unable to set providerID, not able to get vm.OsProfile.ComputerName. Setting ProviderID to nil.")
		machineCopy.Spec.ProviderID = nil
	}

	if err := a.coreClient.Update(context.Background(), machineCopy); err != nil {
		return nil, fmt.Errorf("%s: error updating machine spec ProviderID: %v", machineCopy.Name, err)
	}

	return machineCopy, nil
}
