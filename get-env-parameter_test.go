package core_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestGetEnvParameter(t *testing.T) {
	tests := []struct {
		name        string
		envKey      string
		envValue    string
		want        string
		expectError bool
	}{
		{
			name:        "existing env variable",
			envKey:      "TEST_VAR",
			envValue:    "test_value",
			want:        "test_value",
			expectError: false,
		},
		{
			name:        "non-existent env variable",
			envKey:      "NONEXISTENT_VAR",
			envValue:    "",
			want:        "",
			expectError: true,
		},
		{
			name:        "empty env variable",
			envKey:      "EMPTY_VAR",
			envValue:    "",
			want:        "",
			expectError: true,
		},
		{
			name:        "special characters in value",
			envKey:      "SPECIAL_CHARS",
			envValue:    "test!@#$%^&*()",
			want:        "test!@#$%^&*()",
			expectError: false,
		},
		{
			name:        "whitespace value",
			envKey:      "WHITESPACE",
			envValue:    "  ",
			want:        "  ",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Setup
			if tt.envValue != "" {
				os.Setenv(tt.envKey, tt.envValue)
				defer os.Unsetenv(tt.envKey)
			}

			// Test
			got, err := core.GetEnvParameter(tt.envKey)

			// Assert
			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), "not set")
			} else {
				assert.NoError(t, err)
			}
			assert.Equal(t, tt.want, got)
		})
	}
}
