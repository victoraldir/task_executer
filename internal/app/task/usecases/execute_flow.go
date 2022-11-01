package usecases

import (
	"fmt"

	"github.com/victoraldir/task_executer/internal/app/task/core/ports"
)

type ExecuteFlowUseCase interface {
	Execute() error
}

type executeFlowUseCase struct {
	taskLoaderRepo  ports.LoaderRepository
	taskInvokerRepo ports.InvokerRepository
}

func NewExecuteFlowUseCase(taskLoaderRepo ports.LoaderRepository, taskInvokerRepo ports.InvokerRepository) ExecuteFlowUseCase {
	return executeFlowUseCase{
		taskLoaderRepo:  taskLoaderRepo,
		taskInvokerRepo: taskInvokerRepo,
	}
}

func (e executeFlowUseCase) Execute() error {

	// Load tasks from repository
	tasks, err := e.taskLoaderRepo.LoadTasks()
	if err != nil {
		fmt.Println("Error loading tasks")
		return err
	}

	err = e.taskInvokerRepo.Invoke(*tasks)

	if err != nil {
		fmt.Println("Error invoking tasks")
		return err
	}

	return nil
}
