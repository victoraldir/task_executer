package receivers

import (
	"fmt"
	"os"
)

type removeFileFuncHandler struct {
	params *map[string]string
	osFs   FileSystem
}

func NewRemoveFileFuncHandler(osFs FileSystem) FuncHandler {
	return &removeFileFuncHandler{
		osFs: osFs,
	}
}

func (i *removeFileFuncHandler) Handle() error {

	if err := i.Validate(); err != nil {
		return err
	}

	if err := i.PreCheckValidate(); err != nil {
		fmt.Println(err)
		return nil
	}

	err := os.Remove((*i.params)[PathKey])
	if err != nil {
		return err
	}

	return nil
}

func (i *removeFileFuncHandler) SetParams(params *map[string]string) {
	i.params = params
}

func (i *removeFileFuncHandler) Validate() error {

	if i.params == nil {
		return fmt.Errorf("params are required")
	}

	if (*i.params)[PathKey] == "" {
		return fmt.Errorf("path is required")
	}

	return nil
}

func (i *removeFileFuncHandler) PreCheckValidate() error {

	if _, err := os.Stat((*i.params)[PathKey]); os.IsNotExist(err) {
		return fmt.Errorf("path %s does not exist", (*i.params)[PathKey])
	}

	return nil
}
