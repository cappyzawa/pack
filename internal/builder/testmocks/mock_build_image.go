// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/buildpack/pack/internal/builder (interfaces: BuildImage)

// Package testmocks is a generated GoMock package.
package testmocks

import (
	io "io"
	reflect "reflect"
	time "time"

	imgutil "github.com/buildpack/imgutil"
	gomock "github.com/golang/mock/gomock"
)

// MockBuildImage is a mock of BuildImage interface
type MockBuildImage struct {
	ctrl     *gomock.Controller
	recorder *MockBuildImageMockRecorder
}

// MockBuildImageMockRecorder is the mock recorder for MockBuildImage
type MockBuildImageMockRecorder struct {
	mock *MockBuildImage
}

// NewMockBuildImage creates a new mock instance
func NewMockBuildImage(ctrl *gomock.Controller) *MockBuildImage {
	mock := &MockBuildImage{ctrl: ctrl}
	mock.recorder = &MockBuildImageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBuildImage) EXPECT() *MockBuildImageMockRecorder {
	return m.recorder
}

// AddLayer mocks base method
func (m *MockBuildImage) AddLayer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddLayer indicates an expected call of AddLayer
func (mr *MockBuildImageMockRecorder) AddLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddLayer", reflect.TypeOf((*MockBuildImage)(nil).AddLayer), arg0)
}

// BuildOnlyMixins mocks base method
func (m *MockBuildImage) BuildOnlyMixins() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuildOnlyMixins")
	ret0, _ := ret[0].([]string)
	return ret0
}

// BuildOnlyMixins indicates an expected call of BuildOnlyMixins
func (mr *MockBuildImageMockRecorder) BuildOnlyMixins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuildOnlyMixins", reflect.TypeOf((*MockBuildImage)(nil).BuildOnlyMixins))
}

// CommonMixins mocks base method
func (m *MockBuildImage) CommonMixins() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommonMixins")
	ret0, _ := ret[0].([]string)
	return ret0
}

// CommonMixins indicates an expected call of CommonMixins
func (mr *MockBuildImageMockRecorder) CommonMixins() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommonMixins", reflect.TypeOf((*MockBuildImage)(nil).CommonMixins))
}

// CreatedAt mocks base method
func (m *MockBuildImage) CreatedAt() (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatedAt")
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatedAt indicates an expected call of CreatedAt
func (mr *MockBuildImageMockRecorder) CreatedAt() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatedAt", reflect.TypeOf((*MockBuildImage)(nil).CreatedAt))
}

// Delete mocks base method
func (m *MockBuildImage) Delete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete")
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockBuildImageMockRecorder) Delete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockBuildImage)(nil).Delete))
}

// Env mocks base method
func (m *MockBuildImage) Env(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Env", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Env indicates an expected call of Env
func (mr *MockBuildImageMockRecorder) Env(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Env", reflect.TypeOf((*MockBuildImage)(nil).Env), arg0)
}

// Found mocks base method
func (m *MockBuildImage) Found() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Found")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Found indicates an expected call of Found
func (mr *MockBuildImageMockRecorder) Found() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Found", reflect.TypeOf((*MockBuildImage)(nil).Found))
}

// GetLayer mocks base method
func (m *MockBuildImage) GetLayer(arg0 string) (io.ReadCloser, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetLayer", arg0)
	ret0, _ := ret[0].(io.ReadCloser)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetLayer indicates an expected call of GetLayer
func (mr *MockBuildImageMockRecorder) GetLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetLayer", reflect.TypeOf((*MockBuildImage)(nil).GetLayer), arg0)
}

