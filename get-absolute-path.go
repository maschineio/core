package core

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetAbsDirectory(directoryName string) (s string, err error) {
	if filepath.IsAbs(directoryName) {
		return directoryName, nil
	}
	workDir, err := os.Getwd()
	if err != nil {
		return s, err
	}

	return filepath.Clean(filepath.Join(workDir, directoryName)), nil
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
