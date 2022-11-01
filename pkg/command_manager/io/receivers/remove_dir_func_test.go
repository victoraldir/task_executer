package receivers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDirFuncHandler_Handle(t *testing.T) {

	t.Run("should remove directory", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		directoryPath := "testdata/to_be_deleted_no_content"
		os.RemoveAll(directoryPath)
		os.Mkdir(directoryPath, 0777)

		removeDirHandler := NewRemoveDirFuncHandler(fileSystemMock)
		removeDirHandler.SetParams(&map[string]string{
			PathKey:      directoryPath,
			RecursiveKey: "false",
		})

		// When
		err := removeDirHandler.Handle()

		// Then
		assert.Nil(t, err)
	})

	t.Run("should not remove directory. Directory not empty", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		directoryPath := "testdata/to_be_deleted_with_content"
		os.RemoveAll(directoryPath)
		os.Mkdir(directoryPath, 0777)
		os.Create(directoryPath + "/test1.txt")
		os.Create(directoryPath + "/test2.txt")

		removeDirHandler := NewRemoveDirFuncHandler(fileSystemMock)
		removeDirHandler.SetParams(&map[string]string{
			PathKey:      directoryPath,
			RecursiveKey: "false",
		})

		// When
		err := removeDirHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "remove "+directoryPath+": directory not empty", err.Error())
	})

	t.Run("should remove directory recursively", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		directoryPath := "testdata/to_be_deleted_with_content"
		os.RemoveAll(directoryPath)
		os.Mkdir(directoryPath, 0777)
		os.Create(directoryPath + "/test1.txt")
		os.Create(directoryPath + "/test2.txt")

		removeDirHandler := NewRemoveDirFuncHandler(fileSystemMock)
		removeDirHandler.SetParams(&map[string]string{
			PathKey:      directoryPath,
			RecursiveKey: "true",
		})

		// When
		err := removeDirHandler.Handle()

		// Then
		assert.Nil(t, err)
	})
}
