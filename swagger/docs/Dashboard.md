# Dashboard

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AutomaticRefreshMs** | **int32** | The automatic refresh interval for the dashboard in milliseconds. | [optional] [default to null]
**Description** | **string** |  | [optional] [default to null]
**EditPermissions** | [**[]SharePermission**](SharePermission.md) | The details of any edit share permissions for the dashboard. | [optional] [default to null]
**Id** | **string** | The ID of the dashboard. | [optional] [default to null]
**IsFavourite** | **bool** | Whether the dashboard is selected as a favorite by the user. | [optional] [default to null]
**IsWritable** | **bool** | Whether the current user has permission to edit the dashboard. | [optional] [default to null]
**Name** | **string** | The name of the dashboard. | [optional] [default to null]
**Owner** | [***AllOfDashboardOwner**](AllOfDashboardOwner.md) | The owner of the dashboard. | [optional] [default to null]
**Popularity** | **int64** | The number of users who have this dashboard as a favorite. | [optional] [default to null]
**Rank** | **int32** | The rank of this dashboard. | [optional] [default to null]
**Self** | **string** | The URL of these dashboard details. | [optional] [default to null]
**SharePermissions** | [**[]SharePermission**](SharePermission.md) | The details of any view share permissions for the dashboard. | [optional] [default to null]
**SystemDashboard** | **bool** | Whether the current dashboard is system dashboard. | [optional] [default to null]
**View** | **string** | The URL of the dashboard. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

