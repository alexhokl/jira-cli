# WorkflowReferenceStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApprovalConfiguration** | Pointer to [**NullableApprovalConfiguration**](ApprovalConfiguration.md) |  | [optional] 
**Deprecated** | Pointer to **bool** | Indicates if the status is deprecated. | [optional] 
**Layout** | Pointer to [**NullableWorkflowStatusLayout**](WorkflowStatusLayout.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** | The properties associated with the status. | [optional] 
**StatusReference** | Pointer to **string** | The reference of the status. | [optional] 

## Methods

### NewWorkflowReferenceStatus

`func NewWorkflowReferenceStatus() *WorkflowReferenceStatus`

NewWorkflowReferenceStatus instantiates a new WorkflowReferenceStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowReferenceStatusWithDefaults

`func NewWorkflowReferenceStatusWithDefaults() *WorkflowReferenceStatus`

NewWorkflowReferenceStatusWithDefaults instantiates a new WorkflowReferenceStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovalConfiguration

`func (o *WorkflowReferenceStatus) GetApprovalConfiguration() ApprovalConfiguration`

GetApprovalConfiguration returns the ApprovalConfiguration field if non-nil, zero value otherwise.

### GetApprovalConfigurationOk

`func (o *WorkflowReferenceStatus) GetApprovalConfigurationOk() (*ApprovalConfiguration, bool)`

GetApprovalConfigurationOk returns a tuple with the ApprovalConfiguration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalConfiguration

`func (o *WorkflowReferenceStatus) SetApprovalConfiguration(v ApprovalConfiguration)`

SetApprovalConfiguration sets ApprovalConfiguration field to given value.

### HasApprovalConfiguration

`func (o *WorkflowReferenceStatus) HasApprovalConfiguration() bool`

HasApprovalConfiguration returns a boolean if a field has been set.

### SetApprovalConfigurationNil

`func (o *WorkflowReferenceStatus) SetApprovalConfigurationNil(b bool)`

 SetApprovalConfigurationNil sets the value for ApprovalConfiguration to be an explicit nil

### UnsetApprovalConfiguration
`func (o *WorkflowReferenceStatus) UnsetApprovalConfiguration()`

UnsetApprovalConfiguration ensures that no value is present for ApprovalConfiguration, not even an explicit nil
### GetDeprecated

`func (o *WorkflowReferenceStatus) GetDeprecated() bool`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *WorkflowReferenceStatus) GetDeprecatedOk() (*bool, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *WorkflowReferenceStatus) SetDeprecated(v bool)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *WorkflowReferenceStatus) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetLayout

`func (o *WorkflowReferenceStatus) GetLayout() WorkflowStatusLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *WorkflowReferenceStatus) GetLayoutOk() (*WorkflowStatusLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *WorkflowReferenceStatus) SetLayout(v WorkflowStatusLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *WorkflowReferenceStatus) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### SetLayoutNil

`func (o *WorkflowReferenceStatus) SetLayoutNil(b bool)`

 SetLayoutNil sets the value for Layout to be an explicit nil

### UnsetLayout
`func (o *WorkflowReferenceStatus) UnsetLayout()`

UnsetLayout ensures that no value is present for Layout, not even an explicit nil
### GetProperties

`func (o *WorkflowReferenceStatus) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowReferenceStatus) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowReferenceStatus) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowReferenceStatus) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatusReference

`func (o *WorkflowReferenceStatus) GetStatusReference() string`

GetStatusReference returns the StatusReference field if non-nil, zero value otherwise.

### GetStatusReferenceOk

`func (o *WorkflowReferenceStatus) GetStatusReferenceOk() (*string, bool)`

GetStatusReferenceOk returns a tuple with the StatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusReference

`func (o *WorkflowReferenceStatus) SetStatusReference(v string)`

SetStatusReference sets StatusReference field to given value.

### HasStatusReference

`func (o *WorkflowReferenceStatus) HasStatusReference() bool`

HasStatusReference returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


