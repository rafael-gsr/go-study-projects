package main

import (
	"flag"
	"fmt"
)

type Flags struct {
	persistence *Persistence
	tasks       *TaskList
}

func NewFlags(persistence *Persistence, tasks *TaskList) *Flags {
	return &Flags{persistence, tasks}
}

func (f *Flags) Setup() {
	addFlag := flag.String("add", nil, "--add \"[description]\" - creates an input")
	updateFlag := flag.String("update", nil, "--add \"[old]:[new]\"  - updates an input")

	if &addFlag != nil {
		// TODO
		fmt.Println(&addFlag)
	}

	if &updateFlag != nil {
		// TODO
		fmt.Println(&updateFlag)
	}
}

func (f *Flags) add()    {}
func (f *Flags) update() {}
func (f *Flags) remove() {}
func (f *Flags) clean()  {}
