package adapters

import (
	"testing"

	"github.com/golang/mock/gomock"
	adapters "github.com/victoraldir/task_executer/internal/app/task/adapters/mocks"
	command_mock "github.com/victoraldir/task_executer/pkg/command_manager/mocks"
)

var commandCreateDirMock *command_mock.MockCommand
var readerMock *adapters.MockReader

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	commandCreateDirMock = command_mock.NewMockCommand(ctrl)
	readerMock = adapters.NewMockReader(ctrl)

	return func() {
		defer ctrl.Finish()
	}
}
