# WorkflowPreviewRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeIds** | Pointer to **[]string** | The list of issue type IDs. At most 25 issue type IDs can be specified. | [optional] 
**ProjectId** | **string** | The projectId parameter is required and will be used for permission checks. In addition, you must supply at least one of the following lookup terms: *workflowNames*, *workflowIds*, or *issueTypeIds*. The specified workflows must be associated with the given project. | 
**WorkflowIds** | Pointer to **[]string** | The list of workflow IDs to be returned. At most 25 workflow IDs can be specified. | [optional] 
**WorkflowNames** | Pointer to **[]string** | The list of workflow names to be returned. At most 25 workflow names can be specified. | [optional] 

## Methods

### NewWorkflowPreviewRequest

`func NewWorkflowPreviewRequest(projectId string, ) *WorkflowPreviewRequest`

NewWorkflowPreviewRequest instantiates a new WorkflowPreviewRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowPreviewRequestWithDefaults

`func NewWorkflowPreviewRequestWithDefaults() *WorkflowPreviewRequest`

NewWorkflowPreviewRequestWithDefaults instantiates a new WorkflowPreviewRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeIds

`func (o *WorkflowPreviewRequest) GetIssueTypeIds() []string`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *WorkflowPreviewRequest) GetIssueTypeIdsOk() (*[]string, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *WorkflowPreviewRequest) SetIssueTypeIds(v []string)`

SetIssueTypeIds sets IssueTypeIds field to given value.

### HasIssueTypeIds

`func (o *WorkflowPreviewRequest) HasIssueTypeIds() bool`

HasIssueTypeIds returns a boolean if a field has been set.

### GetProjectId

`func (o *WorkflowPreviewRequest) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *WorkflowPreviewRequest) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *WorkflowPreviewRequest) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.


### GetWorkflowIds

`func (o *WorkflowPreviewRequest) GetWorkflowIds() []string`

GetWorkflowIds returns the WorkflowIds field if non-nil, zero value otherwise.

### GetWorkflowIdsOk

`func (o *WorkflowPreviewRequest) GetWorkflowIdsOk() (*[]string, bool)`

GetWorkflowIdsOk returns a tuple with the WorkflowIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowIds

`func (o *WorkflowPreviewRequest) SetWorkflowIds(v []string)`

SetWorkflowIds sets WorkflowIds field to given value.

### HasWorkflowIds

`func (o *WorkflowPreviewRequest) HasWorkflowIds() bool`

HasWorkflowIds returns a boolean if a field has been set.

### GetWorkflowNames

`func (o *WorkflowPreviewRequest) GetWorkflowNames() []string`

GetWorkflowNames returns the WorkflowNames field if non-nil, zero value otherwise.

### GetWorkflowNamesOk

`func (o *WorkflowPreviewRequest) GetWorkflowNamesOk() (*[]string, bool)`

GetWorkflowNamesOk returns a tuple with the WorkflowNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowNames

`func (o *WorkflowPreviewRequest) SetWorkflowNames(v []string)`

SetWorkflowNames sets WorkflowNames field to given value.

### HasWorkflowNames

`func (o *WorkflowPreviewRequest) HasWorkflowNames() bool`

HasWorkflowNames returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


