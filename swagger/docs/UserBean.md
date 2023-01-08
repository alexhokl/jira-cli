# UserBean

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] [default to null]
**Active** | **bool** | Whether the user is active. | [optional] [default to null]
**AvatarUrls** | [***AllOfUserBeanAvatarUrls**](AllOfUserBeanAvatarUrls.md) | The avatars of the user. | [optional] [default to null]
**DisplayName** | **string** | The display name of the user. Depending on the user’s privacy setting, this may return an alternative value. | [optional] [default to null]
**Key** | **string** | This property is deprecated in favor of &#x60;accountId&#x60; because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.   The key of the user. | [optional] [default to null]
**Name** | **string** | This property is deprecated in favor of &#x60;accountId&#x60; because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.   The username of the user. | [optional] [default to null]
**Self** | **string** | The URL of the user. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

