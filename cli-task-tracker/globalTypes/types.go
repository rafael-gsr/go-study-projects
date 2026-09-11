// Package globaltypes contains the program types
package globaltypes

type ITask struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type ITaskList interface {
	GetTasks() []ITask
	Add(description string)
	Remove(description string)
	Update(id string, newDescription string)
	MarkAsDone(id string)
	MarkAsPenging(id string)
}

type IPersistence interface {
	Write(data any)
	Remove()
	Read() []byte
}
