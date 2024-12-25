package core_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestYAMLToJSONFromFile(t *testing.T) {
	// Create a temporary YAML file
	yamlFileName := "test.yaml"
	yamlContent := []byte("key: value\nnested:\n  key: nestedValue")
	err := os.WriteFile(yamlFileName, yamlContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(yamlFileName)

	// Test the function
	result, err := core.YAMLToJSONFromFile(yamlFileName)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"key":"value","nested":{"key":"nestedValue"}}`, string(result))
}

func TestYAMLToJSONFromFile_FileNotFound(t *testing.T) {
	_, err := core.YAMLToJSONFromFile("nonexistent.yaml")
	assert.Error(t, err)
}

func TestJSONToYAMLFromFile(t *testing.T) {
	// Create a temporary JSON file
	jsonFileName := "test.json"
	jsonContent := []byte(`{"key":"value","nested":{"key":"nestedValue"}}`)
	err := os.WriteFile(jsonFileName, jsonContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(jsonFileName)

	// Test the function
	result, err := core.JSONToYAMLFromFile(jsonFileName)
	assert.NoError(t, err)
	assert.Equal(t, "key: value\nnested:\n  key: nestedValue\n", string(result))
}

func TestJSONToYAMLFromFile_FileNotFound(t *testing.T) {
	_, err := core.JSONToYAMLFromFile("nonexistent.json")
	assert.Error(t, err)
}

func TestReadJSONFile(t *testing.T) {
	// Create a temporary JSON file
	jsonFileName := "test.json"
	jsonContent := []byte(`{"key":"value"}`)
	err := os.WriteFile(jsonFileName, jsonContent, 0644)
	assert.NoError(t, err)
	defer os.Remove(jsonFileName)

	// Test the function
	result, err := core.ReadJSONFile(jsonFileName)
	assert.NoError(t, err)
	assert.Equal(t, jsonContent, result)
}

func TestReadJSONFile_FileNotFound(t *testing.T) {
	_, err := core.ReadJSONFile("nonexistent.json")
	assert.Error(t, err)
}
