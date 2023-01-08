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

// A paginated list of the users that are members of the group. A maximum of 50 users is returned in the list, to access additional users append `[start-index:end-index]` to the expand request. For example, to access the next 50 users, use`?expand=users[51:100]`.
type AllOfGroupUsers struct {
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
