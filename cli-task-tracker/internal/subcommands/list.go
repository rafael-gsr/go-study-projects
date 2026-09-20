package subcommands

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"cli-stack-tracker/internal/globaltypes"
)

type ListSubcommand struct {
	p  *globaltypes.IPersistence
	tl *globaltypes.ITaskList
}

func (ls *ListSubcommand) SetPersistence(p *globaltypes.IPersistence) {
	ls.p = p
}

func (ls *ListSubcommand) persistence() *globaltypes.IPersistence {
	return ls.p
}

func (ls *ListSubcommand) SetTasks(tl *globaltypes.ITaskList) {
	ls.tl = tl
}

func (ls *ListSubcommand) tasks() *globaltypes.ITaskList {
	return ls.tl
}

func (ls *ListSubcommand) Command() string {
	return "list"
}

func (ls *ListSubcommand) Exec() {
	AddCmd := flag.NewFlagSet("list", flag.ExitOnError)
	ExtraInfo := AddCmd.Bool("i", false, "list -i -- shows more info")

	AddCmd.Parse(os.Args[2:])

	tasks := (*ls.tl).GetTasks()
	for _, task := range tasks {
		fmt.Printf("ID: %s | Description: %s | Status: %s\n", task.ID, task.Description, task.Status)
	}

	if *ExtraInfo {
		filename, path := (*ls.p).ExtraInfo()
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error: cannot get working directory")
			return
		}

		path = filepath.Join(wd, path, filename)

		fmt.Printf("\nFile: %s \nPath: %s", filename, path)
	}
}
