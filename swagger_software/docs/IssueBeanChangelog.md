# IssueBeanChangelog

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Histories** | Pointer to [**[]IssueBeanChangelogAllOfHistoriesInner**](IssueBeanChangelogAllOfHistoriesInner.md) | The list of changelogs. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**StartAt** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of results on the page. | [optional] [readonly] 

## Methods

### NewIssueBeanChangelog

`func NewIssueBeanChangelog() *IssueBeanChangelog`

NewIssueBeanChangelog instantiates a new IssueBeanChangelog object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBeanChangelogWithDefaults

`func NewIssueBeanChangelogWithDefaults() *IssueBeanChangelog`

NewIssueBeanChangelogWithDefaults instantiates a new IssueBeanChangelog object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHistories

`func (o *IssueBeanChangelog) GetHistories() []IssueBeanChangelogAllOfHistoriesInner`

GetHistories returns the Histories field if non-nil, zero value otherwise.

### GetHistoriesOk

`func (o *IssueBeanChangelog) GetHistoriesOk() (*[]IssueBeanChangelogAllOfHistoriesInner, bool)`

GetHistoriesOk returns a tuple with the Histories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistories

`func (o *IssueBeanChangelog) SetHistories(v []IssueBeanChangelogAllOfHistoriesInner)`

SetHistories sets Histories field to given value.

### HasHistories

`func (o *IssueBeanChangelog) HasHistories() bool`

HasHistories returns a boolean if a field has been set.

### GetMaxResults

`func (o *IssueBeanChangelog) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *IssueBeanChangelog) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *IssueBeanChangelog) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *IssueBeanChangelog) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *IssueBeanChangelog) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *IssueBeanChangelog) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *IssueBeanChangelog) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *IssueBeanChangelog) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *IssueBeanChangelog) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *IssueBeanChangelog) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *IssueBeanChangelog) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *IssueBeanChangelog) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


