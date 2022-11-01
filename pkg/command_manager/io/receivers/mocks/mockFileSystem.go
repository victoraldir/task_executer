// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/victoraldir/task_executer/pkg/command_manager/io/receivers (interfaces: FileSystem)

// Package receivers is a generated GoMock package.
package receivers

import (
	fs "io/fs"
	os "os"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFileSystem is a mock of FileSystem interface.
type MockFileSystem struct {
	ctrl     *gomock.Controller
	recorder *MockFileSystemMockRecorder
}

// MockFileSystemMockRecorder is the mock recorder for MockFileSystem.
type MockFileSystemMockRecorder struct {
	mock *MockFileSystem
}

// NewMockFileSystem creates a new mock instance.
func NewMockFileSystem(ctrl *gomock.Controller) *MockFileSystem {
	mock := &MockFileSystem{ctrl: ctrl}
	mock.recorder = &MockFileSystemMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFileSystem) EXPECT() *MockFileSystemMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockFileSystem) Create(arg0 string) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockFileSystemMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockFileSystem)(nil).Create), arg0)
}

// IsNotExist mocks base method.
func (m *MockFileSystem) IsNotExist(arg0 error) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsNotExist", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsNotExist indicates an expected call of IsNotExist.
func (mr *MockFileSystemMockRecorder) IsNotExist(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsNotExist", reflect.TypeOf((*MockFileSystem)(nil).IsNotExist), arg0)
}

// Mkdir mocks base method.
func (m *MockFileSystem) Mkdir(arg0 string, arg1 fs.FileMode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mkdir", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Mkdir indicates an expected call of Mkdir.
func (mr *MockFileSystemMockRecorder) Mkdir(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mkdir", reflect.TypeOf((*MockFileSystem)(nil).Mkdir), arg0, arg1)
}

// OpenFile mocks base method.
func (m *MockFileSystem) OpenFile(arg0 string, arg1 int, arg2 fs.FileMode) (*os.File, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "OpenFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(*os.File)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// OpenFile indicates an expected call of OpenFile.
func (mr *MockFileSystemMockRecorder) OpenFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OpenFile", reflect.TypeOf((*MockFileSystem)(nil).OpenFile), arg0, arg1, arg2)
}

// Remove mocks base method.
func (m *MockFileSystem) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockFileSystemMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockFileSystem)(nil).Remove), arg0)
}

// RemoveAll mocks base method.
func (m *MockFileSystem) RemoveAll(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAll", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAll indicates an expected call of RemoveAll.
func (mr *MockFileSystemMockRecorder) RemoveAll(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAll", reflect.TypeOf((*MockFileSystem)(nil).RemoveAll), arg0)
}

// Stat mocks base method.
func (m *MockFileSystem) Stat(arg0 string) (fs.FileInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stat", arg0)
	ret0, _ := ret[0].(fs.FileInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stat indicates an expected call of Stat.
func (mr *MockFileSystemMockRecorder) Stat(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stat", reflect.TypeOf((*MockFileSystem)(nil).Stat), arg0)
}
