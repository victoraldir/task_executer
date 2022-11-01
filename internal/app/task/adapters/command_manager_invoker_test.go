package adapters

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/victoraldir/task_executer/internal/app/task/core/domains"
	"github.com/victoraldir/task_executer/pkg/command_manager/io/receivers"
)

func TestCommandManagerInvokerRepository_Invoke(t *testing.T) {
	teardown := setup(t)
	defer teardown()

	t.Run("should invoke the commands", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		tasks := []domains.Task{
			{
				Name:        "test",
				Type:        "create_dir",
				AbortOnFail: true,
				Args: map[string]string{
					"path": "/tmp/test",
				},
			},
		}

		invokerRepo := NewCommandManagerInvokerRepository(receivers.NewOsFS())

		// When
		err := invokerRepo.Invoke(tasks)

		// Then
		assert.Nil(t, err)
	})

	t.Run("should not invoke command. Invalid task type", func(t *testing.T) {
		// Given
		teardown := setup(t)
		defer teardown()

		tasks := []domains.Task{
			{
				Name:        "test",
				Type:        "type_not_found",
				AbortOnFail: true,
				Args: map[string]string{
					"path": "/tmp",
				},
			},
		}

		invokerRepo := NewCommandManagerInvokerRepository(receivers.NewOsFS())

		// When
		err := invokerRepo.Invoke(tasks)

		// Then
		assert.NotNil(t, err)
		assert.Equal(t, "invalid command type", err.Error())
	})
}
