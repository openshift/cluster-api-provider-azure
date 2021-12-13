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
// Source: ../machinepoolmachine.go

// Package mock_scope is a generated GoMock package.
package mock_scope

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "k8s.io/api/core/v1"
)

// MocknodeGetter is a mock of nodeGetter interface.
type MocknodeGetter struct {
	ctrl     *gomock.Controller
	recorder *MocknodeGetterMockRecorder
}

// MocknodeGetterMockRecorder is the mock recorder for MocknodeGetter.
type MocknodeGetterMockRecorder struct {
	mock *MocknodeGetter
}

// NewMocknodeGetter creates a new mock instance.
func NewMocknodeGetter(ctrl *gomock.Controller) *MocknodeGetter {
	mock := &MocknodeGetter{ctrl: ctrl}
	mock.recorder = &MocknodeGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocknodeGetter) EXPECT() *MocknodeGetterMockRecorder {
	return m.recorder
}

// GetNodeByObjectReference mocks base method.
func (m *MocknodeGetter) GetNodeByObjectReference(ctx context.Context, nodeRef v1.ObjectReference) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeByObjectReference", ctx, nodeRef)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeByObjectReference indicates an expected call of GetNodeByObjectReference.
func (mr *MocknodeGetterMockRecorder) GetNodeByObjectReference(ctx, nodeRef interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeByObjectReference", reflect.TypeOf((*MocknodeGetter)(nil).GetNodeByObjectReference), ctx, nodeRef)
}

// GetNodeByProviderID mocks base method.
func (m *MocknodeGetter) GetNodeByProviderID(ctx context.Context, providerID string) (*v1.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeByProviderID", ctx, providerID)
	ret0, _ := ret[0].(*v1.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeByProviderID indicates an expected call of GetNodeByProviderID.
func (mr *MocknodeGetterMockRecorder) GetNodeByProviderID(ctx, providerID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeByProviderID", reflect.TypeOf((*MocknodeGetter)(nil).GetNodeByProviderID), ctx, providerID)
}
