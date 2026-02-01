package cmd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// buildUsersResponse creates a JSON response for users API (returns array directly)
func buildUsersResponse(users []map[string]interface{}) string {
	data, _ := json.Marshal(users)
	return string(data)
}

// buildUser creates a user object for testing
func buildUser(accountId, displayName, email, accountType string, active bool) map[string]interface{} {
	user := map[string]interface{}{
		"accountId":    accountId,
		"displayName":  displayName,
		"emailAddress": email,
		"accountType":  accountType,
		"active":       active,
		"self":         "https://test.atlassian.net/rest/api/3/user?accountId=" + accountId,
	}
	return user
}

func TestUsersAPI_SinglePage(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Expected GET request, got %s", r.Method)
		}
		if !strings.Contains(r.URL.Path, "/users") {
			t.Errorf("Unexpected path: %s", r.URL.Path)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		users := []map[string]interface{}{
			buildUser("user1", "John Doe", "john@example.com", "atlassian", true),
			buildUser("user2", "Jane Smith", "jane@example.com", "atlassian", true),
			buildUser("user3", "Bob Wilson", "bob@example.com", "atlassian", false),
		}
		w.Write([]byte(buildUsersResponse(users)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 users, got %d", len(result))
	}
}

func TestUsersAPI_Pagination(t *testing.T) {
	requestCount := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		startAt := r.URL.Query().Get("startAt")
		requestCount++

		if startAt == "" || startAt == "0" {
			users := []map[string]interface{}{
				buildUser("user1", "John Doe", "john@example.com", "atlassian", true),
				buildUser("user2", "Jane Smith", "jane@example.com", "atlassian", true),
			}
			w.Write([]byte(buildUsersResponse(users)))
		} else {
			users := []map[string]interface{}{
				buildUser("user3", "Bob Wilson", "bob@example.com", "atlassian", false),
			}
			w.Write([]byte(buildUsersResponse(users)))
		}
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	// First page
	result1, _, err := client.UsersAPI.GetAllUsers(ctx).StartAt(0).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("First page failed: %v", err)
	}
	if len(result1) != 2 {
		t.Errorf("Expected 2 users on first page, got %d", len(result1))
	}

	// Second page
	result2, _, err := client.UsersAPI.GetAllUsers(ctx).StartAt(2).MaxResults(2).Execute()
	if err != nil {
		t.Fatalf("Second page failed: %v", err)
	}
	if len(result2) != 1 {
		t.Errorf("Expected 1 user on second page, got %d", len(result2))
	}

	if requestCount != 2 {
		t.Errorf("Expected 2 requests, got %d", requestCount)
	}
}

func TestUsersAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildUsersResponse([]map[string]interface{}{})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 0 {
		t.Errorf("Expected 0 users, got %d", len(result))
	}
}

func TestUsersAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Unauthorized", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err == nil {
		t.Error("Expected error for unauthorized request")
	}
}

func TestUsersAPI_Forbidden(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte(buildErrorResponse("You do not have permission to browse users", http.StatusForbidden)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err == nil {
		t.Error("Expected error for forbidden request")
	}
}

func TestUsersAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err == nil {
		t.Error("Expected error for server error")
	}
}

func TestUsersAPI_WithDifferentAccountTypes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		users := []map[string]interface{}{
			buildUser("user1", "Regular User", "user@example.com", "atlassian", true),
			buildUser("app1", "Service Account", "", "app", true),
			buildUser("customer1", "Customer User", "customer@example.com", "customer", true),
		}
		w.Write([]byte(buildUsersResponse(users)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 3 {
		t.Errorf("Expected 3 users, got %d", len(result))
	}

	// Verify account types
	accountTypes := make(map[string]bool)
	for _, user := range result {
		accountTypes[user.GetAccountType()] = true
	}
	if !accountTypes["atlassian"] || !accountTypes["app"] || !accountTypes["customer"] {
		t.Error("Missing expected account types")
	}
}

func TestUsersAPI_VerifyRequestParameters(t *testing.T) {
	var capturedStartAt, capturedMaxResults string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		capturedStartAt = r.URL.Query().Get("startAt")
		capturedMaxResults = r.URL.Query().Get("maxResults")

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(buildUsersResponse([]map[string]interface{}{})))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.UsersAPI.GetAllUsers(ctx).StartAt(50).MaxResults(25).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}

	if capturedStartAt != "50" {
		t.Errorf("Expected startAt=50, got %s", capturedStartAt)
	}
	if capturedMaxResults != "25" {
		t.Errorf("Expected maxResults=25, got %s", capturedMaxResults)
	}
}

func TestUsersAPI_LargeDataset(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		users := make([]map[string]interface{}, 100)
		for i := 0; i < 100; i++ {
			active := i%10 != 0 // 10% inactive
			accountType := "atlassian"
			if i%20 == 0 {
				accountType = "app"
			}
			users[i] = buildUser(
				fmt.Sprintf("user%d", i+1),
				fmt.Sprintf("User %d", i+1),
				fmt.Sprintf("user%d@example.com", i+1),
				accountType,
				active,
			)
		}
		w.Write([]byte(buildUsersResponse(users)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, resp, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", resp.StatusCode)
	}
	if len(result) != 100 {
		t.Errorf("Expected 100 users, got %d", len(result))
	}
}

func TestUsersAPI_ActiveInactiveUsers(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		users := []map[string]interface{}{
			buildUser("user1", "Active User 1", "active1@example.com", "atlassian", true),
			buildUser("user2", "Active User 2", "active2@example.com", "atlassian", true),
			buildUser("user3", "Inactive User", "inactive@example.com", "atlassian", false),
		}
		w.Write([]byte(buildUsersResponse(users)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.UsersAPI.GetAllUsers(ctx).Execute()
	if err != nil {
		t.Fatalf("GetAllUsers failed: %v", err)
	}

	activeCount := 0
	inactiveCount := 0
	for _, user := range result {
		if user.GetActive() {
			activeCount++
		} else {
			inactiveCount++
		}
	}

	if activeCount != 2 {
		t.Errorf("Expected 2 active users, got %d", activeCount)
	}
	if inactiveCount != 1 {
		t.Errorf("Expected 1 inactive user, got %d", inactiveCount)
	}
}
