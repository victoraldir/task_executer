// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/victoraldir/task_executer/internal/app/task/core/ports (interfaces: LoaderRepository)

// Package ports is a generated GoMock package.
package ports

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	domains "github.com/victoraldir/task_executer/internal/app/task/core/domains"
)

// MockLoaderRepository is a mock of LoaderRepository interface.
type MockLoaderRepository struct {
	ctrl     *gomock.Controller
	recorder *MockLoaderRepositoryMockRecorder
}

// MockLoaderRepositoryMockRecorder is the mock recorder for MockLoaderRepository.
type MockLoaderRepositoryMockRecorder struct {
	mock *MockLoaderRepository
}

// NewMockLoaderRepository creates a new mock instance.
func NewMockLoaderRepository(ctrl *gomock.Controller) *MockLoaderRepository {
	mock := &MockLoaderRepository{ctrl: ctrl}
	mock.recorder = &MockLoaderRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockLoaderRepository) EXPECT() *MockLoaderRepositoryMockRecorder {
	return m.recorder
}

// LoadTasks mocks base method.
func (m *MockLoaderRepository) LoadTasks() (*[]domains.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadTasks")
	ret0, _ := ret[0].(*[]domains.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LoadTasks indicates an expected call of LoadTasks.
func (mr *MockLoaderRepositoryMockRecorder) LoadTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadTasks", reflect.TypeOf((*MockLoaderRepository)(nil).LoadTasks))
}