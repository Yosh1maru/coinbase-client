// Code generated by MockGen. DO NOT EDIT.
// Source: source.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockSocketSource is a mock of SocketSource interface.
type MockSocketSource struct {
	ctrl     *gomock.Controller
	recorder *MockSocketSourceMockRecorder
}

// MockSocketSourceMockRecorder is the mock recorder for MockSocketSource.
type MockSocketSourceMockRecorder struct {
	mock *MockSocketSource
}

// NewMockSocketSource creates a new mock instance.
func NewMockSocketSource(ctrl *gomock.Controller) *MockSocketSource {
	mock := &MockSocketSource{ctrl: ctrl}
	mock.recorder = &MockSocketSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSocketSource) EXPECT() *MockSocketSourceMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockSocketSource) Read(ctx context.Context) (chan []byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", ctx)
	ret0, _ := ret[0].(chan []byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockSocketSourceMockRecorder) Read(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockSocketSource)(nil).Read), ctx)
}
