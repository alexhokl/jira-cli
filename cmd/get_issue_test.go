package cmd

import (
	"reflect"
	"testing"
)

func TestExtractLabels(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []string
	}{
		{
			name:     "valid labels array",
			input:    []any{"bug", "critical", "frontend"},
			expected: []string{"bug", "critical", "frontend"},
		},
		{
			name:     "empty array",
			input:    []any{},
			expected: nil,
		},
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "not an array",
			input:    "single label",
			expected: nil,
		},
		{
			name:     "mixed types in array",
			input:    []any{"label1", 123, "label2"},
			expected: []string{"label1", "label2"},
		},
		{
			name:     "single label",
			input:    []any{"only-one"},
			expected: []string{"only-one"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractLabels(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("extractLabels(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractUserDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "user with displayName",
			input: map[string]any{
				"displayName":  "John Doe",
				"emailAddress": "john@example.com",
			},
			expected: "John Doe",
		},
		{
			name: "user with only emailAddress",
			input: map[string]any{
				"emailAddress": "john@example.com",
			},
			expected: "john@example.com",
		},
		{
			name:     "nil input",
			input:    nil,
			expected: "",
		},
		{
			name:     "not a map",
			input:    "invalid",
			expected: "",
		},
		{
			name:     "empty map",
			input:    map[string]any{},
			expected: "",
		},
		{
			name: "displayName is not a string",
			input: map[string]any{
				"displayName": 123,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractUserDisplayName(tt.input)
			if result != tt.expected {
				t.Errorf("extractUserDisplayName(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractIssueLinks(t *testing.T) {
	tests := []struct {
		name       string
		input      any
		currentKey string
		expected   []issueLink
	}{
		{
			name: "inward link",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward":  "is blocked by",
						"outward": "blocks",
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-456",
						"fields": map[string]any{
							"summary": "Blocking issue",
						},
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "is blocked by", key: "PROJ-456", summary: "Blocking issue"},
			},
		},
		{
			name: "outward link",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward":  "is blocked by",
						"outward": "blocks",
					},
					"outwardIssue": map[string]any{
						"key": "PROJ-789",
						"fields": map[string]any{
							"summary": "Blocked issue",
						},
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "blocks", key: "PROJ-789", summary: "Blocked issue"},
			},
		},
		{
			name:       "nil input",
			input:      nil,
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name:       "empty array",
			input:      []any{},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name:       "not an array",
			input:      "invalid",
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "link to self is filtered out",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward":  "relates to",
						"outward": "relates to",
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-123",
						"fields": map[string]any{
							"summary": "Self reference",
						},
					},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "multiple links",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward":  "is blocked by",
						"outward": "blocks",
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-1",
						"fields": map[string]any{
							"summary": "First",
						},
					},
				},
				map[string]any{
					"type": map[string]any{
						"inward":  "duplicates",
						"outward": "is duplicated by",
					},
					"outwardIssue": map[string]any{
						"key": "PROJ-2",
						"fields": map[string]any{
							"summary": "Second",
						},
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "is blocked by", key: "PROJ-1", summary: "First"},
				{linkType: "is duplicated by", key: "PROJ-2", summary: "Second"},
			},
		},
		// Additional tests for uncovered branches
		{
			name: "link item is not a map",
			input: []any{
				"not a map",
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "link missing type field",
			input: []any{
				map[string]any{
					"inwardIssue": map[string]any{
						"key": "PROJ-456",
					},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "link type is not a map",
			input: []any{
				map[string]any{
					"type":        "invalid type",
					"inwardIssue": map[string]any{"key": "PROJ-456"},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "inward link without key",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward": "is blocked by",
					},
					"inwardIssue": map[string]any{
						"fields": map[string]any{"summary": "No key"},
					},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "outward link without key",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"outward": "blocks",
					},
					"outwardIssue": map[string]any{
						"fields": map[string]any{"summary": "No key"},
					},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "inward link with key as non-string",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward": "is blocked by",
					},
					"inwardIssue": map[string]any{
						"key": 12345,
					},
				},
			},
			currentKey: "PROJ-123",
			expected:   nil,
		},
		{
			name: "inward link without fields",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward": "is blocked by",
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-456",
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "is blocked by", key: "PROJ-456", summary: ""},
			},
		},
		{
			name: "inward link with fields but no summary",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward": "is blocked by",
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-456",
						"fields": map[string]any{
							"status": "Open",
						},
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "is blocked by", key: "PROJ-456", summary: ""},
			},
		},
		{
			name: "outward link without fields",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"outward": "blocks",
					},
					"outwardIssue": map[string]any{
						"key": "PROJ-789",
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "blocks", key: "PROJ-789", summary: ""},
			},
		},
		{
			name: "inward link type not a string",
			input: []any{
				map[string]any{
					"type": map[string]any{
						"inward": 123,
					},
					"inwardIssue": map[string]any{
						"key": "PROJ-456",
					},
				},
			},
			currentKey: "PROJ-123",
			expected: []issueLink{
				{linkType: "", key: "PROJ-456", summary: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractIssueLinks(tt.input, tt.currentKey)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("extractIssueLinks(%v, %q) = %v, want %v", tt.input, tt.currentKey, result, tt.expected)
			}
		})
	}
}

