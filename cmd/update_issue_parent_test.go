package cmd

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func buildIssueResponseForSubtaskTest(subtask bool, includeSubtaskField bool) string {
	issueType := map[string]interface{}{
		"id":   "10002",
		"name": "Sub-task",
	}
	if includeSubtaskField {
		issueType["subtask"] = subtask
	}
	response := map[string]interface{}{
		"key":    "PROJ-123",
		"id":     "10001",
		"fields": map[string]interface{}{"issuetype": issueType},
	}
	data, _ := json.Marshal(response)
	return string(data)
}

func TestFetchIssueIsSubtask(t *testing.T) {
	tests := []struct {
		name                string
		response            string
		expectedSubtask     bool
		expectedError       bool
		expectedErrorSubstr string
	}{
		{
			name:            "subtask issue type",
			response:        buildIssueResponseForSubtaskTest(true, true),
			expectedSubtask: true,
		},
		{
			name:            "standard issue type",
			response:        buildIssueResponseForSubtaskTest(false, true),
			expectedSubtask: false,
		},
		{
			name: "missing subtask field treated as standard",
			response: `{
				"key": "PROJ-123",
				"id": "10001",
				"fields": {"issuetype": {"id": "10001", "name": "Task"}}
			}`,
			expectedSubtask: false,
		},
		{
			name: "missing issuetype field treated as standard",
			response: `{
				"key": "PROJ-123",
				"id": "10001",
				"fields": {"summary": "Hello"}
			}`,
			expectedSubtask: false,
		},
		{
			name:                "API error",
			response:            `{"errorMessages": ["Issue does not exist"]}`,
			expectedError:       true,
			expectedErrorSubstr: "Issue does not exist",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				if r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123" {
					statusCode := http.StatusOK
					if tt.expectedError {
						statusCode = http.StatusNotFound
					}
					w.WriteHeader(statusCode)
					w.Write([]byte(tt.response))
					return
				}
				w.WriteHeader(http.StatusNotFound)
			}))
			defer server.Close()

			client := newMockClientWithServer(server.URL)
			ctx := mockAuthContext()

			subtask, err := fetchIssueIsSubtask(client, ctx, "PROJ-123")
			if tt.expectedError {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.expectedErrorSubstr) {
					t.Errorf("expected error containing %q, got: %v", tt.expectedErrorSubstr, err)
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if subtask != tt.expectedSubtask {
				t.Errorf("expected subtask=%v, got %v", tt.expectedSubtask, subtask)
			}
		})
	}
}

func TestRunUpdateIssue_ParentNone_SubtaskWithoutType(t *testing.T) {
	var putCount int

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			putCount++
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	err := runUpdateIssue(nil, nil)
	if err == nil {
		t.Fatal("expected error for subtask parent removal, got nil")
	}
	if !strings.Contains(err.Error(), "--parent none is not supported for subtasks") {
		t.Errorf("expected guidance error, got: %v", err)
	}
	if !strings.Contains(err.Error(), "--type Task --parent none") {
		t.Errorf("expected error to suggest --type conversion, got: %v", err)
	}
	if putCount != 0 {
		t.Errorf("expected no PUT requests, got %d", putCount)
	}
}

