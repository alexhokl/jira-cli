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

// A page containing dashboard details.
type PageOfDashboards struct {
	// List of dashboards.
	Dashboards []Dashboard `json:"dashboards,omitempty"`
	// The maximum number of results that could be on the page.
	MaxResults int32 `json:"maxResults,omitempty"`
	// The URL of the next page of results, if any.
	Next string `json:"next,omitempty"`
	// The URL of the previous page of results, if any.
	Prev string `json:"prev,omitempty"`
	// The index of the first item returned on the page.
	StartAt int32 `json:"startAt,omitempty"`
	// The number of results on the page.
	Total int32 `json:"total,omitempty"`
}
