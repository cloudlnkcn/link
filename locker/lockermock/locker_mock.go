// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Scalingo/link/locker (interfaces: Locker)

// Package lockermock is a generated GoMock package.
package lockermock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockLocker is a mock of Locker interface
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *MockLockerMockRecorder
}

// MockLockerMockRecorder is the mock recorder for MockLocker
type MockLockerMockRecorder struct {
	mock *MockLocker
}

// NewMockLocker creates a new mock instance
func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &MockLockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocker) EXPECT() *MockLockerMockRecorder {
	return m.recorder
}

// IsMaster mocks base method
func (m *MockLocker) IsMaster(arg0 context.Context) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsMaster", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// IsMaster indicates an expected call of IsMaster
func (mr *MockLockerMockRecorder) IsMaster(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsMaster", reflect.TypeOf((*MockLocker)(nil).IsMaster), arg0)
}

// Refresh mocks base method
func (m *MockLocker) Refresh(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Refresh indicates an expected call of Refresh
func (mr *MockLockerMockRecorder) Refresh(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockLocker)(nil).Refresh), arg0)
}

// Stop mocks base method
func (m *MockLocker) Stop(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (mr *MockLockerMockRecorder) Stop(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockLocker)(nil).Stop), arg0)
}

// Unlock mocks base method
func (m *MockLocker) Unlock(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unlock indicates an expected call of Unlock
func (mr *MockLockerMockRecorder) Unlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unlock", reflect.TypeOf((*MockLocker)(nil).Unlock), arg0)
}
