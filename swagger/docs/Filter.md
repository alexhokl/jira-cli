# Filter

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | A description of the filter. | [optional] [default to null]
**EditPermissions** | [**[]SharePermission**](SharePermission.md) | The groups and projects that can edit the filter. | [optional] [default to null]
**Favourite** | **bool** | Whether the filter is selected as a favorite. | [optional] [default to null]
**FavouritedCount** | **int64** | The count of how many users have selected this filter as a favorite, including the filter owner. | [optional] [default to null]
**Id** | **string** | The unique identifier for the filter. | [optional] [default to null]
**Jql** | **string** | The JQL query for the filter. For example, *project &#x3D; SSP AND issuetype &#x3D; Bug*. | [optional] [default to null]
**Name** | **string** | The name of the filter. Must be unique. | [default to null]
**Owner** | [***AllOfFilterOwner**](AllOfFilterOwner.md) | The user who owns the filter. This is defaulted to the creator of the filter, however Jira administrators can change the owner of a shared filter in the admin settings. | [optional] [default to null]
**SearchUrl** | **string** | A URL to view the filter results in Jira, using the [Search for issues using JQL](#api-rest-api-3-filter-search-get) operation with the filter&#x27;s JQL string to return the filter results. For example, *https://your-domain.atlassian.net/rest/api/3/search?jql&#x3D;project+%3D+SSP+AND+issuetype+%3D+Bug*. | [optional] [default to null]
**Self** | **string** | The URL of the filter. | [optional] [default to null]
**SharePermissions** | [**[]SharePermission**](SharePermission.md) | The groups and projects that the filter is shared with. | [optional] [default to null]
**SharedUsers** | [***AllOfFilterSharedUsers**](AllOfFilterSharedUsers.md) | A paginated list of the users that the filter is shared with. This includes users that are members of the groups or can browse the projects that the filter is shared with. | [optional] [default to null]
**Subscriptions** | [***AllOfFilterSubscriptions**](AllOfFilterSubscriptions.md) | A paginated list of the users that are subscribed to the filter. | [optional] [default to null]
**ViewUrl** | **string** | A URL to view the filter results in Jira, using the ID of the filter. For example, *https://your-domain.atlassian.net/issues/?filter&#x3D;10100*. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

