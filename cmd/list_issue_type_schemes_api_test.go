package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// buildIssueTypeSchemesResponse creates a JSON response for issue type schemes API
func buildIssueTypeSchemesResponse(schemes []map[string]interface{}, isLast bool, startAt int64, maxResults int32, total int64) string {
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

// buildIssueTypeScheme creates an issue type scheme object for testing
func buildIssueTypeScheme(id, name, description string, isDefault bool) map[string]interface{} {
	return map[string]interface{}{
		"id":          id,
		"name":        name,
		"description": description,
		"isDefault":   isDefault,
	}
}

func TestIssueTypeSchemesAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Handle different API paths
		if strings.Contains(r.URL.Path, "/issuetypescheme") {
			if strings.Contains(r.URL.Path, "/project") {
				// Get scheme for projects - return empty
				w.Write([]byte(buildIssueTypeSchemesResponse([]map[string]interface{}{}, true, 0, 50, 0)))
			} else {
				// Get all schemes
				schemes := []map[string]interface{}{
					buildIssueTypeScheme("10000", "Default Issue Type Scheme", "Default issue type scheme", true),
					buildIssueTypeScheme("10001", "Scrum Issue Type Scheme", "Scheme for Scrum projects", false),
					buildIssueTypeScheme("10002", "Kanban Issue Type Scheme", "Scheme for Kanban projects", false),
				}
				w.Write([]byte(buildIssueTypeSchemesResponse(schemes, true, 0, 50, 3)))
			}
		} else if strings.Contains(r.URL.Path, "/project") {
			// Get all projects - return empty
			w.Write([]byte("[]"))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllIssueTypeSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 3 {
		t.Errorf("Expected 3 schemes, got %d", len(result.GetValues()))
	}
}

func TestIssueTypeSchemesAPI_Pagination(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if !strings.Contains(r.URL.Path, "/issuetypescheme") {
			w.Write([]byte("[]"))
			return
		}

		startAt := r.URL.Query().Get("startAt")
		requestCount++

		if startAt == "" || startAt == "0" {
			schemes := []map[string]interface{}{
				buildIssueTypeScheme("10000", "Scheme 1", "First scheme", true),
				buildIssueTypeScheme("10001", "Scheme 2", "Second scheme", false),
			}
			w.Write([]byte(buildIssueTypeSchemesResponse(schemes, false, 0, 2, 3)))
		} else {
			schemes := []map[string]interface{}{
				buildIssueTypeScheme("10002", "Scheme 3", "Third scheme", false),
			}
			w.Write([]byte(buildIssueTypeSchemesResponse(schemes, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// First page
	result1, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).StartAt(0).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("First page failed: %v", err)
	}
	if len(result1.GetValues()) != 2 {
		t.Errorf("Expected 2 schemes on first page, got %d", len(result1.GetValues()))
	}

	// Second page
	result2, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).StartAt(2).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("Second page failed: %v", err)
	}
	if len(result2.GetValues()) != 1 {
		t.Errorf("Expected 1 scheme on second page, got %d", len(result2.GetValues()))
	}
}

func TestIssueTypeSchemesAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildIssueTypeSchemesResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllIssueTypeSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 0 {
		t.Errorf("Expected 0 schemes, got %d", len(result.GetValues()))
	}
}

func TestIssueTypeSchemesAPI_WithOrderBy(t *testing.T) {
	var capturedOrderBy string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedOrderBy = r.URL.Query().Get("orderBy")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildIssueTypeSchemesResponse([]map[string]interface{}{}, true, 0, 50, 0)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).OrderBy("name").Execute()
	if err != nil {
		t.Fatalf("GetAllIssueTypeSchemes failed: %v", err)
	}

	if capturedOrderBy != "name" {
		t.Errorf("Expected orderBy=name, got %s", capturedOrderBy)
	}
}

func TestIssueTypeSchemesAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestIssueTypeSchemesAPI_Forbidden(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(buildErrorResponse("Requires Administer Jira permission", http.StatusForbidden)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for forbidden request")
	}
}

func TestIssueTypeSchemesAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestIssueTypeSchemesAPI_DefaultScheme(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		schemes := []map[string]interface{}{
			buildIssueTypeScheme("10000", "Default Scheme", "The default scheme", true),
			buildIssueTypeScheme("10001", "Custom Scheme 1", "Custom scheme", false),
			buildIssueTypeScheme("10002", "Custom Scheme 2", "Another custom scheme", false),
		}
		w.Write([]byte(buildIssueTypeSchemesResponse(schemes, true, 0, 50, 3)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllIssueTypeSchemes failed: %v", err)
	}

	defaultCount := 0
	for _, scheme := range result.GetValues() {
		if scheme.GetIsDefault() {
			defaultCount++
		}
	}
	if defaultCount != 1 {
		t.Errorf("Expected 1 default scheme, got %d", defaultCount)
	}
}

func TestIssueTypeSchemesAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		schemes := make([]map[string]interface{}, 50)
		for i := 0; i < 50; i++ {
			isDefault := i == 0
			schemes[i] = buildIssueTypeScheme(
				fmt.Sprintf("%d", 10000+i),
				fmt.Sprintf("Scheme %d", i+1),
				fmt.Sprintf("Description for scheme %d", i+1),
				isDefault,
			)
		}
		w.Write([]byte(buildIssueTypeSchemesResponse(schemes, true, 0, 50, 50)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllIssueTypeSchemes failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result.GetValues()) != 50 {
		t.Errorf("Expected 50 schemes, got %d", len(result.GetValues()))
	}
}
