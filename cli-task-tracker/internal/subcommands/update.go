package subcommands

import globaltypes "cli-stack-tracker/globalTypes"

type UpdateSubcommand struct {
	p  globaltypes.IPersistence
	tl globaltypes.ITaskList
}

func (us *UpdateSubcommand) setpersistence(p globaltypes.IPersistence) {
	us.p = p
}

func (us *UpdateSubcommand) persistence() globaltypes.IPersistence {
	return us.p
}

func (us *UpdateSubcommand) settasks(tl globaltypes.ITaskList) {
	us.tl = tl
}

func (us *UpdateSubcommand) tasks() globaltypes.ITaskList {
	return us.tl
}

func (us *UpdateSubcommand) command() string {
	return "list"
}

func (us *UpdateSubcommand) exec() {
}
