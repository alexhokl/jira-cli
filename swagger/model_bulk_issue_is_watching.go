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

// A container for the watch status of a list of issues.
type BulkIssueIsWatching struct {
	// The map of issue ID to boolean watch status.
	IssuesIsWatching map[string]bool `json:"issuesIsWatching,omitempty"`
}
