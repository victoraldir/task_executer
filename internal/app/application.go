package app

import (
	"os"

	"github.com/victoraldir/task_executer/internal/app/task/adapters"
	"github.com/victoraldir/task_executer/internal/app/task/usecases"
	"github.com/victoraldir/task_executer/pkg/command_manager/io/receivers"
)

type application struct {
	ExecuteFlowUseCase usecases.ExecuteFlowUseCase
}

func NewApplication() application {

	// File system
	fs := receivers.NewOsFS()

	//Adapters
	commandManagerInvokerRepo := adapters.NewCommandManagerInvokerRepository(fs)
	stdinTaskLoaderRepo := adapters.NewStdinTaskLoaderRepository(os.Stdin)

	// Use cases
	executeFlowUseCase := usecases.NewExecuteFlowUseCase(stdinTaskLoaderRepo, commandManagerInvokerRepo)

	return application{
		ExecuteFlowUseCase: executeFlowUseCase,
	}
}
