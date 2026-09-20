package persistence

import (
	"errors"
	"os"
)

type FileSystem struct{}

func (f FileSystem) Write(path string, data string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(data)
	return nil
}

func (f FileSystem) RemoveAll(path string) error { return os.RemoveAll(path) }

func (f FileSystem) Getwd() (dir string, err error) { return os.Getwd() }

func (f FileSystem) ReadFile(completePath string) ([]byte, error) {
	if _, err := os.Stat(completePath); err != nil && errors.Is(err, os.ErrNotExist) {
		os.Create(completePath)
	}

	return os.ReadFile(completePath)
}
