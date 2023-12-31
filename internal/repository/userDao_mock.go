// Code generated by MockGen. DO NOT EDIT.
// Source: userDao.go
//
// Generated by this command:
//
//	mockgen -source=userDao.go -destination=userDao_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	model "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockUsers is a mock of Users interface.
type MockUsers struct {
	ctrl     *gomock.Controller
	recorder *MockUsersMockRecorder
}

// MockUsersMockRecorder is the mock recorder for MockUsers.
type MockUsersMockRecorder struct {
	mock *MockUsers
}

// NewMockUsers creates a new mock instance.
func NewMockUsers(ctrl *gomock.Controller) *MockUsers {
	mock := &MockUsers{ctrl: ctrl}
	mock.recorder = &MockUsersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsers) EXPECT() *MockUsersMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockUsers) CreateUser(arg0 model.User) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockUsersMockRecorder) CreateUser(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUsers)(nil).CreateUser), arg0)
}

// FetchUserByEmail mocks base method.
func (m *MockUsers) FetchUserByEmail(arg0 string) (model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchUserByEmail", arg0)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchUserByEmail indicates an expected call of FetchUserByEmail.
func (mr *MockUsersMockRecorder) FetchUserByEmail(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchUserByEmail", reflect.TypeOf((*MockUsers)(nil).FetchUserByEmail), arg0)
}
