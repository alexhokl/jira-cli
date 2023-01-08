# DefaultWorkflow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UpdateDraftIfNeeded** | Pointer to **bool** | Whether a draft workflow scheme is created or updated when updating an active workflow scheme. The draft is updated with the new default workflow. Defaults to &#x60;false&#x60;. | [optional] 
**Workflow** | **string** | The name of the workflow to set as the default workflow. | 

## Methods

### NewDefaultWorkflow

`func NewDefaultWorkflow(workflow string, ) *DefaultWorkflow`

NewDefaultWorkflow instantiates a new DefaultWorkflow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDefaultWorkflowWithDefaults

`func NewDefaultWorkflowWithDefaults() *DefaultWorkflow`

NewDefaultWorkflowWithDefaults instantiates a new DefaultWorkflow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdateDraftIfNeeded

`func (o *DefaultWorkflow) GetUpdateDraftIfNeeded() bool`

GetUpdateDraftIfNeeded returns the UpdateDraftIfNeeded field if non-nil, zero value otherwise.

### GetUpdateDraftIfNeededOk

`func (o *DefaultWorkflow) GetUpdateDraftIfNeededOk() (*bool, bool)`

GetUpdateDraftIfNeededOk returns a tuple with the UpdateDraftIfNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateDraftIfNeeded

`func (o *DefaultWorkflow) SetUpdateDraftIfNeeded(v bool)`

SetUpdateDraftIfNeeded sets UpdateDraftIfNeeded field to given value.

### HasUpdateDraftIfNeeded

`func (o *DefaultWorkflow) HasUpdateDraftIfNeeded() bool`

HasUpdateDraftIfNeeded returns a boolean if a field has been set.

### GetWorkflow

`func (o *DefaultWorkflow) GetWorkflow() string`

GetWorkflow returns the Workflow field if non-nil, zero value otherwise.

### GetWorkflowOk

`func (o *DefaultWorkflow) GetWorkflowOk() (*string, bool)`

GetWorkflowOk returns a tuple with the Workflow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflow

`func (o *DefaultWorkflow) SetWorkflow(v string)`

SetWorkflow sets Workflow field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


