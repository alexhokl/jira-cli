# WorkflowTransitionRulesDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkflowId** | [**WorkflowId**](WorkflowId.md) |  | 
**WorkflowRuleIds** | **[]string** | The list of connect workflow rule IDs. | 

## Methods

### NewWorkflowTransitionRulesDetails

`func NewWorkflowTransitionRulesDetails(workflowId WorkflowId, workflowRuleIds []string, ) *WorkflowTransitionRulesDetails`

NewWorkflowTransitionRulesDetails instantiates a new WorkflowTransitionRulesDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionRulesDetailsWithDefaults

`func NewWorkflowTransitionRulesDetailsWithDefaults() *WorkflowTransitionRulesDetails`

NewWorkflowTransitionRulesDetailsWithDefaults instantiates a new WorkflowTransitionRulesDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkflowId

`func (o *WorkflowTransitionRulesDetails) GetWorkflowId() WorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowTransitionRulesDetails) GetWorkflowIdOk() (*WorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowTransitionRulesDetails) SetWorkflowId(v WorkflowId)`

SetWorkflowId sets WorkflowId field to given value.


### GetWorkflowRuleIds

`func (o *WorkflowTransitionRulesDetails) GetWorkflowRuleIds() []string`

GetWorkflowRuleIds returns the WorkflowRuleIds field if non-nil, zero value otherwise.

### GetWorkflowRuleIdsOk

`func (o *WorkflowTransitionRulesDetails) GetWorkflowRuleIdsOk() (*[]string, bool)`

GetWorkflowRuleIdsOk returns a tuple with the WorkflowRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowRuleIds

`func (o *WorkflowTransitionRulesDetails) SetWorkflowRuleIds(v []string)`

SetWorkflowRuleIds sets WorkflowRuleIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


