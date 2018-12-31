package storage

import (
	"fmt"

	"github.com/davidtom/tarea/pkg/task"
)

// Set persists a task in the appropriate storage location
func Set(t *task.Task) {
	var storage Storage

	// determine correct storage type based on config
	// TODO: implement some opts or config set up here, even if it only has one path
	// switch opts.storageType
	storage = new(LocalStorage)

	// persist task
	storage.SetTask(t)
}

// SetTask persists a task locally
func (s *LocalStorage) SetTask(t *task.Task) {
	fmt.Printf("Storing locally: %+v", t)
}
