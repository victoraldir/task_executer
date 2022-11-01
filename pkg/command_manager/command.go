package command

type CommandType string

const (
	CreateDir  CommandType = "create_dir"
	PutContent CommandType = "put_content"
	RemoveDir  CommandType = "rm_dir"
	CreateFile CommandType = "create_file"
	RemoveFile CommandType = "rm_file"
)

//go:generate mockgen -destination=../command_manager/mocks/mockCommand.go -package=command github.com/victoraldir/task_executer/pkg/command_manager Command
type Command interface {
	Execute() error
}
