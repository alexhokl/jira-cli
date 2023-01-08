# WorkflowReadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectAndIssueTypes** | Pointer to [**[]ProjectAndIssueTypePair**](ProjectAndIssueTypePair.md) | The list of projects and issue types to query. | [optional] 
**WorkflowIds** | Pointer to **[]string** | The list of workflow IDs to query. | [optional] 
**WorkflowNames** | Pointer to **[]string** | The list of workflow names to query. | [optional] 

## Methods

### NewWorkflowReadRequest

`func NewWorkflowReadRequest() *WorkflowReadRequest`

NewWorkflowReadRequest instantiates a new WorkflowReadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowReadRequestWithDefaults

`func NewWorkflowReadRequestWithDefaults() *WorkflowReadRequest`

NewWorkflowReadRequestWithDefaults instantiates a new WorkflowReadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjectAndIssueTypes

`func (o *WorkflowReadRequest) GetProjectAndIssueTypes() []ProjectAndIssueTypePair`

GetProjectAndIssueTypes returns the ProjectAndIssueTypes field if non-nil, zero value otherwise.

### GetProjectAndIssueTypesOk

`func (o *WorkflowReadRequest) GetProjectAndIssueTypesOk() (*[]ProjectAndIssueTypePair, bool)`

GetProjectAndIssueTypesOk returns a tuple with the ProjectAndIssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectAndIssueTypes

`func (o *WorkflowReadRequest) SetProjectAndIssueTypes(v []ProjectAndIssueTypePair)`

SetProjectAndIssueTypes sets ProjectAndIssueTypes field to given value.

### HasProjectAndIssueTypes

`func (o *WorkflowReadRequest) HasProjectAndIssueTypes() bool`

HasProjectAndIssueTypes returns a boolean if a field has been set.

### GetWorkflowIds

`func (o *WorkflowReadRequest) GetWorkflowIds() []string`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *WorkflowReadRequest) GetWorkflowIdsOk() (*[]string, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *WorkflowReadRequest) SetWorkflowIds(v []string)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *WorkflowReadRequest) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.

### GetWorkflowNames

`func (o *WorkflowReadRequest) GetWorkflowNames() []string`

GetWorkflowNames returns the WorkflowNames field if non-nil, zero value otherwise.

### GetWorkflowNamesOk

`func (o *WorkflowReadRequest) GetWorkflowNamesOk() (*[]string, bool)`

GetWorkflowNamesOk returns a tuple with the WorkflowNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNames

`func (o *WorkflowReadRequest) SetWorkflowNames(v []string)`

SetWorkflowNames sets WorkflowNames field to given value.

### HasWorkflowNames

`func (o *WorkflowReadRequest) HasWorkflowNames() bool`

HasWorkflowNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


