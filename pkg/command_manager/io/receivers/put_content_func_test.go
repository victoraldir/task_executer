package receivers

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestPutContentFuncHandler_Handle(t *testing.T) {

	t.Run("should fail when putting content to a file. Params missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		putContentHandler := NewPutContentFuncHandler(fileSystemMock)

		// When
		err := putContentHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})

	t.Run("should fail when putting content to a file. Path param missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		putContentHandler := NewPutContentFuncHandler(fileSystemMock)
		putContentHandler.SetParams(&map[string]string{
			ContentKey: "test",
		})

		// When
		err := putContentHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "path is required", err.Error())
	})

	t.Run("should put content to a file", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		directory := "testdata"
		fileName := "test.txt"
		filePath := directory + "/" + fileName
		os.RemoveAll(directory)
		os.Mkdir(directory, os.ModePerm)
		file, _ := os.Create(filePath)

		putContentHandler := NewPutContentFuncHandler(fileSystemMock)
		putContentHandler.SetParams(&map[string]string{
			PathKey:    filePath,
			ContentKey: "test",
		})

		fileSystemMock.EXPECT().Create(gomock.Any()).Return(file, nil).Times(1)

		// When
		err := putContentHandler.Handle()

		// Then
		assert.Nil(t, err)
	})

}
