# IssueTypesWorkflowMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultMapping** | Pointer to **bool** | Whether the workflow is the default workflow for the workflow scheme. | [optional] 
**IssueTypes** | Pointer to **[]string** | The list of issue type IDs. | [optional] 
**UpdateDraftIfNeeded** | Pointer to **bool** | Whether a draft workflow scheme is created or updated when updating an active workflow scheme. The draft is updated with the new workflow-issue types mapping. Defaults to &#x60;false&#x60;. | [optional] 
**Workflow** | Pointer to **string** | The name of the workflow. Optional if updating the workflow-issue types mapping. | [optional] 

## Methods

### NewIssueTypesWorkflowMapping

`func NewIssueTypesWorkflowMapping() *IssueTypesWorkflowMapping`

NewIssueTypesWorkflowMapping instantiates a new IssueTypesWorkflowMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypesWorkflowMappingWithDefaults

`func NewIssueTypesWorkflowMappingWithDefaults() *IssueTypesWorkflowMapping`

NewIssueTypesWorkflowMappingWithDefaults instantiates a new IssueTypesWorkflowMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultMapping

`func (o *IssueTypesWorkflowMapping) GetDefaultMapping() bool`

GetDefaultMapping returns the DefaultMapping field if non-nil, zero value otherwise.

### GetDefaultMappingOk

`func (o *IssueTypesWorkflowMapping) GetDefaultMappingOk() (*bool, bool)`

GetDefaultMappingOk returns a tuple with the DefaultMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultMapping

`func (o *IssueTypesWorkflowMapping) SetDefaultMapping(v bool)`

SetDefaultMapping sets DefaultMapping field to given value.

### HasDefaultMapping

`func (o *IssueTypesWorkflowMapping) HasDefaultMapping() bool`

HasDefaultMapping returns a boolean if a field has been set.

### GetIssueTypes

`func (o *IssueTypesWorkflowMapping) GetIssueTypes() []string`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *IssueTypesWorkflowMapping) GetIssueTypesOk() (*[]string, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *IssueTypesWorkflowMapping) SetIssueTypes(v []string)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *IssueTypesWorkflowMapping) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetUpdateDraftIfNeeded

`func (o *IssueTypesWorkflowMapping) GetUpdateDraftIfNeeded() bool`

GetUpdateDraftIfNeeded returns the UpdateDraftIfNeeded field if non-nil, zero value otherwise.

### GetUpdateDraftIfNeededOk

`func (o *IssueTypesWorkflowMapping) GetUpdateDraftIfNeededOk() (*bool, bool)`

GetUpdateDraftIfNeededOk returns a tuple with the UpdateDraftIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDraftIfNeeded

`func (o *IssueTypesWorkflowMapping) SetUpdateDraftIfNeeded(v bool)`

SetUpdateDraftIfNeeded sets UpdateDraftIfNeeded field to given value.

### HasUpdateDraftIfNeeded

`func (o *IssueTypesWorkflowMapping) HasUpdateDraftIfNeeded() bool`

HasUpdateDraftIfNeeded returns a boolean if a field has been set.

### GetWorkflow

`func (o *IssueTypesWorkflowMapping) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *IssueTypesWorkflowMapping) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *IssueTypesWorkflowMapping) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.

### HasWorkflow

`func (o *IssueTypesWorkflowMapping) HasWorkflow() bool`

HasWorkflow returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


