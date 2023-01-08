# BulkProjectPermissionGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | **[]int64** | IDs of the issues the user has the permission for. | 
**Permission** | **string** | A project permission, | 
**Projects** | **[]int64** | IDs of the projects the user has the permission for. | 

## Methods

### NewBulkProjectPermissionGrants

`func NewBulkProjectPermissionGrants(issues []int64, permission string, projects []int64, ) *BulkProjectPermissionGrants`

NewBulkProjectPermissionGrants instantiates a new BulkProjectPermissionGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkProjectPermissionGrantsWithDefaults

`func NewBulkProjectPermissionGrantsWithDefaults() *BulkProjectPermissionGrants`

NewBulkProjectPermissionGrantsWithDefaults instantiates a new BulkProjectPermissionGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *BulkProjectPermissionGrants) GetIssues() []int64`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *BulkProjectPermissionGrants) GetIssuesOk() (*[]int64, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *BulkProjectPermissionGrants) SetIssues(v []int64)`

SetIssues sets Issues field to given value.


### GetPermission

`func (o *BulkProjectPermissionGrants) GetPermission() string`

GetPermission returns the Permission field if non-nil, zero value otherwise.

### GetPermissionOk

`func (o *BulkProjectPermissionGrants) GetPermissionOk() (*string, bool)`

GetPermissionOk returns a tuple with the Permission field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermission

`func (o *BulkProjectPermissionGrants) SetPermission(v string)`

SetPermission sets Permission field to given value.


### GetProjects

`func (o *BulkProjectPermissionGrants) GetProjects() []int64`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BulkProjectPermissionGrants) GetProjectsOk() (*[]int64, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BulkProjectPermissionGrants) SetProjects(v []int64)`

SetProjects sets Projects field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


