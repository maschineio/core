package core_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"maschine.io/core"
)

func TestGetAbsFilePath(t *testing.T) {
	// Setup test environment
	tmpDir := t.TempDir()
	currentDir, err := os.Getwd()
	require.NoError(t, err)

	// Create test file
	testFile := filepath.Join(tmpDir, "test.txt")
	err = os.WriteFile(testFile, []byte("test"), 0644)
	require.NoError(t, err)

	// Create test directory
	testDir := filepath.Join(tmpDir, "testdir")
	err = os.Mkdir(testDir, 0755)
	require.NoError(t, err)

	tests := []struct {
		name        string
		input       string
		check       bool
		expected    string
		expectError bool
	}{
		{
			name:        "absolute path with check true",
			input:       testFile,
			check:       true,
			expected:    testFile,
			expectError: false,
		},
		{
			name:        "absolute path with check false",
			input:       filepath.Join(currentDir, "nonexistent.txt"),
			check:       false,
			expected:    filepath.Join(currentDir, "nonexistent.txt"),
			expectError: false,
		},
		{
			name:        "relative path with check true",
			input:       filepath.Base(testFile),
			check:       true,
			expected:    filepath.Join(currentDir, filepath.Base(testFile)),
			expectError: true, // File doesn't exist in current dir
		},
		{
			name:        "relative path with check false",
			input:       "test.txt",
			check:       false,
			expected:    filepath.Join(currentDir, "test.txt"),
			expectError: false,
		},
		{
			name:        "empty path with check false",
			input:       "",
			check:       false,
			expected:    currentDir,
			expectError: false,
		},
		{
			name:        "path with special characters",
			input:       "./test/../test.txt",
			check:       false,
			expected:    filepath.Join(currentDir, "test.txt"),
			expectError: false,
		},
		{
			name:        "directory path with check true",
			input:       testDir,
			check:       true,
			expected:    "",
			expectError: true,
		},
		{
			name:        "non-existent file with check true",
			input:       filepath.Join(tmpDir, "nonexistent.txt"),
			check:       true,
			expected:    "",
			expectError: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := core.GetAbsFilePath(tt.input, tt.check)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}

func TestGetAbsDirectory(t *testing.T) {
	currentDir, err := os.Getwd()
	require.NoError(t, err)

	tests := []struct {
		name        string
		input       string
		expected    string
		expectError bool
	}{
		{
			name:        "absolute path",
			input:       filepath.Join(currentDir, "testdata"),
			expected:    filepath.Join(currentDir, "testdata"),
			expectError: false,
		},
		{
			name:        "relative path",
			input:       "testdata",
			expected:    filepath.Join(currentDir, "testdata"),
			expectError: false,
		},
		{
			name:        "empty path",
			input:       "",
			expected:    currentDir,
			expectError: false,
		},
		{
			name:        "path with special characters",
			input:       "./test/../testdata",
			expected:    filepath.Join(currentDir, "testdata"),
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := core.GetAbsDirectory(tt.input)

			if tt.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
