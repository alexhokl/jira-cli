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

// The tab for the screen.
type AllOfScreenWithTabTab struct {
	// The ID of the screen tab.
	Id int64 `json:"id,omitempty"`
	// The name of the screen tab. The maximum length is 255 characters.
	Name string `json:"name"`
}