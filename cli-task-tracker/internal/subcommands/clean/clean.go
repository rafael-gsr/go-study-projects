package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type CleanSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (cs *CleanSubcommand) SetPersistence(p globaltypes.IPersistence) {
	cs.p = p
}

func (cs *CleanSubcommand) persistence() globaltypes.IPersistence {
	return cs.p
}

func (cs *CleanSubcommand) SetTasks(tl globaltypes.ITaskList) {
	cs.tl = tl
}

func (cs *CleanSubcommand) tasks() globaltypes.ITaskList {
	return cs.tl
}

func (cs *CleanSubcommand) Command() string {
	return "clean"
}

func (cs *CleanSubcommand) Exec() {
}
