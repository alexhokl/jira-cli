# PermissionDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditPermissions** | [**[]SharePermission**](SharePermission.md) | The edit permissions for the shareable entities. | 
**SharePermissions** | [**[]SharePermission**](SharePermission.md) | The share permissions for the shareable entities. | 

## Methods

### NewPermissionDetails

`func NewPermissionDetails(editPermissions []SharePermission, sharePermissions []SharePermission, ) *PermissionDetails`

NewPermissionDetails instantiates a new PermissionDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionDetailsWithDefaults

`func NewPermissionDetailsWithDefaults() *PermissionDetails`

NewPermissionDetailsWithDefaults instantiates a new PermissionDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditPermissions

`func (o *PermissionDetails) GetEditPermissions() []SharePermission`

GetEditPermissions returns the EditPermissions field if non-nil, zero value otherwise.

### GetEditPermissionsOk

`func (o *PermissionDetails) GetEditPermissionsOk() (*[]SharePermission, bool)`

GetEditPermissionsOk returns a tuple with the EditPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditPermissions

`func (o *PermissionDetails) SetEditPermissions(v []SharePermission)`

SetEditPermissions sets EditPermissions field to given value.


### GetSharePermissions

`func (o *PermissionDetails) GetSharePermissions() []SharePermission`

GetSharePermissions returns the SharePermissions field if non-nil, zero value otherwise.

### GetSharePermissionsOk

`func (o *PermissionDetails) GetSharePermissionsOk() (*[]SharePermission, bool)`

GetSharePermissionsOk returns a tuple with the SharePermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharePermissions

`func (o *PermissionDetails) SetSharePermissions(v []SharePermission)`

SetSharePermissions sets SharePermissions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


