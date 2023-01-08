# WorkflowRules

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionsTree** | Pointer to [**WorkflowCondition**](WorkflowCondition.md) |  | [optional] 
**PostFunctions** | Pointer to [**[]WorkflowTransitionRule**](WorkflowTransitionRule.md) | The workflow post functions. | [optional] 
**Validators** | Pointer to [**[]WorkflowTransitionRule**](WorkflowTransitionRule.md) | The workflow validators. | [optional] 

## Methods

### NewWorkflowRules

`func NewWorkflowRules() *WorkflowRules`

NewWorkflowRules instantiates a new WorkflowRules object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRulesWithDefaults

`func NewWorkflowRulesWithDefaults() *WorkflowRules`

NewWorkflowRulesWithDefaults instantiates a new WorkflowRules object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionsTree

`func (o *WorkflowRules) GetConditionsTree() WorkflowCondition`

GetConditionsTree returns the ConditionsTree field if non-nil, zero value otherwise.

### GetConditionsTreeOk

`func (o *WorkflowRules) GetConditionsTreeOk() (*WorkflowCondition, bool)`

GetConditionsTreeOk returns a tuple with the ConditionsTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionsTree

`func (o *WorkflowRules) SetConditionsTree(v WorkflowCondition)`

SetConditionsTree sets ConditionsTree field to given value.

### HasConditionsTree

`func (o *WorkflowRules) HasConditionsTree() bool`

HasConditionsTree returns a boolean if a field has been set.

### GetPostFunctions

`func (o *WorkflowRules) GetPostFunctions() []WorkflowTransitionRule`

GetPostFunctions returns the PostFunctions field if non-nil, zero value otherwise.

### GetPostFunctionsOk

`func (o *WorkflowRules) GetPostFunctionsOk() (*[]WorkflowTransitionRule, bool)`

GetPostFunctionsOk returns a tuple with the PostFunctions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostFunctions

`func (o *WorkflowRules) SetPostFunctions(v []WorkflowTransitionRule)`

SetPostFunctions sets PostFunctions field to given value.

### HasPostFunctions

`func (o *WorkflowRules) HasPostFunctions() bool`

HasPostFunctions returns a boolean if a field has been set.

### GetValidators

`func (o *WorkflowRules) GetValidators() []WorkflowTransitionRule`

GetValidators returns the Validators field if non-nil, zero value otherwise.

### GetValidatorsOk

`func (o *WorkflowRules) GetValidatorsOk() (*[]WorkflowTransitionRule, bool)`

GetValidatorsOk returns a tuple with the Validators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValidators

`func (o *WorkflowRules) SetValidators(v []WorkflowTransitionRule)`

SetValidators sets Validators field to given value.

### HasValidators

`func (o *WorkflowRules) HasValidators() bool`

HasValidators returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


