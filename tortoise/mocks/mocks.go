// Code generated by MockGen. DO NOT EDIT.
// Source: ./interfaces.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/spacemeshos/go-spacemesh/common/types"
)

// MockblockValidityUpdater is a mock of blockValidityUpdater interface.
type MockblockValidityUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockblockValidityUpdaterMockRecorder
}

// MockblockValidityUpdaterMockRecorder is the mock recorder for MockblockValidityUpdater.
type MockblockValidityUpdaterMockRecorder struct {
	mock *MockblockValidityUpdater
}

// NewMockblockValidityUpdater creates a new mock instance.
func NewMockblockValidityUpdater(ctrl *gomock.Controller) *MockblockValidityUpdater {
	mock := &MockblockValidityUpdater{ctrl: ctrl}
	mock.recorder = &MockblockValidityUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockblockValidityUpdater) EXPECT() *MockblockValidityUpdaterMockRecorder {
	return m.recorder
}

// UpdateBlockValidity mocks base method.
func (m *MockblockValidityUpdater) UpdateBlockValidity(arg0 types.BlockID, arg1 types.LayerID, arg2 bool) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateBlockValidity", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateBlockValidity indicates an expected call of UpdateBlockValidity.
func (mr *MockblockValidityUpdaterMockRecorder) UpdateBlockValidity(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateBlockValidity", reflect.TypeOf((*MockblockValidityUpdater)(nil).UpdateBlockValidity), arg0, arg1, arg2)
}
