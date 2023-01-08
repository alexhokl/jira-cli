# WorkflowMetadataAndIssueTypeRestModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueTypeIds** | **[]string** | The list of issue type IDs for the mapping. | 
**Workflow** | [**WorkflowMetadataRestModel**](WorkflowMetadataRestModel.md) |  | 

## Methods

### NewWorkflowMetadataAndIssueTypeRestModel

`func NewWorkflowMetadataAndIssueTypeRestModel(issueTypeIds []string, workflow WorkflowMetadataRestModel, ) *WorkflowMetadataAndIssueTypeRestModel`

NewWorkflowMetadataAndIssueTypeRestModel instantiates a new WorkflowMetadataAndIssueTypeRestModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowMetadataAndIssueTypeRestModelWithDefaults

`func NewWorkflowMetadataAndIssueTypeRestModelWithDefaults() *WorkflowMetadataAndIssueTypeRestModel`

NewWorkflowMetadataAndIssueTypeRestModelWithDefaults instantiates a new WorkflowMetadataAndIssueTypeRestModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueTypeIds

`func (o *WorkflowMetadataAndIssueTypeRestModel) GetIssueTypeIds() []string`

GetIssueTypeIds returns the IssueTypeIds field if non-nil, zero value otherwise.

### GetIssueTypeIdsOk

`func (o *WorkflowMetadataAndIssueTypeRestModel) GetIssueTypeIdsOk() (*[]string, bool)`

GetIssueTypeIdsOk returns a tuple with the IssueTypeIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypeIds

`func (o *WorkflowMetadataAndIssueTypeRestModel) SetIssueTypeIds(v []string)`

SetIssueTypeIds sets IssueTypeIds field to given value.


### GetWorkflow

`func (o *WorkflowMetadataAndIssueTypeRestModel) GetWorkflow() WorkflowMetadataRestModel`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *WorkflowMetadataAndIssueTypeRestModel) GetWorkflowOk() (*WorkflowMetadataRestModel, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *WorkflowMetadataAndIssueTypeRestModel) SetWorkflow(v WorkflowMetadataRestModel)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


