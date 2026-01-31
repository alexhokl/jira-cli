package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
	"github.com/alexhokl/jira-cli/swagger_software"
	"github.com/spf13/viper"
)

// timeFixTransport is a custom HTTP transport that fixes Jira's timestamp format
// by adding a colon to timezone offsets (e.g., +0800 -> +08:00)
type timeFixTransport struct {
	transport http.RoundTripper
}

// timezoneRegex matches timezone offsets without colons (e.g., +0800, -0530)
var timezoneRegex = regexp.MustCompile(`(\d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(?:\.\d+)?)([\+\-])(\d{2})(\d{2})`)

func (t *timeFixTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	transport := t.transport
	if transport == nil {
		transport = http.DefaultTransport
	}

	resp, err := transport.RoundTrip(req)
	if err != nil {
		return nil, err
	}

	// Only process JSON responses
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		return resp, nil
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	// Fix timezone format: +0800 -> +08:00
	fixedBody := timezoneRegex.ReplaceAll(body, []byte("${1}${2}${3}:${4}"))

	// Replace the response body with the fixed content
	resp.Body = io.NopCloser(bytes.NewReader(fixedBody))
	resp.ContentLength = int64(len(fixedBody))

	return resp, nil
}

func getAuthContext() context.Context {
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	auth := swagger.BasicAuth{
		UserName: email,
		Password: apiKey,
	}

	ctx := context.WithValue(context.Background(), swagger.ContextBasicAuth, auth)
	return ctx
}

func getConfiguration() *swagger.Configuration {
	configuration := swagger.NewConfiguration()
	configuration.Servers[0].URL = fmt.Sprintf("https://%s.atlassian.net", viper.GetString("organization"))
	configuration.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return configuration
}

func newClient() *swagger.APIClient {
	return swagger.NewAPIClient(getConfiguration())
}

func getSoftwareAuthContext() context.Context {
	email := viper.GetString("email")
	apiKey := viper.GetString("api_key")

	auth := swagger_software.BasicAuth{
		UserName: email,
		Password: apiKey,
	}

	ctx := context.WithValue(context.Background(), swagger_software.ContextBasicAuth, auth)
	return ctx
}

func getSoftwareConfiguration() *swagger_software.Configuration {
	configuration := swagger_software.NewConfiguration()
	configuration.Servers[0].URL = fmt.Sprintf("https://%s.atlassian.net", viper.GetString("organization"))
	configuration.HTTPClient = &http.Client{
		Transport: &timeFixTransport{},
	}
	return configuration
}

func newSoftwareClient() *swagger_software.APIClient {
	return swagger_software.NewAPIClient(getSoftwareConfiguration())
}

func getMessageFromEditor() (string, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		return "", fmt.Errorf("EDITOR environment variable is not set")
	}

	tmpFile, err := os.CreateTemp("", "jira-comment-*.txt")
	if err != nil {
		return "", fmt.Errorf("failed to create temporary file: %w", err)
	}
	tmpFileName := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpFileName)

	cmd := exec.Command(editor, tmpFileName)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to run editor: %w", err)
	}

	content, err := os.ReadFile(tmpFileName)
	if err != nil {
		return "", fmt.Errorf("failed to read temporary file: %w", err)
	}

	return strings.TrimSpace(string(content)), nil
}
