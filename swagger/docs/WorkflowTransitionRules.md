# WorkflowTransitionRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]AppWorkflowTransitionRule**](AppWorkflowTransitionRule.md) | The list of conditions within the workflow. | [optional] 
**PostFunctions** | Pointer to [**[]AppWorkflowTransitionRule**](AppWorkflowTransitionRule.md) | The list of post functions within the workflow. | [optional] 
**Validators** | Pointer to [**[]AppWorkflowTransitionRule**](AppWorkflowTransitionRule.md) | The list of validators within the workflow. | [optional] 
**WorkflowId** | [**WorkflowId**](WorkflowId.md) |  | 

## Methods

### NewWorkflowTransitionRules

`func NewWorkflowTransitionRules(workflowId WorkflowId, ) *WorkflowTransitionRules`

NewWorkflowTransitionRules instantiates a new WorkflowTransitionRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionRulesWithDefaults

`func NewWorkflowTransitionRulesWithDefaults() *WorkflowTransitionRules`

NewWorkflowTransitionRulesWithDefaults instantiates a new WorkflowTransitionRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *WorkflowTransitionRules) GetConditions() []AppWorkflowTransitionRule`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowTransitionRules) GetConditionsOk() (*[]AppWorkflowTransitionRule, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowTransitionRules) SetConditions(v []AppWorkflowTransitionRule)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *WorkflowTransitionRules) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetPostFunctions

`func (o *WorkflowTransitionRules) GetPostFunctions() []AppWorkflowTransitionRule`

GetPostFunctions returns the PostFunctions field if non-nil, zero value otherwise.

### GetPostFunctionsOk

`func (o *WorkflowTransitionRules) GetPostFunctionsOk() (*[]AppWorkflowTransitionRule, bool)`

GetPostFunctionsOk returns a tuple with the PostFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFunctions

`func (o *WorkflowTransitionRules) SetPostFunctions(v []AppWorkflowTransitionRule)`

SetPostFunctions sets PostFunctions field to given value.

### HasPostFunctions

`func (o *WorkflowTransitionRules) HasPostFunctions() bool`

HasPostFunctions returns a boolean if a field has been set.

### GetValidators

`func (o *WorkflowTransitionRules) GetValidators() []AppWorkflowTransitionRule`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *WorkflowTransitionRules) GetValidatorsOk() (*[]AppWorkflowTransitionRule, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *WorkflowTransitionRules) SetValidators(v []AppWorkflowTransitionRule)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *WorkflowTransitionRules) HasValidators() bool`

HasValidators returns a boolean if a field has been set.

### GetWorkflowId

`func (o *WorkflowTransitionRules) GetWorkflowId() WorkflowId`

GetWorkflowId returns the WorkflowId field if non-nil, zero value otherwise.

### GetWorkflowIdOk

`func (o *WorkflowTransitionRules) GetWorkflowIdOk() (*WorkflowId, bool)`

GetWorkflowIdOk returns a tuple with the WorkflowId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkflowId

`func (o *WorkflowTransitionRules) SetWorkflowId(v WorkflowId)`

SetWorkflowId sets WorkflowId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


