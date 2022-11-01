package commands

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

func TestPutContentCommand_Execute(t *testing.T) {
	t.Run("should execute put_content command", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		command := NewPutContentCommand(
			"Put content",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.PutContent,
			true)

		receiverMock.EXPECT().PutContent(gomock.Any()).Return(nil)

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})

	t.Run("should not execute put_content and abort", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		abortOnFail := true

		command := NewPutContentCommand(
			"Put content",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.PutContent,
			abortOnFail)

		receiverMock.EXPECT().PutContent(gomock.Any()).Return(errors.New("error"))

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

		command := NewPutContentCommand(
			"Put content",
			receiverMock,
			&map[string]string{"path": "/tmp/file.txt"},
			command.PutContent,
			abortOnFail)

		receiverMock.EXPECT().PutContent(gomock.Any()).Return(errors.New("error"))

		// When
		err := command.Execute()

		// Then
		assert.Nil(t, err)

	})
}
