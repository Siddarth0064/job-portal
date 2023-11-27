// Code generated by MockGen. DO NOT EDIT.
// Source: cache.go
//
// Generated by this command:
//
//	mockgen -source=cache.go -destination=cache_mock.go -package=cache
//
// Package cache is a generated GoMock package.
package cache

import (
	context "context"
	model "job-portal-api/internal/models"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCaching is a mock of Caching interface.
type MockCaching struct {
	ctrl     *gomock.Controller
	recorder *MockCachingMockRecorder
}

// MockCachingMockRecorder is the mock recorder for MockCaching.
type MockCachingMockRecorder struct {
	mock *MockCaching
}

// NewMockCaching creates a new mock instance.
func NewMockCaching(ctrl *gomock.Controller) *MockCaching {
	mock := &MockCaching{ctrl: ctrl}
	mock.recorder = &MockCachingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCaching) EXPECT() *MockCachingMockRecorder {
	return m.recorder
}

// AddToCache mocks base method.
func (m *MockCaching) AddToCache(ctx context.Context, jid uint, jobData model.Job) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddToCache", ctx, jid, jobData)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddToCache indicates an expected call of AddToCache.
func (mr *MockCachingMockRecorder) AddToCache(ctx, jid, jobData any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddToCache", reflect.TypeOf((*MockCaching)(nil).AddToCache), ctx, jid, jobData)
}

// GetCacheData mocks base method.
func (m *MockCaching) GetCacheData(ctx context.Context, jid uint) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCacheData", ctx, jid)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCacheData indicates an expected call of GetCacheData.
func (mr *MockCachingMockRecorder) GetCacheData(ctx, jid any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCacheData", reflect.TypeOf((*MockCaching)(nil).GetCacheData), ctx, jid)
}
