package main

import (
	"github.com/victoraldir/task_executer/internal/app"
)

func main() {

	app := app.NewApplication()

	// Use case
	err := app.ExecuteFlowUseCase.Execute()

	if err != nil {
		panic(err)
	}
}
