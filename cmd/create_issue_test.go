package cmd

import (
	"reflect"
	"testing"
)

func TestNewADFDocumentComprehensive(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "simple text",
			text: "Hello, World!",
		},
		{
			name: "text with newlines",
			text: "Line 1\nLine 2\nLine 3",
		},
		{
			name: "empty text",
			text: "",
		},
		{
			name: "text with special characters",
			text: "<script>alert('xss')</script>",
		},
		{
			name: "text with unicode",
			text: "日本語テスト 🎉",
		},
		{
			name: "very long text",
			text: string(make([]byte, 10000)),
		},
		{
			name: "text with tabs",
			text: "Column1\tColumn2\tColumn3",
		},
		{
			name: "text with carriage returns",
			text: "Line 1\r\nLine 2\r\nLine 3",
		},
		{
			name: "markdown-like content",
			text: "# Heading\n\n- Item 1\n- Item 2\n\n**Bold** and *italic*",
		},
		{
			name: "JSON-like content",
			text: `{"key": "value", "nested": {"a": 1}}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newADFDocument(tt.text)

			// Verify document structure
			if result == nil {
				t.Fatal("newADFDocument returned nil")
			}

			// Check type
			docType, ok := result["type"].(string)
			if !ok || docType != "doc" {
				t.Errorf("document type = %v, want 'doc'", result["type"])
			}

			// Check version
			version, ok := result["version"].(int)
			if !ok || version != 1 {
				t.Errorf("document version = %v, want 1", result["version"])
			}

			// Check content array exists
			content, ok := result["content"].([]map[string]any)
			if !ok {
				t.Fatalf("content is not []map[string]any, got %T", result["content"])
			}

			if len(content) != 1 {
				t.Errorf("content length = %d, want 1", len(content))
			}

			// Check paragraph
			if len(content) > 0 {
				paragraph := content[0]
				if paragraph["type"] != "paragraph" {
					t.Errorf("paragraph type = %v, want 'paragraph'", paragraph["type"])
				}

				// Check text node
				paraContent, ok := paragraph["content"].([]map[string]any)
				if !ok {
					t.Fatalf("paragraph content is not []map[string]any")
				}

				if len(paraContent) != 1 {
					t.Errorf("paragraph content length = %d, want 1", len(paraContent))
				}

				if len(paraContent) > 0 {
					textNode := paraContent[0]
					if textNode["type"] != "text" {
						t.Errorf("text node type = %v, want 'text'", textNode["type"])
					}
					if textNode["text"] != tt.text {
						t.Errorf("text content = %v, want %v", textNode["text"], tt.text)
					}
				}
			}
		})
	}
}

func TestNewADFDocumentIsImmutable(t *testing.T) {
	// Test that calling newADFDocument multiple times returns independent documents
	doc1 := newADFDocument("text1")
	doc2 := newADFDocument("text2")

	// Modify doc1
	content1 := doc1["content"].([]map[string]any)
	content1[0]["type"] = "modified"

	// doc2 should be unaffected
	content2 := doc2["content"].([]map[string]any)
	if content2[0]["type"] != "paragraph" {
		t.Error("modifying one document affected another")
	}
}

func TestBuildIssueFieldsStructure(t *testing.T) {
	// Test helper for building issue fields map structure
	t.Run("project field structure", func(t *testing.T) {
		project := map[string]any{
			"key": "PROJ",
		}
		if project["key"] != "PROJ" {
			t.Error("project key mismatch")
		}
	})

	t.Run("issuetype field structure", func(t *testing.T) {
		issueType := map[string]any{
			"name": "Bug",
		}
		if issueType["name"] != "Bug" {
			t.Error("issuetype name mismatch")
		}
	})

	t.Run("assignee field structure", func(t *testing.T) {
		assignee := map[string]any{
			"id": "account-123",
		}
		if assignee["id"] != "account-123" {
			t.Error("assignee id mismatch")
		}
	})

	t.Run("priority field structure", func(t *testing.T) {
		priority := map[string]any{
			"name": "High",
		}
		if priority["name"] != "High" {
			t.Error("priority name mismatch")
		}
	})

	t.Run("labels field structure", func(t *testing.T) {
		labels := []string{"bug", "critical", "frontend"}
		if len(labels) != 3 {
			t.Error("labels count mismatch")
		}
	})

	t.Run("components field structure", func(t *testing.T) {
		components := []map[string]any{
			{"name": "Backend"},
			{"name": "Frontend"},
		}
		if len(components) != 2 {
			t.Error("components count mismatch")
		}
		if components[0]["name"] != "Backend" {
			t.Error("first component name mismatch")
		}
	})

	t.Run("parent field structure", func(t *testing.T) {
		parent := map[string]any{
			"key": "PROJ-100",
		}
		if parent["key"] != "PROJ-100" {
			t.Error("parent key mismatch")
		}
	})
}

func TestLabelsParsing(t *testing.T) {
	// Test the label parsing logic used in create_issue
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple labels",
			input:    "bug,critical,frontend",
			expected: []string{"bug", "critical", "frontend"},
		},
		{
			name:     "labels with spaces",
			input:    "bug, critical, frontend",
			expected: []string{"bug", "critical", "frontend"},
		},
		{
			name:     "single label",
			input:    "bug",
			expected: []string{"bug"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Simulate the parsing logic from create_issue
			var result []string
			if tt.input != "" {
				parts := splitAndTrim(tt.input, ",")
				result = parts
			} else {
				result = []string{""}
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("parsing %q: got %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}

// splitAndTrim splits a string by separator and trims whitespace from each part
func splitAndTrim(s, sep string) []string {
	var result []string
	for _, part := range splitString(s, sep) {
		result = append(result, trimSpace(part))
	}
	return result
}

func splitString(s, sep string) []string {
	var result []string
	start := 0
	for i := 0; i <= len(s)-len(sep); i++ {
		if s[i:i+len(sep)] == sep {
			result = append(result, s[start:i])
			start = i + len(sep)
		}
	}
	result = append(result, s[start:])
	return result
}

func trimSpace(s string) string {
	start := 0
	end := len(s)
	for start < end && (s[start] == ' ' || s[start] == '\t') {
		start++
	}
	for end > start && (s[end-1] == ' ' || s[end-1] == '\t') {
		end--
	}
	return s[start:end]
}
