# WorkflowSchemeProjectUsageDTO

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Projects** | Pointer to [**ProjectUsagePage**](ProjectUsagePage.md) |  | [optional] 
**WorkflowSchemeId** | Pointer to **string** | The workflow scheme ID. | [optional] 

## Methods

### NewWorkflowSchemeProjectUsageDTO

`func NewWorkflowSchemeProjectUsageDTO() *WorkflowSchemeProjectUsageDTO`

NewWorkflowSchemeProjectUsageDTO instantiates a new WorkflowSchemeProjectUsageDTO object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSchemeProjectUsageDTOWithDefaults

`func NewWorkflowSchemeProjectUsageDTOWithDefaults() *WorkflowSchemeProjectUsageDTO`

NewWorkflowSchemeProjectUsageDTOWithDefaults instantiates a new WorkflowSchemeProjectUsageDTO object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProjects

`func (o *WorkflowSchemeProjectUsageDTO) GetProjects() ProjectUsagePage`

GetProjects returns the Projects field if non-nil, zero value otherwise.

### GetProjectsOk

`func (o *WorkflowSchemeProjectUsageDTO) GetProjectsOk() (*ProjectUsagePage, bool)`

GetProjectsOk returns a tuple with the Projects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjects

`func (o *WorkflowSchemeProjectUsageDTO) SetProjects(v ProjectUsagePage)`

SetProjects sets Projects field to given value.

### HasProjects

`func (o *WorkflowSchemeProjectUsageDTO) HasProjects() bool`

HasProjects returns a boolean if a field has been set.

### GetWorkflowSchemeId

`func (o *WorkflowSchemeProjectUsageDTO) GetWorkflowSchemeId() string`

GetWorkflowSchemeId returns the WorkflowSchemeId field if non-nil, zero value otherwise.

### GetWorkflowSchemeIdOk

`func (o *WorkflowSchemeProjectUsageDTO) GetWorkflowSchemeIdOk() (*string, bool)`

GetWorkflowSchemeIdOk returns a tuple with the WorkflowSchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowSchemeId

`func (o *WorkflowSchemeProjectUsageDTO) SetWorkflowSchemeId(v string)`

SetWorkflowSchemeId sets WorkflowSchemeId field to given value.

### HasWorkflowSchemeId

`func (o *WorkflowSchemeProjectUsageDTO) HasWorkflowSchemeId() bool`

HasWorkflowSchemeId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


