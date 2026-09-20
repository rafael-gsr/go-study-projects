package main

import (
	"cli-stack-tracker/internal/persistence"
	"cli-stack-tracker/internal/subcommands"
	"cli-stack-tracker/internal/task"
)

func main() {
	pers := persistence.NewPersistence("tasks.json", "")
	storedContent := pers.Read()

	tasks, err := task.NewTaskList(storedContent)
	if err != nil {
		pers.Write([]string{})
	}

	subcmds := subcommands.NewSubcommandsParser(pers, tasks)
	subcmds.Parse()
}
