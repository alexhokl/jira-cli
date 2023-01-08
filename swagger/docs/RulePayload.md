# RulePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to **map[string]string** | The parameters of the rule | [optional] 
**RuleKey** | Pointer to **string** | The key of the rule. See https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflows/\\#api-rest-api-3-workflows-capabilities-get | [optional] 

## Methods

### NewRulePayload

`func NewRulePayload() *RulePayload`

NewRulePayload instantiates a new RulePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRulePayloadWithDefaults

`func NewRulePayloadWithDefaults() *RulePayload`

NewRulePayloadWithDefaults instantiates a new RulePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *RulePayload) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *RulePayload) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *RulePayload) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *RulePayload) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRuleKey

`func (o *RulePayload) GetRuleKey() string`

GetRuleKey returns the RuleKey field if non-nil, zero value otherwise.

### GetRuleKeyOk

`func (o *RulePayload) GetRuleKeyOk() (*string, bool)`

GetRuleKeyOk returns a tuple with the RuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleKey

`func (o *RulePayload) SetRuleKey(v string)`

SetRuleKey sets RuleKey field to given value.

### HasRuleKey

`func (o *RulePayload) HasRuleKey() bool`

HasRuleKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


