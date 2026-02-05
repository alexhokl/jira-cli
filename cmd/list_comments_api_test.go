package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListCommentsAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issue/PROJ-123/comment" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		comments := []map[string]interface{}{
			buildComment("1", "First comment", "John Doe", "2024-01-15T10:30:00.000+0000"),
			buildComment("2", "Second comment", "Jane Smith", "2024-01-16T14:45:00.000+0000"),
		}
		response := buildCommentsResponse(comments, 0, 50, 2)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	comments := result.GetComments()
	if len(comments) != 2 {
		t.Errorf("expected 2 comments, got %d", len(comments))
	}
}

func TestListCommentsAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildCommentsResponse([]map[string]interface{}{}, 0, 50, 0)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	comments := result.GetComments()
	if len(comments) != 0 {
		t.Errorf("expected 0 comments, got %d", len(comments))
	}
}

func TestListCommentsAPI_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(buildErrorResponse("Issue not found", http.StatusNotFound)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-999").Execute()
	if err == nil {
		t.Error("expected error for not found issue, got nil")
	}
}

func TestListCommentsAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestListCommentsAPI_WithADFBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"comments": []interface{}{
				map[string]interface{}{
					"id":   "1",
					"self": "https://test.atlassian.net/rest/api/3/issue/10001/comment/1",
					"body": map[string]interface{}{
						"type":    "doc",
						"version": 1,
						"content": []interface{}{
							map[string]interface{}{
								"type": "paragraph",
								"content": []interface{}{
									map[string]interface{}{
										"type": "text",
										"text": "This is a comment with ADF format",
									},
								},
							},
						},
					},
					"author": map[string]interface{}{
						"displayName": "Test User",
					},
					"created": "2024-01-15T10:30:00.000+00:00",
				},
			},
			"startAt":    0,
			"maxResults": 50,
			"total":      1,
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	comments := result.GetComments()
	if len(comments) != 1 {
		t.Fatalf("expected 1 comment, got %d", len(comments))
	}

	// Test converting ADF body to markdown
	body := convertADFToMarkdown(comments[0].Body)
	expected := "This is a comment with ADF format\n"
	if body != expected {
		t.Errorf("expected body text %q, got %q", expected, body)
	}
}

func TestListCommentsAPI_MultiplePages(t *testing.T) {
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			comments := []map[string]interface{}{
				buildComment("1", "Comment 1", "User 1", "2024-01-15T10:30:00.000+0000"),
				buildComment("2", "Comment 2", "User 2", "2024-01-15T11:30:00.000+0000"),
			}
			response = buildCommentsResponse(comments, 0, 2, 4)
		} else {
			comments := []map[string]interface{}{
				buildComment("3", "Comment 3", "User 3", "2024-01-15T12:30:00.000+0000"),
				buildComment("4", "Comment 4", "User 4", "2024-01-15T13:30:00.000+0000"),
			}
			response = buildCommentsResponse(comments, 2, 2, 4)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// First request
	result, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	comments := result.GetComments()
	if len(comments) != 2 {
		t.Errorf("expected 2 comments in first page, got %d", len(comments))
	}

	if result.GetTotal() != 4 {
		t.Errorf("expected total 4, got %d", result.GetTotal())
	}
}

func TestListCommentsAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComments(ctx, "PROJ-123").Execute()
	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}
