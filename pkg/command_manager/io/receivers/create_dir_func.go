package receivers

import (
	"fmt"
	"os"
)

type createDirFuncHanlder struct {
	params *map[string]string
	osFs   FileSystem
}

func NewCreateDirFuncHandler(osFs FileSystem) FuncHandler {
	return &createDirFuncHanlder{
		osFs: osFs,
	}
}

func (i *createDirFuncHanlder) Handle() error {

	if err := i.Validate(); err != nil {
		return err
	}

	if err := i.PreCheckValidate(); err != nil {
		return nil
	}

	if err := i.osFs.Mkdir((*i.params)[PathKey], os.ModePerm); err != nil {
		return err
	}
	return nil
}

func (i *createDirFuncHanlder) SetParams(params *map[string]string) {
	i.params = params
}

func (i *createDirFuncHanlder) Validate() error {

	if i.params == nil {
		return fmt.Errorf("params are required")
	}

	if (*i.params)[PathKey] == "" {
		return fmt.Errorf("path is required")
	}

	return nil
}

func (i *createDirFuncHanlder) PreCheckValidate() error {

	// Check if the dir already exists
	if _, err := i.osFs.Stat((*i.params)[PathKey]); err == nil {
		return fmt.Errorf("dir already exists %s", (*i.params)[PathKey])
	}

	return nil
}
