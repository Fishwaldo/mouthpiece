// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Fishwaldo/mouthpiece/pkg/interfaces (interfaces: UserI)

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"

	interfaces "github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	logr "github.com/go-logr/logr"
	gomock "github.com/golang/mock/gomock"
)

// MockUserI is a mock of UserI interface.
type MockUserI struct {
	ctrl     *gomock.Controller
	recorder *MockUserIMockRecorder
}

// MockUserIMockRecorder is the mock recorder for MockUserI.
type MockUserIMockRecorder struct {
	mock *MockUserI
}

// NewMockUserI creates a new mock instance.
func NewMockUserI(ctrl *gomock.Controller) *MockUserI {
	mock := &MockUserI{ctrl: ctrl}
	mock.recorder = &MockUserIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserI) EXPECT() *MockUserIMockRecorder {
	return m.recorder
}

// AddFilter mocks base method.
func (m *MockUserI) AddFilter(arg0 context.Context, arg1 interfaces.FilterI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFilter indicates an expected call of AddFilter.
func (mr *MockUserIMockRecorder) AddFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilter", reflect.TypeOf((*MockUserI)(nil).AddFilter), arg0, arg1)
}

// AddTransportRecipient mocks base method.
func (m *MockUserI) AddTransportRecipient(arg0 context.Context, arg1 interfaces.TransportRecipient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTransportRecipient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTransportRecipient indicates an expected call of AddTransportRecipient.
func (mr *MockUserIMockRecorder) AddTransportRecipient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTransportRecipient", reflect.TypeOf((*MockUserI)(nil).AddTransportRecipient), arg0, arg1)
}

// DelFilter mocks base method.
func (m *MockUserI) DelFilter(arg0 context.Context, arg1 interfaces.FilterI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelFilter", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelFilter indicates an expected call of DelFilter.
func (mr *MockUserIMockRecorder) DelFilter(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelFilter", reflect.TypeOf((*MockUserI)(nil).DelFilter), arg0, arg1)
}

// DelTransportRecipient mocks base method.
func (m *MockUserI) DelTransportRecipient(arg0 context.Context, arg1 interfaces.TransportRecipient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DelTransportRecipient", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DelTransportRecipient indicates an expected call of DelTransportRecipient.
func (mr *MockUserIMockRecorder) DelTransportRecipient(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DelTransportRecipient", reflect.TypeOf((*MockUserI)(nil).DelTransportRecipient), arg0, arg1)
}

// GetDescription mocks base method.
func (m *MockUserI) GetDescription() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDescription")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetDescription indicates an expected call of GetDescription.
func (mr *MockUserIMockRecorder) GetDescription() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDescription", reflect.TypeOf((*MockUserI)(nil).GetDescription))
}

// GetEmail mocks base method.
func (m *MockUserI) GetEmail() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetEmail")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetEmail indicates an expected call of GetEmail.
func (mr *MockUserIMockRecorder) GetEmail() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetEmail", reflect.TypeOf((*MockUserI)(nil).GetEmail))
}

// GetField mocks base method.
func (m *MockUserI) GetField(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetField", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetField indicates an expected call of GetField.
func (mr *MockUserIMockRecorder) GetField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockUserI)(nil).GetField), arg0, arg1)
}

// GetFields mocks base method.
func (m *MockUserI) GetFields(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFields", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFields indicates an expected call of GetFields.
func (mr *MockUserIMockRecorder) GetFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFields", reflect.TypeOf((*MockUserI)(nil).GetFields), arg0)
}

// GetFilters mocks base method.
func (m *MockUserI) GetFilters(arg0 context.Context) ([]interfaces.FilterI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilters", arg0)
	ret0, _ := ret[0].([]interfaces.FilterI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilters indicates an expected call of GetFilters.
func (mr *MockUserIMockRecorder) GetFilters(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilters", reflect.TypeOf((*MockUserI)(nil).GetFilters), arg0)
}

// GetID mocks base method.
func (m *MockUserI) GetID() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockUserIMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockUserI)(nil).GetID))
}

// GetName mocks base method.
func (m *MockUserI) GetName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockUserIMockRecorder) GetName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockUserI)(nil).GetName))
}

// GetTransportRecipients mocks base method.
func (m *MockUserI) GetTransportRecipients(arg0 context.Context) []interfaces.TransportRecipient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransportRecipients", arg0)
	ret0, _ := ret[0].([]interfaces.TransportRecipient)
	return ret0
}

// GetTransportRecipients indicates an expected call of GetTransportRecipients.
func (mr *MockUserIMockRecorder) GetTransportRecipients(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransportRecipients", reflect.TypeOf((*MockUserI)(nil).GetTransportRecipients), arg0)
}

// Load mocks base method.
func (m *MockUserI) Load(arg0 context.Context, arg1 logr.Logger, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockUserIMockRecorder) Load(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockUserI)(nil).Load), arg0, arg1, arg2)
}

// ProcessMessage mocks base method.
func (m *MockUserI) ProcessMessage(arg0 context.Context, arg1 interfaces.MessageI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockUserIMockRecorder) ProcessMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockUserI)(nil).ProcessMessage), arg0, arg1)
}

// Save mocks base method.
func (m *MockUserI) Save(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockUserIMockRecorder) Save(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockUserI)(nil).Save), arg0)
}

// SetDescription mocks base method.
func (m *MockUserI) SetDescription(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetDescription", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetDescription indicates an expected call of SetDescription.
func (mr *MockUserIMockRecorder) SetDescription(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetDescription", reflect.TypeOf((*MockUserI)(nil).SetDescription), arg0, arg1)
}

// SetEmail mocks base method.
func (m *MockUserI) SetEmail(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEmail", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEmail indicates an expected call of SetEmail.
func (mr *MockUserIMockRecorder) SetEmail(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEmail", reflect.TypeOf((*MockUserI)(nil).SetEmail), arg0, arg1)
}

// SetField mocks base method.
func (m *MockUserI) SetField(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetField", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetField indicates an expected call of SetField.
func (mr *MockUserIMockRecorder) SetField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetField", reflect.TypeOf((*MockUserI)(nil).SetField), arg0, arg1, arg2)
}

// SetFields mocks base method.
func (m *MockUserI) SetFields(arg0 context.Context, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFields", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFields indicates an expected call of SetFields.
func (mr *MockUserIMockRecorder) SetFields(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFields", reflect.TypeOf((*MockUserI)(nil).SetFields), arg0, arg1)
}

// SetName mocks base method.
func (m *MockUserI) SetName(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetName", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetName indicates an expected call of SetName.
func (mr *MockUserIMockRecorder) SetName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetName", reflect.TypeOf((*MockUserI)(nil).SetName), arg0, arg1)
}
