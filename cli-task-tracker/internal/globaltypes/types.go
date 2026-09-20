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
	Remove(id string)
	Update(id string, newDescription string)
	MarkAsDone(id string)
	MarkAsPending(id string)
}

type IPersistence interface {
	Write(data any)
	Remove()
	Read() []byte
	ExtraInfo() (filename string, path string)
}

type ISubcommands interface {
	SetPersistence(p *IPersistence)
	SetTasks(tl *ITaskList)

	Command() string
	Exec()
}

type IFileSystem interface {
	Write(path string, name string) error

	RemoveAll(path string) error

	Getwd() (dir string, err error)

	ReadFile(completePath string) ([]byte, error)
}
