package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davidtom/tarea/pkg/config"
	"github.com/davidtom/tarea/pkg/task"
)

// Set persists a task in the appropriate storage location
func Set(t *task.Task) int {
	var storage Storage

	// determine correct storage type based on config
	// TODO: implement some opts or config set up here, even if it only has one path
	// switch opts.storageType
	storage = new(LocalStorage)

	// persist task
	id := storage.SetTask(t)

	return id
}

// ********************
// LocalStorage Methods
// ********************

// SetTask persists a task locally
func (s *LocalStorage) SetTask(t *task.Task) int {
	var tasks []task.Task

	c := config.LoadConfig()

	readBs, readErr := ioutil.ReadFile(c.TaskFile)
	if readErr != nil {
		// TODO: error handling
		panic(readErr)
	}

	// TODO: this breaks unless there is a file there already, with at least []
	if unmarshErr := json.Unmarshal(readBs, &tasks); unmarshErr != nil {
		panic(unmarshErr)
	}

	// when adding a new task, its id will be automatically assigned based on the length of tasks
	id := len(tasks) + 1
	t.ID = id

	tasks = append(tasks, *t)

	writeBs, marshErr := json.Marshal(tasks)
	if marshErr != nil {
		panic(marshErr)
	}

	// TODO: understand this permission
	if writeErr := ioutil.WriteFile(c.TaskFile, writeBs, 0666); writeErr != nil {
		// TODO: error handling
		panic(writeErr)
	}

	return id
}
