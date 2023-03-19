package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/alexhokl/jira-cli/swagger"
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
