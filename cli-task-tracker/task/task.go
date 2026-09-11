package task

import (
	"encoding/json"
	"fmt"

	globaltypes "cli-stack-tracker/globalTypes"

	"github.com/google/uuid"
)

type TaskList struct {
	tasks []globaltypes.ITask
}

func (tl *TaskList) Add(description string) {
	taskID := uuid.New()
	newTask := globaltypes.ITask{ID: taskID.String(), Description: description, Status: "Pending"}

	tl.tasks = append(tl.tasks, newTask)
}

func (tl *TaskList) Remove(id string) {
	for idx, task := range tl.tasks {
		if task.ID == id {
			tl.tasks = append(tl.tasks[:idx], tl.tasks[idx+1:]...)
			return
		}
	}
}

func (tl *TaskList) Update(id string, newDescription string) {
	for idx, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[idx].Description = newDescription
			return
		}
	}
}

func (tl *TaskList) MarkAsDone(id string) {
	for idx, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[idx].Status = "Done"
			return
		}
	}
}

func (tl *TaskList) MarkAsPenging(id string) {
	for idx, task := range tl.tasks {
		if task.ID == id {
			tl.tasks[idx].Status = "Pending"
			return
		}
	}
}

func (tl *TaskList) GetTasks() []globaltypes.ITask {
	return tl.tasks
}

func NewTaskList(rawData []byte) *TaskList {
	var tl TaskList
	err := json.Unmarshal(rawData, &tl.tasks)
	if err != nil {
		fmt.Println("Error to retrieve the stored data", err)
		return &TaskList{[]globaltypes.ITask{}}
	}

	return &tl
}
