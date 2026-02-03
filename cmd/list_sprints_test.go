package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestFormatDateString(t *testing.T) {
	// Get local timezone for expected results
	loc := time.Local

	tests := []struct {
		name     string
		input    string
		expected string
		isLocal  bool // if true, expected is in local timezone format
	}{
		{
			name:     "RFC3339 format",
			input:    "2024-01-15T10:30:00Z",
			expected: "2024-01-15",
			isLocal:  true,
		},
		{
			name:     "RFC3339 with timezone offset",
			input:    "2024-06-20T14:00:00+05:30",
			expected: "2024-06-20",
			isLocal:  true,
		},
		{
			name:     "RFC3339Nano format",
			input:    "2024-03-10T08:15:30.123456789Z",
			expected: "2024-03-10",
			isLocal:  true,
		},
		{
			name:     "Jira format with milliseconds and Z",
			input:    "2024-12-25T00:00:00.000Z",
			expected: "2024-12-25",
			isLocal:  true,
		},
		{
			name:     "Jira format with milliseconds and offset",
			input:    "2024-07-04T12:00:00.000-0700",
			expected: "2024-07-04",
			isLocal:  true,
		},
		{
			name:     "Simple ISO format with Z",
			input:    "2024-09-01T00:00:00Z",
			expected: "2024-09-01",
			isLocal:  true,
		},
		{
			name:     "Date only format",
			input:    "2024-02-29",
			expected: "2024-02-29",
			isLocal:  false, // date-only doesn't have timezone conversion
		},
		{
			name:     "Invalid format returns original",
			input:    "not-a-date",
			expected: "not-a-date",
			isLocal:  false,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: "",
			isLocal:  false,
		},
		{
			name:     "Partial date format",
			input:    "2024-01",
			expected: "2024-01",
			isLocal:  false,
		},
		{
			name:     "Unix timestamp format (not supported)",
			input:    "1705312200",
			expected: "1705312200",
			isLocal:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDateString(tt.input)

			if tt.isLocal && tt.input != "" {
				// For timezone-aware inputs, we need to account for local timezone
				// The result should be a valid date format
				if !isValidDateFormat(result) {
					t.Errorf("formatDateString(%q) = %q, expected valid date format", tt.input, result)
				}
			} else {
				if result != tt.expected {
					t.Errorf("formatDateString(%q) = %q, want %q", tt.input, result, tt.expected)
				}
			}
		})
	}

	// Test that UTC dates are converted to local timezone correctly
	t.Run("UTC to local conversion", func(t *testing.T) {
		// Use a fixed UTC time
		input := "2024-01-15T00:00:00Z"
		result := formatDateString(input)

		// Parse the input as UTC
		utcTime, _ := time.Parse(time.RFC3339, input)
		localTime := utcTime.In(loc)
		expectedDate := localTime.Format("2006-01-02")

		if result != expectedDate {
			t.Errorf("formatDateString(%q) = %q, want %q (local timezone)", input, result, expectedDate)
		}
	})
}

// isValidDateFormat checks if a string is in yyyy-MM-dd format
func isValidDateFormat(s string) bool {
	if len(s) != 10 {
		return false
	}
	_, err := time.Parse("2006-01-02", s)
	return err == nil
}

func TestFormatDateStringEdgeCases(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "date with extra whitespace",
			input: " 2024-01-15 ",
		},
		{
			name:  "mixed case timezone",
			input: "2024-01-15T10:30:00z", // lowercase z
		},
		{
			name:  "date with time but no timezone",
			input: "2024-01-15T10:30:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := formatDateString(tt.input)
			// These should either be parsed or return original
			// Just ensure no panic
			if result == "" && tt.input != "" {
				// If input was trimmed to empty, this is unexpected
				if strings.TrimSpace(tt.input) != "" {
					t.Logf("formatDateString(%q) returned empty string", tt.input)
				}
			}
		})
	}
}

func TestSprintStateFiltering(t *testing.T) {
	// Test the state filter logic used in runListSprints
	tests := []struct {
		name        string
		stateFilter []string
		sprintState string
		shouldMatch bool
	}{
		{
			name:        "empty filter matches all",
			stateFilter: nil,
			sprintState: "active",
			shouldMatch: true,
		},
		{
			name:        "active filter matches active",
			stateFilter: []string{"active"},
			sprintState: "active",
			shouldMatch: true,
		},
		{
			name:        "active filter does not match closed",
			stateFilter: []string{"active"},
			sprintState: "closed",
			shouldMatch: false,
		},
		{
			name:        "closed filter matches closed",
			stateFilter: []string{"closed"},
			sprintState: "closed",
			shouldMatch: true,
		},
		{
			name:        "future filter matches future",
			stateFilter: []string{"future"},
			sprintState: "future",
			shouldMatch: true,
		},
		{
			name:        "multiple state options matches first",
			stateFilter: []string{"active", "closed"},
			sprintState: "active",
			shouldMatch: true,
		},
		{
			name:        "multiple state options matches second",
			stateFilter: []string{"active", "closed"},
			sprintState: "closed",
			shouldMatch: true,
		},
		{
			name:        "multiple state options does not match other",
			stateFilter: []string{"active", "closed"},
			sprintState: "future",
			shouldMatch: false,
		},
		{
			name:        "all three states matches any",
			stateFilter: []string{"active", "closed", "future"},
			sprintState: "future",
			shouldMatch: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Build state filter map (same logic as in runListSprints)
			stateFilterMap := make(map[string]bool)
			for _, s := range tt.stateFilter {
				stateFilterMap[strings.TrimSpace(s)] = true
			}

			// Check if sprint matches filter
			matches := len(stateFilterMap) == 0 || stateFilterMap[tt.sprintState]

			if matches != tt.shouldMatch {
				t.Errorf("state filter %v with sprint state %q: got match=%v, want %v",
					tt.stateFilter, tt.sprintState, matches, tt.shouldMatch)
			}
		})
	}
}

