package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"runtime"
	"strings"
	"text/tabwriter"
	"unicode/utf8"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/alexhokl/jira-cli/swagger_software"
	"github.com/spf13/viper"
)

// timeFixTransport is a custom HTTP transport that fixes Jira's timestamp format
// by adding a colon to timezone offsets (e.g., +0800 -> +08:00)
type timeFixTransport struct {
	transport http.RoundTripper
}

// timezoneRegex matches timezone offsets without colons (e.g., +0800, -0530)
var timezoneRegex = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?)([\+\-])(\d{2})(\d{2})`)

func (t *timeFixTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Only process JSON responses
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return resp, nil
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Fix timezone format: +0800 -> +08:00
	fixedBody := timezoneRegex.ReplaceAll(body, []byte("${1}${2}${3}:${4}"))

	// Replace the response body with the fixed content
	resp.Body = io.NopCloser(bytes.NewReader(fixedBody))
	resp.ContentLength = int64(len(fixedBody))

	return resp, nil
}

func getAuthContext() context.Context {
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	auth := swagger.BasicAuth{
		UserName: email,
		Password: apiKey,
	}

	ctx := context.WithValue(context.Background(), swagger.ContextBasicAuth, auth)
	return ctx
}

func getConfiguration() *swagger.Configuration {
	configuration := swagger.NewConfiguration()
	configuration.Servers[0].URL = fmt.Sprintf("https://%s.atlassian.net", viper.GetString("organization"))
	configuration.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return configuration
}

func newClient() *swagger.APIClient {
	return swagger.NewAPIClient(getConfiguration())
}

func getSoftwareAuthContext() context.Context {
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	auth := swagger_software.BasicAuth{
		UserName: email,
		Password: apiKey,
	}

	ctx := context.WithValue(context.Background(), swagger_software.ContextBasicAuth, auth)
	return ctx
}

func getSoftwareConfiguration() *swagger_software.Configuration {
	configuration := swagger_software.NewConfiguration()
	configuration.Servers[0].URL = fmt.Sprintf("https://%s.atlassian.net", viper.GetString("organization"))
	configuration.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return configuration
}

func newSoftwareClient() *swagger_software.APIClient {
	return swagger_software.NewAPIClient(getSoftwareConfiguration())
}

func getMessageFromEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "", fmt.Errorf("EDITOR environment variable is not set")
	}

	tmpFile, err := os.CreateTemp("", "jira-comment-*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFileName)

	cmd := exec.Command(editor, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run editor: %w", err)
	}

	content, err := os.ReadFile(tmpFileName)
	if err != nil {
		return "", fmt.Errorf("failed to read temporary file: %w", err)
	}

	return strings.TrimSpace(string(content)), nil
}

// ansiRegex matches ANSI escape sequences
var ansiRegex = regexp.MustCompile(`\x1b\[[0-9;]*m`)

// stripAnsi removes ANSI escape sequences from a string
func stripAnsi(s string) string {
	return ansiRegex.ReplaceAllString(s, "")
}

// tableWriter provides ANSI-aware tabular output
// It calculates column widths ignoring ANSI escape codes
type tableWriter struct {
	out      io.Writer
	rows     [][]string
	minWidth int
	padding  int
	padChar  byte
}

// newTableWriter creates a new tableWriter
// minWidth is the minimum column width, padding is the space between columns
func newTableWriter(out io.Writer, minWidth, padding int) *tableWriter {
	return &tableWriter{
		out:      out,
		rows:     make([][]string, 0),
		minWidth: minWidth,
		padding:  padding,
		padChar:  ' ',
	}
}

// row adds a row of cells to the table
func (t *tableWriter) row(cells ...string) {
	t.rows = append(t.rows, cells)
}

// flush calculates column widths and writes the formatted table
func (t *tableWriter) flush() {
	if len(t.rows) == 0 {
		return
	}

	// Calculate max visible width for each column
	maxCols := 0
	for _, row := range t.rows {
		if len(row) > maxCols {
			maxCols = len(row)
		}
	}

	colWidths := make([]int, maxCols)
	for _, row := range t.rows {
		for i, cell := range row {
			// Calculate visible width (without ANSI codes)
			visibleWidth := utf8.RuneCountInString(stripAnsi(cell))
			if visibleWidth > colWidths[i] {
				colWidths[i] = visibleWidth
			}
		}
	}

	// Apply minimum width
	for i := range colWidths {
		if colWidths[i] < t.minWidth {
			colWidths[i] = t.minWidth
		}
	}

	// Write rows with proper padding
	for _, row := range t.rows {
		for i, cell := range row {
			fmt.Fprint(t.out, cell)

			// Add padding (except for last column)
			if i < len(row)-1 {
				visibleWidth := utf8.RuneCountInString(stripAnsi(cell))
				padding := colWidths[i] - visibleWidth + t.padding
				for j := 0; j < padding; j++ {
					fmt.Fprint(t.out, string(t.padChar))
				}
			}
		}
		fmt.Fprintln(t.out)
	}
}

// legacyTabWriter returns a standard tabwriter for cases where ANSI codes aren't used
// This is kept for backward compatibility
func legacyTabWriter() *tabwriter.Writer {
	return tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
}

// normalizeUserValue converts user-friendly aliases to their API equivalents.
// Specifically, it converts "me" to "currentUser()" for use in JQL queries.
func normalizeUserValue(value string) string {
	if value == "me" {
		return "currentUser()"
	}
	return value
}

// resolveUserAlias resolves user aliases like "me" to actual account IDs.
// If the value is "me", it fetches and returns the current user's account ID.
// Otherwise, it returns the original value unchanged.
func resolveUserAlias(value string) (string, error) {
	if value != "me" {
		return value, nil
	}

	client := newClient()
	ctx := getAuthContext()

	currentUser, _, err := client.MyselfAPI.GetCurrentUser(ctx).Execute()
	if err != nil {
		return "", fmt.Errorf("failed to get current user: %w", err)
	}

	return currentUser.GetAccountId(), nil
}

// wrapAPIError wraps an API error with detailed error body if available.
// This works with both swagger.GenericOpenAPIError and swagger_software.GenericOpenAPIError.
func wrapAPIError(err error) error {
	if err == nil {
		return nil
	}

	// Try swagger.GenericOpenAPIError
	var swaggerErr *swagger.GenericOpenAPIError
	if errors.As(err, &swaggerErr) {
		body := string(swaggerErr.Body())
		if body != "" {
			return fmt.Errorf("%s\n\n%s", err, body)
		}
	}

	// Try swagger_software.GenericOpenAPIError
	var softwareErr *swagger_software.GenericOpenAPIError
	if errors.As(err, &softwareErr) {
		body := string(softwareErr.Body())
		if body != "" {
			return fmt.Errorf("%s\n\n%s", err, body)
		}
	}

	return err
}

// sprintInfo holds basic information about a sprint
type sprintInfo struct {
	id   int
	name string
}

// lookupSprintByName searches for a sprint by name in the boards associated with a project.
// It returns the sprint ID if found, or an error if not found or if multiple matches exist.
func lookupSprintByName(projectKey string, sprintName string) (int, error) {
	client := newSoftwareClient()
	ctx := getSoftwareAuthContext()

	// Find boards for the project
	boards, err := getBoardsForProject(client, ctx, projectKey)
	if err != nil {
		return 0, fmt.Errorf("failed to find boards for project %s: %w", projectKey, err)
	}

	if len(boards) == 0 {
		return 0, fmt.Errorf("no boards found for project %s; use sprint ID instead (run 'jira-cli list sprints --board-id <id>')", projectKey)
	}

	// Search for the sprint in all boards
	var matchingSprints []sprintInfo
	var searchedBoardIds []int64

	for _, board := range boards {
		boardId := board.GetId()
		searchedBoardIds = append(searchedBoardIds, boardId)

		sprints, err := getSprintsForBoard(client, ctx, boardId)
		if err != nil {
			// Skip boards that fail (may not have sprints enabled)
			continue
		}

		for _, sprint := range sprints {
			if strings.EqualFold(sprint.GetName(), sprintName) {
				matchingSprints = append(matchingSprints, sprintInfo{
					id:   int(sprint.GetId()),
					name: sprint.GetName(),
				})
			}
		}
	}

	if len(matchingSprints) == 0 {
		return 0, fmt.Errorf("sprint %q not found in project %s; use 'jira-cli list sprints --board-id <id>' to find available sprints", sprintName, projectKey)
	}

	if len(matchingSprints) > 1 {
		// Check if all matches are the same sprint ID (same sprint on multiple boards)
		firstId := matchingSprints[0].id
		allSame := true
		for _, s := range matchingSprints[1:] {
			if s.id != firstId {
				allSame = false
				break
			}
		}
		if !allSame {
			return 0, fmt.Errorf("multiple sprints named %q found; use sprint ID instead", sprintName)
		}
	}

	return matchingSprints[0].id, nil
}

// getBoardsForProject retrieves all boards associated with a project
func getBoardsForProject(client *swagger_software.APIClient, ctx context.Context, projectKey string) ([]swagger_software.GetAllBoards200ResponseValuesInner, error) {
	var allBoards []swagger_software.GetAllBoards200ResponseValuesInner
	var startAt int64 = 0

	for {
		result, _, err := client.BoardAPI.GetAllBoards(ctx).
			ProjectKeyOrId(projectKey).
			StartAt(startAt).
			Execute()
		if err != nil {
			return nil, err
		}

		allBoards = append(allBoards, result.GetValues()...)

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	return allBoards, nil
}

// getSprintsForBoard retrieves all sprints for a board
func getSprintsForBoard(client *swagger_software.APIClient, ctx context.Context, boardId int64) ([]swagger_software.SprintBean, error) {
	var allSprints []swagger_software.SprintBean
	var startAt int64 = 0

	for {
		resp, err := client.BoardAPI.GetAllSprints(ctx, boardId).StartAt(startAt).Execute()
		if err != nil {
			return nil, err
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("failed to read response body: %w", err)
		}

		var result sprintsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			return nil, fmt.Errorf("failed to parse response: %w", err)
		}

		allSprints = append(allSprints, result.Values...)

		if result.IsLast {
			break
		}

		startAt = int64(result.StartAt + result.MaxResults)
	}

	return allSprints, nil
}

// newADFDocument creates an Atlassian Document Format (ADF) document from plain text.
// ADF is the format used by Jira Cloud for rich text fields like description and comments.
func newADFDocument(text string) map[string]any {
	return map[string]any{
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
}

// convertMarkdownToADF converts markdown text to Atlassian Document Format (ADF).
// This preserves markdown formatting like headings, lists, code blocks, etc.
func convertMarkdownToADF(markdown string) map[string]any {
	lines := strings.Split(markdown, "\n")
	var content []map[string]any

	i := 0
	for i < len(lines) {
		line := lines[i]

		// Skip empty lines
		if strings.TrimSpace(line) == "" {
			i++
			continue
		}

		// Check for code block
		if strings.HasPrefix(line, "```") {
			language := strings.TrimPrefix(line, "```")
			var codeLines []string
			i++
			for i < len(lines) && !strings.HasPrefix(lines[i], "```") {
				codeLines = append(codeLines, lines[i])
				i++
			}
			i++ // skip closing ```

			codeBlock := map[string]any{
				"type": "codeBlock",
				"content": []map[string]any{
					{
						"type": "text",
						"text": strings.Join(codeLines, "\n"),
					},
				},
			}
			if language != "" {
				codeBlock["attrs"] = map[string]any{
					"language": language,
				}
			}
			content = append(content, codeBlock)
			continue
		}

		// Check for heading
		if strings.HasPrefix(line, "#") {
			level := 0
			for _, c := range line {
				if c == '#' {
					level++
				} else {
					break
				}
			}
			if level > 0 && level <= 6 {
				text := strings.TrimSpace(strings.TrimLeft(line, "#"))
				content = append(content, map[string]any{
					"type": "heading",
					"attrs": map[string]any{
						"level": level,
					},
					"content": parseInlineMarkdown(text),
				})
				i++
				continue
			}
		}

		// Check for horizontal rule
		if line == "---" || line == "***" || line == "___" {
			content = append(content, map[string]any{
				"type": "rule",
			})
			i++
			continue
		}

		// Check for blockquote
		if strings.HasPrefix(line, "> ") {
			var quoteContent []map[string]any
			for i < len(lines) && strings.HasPrefix(lines[i], "> ") {
				quoteLine := strings.TrimPrefix(lines[i], "> ")
				quoteContent = append(quoteContent, map[string]any{
					"type":    "paragraph",
					"content": parseInlineMarkdown(quoteLine),
				})
				i++
			}
			content = append(content, map[string]any{
				"type":    "blockquote",
				"content": quoteContent,
			})
			continue
		}

		// Check for unordered list
		if strings.HasPrefix(line, "- ") || strings.HasPrefix(line, "* ") {
			listItems, newIndex := parseList(lines, i, false)
			content = append(content, map[string]any{
				"type":    "bulletList",
				"content": listItems,
			})
			i = newIndex
			continue
		}

		// Check for ordered list
		if isOrderedListItem(line) {
			listItems, newIndex := parseList(lines, i, true)
			content = append(content, map[string]any{
				"type":    "orderedList",
				"content": listItems,
			})
			i = newIndex
			continue
		}

		// Check for table
		if strings.HasPrefix(line, "|") && strings.HasSuffix(strings.TrimSpace(line), "|") {
			tableRows, newIndex := parseTable(lines, i)
			if len(tableRows) > 0 {
				content = append(content, map[string]any{
					"type":    "table",
					"attrs":   map[string]any{"isNumberColumnEnabled": false, "layout": "default"},
					"content": tableRows,
				})
			}
			i = newIndex
			continue
		}

		// Regular paragraph
		content = append(content, map[string]any{
			"type":    "paragraph",
			"content": parseInlineMarkdown(line),
		})
		i++
	}

	return map[string]any{
		"type":    "doc",
		"version": 1,
		"content": content,
	}
}

