package cmd

import (
	"testing"

	"github.com/spf13/viper"
)

func TestValidateConfiguration(t *testing.T) {
	// Save original viper state and restore after tests
	originalEmail := viper.GetString("email")
	originalApiKey := viper.GetString("api_key")
	originalOrg := viper.GetString("organization")
	defer func() {
		viper.Set("email", originalEmail)
		viper.Set("api_key", originalApiKey)
		viper.Set("organization", originalOrg)
	}()

	tests := []struct {
		name        string
		email       string
		apiKey      string
		org         string
		expectError bool
		errorMsg    string
	}{
		{
			name:        "all fields configured",
			email:       "user@example.com",
			apiKey:      "test-api-key",
			org:         "my-org",
			expectError: false,
		},
		{
			name:        "missing email",
			email:       "",
			apiKey:      "test-api-key",
			org:         "my-org",
			expectError: true,
			errorMsg:    "email is not configured",
		},
		{
			name:        "missing api_key",
			email:       "user@example.com",
			apiKey:      "",
			org:         "my-org",
			expectError: true,
			errorMsg:    "api_key is not configured",
		},
		{
			name:        "missing organization",
			email:       "user@example.com",
			apiKey:      "test-api-key",
			org:         "",
			expectError: true,
			errorMsg:    "organization is not configured",
		},
		{
			name:        "all fields missing",
			email:       "",
			apiKey:      "",
			org:         "",
			expectError: true,
			errorMsg:    "email is not configured", // First check fails
		},
		{
			name:        "email and api_key missing",
			email:       "",
			apiKey:      "",
			org:         "my-org",
			expectError: true,
			errorMsg:    "email is not configured",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set up viper configuration
			viper.Set("email", tt.email)
			viper.Set("api_key", tt.apiKey)
			viper.Set("organization", tt.org)

			err := validateConfiguration(nil, nil)

			if tt.expectError {
				if err == nil {
					t.Errorf("validateConfiguration() expected error, got nil")
					return
				}
				if err.Error() != tt.errorMsg {
					t.Errorf("validateConfiguration() error = %q, want %q", err.Error(), tt.errorMsg)
				}
			} else {
				if err != nil {
					t.Errorf("validateConfiguration() unexpected error: %v", err)
				}
			}
		})
	}
}

func TestValidateConfigurationOrder(t *testing.T) {
	// Test that validation checks in correct order: email -> api_key -> organization
	originalEmail := viper.GetString("email")
	originalApiKey := viper.GetString("api_key")
	originalOrg := viper.GetString("organization")
	defer func() {
		viper.Set("email", originalEmail)
		viper.Set("api_key", originalApiKey)
		viper.Set("organization", originalOrg)
	}()

	// When all are missing, email error should come first
	viper.Set("email", "")
	viper.Set("api_key", "")
	viper.Set("organization", "")

	err := validateConfiguration(nil, nil)
	if err == nil || err.Error() != "email is not configured" {
		t.Errorf("Expected email error first, got: %v", err)
	}

	// When email is set, api_key error should come second
	viper.Set("email", "user@example.com")

	err = validateConfiguration(nil, nil)
	if err == nil || err.Error() != "api_key is not configured" {
		t.Errorf("Expected api_key error second, got: %v", err)
	}

	// When email and api_key are set, organization error should come third
	viper.Set("api_key", "test-key")

	err = validateConfiguration(nil, nil)
	if err == nil || err.Error() != "organization is not configured" {
		t.Errorf("Expected organization error third, got: %v", err)
	}
}
