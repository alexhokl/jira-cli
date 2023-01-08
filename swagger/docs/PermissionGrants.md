# PermissionGrants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Expand options that include additional permission grant details in the response. | [optional] [readonly] 
**Permissions** | Pointer to [**[]PermissionGrant**](PermissionGrant.md) | Permission grants list. | [optional] [readonly] 

## Methods

### NewPermissionGrants

`func NewPermissionGrants() *PermissionGrants`

NewPermissionGrants instantiates a new PermissionGrants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPermissionGrantsWithDefaults

`func NewPermissionGrantsWithDefaults() *PermissionGrants`

NewPermissionGrantsWithDefaults instantiates a new PermissionGrants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *PermissionGrants) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *PermissionGrants) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *PermissionGrants) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *PermissionGrants) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetPermissions

`func (o *PermissionGrants) GetPermissions() []PermissionGrant`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *PermissionGrants) GetPermissionsOk() (*[]PermissionGrant, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *PermissionGrants) SetPermissions(v []PermissionGrant)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *PermissionGrants) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


