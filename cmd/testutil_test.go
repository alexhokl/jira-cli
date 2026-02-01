package cmd

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/alexhokl/jira-cli/swagger_software"
	"github.com/spf13/viper"
)

// mockServerConfig holds configuration for creating mock test servers
type mockServerConfig struct {
	handler http.HandlerFunc
}

// newMockServer creates a new httptest.Server with the given handler
func newMockServer(handler http.HandlerFunc) *httptest.Server {
	return httptest.NewServer(handler)
}

// newMockClientWithServer creates a swagger API client configured to use the mock server
func newMockClientWithServer(serverURL string) *swagger.APIClient {
	cfg := swagger.NewConfiguration()
	cfg.Servers[0].URL = serverURL
	cfg.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return swagger.NewAPIClient(cfg)
}

// newMockSoftwareClientWithServer creates a swagger_software API client configured to use the mock server
func newMockSoftwareClientWithServer(serverURL string) *swagger_software.APIClient {
	cfg := swagger_software.NewConfiguration()
	cfg.Servers[0].URL = serverURL
	cfg.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return swagger_software.NewAPIClient(cfg)
}

// setupTestViper configures viper with test values
func setupTestViper() {
	viper.Set("email", "test@example.com")
	viper.Set("api_key", "test-api-key")
	viper.Set("organization", "testorg")
}

// resetTestViper resets viper to clean state
func resetTestViper() {
	viper.Reset()
}

// mockAuthContext returns a context with test authentication
func mockAuthContext() context.Context {
	auth := swagger.BasicAuth{
		UserName: "test@example.com",
		Password: "test-api-key",
	}
	return context.WithValue(context.Background(), swagger.ContextBasicAuth, auth)
}

// mockSoftwareAuthContext returns a context with test authentication for software API
func mockSoftwareAuthContext() context.Context {
	auth := swagger_software.BasicAuth{
		UserName: "test@example.com",
		Password: "test-api-key",
	}
	return context.WithValue(context.Background(), swagger_software.ContextBasicAuth, auth)
}

// JSON response builders for common Jira API responses

