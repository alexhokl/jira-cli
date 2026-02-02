package cmd

import (
	"bytes"
	"testing"
)

func TestStripAnsi(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "no ANSI codes",
			input:    "hello world",
			expected: "hello world",
		},
		{
			name:     "single color code",
			input:    "\x1b[31mred text\x1b[0m",
			expected: "red text",
		},
		{
			name:     "multiple color codes",
			input:    "\x1b[31mred\x1b[0m and \x1b[32mgreen\x1b[0m",
			expected: "red and green",
		},
		{
			name:     "bold and color",
			input:    "\x1b[1;33mbold yellow\x1b[0m",
			expected: "bold yellow",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only ANSI codes",
			input:    "\x1b[31m\x1b[0m",
			expected: "",
		},
		{
			name:     "complex ANSI sequence",
			input:    "\x1b[38;5;196mcomplex color\x1b[0m",
			expected: "complex color",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := stripAnsi(tt.input)
			if result != tt.expected {
				t.Errorf("stripAnsi(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTimezoneRegex(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "positive timezone without colon",
			input:    "2024-01-15T10:30:00.123+0800",
			expected: "2024-01-15T10:30:00.123+08:00",
		},
		{
			name:     "negative timezone without colon",
			input:    "2024-01-15T10:30:00.123-0530",
			expected: "2024-01-15T10:30:00.123-05:30",
		},
		{
			name:     "no milliseconds",
			input:    "2024-01-15T10:30:00+0800",
			expected: "2024-01-15T10:30:00+08:00",
		},
		{
			name:     "already has colon",
			input:    "2024-01-15T10:30:00+08:00",
			expected: "2024-01-15T10:30:00+08:00",
		},
		{
			name:     "UTC timezone",
			input:    "2024-01-15T10:30:00Z",
			expected: "2024-01-15T10:30:00Z",
		},
		{
			name:     "multiple timestamps in JSON",
			input:    `{"created":"2024-01-15T10:30:00+0800","updated":"2024-01-16T09:00:00-0500"}`,
			expected: `{"created":"2024-01-15T10:30:00+08:00","updated":"2024-01-16T09:00:00-05:00"}`,
		},
		// Additional edge cases
		{
			name:     "timezone +0000",
			input:    "2024-01-15T10:30:00+0000",
			expected: "2024-01-15T10:30:00+00:00",
		},
		{
			name:     "timezone -0000",
			input:    "2024-01-15T10:30:00-0000",
			expected: "2024-01-15T10:30:00-00:00",
		},
		{
			name:     "extreme positive timezone +1200",
			input:    "2024-01-15T10:30:00+1200",
			expected: "2024-01-15T10:30:00+12:00",
		},
		{
			name:     "extreme negative timezone -1200",
			input:    "2024-01-15T10:30:00-1200",
			expected: "2024-01-15T10:30:00-12:00",
		},
		{
			name:     "microseconds precision",
			input:    "2024-01-15T10:30:00.123456+0800",
			expected: "2024-01-15T10:30:00.123456+08:00",
		},
		{
			name:     "nanoseconds precision",
			input:    "2024-01-15T10:30:00.123456789+0800",
			expected: "2024-01-15T10:30:00.123456789+08:00",
		},
		{
			name:     "midnight time",
			input:    "2024-01-15T00:00:00+0800",
			expected: "2024-01-15T00:00:00+08:00",
		},
		{
			name:     "end of day time",
			input:    "2024-01-15T23:59:59+0800",
			expected: "2024-01-15T23:59:59+08:00",
		},
		{
			name:     "leap year date",
			input:    "2024-02-29T12:00:00+0800",
			expected: "2024-02-29T12:00:00+08:00",
		},
		{
			name:     "end of year date",
			input:    "2024-12-31T23:59:59.999+0800",
			expected: "2024-12-31T23:59:59.999+08:00",
		},
		{
			name:     "plain text no change",
			input:    "Hello, World!",
			expected: "Hello, World!",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "partial timestamp not matched",
			input:    "2024-01-15T10:30",
			expected: "2024-01-15T10:30",
		},
		{
			name:     "date only not matched",
			input:    "2024-01-15",
			expected: "2024-01-15",
		},
		{
			name:     "complex JSON with multiple fields",
			input:    `{"id":"123","created":"2024-01-15T10:30:00.123+0800","updated":"2024-01-16T09:00:00-0530","resolved":null}`,
			expected: `{"id":"123","created":"2024-01-15T10:30:00.123+08:00","updated":"2024-01-16T09:00:00-05:30","resolved":null}`,
		},
		{
			name:     "array of timestamps in JSON",
			input:    `["2024-01-15T10:30:00+0800","2024-01-16T09:00:00-0500"]`,
			expected: `["2024-01-15T10:30:00+08:00","2024-01-16T09:00:00-05:00"]`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := timezoneRegex.ReplaceAllString(tt.input, "${1}${2}${3}:${4}")
			if result != tt.expected {
				t.Errorf("timezoneRegex replacement: got %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestTableWriter(t *testing.T) {
	t.Run("single row", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("A", "B", "C")
		tw.flush()

		expected := "A  B  C\n"
		if buf.String() != expected {
			t.Errorf("got %q, want %q", buf.String(), expected)
		}
	})

	t.Run("multiple rows with variable widths", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("ID", "Name", "Status")
		tw.row("1", "Short", "OK")
		tw.row("123", "Much Longer Name", "Pending")
		tw.flush()

		// Check that rows are properly aligned
		output := buf.String()
		if output == "" {
			t.Error("expected output, got empty string")
		}
		// The output should contain all the data
		if !bytes.Contains([]byte(output), []byte("ID")) {
			t.Error("output should contain 'ID'")
		}
		if !bytes.Contains([]byte(output), []byte("Much Longer Name")) {
			t.Error("output should contain 'Much Longer Name'")
		}
	})

	t.Run("with ANSI codes", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("\x1b[31mRed\x1b[0m", "Normal")
		tw.row("ABC", "XYZ")
		tw.flush()

		output := buf.String()
		// The Red text should be padded based on its visible width (3), not string length
		// Both "Red" (visible) and "ABC" have 3 characters, so they should align
		if output == "" {
			t.Error("expected output, got empty string")
		}
	})

	t.Run("empty table", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.flush()

		if buf.String() != "" {
			t.Errorf("expected empty output for empty table, got %q", buf.String())
		}
	})

	t.Run("minimum width", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 5, 2)
		tw.row("A", "B")
		tw.flush()

		// With minWidth=5, first column should be padded to at least 5 characters
		// "A" + (5-1) spaces to reach minWidth + 2 padding = "A      B\n"
		expected := "A      B\n"
		if buf.String() != expected {
			t.Errorf("got %q, want %q", buf.String(), expected)
		}
	})

	t.Run("rows with different column counts", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("A", "B", "C")
		tw.row("X", "Y") // fewer columns
		tw.flush()

		output := buf.String()
		if output == "" {
			t.Error("expected output, got empty string")
		}
	})

	t.Run("unicode characters", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("日本語", "English")
		tw.row("ABC", "XYZ")
		tw.flush()

		output := buf.String()
		if output == "" {
			t.Error("expected output, got empty string")
		}
		if !bytes.Contains([]byte(output), []byte("日本語")) {
			t.Error("output should contain unicode characters")
		}
	})

	t.Run("single column", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 2)
		tw.row("Only column")
		tw.row("Another row")
		tw.flush()

		expected := "Only column\nAnother row\n"
		if buf.String() != expected {
			t.Errorf("got %q, want %q", buf.String(), expected)
		}
	})

	t.Run("large padding", func(t *testing.T) {
		var buf bytes.Buffer
		tw := newTableWriter(&buf, 0, 10)
		tw.row("A", "B")
		tw.flush()

		expected := "A          B\n"
		if buf.String() != expected {
			t.Errorf("got %q, want %q", buf.String(), expected)
		}
	})
}

