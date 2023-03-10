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

// Details of the settings for a dashboard gadget.
type DashboardGadgetSettings struct {
	// The color of the gadget. Should be one of `blue`, `red`, `yellow`, `green`, `cyan`, `purple`, `gray`, or `white`.
	Color string `json:"color,omitempty"`
	// Whether to ignore the validation of module key and URI. For example, when a gadget is created that is a part of an application that isn't installed.
	IgnoreUriAndModuleKeyValidation bool `json:"ignoreUriAndModuleKeyValidation,omitempty"`
	// The module key of the gadget type. Can't be provided with `uri`.
	ModuleKey string `json:"moduleKey,omitempty"`
	// The position of the gadget. When the gadget is placed into the position, other gadgets in the same column are moved down to accommodate it.
	Position *AllOfDashboardGadgetSettingsPosition `json:"position,omitempty"`
	// The title of the gadget.
	Title string `json:"title,omitempty"`
	// The URI of the gadget type. Can't be provided with `moduleKey`.
	Uri string `json:"uri,omitempty"`
}
