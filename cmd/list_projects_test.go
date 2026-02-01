package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRunListProjects_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify the request path
		if r.URL.Path != "/rest/api/3/project/search" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		projects := []map[string]interface{}{
			buildProject("PROJ1", "Project One", "software"),
			buildProject("PROJ2", "Project Two", "business"),
			buildProject("PROJ3", "Project Three", "service_desk"),
		}

		response := buildProjectsResponse(projects, true, 0, 50, 3)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).
		StartAt(0).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 3 {
		t.Errorf("expected 3 projects, got %d", len(projects))
	}

	// Verify first project
	if projects[0].GetKey() != "PROJ1" {
		t.Errorf("expected first project key to be PROJ1, got %s", projects[0].GetKey())
	}
	if projects[0].GetName() != "Project One" {
		t.Errorf("expected first project name to be 'Project One', got %s", projects[0].GetName())
	}
	if projects[0].GetProjectTypeKey() != "software" {
		t.Errorf("expected first project type to be 'software', got %s", projects[0].GetProjectTypeKey())
	}
}

func TestRunListProjects_MultiplePagesWithPagination(t *testing.T) {
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			projects := []map[string]interface{}{
				buildProject("PROJ1", "Project One", "software"),
				buildProject("PROJ2", "Project Two", "business"),
			}
			response = buildProjectsResponse(projects, false, 0, 2, 4)
		} else {
			projects := []map[string]interface{}{
				buildProject("PROJ3", "Project Three", "service_desk"),
				buildProject("PROJ4", "Project Four", "software"),
			}
			response = buildProjectsResponse(projects, true, 2, 2, 4)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// Simulate pagination loop from runListProjects
	var allProjects []string
	var startAt int64 = 0

	for {
		result, _, err := client.ProjectsAPI.SearchProjects(ctx).
			StartAt(startAt).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, p := range result.GetValues() {
			allProjects = append(allProjects, p.GetKey())
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) != 4 {
		t.Errorf("expected 4 projects, got %d", len(allProjects))
	}

	expectedKeys := []string{"PROJ1", "PROJ2", "PROJ3", "PROJ4"}
	for i, key := range allProjects {
		if key != expectedKeys[i] {
			t.Errorf("expected project %d key to be %s, got %s", i, expectedKeys[i], key)
		}
	}

	if requestCount != 2 {
		t.Errorf("expected 2 requests, got %d", requestCount)
	}
}

func TestRunListProjects_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildProjectsResponse([]map[string]interface{}{}, true, 0, 50, 0)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).
		StartAt(0).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 0 {
		t.Errorf("expected 0 projects, got %d", len(projects))
	}
}

func TestRunListProjects_WithQueryFilter(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameter
		query := r.URL.Query().Get("query")
		if query != "test" {
			t.Errorf("expected query=test, got query=%s", query)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		projects := []map[string]interface{}{
			buildProject("TEST", "Test Project", "software"),
		}
		response := buildProjectsResponse(projects, true, 0, 50, 1)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).
		Query("test").
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 1 {
		t.Errorf("expected 1 project, got %d", len(projects))
	}

	if projects[0].GetKey() != "TEST" {
		t.Errorf("expected project key TEST, got %s", projects[0].GetKey())
	}
}

func TestRunListProjects_WithOrderBy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify orderBy parameter
		orderBy := r.URL.Query().Get("orderBy")
		if orderBy != "name" {
			t.Errorf("expected orderBy=name, got orderBy=%s", orderBy)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		// Projects ordered by name
		projects := []map[string]interface{}{
			buildProject("ALPHA", "Alpha Project", "software"),
			buildProject("BETA", "Beta Project", "software"),
			buildProject("GAMMA", "Gamma Project", "business"),
		}
		response := buildProjectsResponse(projects, true, 0, 50, 3)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).
		OrderBy("name").
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 3 {
		t.Errorf("expected 3 projects, got %d", len(projects))
	}

	// Verify ordering
	expectedKeys := []string{"ALPHA", "BETA", "GAMMA"}
	for i, key := range expectedKeys {
		if projects[i].GetKey() != key {
			t.Errorf("expected project %d to be %s, got %s", i, key, projects[i].GetKey())
		}
	}
}

