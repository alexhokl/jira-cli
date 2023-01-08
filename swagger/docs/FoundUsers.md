# FoundUsers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Header** | Pointer to **string** | Header text indicating the number of users in the response and the total number of users found in the search. | [optional] 
**Total** | Pointer to **int32** | The total number of users found in the search. | [optional] 
**Users** | Pointer to [**[]UserPickerUser**](UserPickerUser.md) |  | [optional] 

## Methods

### NewFoundUsers

`func NewFoundUsers() *FoundUsers`

NewFoundUsers instantiates a new FoundUsers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundUsersWithDefaults

`func NewFoundUsersWithDefaults() *FoundUsers`

NewFoundUsersWithDefaults instantiates a new FoundUsers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeader

`func (o *FoundUsers) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FoundUsers) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FoundUsers) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FoundUsers) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTotal

`func (o *FoundUsers) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FoundUsers) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FoundUsers) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FoundUsers) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetUsers

`func (o *FoundUsers) GetUsers() []UserPickerUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *FoundUsers) GetUsersOk() (*[]UserPickerUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *FoundUsers) SetUsers(v []UserPickerUser)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *FoundUsers) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


