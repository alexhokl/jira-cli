package cmd

import (
	"reflect"
	"testing"
)

func TestNewComment(t *testing.T) {
	tests := []struct {
		name           string
		comment        string
		expectedBlocks int
		firstBlockType string
	}{
		{
			name:           "simple comment",
			comment:        "This is a test comment",
			expectedBlocks: 1,
			firstBlockType: "paragraph",
		},
		{
			name:           "comment with special characters",
			comment:        "Comment with 'quotes' and \"double quotes\"",
			expectedBlocks: 1,
			firstBlockType: "paragraph",
		},
		{
			name:           "multiline comment",
			comment:        "Line 1\nLine 2\nLine 3",
			expectedBlocks: 3,
			firstBlockType: "paragraph",
		},
		{
			name:           "empty comment",
			comment:        "",
			expectedBlocks: 0,
			firstBlockType: "",
		},
		{
			name:           "comment with unicode",
			comment:        "Test with unicode: 日本語 and emoji 🎉",
			expectedBlocks: 1,
			firstBlockType: "paragraph",
		},
		{
			name:           "comment with markdown heading",
			comment:        "# Heading",
			expectedBlocks: 1,
			firstBlockType: "heading",
		},
		{
			name:           "comment with code block",
			comment:        "```\ncode\n```",
			expectedBlocks: 1,
			firstBlockType: "codeBlock",
		},
		{
			name:           "comment with bullet list",
			comment:        "- item 1\n- item 2",
			expectedBlocks: 1,
			firstBlockType: "bulletList",
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

			body, ok := result.Body.(map[string]any)
			if !ok {
				t.Fatalf("Body is not map[string]any, got %T", result.Body)
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
			content, ok := body["content"].([]map[string]any)
			if !ok {
				t.Fatalf("Body content is not []map[string]any, got %T", body["content"])
			}

			if len(content) != tt.expectedBlocks {
				t.Fatalf("Expected %d content block(s), got %d", tt.expectedBlocks, len(content))
			}

			// Check first block type if we expect content
			if tt.expectedBlocks > 0 {
				firstBlock := content[0]
				if blockType, ok := firstBlock["type"].(string); !ok || blockType != tt.firstBlockType {
					t.Errorf("First content block type = %v, want '%s'", firstBlock["type"], tt.firstBlockType)
				}
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
