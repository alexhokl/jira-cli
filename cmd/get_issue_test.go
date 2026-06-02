package cmd

import (
	"testing"
)

func adfDoc(blocks ...any) map[string]any {
	return map[string]any{
		"version": float64(1),
		"type":    "doc",
		"content": blocks,
	}
}

func adfParagraph(text string) map[string]any {
	return map[string]any{
		"type": "paragraph",
		"content": []any{
			map[string]any{
				"type": "text",
				"text": text,
			},
		},
	}
}

func adfExpand(nodeType, title string, children ...any) map[string]any {
	return map[string]any{
		"type": nodeType,
		"attrs": map[string]any{
			"title": title,
		},
		"content": children,
	}
}

func TestConvertADFToMarkdown_Expand(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name: "expand with title and paragraph child",
			input: adfDoc(
				adfExpand("expand", "Section Title",
					adfParagraph("Some content inside the expand."),
				),
			),
			expected: "**Section Title**\n  Some content inside the expand.\n\n",
		},
		{
			name: "expand with no title",
			input: adfDoc(
				adfExpand("expand", "",
					adfParagraph("Content with no title."),
				),
			),
			expected: "  Content with no title.\n\n",
		},
		{
			name: "expand with multiple child blocks",
			input: adfDoc(
				adfExpand("expand", "Multi Block",
					adfParagraph("First paragraph."),
					adfParagraph("Second paragraph."),
				),
			),
			// Each paragraph produces "text\n"; after TrimRight+Split they are
			// consecutive lines with no blank line between them.
			expected: "**Multi Block**\n  First paragraph.\n  Second paragraph.\n\n",
		},
		{
			name: "nestedExpand renders same as expand",
			input: adfDoc(
				adfExpand("nestedExpand", "Nested Title",
					adfParagraph("Nested content."),
				),
			),
			expected: "**Nested Title**\n  Nested content.\n\n",
		},
		{
			name: "expand surrounded by other blocks",
			input: adfDoc(
				adfParagraph("Before expand."),
				adfExpand("expand", "Collapsible",
					adfParagraph("Inside expand."),
				),
				adfParagraph("After expand."),
			),
			// convertADFToMarkdown adds "\n" between blocks (i > 0 separator),
			// and the expand block itself ends with "\n", producing an extra blank
			// line before the following paragraph.
			expected: "Before expand.\n\n**Collapsible**\n  Inside expand.\n\n\nAfter expand.\n",
		},
		{
			name: "expand with no content",
			input: adfDoc(
				adfExpand("expand", "Empty Section"),
			),
			// Empty children list produces no indented lines; only the title and
			// the trailing newline are written.
			expected: "**Empty Section**\n\n",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := convertADFToMarkdown(tc.input)
			if got != tc.expected {
				t.Errorf("convertADFToMarkdown()\ngot:\n%q\nwant:\n%q", got, tc.expected)
			}
		})
	}
}