func TestRunUpdateIssue_ParentNone_StandardIssueSinglePut(t *testing.T) {
	var putBodies []map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(false, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			body := map[string]interface{}{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			putBodies = append(putBodies, body)
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	if err := runUpdateIssue(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(putBodies) != 1 {
		t.Fatalf("expected 1 PUT request, got %d", len(putBodies))
	}
	updates, ok := putBodies[0]["update"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'update' in request body, got: %v", putBodies[0])
	}
	parentOps, ok := updates["parent"].([]interface{})
	if !ok || len(parentOps) != 1 {
		t.Fatalf("expected one parent update op, got: %v", updates["parent"])
	}
	parentOp, ok := parentOps[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected parent op object, got: %v", parentOps[0])
	}
	setVal, ok := parentOp["set"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'set' in parent op, got: %v", parentOp)
	}
	if setVal["none"] != true {
		t.Errorf("expected set.none=true, got: %v", setVal)
	}
}

func TestRunUpdateIssue_ParentNone_SubtaskWithTypeTwoPuts(t *testing.T) {
	var putBodies []map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
				{"id": "10002", "name": "Sub-task", "subtask": true},
			})))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/project/PROJ":
			w.WriteHeader(http.StatusNotFound)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			body := map[string]interface{}{}
			if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			putBodies = append(putBodies, body)
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Task"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	if err := runUpdateIssue(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(putBodies) != 2 {
		t.Fatalf("expected 2 PUT requests, got %d", len(putBodies))
	}

	// First PUT: issue type conversion only
	firstFields, ok := putBodies[0]["fields"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'fields' in first PUT body, got: %v", putBodies[0])
	}
	issueType, ok := firstFields["issuetype"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'issuetype' in first PUT fields, got: %v", firstFields)
	}
	if issueType["id"] != "10001" {
		t.Errorf("expected issuetype id 10001 in first PUT, got %v", issueType["id"])
	}
	if _, hasUpdate := putBodies[0]["update"]; hasUpdate {
		t.Errorf("expected no 'update' in first PUT body, got: %v", putBodies[0]["update"])
	}

	// Second PUT: parent removal only
	secondUpdates, ok := putBodies[1]["update"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'update' in second PUT body, got: %v", putBodies[1])
	}
	parentOps, ok := secondUpdates["parent"].([]interface{})
	if !ok || len(parentOps) != 1 {
		t.Fatalf("expected one parent update op in second PUT, got: %v", secondUpdates["parent"])
	}
	parentOp, ok := parentOps[0].(map[string]interface{})
	if !ok {
		t.Fatalf("expected parent op object, got: %v", parentOps[0])
	}
	setVal, ok := parentOp["set"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'set' in parent op, got: %v", parentOp)
	}
	if setVal["none"] != true {
		t.Errorf("expected set.none=true in second PUT, got: %v", setVal)
	}
	if _, hasFields := putBodies[1]["fields"]; hasFields {
		t.Errorf("expected no 'fields' in second PUT body, got: %v", putBodies[1]["fields"])
	}
}

func TestRunUpdateIssue_ParentNone_SubtaskWithType_TargetIsSubtask(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
				{"id": "10002", "name": "Sub-task", "subtask": true},
			})))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/project/PROJ":
			w.WriteHeader(http.StatusNotFound)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Sub-task"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	err := runUpdateIssue(nil, nil)
	if err == nil {
		t.Fatal("expected error when target type is a subtask, got nil")
	}
	if !strings.Contains(err.Error(), "parent removal is only supported for standard issue types") {
		t.Errorf("expected subtask-target error, got: %v", err)
	}
}

func TestRunUpdateIssue_ParentNone_TypeChangeFails(t *testing.T) {
	putCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
			})))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/project/PROJ":
			w.WriteHeader(http.StatusNotFound)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			putCount++
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"errors":{"issuetype":"Issue type is not available in this project"}}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Task"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	err := runUpdateIssue(nil, nil)
	if err == nil {
		t.Fatal("expected error when type conversion fails, got nil")
	}
	if !strings.Contains(err.Error(), "400 Bad Request") && !strings.Contains(err.Error(), "issuetype") {
		t.Errorf("expected conversion failure error, got: %v", err)
	}
	if putCount != 1 {
		t.Errorf("expected exactly 1 PUT request (no parent-removal retry), got %d", putCount)
	}
}

func TestRunUpdateIssue_ParentNone_NoNotify(t *testing.T) {
	var sawNotifyQuery []string

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
			})))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/project/PROJ":
			w.WriteHeader(http.StatusNotFound)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			sawNotifyQuery = append(sawNotifyQuery, r.URL.Query().Get("notifyUsers"))
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Task"
	updateIssueOpts.parent = "none"
	updateIssueOpts.noNotify = true

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	if err := runUpdateIssue(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(sawNotifyQuery) != 2 {
		t.Fatalf("expected 2 PUT requests, got %d", len(sawNotifyQuery))
	}
	for i, v := range sawNotifyQuery {
		if v != "false" {
			t.Errorf("PUT %d: expected notifyUsers=false, got %q", i+1, v)
		}
	}
}

func TestRunUpdateIssue_ParentNone_GetIssueError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123" {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"errorMessages": ["Internal server error"]}`))
			return
		}
		w.WriteHeader(http.StatusNotFound)
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	err := runUpdateIssue(nil, nil)
	if err == nil {
		t.Fatal("expected error when issue lookup fails, got nil")
	}
	if !strings.Contains(err.Error(), "failed to get issue") {
		t.Errorf("expected issue lookup failure error, got: %v", err)
	}
}

