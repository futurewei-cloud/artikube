package storage

import (
	"io/ioutil"
	"os"

	pathutil "path"
)

// FilesystemBackend is a storage backend for local filesystem storage
type FilesystemBackend struct {
	RootDirectory string
}

// NewLocalFilesystemBackend creates a new instance of LocalFilesystemBackend
func NewFilesystemBackend(rootDirectory string) *FilesystemBackend {
	if _, err := os.Stat(rootDirectory); os.IsNotExist(err) {
		err := os.MkdirAll(rootDirectory, 0777)
		if err != nil {
			panic(err)
		}
	}
	b := &FilesystemBackend{RootDirectory: rootDirectory}
	return b
}

// Lists all objects in root directory only, without recursive
func (fs FilesystemBackend) ListObjects() ([]Object, error) {
	var objects []Object
	files, err := ioutil.ReadDir(fs.RootDirectory)
	if err != nil {
		return objects, err
	}
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		object := Object{Path: f.Name(), Artifact: []byte{}, LastModified: f.ModTime()}
		objects = append(objects, object)
	}
	return objects, nil
}

func (fs FilesystemBackend) GetObject(path string) (Object, error) {
	var object Object
	object.Path = path
	fullpath := pathutil.Join(fs.RootDirectory, path)
	artifact, err := ioutil.ReadFile(fullpath)
	if err != nil {
		return object, err
	}
	object.Artifact = artifact
	info, err := os.Stat(fullpath)
	if err != nil {
		return object, err
	}
	object.LastModified = info.ModTime()
	return object, err
}

func (fs FilesystemBackend) PutObject(path string, content []byte) error {
	fullpath := pathutil.Join(fs.RootDirectory, path)
	err := ioutil.WriteFile(fullpath, content, 0644)
	return err
}

// Delet removes an object from root directory
func (fs FilesystemBackend) DeleteObject(path string) error {
	fullpath := pathutil.Join(fs.RootDirectory, path)
	err := os.Remove(fullpath)
	return err
}
