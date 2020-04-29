package storage

import (
	"fmt"
	"strings"
	"time"
)

type (
	Object struct {
		Path         string
		Artifact     []byte
		LastModified time.Time
	}

	//Provides what has changed since the last time when ListObject was called
	ObjectDiff struct {
		Change  bool
		Removed []Object
		Added   []Object
		Updated []Object
	}
	//generic interface for storage backends
	Backend interface {
		ListObjects() ([]Object, error)
		GetObject(path string) (Object, error)
		PutObject(path string, connect []byte) error
		DeleteObject(path string) error
	}
)

func cleanPrefix(prefix string) string {
	return strings.Trim(prefix, "/")
}

func removePrefixFromObjectPath(prefix string, path string) string {
	if prefix == "" {
		return path
	}
	path = strings.Replace(path, fmt.Sprintf("%s/", prefix), "", 1)
	return path
}

func objectPathIsInvalid(path string) bool {
	return strings.Contains(path, "/") || path == ""
}
