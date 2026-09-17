package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type ListSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (ls *ListSubcommand) SetPersistence(p globaltypes.IPersistence) {
	ls.p = p
}

func (ls *ListSubcommand) persistence() globaltypes.IPersistence {
	return ls.p
}

func (ls *ListSubcommand) SetTasks(tl globaltypes.ITaskList) {
	ls.tl = tl
}

func (ls *ListSubcommand) tasks() globaltypes.ITaskList {
	return ls.tl
}

func (ls *ListSubcommand) Command() string {
	return "list"
}

func (ls *ListSubcommand) Exec() {
}
