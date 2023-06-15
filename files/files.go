package main

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Settings struct {
	FilePath string `json:"filepath"`
}

// Reads the settings.json and returns the struct
func GetSettings() (Settings, error) {
	var settings Settings

	// Read the settings file
	file, err := os.Open("settings.json")
	if err != nil {
		return settings, err
	}
	defer file.Close()

	// Decode the JSON
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&settings); err != nil {
		return settings, err
	}

	return settings, nil
}

// Retrieves all the files from a directory and returns their paths
func GetFiles(dir string) ([]string, error) {
	var paths []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			paths = append(paths, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	return paths, nil
}
