package task

import (
	"encoding/json"
	"fmt"

	"cli-stack-tracker/internal/globaltypes"

	"github.com/google/uuid"
)

type Task struct {
	Id   string `json:"id"`
	Desc string `json:"description"`
	Stat string `json:"status"`
}

func (t *Task) ID() string                        { return t.Id }
func (t *Task) SetID(id string)                   { t.Id = id }
func (t *Task) Description() string               { return t.Desc }
func (t *Task) SetDescription(description string) { t.Desc = description }
func (t *Task) Status() string                    { return t.Stat }
func (t *Task) SetStatus(status string)           { t.Stat = status }

type TaskList struct {
	tasks []globaltypes.ITask
}

func (tl *TaskList) Add(description string) {
	taskID := uuid.New()
	newTask := Task{taskID.String(), description, "Pending"}

	tl.tasks = append(tl.tasks, &newTask)
}

func (tl *TaskList) Remove(id string) {
	for idx, task := range tl.tasks {
		if task.ID() == id {
			tl.tasks = append(tl.tasks[:idx], tl.tasks[idx+1:]...)
			return
		}
	}
}

func (tl *TaskList) Update(id string, newDescription string) {
	for idx, task := range tl.tasks {
		if task.ID() == id {
			tl.tasks[idx].SetDescription(newDescription)
			return
		}
	}
}

func (tl *TaskList) MarkAsDone(id string) {
	for idx, task := range tl.tasks {
		if task.ID() == id {
			tl.tasks[idx].SetStatus("Done")
			return
		}
	}
}

func (tl *TaskList) MarkAsPending(id string) {
	for idx, task := range tl.tasks {
		if task.ID() == id {
			tl.tasks[idx].SetStatus("Pending")
			return
		}
	}
}

func (tl *TaskList) GetTasks() []globaltypes.ITask {
	return tl.tasks
}

func NewTaskList(rawData []byte) (*TaskList, error) {
	var tl TaskList
	emptyList := &TaskList{[]globaltypes.ITask{}}

	if len(rawData) == 0 {
		return emptyList, nil
	}

	err := json.Unmarshal(rawData, &tl.tasks)
	if err != nil {
		fmt.Println("Error to retrieve the stored data", err)
		return emptyList, err
	}

	return &tl, nil
}
