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

// Details of an item associated with the changed record.
type AssociatedItemBean struct {
	// The ID of the associated record.
	Id string `json:"id,omitempty"`
	// The name of the associated record.
	Name string `json:"name,omitempty"`
	// The ID of the associated parent record.
	ParentId string `json:"parentId,omitempty"`
	// The name of the associated parent record.
	ParentName string `json:"parentName,omitempty"`
	// The type of the associated record.
	TypeName string `json:"typeName,omitempty"`
}
