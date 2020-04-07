package utils

import (
	"os"
	"path/filepath"
)

var files []string

func FilesInDir(directory string) ([]string, error) {
	if err := filepath.Walk(directory,
		func(path string, info os.FileInfo, err error) error {
			files = append(files, path)
			return nil
		},
	); err != nil {
		return nil, err
	}

	return files, nil
}
