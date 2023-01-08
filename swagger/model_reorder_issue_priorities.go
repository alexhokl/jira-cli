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

// Change the order of issue priorities.
type ReorderIssuePriorities struct {
	// The ID of the priority. Required if `position` isn't provided.
	After string `json:"after,omitempty"`
	// The list of issue IDs to be reordered. Cannot contain duplicates nor after ID.
	Ids []string `json:"ids"`
	// The position for issue priorities to be moved to. Required if `after` isn't provided.
	Position string `json:"position,omitempty"`
}