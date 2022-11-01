package commands

import (
	"testing"

	"github.com/golang/mock/gomock"
	command "github.com/victoraldir/task_executer/pkg/command_manager/mocks"
)

var receiverMock *command.MockReceiver

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	receiverMock = command.NewMockReceiver(ctrl)

	return func() {
		defer ctrl.Finish()
	}
}
