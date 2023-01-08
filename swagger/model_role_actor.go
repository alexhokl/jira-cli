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

// Details about a user assigned to a project role.
type RoleActor struct {
	ActorGroup *AllOfRoleActorActorGroup `json:"actorGroup,omitempty"`
	ActorUser *AllOfRoleActorActorUser `json:"actorUser,omitempty"`
	// The avatar of the role actor.
	AvatarUrl string `json:"avatarUrl,omitempty"`
	// The display name of the role actor. For users, depending on the user’s privacy setting, this may return an alternative value for the user's name.
	DisplayName string `json:"displayName,omitempty"`
	// The ID of the role actor.
	Id int64 `json:"id,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Name string `json:"name,omitempty"`
	// The type of role actor.
	Type_ string `json:"type,omitempty"`
}