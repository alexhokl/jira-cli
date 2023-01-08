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

// The description of the page of issues loaded by the provided JQL query.
type IssuesJqlMetaDataBean struct {
	// The number of issues that were loaded in this evaluation.
	Count int32 `json:"count"`
	// The maximum number of issues that could be loaded in this evaluation.
	MaxResults int32 `json:"maxResults"`
	// The index of the first issue.
	StartAt int64 `json:"startAt"`
	// The total number of issues the JQL returned.
	TotalCount int64 `json:"totalCount"`
	// Any warnings related to the JQL query. Present only if the validation mode was set to `warn`.
	ValidationWarnings []string `json:"validationWarnings,omitempty"`
}