# WorkflowSchemeUpdateRequiredMappingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflowId** | Pointer to **NullableString** | The ID of the new default workflow for this workflow scheme. Only used in global-scoped workflow schemes. If it isn&#39;t specified, is set to *Jira Workflow (jira)*. | [optional] 
**Id** | **string** | The ID of the workflow scheme. | 
**WorkflowsForIssueTypes** | [**[]WorkflowSchemeAssociation**](WorkflowSchemeAssociation.md) | The new workflow to issue type mappings for this workflow scheme. | 

## Methods

### NewWorkflowSchemeUpdateRequiredMappingsRequest

`func NewWorkflowSchemeUpdateRequiredMappingsRequest(id string, workflowsForIssueTypes []WorkflowSchemeAssociation, ) *WorkflowSchemeUpdateRequiredMappingsRequest`

NewWorkflowSchemeUpdateRequiredMappingsRequest instantiates a new WorkflowSchemeUpdateRequiredMappingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeUpdateRequiredMappingsRequestWithDefaults

`func NewWorkflowSchemeUpdateRequiredMappingsRequestWithDefaults() *WorkflowSchemeUpdateRequiredMappingsRequest`

NewWorkflowSchemeUpdateRequiredMappingsRequestWithDefaults instantiates a new WorkflowSchemeUpdateRequiredMappingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetDefaultWorkflowId() string`

GetDefaultWorkflowId returns the DefaultWorkflowId field if non-nil, zero value otherwise.

### GetDefaultWorkflowIdOk

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetDefaultWorkflowIdOk() (*string, bool)`

GetDefaultWorkflowIdOk returns a tuple with the DefaultWorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) SetDefaultWorkflowId(v string)`

SetDefaultWorkflowId sets DefaultWorkflowId field to given value.

### HasDefaultWorkflowId

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) HasDefaultWorkflowId() bool`

HasDefaultWorkflowId returns a boolean if a field has been set.

### SetDefaultWorkflowIdNil

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) SetDefaultWorkflowIdNil(b bool)`

 SetDefaultWorkflowIdNil sets the value for DefaultWorkflowId to be an explicit nil

### UnsetDefaultWorkflowId
`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) UnsetDefaultWorkflowId()`

UnsetDefaultWorkflowId ensures that no value is present for DefaultWorkflowId, not even an explicit nil
### GetId

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) SetId(v string)`

SetId sets Id field to given value.


### GetWorkflowsForIssueTypes

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetWorkflowsForIssueTypes() []WorkflowSchemeAssociation`

GetWorkflowsForIssueTypes returns the WorkflowsForIssueTypes field if non-nil, zero value otherwise.

### GetWorkflowsForIssueTypesOk

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) GetWorkflowsForIssueTypesOk() (*[]WorkflowSchemeAssociation, bool)`

GetWorkflowsForIssueTypesOk returns a tuple with the WorkflowsForIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsForIssueTypes

`func (o *WorkflowSchemeUpdateRequiredMappingsRequest) SetWorkflowsForIssueTypes(v []WorkflowSchemeAssociation)`

SetWorkflowsForIssueTypes sets WorkflowsForIssueTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


