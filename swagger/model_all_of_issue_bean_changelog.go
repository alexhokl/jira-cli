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

// Details of changelogs associated with the issue.
type AllOfIssueBeanChangelog struct {
	// The list of changelogs.
	Histories []Changelog `json:"histories,omitempty"`
	// The maximum number of results that could be on the page.
	MaxResults int32 `json:"maxResults,omitempty"`
	// The index of the first item returned on the page.
	StartAt int32 `json:"startAt,omitempty"`
	// The number of results on the page.
	Total int32 `json:"total,omitempty"`
}
