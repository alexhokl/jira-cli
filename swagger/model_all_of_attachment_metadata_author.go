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

// Details of the user who attached the file.
type AllOfAttachmentMetadataAuthor struct {
	// The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required in requests.
	AccountId string `json:"accountId,omitempty"`
	// The user account type. Can take the following values:   *  `atlassian` regular Atlassian user account  *  `app` system account used for Connect applications and OAuth to represent external systems  *  `customer` Jira Service Desk account representing an external service desk
	AccountType string `json:"accountType,omitempty"`
	// Whether the user is active.
	Active bool `json:"active,omitempty"`
	// The application roles the user is assigned to.
	ApplicationRoles *Object `json:"applicationRoles,omitempty"`
	// The avatars of the user.
	AvatarUrls *Object `json:"avatarUrls,omitempty"`
	// The display name of the user. Depending on the user’s privacy setting, this may return an alternative value.
	DisplayName string `json:"displayName,omitempty"`
	// The email address of the user. Depending on the user’s privacy setting, this may be returned as null.
	EmailAddress string `json:"emailAddress,omitempty"`
	// Expand options that include additional user details in the response.
	Expand string `json:"expand,omitempty"`
	// The groups that the user belongs to.
	Groups *Object `json:"groups,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Key string `json:"key,omitempty"`
	// The locale of the user. Depending on the user’s privacy setting, this may be returned as null.
	Locale string `json:"locale,omitempty"`
	// This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.
	Name string `json:"name,omitempty"`
	// The URL of the user.
	Self string `json:"self,omitempty"`
	// The time zone specified in the user's profile. Depending on the user’s privacy setting, this may be returned as null.
	TimeZone string `json:"timeZone,omitempty"`
}
