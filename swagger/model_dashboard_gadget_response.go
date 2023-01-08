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

// The list of gadgets on the dashboard.
type DashboardGadgetResponse struct {
	// The list of gadgets.
	Gadgets []DashboardGadget `json:"gadgets"`
}
