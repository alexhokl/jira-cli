# ConditionGroupUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | Pointer to [**[]ConditionGroupUpdate**](ConditionGroupUpdate.md) | The nested conditions of the condition group. | [optional] 
**Conditions** | Pointer to [**[]WorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) | The rules for this condition. | [optional] 
**Operation** | **string** | Determines how the conditions in the group are evaluated. Accepts either &#x60;ANY&#x60; or &#x60;ALL&#x60;. If &#x60;ANY&#x60; is used, at least one condition in the group must be true for the group to evaluate to true. If &#x60;ALL&#x60; is used, all conditions in the group must be true for the group to evaluate to true. | 

## Methods

### NewConditionGroupUpdate

`func NewConditionGroupUpdate(operation string, ) *ConditionGroupUpdate`

NewConditionGroupUpdate instantiates a new ConditionGroupUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionGroupUpdateWithDefaults

`func NewConditionGroupUpdateWithDefaults() *ConditionGroupUpdate`

NewConditionGroupUpdateWithDefaults instantiates a new ConditionGroupUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *ConditionGroupUpdate) GetConditionGroups() []ConditionGroupUpdate`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *ConditionGroupUpdate) GetConditionGroupsOk() (*[]ConditionGroupUpdate, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *ConditionGroupUpdate) SetConditionGroups(v []ConditionGroupUpdate)`

SetConditionGroups sets ConditionGroups field to given value.

### HasConditionGroups

`func (o *ConditionGroupUpdate) HasConditionGroups() bool`

HasConditionGroups returns a boolean if a field has been set.

### GetConditions

`func (o *ConditionGroupUpdate) GetConditions() []WorkflowRuleConfiguration`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ConditionGroupUpdate) GetConditionsOk() (*[]WorkflowRuleConfiguration, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ConditionGroupUpdate) SetConditions(v []WorkflowRuleConfiguration)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ConditionGroupUpdate) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOperation

`func (o *ConditionGroupUpdate) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ConditionGroupUpdate) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ConditionGroupUpdate) SetOperation(v string)`

SetOperation sets Operation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


