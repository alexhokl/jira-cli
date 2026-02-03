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
			if found && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("findCustomFieldByName() result = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestFindCustomFieldByName_EmptyMap(t *testing.T) {
	emptyMap := map[string]customFieldInfo{}
	result, found := findCustomFieldByName(emptyMap, "Team")
	if found {
		t.Error("findCustomFieldByName() should return false for empty map")
	}
	if result.id != "" {
		t.Errorf("findCustomFieldByName() should return empty customFieldInfo, got %v", result)
	}
}

func TestFindCustomFieldByName_SpecialCharacters(t *testing.T) {
	fieldMap := map[string]customFieldInfo{
		"customfield_10001": {id: "customfield_10001", name: "Field (with parens)", schemaType: "string"},
		"customfield_10002": {id: "customfield_10002", name: "Field-with-dashes", schemaType: "string"},
		"customfield_10003": {id: "customfield_10003", name: "Field_with_underscores", schemaType: "string"},
		"customfield_10004": {id: "customfield_10004", name: "Field.with.dots", schemaType: "string"},
	}

	tests := []struct {
		name      string
		fieldName string
		found     bool
	}{
		{"parentheses", "Field (with parens)", true},
		{"parentheses case insensitive", "field (WITH PARENS)", true},
		{"dashes", "Field-with-dashes", true},
		{"underscores", "Field_with_underscores", true},
		{"dots", "Field.with.dots", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, found := findCustomFieldByName(fieldMap, tt.fieldName)
			if found != tt.found {
				t.Errorf("findCustomFieldByName(%q) found = %v, want %v", tt.fieldName, found, tt.found)
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
		{
			name:       "gh-sprint type",
			customType: "com.pyxis.greenhopper.jira:gh-sprint",
			expected:   "gh-sprint",
		},
		{
			name:       "cascadingselect type",
			customType: "com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect",
			expected:   "cascadingselect",
		},
		{
			name:       "only colon",
			customType: ":",
			expected:   "",
		},
		{
			name:       "colon at start",
			customType: ":fieldtype",
			expected:   "fieldtype",
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
		{
			name:     "single value with spaces",
			input:    "  single value  ",
			expected: []string{"single value"},
		},
		{
			name:     "only whitespace",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "whitespace between commas",
			input:    " , , ",
			expected: []string{},
		},
		{
			name:     "tabs and newlines",
			input:    "a,\tb,\nc",
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "numeric values",
			input:    "1, 2, 3",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "mixed alphanumeric",
			input:    "Sprint 1, Sprint 2, Sprint 3",
			expected: []string{"Sprint 1", "Sprint 2", "Sprint 3"},
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

// Note: TestExtractProjectKeyFromIssueKey is defined in helper_test.go

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
			name:     "number field with negative value",
			value:    "-42",
			field:    customFieldInfo{schemaType: "number"},
			expected: float64(-42),
		},
		{
			name:     "number field with zero",
			value:    "0",
			field:    customFieldInfo{schemaType: "number"},
			expected: float64(0),
		},
		{
			name:     "number field with scientific notation",
			value:    "1.5e2",
			field:    customFieldInfo{schemaType: "number"},
			expected: float64(150),
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
			name:  "option field with allowed values",
			value: "Option A",
			field: customFieldInfo{
				schemaType: "option",
				allowedValues: []interface{}{
					map[string]interface{}{"id": "10001", "value": "Option A"},
					map[string]interface{}{"id": "10002", "value": "Option B"},
				},
			},
			expected: map[string]interface{}{
				"id": "10001",
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
			name:  "array of options with allowed values",
			value: "Option A, Option B",
			field: customFieldInfo{
				schemaType: "array",
				itemType:   "option",
				allowedValues: []interface{}{
					map[string]interface{}{"id": "10001", "value": "Option A"},
					map[string]interface{}{"id": "10002", "value": "Option B"},
				},
			},
			expected: []map[string]interface{}{
				{"id": "10001"},
				{"id": "10002"},
			},
		},
		{
			name:     "array of strings",
			value:    "tag1, tag2",
			field:    customFieldInfo{schemaType: "array", itemType: "string"},
			expected: []string{"tag1", "tag2"},
		},
		{
			name:     "array of strings single value",
			value:    "tag1",
			field:    customFieldInfo{schemaType: "array", itemType: "string"},
			expected: []string{"tag1"},
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
			name:  "cascading select - parent and child with spaces",
			value: "  Parent Value  -  Child Value  ",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect"},
			expected: map[string]interface{}{
				"value": "Parent Value",
				"child": map[string]interface{}{
					"value": "Child Value",
				},
			},
		},
		{
			name:  "cascading select - value with dash but no child separator",
			value: "Parent-Value",
			field: customFieldInfo{schemaType: "any", custom: "com.atlassian.jira.plugin.system.customfieldtypes:cascadingselect"},
			expected: map[string]interface{}{
				"value": "Parent-Value",
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
		{
			name:     "sprint field with single ID",
			value:    "123",
			field:    customFieldInfo{schemaType: "array", itemType: "json", custom: "com.pyxis.greenhopper.jira:gh-sprint"},
			expected: []int{123},
		},
		{
			name:     "sprint field with multiple IDs",
			value:    "123, 456",
			field:    customFieldInfo{schemaType: "array", itemType: "json", custom: "com.pyxis.greenhopper.jira:gh-sprint"},
			expected: []int{123, 456},
		},
		{
			name:        "sprint field with invalid ID (name instead of ID)",
			value:       "Sprint 97",
			field:       customFieldInfo{schemaType: "array", itemType: "json", custom: "com.pyxis.greenhopper.jira:gh-sprint"},
			expected:    nil,
			expectError: true,
		},
		{
			name:     "string with special characters",
			value:    "Test <script>alert('xss')</script>",
			field:    customFieldInfo{schemaType: "string"},
			expected: "Test <script>alert('xss')</script>",
		},
		{
			name:     "string with unicode",
			value:    "Test 你好 🎉",
			field:    customFieldInfo{schemaType: "string"},
			expected: "Test 你好 🎉",
		},
		{
			name:     "string with newlines",
			value:    "Line 1\nLine 2\nLine 3",
			field:    customFieldInfo{schemaType: "string"},
			expected: "Line 1\nLine 2\nLine 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Pass empty project key - sprint name lookup tests would need mock server
			result, err := convertCustomFieldValue(tt.value, tt.field, "")
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

func TestLookupOptionId(t *testing.T) {
	allowedValues := []interface{}{
		map[string]interface{}{"id": "10001", "value": "Option A"},
		map[string]interface{}{"id": "10002", "value": "Option B"},
		map[string]interface{}{"id": float64(10003), "value": "Option C"},
		map[string]interface{}{"id": "10004", "name": "Named Option"},
	}

	tests := []struct {
		name       string
		valueName  string
		expectedId string
	}{
		{
			name:       "exact match by value",
			valueName:  "Option A",
			expectedId: "10001",
		},
		{
			name:       "case insensitive match",
			valueName:  "option b",
			expectedId: "10002",
		},
		{
			name:       "numeric id",
			valueName:  "Option C",
			expectedId: "10003",
		},
		{
			name:       "match by name field",
			valueName:  "Named Option",
			expectedId: "10004",
		},
		{
			name:       "not found",
			valueName:  "Nonexistent",
			expectedId: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lookupOptionId(allowedValues, tt.valueName)
			if result != tt.expectedId {
				t.Errorf("lookupOptionId() = %q, want %q", result, tt.expectedId)
			}
		})
	}
}

func TestLookupOptionId_EdgeCases(t *testing.T) {
	tests := []struct {
		name          string
		allowedValues []interface{}
		valueName     string
		expectedId    string
	}{
		{
			name:          "empty allowed values",
			allowedValues: []interface{}{},
			valueName:     "Option A",
			expectedId:    "",
		},
		{
			name:          "nil allowed values",
			allowedValues: nil,
			valueName:     "Option A",
			expectedId:    "",
		},
		{
			name: "allowed value without id",
			allowedValues: []interface{}{
				map[string]interface{}{"value": "Option A"},
			},
			valueName:  "Option A",
			expectedId: "",
		},
		{
			name: "allowed value with non-string/float id",
			allowedValues: []interface{}{
				map[string]interface{}{"id": 123, "value": "Option A"}, // int, not float64
			},
			valueName:  "Option A",
			expectedId: "",
		},
		{
			name: "non-map allowed value",
			allowedValues: []interface{}{
				"not a map",
			},
			valueName:  "not a map",
			expectedId: "",
		},
		{
			name: "match by name with numeric id",
			allowedValues: []interface{}{
				map[string]interface{}{"id": float64(999), "name": "Named Option"},
			},
			valueName:  "Named Option",
			expectedId: "999",
		},
		{
			name: "empty value name",
			allowedValues: []interface{}{
				map[string]interface{}{"id": "10001", "value": ""},
			},
			valueName:  "",
			expectedId: "10001",
		},
		{
			name: "value with leading/trailing spaces in allowed values",
			allowedValues: []interface{}{
				map[string]interface{}{"id": "10001", "value": "Option A"},
			},
			valueName:  "Option A", // exact match required
			expectedId: "10001",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := lookupOptionId(tt.allowedValues, tt.valueName)
			if result != tt.expectedId {
				t.Errorf("lookupOptionId() = %q, want %q", result, tt.expectedId)
			}
		})
	}
}

func TestFormatOptionValue(t *testing.T) {
	allowedValues := []interface{}{
		map[string]interface{}{"id": "10001", "value": "Option A"},
		map[string]interface{}{"id": "10002", "value": "Option B"},
	}

	tests := []struct {
		name          string
		valueName     string
		allowedValues []interface{}
		expected      map[string]interface{}
	}{
		{
			name:          "with allowed values - found",
			valueName:     "Option A",
			allowedValues: allowedValues,
			expected:      map[string]interface{}{"id": "10001"},
		},
		{
			name:          "with allowed values - not found (fallback to value)",
			valueName:     "Unknown",
			allowedValues: allowedValues,
			expected:      map[string]interface{}{"value": "Unknown"},
		},
		{
			name:          "without allowed values (fallback to value)",
			valueName:     "Option A",
			allowedValues: nil,
			expected:      map[string]interface{}{"value": "Option A"},
		},
		{
			name:          "empty allowed values (fallback to value)",
			valueName:     "Option A",
			allowedValues: []interface{}{},
			expected:      map[string]interface{}{"value": "Option A"},
		},
		{
			name:          "empty value name with no allowed values",
			valueName:     "",
			allowedValues: nil,
			expected:      map[string]interface{}{"value": ""},
		},
		{
			name:          "case insensitive match in allowed values",
			valueName:     "OPTION A",
			allowedValues: allowedValues,
			expected:      map[string]interface{}{"id": "10001"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatOptionValue(tt.valueName, tt.allowedValues)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("formatOptionValue() = %v, want %v", result, tt.expected)
			}
		})
	}
}

// Note: TestNewADFDocument is defined in create_comment_test.go

func TestCustomFieldInfo_AllFields(t *testing.T) {
	// Test that customFieldInfo struct works correctly with all fields populated
	info := customFieldInfo{
		id:         "customfield_10001",
		name:       "Test Field",
		schemaType: "option",
		itemType:   "string",
		custom:     "com.atlassian.jira.plugin.system.customfieldtypes:select",
		allowedValues: []interface{}{
			map[string]interface{}{"id": "1", "value": "A"},
		},
	}

	if info.id != "customfield_10001" {
		t.Errorf("customFieldInfo.id = %q, want %q", info.id, "customfield_10001")
	}
	if info.name != "Test Field" {
		t.Errorf("customFieldInfo.name = %q, want %q", info.name, "Test Field")
	}
	if info.schemaType != "option" {
		t.Errorf("customFieldInfo.schemaType = %q, want %q", info.schemaType, "option")
	}
	if info.itemType != "string" {
		t.Errorf("customFieldInfo.itemType = %q, want %q", info.itemType, "string")
	}
	if info.custom != "com.atlassian.jira.plugin.system.customfieldtypes:select" {
		t.Errorf("customFieldInfo.custom = %q, want %q", info.custom, "com.atlassian.jira.plugin.system.customfieldtypes:select")
	}
	if len(info.allowedValues) != 1 {
		t.Errorf("len(customFieldInfo.allowedValues) = %d, want 1", len(info.allowedValues))
	}
}

func TestConvertCustomFieldValue_NumberEdgeCases(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expected    interface{}
		expectError bool
	}{
		{
			name:     "very large number",
			value:    "999999999999999",
			expected: float64(999999999999999),
		},
		{
			name:     "very small decimal",
			value:    "0.000001",
			expected: float64(0.000001),
		},
		{
			name:     "negative decimal",
			value:    "-123.456",
			expected: float64(-123.456),
		},
		{
			name:        "number with letters",
			value:       "123abc",
			expectError: true,
		},
		{
			name:        "number with spaces",
			value:       "12 34",
			expectError: true,
		},
		{
			name:     "positive sign",
			value:    "+42",
			expected: float64(42),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := customFieldInfo{schemaType: "number"}
			result, err := convertCustomFieldValue(tt.value, field, "")

			if tt.expectError {
				if err == nil {
					t.Errorf("convertCustomFieldValue(%q) expected error but got nil", tt.value)
				}
				return
			}

			if err != nil {
				t.Errorf("convertCustomFieldValue(%q) unexpected error: %v", tt.value, err)
				return
			}

			if result != tt.expected {
				t.Errorf("convertCustomFieldValue(%q) = %v, want %v", tt.value, result, tt.expected)
			}
		})
	}
}