// parseInlineMarkdown parses inline markdown formatting (bold, italic, code, links, etc.)
func parseInlineMarkdown(text string) []map[string]any {
	if text == "" {
		return nil
	}

	var result []map[string]any
	remaining := text

	for len(remaining) > 0 {
		// Check for inline code
		if idx := strings.Index(remaining, "`"); idx != -1 {
			endIdx := strings.Index(remaining[idx+1:], "`")
			if endIdx != -1 {
				// Add text before code
				if idx > 0 {
					result = append(result, parseFormattedText(remaining[:idx])...)
				}
				// Add code
				codeText := remaining[idx+1 : idx+1+endIdx]
				result = append(result, map[string]any{
					"type": "text",
					"text": codeText,
					"marks": []map[string]any{
						{"type": "code"},
					},
				})
				remaining = remaining[idx+1+endIdx+1:]
				continue
			}
		}

		// Check for links [text](url)
		if idx := strings.Index(remaining, "["); idx != -1 {
			endBracket := strings.Index(remaining[idx:], "](")
			if endBracket != -1 {
				endParen := strings.Index(remaining[idx+endBracket+2:], ")")
				if endParen != -1 {
					// Add text before link
					if idx > 0 {
						result = append(result, parseFormattedText(remaining[:idx])...)
					}
					linkText := remaining[idx+1 : idx+endBracket]
					linkURL := remaining[idx+endBracket+2 : idx+endBracket+2+endParen]
					result = append(result, map[string]any{
						"type": "text",
						"text": linkText,
						"marks": []map[string]any{
							{
								"type": "link",
								"attrs": map[string]any{
									"href": linkURL,
								},
							},
						},
					})
					remaining = remaining[idx+endBracket+2+endParen+1:]
					continue
				}
			}
		}

		// No more special inline elements, parse the rest for bold/italic
		result = append(result, parseFormattedText(remaining)...)
		break
	}

	return result
}

