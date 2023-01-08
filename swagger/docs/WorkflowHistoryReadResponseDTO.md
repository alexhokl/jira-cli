# WorkflowHistoryReadResponseDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | Pointer to [**[]WorkflowDocumentStatusDTO**](WorkflowDocumentStatusDTO.md) |  | [optional] 
**Workflows** | Pointer to [**[]WorkflowDocumentDTO**](WorkflowDocumentDTO.md) |  | [optional] 

## Methods

### NewWorkflowHistoryReadResponseDTO

`func NewWorkflowHistoryReadResponseDTO() *WorkflowHistoryReadResponseDTO`

NewWorkflowHistoryReadResponseDTO instantiates a new WorkflowHistoryReadResponseDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowHistoryReadResponseDTOWithDefaults

`func NewWorkflowHistoryReadResponseDTOWithDefaults() *WorkflowHistoryReadResponseDTO`

NewWorkflowHistoryReadResponseDTOWithDefaults instantiates a new WorkflowHistoryReadResponseDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *WorkflowHistoryReadResponseDTO) GetStatuses() []WorkflowDocumentStatusDTO`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *WorkflowHistoryReadResponseDTO) GetStatusesOk() (*[]WorkflowDocumentStatusDTO, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *WorkflowHistoryReadResponseDTO) SetStatuses(v []WorkflowDocumentStatusDTO)`

SetStatuses sets Statuses field to given value.

### HasStatuses

`func (o *WorkflowHistoryReadResponseDTO) HasStatuses() bool`

HasStatuses returns a boolean if a field has been set.

### GetWorkflows

`func (o *WorkflowHistoryReadResponseDTO) GetWorkflows() []WorkflowDocumentDTO`

GetWorkflows returns the Workflows field if non-nil, zero value otherwise.

### GetWorkflowsOk

`func (o *WorkflowHistoryReadResponseDTO) GetWorkflowsOk() (*[]WorkflowDocumentDTO, bool)`

GetWorkflowsOk returns a tuple with the Workflows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflows

`func (o *WorkflowHistoryReadResponseDTO) SetWorkflows(v []WorkflowDocumentDTO)`

SetWorkflows sets Workflows field to given value.

### HasWorkflows

`func (o *WorkflowHistoryReadResponseDTO) HasWorkflows() bool`

HasWorkflows returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


