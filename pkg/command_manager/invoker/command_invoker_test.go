package invoker

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
	command_mock "github.com/victoraldir/task_executer/pkg/command_manager/mocks"
)

var commandMock *command_mock.MockCommand

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	commandMock = command_mock.NewMockCommand(ctrl)

	return func() {
		defer ctrl.Finish()
	}
}

func TestCommandInvoker_Invoke(t *testing.T) {

	t.Run("should invoke command successfully", func(t *testing.T) {
		// given
		teardown := setup(t)
		defer teardown()

		commandMock.EXPECT().Execute().Return(nil)

		commandInvoker := NewCommandInvoker([]command.Command{commandMock})

		// when
		err := commandInvoker.Invoke()

		// then
		assert.Nil(t, err)
	})

	t.Run("should return error when command fails", func(t *testing.T) {
		// given
		teardown := setup(t)
		defer teardown()

		commandMock.EXPECT().Execute().Return(errors.New("error"))

		commandInvoker := NewCommandInvoker([]command.Command{commandMock})

		// when
		err := commandInvoker.Invoke()

		// then
		assert.NotNil(t, err)
	})
}
