# BulkProjectPermissions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Issues** | Pointer to **[]int64** | List of issue IDs. | [optional] 
**Permissions** | **[]string** | List of project permissions. | 
**Projects** | Pointer to **[]int64** | List of project IDs. | [optional] 

## Methods

### NewBulkProjectPermissions

`func NewBulkProjectPermissions(permissions []string, ) *BulkProjectPermissions`

NewBulkProjectPermissions instantiates a new BulkProjectPermissions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkProjectPermissionsWithDefaults

`func NewBulkProjectPermissionsWithDefaults() *BulkProjectPermissions`

NewBulkProjectPermissionsWithDefaults instantiates a new BulkProjectPermissions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssues

`func (o *BulkProjectPermissions) GetIssues() []int64`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *BulkProjectPermissions) GetIssuesOk() (*[]int64, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *BulkProjectPermissions) SetIssues(v []int64)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *BulkProjectPermissions) HasIssues() bool`

HasIssues returns a boolean if a field has been set.

### GetPermissions

`func (o *BulkProjectPermissions) GetPermissions() []string`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *BulkProjectPermissions) GetPermissionsOk() (*[]string, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *BulkProjectPermissions) SetPermissions(v []string)`

SetPermissions sets Permissions field to given value.


### GetProjects

`func (o *BulkProjectPermissions) GetProjects() []int64`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *BulkProjectPermissions) GetProjectsOk() (*[]int64, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *BulkProjectPermissions) SetProjects(v []int64)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *BulkProjectPermissions) HasProjects() bool`

HasProjects returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


