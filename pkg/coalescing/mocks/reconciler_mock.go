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
// Source: sigs.k8s.io/controller-runtime/pkg/reconcile (interfaces: Reconciler)
//
// Generated by this command:
//
//	mockgen -destination reconciler_mock.go -package mock_coalescing sigs.k8s.io/controller-runtime/pkg/reconcile Reconciler
//

// Package mock_coalescing is a generated GoMock package.
package mock_coalescing

import (
	context "context"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
	reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// MockReconciler is a mock of Reconciler interface.
type MockReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockReconcilerMockRecorder
}

// MockReconcilerMockRecorder is the mock recorder for MockReconciler.
type MockReconcilerMockRecorder struct {
	mock *MockReconciler
}

// NewMockReconciler creates a new mock instance.
func NewMockReconciler(ctrl *gomock.Controller) *MockReconciler {
	mock := &MockReconciler{ctrl: ctrl}
	mock.recorder = &MockReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReconciler) EXPECT() *MockReconcilerMockRecorder {
	return m.recorder
}

// Reconcile mocks base method.
func (m *MockReconciler) Reconcile(arg0 context.Context, arg1 reconcile.Request) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reconcile", arg0, arg1)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reconcile indicates an expected call of Reconcile.
func (mr *MockReconcilerMockRecorder) Reconcile(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reconcile", reflect.TypeOf((*MockReconciler)(nil).Reconcile), arg0, arg1)
}
