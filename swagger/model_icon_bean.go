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

// An icon.
type IconBean struct {
	// The URL of the tooltip, used only for a status icon.
	Link string `json:"link,omitempty"`
	// The title of the icon, for use as a tooltip on the icon.
	Title string `json:"title,omitempty"`
	// The URL of a 16x16 pixel icon.
	Url16x16 string `json:"url16x16,omitempty"`
}
