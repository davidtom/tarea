package storage

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davidtom/tarea/pkg/config"
	"github.com/davidtom/tarea/pkg/task"
)

// Get returns a list of tasks
func Get() []task.Task {
	var storage Storage

	// determine correct storage type based on config
	// TODO: implement some opts or config set up here, even if it only has one path
	// switch opts.storageType
	// TODO: abstract this method so it can be used in all set/get/etc.
	storage = new(LocalStorage)

	return storage.GetTask()
}

// ********************
// LocalStorage Methods
// ********************

// GetTask returns a list of all persisted tasks
func (s *LocalStorage) GetTask() []task.Task {
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

	// TODO: reassign ids at each read to keep them updated (see task warrior's, quick demo)
	// this means - I think - that each time we read the tasks for list, we need to update the ids
	// and re-save; this ensures that the ids that a user sees stay valid, so that they can delete/complete
	// multiple tasks without needing to figure out new ids
	// I can make this entire process easier by abstracting the read/write operation
	// maybe something like tasks.ReadFromFile(filename), tasks.WriteToFile(filename)
	for i, t := range tasks {
		t.ID = i + 1
		tasks[i] = t
	}

	writeBs, marshErr := json.Marshal(tasks)
	if marshErr != nil {
		panic(marshErr)
	}

	// TODO: understand this permission
	if writeErr := ioutil.WriteFile(c.TaskFile, writeBs, 0666); writeErr != nil {
		// TODO: error handling
		panic(writeErr)
	}

	return tasks
}
