# ConditionGroupConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroups** | Pointer to [**[]ConditionGroupConfiguration**](ConditionGroupConfiguration.md) | The nested conditions of the condition group. | [optional] 
**Conditions** | Pointer to [**[]WorkflowRuleConfiguration**](WorkflowRuleConfiguration.md) | The rules for this condition. | [optional] 
**Operation** | Pointer to **string** | Determines how the conditions in the group are evaluated. Accepts either &#x60;ANY&#x60; or &#x60;ALL&#x60;. If &#x60;ANY&#x60; is used, at least one condition in the group must be true for the group to evaluate to true. If &#x60;ALL&#x60; is used, all conditions in the group must be true for the group to evaluate to true. | [optional] 

## Methods

### NewConditionGroupConfiguration

`func NewConditionGroupConfiguration() *ConditionGroupConfiguration`

NewConditionGroupConfiguration instantiates a new ConditionGroupConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionGroupConfigurationWithDefaults

`func NewConditionGroupConfigurationWithDefaults() *ConditionGroupConfiguration`

NewConditionGroupConfigurationWithDefaults instantiates a new ConditionGroupConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroups

`func (o *ConditionGroupConfiguration) GetConditionGroups() []ConditionGroupConfiguration`

GetConditionGroups returns the ConditionGroups field if non-nil, zero value otherwise.

### GetConditionGroupsOk

`func (o *ConditionGroupConfiguration) GetConditionGroupsOk() (*[]ConditionGroupConfiguration, bool)`

GetConditionGroupsOk returns a tuple with the ConditionGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroups

`func (o *ConditionGroupConfiguration) SetConditionGroups(v []ConditionGroupConfiguration)`

SetConditionGroups sets ConditionGroups field to given value.

### HasConditionGroups

`func (o *ConditionGroupConfiguration) HasConditionGroups() bool`

HasConditionGroups returns a boolean if a field has been set.

### GetConditions

`func (o *ConditionGroupConfiguration) GetConditions() []WorkflowRuleConfiguration`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ConditionGroupConfiguration) GetConditionsOk() (*[]WorkflowRuleConfiguration, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ConditionGroupConfiguration) SetConditions(v []WorkflowRuleConfiguration)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ConditionGroupConfiguration) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOperation

`func (o *ConditionGroupConfiguration) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ConditionGroupConfiguration) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ConditionGroupConfiguration) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ConditionGroupConfiguration) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


