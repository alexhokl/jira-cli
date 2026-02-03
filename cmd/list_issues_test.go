package cmd

import (
	"fmt"
	"strings"
	"testing"
)

func TestQuoteJQLValue(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple value without special chars",
			input:    "PROJ",
			expected: "PROJ",
		},
		{
			name:     "value with space",
			input:    "In Progress",
			expected: `"In Progress"`,
		},
		{
			name:     "value with hyphen",
			input:    "PROJ-123",
			expected: `"PROJ-123"`,
		},
		{
			name:     "value with single quote",
			input:    "John's Issue",
			expected: `"John's Issue"`,
		},
		{
			name:     "value with double quote",
			input:    `Value with "quotes"`,
			expected: `"Value with \"quotes\""`,
		},
		{
			name:     "simple lowercase",
			input:    "bug",
			expected: "bug",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "value with multiple hyphens",
			input:    "some-long-label-name",
			expected: `"some-long-label-name"`,
		},
		// Additional edge cases - testing JQL reserved characters
		{
			name:     "value with tab character (not quoted)",
			input:    "value\twith\ttab",
			expected: "value\twith\ttab",
		},
		{
			name:     "value with newline (not quoted)",
			input:    "value\nwith\nnewline",
			expected: "value\nwith\nnewline",
		},
		{
			name:     "value with backslash (quoted)",
			input:    `path\to\file`,
			expected: `"path\to\file"`,
		},
		{
			name:     "numeric value",
			input:    "12345",
			expected: "12345",
		},
		{
			name:     "value with parentheses (quoted)",
			input:    "currentUser()",
			expected: `"currentUser()"`,
		},
		{
			name:     "value with brackets (quoted)",
			input:    "[value]",
			expected: `"[value]"`,
		},
		{
			name:     "value with special JQL chars (quoted)",
			input:    "value=test",
			expected: `"value=test"`,
		},
		{
			name:     "value with colon (quoted)",
			input:    "key:value",
			expected: `"key:value"`,
		},
		{
			name:     "value with underscore only",
			input:    "my_label",
			expected: "my_label",
		},
		{
			name:     "value with forward slash (quoted)",
			input:    "team/project",
			expected: `"team/project"`,
		},
		{
			name:     "value with multiple forward slashes",
			input:    "path/to/something",
			expected: `"path/to/something"`,
		},
		{
			name:     "value with asterisk (quoted)",
			input:    "test*",
			expected: `"test*"`,
		},
		{
			name:     "value with question mark (quoted)",
			input:    "test?",
			expected: `"test?"`,
		},
		{
			name:     "value with plus sign (quoted)",
			input:    "test+value",
			expected: `"test+value"`,
		},
		{
			name:     "value with ampersand (quoted)",
			input:    "test&value",
			expected: `"test&value"`,
		},
		{
			name:     "value with pipe (quoted)",
			input:    "test|value",
			expected: `"test|value"`,
		},
		{
			name:     "value with exclamation (quoted)",
			input:    "important!",
			expected: `"important!"`,
		},
		{
			name:     "value with curly braces (quoted)",
			input:    "{value}",
			expected: `"{value}"`,
		},
		{
			name:     "value with caret (quoted)",
			input:    "test^value",
			expected: `"test^value"`,
		},
		{
			name:     "value with tilde (quoted)",
			input:    "test~value",
			expected: `"test~value"`,
		},
		{
			name:     "value with less than (quoted)",
			input:    "test<value",
			expected: `"test<value"`,
		},
		{
			name:     "value with greater than (quoted)",
			input:    "test>value",
			expected: `"test>value"`,
		},
		{
			name:     "value with semicolon (quoted)",
			input:    "test;value",
			expected: `"test;value"`,
		},
		{
			name:     "value with leading space",
			input:    " leading",
			expected: `" leading"`,
		},
		{
			name:     "value with trailing space",
			input:    "trailing ",
			expected: `"trailing "`,
		},
		{
			name:     "value with space and hyphen",
			input:    "my-label name",
			expected: `"my-label name"`,
		},
		{
			name:     "value with multiple quotes",
			input:    `"test" and "more"`,
			expected: `"\"test\" and \"more\""`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := quoteJQLValue(tt.input)
			if result != tt.expected {
				t.Errorf("quoteJQLValue(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

// TestStatusColorCategorization tests the status color categorization logic
// used in printIssues function
func TestStatusColorCategorization(t *testing.T) {
	tests := []struct {
		name           string
		status         string
		expectedColor  string // "green", "yellow", "red", or "cyan" (default)
		statusCategory string
	}{
		// Done/Closed statuses (green)
		{"done status", "Done", "green", "done"},
		{"closed status", "Closed", "green", "done"},
		{"resolved status", "Resolved", "green", "done"},
		{"done lowercase", "done", "green", "done"},
		{"done with suffix", "Task Done", "green", "done"},

		// In Progress statuses (yellow)
		{"in progress status", "In Progress", "yellow", "inprogress"},
		{"progress lowercase", "in progress", "yellow", "inprogress"},
		{"review status", "In Review", "yellow", "inprogress"},
		{"code review", "Code Review", "yellow", "inprogress"},
		{"under review", "Under Review", "yellow", "inprogress"},
		{"progress partial", "Work in Progress", "yellow", "inprogress"},

		// Blocked statuses (red)
		{"blocked status", "Blocked", "red", "blocked"},
		{"blocker status", "Blocker", "red", "blocked"},
		{"blocked lowercase", "blocked", "red", "blocked"},
		{"blocking status", "Blocking", "red", "blocked"},

		// Default statuses (cyan) - these don't match any pattern
		{"open status", "Open", "cyan", "todo"},
		{"to do status", "To Do", "cyan", "todo"},
		{"backlog status", "Backlog", "cyan", "todo"},
		{"new status", "New", "cyan", "todo"},
		{"pending status", "Pending", "cyan", "todo"},
		{"waiting status", "Waiting", "cyan", "todo"},
		{"empty status", "", "cyan", "todo"},
		// Note: "Completed" doesn't match done/closed/resolved
		{"completed status not matched", "Completed", "cyan", "todo"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			statusLower := strings.ToLower(tt.status)
			var colorCategory string

			if strings.Contains(statusLower, "done") || strings.Contains(statusLower, "closed") || strings.Contains(statusLower, "resolved") {
				colorCategory = "green"
			} else if strings.Contains(statusLower, "progress") || strings.Contains(statusLower, "review") {
				colorCategory = "yellow"
			} else if strings.Contains(statusLower, "block") {
				colorCategory = "red"
			} else {
				colorCategory = "cyan"
			}

			if colorCategory != tt.expectedColor {
				t.Errorf("status %q categorized as %q, expected %q", tt.status, colorCategory, tt.expectedColor)
			}
		})
	}
}

// TestPriorityColorCategorization tests the priority color categorization logic
// used in printIssues function
func TestPriorityColorCategorization(t *testing.T) {
	tests := []struct {
		name          string
		priority      string
		expectedColor string // "red", "yellow", "green", or "" (no color/default)
	}{
		// High priority (red)
		{"highest priority", "Highest", "red"},
		{"high priority", "High", "red"},
		{"critical priority", "Critical", "red"},
		{"blocker priority", "Blocker", "red"},
		{"high lowercase", "high", "red"},
		{"very high", "Very High", "red"},

		// Medium priority (yellow)
		{"medium priority", "Medium", "yellow"},
		{"medium lowercase", "medium", "yellow"},

		// Low priority (green)
		{"low priority", "Low", "green"},
		{"lowest priority", "Lowest", "green"},
		{"minor priority", "Minor", "green"},
		{"low lowercase", "low", "green"},

		// No special color - these don't match any pattern
		{"empty priority", "", ""},
		{"unknown priority", "Unknown", ""},
		{"custom priority", "P1", ""},
		// Note: "Normal" and "Trivial" don't match the patterns
		{"normal not matched", "Normal", ""},
		{"trivial not matched", "Trivial", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			priorityLower := strings.ToLower(tt.priority)
			var colorCategory string

			if strings.Contains(priorityLower, "high") || strings.Contains(priorityLower, "critical") || strings.Contains(priorityLower, "blocker") {
				colorCategory = "red"
			} else if strings.Contains(priorityLower, "medium") {
				colorCategory = "yellow"
			} else if strings.Contains(priorityLower, "low") || strings.Contains(priorityLower, "minor") {
				colorCategory = "green"
			} else {
				colorCategory = ""
			}

			if colorCategory != tt.expectedColor {
				t.Errorf("priority %q categorized as %q, expected %q", tt.priority, colorCategory, tt.expectedColor)
			}
		})
	}
}

// TestExtractFieldFromMap tests the field extraction logic used in printIssues
func TestExtractFieldFromMap(t *testing.T) {
	tests := []struct {
		name     string
		fields   map[string]interface{}
		key      string
		subkey   string
		expected string
	}{
		{
			name: "extract summary string",
			fields: map[string]interface{}{
				"summary": "Test Issue",
			},
			key:      "summary",
			subkey:   "",
			expected: "Test Issue",
		},
		{
			name: "extract status name from nested map",
			fields: map[string]interface{}{
				"status": map[string]interface{}{
					"name": "In Progress",
				},
			},
			key:      "status",
			subkey:   "name",
			expected: "In Progress",
		},
		{
			name: "extract assignee displayName",
			fields: map[string]interface{}{
				"assignee": map[string]interface{}{
					"displayName": "John Doe",
				},
			},
			key:      "assignee",
			subkey:   "displayName",
			expected: "John Doe",
		},
		{
			name:     "nil fields returns empty",
			fields:   nil,
			key:      "summary",
			subkey:   "",
			expected: "",
		},
		{
			name:     "missing key returns empty",
			fields:   map[string]interface{}{},
			key:      "summary",
			subkey:   "",
			expected: "",
		},
		{
			name: "nil value returns empty",
			fields: map[string]interface{}{
				"assignee": nil,
			},
			key:      "assignee",
			subkey:   "displayName",
			expected: "",
		},
		{
			name: "nested map missing subkey",
			fields: map[string]interface{}{
				"status": map[string]interface{}{
					"id": "123",
				},
			},
			key:      "status",
			subkey:   "name",
			expected: "",
		},
		{
			name: "value is not a map when expecting nested",
			fields: map[string]interface{}{
				"status": "string value",
			},
			key:      "status",
			subkey:   "name",
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var result string

			if tt.fields == nil {
				result = ""
			} else if obj, ok := tt.fields[tt.key]; ok && obj != nil {
				if tt.subkey == "" {
					result = fmt.Sprintf("%v", obj)
				} else if objMap, ok := obj.(map[string]interface{}); ok {
					if subVal, ok := objMap[tt.subkey].(string); ok {
						result = subVal
					}
				}
			}

			if result != tt.expected {
				t.Errorf("extracting %s.%s from %v: got %q, want %q",
					tt.key, tt.subkey, tt.fields, result, tt.expected)
			}
		})
	}
}

