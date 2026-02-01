package cmd

import (
	"testing"

	"github.com/alexhokl/jira-cli/swagger"
)

func TestGetAuthorDisplayName(t *testing.T) {
	tests := []struct {
		name     string
		author   *swagger.UserDetails
		expected string
	}{
		{
			name:     "nil author",
			author:   nil,
			expected: "Unknown",
		},
		{
			name: "author with display name",
			author: func() *swagger.UserDetails {
				u := swagger.NewUserDetails()
				u.SetDisplayName("John Doe")
				u.SetEmailAddress("john@example.com")
				u.SetAccountId("abc123")
				return u
			}(),
			expected: "John Doe",
		},
		{
			name: "author with email only",
			author: func() *swagger.UserDetails {
				u := swagger.NewUserDetails()
				u.SetEmailAddress("jane@example.com")
				u.SetAccountId("xyz789")
				return u
			}(),
			expected: "jane@example.com",
		},
		{
			name: "author with account ID only",
			author: func() *swagger.UserDetails {
				u := swagger.NewUserDetails()
				u.SetAccountId("account-id-only")
				return u
			}(),
			expected: "account-id-only",
		},
		{
			name: "author with empty display name returns empty (HasDisplayName is true)",
			author: func() *swagger.UserDetails {
				u := swagger.NewUserDetails()
				u.SetDisplayName("") // empty string - HasDisplayName() returns true
				u.SetEmailAddress("fallback@example.com")
				return u
			}(),
			expected: "", // HasDisplayName() returns true so GetDisplayName() is used
		},
		{
			name:     "empty author struct",
			author:   swagger.NewUserDetails(),
			expected: "Unknown",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getAuthorDisplayName(tt.author)
			if result != tt.expected {
				t.Errorf("getAuthorDisplayName() = %q, want %q", result, tt.expected)
			}
		})
	}
}

func TestGetAuthorDisplayNamePriority(t *testing.T) {
	// Test that display name takes priority over email over account ID
	t.Run("display name has highest priority", func(t *testing.T) {
		u := swagger.NewUserDetails()
		u.SetDisplayName("Display Name")
		u.SetEmailAddress("email@example.com")
		u.SetAccountId("account-123")

		result := getAuthorDisplayName(u)
		if result != "Display Name" {
			t.Errorf("Expected display name to take priority, got %q", result)
		}
	})

	t.Run("email has second priority", func(t *testing.T) {
		u := swagger.NewUserDetails()
		u.SetEmailAddress("email@example.com")
		u.SetAccountId("account-123")

		result := getAuthorDisplayName(u)
		if result != "email@example.com" {
			t.Errorf("Expected email to take priority over account ID, got %q", result)
		}
	})

	t.Run("account ID is last resort", func(t *testing.T) {
		u := swagger.NewUserDetails()
		u.SetAccountId("account-123")

		result := getAuthorDisplayName(u)
		if result != "account-123" {
			t.Errorf("Expected account ID as fallback, got %q", result)
		}
	})
}
