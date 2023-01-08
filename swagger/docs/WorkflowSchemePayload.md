# WorkflowSchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflow** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Description** | Pointer to **string** | The description of the workflow scheme | [optional] 
**ExplicitMappings** | Pointer to [**map[string]ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) | Association between issuetypes and workflows | [optional] 
**Name** | Pointer to **string** | The name of the workflow scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewWorkflowSchemePayload

`func NewWorkflowSchemePayload() *WorkflowSchemePayload`

NewWorkflowSchemePayload instantiates a new WorkflowSchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemePayloadWithDefaults

`func NewWorkflowSchemePayloadWithDefaults() *WorkflowSchemePayload`

NewWorkflowSchemePayloadWithDefaults instantiates a new WorkflowSchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflow

`func (o *WorkflowSchemePayload) GetDefaultWorkflow() ProjectCreateResourceIdentifier`

GetDefaultWorkflow returns the DefaultWorkflow field if non-nil, zero value otherwise.

### GetDefaultWorkflowOk

`func (o *WorkflowSchemePayload) GetDefaultWorkflowOk() (*ProjectCreateResourceIdentifier, bool)`

GetDefaultWorkflowOk returns a tuple with the DefaultWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflow

`func (o *WorkflowSchemePayload) SetDefaultWorkflow(v ProjectCreateResourceIdentifier)`

SetDefaultWorkflow sets DefaultWorkflow field to given value.

### HasDefaultWorkflow

`func (o *WorkflowSchemePayload) HasDefaultWorkflow() bool`

HasDefaultWorkflow returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExplicitMappings

`func (o *WorkflowSchemePayload) GetExplicitMappings() map[string]ProjectCreateResourceIdentifier`

GetExplicitMappings returns the ExplicitMappings field if non-nil, zero value otherwise.

### GetExplicitMappingsOk

`func (o *WorkflowSchemePayload) GetExplicitMappingsOk() (*map[string]ProjectCreateResourceIdentifier, bool)`

GetExplicitMappingsOk returns a tuple with the ExplicitMappings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExplicitMappings

`func (o *WorkflowSchemePayload) SetExplicitMappings(v map[string]ProjectCreateResourceIdentifier)`

SetExplicitMappings sets ExplicitMappings field to given value.

### HasExplicitMappings

`func (o *WorkflowSchemePayload) HasExplicitMappings() bool`

HasExplicitMappings returns a boolean if a field has been set.

### GetName

`func (o *WorkflowSchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *WorkflowSchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *WorkflowSchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *WorkflowSchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *WorkflowSchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *WorkflowSchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


