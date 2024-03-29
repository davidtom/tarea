package storage

import (
	"github.com/davidtom/tarea/pkg/task"
)

// Storage interfaces abstract storage methods away from storage location (ie local or remote)
type Storage interface {
	GetTask() []task.Task
	SetTask(t *task.Task) int
	// update()
	// delete()
}

// LocalStorage struct interacts with tasks stored locally
type LocalStorage struct{}
