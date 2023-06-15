package main

import (
	"os"
	"path/filepath"
)

const (
	MaxFileSize      = 10 * 1024 * 1024 // 10MB
	AllowedExtension = ".exe"
)

func FilterFiles(filePaths []string) ([]string, error) {
	var filteredPaths []string

	for _, path := range filePaths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		// Check file size
		if fileInfo.Size() > MaxFileSize {
			continue
		}

		// Check file extension
		if filepath.Ext(path) != AllowedExtension {
			continue
		}

		filteredPaths = append(filteredPaths, path)
	}

	return filteredPaths, nil
}
