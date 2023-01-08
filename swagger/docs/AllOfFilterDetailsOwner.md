# AllOfFilterDetailsOwner

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required in requests. | [optional] [default to null]
**AccountType** | **string** | The user account type. Can take the following values:   *  &#x60;atlassian&#x60; regular Atlassian user account  *  &#x60;app&#x60; system account used for Connect applications and OAuth to represent external systems  *  &#x60;customer&#x60; Jira Service Desk account representing an external service desk | [optional] [default to null]
**Active** | **bool** | Whether the user is active. | [optional] [default to null]
**ApplicationRoles** | [***Object**](.md) | The application roles the user is assigned to. | [optional] [default to null]
**AvatarUrls** | [***Object**](.md) | The avatars of the user. | [optional] [default to null]
**DisplayName** | **string** | The display name of the user. Depending on the user’s privacy setting, this may return an alternative value. | [optional] [default to null]
**EmailAddress** | **string** | The email address of the user. Depending on the user’s privacy setting, this may be returned as null. | [optional] [default to null]
**Expand** | **string** | Expand options that include additional user details in the response. | [optional] [default to null]
**Groups** | [***Object**](.md) | The groups that the user belongs to. | [optional] [default to null]
**Key** | **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [default to null]
**Locale** | **string** | The locale of the user. Depending on the user’s privacy setting, this may be returned as null. | [optional] [default to null]
**Name** | **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [default to null]
**Self** | **string** | The URL of the user. | [optional] [default to null]
**TimeZone** | **string** | The time zone specified in the user&#x27;s profile. Depending on the user’s privacy setting, this may be returned as null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

