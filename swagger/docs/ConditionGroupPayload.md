# ConditionGroupPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConditionGroup** | Pointer to [**[]ConditionGroupPayload**](ConditionGroupPayload.md) | The nested conditions of the condition group. | [optional] 
**Conditions** | Pointer to [**[]RulePayload**](RulePayload.md) | The rules for this condition. | [optional] 
**Operation** | Pointer to **string** | Determines how the conditions in the group are evaluated. Accepts either &#x60;ANY&#x60; or &#x60;ALL&#x60;. If &#x60;ANY&#x60; is used, at least one condition in the group must be true for the group to evaluate to true. If &#x60;ALL&#x60; is used, all conditions in the group must be true for the group to evaluate to true. | [optional] 

## Methods

### NewConditionGroupPayload

`func NewConditionGroupPayload() *ConditionGroupPayload`

NewConditionGroupPayload instantiates a new ConditionGroupPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConditionGroupPayloadWithDefaults

`func NewConditionGroupPayloadWithDefaults() *ConditionGroupPayload`

NewConditionGroupPayloadWithDefaults instantiates a new ConditionGroupPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditionGroup

`func (o *ConditionGroupPayload) GetConditionGroup() []ConditionGroupPayload`

GetConditionGroup returns the ConditionGroup field if non-nil, zero value otherwise.

### GetConditionGroupOk

`func (o *ConditionGroupPayload) GetConditionGroupOk() (*[]ConditionGroupPayload, bool)`

GetConditionGroupOk returns a tuple with the ConditionGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionGroup

`func (o *ConditionGroupPayload) SetConditionGroup(v []ConditionGroupPayload)`

SetConditionGroup sets ConditionGroup field to given value.

### HasConditionGroup

`func (o *ConditionGroupPayload) HasConditionGroup() bool`

HasConditionGroup returns a boolean if a field has been set.

### GetConditions

`func (o *ConditionGroupPayload) GetConditions() []RulePayload`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *ConditionGroupPayload) GetConditionsOk() (*[]RulePayload, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *ConditionGroupPayload) SetConditions(v []RulePayload)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *ConditionGroupPayload) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetOperation

`func (o *ConditionGroupPayload) GetOperation() string`

GetOperation returns the Operation field if non-nil, zero value otherwise.

### GetOperationOk

`func (o *ConditionGroupPayload) GetOperationOk() (*string, bool)`

GetOperationOk returns a tuple with the Operation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperation

`func (o *ConditionGroupPayload) SetOperation(v string)`

SetOperation sets Operation field to given value.

### HasOperation

`func (o *ConditionGroupPayload) HasOperation() bool`

HasOperation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


