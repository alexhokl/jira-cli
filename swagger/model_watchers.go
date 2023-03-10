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

// The details of watchers on an issue.
type Watchers struct {
	// Whether the calling user is watching this issue.
	IsWatching bool `json:"isWatching,omitempty"`
	// The URL of these issue watcher details.
	Self string `json:"self,omitempty"`
	// The number of users watching this issue.
	WatchCount int32 `json:"watchCount,omitempty"`
	// Details of the users watching this issue.
	Watchers []UserDetails `json:"watchers,omitempty"`
}