// Identifier mocks base method
func (m *MockBuildImage) Identifier() (imgutil.Identifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Identifier")
	ret0, _ := ret[0].(imgutil.Identifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Identifier indicates an expected call of Identifier
func (mr *MockBuildImageMockRecorder) Identifier() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Identifier", reflect.TypeOf((*MockBuildImage)(nil).Identifier))
}

// Label mocks base method
func (m *MockBuildImage) Label(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Label", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Label indicates an expected call of Label
func (mr *MockBuildImageMockRecorder) Label(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Label", reflect.TypeOf((*MockBuildImage)(nil).Label), arg0)
}

// Name mocks base method
func (m *MockBuildImage) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockBuildImageMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockBuildImage)(nil).Name))
}

// Rebase mocks base method
func (m *MockBuildImage) Rebase(arg0 string, arg1 imgutil.Image) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Rebase", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Rebase indicates an expected call of Rebase
func (mr *MockBuildImageMockRecorder) Rebase(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rebase", reflect.TypeOf((*MockBuildImage)(nil).Rebase), arg0, arg1)
}

// Rename mocks base method
func (m *MockBuildImage) Rename(arg0 string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Rename", arg0)
}

// Rename indicates an expected call of Rename
func (mr *MockBuildImageMockRecorder) Rename(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Rename", reflect.TypeOf((*MockBuildImage)(nil).Rename), arg0)
}

// ReuseLayer mocks base method
func (m *MockBuildImage) ReuseLayer(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReuseLayer", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReuseLayer indicates an expected call of ReuseLayer
func (mr *MockBuildImageMockRecorder) ReuseLayer(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReuseLayer", reflect.TypeOf((*MockBuildImage)(nil).ReuseLayer), arg0)
}

// Save mocks base method
func (m *MockBuildImage) Save(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save
func (mr *MockBuildImageMockRecorder) Save(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockBuildImage)(nil).Save), arg0...)
}

// SetCmd mocks base method
func (m *MockBuildImage) SetCmd(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetCmd", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetCmd indicates an expected call of SetCmd
func (mr *MockBuildImageMockRecorder) SetCmd(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetCmd", reflect.TypeOf((*MockBuildImage)(nil).SetCmd), arg0...)
}

// SetEntrypoint mocks base method
func (m *MockBuildImage) SetEntrypoint(arg0 ...string) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SetEntrypoint", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEntrypoint indicates an expected call of SetEntrypoint
func (mr *MockBuildImageMockRecorder) SetEntrypoint(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEntrypoint", reflect.TypeOf((*MockBuildImage)(nil).SetEntrypoint), arg0...)
}

// SetEnv mocks base method
func (m *MockBuildImage) SetEnv(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetEnv", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetEnv indicates an expected call of SetEnv
func (mr *MockBuildImageMockRecorder) SetEnv(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetEnv", reflect.TypeOf((*MockBuildImage)(nil).SetEnv), arg0, arg1)
}

// SetLabel mocks base method
func (m *MockBuildImage) SetLabel(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetLabel", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetLabel indicates an expected call of SetLabel
func (mr *MockBuildImageMockRecorder) SetLabel(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetLabel", reflect.TypeOf((*MockBuildImage)(nil).SetLabel), arg0, arg1)
}

// SetWorkingDir mocks base method
func (m *MockBuildImage) SetWorkingDir(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetWorkingDir", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetWorkingDir indicates an expected call of SetWorkingDir
func (mr *MockBuildImageMockRecorder) SetWorkingDir(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetWorkingDir", reflect.TypeOf((*MockBuildImage)(nil).SetWorkingDir), arg0)
}

// StackID mocks base method
func (m *MockBuildImage) StackID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StackID")
	ret0, _ := ret[0].(string)
	return ret0
}

// StackID indicates an expected call of StackID
func (mr *MockBuildImageMockRecorder) StackID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StackID", reflect.TypeOf((*MockBuildImage)(nil).StackID))
}

// TopLayer mocks base method
func (m *MockBuildImage) TopLayer() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopLayer")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TopLayer indicates an expected call of TopLayer
func (mr *MockBuildImageMockRecorder) TopLayer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopLayer", reflect.TypeOf((*MockBuildImage)(nil).TopLayer))
}
