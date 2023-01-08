# AppWorkflowTransitionRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | [**RuleConfiguration**](RuleConfiguration.md) |  | 
**Id** | **string** | The ID of the transition rule. | 
**Key** | **string** | The key of the rule, as defined in the Connect or the Forge app descriptor. | [readonly] 
**Transition** | Pointer to [**WorkflowTransition**](WorkflowTransition.md) |  | [optional] [readonly] 

## Methods

### NewAppWorkflowTransitionRule

`func NewAppWorkflowTransitionRule(configuration RuleConfiguration, id string, key string, ) *AppWorkflowTransitionRule`

NewAppWorkflowTransitionRule instantiates a new AppWorkflowTransitionRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppWorkflowTransitionRuleWithDefaults

`func NewAppWorkflowTransitionRuleWithDefaults() *AppWorkflowTransitionRule`

NewAppWorkflowTransitionRuleWithDefaults instantiates a new AppWorkflowTransitionRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *AppWorkflowTransitionRule) GetConfiguration() RuleConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *AppWorkflowTransitionRule) GetConfigurationOk() (*RuleConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *AppWorkflowTransitionRule) SetConfiguration(v RuleConfiguration)`

SetConfiguration sets Configuration field to given value.


### GetId

`func (o *AppWorkflowTransitionRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppWorkflowTransitionRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppWorkflowTransitionRule) SetId(v string)`

SetId sets Id field to given value.


### GetKey

`func (o *AppWorkflowTransitionRule) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *AppWorkflowTransitionRule) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *AppWorkflowTransitionRule) SetKey(v string)`

SetKey sets Key field to given value.


### GetTransition

`func (o *AppWorkflowTransitionRule) GetTransition() WorkflowTransition`

GetTransition returns the Transition field if non-nil, zero value otherwise.

### GetTransitionOk

`func (o *AppWorkflowTransitionRule) GetTransitionOk() (*WorkflowTransition, bool)`

GetTransitionOk returns a tuple with the Transition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransition

`func (o *AppWorkflowTransitionRule) SetTransition(v WorkflowTransition)`

SetTransition sets Transition field to given value.

### HasTransition

`func (o *AppWorkflowTransitionRule) HasTransition() bool`

HasTransition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


