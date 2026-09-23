package persistence

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

type MockedFs struct {
	errorMock       error
	osFileMock      *os.File
	dirMock         string
	readedBytesMock []byte

	name         string
	path         string
	completePath string
	data         string
}

func (f *MockedFs) Write(path string, data string) error {
	f.completePath = path
	f.data = data
	return f.errorMock
}

func (f *MockedFs) RemoveAll(path string) error {
	f.path = path

	return f.errorMock
}

func (f *MockedFs) Getwd() (dir string, err error) { return f.dirMock, f.errorMock }

func (f *MockedFs) ReadFile(completePath string) ([]byte, error) {
	f.path = completePath
	return f.readedBytesMock, f.errorMock
}

func TestExtraInfo(t *testing.T) {
	file, path := "filename", "path"
	pers := Persistence{file, path, &MockedFs{}}

	fn, p := pers.ExtraInfo()
	if fn != file {
		t.Errorf("Filename not match. Expected: %s, received %s", file, fn)
	}

	if p != path {
		t.Errorf("Path not match. Expected: %s, received %s", path, p)
	}
}

func TestRead(t *testing.T) {
	fileContent := "mycontent"
	dir, file, path := "dir", "file", "path"
	fs := MockedFs{readedBytesMock: []byte(fileContent), dirMock: dir}
	pers := Persistence{file, path, &fs}

	readed := string(pers.Read())
	if readed != fileContent {
		t.Errorf("Read failed.\n Expected: %s, received: %s", fileContent, readed)
	}

	expectedPath := filepath.Join(dir, path, file)

	if fs.path != expectedPath {
		t.Errorf("Path not mounted correctly when reading.\n Expected: %s, received: %s", expectedPath, fs.path)
	}
}

func TestRemoveAll(t *testing.T) {
	fileContent := "mycontent"
	dir, file, path := "dir", "file", "path"
	fs := MockedFs{readedBytesMock: []byte(fileContent), dirMock: dir}
	pers := Persistence{file, path, &fs}

	pers.Remove()
	expectedPath := filepath.Join(dir, path, file)

	if fs.path != expectedPath {
		t.Errorf("Path not mounted correctly when removing.\n Expected: %s, received: %s", expectedPath, fs.path)
	}
}

func TestWrite(t *testing.T) {
	fileContent := "mycontent"
	dir, file, path := "dir", "file", "path"
	fs := MockedFs{readedBytesMock: []byte(fileContent), dirMock: dir}

	pers := Persistence{file, path, &fs}

	expectedPath := filepath.Join(dir, path, file)

	valid := "{}"
	pers.Write(valid)

	if fs.completePath != expectedPath {
		t.Errorf("Path not mounted correctly when writing.\n Expected: %s, received: %s", expectedPath, fs.path)
	}

	validInBytes, err := json.Marshal(valid)
	if err != nil {
		t.Error("Error on test string parsing")
	}

	if fs.data != string(validInBytes) {
		t.Errorf("Error on parsing data.\n Expected: %s, received: %s", string(validInBytes), fs.data)
	}
	invalid := "{\"invalid\""
	pers.Write(invalid)

	invalidInBytes, err := json.Marshal(invalid)
	if err != nil {
		t.Error("Error on test string parsing")
	}

	if fs.data != string(invalidInBytes) {
		t.Errorf("Error on parsing data.\n Expected: %s, received: %s", string(invalidInBytes), fs.data)
	}
}
