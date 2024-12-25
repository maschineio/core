package core_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"maschine.io/core"
)

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
