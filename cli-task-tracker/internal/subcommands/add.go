package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type AddSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (as *AddSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	as.p = p
}

func (as *AddSubcommand) persistence() *globaltypes.IPersistence {
	return as.p
}

func (as *AddSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	as.tl = tl
}

func (as *AddSubcommand) tasks() *globaltypes.ITaskList {
	return as.tl
}

func (as *AddSubcommand) Command() string {
	return "add"
}

func (as *AddSubcommand) Exec() {
	AddCmd := flag.NewFlagSet("add", flag.ExitOnError)
	AddCmd.Parse(os.Args[2:])

	description := AddCmd.Arg(0)
	if description == "" {
		AddCmd.Usage()
		fmt.Println("Error: Inform the description.\nUsage: add [string]")
		return
	}

	(*as.tl).Add(description)
	(*as.p).Write((*as.tl).GetTasks())
}
