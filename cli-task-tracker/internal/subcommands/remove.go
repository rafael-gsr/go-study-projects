package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type RemoveSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (rs *RemoveSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	rs.p = p
}

func (rs *RemoveSubcommand) persistence() *globaltypes.IPersistence {
	return rs.p
}

func (rs *RemoveSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	rs.tl = tl
}

func (rs *RemoveSubcommand) tasks() *globaltypes.ITaskList {
	return rs.tl
}

func (rs *RemoveSubcommand) Command() string {
	return "remove"
}

func (rs *RemoveSubcommand) Exec() {
	RemoveCmd := flag.NewFlagSet("remove", flag.ExitOnError)
	RemoveCmd.Parse(os.Args[2:])

	taskId := RemoveCmd.Arg(0)

	if taskId == "" {
		fmt.Println("Error: Inform the task ID.\nUsage: remove [string]")
		return
	}

	(*rs.tl).Remove(taskId)
	(*rs.p).Write((*rs.tl).GetTasks())
}
