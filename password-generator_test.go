package core_test

import (
	"testing"

	"maschine.io/core"
)

func TestGenerateSimplePassword(t *testing.T) {
	tests := []struct {
		name    string
		length  int
		wantLen int
		wantErr bool
	}{
		{
			name:    "valid password length 8",
			length:  8,
			wantLen: 8,
			wantErr: false,
		},
		{
			name:    "valid password length 16",
			length:  16,
			wantLen: 16,
			wantErr: false,
		},
		{
			name:    "invalid zero length",
			length:  0,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "invalid negative length",
			length:  -1,
			wantLen: 0,
			wantErr: true,
		},
		{
			name:    "large password length",
			length:  32,
			wantLen: 32,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := core.GenerateSimplePassword(tt.length)

			if (err != nil) != tt.wantErr {
				t.Errorf("GenerateSimplePassword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && len(got) != tt.wantLen {
				t.Errorf("GenerateSimplePassword() got length = %v, want %v", len(got), tt.wantLen)
			}

			// Only test uniqueness for valid lengths
			if tt.length > 0 && !tt.wantErr {
				another, _ := core.GenerateSimplePassword(tt.length)
				if got == another {
					t.Error("GenerateSimplePassword() generated identical passwords")
				}
			}
		})
	}
}
