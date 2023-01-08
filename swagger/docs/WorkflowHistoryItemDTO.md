# WorkflowHistoryItemDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsIntermediate** | Pointer to **bool** | Whether the version is an intermediate workflow state, sometimes created during workflow updates. | [optional] 
**WorkflowId** | Pointer to **string** |  | [optional] 
**WorkflowVersion** | Pointer to **int64** |  | [optional] 
**WrittenAt** | Pointer to **string** | The timestamp when this workflow version was created. | [optional] 

## Methods

### NewWorkflowHistoryItemDTO

`func NewWorkflowHistoryItemDTO() *WorkflowHistoryItemDTO`

NewWorkflowHistoryItemDTO instantiates a new WorkflowHistoryItemDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowHistoryItemDTOWithDefaults

`func NewWorkflowHistoryItemDTOWithDefaults() *WorkflowHistoryItemDTO`

NewWorkflowHistoryItemDTOWithDefaults instantiates a new WorkflowHistoryItemDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsIntermediate

`func (o *WorkflowHistoryItemDTO) GetIsIntermediate() bool`

GetIsIntermediate returns the IsIntermediate field if non-nil, zero value otherwise.

### GetIsIntermediateOk

`func (o *WorkflowHistoryItemDTO) GetIsIntermediateOk() (*bool, bool)`

GetIsIntermediateOk returns a tuple with the IsIntermediate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsIntermediate

`func (o *WorkflowHistoryItemDTO) SetIsIntermediate(v bool)`

SetIsIntermediate sets IsIntermediate field to given value.

### HasIsIntermediate

`func (o *WorkflowHistoryItemDTO) HasIsIntermediate() bool`

HasIsIntermediate returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowHistoryItemDTO) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowHistoryItemDTO) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowHistoryItemDTO) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowHistoryItemDTO) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowVersion

`func (o *WorkflowHistoryItemDTO) GetWorkflowVersion() int64`

GetWorkflowVersion returns the WorkflowVersion field if non-nil, zero value otherwise.

### GetWorkflowVersionOk

`func (o *WorkflowHistoryItemDTO) GetWorkflowVersionOk() (*int64, bool)`

GetWorkflowVersionOk returns a tuple with the WorkflowVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowVersion

`func (o *WorkflowHistoryItemDTO) SetWorkflowVersion(v int64)`

SetWorkflowVersion sets WorkflowVersion field to given value.

### HasWorkflowVersion

`func (o *WorkflowHistoryItemDTO) HasWorkflowVersion() bool`

HasWorkflowVersion returns a boolean if a field has been set.

### GetWrittenAt

`func (o *WorkflowHistoryItemDTO) GetWrittenAt() string`

GetWrittenAt returns the WrittenAt field if non-nil, zero value otherwise.

### GetWrittenAtOk

`func (o *WorkflowHistoryItemDTO) GetWrittenAtOk() (*string, bool)`

GetWrittenAtOk returns a tuple with the WrittenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrittenAt

`func (o *WorkflowHistoryItemDTO) SetWrittenAt(v string)`

SetWrittenAt sets WrittenAt field to given value.

### HasWrittenAt

`func (o *WorkflowHistoryItemDTO) HasWrittenAt() bool`

HasWrittenAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


