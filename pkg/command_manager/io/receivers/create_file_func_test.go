package receivers

import (
	"errors"
	"io/fs"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateFileFuncHandler_Handle(t *testing.T) {

	t.Run("should fail when creating file. Params missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createFileFuncHandler := NewCreateFileFuncHandler(fileSystemMock)

		// When
		err := createFileFuncHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})

	t.Run("should fail when creating file. Path param missig", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createFileFuncHandler := NewCreateFileFuncHandler(fileSystemMock)
		createFileFuncHandler.SetParams(&map[string]string{
			ContentKey: "test",
		})

		// When
		err := createFileFuncHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "path is required", err.Error())
	})

	t.Run("should create file", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createFileFuncHandler := NewCreateFileFuncHandler(fileSystemMock)
		createFileFuncHandler.SetParams(&map[string]string{
			PathKey: "/tmp/test.txt",
		})

		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, nil)
		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, errors.New("file not found"))
		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, fs.ErrNotExist)
		fileSystemMock.EXPECT().Create(gomock.Any()).Return(nil, nil).Times(1)

		// When
		err := createFileFuncHandler.Handle()

		// Then
		assert.Nil(t, err)
	})

	t.Run("Should PreCheck fail when dir already exists", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		createFileFuncHandler := NewCreateFileFuncHandler(fileSystemMock)
		createFileFuncHandler.SetParams(&map[string]string{
			PathKey: "/tmp/test",
		})

		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, nil)
		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, nil)

		// When
		err := createFileFuncHandler.Handle()

		// Then
		assert.Nil(t, err)
	})

}
