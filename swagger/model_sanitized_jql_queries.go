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

// The sanitized JQL queries for the given account IDs.
type SanitizedJqlQueries struct {
	// The list of sanitized JQL queries.
	Queries []SanitizedJqlQuery `json:"queries,omitempty"`
}