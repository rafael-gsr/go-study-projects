package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type DoneSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (ds *DoneSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	ds.p = p
}

func (ds *DoneSubcommand) persistence() *globaltypes.IPersistence {
	return ds.p
}

func (ds *DoneSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	ds.tl = tl
}

func (ds *DoneSubcommand) tasks() *globaltypes.ITaskList {
	return ds.tl
}

func (ds *DoneSubcommand) Command() string {
	return "done"
}

func (ds *DoneSubcommand) Exec() {
	DoneCmd := flag.NewFlagSet("done", flag.ExitOnError)
	DoneCmd.Parse(os.Args[2:])

	taskID := DoneCmd.Arg(0)
	if taskID == "" {
		fmt.Println("Inform the taskID to be updated.")
		return
	}

	(*ds.tl).MarkAsDone(taskID)
	(*ds.p).Write((*ds.tl).GetTasks())
}
