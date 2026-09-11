package persitence

import (
	"encoding/json"
	"fmt"
	"path/filepath"
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
	fileSystem FileSystem
}

func (p *Persistence) getCompletePath() string {
	pwd, err := p.fileSystem.Getwd()
	errCheck(err)

	return filepath.Join(pwd, p.path, p.filename)
}

func (p *Persistence) Write(data any) {
	f, err := p.fileSystem.Create(p.getCompletePath())
	errCheck(err)
	defer f.Close()

	jsonData, err := json.Marshal(data)
	errCheck(err)

	f.WriteString(string(jsonData))
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

func NewPersistence(filename string, path string) *Persistence {
	return &Persistence{filename, path, FileSystem{}}
}
