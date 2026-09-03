package main

import (
	"reflect"
	"testing"
)

var (
	description        = "description"
	anotherDescription = "anotherDescription"
)

func TestNewTaskList(t *testing.T) {
	taskList := NewTaskList()
	taskListInstance := &TaskList{}

	if reflect.TypeOf(taskList) != reflect.TypeOf(taskListInstance) {
		t.Errorf("The type of NewTaskList is Wrong. \n Expected: %s \n Received: %s", reflect.TypeOf(taskList), reflect.TypeOf(taskListInstance))
	}

	if taskList.tasks == nil {
		t.Error("The internal task element is nil")
	}

	taskArrayInstance := []Task{}

	if reflect.TypeOf(taskList.tasks) != reflect.TypeOf(taskArrayInstance) {
		t.Error("Tasks types is incorrect")
	}
}

func countDescriptionOcurrencies(taskList *TaskList, description string) int {
	countOfMatchDescriptions := 0

	for _, task := range taskList.tasks {
		if task.Description == description {
			countOfMatchDescriptions++
		}
	}

	return countOfMatchDescriptions
}

func TestTaskListAddMethod(t *testing.T) {
	taskList := NewTaskList()

	taskList.Add(description)
	taskList.Add(description)
	taskList.Add(anotherDescription)

	isLenWrong := len(taskList.tasks) != 3
	isCountWrong := countDescriptionOcurrencies(taskList, description) != 2

	if isLenWrong || isCountWrong {
		t.Error("The Add method is not working properly")
	}
}

func TestTaskListRemoveMethod(t *testing.T) {
	taskList := NewTaskList()

	taskList.Add(description)
	taskList.Add(description)
	taskList.Add(anotherDescription)

	taskList.Remove(description)
	taskList.Remove(anotherDescription)

	isDescriptionCountNotDecreasing := countDescriptionOcurrencies(taskList, description) != 1
	isAnotherDescriptionCountNotDecreasing := countDescriptionOcurrencies(taskList, anotherDescription) != 0

	if isDescriptionCountNotDecreasing || isAnotherDescriptionCountNotDecreasing {
		t.Errorf("The Remove method is not working properly \n Description count: %d \n anotherDescription count: %d \n", countDescriptionOcurrencies(taskList, description), countDescriptionOcurrencies(taskList, anotherDescription))
	}
}

func TestTaskListUpdateMethod(t *testing.T) {
	taskList := NewTaskList()

	taskList.Add(description)
	taskList.Update(description, anotherDescription)

	isDescriptionUnchanged := countDescriptionOcurrencies(taskList, anotherDescription) == 0

	if isDescriptionUnchanged {
		t.Errorf("The Update method is not working properly \n Count: %d \n", countDescriptionOcurrencies(taskList, anotherDescription))
	}
}
