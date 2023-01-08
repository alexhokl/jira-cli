# PageBean2FieldAssociationSchemeFieldSearchResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Whether this is the last page. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of items that could be returned. | [optional] [readonly] 
**NextPage** | Pointer to **string** | If there is another page of results, the URL of the next page. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the page. | [optional] [readonly] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] [readonly] 
**Total** | Pointer to **int64** | The number of items returned. | [optional] [readonly] 
**Values** | Pointer to [**[]FieldAssociationSchemeFieldSearchResult**](FieldAssociationSchemeFieldSearchResult.md) | The list of items. | [optional] [readonly] 

## Methods

### NewPageBean2FieldAssociationSchemeFieldSearchResult

`func NewPageBean2FieldAssociationSchemeFieldSearchResult() *PageBean2FieldAssociationSchemeFieldSearchResult`

NewPageBean2FieldAssociationSchemeFieldSearchResult instantiates a new PageBean2FieldAssociationSchemeFieldSearchResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBean2FieldAssociationSchemeFieldSearchResultWithDefaults

`func NewPageBean2FieldAssociationSchemeFieldSearchResultWithDefaults() *PageBean2FieldAssociationSchemeFieldSearchResult`

NewPageBean2FieldAssociationSchemeFieldSearchResultWithDefaults instantiates a new PageBean2FieldAssociationSchemeFieldSearchResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPage

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartAt

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetValues() []FieldAssociationSchemeFieldSearchResult`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) GetValuesOk() (*[]FieldAssociationSchemeFieldSearchResult, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) SetValues(v []FieldAssociationSchemeFieldSearchResult)`

SetValues sets Values field to given value.

### HasValues

`func (o *PageBean2FieldAssociationSchemeFieldSearchResult) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


