package main

import (
	"os"
	"path/filepath"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

type Persistence struct {
	filename string
	path     string
}

func (p *Persistence) getCompletePath() string {
	pwd, err := os.Getwd()
	errCheck(err)
	return filepath.Join(pwd, p.path, p.filename)
}

func (p *Persistence) Write(data string) {
	f, err := os.Create(p.getCompletePath())
	errCheck(err)
	defer f.Close()

	f.WriteString(data)
}

func (p *Persistence) remove() {
	err := os.RemoveAll(p.getCompletePath())
	errCheck(err)
}
