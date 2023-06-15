package main

import (
	"io/ioutil"
	"strings"
)

// Checks for the presence of certain strings in a file
func ScanFiles(filePaths []string, patterns []string) ([]string, error) {
	var suspiciousFiles []string

	for _, path := range filePaths {
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return nil, err
		}
		fileData := string(data)

		for _, pattern := range patterns {
			if strings.Contains(fileData, pattern) {
				suspiciousFiles = append(suspiciousFiles, path)
				break
			}
		}
	}

	return suspiciousFiles, nil
}
