# WorkflowHistoryListResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entries** | Pointer to [**[]WorkflowHistoryItemDTO**](WorkflowHistoryItemDTO.md) |  | [optional] 

## Methods

### NewWorkflowHistoryListResponseDTO

`func NewWorkflowHistoryListResponseDTO() *WorkflowHistoryListResponseDTO`

NewWorkflowHistoryListResponseDTO instantiates a new WorkflowHistoryListResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowHistoryListResponseDTOWithDefaults

`func NewWorkflowHistoryListResponseDTOWithDefaults() *WorkflowHistoryListResponseDTO`

NewWorkflowHistoryListResponseDTOWithDefaults instantiates a new WorkflowHistoryListResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntries

`func (o *WorkflowHistoryListResponseDTO) GetEntries() []WorkflowHistoryItemDTO`

GetEntries returns the Entries field if non-nil, zero value otherwise.

### GetEntriesOk

`func (o *WorkflowHistoryListResponseDTO) GetEntriesOk() (*[]WorkflowHistoryItemDTO, bool)`

GetEntriesOk returns a tuple with the Entries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntries

`func (o *WorkflowHistoryListResponseDTO) SetEntries(v []WorkflowHistoryItemDTO)`

SetEntries sets Entries field to given value.

### HasEntries

`func (o *WorkflowHistoryListResponseDTO) HasEntries() bool`

HasEntries returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