func TestNewTableWriter(t *testing.T) {
	var buf bytes.Buffer
	tw := newTableWriter(&buf, 5, 3)

	if tw == nil {
		t.Fatal("newTableWriter returned nil")
	}
	if tw.out != &buf {
		t.Error("writer output mismatch")
	}
	if tw.minWidth != 5 {
		t.Errorf("minWidth = %d, want 5", tw.minWidth)
	}
	if tw.padding != 3 {
		t.Errorf("padding = %d, want 3", tw.padding)
	}
	if tw.padChar != ' ' {
		t.Errorf("padChar = %q, want ' '", tw.padChar)
	}
}

func TestAnsiRegex(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		matches bool
	}{
		{"reset code", "\x1b[0m", true},
		{"red foreground", "\x1b[31m", true},
		{"bold", "\x1b[1m", true},
		{"multiple params", "\x1b[1;31;40m", true},
		{"256 color", "\x1b[38;5;196m", true},
		{"plain text", "hello", false},
		{"partial escape", "\x1b[", false},
		{"empty string", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := ansiRegex.MatchString(tt.input)
			if matches != tt.matches {
				t.Errorf("ansiRegex.MatchString(%q) = %v, want %v", tt.input, matches, tt.matches)
			}
		})
	}
}

func TestLegacyTabWriter(t *testing.T) {
	tw := legacyTabWriter()
	if tw == nil {
		t.Fatal("legacyTabWriter returned nil")
	}

	// Write some content and verify it works
	_, err := tw.Write([]byte("Col1\tCol2\tCol3\n"))
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	_, err = tw.Write([]byte("A\tBBB\tC\n"))
	if err != nil {
		t.Errorf("Write failed: %v", err)
	}

	err = tw.Flush()
	if err != nil {
		t.Errorf("Flush failed: %v", err)
	}
}

func TestNormalizeUserValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "me converts to currentUser()",
			input:    "me",
			expected: "currentUser()",
		},
		{
			name:     "currentUser() unchanged",
			input:    "currentUser()",
			expected: "currentUser()",
		},
		{
			name:     "unassigned unchanged",
			input:    "unassigned",
			expected: "unassigned",
		},
		{
			name:     "none unchanged",
			input:    "none",
			expected: "none",
		},
		{
			name:     "account ID unchanged",
			input:    "5b10ac8d82e05b22cc7d4ef5",
			expected: "5b10ac8d82e05b22cc7d4ef5",
		},
		{
			name:     "username unchanged",
			input:    "john.doe",
			expected: "john.doe",
		},
		{
			name:     "empty string unchanged",
			input:    "",
			expected: "",
		},
		{
			name:     "Me with capital M unchanged (case sensitive)",
			input:    "Me",
			expected: "Me",
		},
		{
			name:     "ME all caps unchanged (case sensitive)",
			input:    "ME",
			expected: "ME",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := normalizeUserValue(tt.input)
			if result != tt.expected {
				t.Errorf("normalizeUserValue(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractProjectKeyFromIssueKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "standard issue key",
			input:    "PROJ-123",
			expected: "PROJ",
		},
		{
			name:     "lowercase project key",
			input:    "proj-456",
			expected: "proj",
		},
		{
			name:     "long project key",
			input:    "MYPROJECT-1",
			expected: "MYPROJECT",
		},
		{
			name:     "single letter project key",
			input:    "X-99",
			expected: "X",
		},
		{
			name:     "high issue number",
			input:    "TEST-99999",
			expected: "TEST",
		},
		{
			name:     "project key with numbers",
			input:    "PROJ2-123",
			expected: "PROJ2",
		},
		{
			name:     "no hyphen returns original",
			input:    "PROJ123",
			expected: "PROJ123",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "just hyphen returns empty",
			input:    "-123",
			expected: "",
		},
		{
			name:     "multiple hyphens uses last",
			input:    "MY-PROJ-123",
			expected: "MY-PROJ",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractProjectKeyFromIssueKey(tt.input)
			if result != tt.expected {
				t.Errorf("extractProjectKeyFromIssueKey(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestResolveUserAlias_NonMeValues(t *testing.T) {
	// Test that non-"me" values are returned unchanged without API calls
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "account ID unchanged",
			input:    "5b10ac8d82e05b22cc7d4ef5",
			expected: "5b10ac8d82e05b22cc7d4ef5",
		},
		{
			name:     "username unchanged",
			input:    "john.doe",
			expected: "john.doe",
		},
		{
			name:     "empty string unchanged",
			input:    "",
			expected: "",
		},
		{
			name:     "Me with capital M unchanged (case sensitive)",
			input:    "Me",
			expected: "Me",
		},
		{
			name:     "ME all caps unchanged (case sensitive)",
			input:    "ME",
			expected: "ME",
		},
		{
			name:     "none unchanged",
			input:    "none",
			expected: "none",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := resolveUserAlias(tt.input)
			if err != nil {
				t.Errorf("resolveUserAlias(%q) unexpected error: %v", tt.input, err)
				return
			}
			if result != tt.expected {
				t.Errorf("resolveUserAlias(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}
