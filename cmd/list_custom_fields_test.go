package cmd

import (
	"errors"
	"testing"
)

func TestExtractSchemaName(t *testing.T) {
	tests := []struct {
		name       string
		fullSchema string
		expected   string
	}{
		{
			name:       "standard custom field type",
			fullSchema: "com.atlassian.jira.plugin.system.customfieldtypes:textfield",
			expected:   "textfield",
		},
		{
			name:       "multiselect type",
			fullSchema: "com.atlassian.jira.plugin.system.customfieldtypes:multiselect",
			expected:   "multiselect",
		},
		{
			name:       "select type",
			fullSchema: "com.atlassian.jira.plugin.system.customfieldtypes:select",
			expected:   "select",
		},
		{
			name:       "float type",
			fullSchema: "com.atlassian.jira.plugin.system.customfieldtypes:float",
			expected:   "float",
		},
		{
			name:       "no colon - return as is",
			fullSchema: "simpletype",
			expected:   "simpletype",
		},
		{
			name:       "empty string",
			fullSchema: "",
			expected:   "",
		},
		{
			name:       "multiple colons - extract last part",
			fullSchema: "com.atlassian.jira:plugin:customtype",
			expected:   "customtype",
		},
		{
			name:       "ends with colon",
			fullSchema: "com.atlassian.jira.plugin:",
			expected:   "",
		},
		{
			name:       "third-party plugin type",
			fullSchema: "com.pyxis.greenhopper.jira:gh-sprint",
			expected:   "gh-sprint",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractSchemaName(tt.fullSchema)
			if result != tt.expected {
				t.Errorf("extractSchemaName(%q) = %q, want %q", tt.fullSchema, result, tt.expected)
			}
		})
	}
}

func TestIsPermissionError(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected bool
	}{
		{
			name:     "nil error",
			err:      nil,
			expected: false,
		},
		{
			name:     "403 status code",
			err:      errors.New("request failed with status 403"),
			expected: true,
		},
		{
			name:     "Forbidden message",
			err:      errors.New("Access Forbidden"),
			expected: true,
		},
		{
			name:     "forbidden lowercase",
			err:      errors.New("access forbidden to resource"),
			expected: true,
		},
		{
			name:     "not a permission error",
			err:      errors.New("connection timeout"),
			expected: false,
		},
		{
			name:     "404 not found",
			err:      errors.New("resource not found with status 404"),
			expected: false,
		},
		{
			name:     "generic error",
			err:      errors.New("something went wrong"),
			expected: false,
		},
		{
			name:     "403 in context",
			err:      errors.New("API returned error: 403 Forbidden"),
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isPermissionError(tt.err)
			if result != tt.expected {
				t.Errorf("isPermissionError(%v) = %v, want %v", tt.err, result, tt.expected)
			}
		})
	}
}
