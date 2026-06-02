package cmd

import (
	"encoding/json"
	"testing"
)

// jsonKeys unmarshals data into a map and returns the top-level keys of the
// first element when the JSON is an array, or the object itself when it is a
// single object. It is a test helper used to assert that struct fields are
// exported with the expected JSON key names.
func jsonKeys(t *testing.T, v any) map[string]any {
	t.Helper()
	b, err := json.Marshal(v)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	// Try array first
	var arr []map[string]any
	if json.Unmarshal(b, &arr) == nil && len(arr) > 0 {
		return arr[0]
	}
	// Fall back to object
	var obj map[string]any
	if err := json.Unmarshal(b, &obj); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	return obj
}

func assertKey(t *testing.T, m map[string]any, key string, want any) {
	t.Helper()
	v, ok := m[key]
	if !ok {
		t.Errorf("key %q missing from JSON; available keys: %v", key, m)
		return
	}
	// json.Unmarshal uses float64 for numbers; normalise for comparison.
	switch w := want.(type) {
	case int:
		if v != float64(w) {
			t.Errorf("key %q: got %v, want %v", key, v, want)
		}
	case int64:
		if v != float64(w) {
			t.Errorf("key %q: got %v, want %v", key, v, want)
		}
	case bool:
		if v != w {
			t.Errorf("key %q: got %v, want %v", key, v, want)
		}
	default:
		if v != want {
			t.Errorf("key %q: got %v, want %v", key, v, want)
		}
	}
}

// ---------------------------------------------------------------------------
// statusInfo
// ---------------------------------------------------------------------------

func TestStatusInfoJSON(t *testing.T) {
	s := statusInfo{
		ID:          "1",
		Name:        "In Progress",
		Category:    "IN_PROGRESS",
		Scope:       "GLOBAL",
		Project:     "PROJ",
		Description: "Work in progress",
	}
	m := jsonKeys(t, []statusInfo{s})
	assertKey(t, m, "id", "1")
	assertKey(t, m, "name", "In Progress")
	assertKey(t, m, "category", "IN_PROGRESS")
	assertKey(t, m, "scope", "GLOBAL")
	assertKey(t, m, "project", "PROJ")
	assertKey(t, m, "description", "Work in progress")
}

// ---------------------------------------------------------------------------
// workflowInfo
// ---------------------------------------------------------------------------

func TestWorkflowInfoJSON(t *testing.T) {
	w := workflowInfo{Name: "My Workflow", Description: "A workflow"}
	m := jsonKeys(t, []workflowInfo{w})
	assertKey(t, m, "name", "My Workflow")
	assertKey(t, m, "description", "A workflow")
}

// ---------------------------------------------------------------------------
// customFieldDetails
// ---------------------------------------------------------------------------

func TestCustomFieldDetailsJSON(t *testing.T) {
	f := customFieldDetails{
		ID:         "customfield_10001",
		Name:       "Story Points",
		FieldType:  "number",
		Schema:     "float",
		IssueTypes: "Story, Epic",
	}
	m := jsonKeys(t, []customFieldDetails{f})
	assertKey(t, m, "id", "customfield_10001")
	assertKey(t, m, "name", "Story Points")
	assertKey(t, m, "fieldType", "number")
	assertKey(t, m, "schema", "float")
	assertKey(t, m, "issueTypes", "Story, Epic")
}

// ---------------------------------------------------------------------------
// customFieldOptionValue
// ---------------------------------------------------------------------------

func TestCustomFieldOptionValueJSON(t *testing.T) {
	o := customFieldOptionValue{
		ID:          "10001",
		Value:       "High",
		Disabled:    false,
		ContextName: "Default Context",
	}
	m := jsonKeys(t, []customFieldOptionValue{o})
	assertKey(t, m, "id", "10001")
	assertKey(t, m, "value", "High")
	assertKey(t, m, "disabled", false)
	assertKey(t, m, "contextName", "Default Context")
}

// ---------------------------------------------------------------------------
// issueTypeSchemeInfo
// ---------------------------------------------------------------------------

func TestIssueTypeSchemeInfoJSON(t *testing.T) {
	s := issueTypeSchemeInfo{
		ID:          "10000",
		Name:        "Default Issue Type Scheme",
		Description: "Default scheme",
		IsDefault:   "Yes",
		Projects:    "PROJ, OTHER",
	}
	m := jsonKeys(t, []issueTypeSchemeInfo{s})
	assertKey(t, m, "id", "10000")
	assertKey(t, m, "name", "Default Issue Type Scheme")
	assertKey(t, m, "description", "Default scheme")
	assertKey(t, m, "isDefault", "Yes")
	assertKey(t, m, "projects", "PROJ, OTHER")
}

