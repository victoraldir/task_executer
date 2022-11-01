package receivers

import (
	"fmt"
	"os"
)

type removeDirFuncHandler struct {
	params *map[string]string
	osFs   FileSystem
}

func NewRemoveDirFuncHandler(osFs FileSystem) FuncHandler {
	return &removeDirFuncHandler{
		osFs: osFs,
	}
}

func (i *removeDirFuncHandler) Handle() error {

	if err := i.Validate(); err != nil {
		return err
	}

	if err := i.PreCheckValidate(); err != nil {
		fmt.Println(err)
		return nil
	}

	// Remove directory recursively
	if (*i.params)[RecursiveKey] == "true" {
		if err := os.RemoveAll((*i.params)[PathKey]); err != nil {
			return err
		}
	} else {
		if err := os.Remove((*i.params)[PathKey]); err != nil {
			return err
		}
	}

	return nil
}

func (i *removeDirFuncHandler) SetParams(params *map[string]string) {
	i.params = params
}

func (i *removeDirFuncHandler) Validate() error {

	if i.params == nil {
		return fmt.Errorf("params are required")
	}

	if (*i.params)[PathKey] == "" {
		return fmt.Errorf("path is required")
	}

	return nil
}

func (i *removeDirFuncHandler) PreCheckValidate() error {

	return nil
}
