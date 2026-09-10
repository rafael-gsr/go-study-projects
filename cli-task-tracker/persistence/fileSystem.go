package persitence

import "os"

type FileSystem struct{}

func (f FileSystem) Create(name string) (*os.File, error) { return os.Create(name) }

func (f FileSystem) RemoveAll(path string) error { return os.RemoveAll(path) }

func (f FileSystem) Getwd() (dir string, err error) { return os.Getwd() }

func (f FileSystem) ReadFile(name string) ([]byte, error) { return os.ReadFile(name) }
