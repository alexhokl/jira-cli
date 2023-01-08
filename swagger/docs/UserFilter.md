# UserFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | Whether the filter is enabled. | 
**Groups** | Pointer to **[]string** | User groups autocomplete suggestion users must belong to. If not provided, the default values are used. A maximum of 10 groups can be provided. | [optional] 
**RoleIds** | Pointer to **[]int64** | Roles that autocomplete suggestion users must belong to. If not provided, the default values are used. A maximum of 10 roles can be provided. | [optional] 

## Methods

### NewUserFilter

`func NewUserFilter(enabled bool, ) *UserFilter`

NewUserFilter instantiates a new UserFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserFilterWithDefaults

`func NewUserFilterWithDefaults() *UserFilter`

NewUserFilterWithDefaults instantiates a new UserFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserFilter) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserFilter) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserFilter) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetGroups

`func (o *UserFilter) GetGroups() []string`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserFilter) GetGroupsOk() (*[]string, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserFilter) SetGroups(v []string)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserFilter) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetRoleIds

`func (o *UserFilter) GetRoleIds() []int64`

GetRoleIds returns the RoleIds field if non-nil, zero value otherwise.

### GetRoleIdsOk

`func (o *UserFilter) GetRoleIdsOk() (*[]int64, bool)`

GetRoleIdsOk returns a tuple with the RoleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleIds

`func (o *UserFilter) SetRoleIds(v []int64)`

SetRoleIds sets RoleIds field to given value.

### HasRoleIds

`func (o *UserFilter) HasRoleIds() bool`

HasRoleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


