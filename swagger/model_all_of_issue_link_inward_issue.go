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

// Provides details about the linked issue. If presenting this link in a user interface, use the `inward` field of the issue link type to label the link.
type AllOfIssueLinkInwardIssue struct {
	// The fields associated with the issue.
	Fields *Object `json:"fields,omitempty"`
	// The ID of an issue. Required if `key` isn't provided.
	Id string `json:"id,omitempty"`
	// The key of an issue. Required if `id` isn't provided.
	Key string `json:"key,omitempty"`
	// The URL of the issue.
	Self string `json:"self,omitempty"`
}