// buildLabelsResponse creates a JSON response for the labels API
func buildLabelsResponse(labels []string, isLast bool, startAt int64, maxResults int32, total int64) string {
	response := map[string]interface{}{
		"values":     labels,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildProjectsResponse creates a JSON response for the projects search API
func buildProjectsResponse(projects []map[string]interface{}, isLast bool, startAt int64, maxResults int32, total int64) string {
	response := map[string]interface{}{
		"values":     projects,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildProject creates a project object for testing
func buildProject(key, name, projectTypeKey string) map[string]interface{} {
	return map[string]interface{}{
		"key":            key,
		"name":           name,
		"projectTypeKey": projectTypeKey,
		"id":             "10000",
		"self":           "https://test.atlassian.net/rest/api/3/project/10000",
	}
}

// buildErrorResponse creates a JSON error response
func buildErrorResponse(message string, statusCode int) string {
	response := map[string]interface{}{
		"errorMessages": []string{message},
		"errors":        map[string]interface{}{},
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// mockJSONHandler creates a handler that returns JSON with the given status code
func mockJSONHandler(statusCode int, jsonResponse string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(statusCode)
		w.Write([]byte(jsonResponse))
	}
}

// mockRouteHandler creates a handler that routes requests based on path
func mockRouteHandler(routes map[string]http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		for path, handler := range routes {
			if r.URL.Path == path || matchPathPattern(r.URL.Path, path) {
				handler(w, r)
				return
			}
		}
		// Default 404 response
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"errorMessages": ["Not found"]}`))
	}
}

// matchPathPattern checks if the path matches a pattern with wildcards
// For simple prefix matching in tests
func matchPathPattern(path, pattern string) bool {
	// Simple implementation: check if path starts with pattern (without trailing wildcard)
	if len(pattern) > 0 && pattern[len(pattern)-1] == '*' {
		prefix := pattern[:len(pattern)-1]
		return len(path) >= len(prefix) && path[:len(prefix)] == prefix
	}
	return path == pattern
}

// Issue-related response builders

// buildIssueResponse creates a JSON response for a single issue
func buildIssueResponse(key, summary, status, priority, issueType string) string {
	response := map[string]interface{}{
		"key":  key,
		"id":   "10001",
		"self": "https://test.atlassian.net/rest/api/3/issue/10001",
		"fields": map[string]interface{}{
			"summary": summary,
			"status": map[string]interface{}{
				"name": status,
				"statusCategory": map[string]interface{}{
					"colorName": "blue-gray",
				},
			},
			"priority": map[string]interface{}{
				"name": priority,
			},
			"issuetype": map[string]interface{}{
				"name": issueType,
			},
			"created": "2024-01-15T10:30:00.000+0000",
			"updated": "2024-01-16T14:45:00.000+0000",
		},
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildSearchIssuesResponse creates a JSON response for issue search
func buildSearchIssuesResponse(issues []map[string]interface{}, startAt, maxResults, total int) string {
	response := map[string]interface{}{
		"issues":     issues,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildIssueForSearch creates an issue object suitable for search responses
func buildIssueForSearch(key, summary, status, priority, issueType string) map[string]interface{} {
	return map[string]interface{}{
		"key":  key,
		"id":   "10001",
		"self": "https://test.atlassian.net/rest/api/3/issue/10001",
		"fields": map[string]interface{}{
			"summary": summary,
			"status": map[string]interface{}{
				"name": status,
				"statusCategory": map[string]interface{}{
					"colorName": "blue-gray",
				},
			},
			"priority": map[string]interface{}{
				"name": priority,
			},
			"issuetype": map[string]interface{}{
				"name": issueType,
			},
		},
	}
}

// Sprint-related response builders

// buildSprintsResponse creates a JSON response for sprints
func buildSprintsResponse(sprints []map[string]interface{}, isLast bool, startAt, maxResults int) string {
	response := map[string]interface{}{
		"values":     sprints,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildSprint creates a sprint object for testing
func buildSprint(id int, name, state, startDate, endDate string) map[string]interface{} {
	sprint := map[string]interface{}{
		"id":    id,
		"name":  name,
		"state": state,
	}
	if startDate != "" {
		sprint["startDate"] = startDate
	}
	if endDate != "" {
		sprint["endDate"] = endDate
	}
	return sprint
}

// Comment-related response builders

// buildCommentsResponse creates a JSON response for comments
func buildCommentsResponse(comments []map[string]interface{}, startAt, maxResults, total int) string {
	response := map[string]interface{}{
		"comments":   comments,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildComment creates a comment object for testing
func buildComment(id, body, authorName, created string) map[string]interface{} {
	return map[string]interface{}{
		"id":   id,
		"self": "https://test.atlassian.net/rest/api/3/issue/10001/comment/" + id,
		"body": map[string]interface{}{
			"type":    "doc",
			"version": 1,
			"content": []map[string]interface{}{
				{
					"type": "paragraph",
					"content": []map[string]interface{}{
						{
							"type": "text",
							"text": body,
						},
					},
				},
			},
		},
		"author": map[string]interface{}{
			"displayName": authorName,
		},
		"created": created,
	}
}

// Board-related response builders

// buildBoardsResponse creates a JSON response for boards
func buildBoardsResponse(boards []map[string]interface{}, isLast bool, startAt, maxResults int64) string {
	response := map[string]interface{}{
		"values":     boards,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildBoard creates a board object for testing
func buildBoard(id int64, name, boardType, projectKey string) map[string]interface{} {
	board := map[string]interface{}{
		"id":   id,
		"name": name,
		"type": boardType,
		"self": "https://test.atlassian.net/rest/agile/1.0/board/" + string(rune(id)),
	}
	if projectKey != "" {
		board["location"] = map[string]interface{}{
			"projectKey": projectKey,
		}
	}
	return board
}

// Status-related response builders

// buildStatusesResponse creates a JSON response for statuses search
func buildStatusesResponse(statuses []map[string]interface{}, isLast bool, startAt int64, maxResults int32, total int64) string {
	response := map[string]interface{}{
		"values":     statuses,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildStatus creates a status object for testing
func buildStatus(id, name, category, description string) map[string]interface{} {
	return map[string]interface{}{
		"id":             id,
		"name":           name,
		"statusCategory": category,
		"description":    description,
		"self":           "https://test.atlassian.net/rest/api/3/status/" + id,
	}
}

// Field-related response builders

// buildFieldsResponse creates a JSON response for fields
func buildFieldsResponse(fields []map[string]interface{}) string {
	data, _ := json.Marshal(fields)
	return string(data)
}

// buildField creates a field object for testing
func buildField(id, name string, custom bool) map[string]interface{} {
	return map[string]interface{}{
		"id":     id,
		"name":   name,
		"custom": custom,
	}
}
