package main

import (
	"os/exec"
	"strings"
)

// Checks if the file is signed using sigcheck
func FilterSignedFiles(filePaths []string) ([]string, error) {
	var unsignedFiles []string

	for _, path := range filePaths {
		out, err := exec.Command("sigcheck.exe", "-nobanner", "-accepteula", path).Output()
		if err != nil {
			return nil, err
		}
		output := string(out)

		// If the output doesn't contain "Verified: Signed", the file is not signed
		if !strings.Contains(output, "Verified:       Signed") {
			unsignedFiles = append(unsignedFiles, path)
		}
	}

	return unsignedFiles, nil
}
