# PermissionGrant

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | [***AllOfPermissionGrantHolder**](AllOfPermissionGrantHolder.md) | The user or group being granted the permission. It consists of a &#x60;type&#x60;, a type-dependent &#x60;parameter&#x60; and a type-dependent &#x60;value&#x60;. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information. | [optional] [default to null]
**Id** | **int64** | The ID of the permission granted details. | [optional] [default to null]
**Permission** | **string** | The permission to grant. This permission can be one of the built-in permissions or a custom permission added by an app. See [Built-in permissions](../api-group-permission-schemes/#built-in-permissions) in *Get all permission schemes* for more information about the built-in permissions. See the [project permission](https://developer.atlassian.com/cloud/jira/platform/modules/project-permission/) and [global permission](https://developer.atlassian.com/cloud/jira/platform/modules/global-permission/) module documentation for more information about custom permissions. | [optional] [default to null]
**Self** | **string** | The URL of the permission granted details. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