// ---------------------------------------------------------------------------
// workflowSchemeInfo
// ---------------------------------------------------------------------------

func TestWorkflowSchemeInfoJSON(t *testing.T) {
	s := workflowSchemeInfo{
		ID:              42,
		Name:            "Software Workflow Scheme",
		Description:     "For software projects",
		DefaultWorkflow: "Software Simplified Workflow",
		Draft:           false,
		Projects:        "PROJ",
	}
	m := jsonKeys(t, []workflowSchemeInfo{s})
	assertKey(t, m, "id", int64(42))
	assertKey(t, m, "name", "Software Workflow Scheme")
	assertKey(t, m, "description", "For software projects")
	assertKey(t, m, "defaultWorkflow", "Software Simplified Workflow")
	assertKey(t, m, "draft", false)
	assertKey(t, m, "projects", "PROJ")
}

// ---------------------------------------------------------------------------
// issueTypeInfo
// ---------------------------------------------------------------------------

func TestIssueTypeInfoJSON(t *testing.T) {
	i := issueTypeInfo{
		ID:          "10001",
		Name:        "Bug",
		Description: "A bug",
		Subtask:     false,
		Project:     "PROJ",
	}
	m := jsonKeys(t, []issueTypeInfo{i})
	assertKey(t, m, "id", "10001")
	assertKey(t, m, "name", "Bug")
	assertKey(t, m, "description", "A bug")
	assertKey(t, m, "subtask", false)
	assertKey(t, m, "project", "PROJ")
}

// ---------------------------------------------------------------------------
// issueTransitionInfo
// ---------------------------------------------------------------------------

func TestIssueTransitionInfoJSON(t *testing.T) {
	tr := issueTransitionInfo{
		ID:       "21",
		Name:     "In Progress",
		ToStatus: "In Progress",
		Category: "In Progress",
	}
	m := jsonKeys(t, []issueTransitionInfo{tr})
	assertKey(t, m, "id", "21")
	assertKey(t, m, "name", "In Progress")
	assertKey(t, m, "toStatus", "In Progress")
	assertKey(t, m, "category", "In Progress")
}

// ---------------------------------------------------------------------------
// workflowStatusInfo
// ---------------------------------------------------------------------------

func TestWorkflowStatusInfoJSON(t *testing.T) {
	s := workflowStatusInfo{
		ID:          "1",
		Name:        "To Do",
		Category:    "TODO",
		Description: "Not started",
	}
	m := jsonKeys(t, []workflowStatusInfo{s})
	assertKey(t, m, "id", "1")
	assertKey(t, m, "name", "To Do")
	assertKey(t, m, "category", "TODO")
	assertKey(t, m, "description", "Not started")
}

// ---------------------------------------------------------------------------
// statusPropertyInfo
// ---------------------------------------------------------------------------

func TestStatusPropertyInfoJSON(t *testing.T) {
	p := statusPropertyInfo{
		StatusName: "In Progress",
		Key:        "jira.issue.editable",
		Value:      "true",
	}
	m := jsonKeys(t, []statusPropertyInfo{p})
	assertKey(t, m, "statusName", "In Progress")
	assertKey(t, m, "key", "jira.issue.editable")
	assertKey(t, m, "value", "true")
}

// ---------------------------------------------------------------------------
// boardOutput
// ---------------------------------------------------------------------------

func TestBoardOutputJSON(t *testing.T) {
	b := boardOutput{
		ID:         101,
		Name:       "My Scrum Board",
		Type:       "scrum",
		ProjectKey: "PROJ",
	}
	m := jsonKeys(t, []boardOutput{b})
	assertKey(t, m, "id", int64(101))
	assertKey(t, m, "name", "My Scrum Board")
	assertKey(t, m, "type", "scrum")
	assertKey(t, m, "projectKey", "PROJ")
}

// ---------------------------------------------------------------------------
// projectOutput
// ---------------------------------------------------------------------------

func TestProjectOutputJSON(t *testing.T) {
	p := projectOutput{Key: "PROJ", Name: "My Project", Type: "software"}
	m := jsonKeys(t, []projectOutput{p})
	assertKey(t, m, "key", "PROJ")
	assertKey(t, m, "name", "My Project")
	assertKey(t, m, "type", "software")
}

// ---------------------------------------------------------------------------
// sprintOutput
// ---------------------------------------------------------------------------

