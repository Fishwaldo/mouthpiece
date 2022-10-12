// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Fishwaldo/mouthpiece/pkg/interfaces (interfaces: MessageI)

// Package mock_interfaces is a generated GoMock package.
package mock_interfaces

import (
	context "context"
	reflect "reflect"
	time "time"

	interfaces "github.com/Fishwaldo/mouthpiece/pkg/interfaces"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockMessageI is a mock of MessageI interface.
type MockMessageI struct {
	ctrl     *gomock.Controller
	recorder *MockMessageIMockRecorder
}

// MockMessageIMockRecorder is the mock recorder for MockMessageI.
type MockMessageIMockRecorder struct {
	mock *MockMessageI
}

// NewMockMessageI creates a new mock instance.
func NewMockMessageI(ctrl *gomock.Controller) *MockMessageI {
	mock := &MockMessageI{ctrl: ctrl}
	mock.recorder = &MockMessageIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMessageI) EXPECT() *MockMessageIMockRecorder {
	return m.recorder
}

// Clone mocks base method.
func (m *MockMessageI) Clone() interfaces.MessageI {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Clone")
	ret0, _ := ret[0].(interfaces.MessageI)
	return ret0
}

// Clone indicates an expected call of Clone.
func (mr *MockMessageIMockRecorder) Clone() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Clone", reflect.TypeOf((*MockMessageI)(nil).Clone))
}

// GetApp mocks base method.
func (m *MockMessageI) GetApp(arg0 context.Context) (interfaces.AppI, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetApp", arg0)
	ret0, _ := ret[0].(interfaces.AppI)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetApp indicates an expected call of GetApp.
func (mr *MockMessageIMockRecorder) GetApp(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetApp", reflect.TypeOf((*MockMessageI)(nil).GetApp), arg0)
}

// GetField mocks base method.
func (m *MockMessageI) GetField(arg0 context.Context, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetField", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetField indicates an expected call of GetField.
func (mr *MockMessageIMockRecorder) GetField(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetField", reflect.TypeOf((*MockMessageI)(nil).GetField), arg0, arg1)
}

// GetFields mocks base method.
func (m *MockMessageI) GetFields(arg0 context.Context) (map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFields", arg0)
	ret0, _ := ret[0].(map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFields indicates an expected call of GetFields.
func (mr *MockMessageIMockRecorder) GetFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFields", reflect.TypeOf((*MockMessageI)(nil).GetFields), arg0)
}

// GetID mocks base method.
func (m *MockMessageI) GetID() uuid.UUID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetID")
	ret0, _ := ret[0].(uuid.UUID)
	return ret0
}

// GetID indicates an expected call of GetID.
func (mr *MockMessageIMockRecorder) GetID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetID", reflect.TypeOf((*MockMessageI)(nil).GetID))
}

// GetMessage mocks base method.
func (m *MockMessageI) GetMessage() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMessage")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetMessage indicates an expected call of GetMessage.
func (mr *MockMessageIMockRecorder) GetMessage() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMessage", reflect.TypeOf((*MockMessageI)(nil).GetMessage))
}

// GetMetadata mocks base method.
func (m *MockMessageI) GetMetadata(arg0 context.Context, arg1 string) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadata", arg0, arg1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadata indicates an expected call of GetMetadata.
func (mr *MockMessageIMockRecorder) GetMetadata(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadata", reflect.TypeOf((*MockMessageI)(nil).GetMetadata), arg0, arg1)
}

// GetMetadataFields mocks base method.
func (m *MockMessageI) GetMetadataFields(arg0 context.Context) (map[string]interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMetadataFields", arg0)
	ret0, _ := ret[0].(map[string]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMetadataFields indicates an expected call of GetMetadataFields.
func (mr *MockMessageIMockRecorder) GetMetadataFields(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMetadataFields", reflect.TypeOf((*MockMessageI)(nil).GetMetadataFields), arg0)
}

// GetSeverity mocks base method.
func (m *MockMessageI) GetSeverity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSeverity")
	ret0, _ := ret[0].(int)
	return ret0
}

// GetSeverity indicates an expected call of GetSeverity.
func (mr *MockMessageIMockRecorder) GetSeverity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSeverity", reflect.TypeOf((*MockMessageI)(nil).GetSeverity))
}

// GetShortMsg mocks base method.
func (m *MockMessageI) GetShortMsg() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetShortMsg")
	ret0, _ := ret[0].(*string)
	return ret0
}