func TestRunListProjects_WithQueryAndOrderBy(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify both parameters
		query := r.URL.Query().Get("query")
		orderBy := r.URL.Query().Get("orderBy")

		if query != "proj" {
			t.Errorf("expected query=proj, got query=%s", query)
		}
		if orderBy != "-key" {
			t.Errorf("expected orderBy=-key, got orderBy=%s", orderBy)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		projects := []map[string]interface{}{
			buildProject("PROJ3", "Project Three", "software"),
			buildProject("PROJ2", "Project Two", "software"),
			buildProject("PROJ1", "Project One", "business"),
		}
		response := buildProjectsResponse(projects, true, 0, 50, 3)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).
		Query("proj").
		OrderBy("-key").
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 3 {
		t.Errorf("expected 3 projects, got %d", len(projects))
	}
}

func TestRunListProjects_APIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.ProjectsAPI.SearchProjects(ctx).Execute()

	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestRunListProjects_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.ProjectsAPI.SearchProjects(ctx).Execute()

	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestRunListProjects_ForbiddenAccess(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(buildErrorResponse("You do not have permission to access this resource", http.StatusForbidden)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.ProjectsAPI.SearchProjects(ctx).Execute()

	if err == nil {
		t.Error("expected error for forbidden access, got nil")
	}
}

func TestRunListProjects_LargeDataset(t *testing.T) {
	totalProjects := 150
	pageSize := 50
	currentPage := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		start := currentPage * pageSize
		end := start + pageSize
		if end > totalProjects {
			end = totalProjects
		}

		projects := make([]map[string]interface{}, end-start)
		for i := range projects {
			idx := start + i
			projects[i] = buildProject(
				"PROJ"+string(rune('A'+idx%26)),
				"Project "+string(rune('A'+idx%26)),
				"software",
			)
		}

		isLast := end >= totalProjects
		response := buildProjectsResponse(projects, isLast, int64(start), int32(pageSize), int64(totalProjects))
		currentPage++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	var allProjects []string
	var startAt int64 = 0

	for {
		result, _, err := client.ProjectsAPI.SearchProjects(ctx).
			StartAt(startAt).
			Execute()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		for _, p := range result.GetValues() {
			allProjects = append(allProjects, p.GetKey())
		}

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allProjects) != totalProjects {
		t.Errorf("expected %d projects, got %d", totalProjects, len(allProjects))
	}
}

func TestRunListProjects_VerifyProjectFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		projects := []map[string]interface{}{
			{
				"key":            "TEST",
				"name":           "Test Project",
				"projectTypeKey": "software",
				"id":             "12345",
				"self":           "https://test.atlassian.net/rest/api/3/project/12345",
				"style":          "classic",
				"simplified":     false,
			},
		}
		response := buildProjectsResponse(projects, true, 0, 50, 1)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.ProjectsAPI.SearchProjects(ctx).Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	projects := result.GetValues()
	if len(projects) != 1 {
		t.Fatalf("expected 1 project, got %d", len(projects))
	}

	project := projects[0]
	if project.GetKey() != "TEST" {
		t.Errorf("expected key TEST, got %s", project.GetKey())
	}
	if project.GetName() != "Test Project" {
		t.Errorf("expected name 'Test Project', got %s", project.GetName())
	}
	if project.GetProjectTypeKey() != "software" {
		t.Errorf("expected projectTypeKey 'software', got %s", project.GetProjectTypeKey())
	}
}

func TestRunListProjects_RequestVerification(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify HTTP method
		if r.Method != http.MethodGet {
			t.Errorf("expected GET request, got %s", r.Method)
		}

		// Verify path
		if r.URL.Path != "/rest/api/3/project/search" {
			t.Errorf("expected path /rest/api/3/project/search, got %s", r.URL.Path)
		}

		// Verify startAt parameter
		startAt := r.URL.Query().Get("startAt")
		if startAt != "10" {
			t.Errorf("expected startAt=10, got startAt=%s", startAt)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildProjectsResponse([]map[string]interface{}{}, true, 10, 50, 0)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.ProjectsAPI.SearchProjects(ctx).
		StartAt(10).
		Execute()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}
