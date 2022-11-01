package usecases

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/victoraldir/task_executer/internal/app/task/core/domains"
	ports "github.com/victoraldir/task_executer/internal/app/task/core/ports/mocks"
)

var taskLoaderRepoMock *ports.MockLoaderRepository
var taskInvokerRepoMock *ports.MockInvokerRepository

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	taskLoaderRepoMock = ports.NewMockLoaderRepository(ctrl)
	taskInvokerRepoMock = ports.NewMockInvokerRepository(ctrl)

	return func() {
		defer ctrl.Finish()
	}
}

func TestExecuteFlowUseCase_Execute(t *testing.T) {
	t.Run("should execute flow", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		tasks := generateTasks()

		taskLoaderRepoMock.EXPECT().LoadTasks().Return(tasks, nil)
		taskInvokerRepoMock.EXPECT().Invoke(*tasks).Return(nil)

		executeFlowUseCase := NewExecuteFlowUseCase(taskLoaderRepoMock, taskInvokerRepoMock)

		// When
		err := executeFlowUseCase.Execute()

		// Then
		assert.Nil(t, err)
	})

	t.Run("should not execute flow. Loader task loader fails", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		taskLoaderRepoMock.EXPECT().LoadTasks().Return(nil, errors.New("error"))

		executeFlowUseCase := NewExecuteFlowUseCase(taskLoaderRepoMock, taskInvokerRepoMock)

		// When
		err := executeFlowUseCase.Execute()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})

	t.Run("should not execute flow. Task invoker fails", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		tasks := generateTasks()

		taskLoaderRepoMock.EXPECT().LoadTasks().Return(tasks, nil)
		taskInvokerRepoMock.EXPECT().Invoke(*tasks).Return(errors.New("error"))

		executeFlowUseCase := NewExecuteFlowUseCase(taskLoaderRepoMock, taskInvokerRepoMock)

		// When
		err := executeFlowUseCase.Execute()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
}

func generateTasks() *[]domains.Task {
	return &[]domains.Task{
		{
			Name:        "test create dir",
			Type:        "create_dir",
			AbortOnFail: true,
			Args: map[string]string{
				"path": "/tmp/test",
			},
		},
		{
			Name:        "test delete dir",
			Type:        "rm_dir",
			AbortOnFail: true,
			Args: map[string]string{
				"path": "/tmp/test",
			},
		},
	}
}
