package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// buildWorkflowSchemesResponse creates a JSON response for workflow schemes API
func buildWorkflowSchemesResponse(schemes []map[string]interface{}, isLast bool, startAt int64, maxResults int32, total int64) string {
	response := map[string]interface{}{
		"values":     schemes,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildWorkflowScheme creates a workflow scheme object for testing
func buildWorkflowScheme(id int64, name, description, defaultWorkflow string, draft bool) map[string]interface{} {
	return map[string]interface{}{
		"id":              id,
		"name":            name,
		"description":     description,
		"defaultWorkflow": defaultWorkflow,
		"draft":           draft,
		"self":            fmt.Sprintf("https://test.atlassian.net/rest/api/3/workflowscheme/%d", id),
	}
}

func TestWorkflowSchemesAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Handle different API paths
		if strings.Contains(r.URL.Path, "/workflowscheme") {
			if strings.Contains(r.URL.Path, "/project") {
				// Get scheme for projects - return empty
				w.Write([]byte(buildWorkflowSchemesResponse([]map[string]interface{}{}, true, 0, 50, 0)))
			} else {
				// Get all schemes
				schemes := []map[string]interface{}{
					buildWorkflowScheme(10000, "Default Workflow Scheme", "Default scheme", "jira", false),
					buildWorkflowScheme(10001, "Scrum Workflow Scheme", "Scheme for Scrum", "Scrum Workflow", false),
					buildWorkflowScheme(10002, "Kanban Workflow Scheme", "Scheme for Kanban", "Kanban Workflow", false),
				}
				w.Write([]byte(buildWorkflowSchemesResponse(schemes, true, 0, 50, 3)))
			}
		} else if strings.Contains(r.URL.Path, "/project") {
			// Get all projects - return empty
			w.Write([]byte("[]"))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllWorkflowSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 3 {
		t.Errorf("Expected 3 schemes, got %d", len(result.GetValues()))
	}
}

func TestWorkflowSchemesAPI_Pagination(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if !strings.Contains(r.URL.Path, "/workflowscheme") {
			w.Write([]byte("[]"))
			return
		}

		startAt := r.URL.Query().Get("startAt")
		requestCount++

		if startAt == "" || startAt == "0" {
			schemes := []map[string]interface{}{
				buildWorkflowScheme(10000, "Scheme 1", "First scheme", "jira", false),
				buildWorkflowScheme(10001, "Scheme 2", "Second scheme", "Custom Workflow", false),
			}
			w.Write([]byte(buildWorkflowSchemesResponse(schemes, false, 0, 2, 3)))
		} else {
			schemes := []map[string]interface{}{
				buildWorkflowScheme(10002, "Scheme 3", "Third scheme", "Another Workflow", false),
			}
			w.Write([]byte(buildWorkflowSchemesResponse(schemes, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// First page
	result1, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).StartAt(0).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("First page failed: %v", err)
	}
	if len(result1.GetValues()) != 2 {
		t.Errorf("Expected 2 schemes on first page, got %d", len(result1.GetValues()))
	}

	// Second page
	result2, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).StartAt(2).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("Second page failed: %v", err)
	}
	if len(result2.GetValues()) != 1 {
		t.Errorf("Expected 1 scheme on second page, got %d", len(result2.GetValues()))
	}
}

func TestWorkflowSchemesAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildWorkflowSchemesResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllWorkflowSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 0 {
		t.Errorf("Expected 0 schemes, got %d", len(result.GetValues()))
	}
}

func TestWorkflowSchemesAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestWorkflowSchemesAPI_Forbidden(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(buildErrorResponse("Requires Administer Jira permission", http.StatusForbidden)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for forbidden request")
	}
}

func TestWorkflowSchemesAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestWorkflowSchemesAPI_WithDraftScheme(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		schemes := []map[string]interface{}{
			buildWorkflowScheme(10000, "Active Scheme", "Active scheme", "jira", false),
			buildWorkflowScheme(10001, "Draft Scheme", "Draft scheme", "Custom Workflow", true),
		}
		w.Write([]byte(buildWorkflowSchemesResponse(schemes, true, 0, 50, 2)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllWorkflowSchemes failed: %v", err)
	}

	draftCount := 0
	for _, scheme := range result.GetValues() {
		if scheme.GetDraft() {
			draftCount++
		}
	}
	if draftCount != 1 {
		t.Errorf("Expected 1 draft scheme, got %d", draftCount)
	}
}

func TestWorkflowSchemesAPI_VerifyFieldValues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		schemes := []map[string]interface{}{
			{
				"id":              int64(10000),
				"name":            "Test Scheme",
				"description":     "A test workflow scheme",
				"defaultWorkflow": "jira",
				"draft":           false,
				"self":            "https://test.atlassian.net/rest/api/3/workflowscheme/10000",
			},
		}
		w.Write([]byte(buildWorkflowSchemesResponse(schemes, true, 0, 50, 1)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllWorkflowSchemes failed: %v", err)
	}
	if len(result.GetValues()) != 1 {
		t.Fatalf("Expected 1 scheme, got %d", len(result.GetValues()))
	}

	scheme := result.GetValues()[0]
	if scheme.GetName() != "Test Scheme" {
		t.Errorf("Expected name='Test Scheme', got %s", scheme.GetName())
	}
	if scheme.GetDescription() != "A test workflow scheme" {
		t.Errorf("Expected description='A test workflow scheme', got %s", scheme.GetDescription())
	}
	if scheme.GetDefaultWorkflow() != "jira" {
		t.Errorf("Expected defaultWorkflow='jira', got %s", scheme.GetDefaultWorkflow())
	}
	if scheme.GetDraft() {
		t.Error("Expected draft=false")
	}
}

func TestWorkflowSchemesAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		schemes := make([]map[string]interface{}, 50)
		for i := 0; i < 50; i++ {
			draft := i%10 == 0
			schemes[i] = buildWorkflowScheme(
				int64(10000+i),
				fmt.Sprintf("Scheme %d", i+1),
				fmt.Sprintf("Description for scheme %d", i+1),
				"jira",
				draft,
			)
		}
		w.Write([]byte(buildWorkflowSchemesResponse(schemes, true, 0, 50, 50)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowSchemesAPI.GetAllWorkflowSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllWorkflowSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 50 {
		t.Errorf("Expected 50 schemes, got %d", len(result.GetValues()))
	}
}
