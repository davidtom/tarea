package task

// Task represents a task used throughout the application
type Task struct {
	Description string
	Priority    string
}

// NewTask creates a pointer to a new task from argument and flags
// TODO: change second arg from string to flagSet
// TODO: where do I define this? main? pkg?
func NewTask(arg string, priority string) *Task {
	return &Task{
		Description: arg,
		Priority:    priority,
	}
}
