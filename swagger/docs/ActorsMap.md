# ActorsMap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **[]string** | The name of the group to add. This parameter cannot be used with the &#x60;groupId&#x60; parameter. As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended. | [optional] 
**GroupId** | Pointer to **[]string** | The ID of the group to add. This parameter cannot be used with the &#x60;group&#x60; parameter. | [optional] 
**User** | Pointer to **[]string** | The user account ID of the user to add. | [optional] 

## Methods

### NewActorsMap

`func NewActorsMap() *ActorsMap`

NewActorsMap instantiates a new ActorsMap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorsMapWithDefaults

`func NewActorsMapWithDefaults() *ActorsMap`

NewActorsMapWithDefaults instantiates a new ActorsMap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ActorsMap) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ActorsMap) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ActorsMap) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ActorsMap) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupId

`func (o *ActorsMap) GetGroupId() []string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ActorsMap) GetGroupIdOk() (*[]string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ActorsMap) SetGroupId(v []string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ActorsMap) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUser

`func (o *ActorsMap) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActorsMap) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActorsMap) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *ActorsMap) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


