package subcommands

import "cli-stack-tracker/internal/globaltypes"

type ClearSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (cs *ClearSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	cs.p = p
}

func (cs *ClearSubcommand) persistence() *globaltypes.IPersistence {
	return cs.p
}

func (cs *ClearSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	cs.tl = tl
}

func (cs *ClearSubcommand) tasks() *globaltypes.ITaskList {
	return cs.tl
}

func (cs *ClearSubcommand) Command() string {
	return "clear"
}

func (cs *ClearSubcommand) Exec() {
	(*cs.p).Remove()
}
