package persistence

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"cli-stack-tracker/internal/globaltypes"
)

func errCheck(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

type Persistence struct {
	filename   string
	path       string
	fileSystem globaltypes.IFileSystem
}

func (p *Persistence) getCompletePath() string {
	pwd, err := p.fileSystem.Getwd()
	errCheck(err)

	return filepath.Join(pwd, p.path, p.filename)
}

func (p *Persistence) Write(data any) {
	jsonData, err := json.Marshal(data)
	errCheck(err)
	p.fileSystem.Write(p.getCompletePath(), string(jsonData))
}

func (p *Persistence) Remove() {
	err := p.fileSystem.RemoveAll(p.getCompletePath())
	errCheck(err)
}

func (p *Persistence) Read() []byte {
	file, err := p.fileSystem.ReadFile(p.getCompletePath())
	errCheck(err)

	return file
}

func (p *Persistence) ExtraInfo() (filename string, path string) {
	return p.filename, p.path
}

func NewPersistence(filename string, path string) *Persistence {
	return &Persistence{filename, path, FileSystem{}}
}
