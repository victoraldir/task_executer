package receivers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveFileFuncHandler_Handle(t *testing.T) {

	t.Run("should remove file", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		removeFileHandler := NewRemoveFileFuncHandler(fileSystemMock)

		// When
		removeFileHandler.SetParams(&map[string]string{
			PathKey: "testdata/to_be_deleted.txt",
		})

		err := removeFileHandler.Handle()

		// Then
		assert.Nil(t, err)
	})

	t.Run("should not remove file. Params required", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		removeFileHandler := NewRemoveFileFuncHandler(fileSystemMock)

		// When
		err := removeFileHandler.Handle()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "params are required", err.Error())
	})
}
