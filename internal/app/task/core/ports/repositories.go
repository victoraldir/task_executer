package ports

import "github.com/victoraldir/task_executer/internal/app/task/core/domains"

//go:generate mockgen -destination=../ports/mocks/mockLoaderRepository.go -package=ports github.com/victoraldir/task_executer/internal/app/task/core/ports LoaderRepository
type LoaderRepository interface {
	LoadTasks() (*[]domains.Task, error)
}

//go:generate mockgen -destination=../ports/mocks/mockInvokerRepository.go -package=ports github.com/victoraldir/task_executer/internal/app/task/core/ports InvokerRepository
type InvokerRepository interface {
	Invoke([]domains.Task) error
}