func TestSprintOutputJSON(t *testing.T) {
	s := sprintOutput{
		ID:        55,
		Name:      "Sprint 1",
		State:     "active",
		StartDate: "2024-01-01",
		EndDate:   "2024-01-14",
	}
	m := jsonKeys(t, []sprintOutput{s})
	assertKey(t, m, "id", int64(55))
	assertKey(t, m, "name", "Sprint 1")
	assertKey(t, m, "state", "active")
	assertKey(t, m, "startDate", "2024-01-01")
	assertKey(t, m, "endDate", "2024-01-14")
}

// ---------------------------------------------------------------------------
// userOutput
// ---------------------------------------------------------------------------

func TestUserOutputJSON(t *testing.T) {
	u := userOutput{
		DisplayName: "Alice",
		Email:       "alice@example.com",
		AccountType: "atlassian",
		Active:      true,
	}
	m := jsonKeys(t, []userOutput{u})
	assertKey(t, m, "displayName", "Alice")
	assertKey(t, m, "email", "alice@example.com")
	assertKey(t, m, "accountType", "atlassian")
	assertKey(t, m, "active", true)
}

// ---------------------------------------------------------------------------
// linkTypeOutput
// ---------------------------------------------------------------------------

func TestLinkTypeOutputJSON(t *testing.T) {
	l := linkTypeOutput{
		Name:    "Blocks",
		Inward:  "is blocked by",
		Outward: "blocks",
	}
	m := jsonKeys(t, []linkTypeOutput{l})
	assertKey(t, m, "name", "Blocks")
	assertKey(t, m, "inward", "is blocked by")
	assertKey(t, m, "outward", "blocks")
}

// ---------------------------------------------------------------------------
// issueOutput
// ---------------------------------------------------------------------------

func TestIssueOutputJSON(t *testing.T) {
	o := issueOutput{
		Key:      "PROJ-1",
		Type:     "Bug",
		Summary:  "Something broke",
		Status:   "In Progress",
		Assignee: "Alice",
		Reporter: "Bob",
		Priority: "High",
	}
	m := jsonKeys(t, []issueOutput{o})
	assertKey(t, m, "key", "PROJ-1")
	assertKey(t, m, "type", "Bug")
	assertKey(t, m, "summary", "Something broke")
	assertKey(t, m, "status", "In Progress")
	assertKey(t, m, "assignee", "Alice")
	assertKey(t, m, "reporter", "Bob")
	assertKey(t, m, "priority", "High")
}

// ---------------------------------------------------------------------------
// commentOutput
// ---------------------------------------------------------------------------

func TestCommentOutputJSON(t *testing.T) {
	adfBody := map[string]any{
		"type":    "doc",
		"version": 1,
		"content": []any{},
	}
	c := commentOutput{
		ID:      "10001",
		Author:  "Alice",
		Created: "2024-01-01",
		Body:    adfBody,
	}
	m := jsonKeys(t, []commentOutput{c})
	assertKey(t, m, "id", "10001")
	assertKey(t, m, "author", "Alice")
	assertKey(t, m, "created", "2024-01-01")
	if _, ok := m["body"]; !ok {
		t.Error("key \"body\" missing from JSON")
	}
}

// ---------------------------------------------------------------------------
// Labels — []string passthrough
// ---------------------------------------------------------------------------

func TestLabelsJSONPassthrough(t *testing.T) {
	labels := []string{"bug", "enhancement"}
	b, err := json.Marshal(labels)
	if err != nil {
		t.Fatalf("json.Marshal: %v", err)
	}
	var decoded []string
	if err := json.Unmarshal(b, &decoded); err != nil {
		t.Fatalf("json.Unmarshal: %v", err)
	}
	if len(decoded) != 2 || decoded[0] != "bug" || decoded[1] != "enhancement" {
		t.Errorf("unexpected decoded labels: %v", decoded)
	}
}

// ---------------------------------------------------------------------------
// list_issues mutual exclusion guard
// ---------------------------------------------------------------------------

func TestListIssuesMutualExclusion(t *testing.T) {
	// Save and restore opts
	saved := listIssuesOpts
	defer func() { listIssuesOpts = saved }()

	listIssuesOpts.idOnly = true
	listIssuesOpts.format = "json"

	err := runListIssues(nil, nil)
	if err == nil {
		t.Fatal("expected error when --id-only and --format json are used together, got nil")
	}
	if err.Error() != "--id-only and --format json cannot be used together" {
		t.Errorf("unexpected error message: %v", err)
	}
}
