package core_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestWriteFileSuccess(t *testing.T) {
	fileName := "testfile.txt"
	data := []byte("Hello, World!")

	err := core.WriteFile(fileName, data)
	assert.NoError(t, err)

	// Verify the file content
	content, err := os.ReadFile(fileName)
	assert.NoError(t, err)
	assert.Equal(t, data, content)

	// Clean up
	os.Remove(fileName)
}

func TestWriteFileEmptyData(t *testing.T) {
	fileName := "emptyfile.txt"
	data := []byte("")

	err := core.WriteFile(fileName, data)
	assert.NoError(t, err)

	// Verify the file content
	content, err := os.ReadFile(fileName)
	assert.NoError(t, err)
	assert.Equal(t, data, content)

	// Clean up
	os.Remove(fileName)
}

func TestWriteFileInvalidPath(t *testing.T) {
	fileName := "/invalidpath/testfile.txt"
	data := []byte("Hello, World!")

	err := core.WriteFile(fileName, data)
	assert.Error(t, err)
}
