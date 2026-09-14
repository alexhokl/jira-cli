package cmd

import (
	"encoding/json"
	"os"
	"testing"
)

// getExpand is a helper that asserts the doc has exactly one top-level node,
// that it is of type "expand", and returns it.
func getExpand(t *testing.T, doc map[string]any) map[string]any {
	t.Helper()
	content, ok := doc["content"].([]map[string]any)
	if !ok || len(content) == 0 {
		t.Fatalf("expected at least one content node, got %v", doc["content"])
	}
	node := content[0]
	if node["type"] != "expand" {
		t.Fatalf("expected node type 'expand', got %q", node["type"])
	}
	return node
}

func TestPrintJSON(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		wantKeys []string
	}{
		{
			name:     "struct with exported fields",
			input:    struct{ Name, Value string }{"hello", "world"},
			wantKeys: []string{"Name", "Value"},
		},
		{
			name:  "slice of strings",
			input: []string{"a", "b", "c"},
		},
		{
			name:  "nil slice",
			input: []string(nil),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Redirect stdout to a pipe so we can capture printJSON output
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe: %v", err)
			}
			origStdout := os.Stdout
			os.Stdout = w

			printErr := printJSON(tt.input)

			w.Close()
			os.Stdout = origStdout

			if printErr != nil {
				t.Fatalf("printJSON returned error: %v", printErr)
			}

			buf := make([]byte, 4096)
			n, _ := r.Read(buf)
			r.Close()
			output := buf[:n]

			// Verify the output is valid JSON
			var decoded any
			if err := json.Unmarshal(output, &decoded); err != nil {
				t.Fatalf("output is not valid JSON: %v\noutput: %s", err, output)
			}

			// Verify expected keys are present (for struct inputs)
			if len(tt.wantKeys) > 0 {
				m, ok := decoded.(map[string]any)
				if !ok {
					t.Fatalf("expected JSON object, got %T", decoded)
				}
				for _, key := range tt.wantKeys {
					if _, exists := m[key]; !exists {
						t.Errorf("expected key %q in JSON output, keys present: %v", key, m)
					}
				}
			}
		})
	}
}

func TestConvertMarkdownToADF_Expand(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, doc map[string]any)
	}{
		{
			name:  "basic expand with title",
			input: ":::expand My Title\nSome text\n:::",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				attrs, _ := node["attrs"].(map[string]any)
				if attrs["title"] != "My Title" {
					t.Errorf("expected title %q, got %q", "My Title", attrs["title"])
				}

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 1 {
					t.Fatalf("expected 1 child node, got %d", len(inner))
				}
				if inner[0]["type"] != "paragraph" {
					t.Errorf("expected child type 'paragraph', got %q", inner[0]["type"])
				}
			},
		},
		{
			name:  "expand with empty title",
			input: ":::expand\nContent\n:::",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				attrs, _ := node["attrs"].(map[string]any)
				if attrs["title"] != "" {
					t.Errorf("expected empty title, got %q", attrs["title"])
				}

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 1 {
					t.Fatalf("expected 1 child node, got %d", len(inner))
				}
			},
		},
		{
			name:  "expand with multiple paragraphs in body",
			input: ":::expand Multi\nFirst paragraph.\n\nSecond paragraph.\n:::",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 2 {
					t.Fatalf("expected 2 child nodes, got %d", len(inner))
				}
				for i, child := range inner {
					if child["type"] != "paragraph" {
						t.Errorf("child[%d]: expected 'paragraph', got %q", i, child["type"])
					}
				}
			},
		},
		{
			name:  "expand with bullet list inside",
			input: ":::expand List Section\n- item1\n- item2\n:::",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 1 {
					t.Fatalf("expected 1 child node, got %d", len(inner))
				}
				if inner[0]["type"] != "bulletList" {
					t.Errorf("expected child type 'bulletList', got %q", inner[0]["type"])
				}
			},
		},
		{
			name:  "expand with code block inside",
			input: ":::expand Code Section\n```go\nfmt.Println(\"hello\")\n```\n:::",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 1 {
					t.Fatalf("expected 1 child node, got %d", len(inner))
				}
				if inner[0]["type"] != "codeBlock" {
					t.Errorf("expected child type 'codeBlock', got %q", inner[0]["type"])
				}
				attrs, _ := inner[0]["attrs"].(map[string]any)
				if attrs["language"] != "go" {
					t.Errorf("expected language 'go', got %q", attrs["language"])
				}
			},
		},
		{
			name:  "expand surrounded by other blocks",
			input: "Before expand.\n:::expand Middle\nInside expand.\n:::\nAfter expand.",
			check: func(t *testing.T, doc map[string]any) {
				nodes, ok := doc["content"].([]map[string]any)
				if !ok || len(nodes) != 3 {
					t.Fatalf("expected 3 top-level nodes, got %d", len(nodes))
				}
				if nodes[0]["type"] != "paragraph" {
					t.Errorf("nodes[0]: expected 'paragraph', got %q", nodes[0]["type"])
				}
				if nodes[1]["type"] != "expand" {
					t.Errorf("nodes[1]: expected 'expand', got %q", nodes[1]["type"])
				}
				if nodes[2]["type"] != "paragraph" {
					t.Errorf("nodes[2]: expected 'paragraph', got %q", nodes[2]["type"])
				}

				attrs, _ := nodes[1]["attrs"].(map[string]any)
				if attrs["title"] != "Middle" {
					t.Errorf("expected title 'Middle', got %q", attrs["title"])
				}
			},
		},
		{
			name:  "unclosed expand emits node with captured body",
			input: ":::expand Unclosed\nSome content without closing fence",
			check: func(t *testing.T, doc map[string]any) {
				node := getExpand(t, doc)

				attrs, _ := node["attrs"].(map[string]any)
				if attrs["title"] != "Unclosed" {
					t.Errorf("expected title 'Unclosed', got %q", attrs["title"])
				}

				inner, _ := node["content"].([]map[string]any)
				if len(inner) != 1 {
					t.Fatalf("expected 1 child node, got %d", len(inner))
				}
				if inner[0]["type"] != "paragraph" {
					t.Errorf("expected child type 'paragraph', got %q", inner[0]["type"])
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doc := convertMarkdownToADF(tc.input)
			tc.check(t, doc)
		})
	}
}

