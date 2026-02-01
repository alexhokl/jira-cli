package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunListLabels_SinglePage(t *testing.T) {
	// Create mock server that returns a single page of labels
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request path
		if r.URL.Path != "/rest/api/3/label" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildLabelsResponse(
			[]string{"bug", "enhancement", "documentation"},
			true, // isLast
			0,    // startAt
			1000, // maxResults
			3,    // total
		)
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create client pointing to mock server
	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// Call the API directly (testing the API call pattern)
	result, _, err := client.LabelsAPI.GetAllLabels(ctx).
		StartAt(0).
		MaxResults(1000).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	labels := result.GetValues()
	if len(labels) != 3 {
		t.Errorf("expected 3 labels, got %d", len(labels))
	}

	expectedLabels := []string{"bug", "enhancement", "documentation"}
	for i, label := range labels {
		if label != expectedLabels[i] {
			t.Errorf("expected label %d to be %q, got %q", i, expectedLabels[i], label)
		}
	}

	if !result.GetIsLast() {
		t.Error("expected isLast to be true")
	}
}

func TestRunListLabels_MultiplePagesWithPagination(t *testing.T) {
	// Track request count to return different pages
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/label" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			// First page
			response = buildLabelsResponse(
				[]string{"label1", "label2"},
				false, // not last page
				0,
				2,
				4,
			)
		} else {
			// Second page (last)
			response = buildLabelsResponse(
				[]string{"label3", "label4"},
				true, // last page
				2,
				2,
				4,
			)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// Simulate the pagination loop from runListLabels
	var allLabels []string
	var startAt int64 = 0
	const pageSize int32 = 2

	for {
		result, _, err := client.LabelsAPI.GetAllLabels(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		allLabels = append(allLabels, result.GetValues()...)

		if result.GetIsLast() {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allLabels) != 4 {
		t.Errorf("expected 4 labels, got %d", len(allLabels))
	}

	expectedLabels := []string{"label1", "label2", "label3", "label4"}
	for i, label := range allLabels {
		if label != expectedLabels[i] {
			t.Errorf("expected label %d to be %q, got %q", i, expectedLabels[i], label)
		}
	}

	if requestCount != 2 {
		t.Errorf("expected 2 requests, got %d", requestCount)
	}
}

func TestRunListLabels_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildLabelsResponse([]string{}, true, 0, 1000, 0)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.LabelsAPI.GetAllLabels(ctx).
		StartAt(0).
		MaxResults(1000).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	labels := result.GetValues()
	if len(labels) != 0 {
		t.Errorf("expected 0 labels, got %d", len(labels))
	}
}

func TestRunListLabels_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.LabelsAPI.GetAllLabels(ctx).
		StartAt(0).
		MaxResults(1000).
		Execute()

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestRunListLabels_MaxResultsLimit(t *testing.T) {
	// Test the max results limiting logic
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			response = buildLabelsResponse(
				[]string{"label1", "label2", "label3"},
				false,
				0,
				3,
				10,
			)
		} else {
			response = buildLabelsResponse(
				[]string{"label4", "label5", "label6"},
				false,
				3,
				3,
				10,
			)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// Simulate the logic from runListLabels with maxResults limit
	var allLabels []string
	var startAt int64 = 0
	const pageSize int32 = 3
	maxResults := int32(5) // User wants only 5 labels

	for {
		result, _, err := client.LabelsAPI.GetAllLabels(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		allLabels = append(allLabels, result.GetValues()...)

		// Apply user's max limit
		if maxResults > 0 && int32(len(allLabels)) >= maxResults {
			allLabels = allLabels[:maxResults]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allLabels) != 5 {
		t.Errorf("expected 5 labels, got %d", len(allLabels))
	}
}

func TestRunListLabels_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.LabelsAPI.GetAllLabels(ctx).
		StartAt(0).
		MaxResults(1000).
		Execute()

	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestRunListLabels_LargeDataset(t *testing.T) {
	// Generate a large number of labels across multiple pages
	totalLabels := 2500
	pageSize := 1000
	currentPage := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		start := currentPage * pageSize
		end := start + pageSize
		if end > totalLabels {
			end = totalLabels
		}

		labels := make([]string, end-start)
		for i := range labels {
			labels[i] = "label" + string(rune('0'+((start+i)%10)))
		}

		isLast := end >= totalLabels
		response := buildLabelsResponse(labels, isLast, int64(start), int32(pageSize), int64(totalLabels))
		currentPage++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	var allLabels []string
	var startAt int64 = 0

	for {
		result, _, err := client.LabelsAPI.GetAllLabels(ctx).
			StartAt(startAt).
			MaxResults(int32(pageSize)).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		allLabels = append(allLabels, result.GetValues()...)

		if result.GetIsLast() {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allLabels) != totalLabels {
		t.Errorf("expected %d labels, got %d", totalLabels, len(allLabels))
	}
}

func TestRunListLabels_RequestVerification(t *testing.T) {
	// Verify the correct query parameters are sent
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify HTTP method
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}

		// Verify path
		if r.URL.Path != "/rest/api/3/label" {
			t.Errorf("expected path /rest/api/3/label, got %s", r.URL.Path)
		}

		// Verify query parameters
		startAt := r.URL.Query().Get("startAt")
		if startAt != "0" {
			t.Errorf("expected startAt=0, got startAt=%s", startAt)
		}

		maxResults := r.URL.Query().Get("maxResults")
		if maxResults != "100" {
			t.Errorf("expected maxResults=100, got maxResults=%s", maxResults)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildLabelsResponse([]string{"test"}, true, 0, 100, 1)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.LabelsAPI.GetAllLabels(ctx).
		StartAt(0).
		MaxResults(100).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
