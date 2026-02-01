package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestParseProjectId(t *testing.T) {
	tests := []struct {
		name        string
		id          string
		expected    int64
		expectError bool
	}{
		{
			name:        "valid numeric ID",
			id:          "10001",
			expected:    10001,
			expectError: false,
		},
		{
			name:        "valid large ID",
			id:          "123456789",
			expected:    123456789,
			expectError: false,
		},
		{
			name:        "valid single digit ID",
			id:          "5",
			expected:    5,
			expectError: false,
		},
		{
			name:        "zero ID",
			id:          "0",
			expected:    0,
			expectError: false,
		},
		{
			name:        "invalid - non-numeric",
			id:          "abc",
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid - empty string",
			id:          "",
			expected:    0,
			expectError: true,
		},
		{
			name:        "invalid - mixed alphanumeric",
			id:          "123abc",
			expected:    123,
			expectError: false, // Sscanf will parse the numeric prefix
		},
		{
			name:        "invalid - special characters",
			id:          "!@#",
			expected:    0,
			expectError: true,
		},
		{
			name:        "negative number",
			id:          "-100",
			expected:    -100,
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parseProjectId(tt.id)
			if tt.expectError {
				if err == nil {
					t.Errorf("parseProjectId(%q) expected error but got nil", tt.id)
				}
				return
			}
			if err != nil {
				t.Errorf("parseProjectId(%q) unexpected error: %v", tt.id, err)
				return
			}
			if result != tt.expected {
				t.Errorf("parseProjectId(%q) = %d, want %d", tt.id, result, tt.expected)
			}
		})
	}
}

// buildIssueTypesResponse creates a JSON response for issue types API (returns array directly)
func buildIssueTypesResponse(issueTypes []map[string]interface{}) string {
	data, _ := json.Marshal(issueTypes)
	return string(data)
}

// buildIssueType creates an issue type object for testing
func buildIssueType(id, name, description string, subtask bool, projectKey string) map[string]interface{} {
	issueType := map[string]interface{}{
		"id":          id,
		"name":        name,
		"description": description,
		"subtask":     subtask,
		"self":        "https://test.atlassian.net/rest/api/3/issuetype/" + id,
	}
	if projectKey != "" {
		issueType["scope"] = map[string]interface{}{
			"project": map[string]interface{}{
				"key": projectKey,
			},
		}
	}
	return issueType
}

func TestIssueTypesAPI_GetAllTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.HasSuffix(r.URL.Path, "/issuetype") {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		issueTypes := []map[string]interface{}{
			buildIssueType("10001", "Bug", "A problem or error", false, ""),
			buildIssueType("10002", "Story", "A user story", false, ""),
			buildIssueType("10003", "Task", "A task that needs to be done", false, ""),
			buildIssueType("10004", "Sub-task", "A sub-task of another issue", true, ""),
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 4 {
		t.Errorf("Expected 4 issue types, got %d", len(result))
	}
}

func TestIssueTypesAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildIssueTypesResponse([]map[string]interface{}{})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 issue types, got %d", len(result))
	}
}

func TestIssueTypesAPI_WithSubtasks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		issueTypes := []map[string]interface{}{
			buildIssueType("10001", "Story", "A user story", false, ""),
			buildIssueType("10002", "Sub-task", "A sub-task", true, ""),
			buildIssueType("10003", "Technical Task", "A technical sub-task", true, ""),
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}

	subtaskCount := 0
	for _, it := range result {
		if it.GetSubtask() {
			subtaskCount++
		}
	}
	if subtaskCount != 2 {
		t.Errorf("Expected 2 subtask issue types, got %d", subtaskCount)
	}
}

func TestIssueTypesAPI_WithProjectScope(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		issueTypes := []map[string]interface{}{
			buildIssueType("10001", "Global Bug", "A global bug type", false, ""),
			buildIssueType("10002", "Project Bug", "A project-specific bug type", false, "PROJ"),
			buildIssueType("10003", "Another Project Bug", "Another project-specific type", false, "OTHER"),
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 issue types, got %d", len(result))
	}
}

func TestIssueTypesAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestIssueTypesAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestIssueTypesAPI_GetTypesForProject(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify projectId is passed
		projectId := r.URL.Query().Get("projectId")
		if projectId == "" {
			t.Error("Expected projectId query parameter")
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		issueTypes := []map[string]interface{}{
			buildIssueType("10001", "Bug", "A bug for this project", false, "PROJ"),
			buildIssueType("10002", "Story", "A story for this project", false, "PROJ"),
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypesAPI.GetIssueTypesForProject(ctx).ProjectId(10001).Execute()
	if err != nil {
		t.Fatalf("GetIssueTypesForProject failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 2 {
		t.Errorf("Expected 2 issue types, got %d", len(result))
	}
}

func TestIssueTypesAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issueTypes := make([]map[string]interface{}, 50)
		for i := 0; i < 50; i++ {
			subtask := i%5 == 0
			projectKey := ""
			if i%3 == 0 {
				projectKey = fmt.Sprintf("PROJ%d", i%10)
			}
			issueTypes[i] = buildIssueType(
				fmt.Sprintf("%d", 10000+i),
				fmt.Sprintf("Issue Type %d", i+1),
				fmt.Sprintf("Description for issue type %d", i+1),
				subtask,
				projectKey,
			)
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 50 {
		t.Errorf("Expected 50 issue types, got %d", len(result))
	}
}

func TestIssueTypesAPI_VerifyFieldValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		issueTypes := []map[string]interface{}{
			{
				"id":          "10001",
				"name":        "Epic",
				"description": "A large user story",
				"subtask":     false,
				"self":        "https://test.atlassian.net/rest/api/3/issuetype/10001",
				"iconUrl":     "https://test.atlassian.net/secure/viewavatar?size=medium&avatarId=10307&avatarType=issuetype",
			},
		}
		w.Write([]byte(buildIssueTypesResponse(issueTypes)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueTypesAPI.GetIssueAllTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetIssueAllTypes failed: %v", err)
	}
	if len(result) != 1 {
		t.Fatalf("Expected 1 issue type, got %d", len(result))
	}

	issueType := result[0]
	if issueType.GetId() != "10001" {
		t.Errorf("Expected id=10001, got %s", issueType.GetId())
	}
	if issueType.GetName() != "Epic" {
		t.Errorf("Expected name=Epic, got %s", issueType.GetName())
	}
	if issueType.GetDescription() != "A large user story" {
		t.Errorf("Expected description='A large user story', got %s", issueType.GetDescription())
	}
	if issueType.GetSubtask() {
		t.Error("Expected subtask=false")
	}
}
