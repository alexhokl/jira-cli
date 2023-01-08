# FoundUsersAndGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**FoundGroups**](FoundGroups.md) |  | [optional] 
**Users** | Pointer to [**FoundUsers**](FoundUsers.md) |  | [optional] 

## Methods

### NewFoundUsersAndGroups

`func NewFoundUsersAndGroups() *FoundUsersAndGroups`

NewFoundUsersAndGroups instantiates a new FoundUsersAndGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundUsersAndGroupsWithDefaults

`func NewFoundUsersAndGroupsWithDefaults() *FoundUsersAndGroups`

NewFoundUsersAndGroupsWithDefaults instantiates a new FoundUsersAndGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *FoundUsersAndGroups) GetGroups() FoundGroups`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FoundUsersAndGroups) GetGroupsOk() (*FoundGroups, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FoundUsersAndGroups) SetGroups(v FoundGroups)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FoundUsersAndGroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *FoundUsersAndGroups) GetUsers() FoundUsers`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FoundUsersAndGroups) GetUsersOk() (*FoundUsers, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FoundUsersAndGroups) SetUsers(v FoundUsers)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *FoundUsersAndGroups) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