// resetUpdateIssueOpts resets the package-level update issue options.
func resetUpdateIssueOpts() {
	updateIssueOpts = updateIssueOptions{}
}

func TestUpdateIssueOptions_TypeChangeRequested(t *testing.T) {
	tests := []struct {
		name      string
		issueType string
		expected  bool
	}{
		{"type set", "Task", true},
		{"type empty", "", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			opts := updateIssueOptions{issueType: tt.issueType}
			if got := opts.typeChangeRequested(); got != tt.expected {
				t.Errorf("typeChangeRequested() = %v, want %v", got, tt.expected)
			}
		})
	}
}

func TestResolveIssueTypeId_WrapsResolveIssueType(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issuetype" {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
			{"id": "10001", "name": "Task", "subtask": false},
		})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	typeDetails, err := resolveIssueType(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if typeDetails.GetId() != "10001" {
		t.Errorf("expected id 10001, got %s", typeDetails.GetId())
	}
	if typeDetails.GetSubtask() {
		t.Error("expected Task to not be a subtask")
	}

	id, err := resolveIssueTypeId(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if id != "10001" {
		t.Errorf("expected id 10001, got %s", id)
	}
}

func TestResolveIssueType_SubtaskFlag(t *testing.T) {
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

	subtaskType, err := resolveIssueType(client, ctx, "PROJ", "Sub-task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if !subtaskType.GetSubtask() {
		t.Error("expected Sub-task to be flagged as subtask")
	}

	standardType, err := resolveIssueType(client, ctx, "PROJ", "Task")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if standardType.GetSubtask() {
		t.Error("expected Task to not be flagged as subtask")
	}
}

func TestRunUpdateIssue_ParentNone_ConversionOnlyOutput(t *testing.T) {
	// Verify the combined conversion path succeeds end-to-end without
	// triggering the sprint/transition/linking follow-up branches
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.Write([]byte(buildIssueTypesResponseForTest([]map[string]interface{}{
				{"id": "10001", "name": "Task", "subtask": false},
			})))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/project/PROJ":
			w.WriteHeader(http.StatusNotFound)
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.WriteHeader(http.StatusNoContent)
		default:
			t.Errorf("unexpected request: %s %s", r.Method, r.URL.Path)
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Task"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	if err := runUpdateIssue(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestRunUpdateIssue_ParentSetOnSubtaskStillWorks(t *testing.T) {
	var putBody map[string]interface{}

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodPut && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			if err := json.NewDecoder(r.Body).Decode(&putBody); err != nil {
				t.Errorf("failed to decode request body: %v", err)
			}
			w.WriteHeader(http.StatusNoContent)
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.parent = "PROJ-100"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	if err := runUpdateIssue(nil, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	fields, ok := putBody["fields"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'fields' in request body, got: %v", putBody)
	}
	parent, ok := fields["parent"].(map[string]interface{})
	if !ok {
		t.Fatalf("expected 'parent' in fields, got: %v", fields)
	}
	if parent["key"] != "PROJ-100" {
		t.Errorf("expected parent key PROJ-100, got %v", parent["key"])
	}
}

func TestRunUpdateIssue_ParentNone_IssueTypeError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		switch {
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issue/PROJ-123":
			w.Write([]byte(buildIssueResponseForSubtaskTest(true, true)))
		case r.Method == http.MethodGet && r.URL.Path == "/rest/api/3/issuetype":
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"errorMessages": ["boom"]}`))
		default:
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	resetUpdateIssueOpts()
	defer resetUpdateIssueOpts()
	updateIssueOpts.issueKey = "PROJ-123"
	updateIssueOpts.issueType = "Task"
	updateIssueOpts.parent = "none"

	overrideServerURL = server.URL
	defer func() { overrideServerURL = "" }()

	err := runUpdateIssue(nil, nil)
	if err == nil {
		t.Fatal("expected error when issue type lookup fails, got nil")
	}
	if !strings.Contains(err.Error(), "failed to get issue types") {
		t.Errorf("expected issue type lookup failure error, got: %v", err)
	}
}
