package receivers

import (
	"fmt"
	"os"

	command "github.com/victoraldir/task_executer/pkg/command_manager"
)

const (
	PathKey      string = "path"
	ContentKey   string = "content"
	AppendKey    string = "append"
	RecursiveKey string = "recursive"
)

//go:generate mockgen -destination=../receivers/mocks/mockFileSystem.go -package=receivers github.com/victoraldir/task_executer/pkg/command_manager/io/receivers FileSystem
type FileSystem interface {
	Stat(name string) (os.FileInfo, error)
	Mkdir(name string, perm os.FileMode) error
	Create(name string) (*os.File, error)
	IsNotExist(err error) bool
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
	RemoveAll(path string) error
	Remove(name string) error
}

type OperatingSystemReceiver struct {
	createDir  FuncHandler
	createFile FuncHandler
	putContent FuncHandler
	removeFile FuncHandler
	removeDir  FuncHandler
}

type Params struct {
	Path string
}

type FuncHandler interface {
	SetParams(params *map[string]string)
	Validate() error
	PreCheckValidate() error
	Handle() error
}

// osFS implements fileSystem using the local disk.
type osFS struct{}

func (osFS) Stat(name string) (os.FileInfo, error)     { return os.Stat(name) }
func (osFS) Mkdir(name string, perm os.FileMode) error { return os.Mkdir(name, perm) }
func (osFS) Create(name string) (*os.File, error)      { return os.Create(name) }
func (osFS) IsNotExist(err error) bool                 { return os.IsNotExist(err) }
func (osFS) RemoveAll(path string) error               { return os.RemoveAll(path) }
func (osFS) Remove(name string) error                  { return os.Remove(name) }
func (osFS) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)

}

func NewOsFS() FileSystem {
	return &osFS{}
}

func NewOperatingSystemReceiver(fs FileSystem) command.Receiver {
	return OperatingSystemReceiver{
		createDir:  NewCreateDirFuncHandler(fs),
		createFile: NewCreateFileFuncHandler(fs),
		putContent: NewPutContentFuncHandler(fs),
		removeFile: NewRemoveFileFuncHandler(fs),
		removeDir:  NewRemoveDirFuncHandler(fs),
	}
}

func (i OperatingSystemReceiver) CreateDir(params *map[string]string) error {

	i.createDir.SetParams(params)

	err := i.createDir.Handle()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	return err
}

func (i OperatingSystemReceiver) CreateFile(params *map[string]string) error {

	i.createFile.SetParams(params)

	err := i.createFile.Handle()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	return nil
}

func (i OperatingSystemReceiver) RemoveFile(params *map[string]string) error {

	i.removeFile.SetParams(params)

	err := i.removeFile.Handle()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	return nil
}

func (i OperatingSystemReceiver) RemoveDir(params *map[string]string) error {

	i.removeDir.SetParams(params)

	err := i.removeDir.Handle()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	return nil
}

func (i OperatingSystemReceiver) PutContent(params *map[string]string) error {

	i.putContent.SetParams(params)

	err := i.putContent.Handle()

	if err != nil {
		fmt.Printf("Error: %s", err)
		return err
	}

	return nil

}
