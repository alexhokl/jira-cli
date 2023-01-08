# BulkPermissionGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GlobalPermissions** | **[]string** | List of permissions granted to the user. | 
**ProjectPermissions** | [**[]BulkProjectPermissionGrants**](BulkProjectPermissionGrants.md) | List of project permissions and the projects and issues those permissions provide access to. | 

## Methods

### NewBulkPermissionGrants

`func NewBulkPermissionGrants(globalPermissions []string, projectPermissions []BulkProjectPermissionGrants, ) *BulkPermissionGrants`

NewBulkPermissionGrants instantiates a new BulkPermissionGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkPermissionGrantsWithDefaults

`func NewBulkPermissionGrantsWithDefaults() *BulkPermissionGrants`

NewBulkPermissionGrantsWithDefaults instantiates a new BulkPermissionGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGlobalPermissions

`func (o *BulkPermissionGrants) GetGlobalPermissions() []string`

GetGlobalPermissions returns the GlobalPermissions field if non-nil, zero value otherwise.

### GetGlobalPermissionsOk

`func (o *BulkPermissionGrants) GetGlobalPermissionsOk() (*[]string, bool)`

GetGlobalPermissionsOk returns a tuple with the GlobalPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalPermissions

`func (o *BulkPermissionGrants) SetGlobalPermissions(v []string)`

SetGlobalPermissions sets GlobalPermissions field to given value.


### GetProjectPermissions

`func (o *BulkPermissionGrants) GetProjectPermissions() []BulkProjectPermissionGrants`

GetProjectPermissions returns the ProjectPermissions field if non-nil, zero value otherwise.

### GetProjectPermissionsOk

`func (o *BulkPermissionGrants) GetProjectPermissionsOk() (*[]BulkProjectPermissionGrants, bool)`

GetProjectPermissionsOk returns a tuple with the ProjectPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectPermissions

`func (o *BulkPermissionGrants) SetProjectPermissions(v []BulkProjectPermissionGrants)`

SetProjectPermissions sets ProjectPermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


