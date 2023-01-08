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

// The list of JQL queries to sanitize for the given account IDs.
type JqlQueriesToSanitize struct {
	// The list of JQL queries to sanitize. Must contain unique values. Maximum of 20 queries.
	Queries []JqlQueryToSanitize `json:"queries"`
}
