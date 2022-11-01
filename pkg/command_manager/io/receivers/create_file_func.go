package receivers

import (
	"fmt"
	"os"
	"strings"
)

type createFileFuncHandler struct {
	params *map[string]string
	osFs   FileSystem
}

func NewCreateFileFuncHandler(osFs FileSystem) FuncHandler {
	return &createFileFuncHandler{
		osFs: osFs,
	}
}

func (i *createFileFuncHandler) Handle() error {

	if err := i.Validate(); err != nil {
		return err
	}

	if err := i.PreCheckValidate(); err != nil {
		return nil
	}

	if _, err := i.osFs.Stat((*i.params)[PathKey]); os.IsNotExist(err) {
		i.osFs.Create((*i.params)[PathKey])
	}

	return nil
}

func (i *createFileFuncHandler) SetParams(params *map[string]string) {
	i.params = params
}

func (i *createFileFuncHandler) Validate() error {

	if i.params == nil {
		return fmt.Errorf("params are required")
	}

	if (*i.params)[PathKey] == "" {
		return fmt.Errorf("path is required")
	}

	dirSplit := strings.Split((*i.params)[PathKey], "/")
	directory := strings.Join(dirSplit[:len(dirSplit)-1], "/")

	if _, err := i.osFs.Stat(directory); err != nil {
		return fmt.Errorf(fmt.Sprintf("directory %s does not exist", directory))
	}

	return nil
}

func (i *createFileFuncHandler) PreCheckValidate() error {

	if _, err := i.osFs.Stat((*i.params)[PathKey]); err == nil {
		return fmt.Errorf("file already exists %s", (*i.params)[PathKey])
	}

	return nil
}