func TestConvertMarkdownToADF_Table(t *testing.T) {
	tests := []struct {
		name  string
		input string
		check func(t *testing.T, doc map[string]any)
	}{
		{
			name:  "table with header and data rows",
			input: "| Container | Id |\n| --- | --- |\n| 350885 | 1 |\n| 351232 | 3 |",
			check: func(t *testing.T, doc map[string]any) {
				content, ok := doc["content"].([]map[string]any)
				if !ok || len(content) != 1 || content[0]["type"] != "table" {
					t.Fatalf("expected a single table node, got %v", doc["content"])
				}

				rows, _ := content[0]["content"].([]map[string]any)
				if len(rows) != 3 {
					t.Fatalf("expected 3 tableRow nodes, got %d", len(rows))
				}

				for rowIdx, row := range rows {
					if row["type"] != "tableRow" {
						t.Errorf("rows[%d]: expected type 'tableRow', got %q", rowIdx, row["type"])
					}

					cells, ok := row["content"].([]map[string]any)
					if !ok {
						t.Fatalf("rows[%d]: expected content to be a slice, got %T", rowIdx, row["content"])
					}
					if len(cells) != 2 {
						t.Errorf("rows[%d]: expected 2 cells, got %d", rowIdx, len(cells))
					}

					for cellIdx, cell := range cells {
						if cell["type"] != "tableCell" {
							t.Errorf("rows[%d].cells[%d]: expected type 'tableCell', got %q", rowIdx, cellIdx, cell["type"])
						}

						attrs, ok := cell["attrs"].(map[string]any)
						if !ok {
							t.Fatalf("rows[%d].cells[%d]: expected attrs map, got %T", rowIdx, cellIdx, cell["attrs"])
						}

						// Jira's ADF schema requires colwidth to be an array of
						// numbers; emitting nil (JSON null) is rejected with a
						// 400 Bad Request.
						if _, exists := attrs["colwidth"]; exists {
							t.Errorf("rows[%d].cells[%d]: attrs must not contain 'colwidth', got %v", rowIdx, cellIdx, attrs["colwidth"])
						}
						if attrs["colspan"] != 1 {
							t.Errorf("rows[%d].cells[%d]: expected colspan 1, got %v", rowIdx, cellIdx, attrs["colspan"])
						}
						if attrs["rowspan"] != 1 {
							t.Errorf("rows[%d].cells[%d]: expected rowspan 1, got %v", rowIdx, cellIdx, attrs["rowspan"])
						}

						paras, ok := cell["content"].([]map[string]any)
						if !ok || len(paras) != 1 || paras[0]["type"] != "paragraph" {
							t.Errorf("rows[%d].cells[%d]: expected one paragraph child, got %v", rowIdx, cellIdx, cell["content"])
						}
					}
				}
			},
		},
		{
			name:  "table cell with inline code",
			input: "| Col |\n| --- |\n| `api-transport` |",
			check: func(t *testing.T, doc map[string]any) {
				content, ok := doc["content"].([]map[string]any)
				if !ok || len(content) != 1 || content[0]["type"] != "table" {
					t.Fatalf("expected a single table node, got %v", doc["content"])
				}

				rows, _ := content[0]["content"].([]map[string]any)
				if len(rows) != 2 {
					t.Fatalf("expected 2 tableRow nodes, got %d", len(rows))
				}

				dataRow := rows[1]
				cells, _ := dataRow["content"].([]map[string]any)
				if len(cells) != 1 {
					t.Fatalf("expected 1 cell in data row, got %d", len(cells))
				}

				paras, _ := cells[0]["content"].([]map[string]any)
				if len(paras) != 1 {
					t.Fatalf("expected 1 paragraph in cell, got %d", len(paras))
				}

				texts, _ := paras[0]["content"].([]map[string]any)
				if len(texts) != 1 {
					t.Fatalf("expected 1 text node in paragraph, got %d", len(texts))
				}

				marks, _ := texts[0]["marks"].([]map[string]any)
				if len(marks) != 1 || marks[0]["type"] != "code" {
					t.Errorf("expected code mark on cell text, got %v", marks)
				}
				if texts[0]["text"] != "api-transport" {
					t.Errorf("expected text 'api-transport', got %q", texts[0]["text"])
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doc := convertMarkdownToADF(tc.input)
			tc.check(t, doc)
		})
	}
}
