// Code generated by MockGen. DO NOT EDIT.
// Source: storage.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entity "coinbase-client/internal/entity"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSocketStorage is a mock of SocketStorage interface.
type MockSocketStorage struct {
	ctrl     *gomock.Controller
	recorder *MockSocketStorageMockRecorder
}

// MockSocketStorageMockRecorder is the mock recorder for MockSocketStorage.
type MockSocketStorageMockRecorder struct {
	mock *MockSocketStorage
}

// NewMockSocketStorage creates a new mock instance.
func NewMockSocketStorage(ctrl *gomock.Controller) *MockSocketStorage {
	mock := &MockSocketStorage{ctrl: ctrl}
	mock.recorder = &MockSocketStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocketStorage) EXPECT() *MockSocketStorageMockRecorder {
	return m.recorder
}

// SaveTicker mocks base method.
func (m *MockSocketStorage) SaveTicker(ticker entity.Ticker) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveTicker", ticker)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveTicker indicates an expected call of SaveTicker.
func (mr *MockSocketStorageMockRecorder) SaveTicker(ticker interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveTicker", reflect.TypeOf((*MockSocketStorage)(nil).SaveTicker), ticker)
}
