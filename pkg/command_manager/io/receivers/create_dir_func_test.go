package receivers

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDirFuncHanlder_Handle(t *testing.T) {

	t.Run("should fail when creating dir. Params missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createDirFuncHanlder := NewCreateDirFuncHandler(fileSystemMock)

		// When
		err := createDirFuncHanlder.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})

	t.Run("should fail when creating dir. Path param missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createDirFuncHanlder := NewCreateDirFuncHandler(fileSystemMock)
		createDirFuncHanlder.SetParams(&map[string]string{
			ContentKey: "test",
		})

		// When
		err := createDirFuncHanlder.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "path is required", err.Error())
	})

	t.Run("should create directory", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createDirFuncHanlder := NewCreateDirFuncHandler(fileSystemMock)
		createDirFuncHanlder.SetParams(&map[string]string{
			PathKey: "/tmp/test",
		})

		fileSystemMock.EXPECT().Stat("/tmp/test").Return(nil, os.ErrNotExist)
		fileSystemMock.EXPECT().Mkdir("/tmp/test", os.ModePerm).Return(nil).Times(1)

		// When
		err := createDirFuncHanlder.Handle()

		// Then
		assert.Nil(t, err)
	})

	t.Run("Should PreCheck fail when dir already exists", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createDirFuncHanlder := NewCreateDirFuncHandler(fileSystemMock)
		createDirFuncHanlder.SetParams(&map[string]string{
			PathKey: "/tmp/test",
		})

		fileSystemMock.EXPECT().Stat("/tmp/test").Return(nil, nil)

		// When
		err := createDirFuncHanlder.Handle()

		// Then
		assert.Nil(t, err)
	})

}
