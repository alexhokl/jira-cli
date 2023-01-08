# WorkflowSchemeReadResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultWorkflow** | Pointer to [**WorkflowMetadataRestModel**](WorkflowMetadataRestModel.md) |  | [optional] 
**Description** | Pointer to **NullableString** | The description of the workflow scheme. | [optional] 
**Id** | **string** | The ID of the workflow scheme. | 
**Name** | **string** | The name of the workflow scheme. | 
**Scope** | [**WorkflowScope**](WorkflowScope.md) |  | 
**TaskId** | Pointer to **NullableString** | Indicates if there&#39;s an [asynchronous task](#async-operations) for this workflow scheme. | [optional] 
**Version** | [**DocumentVersion**](DocumentVersion.md) |  | 
**WorkflowsForIssueTypes** | [**[]WorkflowMetadataAndIssueTypeRestModel**](WorkflowMetadataAndIssueTypeRestModel.md) | Mappings from workflows to issue types. | 

## Methods

### NewWorkflowSchemeReadResponse

`func NewWorkflowSchemeReadResponse(id string, name string, scope WorkflowScope, version DocumentVersion, workflowsForIssueTypes []WorkflowMetadataAndIssueTypeRestModel, ) *WorkflowSchemeReadResponse`

NewWorkflowSchemeReadResponse instantiates a new WorkflowSchemeReadResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeReadResponseWithDefaults

`func NewWorkflowSchemeReadResponseWithDefaults() *WorkflowSchemeReadResponse`

NewWorkflowSchemeReadResponseWithDefaults instantiates a new WorkflowSchemeReadResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultWorkflow

`func (o *WorkflowSchemeReadResponse) GetDefaultWorkflow() WorkflowMetadataRestModel`

GetDefaultWorkflow returns the DefaultWorkflow field if non-nil, zero value otherwise.

### GetDefaultWorkflowOk

`func (o *WorkflowSchemeReadResponse) GetDefaultWorkflowOk() (*WorkflowMetadataRestModel, bool)`

GetDefaultWorkflowOk returns a tuple with the DefaultWorkflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultWorkflow

`func (o *WorkflowSchemeReadResponse) SetDefaultWorkflow(v WorkflowMetadataRestModel)`

SetDefaultWorkflow sets DefaultWorkflow field to given value.

### HasDefaultWorkflow

`func (o *WorkflowSchemeReadResponse) HasDefaultWorkflow() bool`

HasDefaultWorkflow returns a boolean if a field has been set.

### GetDescription

`func (o *WorkflowSchemeReadResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *WorkflowSchemeReadResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *WorkflowSchemeReadResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *WorkflowSchemeReadResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *WorkflowSchemeReadResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *WorkflowSchemeReadResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *WorkflowSchemeReadResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowSchemeReadResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowSchemeReadResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *WorkflowSchemeReadResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *WorkflowSchemeReadResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *WorkflowSchemeReadResponse) SetName(v string)`

SetName sets Name field to given value.


### GetScope

`func (o *WorkflowSchemeReadResponse) GetScope() WorkflowScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *WorkflowSchemeReadResponse) GetScopeOk() (*WorkflowScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *WorkflowSchemeReadResponse) SetScope(v WorkflowScope)`

SetScope sets Scope field to given value.


### GetTaskId

`func (o *WorkflowSchemeReadResponse) GetTaskId() string`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *WorkflowSchemeReadResponse) GetTaskIdOk() (*string, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *WorkflowSchemeReadResponse) SetTaskId(v string)`

SetTaskId sets TaskId field to given value.

### HasTaskId

`func (o *WorkflowSchemeReadResponse) HasTaskId() bool`

HasTaskId returns a boolean if a field has been set.

### SetTaskIdNil

`func (o *WorkflowSchemeReadResponse) SetTaskIdNil(b bool)`

 SetTaskIdNil sets the value for TaskId to be an explicit nil

### UnsetTaskId
`func (o *WorkflowSchemeReadResponse) UnsetTaskId()`

UnsetTaskId ensures that no value is present for TaskId, not even an explicit nil
### GetVersion

`func (o *WorkflowSchemeReadResponse) GetVersion() DocumentVersion`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *WorkflowSchemeReadResponse) GetVersionOk() (*DocumentVersion, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *WorkflowSchemeReadResponse) SetVersion(v DocumentVersion)`

SetVersion sets Version field to given value.


### GetWorkflowsForIssueTypes

`func (o *WorkflowSchemeReadResponse) GetWorkflowsForIssueTypes() []WorkflowMetadataAndIssueTypeRestModel`

GetWorkflowsForIssueTypes returns the WorkflowsForIssueTypes field if non-nil, zero value otherwise.

### GetWorkflowsForIssueTypesOk

`func (o *WorkflowSchemeReadResponse) GetWorkflowsForIssueTypesOk() (*[]WorkflowMetadataAndIssueTypeRestModel, bool)`

GetWorkflowsForIssueTypesOk returns a tuple with the WorkflowsForIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowsForIssueTypes

`func (o *WorkflowSchemeReadResponse) SetWorkflowsForIssueTypes(v []WorkflowMetadataAndIssueTypeRestModel)`

SetWorkflowsForIssueTypes sets WorkflowsForIssueTypes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


