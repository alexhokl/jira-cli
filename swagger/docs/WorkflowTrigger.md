# WorkflowTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the trigger. | [optional] 
**Parameters** | **map[string]string** | The parameters of the trigger. | 
**RuleKey** | **string** | The rule key of the trigger. | 

## Methods

### NewWorkflowTrigger

`func NewWorkflowTrigger(parameters map[string]string, ruleKey string, ) *WorkflowTrigger`

NewWorkflowTrigger instantiates a new WorkflowTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowTriggerWithDefaults

`func NewWorkflowTriggerWithDefaults() *WorkflowTrigger`

NewWorkflowTriggerWithDefaults instantiates a new WorkflowTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *WorkflowTrigger) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *WorkflowTrigger) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *WorkflowTrigger) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *WorkflowTrigger) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParameters

`func (o *WorkflowTrigger) GetParameters() map[string]string`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *WorkflowTrigger) GetParametersOk() (*map[string]string, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *WorkflowTrigger) SetParameters(v map[string]string)`

SetParameters sets Parameters field to given value.


### GetRuleKey

`func (o *WorkflowTrigger) GetRuleKey() string`

GetRuleKey returns the RuleKey field if non-nil, zero value otherwise.

### GetRuleKeyOk

`func (o *WorkflowTrigger) GetRuleKeyOk() (*string, bool)`

GetRuleKeyOk returns a tuple with the RuleKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleKey

`func (o *WorkflowTrigger) SetRuleKey(v string)`

SetRuleKey sets RuleKey field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


