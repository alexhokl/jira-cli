# PageOfChangelogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histories** | Pointer to [**[]IssueBeanChangelogAllOfHistoriesInner**](IssueBeanChangelogAllOfHistoriesInner.md) | The list of changelogs. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**StartAt** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of results on the page. | [optional] [readonly] 

## Methods

### NewPageOfChangelogs

`func NewPageOfChangelogs() *PageOfChangelogs`

NewPageOfChangelogs instantiates a new PageOfChangelogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfChangelogsWithDefaults

`func NewPageOfChangelogsWithDefaults() *PageOfChangelogs`

NewPageOfChangelogsWithDefaults instantiates a new PageOfChangelogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistories

`func (o *PageOfChangelogs) GetHistories() []IssueBeanChangelogAllOfHistoriesInner`

GetHistories returns the Histories field if non-nil, zero value otherwise.

### GetHistoriesOk

`func (o *PageOfChangelogs) GetHistoriesOk() (*[]IssueBeanChangelogAllOfHistoriesInner, bool)`

GetHistoriesOk returns a tuple with the Histories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistories

`func (o *PageOfChangelogs) SetHistories(v []IssueBeanChangelogAllOfHistoriesInner)`

SetHistories sets Histories field to given value.

### HasHistories

`func (o *PageOfChangelogs) HasHistories() bool`

HasHistories returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageOfChangelogs) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfChangelogs) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfChangelogs) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfChangelogs) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfChangelogs) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfChangelogs) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfChangelogs) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfChangelogs) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfChangelogs) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfChangelogs) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfChangelogs) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfChangelogs) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


