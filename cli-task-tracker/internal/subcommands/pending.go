package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type PendingSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (ps *PendingSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	ps.p = p
}

func (ps *PendingSubcommand) persistence() *globaltypes.IPersistence {
	return ps.p
}

func (ps *PendingSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	ps.tl = tl
}

func (ps *PendingSubcommand) tasks() *globaltypes.ITaskList {
	return ps.tl
}

func (ps *PendingSubcommand) Command() string {
	return "pending"
}

func (ps *PendingSubcommand) Exec() {
	PendingCmd := flag.NewFlagSet("pending", flag.ExitOnError)
	PendingCmd.Parse(os.Args[2:])

	taskID := PendingCmd.Arg(0)
	if taskID == "" {
		fmt.Println("Inform the taskID to be updated.\nUsage: pending [string]")
		return
	}

	(*ps.tl).MarkAsPending(taskID)
	(*ps.p).Write((*ps.tl).GetTasks())
}
