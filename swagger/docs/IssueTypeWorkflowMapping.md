# IssueTypeWorkflowMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IssueType** | Pointer to **string** | The ID of the issue type. Not required if updating the issue type-workflow mapping. | [optional] 
**UpdateDraftIfNeeded** | Pointer to **bool** | Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to &#x60;false&#x60;. Only applicable when updating the workflow-issue types mapping. | [optional] 
**Workflow** | Pointer to **string** | The name of the workflow. | [optional] 

## Methods

### NewIssueTypeWorkflowMapping

`func NewIssueTypeWorkflowMapping() *IssueTypeWorkflowMapping`

NewIssueTypeWorkflowMapping instantiates a new IssueTypeWorkflowMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeWorkflowMappingWithDefaults

`func NewIssueTypeWorkflowMappingWithDefaults() *IssueTypeWorkflowMapping`

NewIssueTypeWorkflowMappingWithDefaults instantiates a new IssueTypeWorkflowMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIssueType

`func (o *IssueTypeWorkflowMapping) GetIssueType() string`

GetIssueType returns the IssueType field if non-nil, zero value otherwise.

### GetIssueTypeOk

`func (o *IssueTypeWorkflowMapping) GetIssueTypeOk() (*string, bool)`

GetIssueTypeOk returns a tuple with the IssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueType

`func (o *IssueTypeWorkflowMapping) SetIssueType(v string)`

SetIssueType sets IssueType field to given value.

### HasIssueType

`func (o *IssueTypeWorkflowMapping) HasIssueType() bool`

HasIssueType returns a boolean if a field has been set.

### GetUpdateDraftIfNeeded

`func (o *IssueTypeWorkflowMapping) GetUpdateDraftIfNeeded() bool`

GetUpdateDraftIfNeeded returns the UpdateDraftIfNeeded field if non-nil, zero value otherwise.

### GetUpdateDraftIfNeededOk

`func (o *IssueTypeWorkflowMapping) GetUpdateDraftIfNeededOk() (*bool, bool)`

GetUpdateDraftIfNeededOk returns a tuple with the UpdateDraftIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDraftIfNeeded

`func (o *IssueTypeWorkflowMapping) SetUpdateDraftIfNeeded(v bool)`

SetUpdateDraftIfNeeded sets UpdateDraftIfNeeded field to given value.

### HasUpdateDraftIfNeeded

`func (o *IssueTypeWorkflowMapping) HasUpdateDraftIfNeeded() bool`

HasUpdateDraftIfNeeded returns a boolean if a field has been set.

### GetWorkflow

`func (o *IssueTypeWorkflowMapping) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *IssueTypeWorkflowMapping) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *IssueTypeWorkflowMapping) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *IssueTypeWorkflowMapping) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


