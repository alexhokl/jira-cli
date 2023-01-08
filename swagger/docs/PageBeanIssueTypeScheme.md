# PageBeanIssueTypeScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Whether this is the last page. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of items that could be returned. | [optional] [readonly] 
**NextPage** | Pointer to **string** | If there is another page of results, the URL of the next page. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the page. | [optional] [readonly] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] [readonly] 
**Total** | Pointer to **int64** | The number of items returned. | [optional] [readonly] 
**Values** | Pointer to [**[]IssueTypeScheme**](IssueTypeScheme.md) | The list of items. | [optional] [readonly] 

## Methods

### NewPageBeanIssueTypeScheme

`func NewPageBeanIssueTypeScheme() *PageBeanIssueTypeScheme`

NewPageBeanIssueTypeScheme instantiates a new PageBeanIssueTypeScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBeanIssueTypeSchemeWithDefaults

`func NewPageBeanIssueTypeSchemeWithDefaults() *PageBeanIssueTypeScheme`

NewPageBeanIssueTypeSchemeWithDefaults instantiates a new PageBeanIssueTypeScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *PageBeanIssueTypeScheme) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *PageBeanIssueTypeScheme) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *PageBeanIssueTypeScheme) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *PageBeanIssueTypeScheme) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageBeanIssueTypeScheme) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageBeanIssueTypeScheme) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageBeanIssueTypeScheme) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageBeanIssueTypeScheme) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPage

`func (o *PageBeanIssueTypeScheme) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageBeanIssueTypeScheme) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageBeanIssueTypeScheme) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageBeanIssueTypeScheme) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *PageBeanIssueTypeScheme) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PageBeanIssueTypeScheme) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PageBeanIssueTypeScheme) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PageBeanIssueTypeScheme) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartAt

`func (o *PageBeanIssueTypeScheme) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageBeanIssueTypeScheme) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageBeanIssueTypeScheme) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageBeanIssueTypeScheme) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageBeanIssueTypeScheme) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageBeanIssueTypeScheme) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageBeanIssueTypeScheme) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageBeanIssueTypeScheme) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *PageBeanIssueTypeScheme) GetValues() []IssueTypeScheme`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PageBeanIssueTypeScheme) GetValuesOk() (*[]IssueTypeScheme, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PageBeanIssueTypeScheme) SetValues(v []IssueTypeScheme)`

SetValues sets Values field to given value.

### HasValues

`func (o *PageBeanIssueTypeScheme) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


