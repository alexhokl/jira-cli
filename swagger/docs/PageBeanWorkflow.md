# PageBeanWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** | Whether this is the last page. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of items that could be returned. | [optional] [readonly] 
**NextPage** | Pointer to **string** | If there is another page of results, the URL of the next page. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the page. | [optional] [readonly] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] [readonly] 
**Total** | Pointer to **int64** | The number of items returned. | [optional] [readonly] 
**Values** | Pointer to [**[]Workflow**](Workflow.md) | The list of items. | [optional] [readonly] 

## Methods

### NewPageBeanWorkflow

`func NewPageBeanWorkflow() *PageBeanWorkflow`

NewPageBeanWorkflow instantiates a new PageBeanWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBeanWorkflowWithDefaults

`func NewPageBeanWorkflowWithDefaults() *PageBeanWorkflow`

NewPageBeanWorkflowWithDefaults instantiates a new PageBeanWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *PageBeanWorkflow) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *PageBeanWorkflow) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *PageBeanWorkflow) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *PageBeanWorkflow) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageBeanWorkflow) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageBeanWorkflow) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageBeanWorkflow) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageBeanWorkflow) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetNextPage

`func (o *PageBeanWorkflow) GetNextPage() string`

GetNextPage returns the NextPage field if non-nil, zero value otherwise.

### GetNextPageOk

`func (o *PageBeanWorkflow) GetNextPageOk() (*string, bool)`

GetNextPageOk returns a tuple with the NextPage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextPage

`func (o *PageBeanWorkflow) SetNextPage(v string)`

SetNextPage sets NextPage field to given value.

### HasNextPage

`func (o *PageBeanWorkflow) HasNextPage() bool`

HasNextPage returns a boolean if a field has been set.

### GetSelf

`func (o *PageBeanWorkflow) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PageBeanWorkflow) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PageBeanWorkflow) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PageBeanWorkflow) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartAt

`func (o *PageBeanWorkflow) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageBeanWorkflow) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageBeanWorkflow) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageBeanWorkflow) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageBeanWorkflow) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageBeanWorkflow) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageBeanWorkflow) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageBeanWorkflow) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *PageBeanWorkflow) GetValues() []Workflow`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PageBeanWorkflow) GetValuesOk() (*[]Workflow, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PageBeanWorkflow) SetValues(v []Workflow)`

SetValues sets Values field to given value.

### HasValues

`func (o *PageBeanWorkflow) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


