package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alexhokl/jira-cli/swagger"
)

func TestEditIssueAPI_WithIssueType(t *testing.T) {
	var capturedBody map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			body := map[string]interface{}{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			capturedBody = body
			w.WriteHeader(http.StatusNoContent)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
			})))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	typeId, err := resolveIssueTypeId(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("failed to resolve issue type: %v", err)
	}

	fields := map[string]interface{}{
		"issuetype": map[string]interface{}{
			"id": typeId,
		},
	}
	updateDetails := swagger.NewIssueUpdateDetails()
	updateDetails.SetFields(fields)

	_, _, err = client.IssuesAPI.EditIssue(ctx, "PROJ-123").
		IssueUpdateDetails(*updateDetails).
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if capturedBody == nil {
		t.Fatal("PUT request body was not captured")
	}

	fieldsMap, ok := capturedBody["fields"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'fields' in request body, got: %v", capturedBody)
	}
	issueType, ok := fieldsMap["issuetype"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'issuetype' in fields, got: %v", fieldsMap)
	}
	if issueType["id"] != "10001" {
		t.Errorf("expected issuetype id 10001, got %v", issueType["id"])
	}
}

func TestEditIssueAPI_IssueTypeUpdateValidation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPut {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"errors":{"issuetype":"Issue type is not available in this project"}}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	fields := map[string]interface{}{
		"issuetype": map[string]interface{}{
			"id": "10002",
		},
	}
	updateDetails := swagger.NewIssueUpdateDetails()
	updateDetails.SetFields(fields)

	_, _, err := client.IssuesAPI.EditIssue(ctx, "PROJ-123").
		IssueUpdateDetails(*updateDetails).
		Execute()
	if err == nil {
		t.Fatal("expected error for unavailable issue type, got nil")
	}
	if !strings.Contains(err.Error(), "issuetype") && !strings.Contains(err.Error(), "400") {
		t.Errorf("expected error mentioning issuetype or 400, got: %v", err)
	}
}
