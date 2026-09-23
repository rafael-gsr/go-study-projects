// Package subcommands contains the application subcommands and its actions
package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type SubcommandsParser struct {
	persistence globaltypes.IPersistence
	tasks       globaltypes.ITaskList
	subMap      map[string]globaltypes.ISubcommands
}

func (s *SubcommandsParser) Parse() {
	helpFlag := flag.Bool("h", false, "-h -- shows help")
	if *helpFlag {
		s.printHelper()
		return
	}

	if len(os.Args) < 2 {
		fmt.Println("Missing user commands, user -h to see how to use the tool")
		os.Exit(1)
	}

	userCommand := os.Args[1]
	s.Execute(userCommand)
}

func (s *SubcommandsParser) Execute(command string) {
	if sub, valid := s.subMap[command]; valid {
		sub.Exec()
	} else {
		fmt.Println("Invalid command")
		s.printHelper()
	}
}

func (s SubcommandsParser) printHelper() {
	fmt.Println("Usage:")
	fmt.Println("  list -- list the stored tasks")
	fmt.Println("\n  add [description] -- adds a new task ")
	fmt.Println("    description: string")
	fmt.Println("\n  clear -- removes all tasks")
	fmt.Println("\n  done [taskID] -- marks task as done")
	fmt.Println("    taskID: string")
	fmt.Println("\n  pending [taskID] -- marks task as pending")
	fmt.Println("    taskID: string")
	fmt.Println("\n  remove [taskID] -- removes a task from list")
	fmt.Println("    taskID: string")
	fmt.Println("\n  update -i [taskID] -d [description] -- updates task description")
	fmt.Println("    taskID: string")
	fmt.Println("    description: string")
}

func NewSubcommandsParser(
	persistence globaltypes.IPersistence,
	tasks globaltypes.ITaskList,
) *SubcommandsParser {
	addCmd := &AddSubcommand{&persistence, &tasks}
	clearCmd := &ClearSubcommand{&persistence, &tasks}
	doneCmd := &DoneSubcommand{&persistence, &tasks}
	listCmd := &ListSubcommand{&persistence, &tasks}

	pendCmd := &PendingSubcommand{&persistence, &tasks}
	remCmd := &RemoveSubcommand{&persistence, &tasks}
	updCmd := &UpdateSubcommand{&persistence, &tasks}

	subMap := map[string]globaltypes.ISubcommands{
		addCmd.Command():   addCmd,
		clearCmd.Command(): clearCmd,
		doneCmd.Command():  doneCmd,
		listCmd.Command():  listCmd,
		pendCmd.Command():  pendCmd,
		remCmd.Command():   remCmd,
		updCmd.Command():   updCmd,
	}

	return &SubcommandsParser{persistence, tasks, subMap}
}
