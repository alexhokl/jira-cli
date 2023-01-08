# PageOfDashboards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dashboards** | Pointer to [**[]Dashboard**](Dashboard.md) | List of dashboards. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of results that could be on the page. | [optional] [readonly] 
**Next** | Pointer to **string** | The URL of the next page of results, if any. | [optional] [readonly] 
**Prev** | Pointer to **string** | The URL of the previous page of results, if any. | [optional] [readonly] 
**StartAt** | Pointer to **int32** | The index of the first item returned on the page. | [optional] [readonly] 
**Total** | Pointer to **int32** | The number of results on the page. | [optional] [readonly] 

## Methods

### NewPageOfDashboards

`func NewPageOfDashboards() *PageOfDashboards`

NewPageOfDashboards instantiates a new PageOfDashboards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfDashboardsWithDefaults

`func NewPageOfDashboardsWithDefaults() *PageOfDashboards`

NewPageOfDashboardsWithDefaults instantiates a new PageOfDashboards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDashboards

`func (o *PageOfDashboards) GetDashboards() []Dashboard`

GetDashboards returns the Dashboards field if non-nil, zero value otherwise.

### GetDashboardsOk

`func (o *PageOfDashboards) GetDashboardsOk() (*[]Dashboard, bool)`

GetDashboardsOk returns a tuple with the Dashboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDashboards

`func (o *PageOfDashboards) SetDashboards(v []Dashboard)`

SetDashboards sets Dashboards field to given value.

### HasDashboards

`func (o *PageOfDashboards) HasDashboards() bool`

HasDashboards returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageOfDashboards) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfDashboards) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfDashboards) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfDashboards) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNext

`func (o *PageOfDashboards) GetNext() string`

GetNext returns the Next field if non-nil, zero value otherwise.

### GetNextOk

`func (o *PageOfDashboards) GetNextOk() (*string, bool)`

GetNextOk returns a tuple with the Next field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNext

`func (o *PageOfDashboards) SetNext(v string)`

SetNext sets Next field to given value.

### HasNext

`func (o *PageOfDashboards) HasNext() bool`

HasNext returns a boolean if a field has been set.

### GetPrev

`func (o *PageOfDashboards) GetPrev() string`

GetPrev returns the Prev field if non-nil, zero value otherwise.

### GetPrevOk

`func (o *PageOfDashboards) GetPrevOk() (*string, bool)`

GetPrevOk returns a tuple with the Prev field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrev

`func (o *PageOfDashboards) SetPrev(v string)`

SetPrev sets Prev field to given value.

### HasPrev

`func (o *PageOfDashboards) HasPrev() bool`

HasPrev returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfDashboards) GetStartAt() int32`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfDashboards) GetStartAtOk() (*int32, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfDashboards) SetStartAt(v int32)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfDashboards) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfDashboards) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfDashboards) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfDashboards) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfDashboards) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


