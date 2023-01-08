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

// Details of an application role.
type ApplicationRole struct {
	// The groups that are granted default access for this application role. As a group's name can change, use of `defaultGroupsDetails` is recommended to identify a groups.
	DefaultGroups []string `json:"defaultGroups,omitempty"`
	// The groups that are granted default access for this application role.
	DefaultGroupsDetails []GroupName `json:"defaultGroupsDetails,omitempty"`
	// Deprecated.
	Defined bool `json:"defined,omitempty"`
	// The groups associated with the application role.
	GroupDetails []GroupName `json:"groupDetails,omitempty"`
	// The groups associated with the application role. As a group's name can change, use of `groupDetails` is recommended to identify a groups.
	Groups []string `json:"groups,omitempty"`
	HasUnlimitedSeats bool `json:"hasUnlimitedSeats,omitempty"`
	// The key of the application role.
	Key string `json:"key,omitempty"`
	// The display name of the application role.
	Name string `json:"name,omitempty"`
	// The maximum count of users on your license.
	NumberOfSeats int32 `json:"numberOfSeats,omitempty"`
	// Indicates if the application role belongs to Jira platform (`jira-core`).
	Platform bool `json:"platform,omitempty"`
	// The count of users remaining on your license.
	RemainingSeats int32 `json:"remainingSeats,omitempty"`
	// Determines whether this application role should be selected by default on user creation.
	SelectedByDefault bool `json:"selectedByDefault,omitempty"`
	// The number of users counting against your license.
	UserCount int32 `json:"userCount,omitempty"`
	// The [type of users](https://confluence.atlassian.com/x/lRW3Ng) being counted against your license.
	UserCountDescription string `json:"userCountDescription,omitempty"`
}