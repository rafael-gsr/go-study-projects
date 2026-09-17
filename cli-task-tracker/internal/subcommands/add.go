package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type AddSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (as *AddSubcommand) SetPersistence(p globaltypes.IPersistence) {
	as.p = p
}

func (as *AddSubcommand) persistence() globaltypes.IPersistence {
	return as.p
}

func (as *AddSubcommand) SetTasks(tl globaltypes.ITaskList) {
	as.tl = tl
}

func (as *AddSubcommand) tasks() globaltypes.ITaskList {
	return as.tl
}

func (as *AddSubcommand) Command() string {
	return "add"
}

func (as *AddSubcommand) Exec() {
}
