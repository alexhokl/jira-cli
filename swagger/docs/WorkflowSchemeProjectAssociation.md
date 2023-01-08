# WorkflowSchemeProjectAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectId** | **string** | The ID of the project. | 
**WorkflowSchemeId** | Pointer to **string** | The ID of the workflow scheme. If the workflow scheme ID is &#x60;null&#x60;, the operation assigns the default workflow scheme. | [optional] 

## Methods

### NewWorkflowSchemeProjectAssociation

`func NewWorkflowSchemeProjectAssociation(projectId string, ) *WorkflowSchemeProjectAssociation`

NewWorkflowSchemeProjectAssociation instantiates a new WorkflowSchemeProjectAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeProjectAssociationWithDefaults

`func NewWorkflowSchemeProjectAssociationWithDefaults() *WorkflowSchemeProjectAssociation`

NewWorkflowSchemeProjectAssociationWithDefaults instantiates a new WorkflowSchemeProjectAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectId

`func (o *WorkflowSchemeProjectAssociation) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkflowSchemeProjectAssociation) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkflowSchemeProjectAssociation) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetWorkflowSchemeId

`func (o *WorkflowSchemeProjectAssociation) GetWorkflowSchemeId() string`

GetWorkflowSchemeId returns the WorkflowSchemeId field if non-nil, zero value otherwise.

### GetWorkflowSchemeIdOk

`func (o *WorkflowSchemeProjectAssociation) GetWorkflowSchemeIdOk() (*string, bool)`

GetWorkflowSchemeIdOk returns a tuple with the WorkflowSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSchemeId

`func (o *WorkflowSchemeProjectAssociation) SetWorkflowSchemeId(v string)`

SetWorkflowSchemeId sets WorkflowSchemeId field to given value.

### HasWorkflowSchemeId

`func (o *WorkflowSchemeProjectAssociation) HasWorkflowSchemeId() bool`

HasWorkflowSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


