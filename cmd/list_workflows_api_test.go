package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// buildWorkflowsResponse creates a JSON response for workflows API
func buildWorkflowsResponse(workflows []map[string]interface{}, isLast bool, startAt int64, maxResults int32, total int64) string {
	response := map[string]interface{}{
		"values":     workflows,
		"isLast":     isLast,
		"startAt":    startAt,
		"maxResults": maxResults,
		"total":      total,
	}
	data, _ := json.Marshal(response)
	return string(data)
}

// buildWorkflow creates a workflow object for testing
func buildWorkflow(name, description string, isDefault bool) map[string]interface{} {
	return map[string]interface{}{
		"id": map[string]interface{}{
			"name":     name,
			"entityId": fmt.Sprintf("workflow-%s", strings.ReplaceAll(strings.ToLower(name), " ", "-")),
		},
		"description": description,
		"isDefault":   isDefault,
	}
}

func TestWorkflowsAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/workflow") {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		workflows := []map[string]interface{}{
			buildWorkflow("jira", "Default Jira Workflow", true),
			buildWorkflow("Scrum Workflow", "Workflow for Scrum projects", false),
			buildWorkflow("Kanban Workflow", "Workflow for Kanban projects", false),
		}
		w.Write([]byte(buildWorkflowsResponse(workflows, true, 0, 50, 3)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 3 {
		t.Errorf("Expected 3 workflows, got %d", len(result.GetValues()))
	}
}

func TestWorkflowsAPI_Pagination(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		startAt := r.URL.Query().Get("startAt")
		requestCount++

		if startAt == "" || startAt == "0" {
			workflows := []map[string]interface{}{
				buildWorkflow("Workflow 1", "First workflow", true),
				buildWorkflow("Workflow 2", "Second workflow", false),
			}
			w.Write([]byte(buildWorkflowsResponse(workflows, false, 0, 2, 3)))
		} else {
			workflows := []map[string]interface{}{
				buildWorkflow("Workflow 3", "Third workflow", false),
			}
			w.Write([]byte(buildWorkflowsResponse(workflows, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// First page
	result1, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).StartAt(0).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("First page failed: %v", err)
	}
	if len(result1.GetValues()) != 2 {
		t.Errorf("Expected 2 workflows on first page, got %d", len(result1.GetValues()))
	}
	if result1.GetIsLast() {
		t.Error("Expected isLast=false on first page")
	}

	// Second page
	result2, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).StartAt(2).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("Second page failed: %v", err)
	}
	if len(result2.GetValues()) != 1 {
		t.Errorf("Expected 1 workflow on second page, got %d", len(result2.GetValues()))
	}
	if !result2.GetIsLast() {
		t.Error("Expected isLast=true on second page")
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests, got %d", requestCount)
	}
}

func TestWorkflowsAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildWorkflowsResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 0 {
		t.Errorf("Expected 0 workflows, got %d", len(result.GetValues()))
	}
}

func TestWorkflowsAPI_WithQueryFilter(t *testing.T) {
	var capturedQuery string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedQuery = r.URL.Query().Get("queryString")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		workflows := []map[string]interface{}{
			buildWorkflow("Scrum Workflow", "Workflow for Scrum projects", false),
		}
		w.Write([]byte(buildWorkflowsResponse(workflows, true, 0, 50, 1)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).QueryString("scrum").Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}

	if capturedQuery != "scrum" {
		t.Errorf("Expected queryString=scrum, got %s", capturedQuery)
	}
}

func TestWorkflowsAPI_WithActiveFilter(t *testing.T) {
	var capturedIsActive string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedIsActive = r.URL.Query().Get("isActive")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildWorkflowsResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).IsActive(true).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}

	if capturedIsActive != "true" {
		t.Errorf("Expected isActive=true, got %s", capturedIsActive)
	}
}

func TestWorkflowsAPI_WithOrderBy(t *testing.T) {
	var capturedOrderBy string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedOrderBy = r.URL.Query().Get("orderBy")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildWorkflowsResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).OrderBy("name").Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}

	if capturedOrderBy != "name" {
		t.Errorf("Expected orderBy=name, got %s", capturedOrderBy)
	}
}

func TestWorkflowsAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestWorkflowsAPI_Forbidden(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(buildErrorResponse("You do not have permission to view workflows", http.StatusForbidden)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err == nil {
		t.Error("Expected error for forbidden request")
	}
}

func TestWorkflowsAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestWorkflowsAPI_WithLongDescription(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		longDesc := strings.Repeat("A very long description. ", 20)
		workflows := []map[string]interface{}{
			buildWorkflow("Complex Workflow", longDesc, false),
		}
		w.Write([]byte(buildWorkflowsResponse(workflows, true, 0, 50, 1)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}
	if len(result.GetValues()) != 1 {
		t.Errorf("Expected 1 workflow, got %d", len(result.GetValues()))
	}
}

func TestWorkflowsAPI_DefaultWorkflow(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		workflows := []map[string]interface{}{
			buildWorkflow("jira", "Default Jira Workflow", true),
			buildWorkflow("Custom Workflow", "A custom workflow", false),
		}
		w.Write([]byte(buildWorkflowsResponse(workflows, true, 0, 50, 2)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}

	defaultCount := 0
	for _, wf := range result.GetValues() {
		if wf.GetIsDefault() {
			defaultCount++
		}
	}
	if defaultCount != 1 {
		t.Errorf("Expected 1 default workflow, got %d", defaultCount)
	}
}

func TestWorkflowsAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		workflows := make([]map[string]interface{}, 50)
		for i := 0; i < 50; i++ {
			isDefault := i == 0
			workflows[i] = buildWorkflow(
				fmt.Sprintf("Workflow %d", i+1),
				fmt.Sprintf("Description for workflow %d", i+1),
				isDefault,
			)
		}
		w.Write([]byte(buildWorkflowsResponse(workflows, true, 0, 50, 50)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).Execute()
	if err != nil {
		t.Fatalf("GetWorkflowsPaginated failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 50 {
		t.Errorf("Expected 50 workflows, got %d", len(result.GetValues()))
	}
}
