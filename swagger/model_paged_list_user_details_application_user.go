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

// A paged list. To access additional details append `[start-index:end-index]` to the expand request. For example, `?expand=sharedUsers[10:40]` returns a list starting at item 10 and finishing at item 40.
type PagedListUserDetailsApplicationUser struct {
	// The index of the last item returned on the page.
	EndIndex int32 `json:"end-index,omitempty"`
	// The list of items.
	Items []UserDetails `json:"items,omitempty"`
	// The maximum number of results that could be on the page.
	MaxResults int32 `json:"max-results,omitempty"`
	// The number of items on the page.
	Size int32 `json:"size,omitempty"`
	// The index of the first item returned on the page.
	StartIndex int32 `json:"start-index,omitempty"`
}
