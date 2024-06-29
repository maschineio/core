package core

import (
	"path/filepath"
	"strings"
)

func GetFileBasenameAndSuffix(fileName string) (string, string) {
	extension := filepath.Ext(fileName)
	return strings.TrimSuffix(fileName, extension), extension
}
