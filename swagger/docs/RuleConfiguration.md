# RuleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Disabled** | Pointer to **bool** | Whether the rule is disabled. | [optional] [default to false]
**Tag** | Pointer to **string** | A tag used to filter rules in [Get workflow transition rule configurations](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-transition-rules/#api-rest-api-3-workflow-rule-config-get). | [optional] 
**Value** | **string** | Configuration of the rule, as it is stored by the Connect or the Forge app on the rule configuration page. | 

## Methods

### NewRuleConfiguration

`func NewRuleConfiguration(value string, ) *RuleConfiguration`

NewRuleConfiguration instantiates a new RuleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRuleConfigurationWithDefaults

`func NewRuleConfigurationWithDefaults() *RuleConfiguration`

NewRuleConfigurationWithDefaults instantiates a new RuleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDisabled

`func (o *RuleConfiguration) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *RuleConfiguration) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *RuleConfiguration) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *RuleConfiguration) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetTag

`func (o *RuleConfiguration) GetTag() string`

GetTag returns the Tag field if non-nil, zero value otherwise.

### GetTagOk

`func (o *RuleConfiguration) GetTagOk() (*string, bool)`

GetTagOk returns a tuple with the Tag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTag

`func (o *RuleConfiguration) SetTag(v string)`

SetTag sets Tag field to given value.

### HasTag

`func (o *RuleConfiguration) HasTag() bool`

HasTag returns a boolean if a field has been set.

### GetValue

`func (o *RuleConfiguration) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *RuleConfiguration) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *RuleConfiguration) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


