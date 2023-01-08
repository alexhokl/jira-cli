# WorkflowSchemeAssociations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | **[]string** | The list of projects that use the workflow scheme. | 
**WorkflowScheme** | [**WorkflowScheme**](WorkflowScheme.md) | The workflow scheme. | 

## Methods

### NewWorkflowSchemeAssociations

`func NewWorkflowSchemeAssociations(projectIds []string, workflowScheme WorkflowScheme, ) *WorkflowSchemeAssociations`

NewWorkflowSchemeAssociations instantiates a new WorkflowSchemeAssociations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeAssociationsWithDefaults

`func NewWorkflowSchemeAssociationsWithDefaults() *WorkflowSchemeAssociations`

NewWorkflowSchemeAssociationsWithDefaults instantiates a new WorkflowSchemeAssociations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *WorkflowSchemeAssociations) GetProjectIds() []string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkflowSchemeAssociations) GetProjectIdsOk() (*[]string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkflowSchemeAssociations) SetProjectIds(v []string)`

SetProjectIds sets ProjectIds field to given value.


### GetWorkflowScheme

`func (o *WorkflowSchemeAssociations) GetWorkflowScheme() WorkflowScheme`

GetWorkflowScheme returns the WorkflowScheme field if non-nil, zero value otherwise.

### GetWorkflowSchemeOk

`func (o *WorkflowSchemeAssociations) GetWorkflowSchemeOk() (*WorkflowScheme, bool)`

GetWorkflowSchemeOk returns a tuple with the WorkflowScheme field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowScheme

`func (o *WorkflowSchemeAssociations) SetWorkflowScheme(v WorkflowScheme)`

SetWorkflowScheme sets WorkflowScheme field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


