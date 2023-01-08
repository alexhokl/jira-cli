# CreateWorkflowCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | Pointer to [**[]CreateWorkflowCondition**](CreateWorkflowCondition.md) | The list of workflow conditions. | [optional] 
**Configuration** | Pointer to **map[string]interface{}** | EXPERIMENTAL. The configuration of the transition rule. | [optional] 
**Operator** | Pointer to **string** | The compound condition operator. | [optional] 
**Type** | Pointer to **string** | The type of the transition rule. | [optional] 

## Methods

### NewCreateWorkflowCondition

`func NewCreateWorkflowCondition() *CreateWorkflowCondition`

NewCreateWorkflowCondition instantiates a new CreateWorkflowCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateWorkflowConditionWithDefaults

`func NewCreateWorkflowConditionWithDefaults() *CreateWorkflowCondition`

NewCreateWorkflowConditionWithDefaults instantiates a new CreateWorkflowCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *CreateWorkflowCondition) GetConditions() []CreateWorkflowCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *CreateWorkflowCondition) GetConditionsOk() (*[]CreateWorkflowCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *CreateWorkflowCondition) SetConditions(v []CreateWorkflowCondition)`

SetConditions sets Conditions field to given value.

### HasConditions

`func (o *CreateWorkflowCondition) HasConditions() bool`

HasConditions returns a boolean if a field has been set.

### GetConfiguration

`func (o *CreateWorkflowCondition) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *CreateWorkflowCondition) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *CreateWorkflowCondition) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *CreateWorkflowCondition) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetOperator

`func (o *CreateWorkflowCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *CreateWorkflowCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *CreateWorkflowCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *CreateWorkflowCondition) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetType

`func (o *CreateWorkflowCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateWorkflowCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateWorkflowCondition) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateWorkflowCondition) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


