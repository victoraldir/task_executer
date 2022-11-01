package adapters

import (
	"errors"

	"github.com/victoraldir/task_executer/internal/app/task/core/domains"
	"github.com/victoraldir/task_executer/internal/app/task/core/ports"
	command "github.com/victoraldir/task_executer/pkg/command_manager"
	"github.com/victoraldir/task_executer/pkg/command_manager/invoker"
	"github.com/victoraldir/task_executer/pkg/command_manager/io/commands"
	"github.com/victoraldir/task_executer/pkg/command_manager/io/receivers"
)

type commandManagerInvokerRepository struct {
	fs receivers.FileSystem
}

func NewCommandManagerInvokerRepository(fs receivers.FileSystem) ports.InvokerRepository {
	return &commandManagerInvokerRepository{
		fs: fs,
	}
}

func (c *commandManagerInvokerRepository) Invoke(tasks []domains.Task) error {

	var commands []command.Command

	for _, t := range tasks {

		cmd, err := ToCommand(t, c.fs)
		if err != nil {
			return err
		}

		commands = append(commands, cmd)
	}

	commandInvoker := invoker.NewCommandInvoker(commands)

	return commandInvoker.Invoke()
}

func ToCommand(t domains.Task, fs receivers.FileSystem) (command.Command, error) {
	switch t.Type {
	case string(command.CreateDir):
		return commands.NewCreateDirCommand(
			t.Name,
			receivers.NewOperatingSystemReceiver(fs),
			&t.Args,
			command.CreateDir,
			t.AbortOnFail,
		), nil
	case string(command.CreateFile):
		return commands.NewCreateFileCommand(
			t.Name,
			receivers.NewOperatingSystemReceiver(fs),
			&t.Args,
			command.CreateFile,
			t.AbortOnFail,
		), nil
	case string(command.PutContent):
		return commands.NewPutContentCommand(
			t.Name,
			receivers.NewOperatingSystemReceiver(fs),
			&t.Args,
			command.PutContent,
			t.AbortOnFail,
		), nil
	case string(command.RemoveDir):
		return commands.NewRemoveDirCommand(
			t.Name,
			receivers.NewOperatingSystemReceiver(fs),
			&t.Args,
			command.RemoveDir,
			t.AbortOnFail,
		), nil
	case string(command.RemoveFile):
		return commands.NewRemoveFileCommand(
			t.Name,
			receivers.NewOperatingSystemReceiver(fs),
			&t.Args,
			command.RemoveFile,
			t.AbortOnFail,
		), nil
	default:
		return nil, errors.New("invalid command type")
	}
}
