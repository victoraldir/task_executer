package receivers

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestOperatingSystemReceiver_CreateDir(t *testing.T) {

	t.Run("should create dir", func(t *testing.T) {
		// given

		teardown := setup(t)
		defer teardown()

		receiver := NewOperatingSystemReceiver(fileSystemMock)
		// when

		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, errors.New("error"))
		fileSystemMock.EXPECT().Mkdir(gomock.Any(), gomock.Any()).Return(nil)

		err := receiver.CreateDir(&map[string]string{
			PathKey: "/test",
		})

		// then
		assert.Nil(t, err)
	})

	t.Run("should not create dir. Path required", func(t *testing.T) {
		// given
		teardown := setup(t)
		defer teardown()

		receiver := NewOperatingSystemReceiver(fileSystemMock)
		// when

		err := receiver.CreateDir(nil)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})

	t.Run("should not create dir. Error creating directory", func(t *testing.T) {
		// given

		teardown := setup(t)
		defer teardown()

		receiver := NewOperatingSystemReceiver(fileSystemMock)
		// when

		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, errors.New("error"))
		fileSystemMock.EXPECT().Mkdir(gomock.Any(), gomock.Any()).Return(errors.New("error"))

		err := receiver.CreateDir(&map[string]string{
			PathKey: "/test",
		})

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "error", err.Error())
	})
}

func TestOperatingSystemReceiver_CreateFile(t *testing.T) {

	t.Run("should not create file. Path required", func(t *testing.T) {
		// given
		teardown := setup(t)
		defer teardown()

		receiver := NewOperatingSystemReceiver(fileSystemMock)
		// when

		err := receiver.CreateFile(nil)

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})

	t.Run("should not create file. Error creating file", func(t *testing.T) {
		// given
		teardown := setup(t)
		defer teardown()

		receiver := NewOperatingSystemReceiver(fileSystemMock)

		// when
		fileSystemMock.EXPECT().Stat(gomock.Any()).Return(nil, errors.New("error"))

		err := receiver.CreateFile(&map[string]string{
			PathKey: "/test",
		})

		// then
		assert.NotNil(t, err)
		assert.Equal(t, "directory  does not exist", err.Error())
	})
}
