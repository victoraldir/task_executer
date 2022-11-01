package commands

import (
	"fmt"

	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

type createDirCommand struct {
	receiver    command.Receiver
	name        string
	params      *map[string]string
	commandType command.CommandType
	abortOnFail bool
}

func NewCreateDirCommand(
	name string,
	receiver command.Receiver,
	params *map[string]string,
	commandType command.CommandType,
	abortOnFail bool) command.Command {

	return createDirCommand{
		name:        name,
		params:      params,
		receiver:    receiver,
		commandType: commandType,
		abortOnFail: abortOnFail,
	}
}

func (c createDirCommand) Execute() error {
	fmt.Println("Executing command: ", c.name)
	err := c.receiver.CreateDir(c.params)

	if err != nil && c.abortOnFail {
		fmt.Println("Aborting execution")
		return err
	}

	return nil
}
