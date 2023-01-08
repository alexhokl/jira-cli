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

// The details of the gadget to update.
type DashboardGadgetUpdateRequest struct {
	// The color of the gadget. Should be one of `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, or `white`.
	Color string `json:"color,omitempty"`
	// The position of the gadget.
	Position *AllOfDashboardGadgetUpdateRequestPosition `json:"position,omitempty"`
	// The title of the gadget.
	Title string `json:"title,omitempty"`
}
