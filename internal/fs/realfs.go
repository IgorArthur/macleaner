package fs

import (
	"os"
	"path/filepath"
	"strings"
)

type RealFS struct{}

func (RealFS) ExpandPath(path string) (string, error) {
	if strings.HasPrefix(path, "~") {
		home, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		path = filepath.Join(home, path[1:])
	}
	return os.ExpandEnv(path), nil
}

func (RealFS) DirSize(path string) (int64, error) {
	var size int64

	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			size += info.Size()
		}
		return nil
	})

	return size, err
}

func (RealFS) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (RealFS) Exists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
