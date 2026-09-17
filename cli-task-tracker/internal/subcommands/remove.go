package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type RemoveSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (rs *RemoveSubcommand) setpersistence(p globaltypes.IPersistence) {
	rs.p = p
}

func (rs *RemoveSubcommand) persistence() globaltypes.IPersistence {
	return rs.p
}

func (rs *RemoveSubcommand) settasks(tl globaltypes.ITaskList) {
	rs.tl = tl
}

func (rs *RemoveSubcommand) tasks() globaltypes.ITaskList {
	return rs.tl
}

func (rs *RemoveSubcommand) command() string {
	return "list"
}

func (rs *RemoveSubcommand) exec() {
}
