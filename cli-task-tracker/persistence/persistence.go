package persitence

import (
	"encoding/json"
	"os"
	"path/filepath"
)

func errCheck(err error) {
	if err != nil {
		panic(err)
	}
}

type FileSystem struct{}

func (f FileSystem) Create(name string) (*os.File, error) { return os.Create(name) }

func (f FileSystem) RemoveAll(path string) error { return os.RemoveAll(path) }

func (f FileSystem) Getwd() (dir string, err error) { return os.Getwd() }

func (f FileSystem) ReadFile(name string) ([]byte, error) { return os.ReadFile(name) }

type Persistence struct {
	filename   string
	path       string
	FileSystem FileSystem
}

func NewPersistence(filename string, path string) *Persistence {
	return &Persistence{filename, path, FileSystem{}}
}

func (p *Persistence) getCompletePath() string {
	pwd, err := p.FileSystem.Getwd()
	errCheck(err)
	return filepath.Join(pwd, p.path, p.filename)
}

func (p *Persistence) Write(data any) {
	f, err := p.FileSystem.Create(p.getCompletePath())
	errCheck(err)
	defer f.Close()

	jsonData, err := json.Marshal(data)
	errCheck(err)

	f.WriteString(string(jsonData))
}

func (p *Persistence) Remove() {
	err := p.FileSystem.RemoveAll(p.getCompletePath())
	errCheck(err)
}

func (p *Persistence) Read() []byte {
	file, err := p.FileSystem.ReadFile(p.getCompletePath())
	errCheck(err)

	return file
}
