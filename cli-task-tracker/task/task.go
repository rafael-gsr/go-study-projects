package task

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type Task struct {
	ID          string `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

type TaskList struct {
	tasks []Task
}

func (tl *TaskList) Add(description string) {
	taskID := uuid.New()
	newTask := Task{ID: taskID.String(), Description: description, Status: "Pending"}

	tl.tasks = append(tl.tasks, newTask)
}

func (tl *TaskList) Remove(description string) {
	for idx, task := range tl.tasks {
		if task.Description == description {
			tl.tasks = append(tl.tasks[:idx], tl.tasks[idx+1:]...)
			return
		}
	}
}

func (tl *TaskList) Update(oldDescription string, newDescription string) {
	for idx, task := range tl.tasks {
		if task.Description == oldDescription {
			tl.tasks[idx].Description = newDescription
			return
		}
	}
}

func NewTaskList(rawData []byte) *TaskList {
	var tl TaskList
	err := json.Unmarshal(rawData, tl)
	if err != nil {
		fmt.Println("Error to retrieve the stored data")
		return &TaskList{[]Task{}}
	}

	return &tl
}
