# AllOfChangelogAuthor

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] [default to null]
**AccountType** | **string** | The type of account represented by this user. This will be one of &#x27;atlassian&#x27; (normal users), &#x27;app&#x27; (application user) or &#x27;customer&#x27; (Jira Service Desk customer user) | [optional] [default to null]
**Active** | **bool** | Whether the user is active. | [optional] [default to null]
**AvatarUrls** | [***Object**](.md) | The avatars of the user. | [optional] [default to null]
**DisplayName** | **string** | The display name of the user. Depending on the user’s privacy settings, this may return an alternative value. | [optional] [default to null]
**EmailAddress** | **string** | The email address of the user. Depending on the user’s privacy settings, this may be returned as null. | [optional] [default to null]
**Key** | **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [default to null]
**Name** | **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [default to null]
**Self** | **string** | The URL of the user. | [optional] [default to null]
**TimeZone** | **string** | The time zone specified in the user&#x27;s profile. Depending on the user’s privacy settings, this may be returned as null. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

