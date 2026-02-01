package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetIssueAPI_SingleIssue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issue/PROJ-123" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildIssueResponse("PROJ-123", "Test Issue Summary", "In Progress", "High", "Bug")))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if issue.GetKey() != "PROJ-123" {
		t.Errorf("expected key PROJ-123, got %s", issue.GetKey())
	}

	// Check summary field
	if summary, ok := issue.Fields["summary"].(string); ok {
		if summary != "Test Issue Summary" {
			t.Errorf("expected summary 'Test Issue Summary', got %s", summary)
		}
	}
}

func TestGetIssueAPI_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(buildErrorResponse("Issue not found", http.StatusNotFound)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-999").Execute()
	if err == nil {
		t.Error("expected error for not found issue, got nil")
	}
}

func TestGetIssueAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestGetIssueAPI_WithCustomFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Test Issue",
				"status": map[string]interface{}{
					"name": "Open",
				},
				"customfield_10001": "Custom Value",
				"customfield_10002": map[string]interface{}{
					"value": "Option A",
				},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Verify custom fields are present
	if cf1, ok := issue.Fields["customfield_10001"].(string); !ok || cf1 != "Custom Value" {
		t.Errorf("expected customfield_10001='Custom Value', got %v", issue.Fields["customfield_10001"])
	}
}

func TestGetIssueAPI_WithLinkedIssues(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Test Issue",
				"issuelinks": []interface{}{
					map[string]interface{}{
						"type": map[string]interface{}{
							"inward":  "is blocked by",
							"outward": "blocks",
						},
						"inwardIssue": map[string]interface{}{
							"key": "PROJ-100",
							"fields": map[string]interface{}{
								"summary": "Blocking issue",
							},
						},
					},
				},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test the extractIssueLinks function with the returned data
	links := extractIssueLinks(issue.Fields["issuelinks"], "PROJ-123")
	if len(links) != 1 {
		t.Errorf("expected 1 link, got %d", len(links))
	}
	if len(links) > 0 && links[0].key != "PROJ-100" {
		t.Errorf("expected linked issue PROJ-100, got %s", links[0].key)
	}
}

func TestGetIssueAPI_WithSubtasks(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Parent Issue",
				"subtasks": []interface{}{
					map[string]interface{}{
						"key": "PROJ-124",
						"fields": map[string]interface{}{
							"summary": "Subtask 1",
							"status": map[string]interface{}{
								"name": "Done",
							},
						},
					},
					map[string]interface{}{
						"key": "PROJ-125",
						"fields": map[string]interface{}{
							"summary": "Subtask 2",
							"status": map[string]interface{}{
								"name": "In Progress",
							},
						},
					},
				},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test the extractSubtasks function with the returned data
	subtasks := extractSubtasks(issue.Fields["subtasks"])
	if len(subtasks) != 2 {
		t.Errorf("expected 2 subtasks, got %d", len(subtasks))
	}
	if len(subtasks) > 0 && subtasks[0].key != "PROJ-124" {
		t.Errorf("expected first subtask PROJ-124, got %s", subtasks[0].key)
	}
}

func TestGetIssueAPI_WithDescription(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Test Issue",
				"description": map[string]interface{}{
					"type":    "doc",
					"version": 1,
					"content": []interface{}{
						map[string]interface{}{
							"type": "paragraph",
							"content": []interface{}{
								map[string]interface{}{
									"type": "text",
									"text": "This is the description",
								},
							},
						},
					},
				},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test the extractTextFromADF function with the returned data
	description := extractTextFromADF(issue.Fields["description"])
	if description != "This is the description" {
		t.Errorf("expected description 'This is the description', got %q", description)
	}
}

func TestGetIssueAPI_WithAssigneeAndReporter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Test Issue",
				"assignee": map[string]interface{}{
					"displayName":  "John Doe",
					"emailAddress": "john@example.com",
				},
				"reporter": map[string]interface{}{
					"displayName":  "Jane Smith",
					"emailAddress": "jane@example.com",
				},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test the extractUserDisplayName function with the returned data
	assignee := extractUserDisplayName(issue.Fields["assignee"])
	if assignee != "John Doe" {
		t.Errorf("expected assignee 'John Doe', got %q", assignee)
	}

	reporter := extractUserDisplayName(issue.Fields["reporter"])
	if reporter != "Jane Smith" {
		t.Errorf("expected reporter 'Jane Smith', got %q", reporter)
	}
}

func TestGetIssueAPI_WithLabels(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		issue := map[string]interface{}{
			"key": "PROJ-123",
			"id":  "10001",
			"fields": map[string]interface{}{
				"summary": "Test Issue",
				"labels":  []interface{}{"bug", "critical", "frontend"},
			},
		}
		data, _ := json.Marshal(issue)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	issue, _, err := client.IssuesAPI.GetIssue(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Test the extractLabels function with the returned data
	labels := extractLabels(issue.Fields["labels"])
	if len(labels) != 3 {
		t.Errorf("expected 3 labels, got %d", len(labels))
	}
	expectedLabels := []string{"bug", "critical", "frontend"}
	for i, label := range labels {
		if label != expectedLabels[i] {
			t.Errorf("expected label %d to be %q, got %q", i, expectedLabels[i], label)
		}
	}
}
