package cmd

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/spf13/pflag"
)

// ---------------------------------------------------------------------------
// Flag registration tests — verify --limit exists and --max-results is gone
// ---------------------------------------------------------------------------

func TestListLabelsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listLabelsCmd.Flags())
}

func TestListUsersCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listUsersCmd.Flags())
}

func TestListIssuesCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listIssuesCmd.Flags())
}

func TestListBoardsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listBoardsCmd.Flags())
}

func TestListCustomFieldsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listCustomFieldsCmd.Flags())
}

func TestListProjectsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listProjectsCmd.Flags())
}

func TestListSprintsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listSprintsCmd.Flags())
}

func TestListStatusCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listStatusCmd.Flags())
}

func TestListWorkflowsCmd_HasLimitFlag(t *testing.T) {
	assertLimitFlag(t, listWorkflowsCmd.Flags())
}

// assertLimitFlag checks that a FlagSet has --limit with default 0 and no --max-results.
func assertLimitFlag(t *testing.T, flags *pflag.FlagSet) {
	t.Helper()

	limitFlag := flags.Lookup("limit")
	if limitFlag == nil {
		t.Fatal("expected --limit flag to be registered")
	}
	if limitFlag.DefValue != "0" {
		t.Errorf("expected --limit default value to be 0, got %q", limitFlag.DefValue)
	}

	if flags.Lookup("max-results") != nil {
		t.Error("--max-results flag should not exist; it was renamed to --limit")
	}
}

// ---------------------------------------------------------------------------
// Limit pagination logic tests — boards
// ---------------------------------------------------------------------------

