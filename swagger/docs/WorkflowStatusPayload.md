# WorkflowStatusPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Layout** | Pointer to [**WorkflowStatusLayoutPayload**](WorkflowStatusLayoutPayload.md) |  | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** | The properties of the workflow status. | [optional] 

## Methods

### NewWorkflowStatusPayload

`func NewWorkflowStatusPayload() *WorkflowStatusPayload`

NewWorkflowStatusPayload instantiates a new WorkflowStatusPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowStatusPayloadWithDefaults

`func NewWorkflowStatusPayloadWithDefaults() *WorkflowStatusPayload`

NewWorkflowStatusPayloadWithDefaults instantiates a new WorkflowStatusPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLayout

`func (o *WorkflowStatusPayload) GetLayout() WorkflowStatusLayoutPayload`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *WorkflowStatusPayload) GetLayoutOk() (*WorkflowStatusLayoutPayload, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *WorkflowStatusPayload) SetLayout(v WorkflowStatusLayoutPayload)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *WorkflowStatusPayload) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetPcri

`func (o *WorkflowStatusPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *WorkflowStatusPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *WorkflowStatusPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *WorkflowStatusPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetProperties

`func (o *WorkflowStatusPayload) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *WorkflowStatusPayload) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *WorkflowStatusPayload) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *WorkflowStatusPayload) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


