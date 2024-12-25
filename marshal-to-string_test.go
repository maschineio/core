package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestMarshalToString(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    string
		wantErr bool
	}{
		{
			name:    "string value",
			input:   "test",
			want:    `"test"`,
			wantErr: false,
		},
		{
			name:    "integer value",
			input:   123,
			want:    "123",
			wantErr: false,
		},
		{
			name:    "map value",
			input:   map[string]string{"key": "value"},
			want:    `{"key":"value"}`,
			wantErr: false,
		},
		{
			name:    "nil value",
			input:   nil,
			want:    "null",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := core.MarshalToString(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestUnmarshalJsonBytesToAny(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "valid json bytes",
			input:   []byte(`{"key":"value"}`),
			want:    map[string]interface{}{"key": "value"},
			wantErr: false,
		},
		{
			name:    "invalid json bytes",
			input:   []byte(`{invalid}`),
			want:    nil,
			wantErr: true,
		},
		{
			name:    "non-bytes input",
			input:   "not bytes",
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := core.UnmarshalJsonBytesToAny(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTypeToAny(t *testing.T) {
	tests := []struct {
		name    string
		input   interface{}
		want    interface{}
		wantErr bool
	}{
		{
			name:    "string to any",
			input:   "test",
			want:    "test",
			wantErr: false,
		},
		{
			name:    "struct to any",
			input:   struct{ Name string }{"test"},
			want:    map[string]interface{}{"Name": "test"},
			wantErr: false,
		},
		{
			name:    "map to any",
			input:   map[string]int{"value": 123},
			want:    map[string]interface{}{"value": float64(123)},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := core.TypeToAny(tt.input)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, tt.want, got)
		})
	}
}
