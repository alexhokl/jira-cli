# WorkflowTransitionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **interface{}** | EXPERIMENTAL. The configuration of the transition rule. | [optional] 
**Type** | **string** | The type of the transition rule. | 

## Methods

### NewWorkflowTransitionRule

`func NewWorkflowTransitionRule(type_ string, ) *WorkflowTransitionRule`

NewWorkflowTransitionRule instantiates a new WorkflowTransitionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTransitionRuleWithDefaults

`func NewWorkflowTransitionRuleWithDefaults() *WorkflowTransitionRule`

NewWorkflowTransitionRuleWithDefaults instantiates a new WorkflowTransitionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *WorkflowTransitionRule) GetConfiguration() interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *WorkflowTransitionRule) GetConfigurationOk() (*interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *WorkflowTransitionRule) SetConfiguration(v interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *WorkflowTransitionRule) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### SetConfigurationNil

`func (o *WorkflowTransitionRule) SetConfigurationNil(b bool)`

 SetConfigurationNil sets the value for Configuration to be an explicit nil

### UnsetConfiguration
`func (o *WorkflowTransitionRule) UnsetConfiguration()`

UnsetConfiguration ensures that no value is present for Configuration, not even an explicit nil
### GetType

`func (o *WorkflowTransitionRule) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowTransitionRule) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowTransitionRule) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


