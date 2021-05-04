// Code generated by MockGen. DO NOT EDIT.
// Source: graph/model/graph.go

// Package model is a generated GoMock package.
package model

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockVertex is a mock of Vertex interface.
type MockVertex struct {
	ctrl     *gomock.Controller
	recorder *MockVertexMockRecorder
}

// MockVertexMockRecorder is the mock recorder for MockVertex.
type MockVertexMockRecorder struct {
	mock *MockVertex
}

// NewMockVertex creates a new mock instance.
func NewMockVertex(ctrl *gomock.Controller) *MockVertex {
	mock := &MockVertex{ctrl: ctrl}
	mock.recorder = &MockVertexMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVertex) EXPECT() *MockVertexMockRecorder {
	return m.recorder
}

// Key mocks base method.
func (m *MockVertex) Key() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Key")
	ret0, _ := ret[0].(string)
	return ret0
}

// Key indicates an expected call of Key.
func (mr *MockVertexMockRecorder) Key() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Key", reflect.TypeOf((*MockVertex)(nil).Key))
}

// Value mocks base method.
func (m *MockVertex) Value() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Value")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// Value indicates an expected call of Value.
func (mr *MockVertexMockRecorder) Value() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Value", reflect.TypeOf((*MockVertex)(nil).Value))
}
