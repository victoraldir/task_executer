package commands

import (
	"fmt"

	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

type putContentCommand struct {
	name        string
	receiver    command.Receiver
	params      *map[string]string
	commandType command.CommandType
	abortOnFail bool
}

func NewPutContentCommand(
	name string,
	receiver command.Receiver,
	params *map[string]string,
	commandType command.CommandType,
	abortOnFail bool,
) command.Command {
	return putContentCommand{
		name:        name,
		params:      params,
		receiver:    receiver,
		commandType: commandType,
		abortOnFail: abortOnFail,
	}
}

func (c putContentCommand) Execute() error {

	fmt.Println("Executing command: ", c.name)

	err := c.receiver.PutContent(c.params)
	if err != nil && c.abortOnFail {
		fmt.Println("Aborting execution")
		return err
	}

	return nil
}
