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
	"encoding/base64"
	"fmt"
	"strings"
	"testing"

	"github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2018-10-01/compute"
	"github.com/Azure/go-autorest/autorest/to"
	"github.com/golang/mock/gomock"
	clusterapis "github.com/openshift/cluster-api/pkg/apis"
	machinev1 "github.com/openshift/cluster-api/pkg/apis/machine/v1beta1"
	"github.com/openshift/cluster-api/pkg/client/clientset_generated/clientset/fake"
	"github.com/openshift/cluster-api/pkg/controller/machine"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/tools/record"
	"k8s.io/utils/pointer"
	machineproviderv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	providerspecv1 "sigs.k8s.io/cluster-api-provider-azure/pkg/apis/azureprovider/v1beta1"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/actuators"
	mock_azure "sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/mock"
	"sigs.k8s.io/cluster-api-provider-azure/pkg/cloud/azure/services/virtualmachines"
	"sigs.k8s.io/controller-runtime/pkg/client"
	controllerfake "sigs.k8s.io/controller-runtime/pkg/client/fake"
)

var (
	_ machine.Actuator = (*Actuator)(nil)
)

func newFakeScope(t *testing.T) *actuators.Scope {
	return &actuators.Scope{
		ClusterName:          "cluster-test",
		ResourceGroup:        "dummyResourceGroup",
		NetworkResourceGroup: "dummyResourceGroup",
		Location:             "dummyLocation",
	}
}

func newFakeReconciler(t *testing.T, client client.Client, machine *machinev1.Machine, machineConfig *machineproviderv1.AzureMachineProviderSpec) *Reconciler {
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeVMSuccessSvc := &FakeVMService{
		Name:              "machine-test",
		ID:                "machine-test-ID",
		ProvisioningState: "Succeeded",
	}
	return &Reconciler{
		scope:                 newFakeScope(t),
		client:                client,
		machine:               machine,
		machineConfig:         machineConfig,
		availabilityZonesSvc:  fakeSuccessSvc,
		networkInterfacesSvc:  fakeSuccessSvc,
		virtualMachinesSvc:    fakeVMSuccessSvc,
		virtualMachinesExtSvc: fakeSuccessSvc,
		disksSvc:              fakeSuccessSvc,
		publicIPSvc:           fakeSuccessSvc,
	}
}

func newFakeReconcilerWithScope(t *testing.T, scope *actuators.Scope, client client.Client, machine *machinev1.Machine, machineConfig *machineproviderv1.AzureMachineProviderSpec) *Reconciler {
	fakeSuccessSvc := &azure.FakeSuccessService{}
	fakeVMSuccessSvc := &FakeVMService{
		Name:              "machine-test",
		ID:                "machine-test-ID",
		ProvisioningState: "Succeeded",
	}
	return &Reconciler{
		scope:                 scope,
		client:                client,
		machine:               machine,
		machineConfig:         machineConfig,
		availabilityZonesSvc:  fakeSuccessSvc,
		networkInterfacesSvc:  fakeSuccessSvc,
		virtualMachinesSvc:    fakeVMSuccessSvc,
		virtualMachinesExtSvc: fakeSuccessSvc,
	}
}

// FakeVMService generic vm service
type FakeVMService struct {
	Name                    string
	ID                      string
	ProvisioningState       string
	GetCallCount            int
	CreateOrUpdateCallCount int
	DeleteCallCount         int
}

// Get returns fake success.
func (s *FakeVMService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	s.GetCallCount++
	return compute.VirtualMachine{
		ID:   to.StringPtr(s.ID),
		Name: to.StringPtr(s.Name),
		VirtualMachineProperties: &compute.VirtualMachineProperties{
			ProvisioningState: to.StringPtr(s.ProvisioningState),
			HardwareProfile: &compute.HardwareProfile{
				VMSize: compute.VirtualMachineSizeTypesStandardB2ms,
			},
		},
	}, nil
}

// CreateOrUpdate returns fake success.
func (s *FakeVMService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	s.CreateOrUpdateCallCount++
	return nil
}

