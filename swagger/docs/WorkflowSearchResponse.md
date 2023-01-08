# WorkflowSearchResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Whether this is the last page. | [optional] 
**MaxResults** | Pointer to **int32** | The maximum number of items that could be returned. | [optional] 
**NextPage** | Pointer to **string** | If there is another page of results, the URL of the next page. | [optional] 
**Self** | Pointer to **string** | The URL of the page. | [optional] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] 
**Statuses** | Pointer to [**[]JiraWorkflowStatus**](JiraWorkflowStatus.md) | List of statuses. | [optional] 
**Total** | Pointer to **int64** | The number of items returned. | [optional] 
**Values** | Pointer to [**[]JiraWorkflow**](JiraWorkflow.md) | List of workflows. | [optional] 

## Methods

### NewWorkflowSearchResponse

`func NewWorkflowSearchResponse() *WorkflowSearchResponse`

NewWorkflowSearchResponse instantiates a new WorkflowSearchResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSearchResponseWithDefaults

`func NewWorkflowSearchResponseWithDefaults() *WorkflowSearchResponse`

NewWorkflowSearchResponseWithDefaults instantiates a new WorkflowSearchResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *WorkflowSearchResponse) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *WorkflowSearchResponse) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *WorkflowSearchResponse) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *WorkflowSearchResponse) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *WorkflowSearchResponse) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *WorkflowSearchResponse) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *WorkflowSearchResponse) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *WorkflowSearchResponse) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPage

`func (o *WorkflowSearchResponse) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *WorkflowSearchResponse) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *WorkflowSearchResponse) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *WorkflowSearchResponse) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *WorkflowSearchResponse) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *WorkflowSearchResponse) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *WorkflowSearchResponse) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *WorkflowSearchResponse) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartAt

`func (o *WorkflowSearchResponse) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *WorkflowSearchResponse) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *WorkflowSearchResponse) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *WorkflowSearchResponse) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetStatuses

`func (o *WorkflowSearchResponse) GetStatuses() []JiraWorkflowStatus`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowSearchResponse) GetStatusesOk() (*[]JiraWorkflowStatus, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowSearchResponse) SetStatuses(v []JiraWorkflowStatus)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowSearchResponse) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetTotal

`func (o *WorkflowSearchResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *WorkflowSearchResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *WorkflowSearchResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *WorkflowSearchResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *WorkflowSearchResponse) GetValues() []JiraWorkflow`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *WorkflowSearchResponse) GetValuesOk() (*[]JiraWorkflow, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *WorkflowSearchResponse) SetValues(v []JiraWorkflow)`

SetValues sets Values field to given value.

### HasValues

`func (o *WorkflowSearchResponse) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


