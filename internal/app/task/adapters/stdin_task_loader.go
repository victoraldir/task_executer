package adapters

import (
	"fmt"
	"io"

	"github.com/victoraldir/task_executer/internal/app/task/core/domains"
	"github.com/victoraldir/task_executer/internal/app/task/core/ports"
	"gopkg.in/yaml.v2"
)

//go:generate mockgen -destination=../adapters/mocks/mockReader.go -package=adapters github.com/victoraldir/task_executer/internal/app/task/adapters Reader
type Reader interface {
	Read(p []byte) (n int, err error)
}

type stdinTaskLoaderRepository struct {
	ioReader io.Reader
}

func NewStdinTaskLoaderRepository(reader Reader) ports.LoaderRepository {
	return stdinTaskLoaderRepository{
		ioReader: reader,
	}
}

func (y stdinTaskLoaderRepository) LoadTasks() (*[]domains.Task, error) {
	bytes, err := ReadStdin(y.ioReader)

	if err != nil {
		return nil, err
	}

	var tasks []domains.Task

	err = yaml.Unmarshal(bytes, &tasks)

	if err != nil {
		return nil, err
	}

	return &tasks, nil

}

func ReadStdin(ioReader io.Reader) ([]byte, error) {
	bytes, err := io.ReadAll(ioReader)

	if err != nil {
		fmt.Println("Error reading workflow file: ", err)
		return nil, err
	}

	return bytes, nil
}
