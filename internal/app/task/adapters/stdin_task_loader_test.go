package adapters

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStdinTaskLoaderRepository_LoadTasks(t *testing.T) {
	t.Run("should load tasks", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		reader, _ := os.Open("testdata/tasks.yaml")
		taskLoaderRepo := NewStdinTaskLoaderRepository(reader)

		// When
		tasks, err := taskLoaderRepo.LoadTasks()

		// Then
		assert.Nil(t, err)
		assert.Equal(t, 1, len(*tasks))
		assert.Equal(t, "Create root directory", (*tasks)[0].Name)
		assert.Equal(t, "create_dir", (*tasks)[0].Type)
	})

	t.Run("should not load tasks. Invalid yaml", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		reader, _ := os.Open("testdata/tasks_invalid.yaml")
		taskLoaderRepo := NewStdinTaskLoaderRepository(reader)

		// When
		tasks, err := taskLoaderRepo.LoadTasks()

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "yaml: unmarshal errors:\n  line 1: cannot unmarshal !!str `test` into []domains.Task", err.Error())
		assert.Nil(t, tasks)
	})
}
