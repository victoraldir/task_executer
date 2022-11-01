package commands

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

func TestRemoveFileCommand_Execute(t *testing.T) {
	t.Run("should execute rm_file command", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := false
		command := generateRemoveFileCommand(abortOnFail)

		receiverMock.EXPECT().RemoveFile(gomock.Any()).Return(nil)

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})

	t.Run("should not execute rm_file and abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := true
		command := generateRemoveFileCommand(abortOnFail)

		receiverMock.EXPECT().RemoveFile(gomock.Any()).Return(errors.New("error"))

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
		command := generateRemoveFileCommand(abortOnFail)

		receiverMock.EXPECT().RemoveFile(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})
}

func generateRemoveFileCommand(abortOnFail bool) command.Command {
	return NewRemoveFileCommand(
		"Remove file",
		receiverMock,
		&map[string]string{"path": "/tmp/test.txt"},
		command.RemoveFile,
		abortOnFail)
}
