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
