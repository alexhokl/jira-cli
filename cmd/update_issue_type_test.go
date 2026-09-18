package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func buildIssueTypesResponseForTest(types []map[string]interface{}) string {
	var parts []string
	for _, t := range types {
		parts = append(parts, `{"id":"`+t["id"].(string)+`","name":"`+t["name"].(string)+`","subtask":`+boolToString(t["subtask"].(bool))+`}`)
	}
	return "[" + strings.Join(parts, ",") + "]"
}

func boolToString(b bool) string {
	if b {
		return "true"
	}
	return "false"
}

func TestResolveIssueTypeId_ByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issuetype" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
			{"id": "10001", "name": "Task", "subtask": false},
			{"id": "10002", "name": "Sub-task", "subtask": true},
		})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	id, err := resolveIssueTypeId(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "10001" {
		t.Errorf("expected 10001, got %s", id)
	}
}

func TestResolveIssueTypeId_ByNameCaseInsensitive(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issuetype" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
			{"id": "10001", "name": "Task", "subtask": false},
			{"id": "10002", "name": "Bug", "subtask": false},
		})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	id, err := resolveIssueTypeId(client, ctx, "PROJ", "bug")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "10002" {
		t.Errorf("expected 10002, got %s", id)
	}
}

func TestResolveIssueTypeId_ById(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issuetype" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
			{"id": "10001", "name": "Task", "subtask": false},
			{"id": "10002", "name": "Sub-task", "subtask": true},
		})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	id, err := resolveIssueTypeId(client, ctx, "PROJ", "10002")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "10002" {
		t.Errorf("expected 10002, got %s", id)
	}
}

func TestResolveIssueTypeId_NotFound(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issuetype" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
			{"id": "10001", "name": "Task", "subtask": false},
			{"id": "10002", "name": "Bug", "subtask": false},
		})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, err := resolveIssueTypeId(client, ctx, "PROJ", "Nonexistent")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "issue type \"Nonexistent\" not found") {
		t.Errorf("expected 'not found' error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "Task") {
		t.Errorf("expected available types in error message, got: %v", err)
	}
}

func TestResolveIssueTypeId_ApiError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, err := resolveIssueTypeId(client, ctx, "", "Task")
	if err == nil {
		t.Fatal("expected error, got nil")
	}
	if !strings.Contains(err.Error(), "failed to get issue types") {
		t.Errorf("expected 'failed to get issue types' error, got: %v", err)
	}
}

func TestResolveIssueTypeId_ProjectScopedLookup(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.URL.Path == "/rest/api/3/project/PROJ":
			w.Write([]byte(`{"key":"PROJ","name":"Project","id":"10000"}`))
		case r.URL.Path == "/rest/api/3/issuetype/project":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10010", "name": "Task", "subtask": false},
			})))
		case r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "99999", "name": "Task", "subtask": false},
			})))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	id, err := resolveIssueTypeId(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "10010" {
		t.Errorf("expected project-scoped ID 10010, got %s", id)
	}
}
