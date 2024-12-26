package core

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"sigs.k8s.io/yaml"
)

func YAMLToJSONFromFile(fileName string) (result []byte, err error) {
	if err := ValidateFilePath(fileName); err != nil {
		return nil, err
	}

	yamlBytes, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}

	return yaml.YAMLToJSON(yamlBytes)
}

func JSONToYAMLFromFile(fileName string) (result []byte, err error) {
	if err := ValidateFilePath(fileName); err != nil {
		return nil, err
	}

	jsonBytes, err := os.ReadFile(filepath.Clean(fileName))
	if err != nil {
		return nil, err
	}

	return yaml.JSONToYAML(jsonBytes)
}

func ReadJSONFile(fileName string) (result []byte, err error) {
	if err := ValidateFilePath(fileName); err != nil {
		return nil, err
	}

	return os.ReadFile(filepath.Clean(fileName))
}

func ValidateFilePath(fileName string) error {
	// normalize path
	cleanPath := filepath.Clean(fileName)

	// check for directory traversal
	if strings.Contains(cleanPath, "..") {
		return fmt.Errorf("directory traversal ist not allowed")
	}

	// check for absolute paths
	// if filepath.IsAbs(cleanPath) {
	// 	return fmt.Errorf("absolute paths are not allowed")
	// }

	return nil
}
