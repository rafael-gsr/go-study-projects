// Package globaltypes contains the program types
package globaltypes

type ITask struct {
	ID          string
	Description string
	Status      string
}

type ITaskList interface {
	GetTasks() []ITask
	Add(description string)
	Remove(description string)
	Update(oldDescription string, newDescription string)
}

type IPersistence interface {
	Write(data any)
	Remove()
	Read() []byte
}
