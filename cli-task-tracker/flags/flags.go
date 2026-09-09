package flags

import (
	"flag"
	"fmt"
	"strings"
)

type Flags struct {
	persistence *Persistence
	tasks       *TaskList
}

func NewFlags(persistence *Persistence, tasks *TaskList) *Flags {
	return &Flags{persistence, tasks}
}

func (f *Flags) Setup() {
	addFlag := flag.String("add", "", "--add \"[description]\" - creates an input")
	updateFlag := flag.String("update", "", "--update \"[id]:[newDescription]\"  - updates an input")
	cleanFlag := flag.String("clean", "", "--clean - cleans tasks")
	removeFlag := flag.String("remove", "", "--remove [id] - remove task")
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
}

func (f *Flags) add(description string) {
	f.tasks.Add(description)
	f.persistence.Write(*f.tasks)
}

func (f *Flags) update(id string, description string) {
	f.tasks.Update(id, description)
	f.persistence.Write(*f.tasks)
}

func (f *Flags) remove(id string) {
	f.tasks.Remove(id)
	f.persistence.Write(*f.tasks)
}

func (f *Flags) clean() {
	f.persistence.Remove()
}

func (f *Flags) list() {
	for _, task := range f.tasks.tasks {
		fmt.Printf("ID: %s | Description: %s | Status: %s", task.ID, task.Description, task.Status)
	}
}