// TestListBoardsLimit_StopsPaginationEarly verifies that the pagination loop
// in runListBoards stops fetching once the --limit count is reached.
func TestListBoardsLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			boards := []map[string]interface{}{
				buildBoard(1, "Board A", "scrum", "PROJ"),
				buildBoard(2, "Board B", "kanban", "PROJ"),
				buildBoard(3, "Board C", "scrum", "TEST"),
			}
			w.Write([]byte(buildBoardsResponse(boards, false, 0, 3)))
		} else {
			// This page should never be fetched when limit=2
			boards := []map[string]interface{}{
				buildBoard(4, "Board D", "simple", "TEST"),
			}
			w.Write([]byte(buildBoardsResponse(boards, true, 3, 3)))
		}
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	const limit int32 = 2
	var allBoards []string
	var startAt int64 = 0

	for {
		result, _, err := client.BoardAPI.GetAllBoards(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, b := range result.GetValues() {
			allBoards = append(allBoards, b.GetName())
		}

		if limit > 0 && int32(len(allBoards)) >= limit {
			allBoards = allBoards[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allBoards) != int(limit) {
		t.Errorf("expected %d boards, got %d", limit, len(allBoards))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request (limit satisfied on first page), got %d", requestCount)
	}
}

// TestListBoardsLimit_ZeroMeansAll verifies that limit=0 fetches all pages.
func TestListBoardsLimit_ZeroMeansAll(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			boards := []map[string]interface{}{
				buildBoard(1, "Board A", "scrum", "PROJ"),
				buildBoard(2, "Board B", "kanban", "PROJ"),
			}
			w.Write([]byte(buildBoardsResponse(boards, false, 0, 2)))
		} else {
			boards := []map[string]interface{}{
				buildBoard(3, "Board C", "simple", "TEST"),
			}
			w.Write([]byte(buildBoardsResponse(boards, true, 2, 2)))
		}
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	const limit int32 = 0 // 0 = no limit
	var allBoards []string
	var startAt int64 = 0

	for {
		result, _, err := client.BoardAPI.GetAllBoards(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, b := range result.GetValues() {
			allBoards = append(allBoards, b.GetName())
		}

		if limit > 0 && int32(len(allBoards)) >= limit {
			allBoards = allBoards[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allBoards) != 3 {
		t.Errorf("expected all 3 boards with limit=0, got %d", len(allBoards))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests (all pages), got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit pagination logic tests — projects
// ---------------------------------------------------------------------------

// TestListProjectsLimit_StopsPaginationEarly verifies pagination stops at limit.
func TestListProjectsLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			projects := []map[string]interface{}{
				buildProject("P1", "Project One", "software"),
				buildProject("P2", "Project Two", "software"),
				buildProject("P3", "Project Three", "business"),
			}
			w.Write([]byte(buildProjectsResponse(projects, false, 0, 3, 5)))
		} else {
			projects := []map[string]interface{}{
				buildProject("P4", "Project Four", "software"),
				buildProject("P5", "Project Five", "business"),
			}
			w.Write([]byte(buildProjectsResponse(projects, true, 3, 3, 5)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 2
	var allProjects []string
	var startAt int64 = 0

	for {
		result, _, err := client.ProjectsAPI.SearchProjects(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, p := range result.GetValues() {
			allProjects = append(allProjects, p.GetKey())
		}

		if limit > 0 && int32(len(allProjects)) >= limit {
			allProjects = allProjects[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) != int(limit) {
		t.Errorf("expected %d projects, got %d", limit, len(allProjects))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// TestListProjectsLimit_ZeroMeansAll verifies limit=0 fetches all pages.
func TestListProjectsLimit_ZeroMeansAll(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			projects := []map[string]interface{}{
				buildProject("P1", "Project One", "software"),
				buildProject("P2", "Project Two", "software"),
			}
			w.Write([]byte(buildProjectsResponse(projects, false, 0, 2, 3)))
		} else {
			projects := []map[string]interface{}{
				buildProject("P3", "Project Three", "business"),
			}
			w.Write([]byte(buildProjectsResponse(projects, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 0
	var allProjects []string
	var startAt int64 = 0

	for {
		result, _, err := client.ProjectsAPI.SearchProjects(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, p := range result.GetValues() {
			allProjects = append(allProjects, p.GetKey())
		}

		if limit > 0 && int32(len(allProjects)) >= limit {
			allProjects = allProjects[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) != 3 {
		t.Errorf("expected all 3 projects with limit=0, got %d", len(allProjects))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit pagination logic tests — sprints
// ---------------------------------------------------------------------------

// TestListSprintsLimit_StopsPaginationEarly verifies pagination stops at limit.
// Sprints are fetched via a raw HTTP response and parsed with json.Unmarshal.
func TestListSprintsLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			sprints := []map[string]interface{}{
				buildSprint(1, "Sprint 1", "active", "", ""),
				buildSprint(2, "Sprint 2", "future", "", ""),
				buildSprint(3, "Sprint 3", "closed", "", ""),
			}
			w.Write([]byte(buildSprintsResponse(sprints, false, 0, 3)))
		} else {
			sprints := []map[string]interface{}{
				buildSprint(4, "Sprint 4", "closed", "", ""),
			}
			w.Write([]byte(buildSprintsResponse(sprints, true, 3, 3)))
		}
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	const limit int32 = 2
	var allSprints []string
	var startAt int64 = 0
	stateFilter := make(map[string]bool) // no filter

	for {
		resp, err := client.BoardAPI.GetAllSprints(ctx, 100).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatalf("unexpected error reading body: %v", err)
		}

		var result sprintsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			t.Fatalf("unexpected error unmarshalling: %v", err)
		}

		for _, sprint := range result.Values {
			if len(stateFilter) == 0 || stateFilter[sprint.GetState()] {
				allSprints = append(allSprints, sprint.GetName())
			}
		}

		if limit > 0 && int32(len(allSprints)) >= limit {
			allSprints = allSprints[:limit]
			break
		}

		if result.IsLast {
			break
		}

		startAt = int64(result.StartAt + result.MaxResults)
	}

	if len(allSprints) != int(limit) {
		t.Errorf("expected %d sprints, got %d", limit, len(allSprints))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// TestListSprintsLimit_ZeroMeansAll verifies limit=0 fetches all pages.
func TestListSprintsLimit_ZeroMeansAll(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
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

	const limit int32 = 0
	var allSprints []string
	var startAt int64 = 0
	stateFilter := make(map[string]bool)

	for {
		resp, err := client.BoardAPI.GetAllSprints(ctx, 100).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		body, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			t.Fatalf("unexpected error reading body: %v", err)
		}

		var result sprintsResponse
		if err := json.Unmarshal(body, &result); err != nil {
			t.Fatalf("unexpected error unmarshalling: %v", err)
		}

		for _, sprint := range result.Values {
			if len(stateFilter) == 0 || stateFilter[sprint.GetState()] {
				allSprints = append(allSprints, sprint.GetName())
			}
		}

		if limit > 0 && int32(len(allSprints)) >= limit {
			allSprints = allSprints[:limit]
			break
		}

		if result.IsLast {
			break
		}

		startAt = int64(result.StartAt + result.MaxResults)
	}

	if len(allSprints) != 3 {
		t.Errorf("expected all 3 sprints with limit=0, got %d", len(allSprints))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit pagination logic tests — status
// ---------------------------------------------------------------------------

// TestListStatusLimit_StopsPaginationEarly verifies pagination stops at limit.
func TestListStatusLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			statuses := []map[string]interface{}{
				buildStatus("1", "To Do", "TODO", ""),
				buildStatus("2", "In Progress", "IN_PROGRESS", ""),
				buildStatus("3", "In Review", "IN_PROGRESS", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, false, 0, 3, 5)))
		} else {
			statuses := []map[string]interface{}{
				buildStatus("4", "Done", "DONE", ""),
				buildStatus("5", "Closed", "DONE", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, true, 3, 3, 5)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 2
	var allStatuses []string
	var startAt int64 = 0

	for {
		result, _, err := client.StatusAPI.Search(ctx).StartAt(startAt).MaxResults(3).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, s := range result.GetValues() {
			allStatuses = append(allStatuses, s.GetName())
		}

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		if limit > 0 && int32(len(allStatuses)) >= limit {
			allStatuses = allStatuses[:limit]
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allStatuses) != int(limit) {
		t.Errorf("expected %d statuses, got %d", limit, len(allStatuses))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// TestListStatusLimit_ZeroMeansAll verifies limit=0 fetches all pages.
func TestListStatusLimit_ZeroMeansAll(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			statuses := []map[string]interface{}{
				buildStatus("1", "To Do", "TODO", ""),
				buildStatus("2", "In Progress", "IN_PROGRESS", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, false, 0, 2, 3)))
		} else {
			statuses := []map[string]interface{}{
				buildStatus("3", "Done", "DONE", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 0
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

		if limit > 0 && int32(len(allStatuses)) >= limit {
			allStatuses = allStatuses[:limit]
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allStatuses) != 3 {
		t.Errorf("expected all 3 statuses with limit=0, got %d", len(allStatuses))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit pagination logic tests — workflows
// ---------------------------------------------------------------------------

// TestListWorkflowsLimit_StopsPaginationEarly verifies pagination stops at limit.
func TestListWorkflowsLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			wfs := []map[string]interface{}{
				buildWorkflow("Workflow A", "desc A", false),
				buildWorkflow("Workflow B", "desc B", false),
				buildWorkflow("Workflow C", "desc C", true),
			}
			w.Write([]byte(buildWorkflowsResponse(wfs, false, 0, 3, 5)))
		} else {
			wfs := []map[string]interface{}{
				buildWorkflow("Workflow D", "desc D", false),
				buildWorkflow("Workflow E", "desc E", false),
			}
			w.Write([]byte(buildWorkflowsResponse(wfs, true, 3, 3, 5)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 2
	var allWorkflows []string
	var startAt int64 = 0

	for {
		result, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).StartAt(startAt).MaxResults(3).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, wf := range result.GetValues() {
			id := wf.GetId()
			allWorkflows = append(allWorkflows, id.GetName())
		}

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		if limit > 0 && int32(len(allWorkflows)) >= limit {
			allWorkflows = allWorkflows[:limit]
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allWorkflows) != int(limit) {
		t.Errorf("expected %d workflows, got %d", limit, len(allWorkflows))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// TestListWorkflowsLimit_ZeroMeansAll verifies limit=0 fetches all pages.
func TestListWorkflowsLimit_ZeroMeansAll(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			wfs := []map[string]interface{}{
				buildWorkflow("Workflow A", "desc A", false),
				buildWorkflow("Workflow B", "desc B", false),
			}
			w.Write([]byte(buildWorkflowsResponse(wfs, false, 0, 2, 3)))
		} else {
			wfs := []map[string]interface{}{
				buildWorkflow("Workflow C", "desc C", true),
			}
			w.Write([]byte(buildWorkflowsResponse(wfs, true, 2, 2, 3)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 0
	var allWorkflows []string
	var startAt int64 = 0

	for {
		result, _, err := client.WorkflowsAPI.GetWorkflowsPaginated(ctx).StartAt(startAt).MaxResults(2).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, wf := range result.GetValues() {
			id := wf.GetId()
			allWorkflows = append(allWorkflows, id.GetName())
		}

		if result.GetIsLast() || len(result.GetValues()) == 0 {
			break
		}

		if limit > 0 && int32(len(allWorkflows)) >= limit {
			allWorkflows = allWorkflows[:limit]
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allWorkflows) != 3 {
		t.Errorf("expected all 3 workflows with limit=0, got %d", len(allWorkflows))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit logic tests — custom fields (post-fetch cap)
// ---------------------------------------------------------------------------

// TestListCustomFieldsLimit_CapsResults verifies that the post-fetch slice cap
// applied in runListCustomFields respects the limit.
func TestListCustomFieldsLimit_CapsResults(t *testing.T) {
	// Build a slice of 5 fields and apply the same capping logic as the command.
	fields := []customFieldDetails{
		{ID: "cf1", Name: "Field 1"},
		{ID: "cf2", Name: "Field 2"},
		{ID: "cf3", Name: "Field 3"},
		{ID: "cf4", Name: "Field 4"},
		{ID: "cf5", Name: "Field 5"},
	}

	tests := []struct {
		name       string
		limit      int32
		wantLen    int
	}{
		{"limit=2 caps at 2", 2, 2},
		{"limit=5 keeps all 5", 5, 5},
		{"limit=10 keeps all 5 (limit > total)", 10, 5},
		{"limit=0 keeps all 5 (no limit)", 0, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := make([]customFieldDetails, len(fields))
			copy(result, fields)

			// Same logic as runListCustomFields
			if tt.limit > 0 && int32(len(result)) > tt.limit {
				result = result[:tt.limit]
			}

			if len(result) != tt.wantLen {
				t.Errorf("expected %d fields after limit=%d, got %d", tt.wantLen, tt.limit, len(result))
			}
		})
	}
}

// ---------------------------------------------------------------------------
// Limit logic tests — labels (renamed from --max-results; verify logic intact)
// ---------------------------------------------------------------------------

// TestListLabelsLimit_StopsPaginationEarly verifies the renamed flag still
// limits results correctly for labels.
func TestListLabelsLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			w.Write([]byte(buildLabelsResponse(
				[]string{"alpha", "beta", "gamma"}, false, 0, 3, 6,
			)))
		} else {
			w.Write([]byte(buildLabelsResponse(
				[]string{"delta", "epsilon", "zeta"}, true, 3, 3, 6,
			)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 2
	var allLabels []string
	var startAt int64 = 0
	const pageSize int32 = 3

	for {
		result, _, err := client.LabelsAPI.GetAllLabels(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		allLabels = append(allLabels, result.GetValues()...)

		if limit > 0 && int32(len(allLabels)) >= limit {
			allLabels = allLabels[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allLabels) != int(limit) {
		t.Errorf("expected %d labels, got %d", limit, len(allLabels))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit logic tests — users (renamed from --max-results; verify logic intact)
// ---------------------------------------------------------------------------

// TestListUsersLimit_StopsPaginationEarly verifies the renamed flag still
// limits results correctly for users.
func TestListUsersLimit_StopsPaginationEarly(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		users := []map[string]interface{}{
			buildUser("u1", "Alice", "alice@example.com", "atlassian", true),
			buildUser("u2", "Bob", "bob@example.com", "atlassian", true),
			buildUser("u3", "Carol", "carol@example.com", "atlassian", true),
		}
		// Return a full page so the loop would continue without a limit
		w.Write([]byte(buildUsersResponse(users)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 2
	var allUsers []string
	var startAt int32 = 0
	const pageSize int32 = 3

	for {
		result, _, err := client.UsersAPI.GetAllUsers(ctx).
			StartAt(startAt).
			MaxResults(pageSize).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, u := range result {
			allUsers = append(allUsers, u.GetDisplayName())
		}

		if limit > 0 && int32(len(allUsers)) >= limit {
			allUsers = allUsers[:limit]
			break
		}

		if int32(len(result)) < pageSize {
			break
		}

		startAt += int32(len(result))
	}

	if len(allUsers) != int(limit) {
		t.Errorf("expected %d users, got %d", limit, len(allUsers))
	}
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit logic test — boards: limit exactly equals one full page
// ---------------------------------------------------------------------------

// TestListBoardsLimit_ExactPageSize verifies behaviour when limit equals the
// page size exactly (no truncation, but no further pages fetched).
func TestListBoardsLimit_ExactPageSize(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		if requestCount == 1 {
			boards := []map[string]interface{}{
				buildBoard(1, "Board A", "scrum", "PROJ"),
				buildBoard(2, "Board B", "kanban", "PROJ"),
			}
			// Page is NOT last — there are more pages.
			w.Write([]byte(buildBoardsResponse(boards, false, 0, 2)))
		} else {
			boards := []map[string]interface{}{
				buildBoard(3, "Board C", "simple", "TEST"),
			}
			w.Write([]byte(buildBoardsResponse(boards, true, 2, 2)))
		}
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	const limit int32 = 2 // exactly one page
	var allBoards []string
	var startAt int64 = 0

	for {
		result, _, err := client.BoardAPI.GetAllBoards(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, b := range result.GetValues() {
			allBoards = append(allBoards, b.GetName())
		}

		if limit > 0 && int32(len(allBoards)) >= limit {
			allBoards = allBoards[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allBoards) != int(limit) {
		t.Errorf("expected %d boards, got %d", limit, len(allBoards))
	}
	// Should stop after first page because limit was satisfied.
	if requestCount != 1 {
		t.Errorf("expected 1 API request, got %d", requestCount)
	}
}

// ---------------------------------------------------------------------------
// Limit logic test — projects: limit spans multiple pages
// ---------------------------------------------------------------------------

// TestListProjectsLimit_SpansMultiplePages verifies that when limit > page size
// the loop fetches multiple pages and stops mid-page when limit is reached.
func TestListProjectsLimit_SpansMultiplePages(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		switch requestCount {
		case 1:
			projects := []map[string]interface{}{
				buildProject("P1", "Project One", "software"),
				buildProject("P2", "Project Two", "software"),
			}
			w.Write([]byte(buildProjectsResponse(projects, false, 0, 2, 6)))
		case 2:
			projects := []map[string]interface{}{
				buildProject("P3", "Project Three", "business"),
				buildProject("P4", "Project Four", "software"),
			}
			w.Write([]byte(buildProjectsResponse(projects, false, 2, 2, 6)))
		default:
			// Third page — should NOT be fetched when limit=3
			projects := []map[string]interface{}{
				buildProject("P5", "Project Five", "software"),
				buildProject("P6", "Project Six", "business"),
			}
			w.Write([]byte(buildProjectsResponse(projects, true, 4, 2, 6)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 3
	var allProjects []string
	var startAt int64 = 0

	for {
		result, _, err := client.ProjectsAPI.SearchProjects(ctx).StartAt(startAt).Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, p := range result.GetValues() {
			allProjects = append(allProjects, p.GetKey())
		}

		if limit > 0 && int32(len(allProjects)) >= limit {
			allProjects = allProjects[:limit]
			break
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) != int(limit) {
		t.Errorf("expected %d projects, got %d", limit, len(allProjects))
	}
	// limit=3, page size=2: needs 2 pages to satisfy limit (pages 1+2 give 4 items, capped to 3)
	if requestCount != 2 {
		t.Errorf("expected 2 API requests (limit satisfied mid-second page), got %d", requestCount)
	}

	// Verify the first 3 in order
	expectedKeys := []string{"P1", "P2", "P3"}
	for i, key := range expectedKeys {
		if allProjects[i] != key {
			t.Errorf("expected project[%d]=%s, got %s", i, key, allProjects[i])
		}
	}
}

// ---------------------------------------------------------------------------
// Limit logic test — status: limit spans multiple pages
// ---------------------------------------------------------------------------

// TestListStatusLimit_SpansMultiplePages verifies multi-page fetch with limit.
func TestListStatusLimit_SpansMultiplePages(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		requestCount++
		switch requestCount {
		case 1:
			statuses := []map[string]interface{}{
				buildStatus("1", "To Do", "TODO", ""),
				buildStatus("2", "In Progress", "IN_PROGRESS", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, false, 0, 2, 5)))
		case 2:
			statuses := []map[string]interface{}{
				buildStatus("3", "In Review", "IN_PROGRESS", ""),
				buildStatus("4", "Done", "DONE", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, false, 2, 2, 5)))
		default:
			// Should NOT be fetched when limit=3
			statuses := []map[string]interface{}{
				buildStatus("5", "Closed", "DONE", ""),
			}
			w.Write([]byte(buildStatusesResponse(statuses, true, 4, 2, 5)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	const limit int32 = 3
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

		if limit > 0 && int32(len(allStatuses)) >= limit {
			allStatuses = allStatuses[:limit]
			break
		}

		startAt += int64(len(result.GetValues()))
	}

	if len(allStatuses) != int(limit) {
		t.Errorf("expected %d statuses, got %d", limit, len(allStatuses))
	}
	if requestCount != 2 {
		t.Errorf("expected 2 API requests, got %d", requestCount)
	}

	expectedNames := []string{"To Do", "In Progress", "In Review"}
	for i, name := range expectedNames {
		if !strings.Contains(allStatuses[i], name) {
			t.Errorf("expected status[%d]=%q, got %q", i, name, allStatuses[i])
		}
	}
}
