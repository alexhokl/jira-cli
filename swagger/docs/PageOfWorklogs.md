# PageOfWorklogs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**StartAt** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of results on the page. | [optional] [readonly] 
**Worklogs** | Pointer to [**[]Worklog**](Worklog.md) | List of worklogs. | [optional] [readonly] 

## Methods

### NewPageOfWorklogs

`func NewPageOfWorklogs() *PageOfWorklogs`

NewPageOfWorklogs instantiates a new PageOfWorklogs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfWorklogsWithDefaults

`func NewPageOfWorklogsWithDefaults() *PageOfWorklogs`

NewPageOfWorklogsWithDefaults instantiates a new PageOfWorklogs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *PageOfWorklogs) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfWorklogs) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfWorklogs) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfWorklogs) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfWorklogs) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfWorklogs) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfWorklogs) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfWorklogs) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfWorklogs) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfWorklogs) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfWorklogs) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfWorklogs) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetWorklogs

`func (o *PageOfWorklogs) GetWorklogs() []Worklog`

GetWorklogs returns the Worklogs field if non-nil, zero value otherwise.

### GetWorklogsOk

`func (o *PageOfWorklogs) GetWorklogsOk() (*[]Worklog, bool)`

GetWorklogsOk returns a tuple with the Worklogs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorklogs

`func (o *PageOfWorklogs) SetWorklogs(v []Worklog)`

SetWorklogs sets Worklogs field to given value.

### HasWorklogs

`func (o *PageOfWorklogs) HasWorklogs() bool`

HasWorklogs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


