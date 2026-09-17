package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type PendingSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (ps *PendingSubcommand) setpersistence(p globaltypes.IPersistence) {
	ps.p = p
}

func (ps *PendingSubcommand) persistence() globaltypes.IPersistence {
	return ps.p
}

func (ps *PendingSubcommand) settasks(tl globaltypes.ITaskList) {
	ps.tl = tl
}

func (ps *PendingSubcommand) tasks() globaltypes.ITaskList {
	return ps.tl
}

func (ps *PendingSubcommand) command() string {
	return "list"
}

func (ps *PendingSubcommand) exec() {
}
