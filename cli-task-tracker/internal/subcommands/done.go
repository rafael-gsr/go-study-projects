package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type DoneSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (ds *DoneSubcommand) SetPersistence(p globaltypes.IPersistence) {
	ds.p = p
}

func (ds *DoneSubcommand) persistence() globaltypes.IPersistence {
	return ds.p
}

func (ds *DoneSubcommand) SetTasks(tl globaltypes.ITaskList) {
	ds.tl = tl
}

func (ds *DoneSubcommand) tasks() globaltypes.ITaskList {
	return ds.tl
}

func (ds *DoneSubcommand) Command() string {
	return "done"
}

func (ds *DoneSubcommand) Exec() {
}
