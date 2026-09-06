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
	addFlag := flag.String("add", "", "--add \"[description]\" - creates an input")
	updateFlag := flag.String("update", "", "--add \"[old]:[new]\"  - updates an input")

	if *addFlag != "" {
		// TODO
		fmt.Println(&addFlag)
	}

	if *updateFlag != "" {
		// TODO
		fmt.Println(&updateFlag)
	}
}

func (f *Flags) add()    {}
func (f *Flags) update() {}
func (f *Flags) remove() {}
func (f *Flags) clean()  {}
