package commands

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

func TestCreateDirCommand_Execute(t *testing.T) {
	t.Run("should execute create_dir command", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		command := NewCreateDirCommand(
			"Create dir",
			receiverMock,
			&map[string]string{"path": "/tmp"},
			command.CreateDir,
			true)

		receiverMock.EXPECT().CreateDir(gomock.Any()).Return(nil)

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})

	t.Run("should not execute create_dir and abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := true

		command := NewCreateDirCommand(
			"Create dir",
			receiverMock,
			&map[string]string{"path": "/tmp"},
			command.CreateDir,
			abortOnFail)

		receiverMock.EXPECT().CreateDir(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.NotNil(t, err)

	})

	t.Run("should fail command execution but does not abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := false

		command := NewCreateDirCommand(
			"Create dir",
			receiverMock,
			&map[string]string{"path": "/tmp"},
			command.CreateDir,
			abortOnFail)

		receiverMock.EXPECT().CreateDir(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})
}
