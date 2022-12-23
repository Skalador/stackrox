// Code generated by MockGen. DO NOT EDIT.
// Source: telemeter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockTelemeter is a mock of Telemeter interface.
type MockTelemeter struct {
	ctrl     *gomock.Controller
	recorder *MockTelemeterMockRecorder
}

// MockTelemeterMockRecorder is the mock recorder for MockTelemeter.
type MockTelemeterMockRecorder struct {
	mock *MockTelemeter
}

// NewMockTelemeter creates a new mock instance.
func NewMockTelemeter(ctrl *gomock.Controller) *MockTelemeter {
	mock := &MockTelemeter{ctrl: ctrl}
	mock.recorder = &MockTelemeterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTelemeter) EXPECT() *MockTelemeterMockRecorder {
	return m.recorder
}

// Group mocks base method.
func (m *MockTelemeter) Group(groupID, userID string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Group", groupID, userID, props)
}

// Group indicates an expected call of Group.
func (mr *MockTelemeterMockRecorder) Group(groupID, userID, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockTelemeter)(nil).Group), groupID, userID, props)
}

// Identify mocks base method.
func (m *MockTelemeter) Identify(userID string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Identify", userID, props)
}

// Identify indicates an expected call of Identify.
func (mr *MockTelemeterMockRecorder) Identify(userID, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identify", reflect.TypeOf((*MockTelemeter)(nil).Identify), userID, props)
}

// Stop mocks base method.
func (m *MockTelemeter) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockTelemeterMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockTelemeter)(nil).Stop))
}

// Track mocks base method.
func (m *MockTelemeter) Track(event, userID string, props map[string]any) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Track", event, userID, props)
}

// Track indicates an expected call of Track.
func (mr *MockTelemeterMockRecorder) Track(event, userID, props interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Track", reflect.TypeOf((*MockTelemeter)(nil).Track), event, userID, props)
}
