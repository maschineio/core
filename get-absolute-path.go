package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsDirectory(directoryName string) (s string, err error) {
	// Prüfe auf leeren Pfad
	if directoryName == "" {
		workDir, err := os.Getwd()
		if err != nil {
			return s, err
		}
		return workDir, nil
	}

	// Wenn absoluter Pfad, prüfe ob zugreifbar
	if filepath.IsAbs(directoryName) {
		if _, err := os.Stat(directoryName); err != nil {
			return s, err
		}
		return directoryName, nil
	}

	// Hole aktuelles Arbeitsverzeichnis
	workDir, err := os.Getwd()
	if err != nil {
		return s, err
	}

	// Erstelle absoluten Pfad und prüfe Zugriff
	absPath := filepath.Clean(filepath.Join(workDir, directoryName))
	if _, err := os.Stat(absPath); err != nil {
		return s, err
	}

	return absPath, nil
}

func GetAbsFilePath(fileName string, check bool) (s string, err error) {
	var fullPath string

	if filepath.IsAbs(fileName) {
		fullPath = fileName
	} else {
		workDir, err := os.Getwd()
		if err != nil {
			return s, err
		}
		fullPath = filepath.Clean(filepath.Join(workDir, fileName))
	}

	if check {
		fInfo, err := os.Stat(fullPath)

		if err != nil {
			return s, err
		}

		if fInfo.IsDir() {
			return s, fmt.Errorf("unexpected directory; must be a regular file")
		}
	}

	return fullPath, nil
}
