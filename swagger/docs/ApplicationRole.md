# ApplicationRole

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultGroups** | **[]string** | The groups that are granted default access for this application role. As a group&#x27;s name can change, use of &#x60;defaultGroupsDetails&#x60; is recommended to identify a groups. | [optional] [default to null]
**DefaultGroupsDetails** | [**[]GroupName**](GroupName.md) | The groups that are granted default access for this application role. | [optional] [default to null]
**Defined** | **bool** | Deprecated. | [optional] [default to null]
**GroupDetails** | [**[]GroupName**](GroupName.md) | The groups associated with the application role. | [optional] [default to null]
**Groups** | **[]string** | The groups associated with the application role. As a group&#x27;s name can change, use of &#x60;groupDetails&#x60; is recommended to identify a groups. | [optional] [default to null]
**HasUnlimitedSeats** | **bool** |  | [optional] [default to null]
**Key** | **string** | The key of the application role. | [optional] [default to null]
**Name** | **string** | The display name of the application role. | [optional] [default to null]
**NumberOfSeats** | **int32** | The maximum count of users on your license. | [optional] [default to null]
**Platform** | **bool** | Indicates if the application role belongs to Jira platform (&#x60;jira-core&#x60;). | [optional] [default to null]
**RemainingSeats** | **int32** | The count of users remaining on your license. | [optional] [default to null]
**SelectedByDefault** | **bool** | Determines whether this application role should be selected by default on user creation. | [optional] [default to null]
**UserCount** | **int32** | The number of users counting against your license. | [optional] [default to null]
**UserCountDescription** | **string** | The [type of users](https://confluence.atlassian.com/x/lRW3Ng) being counted against your license. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

