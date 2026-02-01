package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListBoardsAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/agile/1.0/board" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		boards := []map[string]interface{}{
			buildBoard(1, "Sprint Board", "scrum", "PROJ"),
			buildBoard(2, "Kanban Board", "kanban", "PROJ"),
			buildBoard(3, "Simple Board", "simple", "TEST"),
		}
		response := buildBoardsResponse(boards, true, 0, 50)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	result, _, err := client.BoardAPI.GetAllBoards(ctx).StartAt(0).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	boards := result.GetValues()
	if len(boards) != 3 {
		t.Errorf("expected 3 boards, got %d", len(boards))
	}

	// Verify first board
	if boards[0].GetId() != 1 {
		t.Errorf("expected first board id 1, got %d", boards[0].GetId())
	}
	if boards[0].GetName() != "Sprint Board" {
		t.Errorf("expected first board name 'Sprint Board', got %s", boards[0].GetName())
	}
	if boards[0].GetType() != "scrum" {
		t.Errorf("expected first board type 'scrum', got %s", boards[0].GetType())
	}
}

func TestListBoardsAPI_WithFilters(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify query parameters
		name := r.URL.Query().Get("name")
		if name != "Sprint" {
			t.Errorf("expected name=Sprint, got name=%s", name)
		}

		projectKeyOrId := r.URL.Query().Get("projectKeyOrId")
		if projectKeyOrId != "PROJ" {
			t.Errorf("expected projectKeyOrId=PROJ, got projectKeyOrId=%s", projectKeyOrId)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		boards := []map[string]interface{}{
			buildBoard(1, "Sprint Board", "scrum", "PROJ"),
		}
		response := buildBoardsResponse(boards, true, 0, 50)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	result, _, err := client.BoardAPI.GetAllBoards(ctx).
		Name("Sprint").
		ProjectKeyOrId("PROJ").
		Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	boards := result.GetValues()
	if len(boards) != 1 {
		t.Errorf("expected 1 board, got %d", len(boards))
	}
}

func TestListBoardsAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		response := buildBoardsResponse([]map[string]interface{}{}, true, 0, 50)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	result, _, err := client.BoardAPI.GetAllBoards(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	boards := result.GetValues()
	if len(boards) != 0 {
		t.Errorf("expected 0 boards, got %d", len(boards))
	}
}

func TestListBoardsAPI_Pagination(t *testing.T) {
	requestCount := 0

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		var response string
		if requestCount == 0 {
			boards := []map[string]interface{}{
				buildBoard(1, "Board 1", "scrum", "PROJ"),
				buildBoard(2, "Board 2", "kanban", "PROJ"),
			}
			response = buildBoardsResponse(boards, false, 0, 2)
		} else {
			boards := []map[string]interface{}{
				buildBoard(3, "Board 3", "simple", "TEST"),
			}
			response = buildBoardsResponse(boards, true, 2, 2)
		}
		requestCount++
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	// Simulate pagination loop
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

		if result.GetIsLast() {
			break
		}

		startAt = result.GetStartAt() + int64(result.GetMaxResults())
	}

	if len(allBoards) != 3 {
		t.Errorf("expected 3 boards, got %d", len(allBoards))
	}

	if requestCount != 2 {
		t.Errorf("expected 2 requests, got %d", requestCount)
	}
}

func TestListBoardsAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, _, err := client.BoardAPI.GetAllBoards(ctx).Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestListBoardsAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	_, _, err := client.BoardAPI.GetAllBoards(ctx).Execute()
	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestListBoardsAPI_WithLocation(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		boards := []map[string]interface{}{
			{
				"id":   int64(1),
				"name": "Project Board",
				"type": "scrum",
				"location": map[string]interface{}{
					"projectId":   int64(10000),
					"projectKey":  "PROJ",
					"projectName": "Project Name",
				},
			},
		}
		response := buildBoardsResponse(boards, true, 0, 50)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	result, _, err := client.BoardAPI.GetAllBoards(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	boards := result.GetValues()
	if len(boards) != 1 {
		t.Fatalf("expected 1 board, got %d", len(boards))
	}

	board := boards[0]
	if !board.HasLocation() {
		t.Error("expected board to have location")
	}

	location := board.GetLocation()
	if location.GetProjectKey() != "PROJ" {
		t.Errorf("expected project key 'PROJ', got %s", location.GetProjectKey())
	}
}

func TestListBoardsAPI_OrderByName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify orderBy parameter
		orderBy := r.URL.Query().Get("orderBy")
		if orderBy != "name" {
			t.Errorf("expected orderBy=name, got orderBy=%s", orderBy)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		boards := []map[string]interface{}{
			buildBoard(1, "Alpha Board", "scrum", "PROJ"),
			buildBoard(2, "Beta Board", "kanban", "PROJ"),
			buildBoard(3, "Gamma Board", "simple", "TEST"),
		}
		response := buildBoardsResponse(boards, true, 0, 50)
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockSoftwareClientWithServer(server.URL)
	ctx := mockSoftwareAuthContext()

	result, _, err := client.BoardAPI.GetAllBoards(ctx).OrderBy("name").Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	boards := result.GetValues()
	if len(boards) != 3 {
		t.Errorf("expected 3 boards, got %d", len(boards))
	}

	// Verify order
	expectedNames := []string{"Alpha Board", "Beta Board", "Gamma Board"}
	for i, name := range expectedNames {
		if boards[i].GetName() != name {
			t.Errorf("expected board %d to be '%s', got '%s'", i, name, boards[i].GetName())
		}
	}
}
