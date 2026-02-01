package cmd

import (
	"reflect"
	"testing"
)

func TestFindCustomFieldByName(t *testing.T) {
	fieldMap := map[string]customFieldInfo{
		"customfield_10001": {id: "customfield_10001", name: "Team", schemaType: "string"},
		"customfield_10002": {id: "customfield_10002", name: "Story Points", schemaType: "number"},
		"customfield_10003": {id: "customfield_10003", name: "Environment", schemaType: "option"},
	}

	tests := []struct {
		name      string
		fieldName string
		expected  customFieldInfo
		found     bool
	}{
		{
			name:      "exact match",
			fieldName: "Team",
			expected:  customFieldInfo{id: "customfield_10001", name: "Team", schemaType: "string"},
			found:     true,
		},
		{
			name:      "case insensitive match",
			fieldName: "team",
			expected:  customFieldInfo{id: "customfield_10001", name: "Team", schemaType: "string"},
			found:     true,
		},
		{
			name:      "uppercase match",
			fieldName: "STORY POINTS",
			expected:  customFieldInfo{id: "customfield_10002", name: "Story Points", schemaType: "number"},
			found:     true,
		},
		{
			name:      "not found",
			fieldName: "Nonexistent Field",
			expected:  customFieldInfo{},
			found:     false,
		},
		{
			name:      "empty field name",
			fieldName: "",
			expected:  customFieldInfo{},
			found:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, found := findCustomFieldByName(fieldMap, tt.fieldName)
			if found != tt.found {
				t.Errorf("findCustomFieldByName() found = %v, want %v", found, tt.found)
			}
			if found && result != tt.expected {
				t.Errorf("findCustomFieldByName() result = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractFieldType(t *testing.T) {
	tests := []struct {
		name       string
		customType string
		expected   string
	}{
		{
			name:       "select field type",
			customType: "com.atlassian.jira.plugin.system.customfieldtypes:select",
			expected:   "select",
		},
		{
			name:       "multiselect field type",
			customType: "com.atlassian.jira.plugin.system.customfieldtypes:multiselect",
			expected:   "multiselect",
		},
		{
			name:       "textfield type",
			customType: "com.atlassian.jira.plugin.system.customfieldtypes:textfield",
			expected:   "textfield",
		},
		{
			name:       "no colon",
			customType: "simpletype",
			expected:   "simpletype",
		},
		{
			name:       "empty string",
			customType: "",
			expected:   "",
		},
		{
			name:       "ends with colon",
			customType: "prefix:",
			expected:   "",
		},
		{
			name:       "multiple colons",
			customType: "com.atlassian.jira:plugin.system:float",
			expected:   "float",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractFieldType(tt.customType)
			if result != tt.expected {
				t.Errorf("extractFieldType(%q) = %q, want %q", tt.customType, result, tt.expected)
			}
		})
	}
}

func TestParseCommaSeparatedValues(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple values",
			input:    "a,b,c",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "values with spaces",
			input:    "a, b, c",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "values with extra spaces",
			input:    "  a  ,  b  ,  c  ",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "single value",
			input:    "single",
			expected: []string{"single"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only commas",
			input:    ",,,",
			expected: []string{},
		},
		{
			name:     "empty values between commas",
			input:    "a,,b",
			expected: []string{"a", "b"},
		},
		{
			name:     "values with internal spaces",
			input:    "Hello World, Goodbye World",
			expected: []string{"Hello World", "Goodbye World"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parseCommaSeparatedValues(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parseCommaSeparatedValues(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertCustomFieldValue(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		field       customFieldInfo
		expected    interface{}
		expectError bool
	}{
		{
			name:     "empty value clears field",
			value:    "",
			field:    customFieldInfo{schemaType: "string"},
			expected: nil,
		},
		{
			name:     "string field",
			value:    "hello",
			field:    customFieldInfo{schemaType: "string"},
			expected: "hello",
		},
		{
			name:     "number field with integer",
			value:    "42",
			field:    customFieldInfo{schemaType: "number"},
			expected: float64(42),
		},
		{
			name:     "number field with float",
			value:    "3.14",
			field:    customFieldInfo{schemaType: "number"},
			expected: float64(3.14),
		},
		{
			name:        "number field with invalid value",
			value:       "not a number",
			field:       customFieldInfo{schemaType: "number"},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "option field",
			value: "Option A",
			field: customFieldInfo{schemaType: "option"},
			expected: map[string]interface{}{
				"value": "Option A",
			},
		},
		{
			name:  "array of options",
			value: "Opt1, Opt2, Opt3",
			field: customFieldInfo{schemaType: "array", itemType: "option"},
			expected: []map[string]interface{}{
				{"value": "Opt1"},
				{"value": "Opt2"},
				{"value": "Opt3"},
			},
		},
		{
			name:     "array of strings",
			value:    "tag1, tag2",
			field:    customFieldInfo{schemaType: "array", itemType: "string"},
			expected: []string{"tag1", "tag2"},
		},
		{
			name:  "array of users",
			value: "user1, user2",
			field: customFieldInfo{schemaType: "array", itemType: "user"},
			expected: []map[string]interface{}{
				{"id": "user1"},
				{"id": "user2"},
			},
		},
		{
			name:  "single user",
			value: "user123",
			field: customFieldInfo{schemaType: "user"},
			expected: map[string]interface{}{
				"id": "user123",
			},
		},
		{
			name:     "date field",
			value:    "2024-01-15",
			field:    customFieldInfo{schemaType: "date"},
			expected: "2024-01-15",
		},
		{
			name:     "datetime field",
			value:    "2024-01-15T10:30:00Z",
			field:    customFieldInfo{schemaType: "datetime"},
			expected: "2024-01-15T10:30:00Z",
		},
		{
			name:  "select custom type",
			value: "Selected",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:select"},
			expected: map[string]interface{}{
				"value": "Selected",
			},
		},
		{
			name:  "radio buttons custom type",
			value: "Option",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:radiobuttons"},
			expected: map[string]interface{}{
				"value": "Option",
			},
		},
		{
			name:  "multiselect custom type",
			value: "A, B",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:multiselect"},
			expected: []map[string]interface{}{
				{"value": "A"},
				{"value": "B"},
			},
		},
		{
			name:  "multi checkboxes custom type",
			value: "X, Y",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:multicheckboxes"},
			expected: []map[string]interface{}{
				{"value": "X"},
				{"value": "Y"},
			},
		},
		{
			name:  "user picker custom type",
			value: "user456",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:userpicker"},
			expected: map[string]interface{}{
				"id": "user456",
			},
		},
		{
			name:  "multi user picker custom type",
			value: "u1, u2",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:multiuserpicker"},
			expected: []map[string]interface{}{
				{"id": "u1"},
				{"id": "u2"},
			},
		},
		{
			name:  "cascading select - parent only",
			value: "Parent",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect"},
			expected: map[string]interface{}{
				"value": "Parent",
			},
		},
		{
			name:  "cascading select - parent and child",
			value: "Parent - Child",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect"},
			expected: map[string]interface{}{
				"value": "Parent",
				"child": map[string]interface{}{
					"value": "Child",
				},
			},
		},
		{
			name:     "float custom type",
			value:    "99.5",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:float"},
			expected: float64(99.5),
		},
		{
			name:        "float custom type with invalid value",
			value:       "invalid",
			field:       customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:float"},
			expected:    nil,
			expectError: true,
		},
		{
			name:     "datepicker custom type",
			value:    "2024-06-15",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:datepicker"},
			expected: "2024-06-15",
		},
		{
			name:     "url custom type",
			value:    "https://example.com",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:url"},
			expected: "https://example.com",
		},
		{
			name:     "textarea custom type",
			value:    "Long text content",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:textarea"},
			expected: "Long text content",
		},
		{
			name:     "textfield custom type",
			value:    "Short text",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:textfield"},
			expected: "Short text",
		},
		{
			name:     "datetime custom type",
			value:    "2024-01-15T10:30:00.000+0000",
			field:    customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:datetime"},
			expected: "2024-01-15T10:30:00.000+0000",
		},
		{
			name:     "unknown schema type defaults to string",
			value:    "default value",
			field:    customFieldInfo{schemaType: "unknown"},
			expected: "default value",
		},
		{
			name:     "array with unknown item type",
			value:    "a, b, c",
			field:    customFieldInfo{schemaType: "array", itemType: "unknown"},
			expected: []string{"a", "b", "c"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := convertCustomFieldValue(tt.value, tt.field)
			if tt.expectError {
				if err == nil {
					t.Errorf("convertCustomFieldValue() expected error but got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("convertCustomFieldValue() unexpected error: %v", err)
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("convertCustomFieldValue(%q, %+v) = %v, want %v", tt.value, tt.field, result, tt.expected)
			}
		})
	}
}
