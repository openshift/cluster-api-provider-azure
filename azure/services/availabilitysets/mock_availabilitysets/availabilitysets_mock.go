/*
Copyright The Kubernetes Authors.

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

// Code generated by MockGen. DO NOT EDIT.
// Source: ../availabilitysets.go
//
// Generated by this command:
//
//	mockgen -destination availabilitysets_mock.go -package mock_availabilitysets -source ../availabilitysets.go AvailabilitySetScope
//

// Package mock_availabilitysets is a generated GoMock package.
package mock_availabilitysets

import (
	reflect "reflect"
	time "time"

	azcore "github.com/Azure/azure-sdk-for-go/sdk/azcore"
	gomock "go.uber.org/mock/gomock"
	v1beta1 "sigs.k8s.io/cluster-api-provider-azure/api/v1beta1"
	azure "sigs.k8s.io/cluster-api-provider-azure/azure"
	v1beta10 "sigs.k8s.io/cluster-api/api/v1beta1"
)

// MockAvailabilitySetScope is a mock of AvailabilitySetScope interface.
type MockAvailabilitySetScope struct {
	ctrl     *gomock.Controller
	recorder *MockAvailabilitySetScopeMockRecorder
}

// MockAvailabilitySetScopeMockRecorder is the mock recorder for MockAvailabilitySetScope.
type MockAvailabilitySetScopeMockRecorder struct {
	mock *MockAvailabilitySetScope
}

// NewMockAvailabilitySetScope creates a new mock instance.
func NewMockAvailabilitySetScope(ctrl *gomock.Controller) *MockAvailabilitySetScope {
	mock := &MockAvailabilitySetScope{ctrl: ctrl}
	mock.recorder = &MockAvailabilitySetScopeMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAvailabilitySetScope) EXPECT() *MockAvailabilitySetScopeMockRecorder {
	return m.recorder
}

// AdditionalTags mocks base method.
func (m *MockAvailabilitySetScope) AdditionalTags() v1beta1.Tags {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AdditionalTags")
	ret0, _ := ret[0].(v1beta1.Tags)
	return ret0
}

// AdditionalTags indicates an expected call of AdditionalTags.
func (mr *MockAvailabilitySetScopeMockRecorder) AdditionalTags() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AdditionalTags", reflect.TypeOf((*MockAvailabilitySetScope)(nil).AdditionalTags))
}

// AvailabilitySetEnabled mocks base method.
func (m *MockAvailabilitySetScope) AvailabilitySetEnabled() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetEnabled")
	ret0, _ := ret[0].(bool)
	return ret0
}

// AvailabilitySetEnabled indicates an expected call of AvailabilitySetEnabled.
func (mr *MockAvailabilitySetScopeMockRecorder) AvailabilitySetEnabled() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetEnabled", reflect.TypeOf((*MockAvailabilitySetScope)(nil).AvailabilitySetEnabled))
}

// AvailabilitySetSpec mocks base method.
func (m *MockAvailabilitySetScope) AvailabilitySetSpec() azure.ResourceSpecGetter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AvailabilitySetSpec")
	ret0, _ := ret[0].(azure.ResourceSpecGetter)
	return ret0
}

// AvailabilitySetSpec indicates an expected call of AvailabilitySetSpec.
func (mr *MockAvailabilitySetScopeMockRecorder) AvailabilitySetSpec() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AvailabilitySetSpec", reflect.TypeOf((*MockAvailabilitySetScope)(nil).AvailabilitySetSpec))
}

// BaseURI mocks base method.
func (m *MockAvailabilitySetScope) BaseURI() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BaseURI")
	ret0, _ := ret[0].(string)
	return ret0
}

// BaseURI indicates an expected call of BaseURI.
func (mr *MockAvailabilitySetScopeMockRecorder) BaseURI() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BaseURI", reflect.TypeOf((*MockAvailabilitySetScope)(nil).BaseURI))
}

// ClientID mocks base method.
func (m *MockAvailabilitySetScope) ClientID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientID indicates an expected call of ClientID.
func (mr *MockAvailabilitySetScopeMockRecorder) ClientID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientID", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ClientID))
}

// ClientSecret mocks base method.
func (m *MockAvailabilitySetScope) ClientSecret() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClientSecret")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClientSecret indicates an expected call of ClientSecret.
func (mr *MockAvailabilitySetScopeMockRecorder) ClientSecret() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClientSecret", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ClientSecret))
}

// CloudEnvironment mocks base method.
func (m *MockAvailabilitySetScope) CloudEnvironment() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudEnvironment")
	ret0, _ := ret[0].(string)
	return ret0
}

// CloudEnvironment indicates an expected call of CloudEnvironment.
func (mr *MockAvailabilitySetScopeMockRecorder) CloudEnvironment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudEnvironment", reflect.TypeOf((*MockAvailabilitySetScope)(nil).CloudEnvironment))
}

// CloudProviderConfigOverrides mocks base method.
func (m *MockAvailabilitySetScope) CloudProviderConfigOverrides() *v1beta1.CloudProviderConfigOverrides {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloudProviderConfigOverrides")
	ret0, _ := ret[0].(*v1beta1.CloudProviderConfigOverrides)
	return ret0
}

// CloudProviderConfigOverrides indicates an expected call of CloudProviderConfigOverrides.
func (mr *MockAvailabilitySetScopeMockRecorder) CloudProviderConfigOverrides() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloudProviderConfigOverrides", reflect.TypeOf((*MockAvailabilitySetScope)(nil).CloudProviderConfigOverrides))
}

// ClusterName mocks base method.
func (m *MockAvailabilitySetScope) ClusterName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClusterName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ClusterName indicates an expected call of ClusterName.
func (mr *MockAvailabilitySetScopeMockRecorder) ClusterName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClusterName", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ClusterName))
}

// DefaultedAzureCallTimeout mocks base method.
func (m *MockAvailabilitySetScope) DefaultedAzureCallTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureCallTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureCallTimeout indicates an expected call of DefaultedAzureCallTimeout.
func (mr *MockAvailabilitySetScopeMockRecorder) DefaultedAzureCallTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureCallTimeout", reflect.TypeOf((*MockAvailabilitySetScope)(nil).DefaultedAzureCallTimeout))
}

// DefaultedAzureServiceReconcileTimeout mocks base method.
func (m *MockAvailabilitySetScope) DefaultedAzureServiceReconcileTimeout() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedAzureServiceReconcileTimeout")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedAzureServiceReconcileTimeout indicates an expected call of DefaultedAzureServiceReconcileTimeout.
func (mr *MockAvailabilitySetScopeMockRecorder) DefaultedAzureServiceReconcileTimeout() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedAzureServiceReconcileTimeout", reflect.TypeOf((*MockAvailabilitySetScope)(nil).DefaultedAzureServiceReconcileTimeout))
}

// DefaultedReconcilerRequeue mocks base method.
func (m *MockAvailabilitySetScope) DefaultedReconcilerRequeue() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultedReconcilerRequeue")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// DefaultedReconcilerRequeue indicates an expected call of DefaultedReconcilerRequeue.
func (mr *MockAvailabilitySetScopeMockRecorder) DefaultedReconcilerRequeue() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultedReconcilerRequeue", reflect.TypeOf((*MockAvailabilitySetScope)(nil).DefaultedReconcilerRequeue))
}

// DeleteLongRunningOperationState mocks base method.
func (m *MockAvailabilitySetScope) DeleteLongRunningOperationState(arg0, arg1, arg2 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteLongRunningOperationState", arg0, arg1, arg2)
}

// DeleteLongRunningOperationState indicates an expected call of DeleteLongRunningOperationState.
func (mr *MockAvailabilitySetScopeMockRecorder) DeleteLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLongRunningOperationState", reflect.TypeOf((*MockAvailabilitySetScope)(nil).DeleteLongRunningOperationState), arg0, arg1, arg2)
}

// ExtendedLocation mocks base method.
func (m *MockAvailabilitySetScope) ExtendedLocation() *v1beta1.ExtendedLocationSpec {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocation")
	ret0, _ := ret[0].(*v1beta1.ExtendedLocationSpec)
	return ret0
}

// ExtendedLocation indicates an expected call of ExtendedLocation.
func (mr *MockAvailabilitySetScopeMockRecorder) ExtendedLocation() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocation", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ExtendedLocation))
}

// ExtendedLocationName mocks base method.
func (m *MockAvailabilitySetScope) ExtendedLocationName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationName")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationName indicates an expected call of ExtendedLocationName.
func (mr *MockAvailabilitySetScopeMockRecorder) ExtendedLocationName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationName", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ExtendedLocationName))
}

// ExtendedLocationType mocks base method.
func (m *MockAvailabilitySetScope) ExtendedLocationType() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExtendedLocationType")
	ret0, _ := ret[0].(string)
	return ret0
}

// ExtendedLocationType indicates an expected call of ExtendedLocationType.
func (mr *MockAvailabilitySetScopeMockRecorder) ExtendedLocationType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendedLocationType", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ExtendedLocationType))
}

// FailureDomains mocks base method.
func (m *MockAvailabilitySetScope) FailureDomains() []*string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FailureDomains")
	ret0, _ := ret[0].([]*string)
	return ret0
}

// FailureDomains indicates an expected call of FailureDomains.
func (mr *MockAvailabilitySetScopeMockRecorder) FailureDomains() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FailureDomains", reflect.TypeOf((*MockAvailabilitySetScope)(nil).FailureDomains))
}

// GetLongRunningOperationState mocks base method.
func (m *MockAvailabilitySetScope) GetLongRunningOperationState(arg0, arg1, arg2 string) *v1beta1.Future {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLongRunningOperationState", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1beta1.Future)
	return ret0
}

// GetLongRunningOperationState indicates an expected call of GetLongRunningOperationState.
func (mr *MockAvailabilitySetScopeMockRecorder) GetLongRunningOperationState(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLongRunningOperationState", reflect.TypeOf((*MockAvailabilitySetScope)(nil).GetLongRunningOperationState), arg0, arg1, arg2)
}

// HashKey mocks base method.
func (m *MockAvailabilitySetScope) HashKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HashKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// HashKey indicates an expected call of HashKey.
func (mr *MockAvailabilitySetScopeMockRecorder) HashKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HashKey", reflect.TypeOf((*MockAvailabilitySetScope)(nil).HashKey))
}

// Location mocks base method.
func (m *MockAvailabilitySetScope) Location() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Location")
	ret0, _ := ret[0].(string)
	return ret0
}

// Location indicates an expected call of Location.
func (mr *MockAvailabilitySetScopeMockRecorder) Location() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Location", reflect.TypeOf((*MockAvailabilitySetScope)(nil).Location))
}

// NodeResourceGroup mocks base method.
func (m *MockAvailabilitySetScope) NodeResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// NodeResourceGroup indicates an expected call of NodeResourceGroup.
func (mr *MockAvailabilitySetScopeMockRecorder) NodeResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeResourceGroup", reflect.TypeOf((*MockAvailabilitySetScope)(nil).NodeResourceGroup))
}

// ResourceGroup mocks base method.
func (m *MockAvailabilitySetScope) ResourceGroup() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ResourceGroup")
	ret0, _ := ret[0].(string)
	return ret0
}

// ResourceGroup indicates an expected call of ResourceGroup.
func (mr *MockAvailabilitySetScopeMockRecorder) ResourceGroup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ResourceGroup", reflect.TypeOf((*MockAvailabilitySetScope)(nil).ResourceGroup))
}

// SetLongRunningOperationState mocks base method.
func (m *MockAvailabilitySetScope) SetLongRunningOperationState(arg0 *v1beta1.Future) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetLongRunningOperationState", arg0)
}

// SetLongRunningOperationState indicates an expected call of SetLongRunningOperationState.
func (mr *MockAvailabilitySetScopeMockRecorder) SetLongRunningOperationState(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLongRunningOperationState", reflect.TypeOf((*MockAvailabilitySetScope)(nil).SetLongRunningOperationState), arg0)
}

// SubscriptionID mocks base method.
func (m *MockAvailabilitySetScope) SubscriptionID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubscriptionID")
	ret0, _ := ret[0].(string)
	return ret0
}

// SubscriptionID indicates an expected call of SubscriptionID.
func (mr *MockAvailabilitySetScopeMockRecorder) SubscriptionID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubscriptionID", reflect.TypeOf((*MockAvailabilitySetScope)(nil).SubscriptionID))
}

// TenantID mocks base method.
func (m *MockAvailabilitySetScope) TenantID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TenantID")
	ret0, _ := ret[0].(string)
	return ret0
}

// TenantID indicates an expected call of TenantID.
func (mr *MockAvailabilitySetScopeMockRecorder) TenantID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TenantID", reflect.TypeOf((*MockAvailabilitySetScope)(nil).TenantID))
}

// Token mocks base method.
func (m *MockAvailabilitySetScope) Token() azcore.TokenCredential {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Token")
	ret0, _ := ret[0].(azcore.TokenCredential)
	return ret0
}

// Token indicates an expected call of Token.
func (mr *MockAvailabilitySetScopeMockRecorder) Token() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Token", reflect.TypeOf((*MockAvailabilitySetScope)(nil).Token))
}

// UpdateDeleteStatus mocks base method.
func (m *MockAvailabilitySetScope) UpdateDeleteStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateDeleteStatus", arg0, arg1, arg2)
}

// UpdateDeleteStatus indicates an expected call of UpdateDeleteStatus.
func (mr *MockAvailabilitySetScopeMockRecorder) UpdateDeleteStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateDeleteStatus", reflect.TypeOf((*MockAvailabilitySetScope)(nil).UpdateDeleteStatus), arg0, arg1, arg2)
}

// UpdatePatchStatus mocks base method.
func (m *MockAvailabilitySetScope) UpdatePatchStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePatchStatus", arg0, arg1, arg2)
}

// UpdatePatchStatus indicates an expected call of UpdatePatchStatus.
func (mr *MockAvailabilitySetScopeMockRecorder) UpdatePatchStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePatchStatus", reflect.TypeOf((*MockAvailabilitySetScope)(nil).UpdatePatchStatus), arg0, arg1, arg2)
}

// UpdatePutStatus mocks base method.
func (m *MockAvailabilitySetScope) UpdatePutStatus(arg0 v1beta10.ConditionType, arg1 string, arg2 error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatePutStatus", arg0, arg1, arg2)
}

// UpdatePutStatus indicates an expected call of UpdatePutStatus.
func (mr *MockAvailabilitySetScopeMockRecorder) UpdatePutStatus(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePutStatus", reflect.TypeOf((*MockAvailabilitySetScope)(nil).UpdatePutStatus), arg0, arg1, arg2)
}
