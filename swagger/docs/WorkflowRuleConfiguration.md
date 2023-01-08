# WorkflowRuleConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **NullableString** | The ID of the rule. | [optional] 
**Parameters** | Pointer to **map[string]string** | The parameters related to the rule. | [optional] 
**RuleKey** | **string** | The rule key of the rule. | 

## Methods

### NewWorkflowRuleConfiguration

`func NewWorkflowRuleConfiguration(ruleKey string, ) *WorkflowRuleConfiguration`

NewWorkflowRuleConfiguration instantiates a new WorkflowRuleConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowRuleConfigurationWithDefaults

`func NewWorkflowRuleConfigurationWithDefaults() *WorkflowRuleConfiguration`

NewWorkflowRuleConfigurationWithDefaults instantiates a new WorkflowRuleConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowRuleConfiguration) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowRuleConfiguration) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowRuleConfiguration) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowRuleConfiguration) HasId() bool`

HasId returns a boolean if a field has been set.

### SetIdNil

`func (o *WorkflowRuleConfiguration) SetIdNil(b bool)`

 SetIdNil sets the value for Id to be an explicit nil

### UnsetId
`func (o *WorkflowRuleConfiguration) UnsetId()`

UnsetId ensures that no value is present for Id, not even an explicit nil
### GetParameters

`func (o *WorkflowRuleConfiguration) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WorkflowRuleConfiguration) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WorkflowRuleConfiguration) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *WorkflowRuleConfiguration) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetRuleKey

`func (o *WorkflowRuleConfiguration) GetRuleKey() string`

GetRuleKey returns the RuleKey field if non-nil, zero value otherwise.

### GetRuleKeyOk

`func (o *WorkflowRuleConfiguration) GetRuleKeyOk() (*string, bool)`

GetRuleKeyOk returns a tuple with the RuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleKey

`func (o *WorkflowRuleConfiguration) SetRuleKey(v string)`

SetRuleKey sets RuleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


