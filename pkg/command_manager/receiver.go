package command

//go:generate mockgen -destination=../command_manager/mocks/mockReceiver.go -package=command github.com/victoraldir/task_executer/pkg/command_manager Receiver
type Receiver interface {
	CreateDir(params *map[string]string) error
	CreateFile(params *map[string]string) error
	RemoveFile(params *map[string]string) error
	RemoveDir(params *map[string]string) error
	PutContent(params *map[string]string) error
}