// Delete returns fake success.
func (s *FakeVMService) Delete(ctx context.Context, spec azure.Spec) error {
	s.DeleteCallCount++
	return nil
}

// FakeVMService generic vm service
type FakeCountService struct {
	GetCallCount            int
	CreateOrUpdateCallCount int
	DeleteCallCount         int
}

// Get returns fake success.
func (s *FakeCountService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	s.GetCallCount++
	return nil, nil
}

// CreateOrUpdate returns fake success.
func (s *FakeCountService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	s.CreateOrUpdateCallCount++
	return nil
}

// Delete returns fake success.
func (s *FakeCountService) Delete(ctx context.Context, spec azure.Spec) error {
	s.DeleteCallCount++
	return nil
}

func TestReconcilerSuccess(t *testing.T) {
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	fakeReconciler := newFakeReconciler(t, controllerfake.NewFakeClient(), machine, stubProviderConfig())

	if _, err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}

	if _, err := fakeReconciler.Update(context.Background()); err != nil {
		t.Errorf("failed to update machine: %+v", err)
	}

	if _, err := fakeReconciler.Exists(context.Background()); err != nil {
		t.Errorf("failed to check if machine exists: %+v", err)
	}

	if err := fakeReconciler.Delete(context.Background()); err != nil {
		t.Errorf("failed to delete machine: %+v", err)
	}
}

func TestReconcileFailure(t *testing.T) {
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	fakeFailureSvc := &azure.FakeFailureService{}
	fakeReconciler := newFakeReconciler(t, controllerfake.NewFakeClient(), machine, stubProviderConfig())
	fakeReconciler.networkInterfacesSvc = fakeFailureSvc
	fakeReconciler.virtualMachinesSvc = fakeFailureSvc
	fakeReconciler.virtualMachinesExtSvc = fakeFailureSvc

	if _, err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}

	if _, err := fakeReconciler.Update(context.Background()); err == nil {
		t.Errorf("expected update to fail")
	}

	if _, err := fakeReconciler.Exists(context.Background()); err == nil {
		t.Errorf("expected exists to fail")
	}

	if err := fakeReconciler.Delete(context.Background()); err == nil {
		t.Errorf("expected delete to fail")
	}
}

func TestReconcileVMFailedState(t *testing.T) {
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	fakeReconciler := newFakeReconciler(t, controllerfake.NewFakeClient(), machine, stubProviderConfig())
	fakeVMService := &FakeVMService{
		Name:              "machine-test",
		ID:                "machine-test-ID",
		ProvisioningState: "Failed",
	}
	fakeReconciler.virtualMachinesSvc = fakeVMService
	fakeDiskService := &FakeCountService{}
	fakeReconciler.disksSvc = fakeDiskService
	fakeNicService := &FakeCountService{}
	fakeReconciler.networkInterfacesSvc = fakeNicService

	if _, err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}

	if fakeVMService.GetCallCount != 1 {
		t.Errorf("expected get to be called just once")
	}

	if fakeVMService.DeleteCallCount != 1 {
		t.Errorf("expected delete to be called just once")
	}

	if fakeDiskService.DeleteCallCount != 1 {
		t.Errorf("expected disk delete to be called just once")
	}

	if fakeNicService.DeleteCallCount != 1 {
		t.Errorf("expected nic delete to be called just once")
	}

	if fakeVMService.CreateOrUpdateCallCount != 0 {
		t.Errorf("expected createorupdate not to be called")
	}
}

func TestReconcileVMSuceededState(t *testing.T) {
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	fakeReconciler := newFakeReconciler(t, controllerfake.NewFakeClient(), machine, stubProviderConfig())
	fakeVMService := &FakeVMService{
		Name:              "machine-test",
		ID:                "machine-test-ID",
		ProvisioningState: "Succeeded",
	}
	fakeReconciler.virtualMachinesSvc = fakeVMService

	if _, err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}

	if fakeVMService.GetCallCount != 1 {
		t.Errorf("expected get to be called just once")
	}

	if fakeVMService.DeleteCallCount != 0 {
		t.Errorf("expected delete not to be called")
	}

	if fakeVMService.CreateOrUpdateCallCount != 0 {
		t.Errorf("expected createorupdate not to be called")
	}
}

