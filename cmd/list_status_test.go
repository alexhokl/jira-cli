package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListStatusAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/statuses/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		statuses := []map[string]interface{}{
			buildStatus("1", "To Do", "TODO", "Initial status"),
			buildStatus("2", "In Progress", "IN_PROGRESS", "Work in progress"),
			buildStatus("3", "Done", "DONE", "Completed"),
		}
		response := buildStatusesResponse(statuses, true, 0, 50, 3)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.StatusAPI.Search(ctx).StartAt(0).MaxResults(50).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statuses := result.GetValues()
	if len(statuses) != 3 {
		t.Errorf("expected 3 statuses, got %d", len(statuses))
	}

	// Verify first status
	if statuses[0].GetName() != "To Do" {
		t.Errorf("expected first status name 'To Do', got %s", statuses[0].GetName())
	}
	if statuses[0].GetStatusCategory() != "TODO" {
		t.Errorf("expected first status category 'TODO', got %s", statuses[0].GetStatusCategory())
	}
}

func TestListStatusAPI_WithQuery(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameter
		searchString := r.URL.Query().Get("searchString")
		if searchString != "progress" {
			t.Errorf("expected searchString=progress, got searchString=%s", searchString)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		statuses := []map[string]interface{}{
			buildStatus("2", "In Progress", "IN_PROGRESS", "Work in progress"),
		}
		response := buildStatusesResponse(statuses, true, 0, 50, 1)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.StatusAPI.Search(ctx).SearchString("progress").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statuses := result.GetValues()
	if len(statuses) != 1 {
		t.Errorf("expected 1 status, got %d", len(statuses))
	}
}

func TestListStatusAPI_WithCategoryFilter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify category parameter
		category := r.URL.Query().Get("statusCategory")
		if category != "IN_PROGRESS" {
			t.Errorf("expected statusCategory=IN_PROGRESS, got statusCategory=%s", category)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		statuses := []map[string]interface{}{
			buildStatus("2", "In Progress", "IN_PROGRESS", "Work in progress"),
			buildStatus("4", "In Review", "IN_PROGRESS", "Under review"),
		}
		response := buildStatusesResponse(statuses, true, 0, 50, 2)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.StatusAPI.Search(ctx).StatusCategory("IN_PROGRESS").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statuses := result.GetValues()
	if len(statuses) != 2 {
		t.Errorf("expected 2 statuses, got %d", len(statuses))
	}

	// Verify all have IN_PROGRESS category
	for _, s := range statuses {
		if s.GetStatusCategory() != "IN_PROGRESS" {
			t.Errorf("expected category 'IN_PROGRESS', got %s", s.GetStatusCategory())
		}
	}
}

func TestListStatusAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildStatusesResponse([]map[string]interface{}{}, true, 0, 50, 0)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.StatusAPI.Search(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statuses := result.GetValues()
	if len(statuses) != 0 {
		t.Errorf("expected 0 statuses, got %d", len(statuses))
	}
}

func TestListStatusAPI_Pagination(t *testing.T) {
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			statuses := []map[string]interface{}{
				buildStatus("1", "To Do", "TODO", ""),
				buildStatus("2", "In Progress", "IN_PROGRESS", ""),
			}
			response = buildStatusesResponse(statuses, false, 0, 2, 4)
		} else {
			statuses := []map[string]interface{}{
				buildStatus("3", "Done", "DONE", ""),
				buildStatus("4", "Closed", "DONE", ""),
			}
			response = buildStatusesResponse(statuses, true, 2, 2, 4)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// Simulate pagination loop
	var allStatuses []string
	var startAt int64 = 0

	for {
		result, _, err := client.StatusAPI.Search(ctx).StartAt(startAt).MaxResults(2).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, s := range result.GetValues() {
			allStatuses = append(allStatuses, s.GetName())
		}

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allStatuses) != 4 {
		t.Errorf("expected 4 statuses, got %d", len(allStatuses))
	}

	if requestCount != 2 {
		t.Errorf("expected 2 requests, got %d", requestCount)
	}
}

func TestListStatusAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.StatusAPI.Search(ctx).Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestListStatusAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.StatusAPI.Search(ctx).Execute()
	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestListStatusAPI_WithScope(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		statuses := []map[string]interface{}{
			{
				"id":             "1",
				"name":           "Project Status",
				"statusCategory": "IN_PROGRESS",
				"description":    "Project-scoped status",
				"scope": map[string]interface{}{
					"type": "PROJECT",
					"project": map[string]interface{}{
						"id": "10000",
					},
				},
			},
		}
		response := buildStatusesResponse(statuses, true, 0, 50, 1)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.StatusAPI.Search(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	statuses := result.GetValues()
	if len(statuses) != 1 {
		t.Fatalf("expected 1 status, got %d", len(statuses))
	}

	// Verify scope exists
	status := statuses[0]
	if status.Scope == nil {
		t.Error("expected status to have scope")
	}
}
