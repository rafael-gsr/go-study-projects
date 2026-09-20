package subcommands

import (
	"flag"
	"fmt"
	"os"

	"cli-stack-tracker/internal/globaltypes"
)

type UpdateSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (us *UpdateSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	us.p = p
}

func (us *UpdateSubcommand) persistence() *globaltypes.IPersistence {
	return us.p
}

func (us *UpdateSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	us.tl = tl
}

func (us *UpdateSubcommand) tasks() *globaltypes.ITaskList {
	return us.tl
}

func (us *UpdateSubcommand) Command() string {
	return "update"
}

func (us *UpdateSubcommand) Exec() {
	UpdateCmd := flag.NewFlagSet("update", flag.ExitOnError)
	ID := UpdateCmd.String("i", "", "-i [string]")
	description := UpdateCmd.String("d", "", "-d [string]")

	UpdateCmd.Parse(os.Args[2:])

	hasMissingParams := *ID == "" && *description == ""
	if hasMissingParams {
		fmt.Println("Error: Missing params.\nInform ID and description using -i and -d flags")
		return
	}

	(*us.tl).Update(*ID, *description)
	(*us.p).Write((*us.tl).GetTasks())
}
