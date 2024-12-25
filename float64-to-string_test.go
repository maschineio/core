package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestFloat64ToString(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "positive integer",
			input:    42.0,
			expected: "42",
		},
		{
			name:     "zero",
			input:    0.0,
			expected: "0",
		},
		{
			name:     "negative integer",
			input:    -42.0,
			expected: "-42",
		},
		{
			name:     "positive decimal truncation",
			input:    1.7,
			expected: "1",
		},
		{
			name:     "negative decimal truncation",
			input:    -1.7,
			expected: "-1",
		},
		{
			name:     "large number",
			input:    1000000.0,
			expected: "1000000",
		},
		{
			name:     "small decimal",
			input:    0.1,
			expected: "0",
		},
		{
			name:     "half number",
			input:    0.5,
			expected: "0",
		},
		{
			name:     "one",
			input:    1.0,
			expected: "1",
		},
		{
			name:     "negative one",
			input:    -1.0,
			expected: "-1",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := core.Float64ToString(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}
