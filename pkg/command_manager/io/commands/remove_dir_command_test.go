package commands

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

func TestRemoveDirCommand_Execute(t *testing.T) {
	t.Run("should execute rm_dir command", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := false
		command := generateRemoveDirCommand(abortOnFail)

		receiverMock.EXPECT().RemoveDir(gomock.Any()).Return(nil)

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})

	t.Run("should not execute rm_dir and abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := true
		command := generateRemoveDirCommand(abortOnFail)

		receiverMock.EXPECT().RemoveDir(gomock.Any()).Return(errors.New("error"))

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
		command := generateRemoveDirCommand(abortOnFail)

		receiverMock.EXPECT().RemoveDir(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})
}

func generateRemoveDirCommand(abortOnFail bool) command.Command {
	return NewRemoveDirCommand(
		"Remove dir",
		receiverMock,
		&map[string]string{"path": "/tmp/"},
		command.RemoveDir,
		abortOnFail)
}
