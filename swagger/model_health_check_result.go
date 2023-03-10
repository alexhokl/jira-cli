/*
 * The Jira Cloud platform REST API
 *
 * Jira Cloud platform REST API documentation
 *
 * API version: 1001.0.0-SNAPSHOT
 * Contact: ecosystem@atlassian.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package swagger

// Jira instance health check results. Deprecated and no longer returned.
type HealthCheckResult struct {
	// The description of the Jira health check item.
	Description string `json:"description,omitempty"`
	// The name of the Jira health check item.
	Name string `json:"name,omitempty"`
	// Whether the Jira health check item passed or failed.
	Passed bool `json:"passed,omitempty"`
}
