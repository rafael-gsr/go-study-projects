// Package subcommands contains the application subcommands and its actions
package subcommands

import (
	globaltypes "cli-stack-tracker/globalTypes"
	"fmt"
	"os"
)

type SubcommandsParser struct {
	persistence globaltypes.IPersistence
	tasks       globaltypes.ITaskList
}

func (s *SubcommandsParser) Parse() {
	subcommands := []globaltypes.ISubcommands{}
	userCommand := os.Args[1]

	if userCommand == "" {
		fmt.Println("Missing user commands, user -h to see how to use the tool")
		os.Exit(1)
	}

	for _, command := range subcommands {
		if userCommand == command.Command() {
			command.Exec()
			return
		}
	}
}

func NewSubcommandsParser(
	persistence globaltypes.IPersistence,
	tasks globaltypes.ITaskList,
) *SubcommandsParser {
	return &SubcommandsParser{persistence, tasks}
}
