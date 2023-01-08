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

type ActorsMap struct {
	// The name of the group to add. This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended.
	Group []string `json:"group,omitempty"`
	// The ID of the group to add. This parameter cannot be used with the `group` parameter.
	GroupId []string `json:"groupId,omitempty"`
	// The user account ID of the user to add.
	User []string `json:"user,omitempty"`
}