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

// The IDs of the screens for the screen types of the screen scheme. Only screens used in classic projects are accepted.
type AllOfScreenSchemeDetailsScreens struct {
	// The ID of the create screen.
	Create int64 `json:"create,omitempty"`
	// The ID of the default screen. Required when creating a screen scheme.
	Default_ int64 `json:"default,omitempty"`
	// The ID of the edit screen.
	Edit int64 `json:"edit,omitempty"`
	// The ID of the view screen.
	View int64 `json:"view,omitempty"`
}
