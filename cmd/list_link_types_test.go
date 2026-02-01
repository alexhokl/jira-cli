package cmd

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListLinkTypesAPI_Success(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/rest/api/3/issueLinkType" {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if r.Method != http.MethodGet {
			t.Errorf("expected GET, got %s", r.Method)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := `{
			"issueLinkTypes": [
				{
					"id": "10000",
					"name": "Blocks",
					"inward": "is blocked by",
					"outward": "blocks"
				},
				{
					"id": "10001",
					"name": "Duplicates",
					"inward": "is duplicated by",
					"outward": "duplicates"
				},
				{
					"id": "10002",
					"name": "Relates",
					"inward": "relates to",
					"outward": "relates to"
				}
			]
		}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	linkTypes := result.GetIssueLinkTypes()
	if len(linkTypes) != 3 {
		t.Errorf("expected 3 link types, got %d", len(linkTypes))
	}

	// Verify first link type
	if linkTypes[0].GetName() != "Blocks" {
		t.Errorf("expected first link type name 'Blocks', got %s", linkTypes[0].GetName())
	}
	if linkTypes[0].GetInward() != "is blocked by" {
		t.Errorf("expected inward 'is blocked by', got %s", linkTypes[0].GetInward())
	}
	if linkTypes[0].GetOutward() != "blocks" {
		t.Errorf("expected outward 'blocks', got %s", linkTypes[0].GetOutward())
	}
}

func TestListLinkTypesAPI_EmptyResponse(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := `{"issueLinkTypes": []}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	linkTypes := result.GetIssueLinkTypes()
	if len(linkTypes) != 0 {
		t.Errorf("expected 0 link types, got %d", len(linkTypes))
	}
}

func TestListLinkTypesAPI_Unauthorized(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte(buildErrorResponse("Authentication required", http.StatusUnauthorized)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err == nil {
		t.Error("expected error for unauthorized, got nil")
	}
}

func TestListLinkTypesAPI_ServerError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(buildErrorResponse("Internal server error", http.StatusInternalServerError)))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	_, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err == nil {
		t.Error("expected error for server error, got nil")
	}
}

func TestListLinkTypesAPI_VerifyAllFields(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		response := `{
			"issueLinkTypes": [
				{
					"id": "10003",
					"name": "Clones",
					"inward": "is cloned by",
					"outward": "clones",
					"self": "https://test.atlassian.net/rest/api/3/issueLinkType/10003"
				}
			]
		}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	client := newMockClientWithServer(server.URL)
	ctx := mockAuthContext()

	result, _, err := client.IssueLinkTypesAPI.GetIssueLinkTypes(ctx).Execute()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	linkTypes := result.GetIssueLinkTypes()
	if len(linkTypes) != 1 {
		t.Fatalf("expected 1 link type, got %d", len(linkTypes))
	}

	lt := linkTypes[0]
	if lt.GetId() != "10003" {
		t.Errorf("expected id '10003', got %s", lt.GetId())
	}
	if lt.GetName() != "Clones" {
		t.Errorf("expected name 'Clones', got %s", lt.GetName())
	}
	if lt.GetInward() != "is cloned by" {
		t.Errorf("expected inward 'is cloned by', got %s", lt.GetInward())
	}
	if lt.GetOutward() != "clones" {
		t.Errorf("expected outward 'clones', got %s", lt.GetOutward())
	}
}
