package main

import (
	"fmt"
	"log"

	"./files"
	"./prerequisites"
	"./scanner"
	"./signed"
)

func main() {
	settings, err := files.GetSettings()
	if err != nil {
		log.Fatalf("Error getting settings: %v", err)
	}

	allFiles, err := files.GetFiles(settings.FilePath)
	if err != nil {
		log.Fatalf("Error getting files: %v", err)
	}

	filteredFiles, err := prerequisites.FilterFiles(allFiles, settings)
	if err != nil {
		log.Fatalf("Error filtering files: %v", err)
	}

	unsignedFiles, err := signed.FilterSignedFiles(filteredFiles)
	if err != nil {
		log.Fatalf("Error filtering signed files: %v", err)
	}

	suspiciousFiles, err := scanner.ScanFiles(unsignedFiles, []string{"processMemoryFuncs", "otherFuncs"})
	if err != nil {
		log.Fatalf("Error scanning files: %v", err)
	}

	for _, file := range suspiciousFiles {
		fmt.Println("Suspicious file:", file)
	}
}