// parseFormattedText parses bold and italic formatting
func parseFormattedText(text string) []map[string]any {
	if text == "" {
		return nil
	}

	var result []map[string]any
	remaining := text

	for len(remaining) > 0 {
		// Check for bold **text**
		if idx := strings.Index(remaining, "**"); idx != -1 {
			endIdx := strings.Index(remaining[idx+2:], "**")
			if endIdx != -1 {
				// Add text before bold
				if idx > 0 {
					result = append(result, map[string]any{
						"type": "text",
						"text": remaining[:idx],
					})
				}
				// Add bold text
				boldText := remaining[idx+2 : idx+2+endIdx]
				result = append(result, map[string]any{
					"type": "text",
					"text": boldText,
					"marks": []map[string]any{
						{"type": "strong"},
					},
				})
				remaining = remaining[idx+2+endIdx+2:]
				continue
			}
		}

		// Check for italic *text* (but not **)
		if idx := strings.Index(remaining, "*"); idx != -1 {
			// Make sure it's not **
			if idx+1 < len(remaining) && remaining[idx+1] != '*' {
				endIdx := strings.Index(remaining[idx+1:], "*")
				if endIdx != -1 && (idx+1+endIdx+1 >= len(remaining) || remaining[idx+1+endIdx+1] != '*') {
					// Add text before italic
					if idx > 0 {
						result = append(result, map[string]any{
							"type": "text",
							"text": remaining[:idx],
						})
					}
					// Add italic text
					italicText := remaining[idx+1 : idx+1+endIdx]
					result = append(result, map[string]any{
						"type": "text",
						"text": italicText,
						"marks": []map[string]any{
							{"type": "em"},
						},
					})
					remaining = remaining[idx+1+endIdx+1:]
					continue
				}
			}
		}

		// Check for strikethrough ~~text~~
		if idx := strings.Index(remaining, "~~"); idx != -1 {
			endIdx := strings.Index(remaining[idx+2:], "~~")
			if endIdx != -1 {
				// Add text before strikethrough
				if idx > 0 {
					result = append(result, map[string]any{
						"type": "text",
						"text": remaining[:idx],
					})
				}
				// Add strikethrough text
				strikeText := remaining[idx+2 : idx+2+endIdx]
				result = append(result, map[string]any{
					"type": "text",
					"text": strikeText,
					"marks": []map[string]any{
						{"type": "strike"},
					},
				})
				remaining = remaining[idx+2+endIdx+2:]
				continue
			}
		}

		// No more formatting, add remaining text
		result = append(result, map[string]any{
			"type": "text",
			"text": remaining,
		})
		break
	}

	return result
}