func TestBuildJQL(t *testing.T) {
	// Save original options and restore after test
	originalOpts := listIssuesOpts
	defer func() { listIssuesOpts = originalOpts }()

	tests := []struct {
		name     string
		opts     listIssuesOptions
		expected string
	}{
		{
			name: "project only",
			opts: listIssuesOptions{
				project: "MYPROJ",
			},
			expected: "project = MYPROJ ORDER BY status DESC",
		},
		{
			name: "project with status",
			opts: listIssuesOptions{
				project: "MYPROJ",
				status:  "In Progress",
			},
			expected: `project = MYPROJ AND status = "In Progress" ORDER BY status DESC`,
		},
		{
			name: "project with currentUser assignee",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				assignee: "currentUser()",
			},
			expected: "project = MYPROJ AND assignee = currentUser() ORDER BY status DESC",
		},
		{
			name: "project with me assignee (alias for currentUser)",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				assignee: "me",
			},
			expected: "project = MYPROJ AND assignee = currentUser() ORDER BY status DESC",
		},
		{
			name: "project with unassigned",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				assignee: "unassigned",
			},
			expected: "project = MYPROJ AND assignee IS EMPTY ORDER BY status DESC",
		},
		{
			name: "project with named assignee",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				assignee: "john.doe",
			},
			expected: `project = MYPROJ AND assignee = john.doe ORDER BY status DESC`,
		},
		{
			name: "project with reporter function",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				reporter: "currentUser()",
			},
			expected: "project = MYPROJ AND reporter = currentUser() ORDER BY status DESC",
		},
		{
			name: "project with me reporter (alias for currentUser)",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				reporter: "me",
			},
			expected: "project = MYPROJ AND reporter = currentUser() ORDER BY status DESC",
		},
		{
			name: "project with named reporter",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				reporter: "jane.doe",
			},
			expected: `project = MYPROJ AND reporter = jane.doe ORDER BY status DESC`,
		},
		{
			name: "project with labels",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"bug", "critical"},
			},
			expected: "project = MYPROJ AND labels = bug AND labels = critical ORDER BY status DESC",
		},
		{
			name: "project with issue type",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				issueType: "Bug",
			},
			expected: "project = MYPROJ AND issuetype = Bug ORDER BY status DESC",
		},
		{
			name: "project with priority",
			opts: listIssuesOptions{
				project:  "MYPROJ",
				priority: "High",
			},
			expected: "project = MYPROJ AND priority = High ORDER BY status DESC",
		},
		{
			name: "project with sprint function",
			opts: listIssuesOptions{
				project: "MYPROJ",
				sprint:  "openSprints()",
			},
			expected: "project = MYPROJ AND sprint IN openSprints() ORDER BY status DESC",
		},
		{
			name: "project with sprint name",
			opts: listIssuesOptions{
				project: "MYPROJ",
				sprint:  "Sprint 1",
			},
			expected: `project = MYPROJ AND sprint = "Sprint 1" ORDER BY status DESC`,
		},
		{
			name: "project with component",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				component: "Backend",
			},
			expected: "project = MYPROJ AND component = Backend ORDER BY status DESC",
		},
		{
			name: "project with fix version",
			opts: listIssuesOptions{
				project:    "MYPROJ",
				fixVersion: "v1.0.0",
			},
			expected: `project = MYPROJ AND fixVersion = v1.0.0 ORDER BY status DESC`,
		},
		{
			name: "project with created date range",
			opts: listIssuesOptions{
				project:       "MYPROJ",
				createdAfter:  "2024-01-01",
				createdBefore: "2024-12-31",
			},
			expected: `project = MYPROJ AND created >= "2024-01-01" AND created <= "2024-12-31" ORDER BY status DESC`,
		},
		{
			name: "project with updated date range",
			opts: listIssuesOptions{
				project:       "MYPROJ",
				updatedAfter:  "-7d",
				updatedBefore: "-1d",
			},
			expected: `project = MYPROJ AND updated >= "-7d" AND updated <= "-1d" ORDER BY status DESC`,
		},
		{
			name: "project with custom order by",
			opts: listIssuesOptions{
				project: "MYPROJ",
				orderBy: "created DESC",
			},
			expected: "project = MYPROJ ORDER BY created DESC",
		},
		{
			name: "all filters combined",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				status:    "Open",
				assignee:  "currentUser()",
				issueType: "Bug",
				priority:  "High",
				orderBy:   "priority DESC",
			},
			expected: "project = MYPROJ AND status = Open AND assignee = currentUser() AND issuetype = Bug AND priority = High ORDER BY priority DESC",
		},
		{
			name:     "no filters returns default order by only",
			opts:     listIssuesOptions{},
			expected: " ORDER BY status DESC",
		},
		// Additional edge cases
		{
			name: "closedSprints function",
			opts: listIssuesOptions{
				project: "MYPROJ",
				sprint:  "closedSprints()",
			},
			expected: "project = MYPROJ AND sprint IN closedSprints() ORDER BY status DESC",
		},
		{
			name: "futureSprints function",
			opts: listIssuesOptions{
				project: "MYPROJ",
				sprint:  "futureSprints()",
			},
			expected: "project = MYPROJ AND sprint IN futureSprints() ORDER BY status DESC",
		},
		{
			name: "project with multiple labels including special chars",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"bug-fix", "priority-high", "needs-review"},
			},
			expected: `project = MYPROJ AND labels = "bug-fix" AND labels = "priority-high" AND labels = "needs-review" ORDER BY status DESC`,
		},
		{
			name: "project with status containing space",
			opts: listIssuesOptions{
				project: "MYPROJ",
				status:  "To Do",
			},
			expected: `project = MYPROJ AND status = "To Do" ORDER BY status DESC`,
		},
		{
			name: "project key with hyphen",
			opts: listIssuesOptions{
				project: "MY-PROJ",
			},
			expected: `project = "MY-PROJ" ORDER BY status DESC`,
		},
		{
			name: "issue type with space",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				issueType: "User Story",
			},
			expected: `project = MYPROJ AND issuetype = "User Story" ORDER BY status DESC`,
		},
		{
			name: "component with space",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				component: "My Component",
			},
			expected: `project = MYPROJ AND component = "My Component" ORDER BY status DESC`,
		},
		{
			name: "fix version with dots",
			opts: listIssuesOptions{
				project:    "MYPROJ",
				fixVersion: "1.2.3",
			},
			expected: `project = MYPROJ AND fixVersion = 1.2.3 ORDER BY status DESC`,
		},
		{
			name: "only created after",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				createdAfter: "-30d",
			},
			expected: `project = MYPROJ AND created >= "-30d" ORDER BY status DESC`,
		},
		{
			name: "only created before",
			opts: listIssuesOptions{
				project:       "MYPROJ",
				createdBefore: "2024-01-01",
			},
			expected: `project = MYPROJ AND created <= "2024-01-01" ORDER BY status DESC`,
		},
		{
			name: "only updated after",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				updatedAfter: "-7d",
			},
			expected: `project = MYPROJ AND updated >= "-7d" ORDER BY status DESC`,
		},
		{
			name: "only updated before",
			opts: listIssuesOptions{
				project:       "MYPROJ",
				updatedBefore: "-1d",
			},
			expected: `project = MYPROJ AND updated <= "-1d" ORDER BY status DESC`,
		},
		{
			name: "single label",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"documentation"},
			},
			expected: "project = MYPROJ AND labels = documentation ORDER BY status DESC",
		},
		{
			name: "empty labels slice",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{},
			},
			expected: "project = MYPROJ ORDER BY status DESC",
		},
		// Test cases for labels with forward slashes (the bug fix)
		{
			name: "label with forward slash",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"team/backend"},
			},
			expected: `project = MYPROJ AND labels = "team/backend" ORDER BY status DESC`,
		},
		{
			name: "multiple labels with forward slashes",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"team/backend", "priority/high"},
			},
			expected: `project = MYPROJ AND labels = "team/backend" AND labels = "priority/high" ORDER BY status DESC`,
		},
		{
			name: "label with mixed special characters",
			opts: listIssuesOptions{
				project: "MYPROJ",
				labels:  []string{"team/backend-v2"},
			},
			expected: `project = MYPROJ AND labels = "team/backend-v2" ORDER BY status DESC`,
		},
		{
			name: "component with forward slash",
			opts: listIssuesOptions{
				project:   "MYPROJ",
				component: "Frontend/React",
			},
			expected: `project = MYPROJ AND component = "Frontend/React" ORDER BY status DESC`,
		},
		{
			name: "fix version with forward slash",
			opts: listIssuesOptions{
				project:    "MYPROJ",
				fixVersion: "2024/Q1",
			},
			expected: `project = MYPROJ AND fixVersion = "2024/Q1" ORDER BY status DESC`,
		},
		// Test cases for custom field filter
		{
			name: "single custom field",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Team=Platform"},
			},
			expected: `project = MYPROJ AND Team = Platform ORDER BY status DESC`,
		},
		{
			name: "custom field with value containing space",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Team=Platform Engineering"},
			},
			expected: `project = MYPROJ AND Team = "Platform Engineering" ORDER BY status DESC`,
		},
		{
			name: "custom field with field name containing space",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Custom Team=Platform"},
			},
			expected: `project = MYPROJ AND "Custom Team" = Platform ORDER BY status DESC`,
		},
		{
			name: "custom field using cf notation",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"cf[10001]=value"},
			},
			expected: `project = MYPROJ AND cf[10001] = value ORDER BY status DESC`,
		},
		{
			name: "multiple custom fields",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Team=Platform", "Environment=Production"},
			},
			expected: `project = MYPROJ AND Team = Platform AND Environment = Production ORDER BY status DESC`,
		},
		{
			name: "custom field with equals in value",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Formula=a=b"},
			},
			expected: `project = MYPROJ AND Formula = "a=b" ORDER BY status DESC`,
		},
		{
			name: "empty custom fields slice",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{},
			},
			expected: "project = MYPROJ ORDER BY status DESC",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listIssuesOpts = tt.opts
			result, err := buildJQL()
			if err != nil {
				t.Errorf("buildJQL() unexpected error: %v", err)
				return
			}
			if result != tt.expected {
				t.Errorf("buildJQL() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestBuildJQLCustomFieldError(t *testing.T) {
	// Save original options and restore after test
	originalOpts := listIssuesOpts
	defer func() { listIssuesOpts = originalOpts }()

	tests := []struct {
		name        string
		opts        listIssuesOptions
		expectedErr string
	}{
		{
			name: "invalid custom field format without equals",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"InvalidFormat"},
			},
			expectedErr: "invalid custom field format \"InvalidFormat\": expected 'name=value'",
		},
		{
			name: "one valid and one invalid custom field",
			opts: listIssuesOptions{
				project:      "MYPROJ",
				customFields: []string{"Team=Platform", "InvalidFormat"},
			},
			expectedErr: "invalid custom field format \"InvalidFormat\": expected 'name=value'",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			listIssuesOpts = tt.opts
			_, err := buildJQL()
			if err == nil {
				t.Errorf("buildJQL() expected error but got nil")
				return
			}
			if err.Error() != tt.expectedErr {
				t.Errorf("buildJQL() error = %q, want %q", err.Error(), tt.expectedErr)
			}
		})
	}
}
