package task

// Task represents a task used throughout the application
type Task struct {
	description string
	priority    string
}

// NewTask creates a pointer to a new task from argument and flags
// TODO: change second arg from string to flagSet
func NewTask(arg string, priority string) *Task {
	return &Task{
		description: arg,
		priority:    priority,
	}
}
