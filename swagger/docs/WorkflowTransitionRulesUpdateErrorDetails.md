# WorkflowTransitionRulesUpdateErrorDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleUpdateErrors** | **map[string][]string** | A list of transition rule update errors, indexed by the transition rule ID. Any transition rule that appears here wasn&#39;t updated. | 
**UpdateErrors** | **[]string** | The list of errors that specify why the workflow update failed. The workflow was not updated if the list contains any entries. | 
**WorkflowId** | [**WorkflowId**](WorkflowId.md) |  | 

## Methods

### NewWorkflowTransitionRulesUpdateErrorDetails

`func NewWorkflowTransitionRulesUpdateErrorDetails(ruleUpdateErrors map[string][]string, updateErrors []string, workflowId WorkflowId, ) *WorkflowTransitionRulesUpdateErrorDetails`

NewWorkflowTransitionRulesUpdateErrorDetails instantiates a new WorkflowTransitionRulesUpdateErrorDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionRulesUpdateErrorDetailsWithDefaults

`func NewWorkflowTransitionRulesUpdateErrorDetailsWithDefaults() *WorkflowTransitionRulesUpdateErrorDetails`

NewWorkflowTransitionRulesUpdateErrorDetailsWithDefaults instantiates a new WorkflowTransitionRulesUpdateErrorDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleUpdateErrors

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetRuleUpdateErrors() map[string][]string`

GetRuleUpdateErrors returns the RuleUpdateErrors field if non-nil, zero value otherwise.

### GetRuleUpdateErrorsOk

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetRuleUpdateErrorsOk() (*map[string][]string, bool)`

GetRuleUpdateErrorsOk returns a tuple with the RuleUpdateErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleUpdateErrors

`func (o *WorkflowTransitionRulesUpdateErrorDetails) SetRuleUpdateErrors(v map[string][]string)`

SetRuleUpdateErrors sets RuleUpdateErrors field to given value.


### GetUpdateErrors

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetUpdateErrors() []string`

GetUpdateErrors returns the UpdateErrors field if non-nil, zero value otherwise.

### GetUpdateErrorsOk

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetUpdateErrorsOk() (*[]string, bool)`

GetUpdateErrorsOk returns a tuple with the UpdateErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateErrors

`func (o *WorkflowTransitionRulesUpdateErrorDetails) SetUpdateErrors(v []string)`

SetUpdateErrors sets UpdateErrors field to given value.


### GetWorkflowId

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetWorkflowId() WorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowTransitionRulesUpdateErrorDetails) GetWorkflowIdOk() (*WorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowTransitionRulesUpdateErrorDetails) SetWorkflowId(v WorkflowId)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


