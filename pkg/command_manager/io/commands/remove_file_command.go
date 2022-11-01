package commands

import (
	"fmt"

	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

type removeFileCommand struct {
	name        string
	receiver    command.Receiver
	params      *map[string]string
	commandType command.CommandType
	abortOnFail bool
}

func NewRemoveFileCommand(
	name string,
	receiver command.Receiver,
	params *map[string]string,
	commandType command.CommandType,
	abortOnFail bool,
) command.Command {
	return removeFileCommand{
		name:        name,
		params:      params,
		receiver:    receiver,
		commandType: commandType,
		abortOnFail: abortOnFail,
	}
}

func (c removeFileCommand) Execute() error {

	fmt.Println("Executing command: ", c.name)

	err := c.receiver.RemoveFile(c.params)
	if err != nil && c.abortOnFail {
		fmt.Println("Aborting execution")
		return err
	}

	return nil
}
