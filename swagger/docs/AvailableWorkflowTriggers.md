# AvailableWorkflowTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableTypes** | [**[]AvailableWorkflowTriggerTypes**](AvailableWorkflowTriggerTypes.md) | The list of available trigger types. | 
**RuleKey** | **string** | The rule key of the rule. | 

## Methods

### NewAvailableWorkflowTriggers

`func NewAvailableWorkflowTriggers(availableTypes []AvailableWorkflowTriggerTypes, ruleKey string, ) *AvailableWorkflowTriggers`

NewAvailableWorkflowTriggers instantiates a new AvailableWorkflowTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAvailableWorkflowTriggersWithDefaults

`func NewAvailableWorkflowTriggersWithDefaults() *AvailableWorkflowTriggers`

NewAvailableWorkflowTriggersWithDefaults instantiates a new AvailableWorkflowTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableTypes

`func (o *AvailableWorkflowTriggers) GetAvailableTypes() []AvailableWorkflowTriggerTypes`

GetAvailableTypes returns the AvailableTypes field if non-nil, zero value otherwise.

### GetAvailableTypesOk

`func (o *AvailableWorkflowTriggers) GetAvailableTypesOk() (*[]AvailableWorkflowTriggerTypes, bool)`

GetAvailableTypesOk returns a tuple with the AvailableTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTypes

`func (o *AvailableWorkflowTriggers) SetAvailableTypes(v []AvailableWorkflowTriggerTypes)`

SetAvailableTypes sets AvailableTypes field to given value.


### GetRuleKey

`func (o *AvailableWorkflowTriggers) GetRuleKey() string`

GetRuleKey returns the RuleKey field if non-nil, zero value otherwise.

### GetRuleKeyOk

`func (o *AvailableWorkflowTriggers) GetRuleKeyOk() (*string, bool)`

GetRuleKeyOk returns a tuple with the RuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleKey

`func (o *AvailableWorkflowTriggers) SetRuleKey(v string)`

SetRuleKey sets RuleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


