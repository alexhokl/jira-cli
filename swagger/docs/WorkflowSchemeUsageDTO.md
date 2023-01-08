# WorkflowSchemeUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | Pointer to **string** | The workflow ID. | [optional] 
**WorkflowSchemes** | Pointer to [**WorkflowSchemeUsagePage**](WorkflowSchemeUsagePage.md) |  | [optional] 

## Methods

### NewWorkflowSchemeUsageDTO

`func NewWorkflowSchemeUsageDTO() *WorkflowSchemeUsageDTO`

NewWorkflowSchemeUsageDTO instantiates a new WorkflowSchemeUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeUsageDTOWithDefaults

`func NewWorkflowSchemeUsageDTOWithDefaults() *WorkflowSchemeUsageDTO`

NewWorkflowSchemeUsageDTOWithDefaults instantiates a new WorkflowSchemeUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowSchemeUsageDTO) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowSchemeUsageDTO) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowSchemeUsageDTO) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowSchemeUsageDTO) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.

### GetWorkflowSchemes

`func (o *WorkflowSchemeUsageDTO) GetWorkflowSchemes() WorkflowSchemeUsagePage`

GetWorkflowSchemes returns the WorkflowSchemes field if non-nil, zero value otherwise.

### GetWorkflowSchemesOk

`func (o *WorkflowSchemeUsageDTO) GetWorkflowSchemesOk() (*WorkflowSchemeUsagePage, bool)`

GetWorkflowSchemesOk returns a tuple with the WorkflowSchemes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSchemes

`func (o *WorkflowSchemeUsageDTO) SetWorkflowSchemes(v WorkflowSchemeUsagePage)`

SetWorkflowSchemes sets WorkflowSchemes field to given value.

### HasWorkflowSchemes

`func (o *WorkflowSchemeUsageDTO) HasWorkflowSchemes() bool`

HasWorkflowSchemes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


