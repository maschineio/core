package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestRemoveSlash(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "path with leading slash",
			input:    "/test/path",
			expected: "test/path",
		},
		{
			name:     "path without leading slash",
			input:    "test/path",
			expected: "test/path",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "single slash",
			input:    "/",
			expected: "",
		},
		{
			name:     "multiple leading slashes",
			input:    "//test/path",
			expected: "/test/path",
		},
		{
			name:     "path with special characters",
			input:    "/test-path_123",
			expected: "test-path_123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := core.RemoveSlash(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
