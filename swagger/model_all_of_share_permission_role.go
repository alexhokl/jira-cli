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

// The project role that the filter is shared with.   For a request, specify the `id` for the role. You must also specify the `project` object and `id` for the project that the role is in.
type AllOfSharePermissionRole struct {
	// The list of users who act in this role.
	Actors []RoleActor `json:"actors,omitempty"`
	// Whether this role is the admin role for the project.
	Admin bool `json:"admin,omitempty"`
	// Whether the calling user is part of this role.
	CurrentUserRole bool `json:"currentUserRole,omitempty"`
	// Whether this role is the default role for the project
	Default_ bool `json:"default,omitempty"`
	// The description of the project role.
	Description string `json:"description,omitempty"`
	// The ID of the project role.
	Id int64 `json:"id,omitempty"`
	// The name of the project role.
	Name string `json:"name,omitempty"`
	// Whether the roles are configurable for this project.
	RoleConfigurable bool `json:"roleConfigurable,omitempty"`
	// The scope of the role. Indicated for roles associated with [next-gen projects](https://confluence.atlassian.com/x/loMyO).
	Scope *Object `json:"scope,omitempty"`
	// The URL the project role details.
	Self string `json:"self,omitempty"`
	// The translated name of the project role.
	TranslatedName string `json:"translatedName,omitempty"`
}
