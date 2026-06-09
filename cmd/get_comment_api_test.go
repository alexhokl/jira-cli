package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/cobra"
)

func TestGetCommentAPI_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issue/PROJ-123/comment/10001" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildSingleCommentResponse("10001", "Hello world", "John Doe", "2024-01-15T10:30:00.000+0000")))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	comment, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if comment.GetId() != "10001" {
		t.Errorf("expected comment ID '10001', got %q", comment.GetId())
	}

	author := comment.GetAuthor()
	if author.GetDisplayName() != "John Doe" {
		t.Errorf("expected author 'John Doe', got %q", author.GetDisplayName())
	}

	body := convertADFToMarkdown(comment.Body)
	expected := "Hello world\n"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}

func TestGetCommentAPI_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(buildErrorResponse("Comment not found", http.StatusNotFound)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "99999").Execute()
	if err == nil {
		t.Error("expected error for not found comment, got nil")
	}
}

func TestGetCommentAPI_IssueNotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(buildErrorResponse("Issue not found", http.StatusNotFound)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-999", "10001").Execute()
	if err == nil {
		t.Error("expected error for not found issue, got nil")
	}
}

func TestGetCommentAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestGetCommentAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestGetCommentAPI_WithADFBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id":   "10001",
			"self": "https://test.atlassian.net/rest/api/3/issue/PROJ-123/comment/10001",
			"body": map[string]interface{}{
				"type":    "doc",
				"version": 1,
				"content": []interface{}{
					map[string]interface{}{
						"type": "paragraph",
						"content": []interface{}{
							map[string]interface{}{
								"type": "text",
								"text": "This is a ",
							},
							map[string]interface{}{
								"type": "text",
								"text": "bold",
								"marks": []interface{}{
									map[string]interface{}{"type": "strong"},
								},
							},
							map[string]interface{}{
								"type": "text",
								"text": " comment",
							},
						},
					},
				},
			},
			"author": map[string]interface{}{
				"displayName": "Test User",
			},
			"created": "2024-01-15T10:30:00.000+00:00",
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	comment, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	body := convertADFToMarkdown(comment.Body)
	expected := "This is a **bold** comment\n"
	if body != expected {
		t.Errorf("expected body %q, got %q", expected, body)
	}
}

func TestGetCommentAPI_AuthorFallbackToEmail(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id":   "10001",
			"self": "https://test.atlassian.net/rest/api/3/issue/PROJ-123/comment/10001",
			"body": map[string]interface{}{
				"type": "doc", "version": 1, "content": []interface{}{},
			},
			"author": map[string]interface{}{
				"emailAddress": "fallback@example.com",
			},
			"created": "2024-01-15T10:30:00.000+00:00",
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	comment, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name := getAuthorDisplayName(comment.Author)
	if name != "fallback@example.com" {
		t.Errorf("expected author fallback to email 'fallback@example.com', got %q", name)
	}
}

func TestGetCommentAPI_AuthorFallbackToAccountId(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := map[string]interface{}{
			"id":   "10001",
			"self": "https://test.atlassian.net/rest/api/3/issue/PROJ-123/comment/10001",
			"body": map[string]interface{}{
				"type": "doc", "version": 1, "content": []interface{}{},
			},
			"author": map[string]interface{}{
				"accountId": "5f3a2b1c0d9e8f7a6b5c4d3e",
			},
			"created": "2024-01-15T10:30:00.000+00:00",
		}
		data, _ := json.Marshal(response)
		w.Write(data)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	comment, _, err := client.IssueCommentsAPI.GetComment(ctx, "PROJ-123", "10001").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	name := getAuthorDisplayName(comment.Author)
	if name != "5f3a2b1c0d9e8f7a6b5c4d3e" {
		t.Errorf("expected author fallback to accountId '5f3a2b1c0d9e8f7a6b5c4d3e', got %q", name)
	}
}

func TestGetCommentCmd_RequiredFlags(t *testing.T) {
	idFlag := getCommentCmd.Flags().Lookup("id")
	if idFlag == nil {
		t.Fatal("expected --id flag to be defined")
	}
	if idFlag.Annotations[cobra.BashCompOneRequiredFlag] == nil {
		t.Error("expected --id flag to be marked required")
	}

	commentIDFlag := getCommentCmd.Flags().Lookup("comment-id")
	if commentIDFlag == nil {
		t.Fatal("expected --comment-id flag to be defined")
	}
	if commentIDFlag.Annotations[cobra.BashCompOneRequiredFlag] == nil {
		t.Error("expected --comment-id flag to be marked required")
	}
}

func TestGetCommentCmd_HasFormatFlag(t *testing.T) {
	flag := getCommentCmd.Flags().Lookup("format")
	if flag == nil {
		t.Fatal("expected --format flag to be defined")
	}
	if flag.DefValue != "table" {
		t.Errorf("expected --format default 'table', got %q", flag.DefValue)
	}
}

func TestGetCommentCmd_HasNoImagesFlag(t *testing.T) {
	flag := getCommentCmd.Flags().Lookup("no-images")
	if flag == nil {
		t.Fatal("expected --no-images flag to be defined")
	}
	if flag.DefValue != "false" {
		t.Errorf("expected --no-images default 'false', got %q", flag.DefValue)
	}
}

func TestGetCommentCmd_HasCommentOnlyFlag(t *testing.T) {
	flag := getCommentCmd.Flags().Lookup("comment-only")
	if flag == nil {
		t.Fatal("expected --comment-only flag to be defined")
	}
	if flag.DefValue != "false" {
		t.Errorf("expected --comment-only default 'false', got %q", flag.DefValue)
	}
}

// buildSingleCommentResponse creates a JSON response for a single comment (GET /issue/{key}/comment/{id})
func buildSingleCommentResponse(id, body, authorName, created string) string {
	response := map[string]interface{}{
		"id":   id,
		"self": "https://test.atlassian.net/rest/api/3/issue/PROJ-123/comment/" + id,
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
	data, _ := json.Marshal(response)
	return string(data)
}
