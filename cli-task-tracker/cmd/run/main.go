package main

import (
	"cli-stack-tracker/flags"
	persitence "cli-stack-tracker/persistence"
	"cli-stack-tracker/task"
)

func main() {
	persistence := persitence.NewPersistence("tasks.txt", "")
	storedContent := persistence.Read()
	tasks := task.NewTaskList(storedContent)
	flags := flags.NewFlags(persistence, tasks)
	flags.Setup()
}