func TestExtractSubtasks(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected []subtaskInfo
	}{
		{
			name: "single subtask with status",
			input: []any{
				map[string]any{
					"key": "PROJ-101",
					"fields": map[string]any{
						"summary": "Subtask one",
						"status": map[string]any{
							"name": "In Progress",
						},
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-101", summary: "Subtask one", status: "In Progress"},
			},
		},
		{
			name: "subtask without status",
			input: []any{
				map[string]any{
					"key": "PROJ-102",
					"fields": map[string]any{
						"summary": "Subtask two",
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-102", summary: "Subtask two", status: ""},
			},
		},
		{
			name:     "nil input",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty array",
			input:    []any{},
			expected: nil,
		},
		{
			name: "multiple subtasks",
			input: []any{
				map[string]any{
					"key": "PROJ-1",
					"fields": map[string]any{
						"summary": "First",
						"status":  map[string]any{"name": "Done"},
					},
				},
				map[string]any{
					"key": "PROJ-2",
					"fields": map[string]any{
						"summary": "Second",
						"status":  map[string]any{"name": "To Do"},
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-1", summary: "First", status: "Done"},
				{key: "PROJ-2", summary: "Second", status: "To Do"},
			},
		},
		{
			name: "subtask without key is skipped",
			input: []any{
				map[string]any{
					"fields": map[string]any{
						"summary": "No key subtask",
					},
				},
			},
			expected: nil,
		},
		// Additional tests for uncovered branches
		{
			name:     "not an array",
			input:    "invalid",
			expected: nil,
		},
		{
			name: "subtask item is not a map",
			input: []any{
				"not a map",
			},
			expected: nil,
		},
		{
			name: "subtask with key as non-string",
			input: []any{
				map[string]any{
					"key": 12345,
					"fields": map[string]any{
						"summary": "Has numeric key",
					},
				},
			},
			expected: nil,
		},
		{
			name: "subtask without fields",
			input: []any{
				map[string]any{
					"key": "PROJ-103",
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-103", summary: "", status: ""},
			},
		},
		{
			name: "subtask with fields not a map",
			input: []any{
				map[string]any{
					"key":    "PROJ-104",
					"fields": "not a map",
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-104", summary: "", status: ""},
			},
		},
		{
			name: "subtask with summary not a string",
			input: []any{
				map[string]any{
					"key": "PROJ-105",
					"fields": map[string]any{
						"summary": 12345,
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-105", summary: "", status: ""},
			},
		},
		{
			name: "subtask with status not a map",
			input: []any{
				map[string]any{
					"key": "PROJ-106",
					"fields": map[string]any{
						"summary": "Valid summary",
						"status":  "not a map",
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-106", summary: "Valid summary", status: ""},
			},
		},
		{
			name: "subtask with status name not a string",
			input: []any{
				map[string]any{
					"key": "PROJ-107",
					"fields": map[string]any{
						"summary": "Valid summary",
						"status": map[string]any{
							"name": 123,
						},
					},
				},
			},
			expected: []subtaskInfo{
				{key: "PROJ-107", summary: "Valid summary", status: ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractSubtasks(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("extractSubtasks(%v) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFormatCustomFieldValue(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "nil value",
			input:    nil,
			expected: "",
		},
		{
			name:     "string value",
			input:    "hello",
			expected: "hello",
		},
		{
			name:     "integer float",
			input:    float64(42),
			expected: "42",
		},
		{
			name:     "decimal float",
			input:    float64(3.14),
			expected: "3.14",
		},
		{
			name:     "boolean true",
			input:    true,
			expected: "true",
		},
		{
			name:     "boolean false",
			input:    false,
			expected: "false",
		},
		{
			name: "map with displayName",
			input: map[string]any{
				"displayName": "John Doe",
				"id":          "123",
			},
			expected: "John Doe",
		},
		{
			name: "map with value",
			input: map[string]any{
				"value": "Option A",
				"id":    "opt-a",
			},
			expected: "Option A",
		},
		{
			name: "map with name",
			input: map[string]any{
				"name": "Status Name",
				"id":   "status-1",
			},
			expected: "Status Name",
		},
		{
			name: "map with only id",
			input: map[string]any{
				"id": "only-id",
			},
			expected: "map[id:only-id]",
		},
		{
			name: "array of maps with value",
			input: []any{
				map[string]any{"value": "Opt1"},
				map[string]any{"value": "Opt2"},
			},
			expected: "Opt1, Opt2",
		},
		{
			name: "array of maps with displayName",
			input: []any{
				map[string]any{"displayName": "User1"},
				map[string]any{"displayName": "User2"},
			},
			expected: "User1, User2",
		},
		{
			name: "array of maps with name",
			input: []any{
				map[string]any{"name": "Item1"},
				map[string]any{"name": "Item2"},
			},
			expected: "Item1, Item2",
		},
		{
			name:     "array of strings",
			input:    []any{"tag1", "tag2", "tag3"},
			expected: "tag1, tag2, tag3",
		},
		{
			name:     "empty array",
			input:    []any{},
			expected: "",
		},
		{
			name:     "array with unhandled types",
			input:    []any{123, 456},
			expected: "",
		},
		{
			name:     "integer",
			input:    int(99),
			expected: "99",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatCustomFieldValue(tt.input)
			if result != tt.expected {
				t.Errorf("formatCustomFieldValue(%v) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestExtractIssueCustomFields(t *testing.T) {
	tests := []struct {
		name     string
		fields   map[string]any
		nameMap  map[string]string
		expected []customFieldValue
	}{
		{
			name: "single custom field",
			fields: map[string]any{
				"customfield_10001": "value1",
				"summary":           "not a custom field",
			},
			nameMap: map[string]string{
				"customfield_10001": "Custom Field 1",
			},
			expected: []customFieldValue{
				{name: "Custom Field 1", value: "value1"},
			},
		},
		{
			name: "custom field not in nameMap uses ID",
			fields: map[string]any{
				"customfield_10002": "value2",
			},
			nameMap: map[string]string{},
			expected: []customFieldValue{
				{name: "customfield_10002", value: "value2"},
			},
		},
		{
			name: "nil custom field value is skipped",
			fields: map[string]any{
				"customfield_10003": nil,
			},
			nameMap:  map[string]string{},
			expected: nil,
		},
		{
			name: "multiple custom fields sorted by name",
			fields: map[string]any{
				"customfield_10001": "B Value",
				"customfield_10002": "A Value",
			},
			nameMap: map[string]string{
				"customfield_10001": "Zebra Field",
				"customfield_10002": "Apple Field",
			},
			expected: []customFieldValue{
				{name: "Apple Field", value: "A Value"},
				{name: "Zebra Field", value: "B Value"},
			},
		},
		{
			name:     "no custom fields",
			fields:   map[string]any{"summary": "regular field"},
			nameMap:  map[string]string{},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractIssueCustomFields(tt.fields, tt.nameMap)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("extractIssueCustomFields() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestExtractTextFromADF(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "simple paragraph",
			input: map[string]any{
				"type":    "doc",
				"version": 1,
				"content": []any{
					map[string]any{
						"type": "paragraph",
						"content": []any{
							map[string]any{
								"type": "text",
								"text": "Hello, World!",
							},
						},
					},
				},
			},
			expected: "Hello, World!",
		},
		{
			name: "multiple paragraphs",
			input: map[string]any{
				"type":    "doc",
				"version": 1,
				"content": []any{
					map[string]any{
						"type": "paragraph",
						"content": []any{
							map[string]any{"type": "text", "text": "First paragraph"},
						},
					},
					map[string]any{
						"type": "paragraph",
						"content": []any{
							map[string]any{"type": "text", "text": "Second paragraph"},
						},
					},
				},
			},
			expected: "First paragraph\nSecond paragraph",
		},
		{
			name: "multiple text nodes in same paragraph",
			input: map[string]any{
				"type":    "doc",
				"version": 1,
				"content": []any{
					map[string]any{
						"type": "paragraph",
						"content": []any{
							map[string]any{"type": "text", "text": "Hello "},
							map[string]any{"type": "text", "text": "World"},
						},
					},
				},
			},
			expected: "Hello World",
		},
		{
			name:     "nil input",
			input:    nil,
			expected: "",
		},
		{
			name:     "not a map",
			input:    "invalid",
			expected: "",
		},
		{
			name: "empty content",
			input: map[string]any{
				"type":    "doc",
				"version": 1,
				"content": []any{},
			},
			expected: "",
		},
		{
			name: "no content key",
			input: map[string]any{
				"type":    "doc",
				"version": 1,
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTextFromADF(tt.input)
			if result != tt.expected {
				t.Errorf("extractTextFromADF() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestExtractTextFromBlock(t *testing.T) {
	tests := []struct {
		name     string
		input    interface{}
		expected string
	}{
		{
			name: "text block",
			input: map[string]interface{}{
				"type": "text",
				"text": "Direct text",
			},
			expected: "Direct text",
		},
		{
			name: "nested content",
			input: map[string]interface{}{
				"type": "paragraph",
				"content": []interface{}{
					map[string]interface{}{"type": "text", "text": "Nested"},
				},
			},
			expected: "Nested",
		},
		{
			name:     "nil input",
			input:    nil,
			expected: "",
		},
		{
			name:     "not a map",
			input:    "string input",
			expected: "",
		},
		{
			name: "empty content array",
			input: map[string]interface{}{
				"type":    "paragraph",
				"content": []interface{}{},
			},
			expected: "",
		},
		// Additional tests for uncovered branches
		{
			name: "content is not an array",
			input: map[string]interface{}{
				"type":    "paragraph",
				"content": "not an array",
			},
			expected: "",
		},
		{
			name: "text field is not a string",
			input: map[string]interface{}{
				"type": "text",
				"text": 12345,
			},
			expected: "",
		},
		{
			name: "deeply nested content",
			input: map[string]interface{}{
				"type": "bulletList",
				"content": []interface{}{
					map[string]interface{}{
						"type": "listItem",
						"content": []interface{}{
							map[string]interface{}{
								"type": "paragraph",
								"content": []interface{}{
									map[string]interface{}{"type": "text", "text": "Item 1"},
								},
							},
						},
					},
					map[string]interface{}{
						"type": "listItem",
						"content": []interface{}{
							map[string]interface{}{
								"type": "paragraph",
								"content": []interface{}{
									map[string]interface{}{"type": "text", "text": "Item 2"},
								},
							},
						},
					},
				},
			},
			expected: "Item 1Item 2",
		},
		{
			name: "no text and no content",
			input: map[string]interface{}{
				"type": "hardBreak",
			},
			expected: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractTextFromBlock(tt.input)
			if result != tt.expected {
				t.Errorf("extractTextFromBlock() = %q, want %q", result, tt.expected)
			}
		})
	}
}
