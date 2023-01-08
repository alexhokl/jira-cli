# WorkflowRulesSearch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | Pointer to **string** | Use expand to include additional information in the response. This parameter accepts &#x60;transition&#x60; which, for each rule, returns information about the transition the rule is assigned to. | [optional] 
**RuleIds** | **[]string** | The list of workflow rule IDs. | 
**WorkflowEntityId** | **string** | The workflow ID. | 

## Methods

### NewWorkflowRulesSearch

`func NewWorkflowRulesSearch(ruleIds []string, workflowEntityId string, ) *WorkflowRulesSearch`

NewWorkflowRulesSearch instantiates a new WorkflowRulesSearch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRulesSearchWithDefaults

`func NewWorkflowRulesSearchWithDefaults() *WorkflowRulesSearch`

NewWorkflowRulesSearchWithDefaults instantiates a new WorkflowRulesSearch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpand

`func (o *WorkflowRulesSearch) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *WorkflowRulesSearch) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *WorkflowRulesSearch) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *WorkflowRulesSearch) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetRuleIds

`func (o *WorkflowRulesSearch) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *WorkflowRulesSearch) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *WorkflowRulesSearch) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.


### GetWorkflowEntityId

`func (o *WorkflowRulesSearch) GetWorkflowEntityId() string`

GetWorkflowEntityId returns the WorkflowEntityId field if non-nil, zero value otherwise.

### GetWorkflowEntityIdOk

`func (o *WorkflowRulesSearch) GetWorkflowEntityIdOk() (*string, bool)`

GetWorkflowEntityIdOk returns a tuple with the WorkflowEntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowEntityId

`func (o *WorkflowRulesSearch) SetWorkflowEntityId(v string)`

SetWorkflowEntityId sets WorkflowEntityId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


