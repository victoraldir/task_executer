package receivers

import (
	"fmt"
	"os"
)

type putContentFuncHandler struct {
	params *map[string]string
	osFs   FileSystem
}

func NewPutContentFuncHandler(osFs FileSystem) FuncHandler {
	return &putContentFuncHandler{
		osFs: osFs,
	}
}

func (i *putContentFuncHandler) Handle() error {

	if err := i.Validate(); err != nil {
		return err
	}

	if err := i.PreCheckValidate(); err != nil {
		fmt.Println(err)
		return nil
	}

	flag := os.O_APPEND | os.O_WRONLY
	var file *os.File
	var err error

	if (*i.params)[AppendKey] == "true" {
		file, err = i.osFs.OpenFile((*i.params)[PathKey], flag, 0666)
	} else {
		file, err = i.osFs.Create((*i.params)[PathKey])
	}

	if err != nil {
		return err
	}

	defer file.Close()

	if _, err := file.WriteString((*i.params)[ContentKey]); err != nil {
		return err
	}

	return nil
}

func (i *putContentFuncHandler) SetParams(params *map[string]string) {
	i.params = params
}

func (i *putContentFuncHandler) Validate() error {

	if i.params == nil {
		return fmt.Errorf("params are required")
	}

	if (*i.params)[PathKey] == "" {
		return fmt.Errorf("path is required")
	}

	if (*i.params)[ContentKey] == "" {
		return fmt.Errorf("content is required")
	}

	return nil
}

func (i *putContentFuncHandler) PreCheckValidate() error {

	return nil
}
