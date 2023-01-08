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

type AllOfRoleActorActorGroup struct {
	// The display name of the group.
	DisplayName string `json:"displayName,omitempty"`
	// The ID of the group.
	GroupId string `json:"groupId,omitempty"`
	// The name of the group. As a group's name can change, use of `groupId` is recommended to identify the group.
	Name string `json:"name,omitempty"`
}