// GetShortMsg indicates an expected call of GetShortMsg.
func (mr *MockMessageIMockRecorder) GetShortMsg() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetShortMsg", reflect.TypeOf((*MockMessageI)(nil).GetShortMsg))
}

// GetTimestamp mocks base method.
func (m *MockMessageI) GetTimestamp() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTimestamp")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// GetTimestamp indicates an expected call of GetTimestamp.
func (mr *MockMessageIMockRecorder) GetTimestamp() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTimestamp", reflect.TypeOf((*MockMessageI)(nil).GetTimestamp))
}

// GetTopic mocks base method.
func (m *MockMessageI) GetTopic() *string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTopic")
	ret0, _ := ret[0].(*string)
	return ret0
}

// GetTopic indicates an expected call of GetTopic.
func (mr *MockMessageIMockRecorder) GetTopic() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTopic", reflect.TypeOf((*MockMessageI)(nil).GetTopic))
}

// Load mocks base method.
func (m *MockMessageI) Load(arg0 context.Context, arg1 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockMessageIMockRecorder) Load(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockMessageI)(nil).Load), arg0, arg1)
}

// ProcessMessage mocks base method.
func (m *MockMessageI) ProcessMessage(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ProcessMessage", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ProcessMessage indicates an expected call of ProcessMessage.
func (mr *MockMessageIMockRecorder) ProcessMessage(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ProcessMessage", reflect.TypeOf((*MockMessageI)(nil).ProcessMessage), arg0)
}

// Save mocks base method.
func (m *MockMessageI) Save(arg0 context.Context, arg1 interfaces.AppI) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockMessageIMockRecorder) Save(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockMessageI)(nil).Save), arg0, arg1)
}

// SetField mocks base method.
func (m *MockMessageI) SetField(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetField", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetField indicates an expected call of SetField.
func (mr *MockMessageIMockRecorder) SetField(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetField", reflect.TypeOf((*MockMessageI)(nil).SetField), arg0, arg1, arg2)
}

// SetFields mocks base method.
func (m *MockMessageI) SetFields(arg0 context.Context, arg1 map[string]string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetFields", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetFields indicates an expected call of SetFields.
func (mr *MockMessageIMockRecorder) SetFields(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetFields", reflect.TypeOf((*MockMessageI)(nil).SetFields), arg0, arg1)
}

// SetMessage mocks base method.
func (m *MockMessageI) SetMessage(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMessage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMessage indicates an expected call of SetMessage.
func (mr *MockMessageIMockRecorder) SetMessage(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMessage", reflect.TypeOf((*MockMessageI)(nil).SetMessage), arg0, arg1)
}

// SetMetadata mocks base method.
func (m *MockMessageI) SetMetadata(arg0 context.Context, arg1 string, arg2 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetMetadata", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetMetadata indicates an expected call of SetMetadata.
func (mr *MockMessageIMockRecorder) SetMetadata(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMetadata", reflect.TypeOf((*MockMessageI)(nil).SetMetadata), arg0, arg1, arg2)
}

// SetSeverity mocks base method.
func (m *MockMessageI) SetSeverity(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSeverity", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetSeverity indicates an expected call of SetSeverity.
func (mr *MockMessageIMockRecorder) SetSeverity(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSeverity", reflect.TypeOf((*MockMessageI)(nil).SetSeverity), arg0, arg1)
}

// SetShortMsg mocks base method.
func (m *MockMessageI) SetShortMsg(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetShortMsg", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetShortMsg indicates an expected call of SetShortMsg.
func (mr *MockMessageIMockRecorder) SetShortMsg(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetShortMsg", reflect.TypeOf((*MockMessageI)(nil).SetShortMsg), arg0, arg1)
}

// SetTimestamp mocks base method.
func (m *MockMessageI) SetTimestamp(arg0 context.Context, arg1 time.Time) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTimestamp", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTimestamp indicates an expected call of SetTimestamp.
func (mr *MockMessageIMockRecorder) SetTimestamp(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimestamp", reflect.TypeOf((*MockMessageI)(nil).SetTimestamp), arg0, arg1)
}

// SetTopic mocks base method.
func (m *MockMessageI) SetTopic(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetTopic", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetTopic indicates an expected call of SetTopic.
func (mr *MockMessageIMockRecorder) SetTopic(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTopic", reflect.TypeOf((*MockMessageI)(nil).SetTopic), arg0, arg1)
}

// String mocks base method.
func (m *MockMessageI) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockMessageIMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockMessageI)(nil).String))
}