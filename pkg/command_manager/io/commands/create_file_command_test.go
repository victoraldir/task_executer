package commands

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

func TestCreateFileCommand_Execute(t *testing.T) {
	t.Run("should execute create_file command", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		command := NewCreateFileCommand(
			"Create file",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.CreateFile,
			true)

		receiverMock.EXPECT().CreateFile(gomock.Any()).Return(nil)

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})

	t.Run("should not execute create_file and abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := true

		command := NewCreateFileCommand(
			"Create file",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.CreateFile,
			abortOnFail)

		receiverMock.EXPECT().CreateFile(gomock.Any()).Return(errors.New("error"))

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

		command := NewCreateFileCommand(
			"Create file",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.CreateFile,
			abortOnFail)

		receiverMock.EXPECT().CreateFile(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})
}
