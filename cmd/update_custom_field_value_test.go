package cmd

import (
	"testing"
)

func TestListMatchingCustomFields(t *testing.T) {
	fieldInfoMap := map[string]customFieldInfo{
		"customfield_10001": {id: "customfield_10001", name: "Team", schemaType: "string"},
		"customfield_10002": {id: "customfield_10002", name: "Story Points", schemaType: "number"},
		"customfield_10003": {id: "customfield_10003", name: "Team Lead", schemaType: "user"},
		"customfield_10004": {id: "customfield_10004", name: "Environment", schemaType: "option"},
	}

	tests := []struct {
		name          string
		query         string
		expectedCount int
		expectedNames []string
	}{
		{
			name:          "exact match",
			query:         "Team",
			expectedCount: 2, // "Team" and "Team Lead"
			expectedNames: []string{"Team", "Team Lead"},
		},
		{
			name:          "partial match",
			query:         "point",
			expectedCount: 1,
			expectedNames: []string{"Story Points"},
		},
		{
			name:          "case insensitive",
			query:         "ENVIRONMENT",
			expectedCount: 1,
			expectedNames: []string{"Environment"},
		},
		{
			name:          "no match",
			query:         "nonexistent",
			expectedCount: 0,
			expectedNames: []string{},
		},
		{
			name:          "empty query matches all",
			query:         "",
			expectedCount: 4,
			expectedNames: []string{"Team", "Story Points", "Team Lead", "Environment"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			matches := listMatchingCustomFields(fieldInfoMap, tt.query)
			if len(matches) != tt.expectedCount {
				t.Errorf("listMatchingCustomFields() returned %d matches, want %d", len(matches), tt.expectedCount)
			}

			// Check that all expected names are present
			matchNames := make(map[string]bool)
			for _, m := range matches {
				matchNames[m.name] = true
			}
			for _, expectedName := range tt.expectedNames {
				if !matchNames[expectedName] {
					t.Errorf("listMatchingCustomFields() missing expected match %q", expectedName)
				}
			}
		})
	}
}

func TestUpdateCustomFieldValueOptionsValidation(t *testing.T) {
	// Test that required fields are properly marked
	cmd := updateCustomFieldValueCmd

	// Check that name, old, new, and enable flags exist
	nameFlag := cmd.Flags().Lookup("name")
	if nameFlag == nil {
		t.Error("name flag not found")
	}

	oldFlag := cmd.Flags().Lookup("old")
	if oldFlag == nil {
		t.Error("old flag not found")
	}

	newFlag := cmd.Flags().Lookup("new")
	if newFlag == nil {
		t.Error("new flag not found")
	}

	enableFlag := cmd.Flags().Lookup("enable")
	if enableFlag == nil {
		t.Error("enable flag not found")
	}
}

func TestUpdateCustomFieldValueShortFlags(t *testing.T) {
	cmd := updateCustomFieldValueCmd

	tests := []struct {
		shorthand string
		longName  string
	}{
		{"n", "name"},
		{"o", "old"},
		{"w", "new"},
		{"e", "enable"},
	}

	for _, tt := range tests {
		t.Run(tt.longName, func(t *testing.T) {
			flag := cmd.Flags().Lookup(tt.longName)
			if flag == nil {
				t.Errorf("flag %q not found", tt.longName)
				return
			}
			if flag.Shorthand != tt.shorthand {
				t.Errorf("flag %q has shorthand %q, want %q", tt.longName, flag.Shorthand, tt.shorthand)
			}
		})
	}
}