// isOrderedListItem checks if a line is an ordered list item (e.g., "1. text")
func isOrderedListItem(line string) bool {
	for i, c := range line {
		if c >= '0' && c <= '9' {
			continue
		}
		if c == '.' && i > 0 && i+1 < len(line) && line[i+1] == ' ' {
			return true
		}
		break
	}
	return false
}

// getListIndent returns the indentation level of a list item
func getListIndent(line string) int {
	indent := 0
	for _, c := range line {
		if c == ' ' {
			indent++
		} else if c == '\t' {
			indent += 2
		} else {
			break
		}
	}
	return indent / 2
}

// parseList parses a list (bullet or ordered) and returns list items
func parseList(lines []string, startIndex int, ordered bool) ([]map[string]any, int) {
	var items []map[string]any
	i := startIndex
	baseIndent := getListIndent(lines[i])

	for i < len(lines) {
		line := lines[i]
		trimmedLine := strings.TrimSpace(line)

		// Empty line ends the list
		if trimmedLine == "" {
			break
		}

		currentIndent := getListIndent(line)

		// If we're back to a lower indent, we're done with this list level
		if currentIndent < baseIndent {
			break
		}

		// Check if this is a list item at our level
		isUnordered := strings.HasPrefix(trimmedLine, "- ") || strings.HasPrefix(trimmedLine, "* ")
		isOrdered := isOrderedListItem(trimmedLine)

		if currentIndent == baseIndent && (isUnordered || isOrdered) {
			// Extract the text after the list marker
			var itemText string
			if isUnordered {
				if strings.HasPrefix(trimmedLine, "- ") {
					itemText = strings.TrimPrefix(trimmedLine, "- ")
				} else {
					itemText = strings.TrimPrefix(trimmedLine, "* ")
				}
			} else {
				// Remove "N. " prefix
				dotIdx := strings.Index(trimmedLine, ". ")
				if dotIdx != -1 {
					itemText = trimmedLine[dotIdx+2:]
				}
			}

			itemContent := []map[string]any{
				{
					"type":    "paragraph",
					"content": parseInlineMarkdown(itemText),
				},
			}

			// Check for nested list
			i++
			if i < len(lines) {
				nextIndent := getListIndent(lines[i])
				nextTrimmed := strings.TrimSpace(lines[i])
				nextIsUnordered := strings.HasPrefix(nextTrimmed, "- ") || strings.HasPrefix(nextTrimmed, "* ")
				nextIsOrdered := isOrderedListItem(nextTrimmed)

				if nextIndent > baseIndent && (nextIsUnordered || nextIsOrdered) {
					nestedItems, newIndex := parseList(lines, i, nextIsOrdered)
					if nextIsOrdered {
						itemContent = append(itemContent, map[string]any{
							"type":    "orderedList",
							"content": nestedItems,
						})
					} else {
						itemContent = append(itemContent, map[string]any{
							"type":    "bulletList",
							"content": nestedItems,
						})
					}
					i = newIndex
				}
			}

			items = append(items, map[string]any{
				"type":    "listItem",
				"content": itemContent,
			})
		} else {
			break
		}
	}

	return items, i
}

