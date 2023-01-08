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

// A screen with tab details.
type ScreenWithTab struct {
	// The description of the screen.
	Description string `json:"description,omitempty"`
	// The ID of the screen.
	Id int64 `json:"id,omitempty"`
	// The name of the screen.
	Name string `json:"name,omitempty"`
	// The scope of the screen.
	Scope *AllOfScreenWithTabScope `json:"scope,omitempty"`
	// The tab for the screen.
	Tab *AllOfScreenWithTabTab `json:"tab,omitempty"`
}
