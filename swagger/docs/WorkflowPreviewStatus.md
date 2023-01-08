# WorkflowPreviewStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalConfiguration** | Pointer to [**ApprovalConfigurationPreview**](ApprovalConfigurationPreview.md) |  | [optional] 
**Deprecated** | Pointer to **bool** | Whether the status is deprecated. | [optional] 
**Layout** | Pointer to [**WorkflowPreviewLayout**](WorkflowPreviewLayout.md) |  | [optional] 
**StatusReference** | Pointer to **string** | The reference of the status. | [optional] 

## Methods

### NewWorkflowPreviewStatus

`func NewWorkflowPreviewStatus() *WorkflowPreviewStatus`

NewWorkflowPreviewStatus instantiates a new WorkflowPreviewStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPreviewStatusWithDefaults

`func NewWorkflowPreviewStatusWithDefaults() *WorkflowPreviewStatus`

NewWorkflowPreviewStatusWithDefaults instantiates a new WorkflowPreviewStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalConfiguration

`func (o *WorkflowPreviewStatus) GetApprovalConfiguration() ApprovalConfigurationPreview`

GetApprovalConfiguration returns the ApprovalConfiguration field if non-nil, zero value otherwise.

### GetApprovalConfigurationOk

`func (o *WorkflowPreviewStatus) GetApprovalConfigurationOk() (*ApprovalConfigurationPreview, bool)`

GetApprovalConfigurationOk returns a tuple with the ApprovalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfiguration

`func (o *WorkflowPreviewStatus) SetApprovalConfiguration(v ApprovalConfigurationPreview)`

SetApprovalConfiguration sets ApprovalConfiguration field to given value.

### HasApprovalConfiguration

`func (o *WorkflowPreviewStatus) HasApprovalConfiguration() bool`

HasApprovalConfiguration returns a boolean if a field has been set.

### GetDeprecated

`func (o *WorkflowPreviewStatus) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowPreviewStatus) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowPreviewStatus) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowPreviewStatus) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetLayout

`func (o *WorkflowPreviewStatus) GetLayout() WorkflowPreviewLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *WorkflowPreviewStatus) GetLayoutOk() (*WorkflowPreviewLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *WorkflowPreviewStatus) SetLayout(v WorkflowPreviewLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *WorkflowPreviewStatus) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStatusReference

`func (o *WorkflowPreviewStatus) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *WorkflowPreviewStatus) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *WorkflowPreviewStatus) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.

### HasStatusReference

`func (o *WorkflowPreviewStatus) HasStatusReference() bool`

HasStatusReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