// FakeVMCheckZonesService generic fake vm zone service
type FakeVMCheckZonesService struct {
	checkZones []string
}

// Get returns fake success.
func (s *FakeVMCheckZonesService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	return nil, errors.New("vm not found")
}

// CreateOrUpdate returns fake success.
func (s *FakeVMCheckZonesService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	vmSpec, ok := spec.(*virtualmachines.Spec)
	if !ok {
		return errors.New("invalid vm specification")
	}

	if len(s.checkZones) <= 0 {
		return nil
	}
	for _, zone := range s.checkZones {
		if strings.EqualFold(zone, vmSpec.Zone) {
			return nil
		}
	}

	return errors.New("invalid input zone")
}

// Delete returns fake success.
func (s *FakeVMCheckZonesService) Delete(ctx context.Context, spec azure.Spec) error {
	return nil
}

// FakeAvailabilityZonesService generic fake availability zones
type FakeAvailabilityZonesService struct {
	zonesResponse []string
}

// Get returns fake success.
func (s *FakeAvailabilityZonesService) Get(ctx context.Context, spec azure.Spec) (interface{}, error) {
	return s.zonesResponse, nil
}

// CreateOrUpdate returns fake success.
func (s *FakeAvailabilityZonesService) CreateOrUpdate(ctx context.Context, spec azure.Spec) error {
	return nil
}

// Delete returns fake success.
func (s *FakeAvailabilityZonesService) Delete(ctx context.Context, spec azure.Spec) error {
	return nil
}

func TestAvailabilityZones(t *testing.T) {
	fakeScope := newFakeScope(t)
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}
	machineConfig := stubProviderConfig()
	machineConfig.UserDataSecret = nil

	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope, controllerfake.NewFakeClient(), machine, machineConfig)

	machineConfig.Zone = to.StringPtr("2")
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{
		checkZones: []string{"2"},
	}
	if _, err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}

	machineConfig.Zone = nil
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{
		checkZones: []string{""},
	}
	if _, err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("failed to create machine: %+v", err)
	}

	machineConfig.Zone = to.StringPtr("1")
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{
		checkZones: []string{"3"},
	}
	if _, err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail due to zone mismatch")
	}
}

func TestGetZone(t *testing.T) {
	testCases := []struct {
		inputZone *string
		expected  string
	}{
		{
			inputZone: nil,
			expected:  "",
		},
		{
			inputZone: pointer.StringPtr("3"),
			expected:  "3",
		},
	}

	for _, tc := range testCases {
		fakeScope := newFakeScope(t)
		machine, err := stubMachine()
		if err != nil {
			t.Fatal(err)
		}

		machineConfig := stubProviderConfig()
		machineConfig.Zone = tc.inputZone
		fakeReconciler := newFakeReconcilerWithScope(t, fakeScope, controllerfake.NewFakeClient(), machine, machineConfig)

		zones := []string{"1", "2", "3"}
		fakeReconciler.availabilityZonesSvc = &FakeAvailabilityZonesService{
			zonesResponse: zones,
		}

		got, err := fakeReconciler.getZone(context.Background())
		if err != nil {
			t.Fatalf("unexpected error getting zone")
		}

		if !strings.EqualFold(tc.expected, got) {
			t.Errorf("expected: %v, got: %v", tc.expected, got)
		}
	}
}

func TestCustomUserData(t *testing.T) {
	fakeScope := newFakeScope(t)
	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	machineConfig := stubProviderConfig()

	userDataSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      machineConfig.UserDataSecret.Name,
			Namespace: machine.Namespace,
		},
		Data: map[string][]byte{
			"userData": []byte("test-userdata"),
		},
	}

	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope, controllerfake.NewFakeClient(userDataSecret), machine, machineConfig)
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{}
	if _, err := fakeReconciler.Create(context.Background()); err != nil {
		t.Errorf("expected create to succeed %v", err)
	}

	userData, err := fakeReconciler.getCustomUserData()
	if err != nil {
		t.Errorf("expected get custom data to succeed %v", err)
	}

	if userData != base64.StdEncoding.EncodeToString([]byte("test-userdata")) {
		t.Errorf("expected userdata to be test-userdata, but found %s", userData)
	}
}

