# WorkflowProjectUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**ProjectUsagePage**](ProjectUsagePage.md) |  | [optional] 
**WorkflowId** | Pointer to **string** | The workflow ID. | [optional] 

## Methods

### NewWorkflowProjectUsageDTO

`func NewWorkflowProjectUsageDTO() *WorkflowProjectUsageDTO`

NewWorkflowProjectUsageDTO instantiates a new WorkflowProjectUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowProjectUsageDTOWithDefaults

`func NewWorkflowProjectUsageDTOWithDefaults() *WorkflowProjectUsageDTO`

NewWorkflowProjectUsageDTOWithDefaults instantiates a new WorkflowProjectUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *WorkflowProjectUsageDTO) GetProjects() ProjectUsagePage`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *WorkflowProjectUsageDTO) GetProjectsOk() (*ProjectUsagePage, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *WorkflowProjectUsageDTO) SetProjects(v ProjectUsagePage)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *WorkflowProjectUsageDTO) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowProjectUsageDTO) GetWorkflowId() string`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowProjectUsageDTO) GetWorkflowIdOk() (*string, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowProjectUsageDTO) SetWorkflowId(v string)`

SetWorkflowId sets WorkflowId field to given value.

### HasWorkflowId

`func (o *WorkflowProjectUsageDTO) HasWorkflowId() bool`

HasWorkflowId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


