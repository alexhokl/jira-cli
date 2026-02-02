package cmd

import "testing"

// TestQuoteJQLValueForBulkUpdateDeleteLabel tests that labels with special characters
// are properly quoted for the bulk-update delete-label command
func TestQuoteJQLValueForBulkUpdateDeleteLabel(t *testing.T) {
	tests := []struct {
		name        string
		labelName   string
		project     string
		expectedJQL string
	}{
		{
			name:        "simple label",
			labelName:   "bug",
			project:     "",
			expectedJQL: "labels = bug",
		},
		{
			name:        "label with forward slash",
			labelName:   "team/backend",
			project:     "",
			expectedJQL: `labels = "team/backend"`,
		},
		{
			name:        "label with hyphen",
			labelName:   "high-priority",
			project:     "",
			expectedJQL: `labels = "high-priority"`,
		},
		{
			name:        "label with forward slash and project",
			labelName:   "team/backend",
			project:     "MYPROJ",
			expectedJQL: `project = MYPROJ AND labels = "team/backend"`,
		},
		{
			name:        "label with special chars and project with hyphen",
			labelName:   "priority/high",
			project:     "MY-PROJ",
			expectedJQL: `project = "MY-PROJ" AND labels = "priority/high"`,
		},
		{
			name:        "label with multiple special characters",
			labelName:   "team/backend-v2",
			project:     "",
			expectedJQL: `labels = "team/backend-v2"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build JQL the same way as runBulkUpdateDeleteLabel
			jql := "labels = " + quoteJQLValue(tt.labelName)
			if tt.project != "" {
				jql = "project = " + quoteJQLValue(tt.project) + " AND " + jql
			}

			if jql != tt.expectedJQL {
				t.Errorf("JQL = %q, want %q", jql, tt.expectedJQL)
			}
		})
	}
}

// TestQuoteJQLValueForBulkUpdateRenameLabel tests that labels with special characters
// are properly quoted for the bulk-update rename-label command
func TestQuoteJQLValueForBulkUpdateRenameLabel(t *testing.T) {
	tests := []struct {
		name        string
		oldLabel    string
		project     string
		expectedJQL string
	}{
		{
			name:        "simple label",
			oldLabel:    "bug",
			project:     "",
			expectedJQL: "labels = bug",
		},
		{
			name:        "label with forward slash",
			oldLabel:    "team/frontend",
			project:     "",
			expectedJQL: `labels = "team/frontend"`,
		},
		{
			name:        "label with colon",
			oldLabel:    "env:production",
			project:     "",
			expectedJQL: `labels = "env:production"`,
		},
		{
			name:        "label with forward slash and project",
			oldLabel:    "team/frontend",
			project:     "MYPROJ",
			expectedJQL: `project = MYPROJ AND labels = "team/frontend"`,
		},
		{
			name:        "label with asterisk",
			oldLabel:    "priority*",
			project:     "",
			expectedJQL: `labels = "priority*"`,
		},
		{
			name:        "label with plus sign",
			oldLabel:    "c++",
			project:     "",
			expectedJQL: `labels = "c++"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build JQL the same way as runBulkUpdateRenameLabel
			jql := "labels = " + quoteJQLValue(tt.oldLabel)
			if tt.project != "" {
				jql = "project = " + quoteJQLValue(tt.project) + " AND " + jql
			}

			if jql != tt.expectedJQL {
				t.Errorf("JQL = %q, want %q", jql, tt.expectedJQL)
			}
		})
	}
}
