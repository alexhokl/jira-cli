package cmd

import (
	"strings"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expectError bool
		checkFormat bool // if true, just verify the output is RFC3339 format
	}{
		{
			name:        "valid date",
			input:       "2024-01-15",
			expectError: false,
			checkFormat: true,
		},
		{
			name:        "leap year date",
			input:       "2024-02-29",
			expectError: false,
			checkFormat: true,
		},
		{
			name:        "end of year",
			input:       "2024-12-31",
			expectError: false,
			checkFormat: true,
		},
		{
			name:        "beginning of year",
			input:       "2024-01-01",
			expectError: false,
			checkFormat: true,
		},
		{
			name:        "invalid format - wrong separator",
			input:       "2024/01/15",
			expectError: true,
		},
		{
			name:        "invalid format - no separators",
			input:       "20240115",
			expectError: true,
		},
		{
			name:        "invalid date - non-leap year Feb 29",
			input:       "2023-02-29",
			expectError: true,
		},
		{
			name:        "invalid date - month 13",
			input:       "2024-13-01",
			expectError: true,
		},
		{
			name:        "invalid date - day 32",
			input:       "2024-01-32",
			expectError: true,
		},
		{
			name:        "empty string",
			input:       "",
			expectError: true,
		},
		{
			name:        "random text",
			input:       "not-a-date",
			expectError: true,
		},
		{
			name:        "partial date",
			input:       "2024-01",
			expectError: true,
		},
		{
			name:        "date with time",
			input:       "2024-01-15T10:30:00",
			expectError: true, // parseDate expects date-only format
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseDate(tt.input)

			if tt.expectError {
				if err == nil {
					t.Errorf("parseDate(%q) expected error, got result %q", tt.input, result)
				}
				return
			}

			if err != nil {
				t.Errorf("parseDate(%q) unexpected error: %v", tt.input, err)
				return
			}

			if tt.checkFormat {
				// Verify the output is valid RFC3339 format
				_, parseErr := time.Parse(time.RFC3339, result)
				if parseErr != nil {
					t.Errorf("parseDate(%q) = %q, not valid RFC3339: %v", tt.input, result, parseErr)
				}

				// Verify the date portion matches the input
				if !strings.HasPrefix(result, tt.input) {
					t.Errorf("parseDate(%q) = %q, date portion doesn't match input", tt.input, result)
				}
			}
		})
	}
}

func TestParseDateTimezoneHandling(t *testing.T) {
	// Test that parseDate uses local timezone
	input := "2024-06-15"
	result, err := parseDate(input)
	if err != nil {
		t.Fatalf("parseDate(%q) unexpected error: %v", input, err)
	}

	// Parse the result back
	parsed, err := time.Parse(time.RFC3339, result)
	if err != nil {
		t.Fatalf("Failed to parse result %q: %v", result, err)
	}

	// Verify the time is midnight in local timezone
	localMidnight, _ := time.ParseInLocation("2006-01-02", input, time.Local)
	if !parsed.Equal(localMidnight) {
		t.Errorf("parseDate(%q) = %q, expected local midnight %v", input, result, localMidnight)
	}
}

func TestParseDateRFC3339Output(t *testing.T) {
	// Verify the output format is exactly RFC3339
	input := "2024-07-04"
	result, err := parseDate(input)
	if err != nil {
		t.Fatalf("parseDate(%q) unexpected error: %v", input, err)
	}

	// RFC3339 format should contain T separator and timezone info
	if !strings.Contains(result, "T") {
		t.Errorf("parseDate(%q) = %q, missing 'T' separator for RFC3339", input, result)
	}

	// Should have timezone offset (either Z or +/-HH:MM)
	hasTimezone := strings.HasSuffix(result, "Z") ||
		strings.Contains(result[len(result)-6:], "+") ||
		strings.Contains(result[len(result)-6:], "-")
	if !hasTimezone {
		t.Errorf("parseDate(%q) = %q, missing timezone info for RFC3339", input, result)
	}
}

func TestParseDateBoundaryDates(t *testing.T) {
	// Test boundary dates
	boundaryDates := []string{
		"1970-01-01", // Unix epoch
		"2000-01-01", // Y2K
		"2038-01-19", // Near 32-bit Unix timestamp limit
		"9999-12-31", // Far future
	}

	for _, date := range boundaryDates {
		t.Run(date, func(t *testing.T) {
			result, err := parseDate(date)
			if err != nil {
				t.Errorf("parseDate(%q) unexpected error: %v", date, err)
				return
			}

			// Verify result is valid
			_, parseErr := time.Parse(time.RFC3339, result)
			if parseErr != nil {
				t.Errorf("parseDate(%q) = %q, not valid RFC3339: %v", date, result, parseErr)
			}
		})
	}
}
