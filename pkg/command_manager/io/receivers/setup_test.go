package receivers

import (
	"testing"

	"github.com/golang/mock/gomock"
	receivers "github.com/victoraldir/task_executer/pkg/command_manager/io/receivers/mocks"
)

var fileSystemMock *receivers.MockFileSystem

func setup(t *testing.T) func() {
	ctrl := gomock.NewController(t)

	fileSystemMock = receivers.NewMockFileSystem(ctrl)

	return func() {
		defer ctrl.Finish()
	}
}
