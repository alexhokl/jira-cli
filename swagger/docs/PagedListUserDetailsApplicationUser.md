# PagedListUserDetailsApplicationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndIndex** | Pointer to **int32** | The index of the last item returned on the page. | [optional] [readonly] 
**Items** | Pointer to [**[]UserDetails**](UserDetails.md) | The list of items. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**Size** | Pointer to **int32** | The number of items on the page. | [optional] [readonly] 
**StartIndex** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 

## Methods

### NewPagedListUserDetailsApplicationUser

`func NewPagedListUserDetailsApplicationUser() *PagedListUserDetailsApplicationUser`

NewPagedListUserDetailsApplicationUser instantiates a new PagedListUserDetailsApplicationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPagedListUserDetailsApplicationUserWithDefaults

`func NewPagedListUserDetailsApplicationUserWithDefaults() *PagedListUserDetailsApplicationUser`

NewPagedListUserDetailsApplicationUserWithDefaults instantiates a new PagedListUserDetailsApplicationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndIndex

`func (o *PagedListUserDetailsApplicationUser) GetEndIndex() int32`

GetEndIndex returns the EndIndex field if non-nil, zero value otherwise.

### GetEndIndexOk

`func (o *PagedListUserDetailsApplicationUser) GetEndIndexOk() (*int32, bool)`

GetEndIndexOk returns a tuple with the EndIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndIndex

`func (o *PagedListUserDetailsApplicationUser) SetEndIndex(v int32)`

SetEndIndex sets EndIndex field to given value.

### HasEndIndex

`func (o *PagedListUserDetailsApplicationUser) HasEndIndex() bool`

HasEndIndex returns a boolean if a field has been set.

### GetItems

`func (o *PagedListUserDetailsApplicationUser) GetItems() []UserDetails`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *PagedListUserDetailsApplicationUser) GetItemsOk() (*[]UserDetails, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *PagedListUserDetailsApplicationUser) SetItems(v []UserDetails)`

SetItems sets Items field to given value.

### HasItems

`func (o *PagedListUserDetailsApplicationUser) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMaxResults

`func (o *PagedListUserDetailsApplicationUser) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PagedListUserDetailsApplicationUser) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PagedListUserDetailsApplicationUser) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PagedListUserDetailsApplicationUser) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetSize

`func (o *PagedListUserDetailsApplicationUser) GetSize() int32`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PagedListUserDetailsApplicationUser) GetSizeOk() (*int32, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PagedListUserDetailsApplicationUser) SetSize(v int32)`

SetSize sets Size field to given value.

### HasSize

`func (o *PagedListUserDetailsApplicationUser) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetStartIndex

`func (o *PagedListUserDetailsApplicationUser) GetStartIndex() int32`

GetStartIndex returns the StartIndex field if non-nil, zero value otherwise.

### GetStartIndexOk

`func (o *PagedListUserDetailsApplicationUser) GetStartIndexOk() (*int32, bool)`

GetStartIndexOk returns a tuple with the StartIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartIndex

`func (o *PagedListUserDetailsApplicationUser) SetStartIndex(v int32)`

SetStartIndex sets StartIndex field to given value.

### HasStartIndex

`func (o *PagedListUserDetailsApplicationUser) HasStartIndex() bool`

HasStartIndex returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


