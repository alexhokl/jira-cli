# FilterSubscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**GroupName**](GroupName.md) | The group subscribing to filter. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the filter subscription. | [optional] [readonly] 
**User** | Pointer to [**User**](User.md) | The user subscribing to filter. | [optional] [readonly] 

## Methods

### NewFilterSubscription

`func NewFilterSubscription() *FilterSubscription`

NewFilterSubscription instantiates a new FilterSubscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFilterSubscriptionWithDefaults

`func NewFilterSubscriptionWithDefaults() *FilterSubscription`

NewFilterSubscriptionWithDefaults instantiates a new FilterSubscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *FilterSubscription) GetGroup() GroupName`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *FilterSubscription) GetGroupOk() (*GroupName, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *FilterSubscription) SetGroup(v GroupName)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *FilterSubscription) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *FilterSubscription) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FilterSubscription) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FilterSubscription) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *FilterSubscription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUser

`func (o *FilterSubscription) GetUser() User`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *FilterSubscription) GetUserOk() (*User, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *FilterSubscription) SetUser(v User)`

SetUser sets User field to given value.

### HasUser

`func (o *FilterSubscription) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


