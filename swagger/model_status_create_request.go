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

// Details of the statuses being created and their scope.
type StatusCreateRequest struct {
	Scope *StatusScope `json:"scope"`
	// Details of the statuses being created.
	Statuses []StatusCreate `json:"statuses"`
}