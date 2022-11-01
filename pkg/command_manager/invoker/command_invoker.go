package invoker

import command "github.com/victoraldir/task_executer/pkg/command_manager"

type CommandInvoker struct {
	commands []command.Command
}

func NewCommandInvoker(commands []command.Command) CommandInvoker {
	return CommandInvoker{
		commands: commands,
	}
}

func (i CommandInvoker) Invoke() error {
	for _, command := range i.commands {
		err := command.Execute()
		if err != nil {
			return err
		}
	}

	return nil
}
