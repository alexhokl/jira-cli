package cmd

import (
	"reflect"
	"testing"
)

func TestNewComment(t *testing.T) {
	tests := []struct {
		name    string
		comment string
	}{
		{
			name:    "simple comment",
			comment: "This is a test comment",
		},
		{
			name:    "comment with special characters",
			comment: "Comment with 'quotes' and \"double quotes\"",
		},
		{
			name:    "multiline comment",
			comment: "Line 1\nLine 2\nLine 3",
		},
		{
			name:    "empty comment",
			comment: "",
		},
		{
			name:    "comment with unicode",
			comment: "Test with unicode: 日本語 and emoji 🎉",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newComment(tt.comment)

			if result == nil {
				t.Fatal("newComment() returned nil")
			}

			if result.Body == nil {
				t.Fatal("newComment().Body is nil")
			}

			body, ok := result.Body.(map[string]interface{})
			if !ok {
				t.Fatalf("Body is not map[string]interface{}, got %T", result.Body)
			}

			// Check document type
			if docType, ok := body["type"].(string); !ok || docType != "doc" {
				t.Errorf("Body type = %v, want 'doc'", body["type"])
			}

			// Check version
			if version, ok := body["version"].(int); !ok || version != 1 {
				t.Errorf("Body version = %v, want 1", body["version"])
			}

			// Check content structure
			content, ok := body["content"].([]map[string]interface{})
			if !ok {
				t.Fatalf("Body content is not []map[string]interface{}, got %T", body["content"])
			}

			if len(content) != 1 {
				t.Fatalf("Expected 1 content block, got %d", len(content))
			}

			paragraph := content[0]
			if paragraph["type"] != "paragraph" {
				t.Errorf("Content block type = %v, want 'paragraph'", paragraph["type"])
			}

			paragraphContent, ok := paragraph["content"].([]map[string]interface{})
			if !ok {
				t.Fatalf("Paragraph content is not []map[string]interface{}, got %T", paragraph["content"])
			}

			if len(paragraphContent) != 1 {
				t.Fatalf("Expected 1 text node, got %d", len(paragraphContent))
			}

			textNode := paragraphContent[0]
			if textNode["type"] != "text" {
				t.Errorf("Text node type = %v, want 'text'", textNode["type"])
			}

			if textNode["text"] != tt.comment {
				t.Errorf("Text node text = %v, want %v", textNode["text"], tt.comment)
			}
		})
	}
}

func TestNewADFDocument(t *testing.T) {
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
			text: "Line 1\nLine 2",
		},
		{
			name: "empty text",
			text: "",
		},
		{
			name: "text with special characters",
			text: "<script>alert('xss')</script>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := newADFDocument(tt.text)

			if result == nil {
				t.Fatal("newADFDocument() returned nil")
			}

			// Check document type
			if docType, ok := result["type"].(string); !ok || docType != "doc" {
				t.Errorf("result type = %v, want 'doc'", result["type"])
			}

			// Check version
			if version, ok := result["version"].(int); !ok || version != 1 {
				t.Errorf("result version = %v, want 1", result["version"])
			}

			// Check content structure
			content, ok := result["content"].([]map[string]any)
			if !ok {
				t.Fatalf("result content is not []map[string]any, got %T", result["content"])
			}

			if len(content) != 1 {
				t.Fatalf("Expected 1 content block, got %d", len(content))
			}

			paragraph := content[0]
			if paragraph["type"] != "paragraph" {
				t.Errorf("Content block type = %v, want 'paragraph'", paragraph["type"])
			}

			paragraphContent, ok := paragraph["content"].([]map[string]any)
			if !ok {
				t.Fatalf("Paragraph content is not []map[string]any, got %T", paragraph["content"])
			}

			if len(paragraphContent) != 1 {
				t.Fatalf("Expected 1 text node, got %d", len(paragraphContent))
			}

			textNode := paragraphContent[0]
			if textNode["type"] != "text" {
				t.Errorf("Text node type = %v, want 'text'", textNode["type"])
			}

			if textNode["text"] != tt.text {
				t.Errorf("Text node text = %v, want %v", textNode["text"], tt.text)
			}
		})
	}
}

func TestNewADFDocumentStructure(t *testing.T) {
	// Test that the returned structure matches the expected ADF format exactly
	text := "Test text"
	result := newADFDocument(text)

	expected := map[string]any{
		"type":    "doc",
		"version": 1,
		"content": []map[string]any{
			{
				"type": "paragraph",
				"content": []map[string]any{
					{
						"type": "text",
						"text": text,
					},
				},
			},
		},
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("newADFDocument(%q) structure doesn't match expected ADF format", text)
		t.Errorf("got: %+v", result)
		t.Errorf("want: %+v", expected)
	}
}
