# WorkflowSchemeReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectIds** | Pointer to **[]string** | The list of project IDs to query. | [optional] 
**WorkflowSchemeIds** | Pointer to **[]string** | The list of workflow scheme IDs to query. | [optional] 

## Methods

### NewWorkflowSchemeReadRequest

`func NewWorkflowSchemeReadRequest() *WorkflowSchemeReadRequest`

NewWorkflowSchemeReadRequest instantiates a new WorkflowSchemeReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeReadRequestWithDefaults

`func NewWorkflowSchemeReadRequestWithDefaults() *WorkflowSchemeReadRequest`

NewWorkflowSchemeReadRequestWithDefaults instantiates a new WorkflowSchemeReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectIds

`func (o *WorkflowSchemeReadRequest) GetProjectIds() []*string`

GetProjectIds returns the ProjectIds field if non-nil, zero value otherwise.

### GetProjectIdsOk

`func (o *WorkflowSchemeReadRequest) GetProjectIdsOk() (*[]*string, bool)`

GetProjectIdsOk returns a tuple with the ProjectIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectIds

`func (o *WorkflowSchemeReadRequest) SetProjectIds(v []*string)`

SetProjectIds sets ProjectIds field to given value.

### HasProjectIds

`func (o *WorkflowSchemeReadRequest) HasProjectIds() bool`

HasProjectIds returns a boolean if a field has been set.

### SetProjectIdsNil

`func (o *WorkflowSchemeReadRequest) SetProjectIdsNil(b bool)`

 SetProjectIdsNil sets the value for ProjectIds to be an explicit nil

### UnsetProjectIds
`func (o *WorkflowSchemeReadRequest) UnsetProjectIds()`

UnsetProjectIds ensures that no value is present for ProjectIds, not even an explicit nil
### GetWorkflowSchemeIds

`func (o *WorkflowSchemeReadRequest) GetWorkflowSchemeIds() []*string`

GetWorkflowSchemeIds returns the WorkflowSchemeIds field if non-nil, zero value otherwise.

### GetWorkflowSchemeIdsOk

`func (o *WorkflowSchemeReadRequest) GetWorkflowSchemeIdsOk() (*[]*string, bool)`

GetWorkflowSchemeIdsOk returns a tuple with the WorkflowSchemeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSchemeIds

`func (o *WorkflowSchemeReadRequest) SetWorkflowSchemeIds(v []*string)`

SetWorkflowSchemeIds sets WorkflowSchemeIds field to given value.

### HasWorkflowSchemeIds

`func (o *WorkflowSchemeReadRequest) HasWorkflowSchemeIds() bool`

HasWorkflowSchemeIds returns a boolean if a field has been set.

### SetWorkflowSchemeIdsNil

`func (o *WorkflowSchemeReadRequest) SetWorkflowSchemeIdsNil(b bool)`

 SetWorkflowSchemeIdsNil sets the value for WorkflowSchemeIds to be an explicit nil

### UnsetWorkflowSchemeIds
`func (o *WorkflowSchemeReadRequest) UnsetWorkflowSchemeIds()`

UnsetWorkflowSchemeIds ensures that no value is present for WorkflowSchemeIds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


