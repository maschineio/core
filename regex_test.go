package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestStringMatchesRegex(t *testing.T) {
	tests := []struct {
		name     string
		pattern  string
		value    string
		expected bool
		wantErr  bool
	}{
		{
			name:     "exact match",
			pattern:  "foo.log",
			value:    "foo.log",
			expected: true,
		},
		{
			name:     "single wildcard prefix",
			pattern:  "*.log",
			value:    "zebra.log",
			expected: true,
		},
		{
			name:     "single wildcard suffix",
			pattern:  "foo*",
			value:    "foobar",
			expected: true,
		},
		{
			name:     "wildcard in middle",
			pattern:  "foo*.log",
			value:    "foo23.log",
			expected: true,
		},
		{
			name:     "multiple wildcards",
			pattern:  "foo*.*",
			value:    "foobar.zebra",
			expected: true,
		},
		{
			name:     "no match",
			pattern:  "foo.log",
			value:    "bar.log",
			expected: false,
		},
		{
			name:     "empty pattern",
			pattern:  "",
			value:    "test",
			expected: true,
			wantErr:  false,
		},
		{
			name:     "empty value",
			pattern:  "test",
			value:    "",
			expected: false,
		},
		{
			name:     "multiple consecutive wildcards",
			pattern:  "foo**.log",
			value:    "foobar.log",
			expected: true,
		},
		{
			name:     "special regex characters",
			pattern:  "test.*.txt",
			value:    "test.abc.txt",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := core.StringMatchesRegex(tt.pattern, tt.value)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, result)
		})
	}
}
