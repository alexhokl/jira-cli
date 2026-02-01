package cmd

import (
	"testing"
)

func TestParseContextId(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		expected    int64
		expectError bool
	}{
		{
			name:        "valid numeric ID",
			id:          "10001",
			expected:    10001,
			expectError: false,
		},
		{
			name:        "valid large ID",
			id:          "123456789",
			expected:    123456789,
			expectError: false,
		},
		{
			name:        "valid single digit ID",
			id:          "5",
			expected:    5,
			expectError: false,
		},
		{
			name:        "zero ID",
			id:          "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "invalid - non-numeric",
			id:          "abc",
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid - empty string",
			id:          "",
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid - special characters",
			id:          "!@#",
			expected:    0,
			expectError: true,
		},
		{
			name:        "negative number",
			id:          "-100",
			expected:    -100,
			expectError: false,
		},
		{
			name:        "very large number",
			id:          "9223372036854775807",
			expected:    9223372036854775807,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseContextId(tt.id)
			if tt.expectError {
				if err == nil {
					t.Errorf("parseContextId(%q) expected error but got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("parseContextId(%q) unexpected error: %v", tt.id, err)
				return
			}
			if result != tt.expected {
				t.Errorf("parseContextId(%q) = %d, want %d", tt.id, result, tt.expected)
			}
		})
	}
}