func TestCustomDataFailures(t *testing.T) {
	fakeScope := newFakeScope(t)
	userDataSecret := &corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "testCustomUserData",
			Namespace: "dummyNamespace",
		},
		Data: map[string][]byte{
			"userData": []byte("test-userdata"),
		},
	}

	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	fakeReconciler := newFakeReconcilerWithScope(t, fakeScope, controllerfake.NewFakeClient(userDataSecret), machine, stubProviderConfig())
	fakeReconciler.virtualMachinesSvc = &FakeVMCheckZonesService{}

	if _, err := fakeReconciler.Create(context.Background()); err == nil {
		t.Errorf("expected create to fail")
	}

	if _, err := fakeReconciler.getCustomUserData(); err == nil {
		t.Errorf("expected get custom data to fail")
	}

	userDataSecret.Data = map[string][]byte{
		"notUserData": []byte("test-notuserdata"),
	}

	if _, err := fakeReconciler.getCustomUserData(); err == nil {
		t.Errorf("expected get custom data to fail, due to missing userdata")
	}
}

func TestMachineEvents(t *testing.T) {
	if err := clusterapis.AddToScheme(scheme.Scheme); err != nil {
		t.Fatal(err)
	}

	machine, err := stubMachine()
	if err != nil {
		t.Fatal(err)
	}

	machinePc := stubProviderConfig()
	machinePc.Subnet = ""
	machinePc.Vnet = ""
	machinePc.VMSize = ""
	providerSpec, err := providerspecv1.EncodeMachineSpec(machinePc)
	if err != nil {
		t.Fatalf("EncodeMachineSpec failed: %v", err)
	}

	invalidMachine := machine.DeepCopy()
	invalidMachine.Spec.ProviderSpec = machinev1.ProviderSpec{Value: providerSpec}

	azureCredentialsSecret := stubAzureCredentialsSecret()
	invalidAzureCredentialsSecret := stubAzureCredentialsSecret()
	delete(invalidAzureCredentialsSecret.Data, "azure_client_id")

	cases := []struct {
		name       string
		machine    *machinev1.Machine
		credSecret *corev1.Secret
		error      string
		operation  func(actuator *Actuator, machine *machinev1.Machine)
		event      string
	}{
		{
			name:       "Create machine event failed (scope)",
			machine:    machine,
			credSecret: invalidAzureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Create(context.TODO(), nil, machine)
			},
			event: "Warning FailedCreate CreateError: azure-actuator-testing-machine: failed to create scope: failed to create scope: Azure client id default/azure-credentials-secret did not contain key azure_client_id",
		},
		{
			name:       "Create machine event failed (reconciler)",
			machine:    invalidMachine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Create(context.TODO(), nil, machine)
			},
			event: "Warning FailedCreate CreateError: failed to reconcile machine \"azure-actuator-testing-machine\"s: failed to create nic azure-actuator-testing-machine-nic for machine azure-actuator-testing-machine: MachineConfig vnet is missing on machine azure-actuator-testing-machine",
		},
		{
			name:       "Create machine event succeed",
			machine:    machine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Create(context.TODO(), nil, machine)
			},
			event: fmt.Sprintf("Normal Created Created machine %q", machine.Name),
		},
		{
			name:       "Update machine event failed (scope)",
			machine:    machine,
			credSecret: invalidAzureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Update(context.TODO(), nil, machine)
			},
			event: "Warning FailedUpdate UpdateError: azure-actuator-testing-machine: failed to create scope: failed to create scope: Azure client id default/azure-credentials-secret did not contain key azure_client_id",
		},
		{
			name:       "Update machine event failed (reconciler)",
			machine:    invalidMachine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Update(context.TODO(), nil, machine)
			},
			event: "Warning FailedUpdate UpdateError: failed to update machine \"azure-actuator-testing-machine\": found attempt to change immutable state",
		},
		{
			name:       "Update machine event succeed",
			machine:    machine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Update(context.TODO(), nil, machine)
			},
			event: fmt.Sprintf("Normal Updated Updated machine %q", machine.Name),
		},
		{
			name:       "Delete machine event failed (scope)",
			machine:    machine,
			credSecret: invalidAzureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Delete(context.TODO(), nil, machine)
			},
			event: "Warning FailedDelete DeleteError: azure-actuator-testing-machine: failed to create scope: failed to create scope: Azure client id default/azure-credentials-secret did not contain key azure_client_id",
		},
		{
			name:       "Delete machine event failed (reconciler)",
			machine:    invalidMachine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Delete(context.TODO(), nil, machine)
			},
			event: "Warning FailedDelete DeleteError: failed to delete machine \"azure-actuator-testing-machine\": MachineConfig vnet is missing on machine azure-actuator-testing-machine",
		},
		{
			name:       "Delete machine event succeed",
			machine:    machine,
			credSecret: azureCredentialsSecret,
			operation: func(actuator *Actuator, machine *machinev1.Machine) {
				actuator.Delete(context.TODO(), nil, machine)
			},
			event: fmt.Sprintf("Normal Deleted Deleted machine %q", machine.Name),
		},
	}

	codec, err := providerspecv1.NewCodec()
	if err != nil {
		t.Fatalf("Unable to create codec: %v", err)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			cs := controllerfake.NewFakeClient(tc.credSecret, tc.machine)

			mockCtrl := gomock.NewController(t)
			azSvc := mock_azure.NewMockService(mockCtrl)
			networkSvc := mock_azure.NewMockService(mockCtrl)
			vmSvc := mock_azure.NewMockService(mockCtrl)
			vmExtSvc := mock_azure.NewMockService(mockCtrl)
			pipSvc := mock_azure.NewMockService(mockCtrl)
			disksSvc := mock_azure.NewMockService(mockCtrl)

			eventsChannel := make(chan string, 1)

			machineActuator := NewActuator(ActuatorParams{
				Client:     fake.NewSimpleClientset(tc.machine).MachineV1beta1(),
				CoreClient: cs,
				Codec:      codec,
				ReconcilerBuilder: func(scope *actuators.Scope, client client.Client, machine *machinev1.Machine, machineConfig *providerspecv1.AzureMachineProviderSpec) *Reconciler {
					return &Reconciler{
						scope:                 scope,
						client:                client,
						machine:               machine,
						machineConfig:         machineConfig,
						availabilityZonesSvc:  azSvc,
						networkInterfacesSvc:  networkSvc,
						virtualMachinesSvc:    vmSvc,
						virtualMachinesExtSvc: vmExtSvc,
						publicIPSvc:           pipSvc,
						disksSvc:              disksSvc,
					}
				},
				// use fake recorder and store an event into one item long buffer for subsequent check
				EventRecorder: &record.FakeRecorder{
					Events: eventsChannel,
				},
			})

			networkSvc.EXPECT().CreateOrUpdate(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			vmSvc.EXPECT().Get(gomock.Any(), gomock.Any()).Return(compute.VirtualMachine{
				ID:   pointer.StringPtr("vm-id"),
				Name: pointer.StringPtr("vm-name"),
				VirtualMachineProperties: &compute.VirtualMachineProperties{
					ProvisioningState: pointer.StringPtr("Succeeded"),
					HardwareProfile: &compute.HardwareProfile{
						VMSize: compute.VirtualMachineSizeTypesStandardB2ms,
					},
					OsProfile: &compute.OSProfile{
						ComputerName: pointer.StringPtr("vm-name"),
					},
				},
			}, nil).AnyTimes()
			vmSvc.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			disksSvc.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			networkSvc.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()
			pipSvc.EXPECT().Delete(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

			tc.operation(machineActuator, tc.machine)

			select {
			case event := <-eventsChannel:
				if event != tc.event {
					t.Errorf("Expected %q event, got %q", tc.event, event)
				}
			default:
				t.Errorf("Expected %q event, got none", tc.event)
			}
		})
	}
}
