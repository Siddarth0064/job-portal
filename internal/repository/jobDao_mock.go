// Code generated by MockGen. DO NOT EDIT.
// Source: jobDao.go
//
// Generated by this command:
//
//	mockgen -source=jobDao.go -destination=jobDao_mock.go -package=repository
//
// Package repository is a generated GoMock package.
package repository

import (
	model "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCompany is a mock of Company interface.
type MockCompany struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyMockRecorder
}

// MockCompanyMockRecorder is the mock recorder for MockCompany.
type MockCompanyMockRecorder struct {
	mock *MockCompany
}

// NewMockCompany creates a new mock instance.
func NewMockCompany(ctrl *gomock.Controller) *MockCompany {
	mock := &MockCompany{ctrl: ctrl}
	mock.recorder = &MockCompanyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompany) EXPECT() *MockCompanyMockRecorder {
	return m.recorder
}

// CreateCompany mocks base method.
func (m *MockCompany) CreateCompany(arg0 model.Company) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCompany", arg0)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateCompany indicates an expected call of CreateCompany.
func (mr *MockCompanyMockRecorder) CreateCompany(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCompany", reflect.TypeOf((*MockCompany)(nil).CreateCompany), arg0)
}

// FetchJobData mocks base method.
func (m *MockCompany) FetchJobData(jid uint64) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchJobData", jid)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchJobData indicates an expected call of FetchJobData.
func (mr *MockCompanyMockRecorder) FetchJobData(jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchJobData", reflect.TypeOf((*MockCompany)(nil).FetchJobData), jid)
}

// GetAllCompany mocks base method.
func (m *MockCompany) GetAllCompany() ([]model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllCompany")
	ret0, _ := ret[0].([]model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllCompany indicates an expected call of GetAllCompany.
func (mr *MockCompanyMockRecorder) GetAllCompany() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllCompany", reflect.TypeOf((*MockCompany)(nil).GetAllCompany))
}

// GetAllJobs mocks base method.
func (m *MockCompany) GetAllJobs() ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllJobs")
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllJobs indicates an expected call of GetAllJobs.
func (mr *MockCompanyMockRecorder) GetAllJobs() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllJobs", reflect.TypeOf((*MockCompany)(nil).GetAllJobs))
}

// GetCompany mocks base method.
func (m *MockCompany) GetCompany(id int64) (model.Company, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCompany", id)
	ret0, _ := ret[0].(model.Company)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCompany indicates an expected call of GetCompany.
func (mr *MockCompanyMockRecorder) GetCompany(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCompany", reflect.TypeOf((*MockCompany)(nil).GetCompany), id)
}

// GetJobs mocks base method.
func (m *MockCompany) GetJobs(id int) ([]model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetJobs", id)
	ret0, _ := ret[0].([]model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetJobs indicates an expected call of GetJobs.
func (mr *MockCompanyMockRecorder) GetJobs(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetJobs", reflect.TypeOf((*MockCompany)(nil).GetJobs), id)
}

// GetTheJobData mocks base method.
func (m *MockCompany) GetTheJobData(jobid uint) (model.Job, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheJobData", jobid)
	ret0, _ := ret[0].(model.Job)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTheJobData indicates an expected call of GetTheJobData.
func (mr *MockCompanyMockRecorder) GetTheJobData(jobid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheJobData", reflect.TypeOf((*MockCompany)(nil).GetTheJobData), jobid)
}

// PostJob mocks base method.
func (m *MockCompany) PostJob(nj model.Job) (model.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PostJob", nj)
	ret0, _ := ret[0].(model.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// PostJob indicates an expected call of PostJob.
func (mr *MockCompanyMockRecorder) PostJob(nj any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PostJob", reflect.TypeOf((*MockCompany)(nil).PostJob), nj)
}