// API Mock Tests for Sprints

func TestSprintsAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.HasPrefix(r.URL.Path, "/rest/agile/1.0/board/") || !strings.HasSuffix(r.URL.Path, "/sprint") {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		sprints := []map[string]interface{}{
			buildSprint(1, "Sprint 1", "active", "2024-01-01T00:00:00.000Z", "2024-01-14T00:00:00.000Z"),
			buildSprint(2, "Sprint 2", "future", "", ""),
			buildSprint(3, "Sprint 3", "closed", "2023-12-01T00:00:00.000Z", "2023-12-14T00:00:00.000Z"),
		}
		w.Write([]byte(buildSprintsResponse(sprints, true, 0, 50)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	resp, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err != nil {
		t.Fatalf("GetAllSprints failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestSprintsAPI_Pagination(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		startAt := r.URL.Query().Get("startAt")
		requestCount++

		if startAt == "" || startAt == "0" {
			sprints := []map[string]interface{}{
				buildSprint(1, "Sprint 1", "active", "", ""),
				buildSprint(2, "Sprint 2", "future", "", ""),
			}
			w.Write([]byte(buildSprintsResponse(sprints, false, 0, 2)))
		} else {
			sprints := []map[string]interface{}{
				buildSprint(3, "Sprint 3", "closed", "", ""),
			}
			w.Write([]byte(buildSprintsResponse(sprints, true, 2, 2)))
		}
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	// First page
	resp, err := client.BoardAPI.GetAllSprints(ctx, 100).StartAt(0).Execute()
	if err != nil {
		t.Fatalf("First page failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	// Second page
	resp, err = client.BoardAPI.GetAllSprints(ctx, 100).StartAt(2).Execute()
	if err != nil {
		t.Fatalf("Second page failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests, got %d", requestCount)
	}
}

func TestSprintsAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildSprintsResponse([]map[string]interface{}{}, true, 0, 50)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	resp, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err != nil {
		t.Fatalf("GetAllSprints failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestSprintsAPI_BoardNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"errorMessages": ["Board does not exist or you do not have permission to see it"]}`))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, err := client.BoardAPI.GetAllSprints(ctx, 99999).Execute()
	if err == nil {
		t.Error("Expected error for non-existent board")
	}
}

func TestSprintsAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestSprintsAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestSprintsAPI_WithDates(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		sprints := []map[string]interface{}{
			{
				"id":            1,
				"name":          "Sprint with full dates",
				"state":         "active",
				"startDate":     "2024-01-15T09:00:00.000Z",
				"endDate":       "2024-01-29T17:00:00.000Z",
				"completeDate":  "2024-01-28T16:30:00.000Z",
				"originBoardId": 100,
				"goal":          "Complete user authentication",
			},
		}
		w.Write([]byte(buildSprintsResponse(sprints, true, 0, 50)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	resp, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err != nil {
		t.Fatalf("GetAllSprints failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}

func TestSprintsAPI_VerifyRequestParameters(t *testing.T) {
	var capturedStartAt string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedStartAt = r.URL.Query().Get("startAt")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildSprintsResponse([]map[string]interface{}{}, true, 0, 50)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, err := client.BoardAPI.GetAllSprints(ctx, 100).StartAt(25).Execute()
	if err != nil {
		t.Fatalf("GetAllSprints failed: %v", err)
	}

	if capturedStartAt != "25" {
		t.Errorf("Expected startAt=25, got %s", capturedStartAt)
	}
}

func TestSprintsAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		sprints := make([]map[string]interface{}, 100)
		for i := 0; i < 100; i++ {
			state := "closed"
			if i == 98 {
				state = "active"
			} else if i == 99 {
				state = "future"
			}
			sprints[i] = buildSprint(i+1, fmt.Sprintf("Sprint %d", i+1), state, "", "")
		}
		w.Write([]byte(buildSprintsResponse(sprints, true, 0, 100)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	resp, err := client.BoardAPI.GetAllSprints(ctx, 100).Execute()
	if err != nil {
		t.Fatalf("GetAllSprints failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
}
