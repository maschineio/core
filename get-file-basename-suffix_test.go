package core_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"maschine.io/core"
)

func TestGetFileBasenameAndSuffix(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		wantBasename  string
		wantExtension string
	}{
		{
			name:          "regular file",
			input:         "document.txt",
			wantBasename:  "document",
			wantExtension: ".txt",
		},
		{
			name:          "no extension",
			input:         "README",
			wantBasename:  "README",
			wantExtension: "",
		},
		{
			name:          "empty string",
			input:         "",
			wantBasename:  "",
			wantExtension: "",
		},
		{
			name:          "multiple dots",
			input:         "archive.tar.gz",
			wantBasename:  "archive.tar",
			wantExtension: ".gz",
		},
		{
			name:          "only extension",
			input:         ".gitignore",
			wantBasename:  "",
			wantExtension: ".gitignore",
		},
		{
			name:          "with path separators",
			input:         "path/to/file.json",
			wantBasename:  "path/to/file",
			wantExtension: ".json",
		},
		{
			name:          "hidden file with extension",
			input:         ".config.json",
			wantBasename:  ".config",
			wantExtension: ".json",
		},
		{
			name:          "source code file",
			input:         "main.go",
			wantBasename:  "main",
			wantExtension: ".go",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBasename, gotExtension := core.GetFileBasenameAndSuffix(tt.input)
			assert.Equal(t, tt.wantBasename, gotBasename)
			assert.Equal(t, tt.wantExtension, gotExtension)
		})
	}
}