// parseTable parses a markdown table and returns table rows
func parseTable(lines []string, startIndex int) ([]map[string]any, int) {
	var rows []map[string]any
	i := startIndex

	for i < len(lines) {
		line := strings.TrimSpace(lines[i])

		// Check if this is still a table row
		if !strings.HasPrefix(line, "|") || !strings.HasSuffix(line, "|") {
			break
		}

		// Skip separator row (|---|---|)
		if strings.Contains(line, "---") {
			i++
			continue
		}

		// Parse cells
		cells := strings.Split(line, "|")
		var cellContent []map[string]any

		for _, cell := range cells {
			cell = strings.TrimSpace(cell)
			if cell == "" {
				continue
			}

			cellContent = append(cellContent, map[string]any{
				"type": "tableCell",
				"attrs": map[string]any{
					"colspan":  1,
					"rowspan":  1,
					"colwidth": nil,
				},
				"content": []map[string]any{
					{
						"type":    "paragraph",
						"content": parseInlineMarkdown(cell),
					},
				},
			})
		}

		if len(cellContent) > 0 {
			rows = append(rows, map[string]any{
				"type":    "tableRow",
				"content": cellContent,
			})
		}

		i++
	}

	return rows, i
}

// extractProjectKeyFromIssueKey extracts the project key from an issue key.
// For example, "PROJ-123" returns "PROJ", "MYPROJECT-456" returns "MYPROJECT".
// If no dash is found, the entire issue key is returned.
func extractProjectKeyFromIssueKey(issueKey string) string {
	if idx := strings.LastIndex(issueKey, "-"); idx != -1 {
		return issueKey[:idx]
	}
	return issueKey
}

// filterBoardsByType filters boards by the specified type.
// This is a workaround for the swagger-generated code which incorrectly serializes
// the type parameter as type[type]=value instead of type=value.
func filterBoardsByType(boards []swagger_software.GetAllBoards200ResponseValuesInner, boardType string) []swagger_software.GetAllBoards200ResponseValuesInner {
	var filtered []swagger_software.GetAllBoards200ResponseValuesInner
	for _, board := range boards {
		if board.GetType() == boardType {
			filtered = append(filtered, board)
		}
	}
	return filtered
}

// openBrowser opens the specified URL in the default browser.
// It uses platform-specific commands: open on macOS, xdg-open on Linux, and cmd /c start on Windows.
func openBrowser(url string) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "darwin":
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	return cmd.Start()
}

// getIssueURL returns the URL to view an issue in the browser.
func getIssueURL(issueKey string) string {
	return fmt.Sprintf("https://%s.atlassian.net/browse/%s", viper.GetString("organization"), issueKey)
}
