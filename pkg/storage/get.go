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

// SetTask persists a task locally
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

	return tasks

	// fmt.Println(tasks)
	// tasks = append(tasks, *t)
	// fmt.Println(tasks)

	// writeBs, marshErr := json.Marshal(tasks)
	// if marshErr != nil {
	// 	panic(marshErr)
	// }

	// // TODO: understand this permission
	// if writeErr := ioutil.WriteFile(c.TaskFile, writeBs, 0666); writeErr != nil {
	// 	panic(writeErr)
	// }
	// // errr := ioutil.WriteFile(c.TaskFile, bs, 0666)
	// // if err != nil {
	// // 	// TODO: error handling
	// // 	panic(errr)
	// // }
}

func readFile(filename string) *[]byte {
	// TODO: where should this file be saved?
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		// TODO: handle errors better?
		panic(err)
	}

	return &bs
}
