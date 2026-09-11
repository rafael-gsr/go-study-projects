package flags

import (
	"flag"
	"fmt"
	"strings"

	globaltypes "cli-stack-tracker/globalTypes"
)

type Flags struct {
	persistence globaltypes.IPersistence
	tasks       globaltypes.ITaskList
}

func NewFlags(persistence globaltypes.IPersistence, tasks globaltypes.ITaskList) *Flags {
	return &Flags{persistence, tasks}
}

func (f *Flags) Setup() {
	addFlag := flag.String("add", "", "--add \"[description]\" - creates an input")
	updateFlag := flag.String("update", "", "--update \"[id]:[newDescription]\"  - updates an input")
	cleanFlag := flag.String("clean", "", "--clean - cleans tasks")
	removeFlag := flag.String("remove", "", "--remove [id] - remove task")
	markAsDoneFlag := flag.String("done", "", "--done [id] - mark task as done")
	markAsPendingFlag := flag.String("pending", "", "--pending [id] - mark task as pending")
	listFlag := flag.Bool("list", false, "--list - list tasks")

	flag.Parse()

	if *addFlag != "" {
		f.add(*addFlag)
	}

	if *updateFlag != "" {
		idAndDesc := strings.Split(*updateFlag, ":")
		if len(idAndDesc) < 2 {
			fmt.Println("Error: missing id or description")
			return
		}

		f.update(idAndDesc[0], idAndDesc[1])
	}

	if *cleanFlag != "" {
		f.clean()
	}

	if *removeFlag != "" {
		f.remove(*removeFlag)
	}

	if *listFlag {
		f.list()
	}

	if *markAsDoneFlag != "" {
		f.done(*markAsDoneFlag)
	}

	if *markAsPendingFlag != "" {
		f.pending(*markAsPendingFlag)
	}
}

func (f *Flags) add(description string) {
	f.tasks.Add(description)
	f.persistence.Write(f.tasks.GetTasks())
}

func (f *Flags) update(id string, description string) {
	f.tasks.Update(id, description)
	f.persistence.Write(f.tasks.GetTasks())
}

func (f *Flags) remove(id string) {
	f.tasks.Remove(id)
	f.persistence.Write(f.tasks.GetTasks())
}

func (f *Flags) clean() {
	f.persistence.Remove()
}

func (f *Flags) list() {
	for _, task := range f.tasks.GetTasks() {
		fmt.Printf("ID: %s | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}
}

func (f *Flags) done(id string) {
	f.tasks.MarkAsDone(id)
	f.persistence.Write(f.tasks.GetTasks())
}

func (f *Flags) pending(id string) {
	f.tasks.MarkAsPenging(id)
	f.persistence.Write(f.tasks.GetTasks())
}
