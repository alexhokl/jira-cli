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

// An [issue](https://developer.atlassian.com/cloud/jira/platform/jira-expressions-type-reference#issue) specified by ID or key. All the fields of the issue object are available in the Jira expression.
type IssueContextVariable struct {
	// The issue ID.
	Id int64 `json:"id,omitempty"`
	// The issue key.
	Key string `json:"key,omitempty"`
	// Type of custom context variable.
	Type_ string `json:"type"`
}
