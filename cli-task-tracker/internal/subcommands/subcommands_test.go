package subcommands

import (
	"fmt"
	"testing"

	"cli-stack-tracker/internal/globaltypes"
)

type TaskImp struct {
	Id  string `json:"id"`
	Des string `json:"description"`
	Sts string `json:"status"`
}

func (t *TaskImp) ID() string                        { return t.Id }
func (t *TaskImp) Description() string               { return t.Des }
func (t *TaskImp) Status() string                    { return t.Sts }
func (t *TaskImp) SetID(id string)                   { t.Id = id }
func (t *TaskImp) SetDescription(description string) { t.Des = description }
func (t *TaskImp) SetStatus(status string)           { t.Sts = status }

type TaskListMock struct {
	description    string
	id             string
	newDescription string
}

func (t *TaskListMock) GetTasks() []globaltypes.ITask {
	return []globaltypes.ITask{
		&TaskImp{"id", "description", "approved"},
	}
}
func (t *TaskListMock) Add(description string) { t.description = description }
func (t *TaskListMock) Remove(id string)       { t.id = id }
func (t *TaskListMock) Update(id string, newDescription string) {
	t.id = id
	t.description = newDescription
}
func (t *TaskListMock) MarkAsDone(id string)    { t.id = id }
func (t *TaskListMock) MarkAsPending(id string) { t.id = id }

type PersistenceMock struct {
	data     string
	filename string
	path     string
}

func (p *PersistenceMock) Write(data any) {
	p.data = fmt.Sprintf("%v", data)
}
func (p *PersistenceMock) Remove() {}
func (p *PersistenceMock) Read() []byte {
	return []byte("")
}

func (p *PersistenceMock) ExtraInfo() (filename string, path string) {
	return p.filename, p.path
}

var didSucommandMockExecuted bool = false

type SubcommandsMock struct {
	task        globaltypes.ITaskList
	persistence globaltypes.IPersistence
}

func (s *SubcommandsMock) SetPersistence(p *globaltypes.IPersistence) { s.persistence = *p }
func (s *SubcommandsMock) SetTasks(tl *globaltypes.ITaskList)         { s.task = *tl }
func (s *SubcommandsMock) Command() string {
	return "mockCommand"
}

func (s *SubcommandsMock) Exec() {
	didSucommandMockExecuted = true
}

func NewSubcommandsParserMock() *SubcommandsParser {
	p := &PersistenceMock{}
	tl := &TaskListMock{}
	s := &SubcommandsMock{tl, p}
	subMap := map[string]globaltypes.ISubcommands{
		s.Command(): s,
	}

	return &SubcommandsParser{p, tl, subMap}
}

func TestExecute(t *testing.T) {
	subcommandParserMock := NewSubcommandsParserMock()
	subcommandParserMock.Execute("mockCommand")

	if !didSucommandMockExecuted {
		t.Error("MockCommand not executed when should")
	}

	didSucommandMockExecuted = false

	subcommandParserMock.Execute("invalid")
	if didSucommandMockExecuted {
		t.Error("MockCommand executed when should not")
	}
}
