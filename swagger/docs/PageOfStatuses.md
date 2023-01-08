# PageOfStatuses

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Whether this is the last page. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of items that could be returned. | [optional] 
**NextPage** | Pointer to **string** | The URL of the next page of results, if any. | [optional] 
**Self** | Pointer to **string** | The URL of this page. | [optional] 
**StartAt** | Pointer to **int64** | The index of the first item returned on the page. | [optional] 
**Total** | Pointer to **int64** | Number of items that satisfy the search. | [optional] 
**Values** | Pointer to [**[]JiraStatus**](JiraStatus.md) | The list of items. | [optional] 

## Methods

### NewPageOfStatuses

`func NewPageOfStatuses() *PageOfStatuses`

NewPageOfStatuses instantiates a new PageOfStatuses object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfStatusesWithDefaults

`func NewPageOfStatusesWithDefaults() *PageOfStatuses`

NewPageOfStatusesWithDefaults instantiates a new PageOfStatuses object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *PageOfStatuses) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *PageOfStatuses) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *PageOfStatuses) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *PageOfStatuses) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageOfStatuses) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfStatuses) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfStatuses) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfStatuses) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPage

`func (o *PageOfStatuses) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageOfStatuses) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageOfStatuses) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageOfStatuses) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *PageOfStatuses) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PageOfStatuses) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PageOfStatuses) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PageOfStatuses) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfStatuses) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfStatuses) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfStatuses) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfStatuses) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfStatuses) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfStatuses) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfStatuses) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfStatuses) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *PageOfStatuses) GetValues() []JiraStatus`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PageOfStatuses) GetValuesOk() (*[]JiraStatus, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PageOfStatuses) SetValues(v []JiraStatus)`

SetValues sets Values field to given value.

### HasValues

`func (o *PageOfStatuses) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


