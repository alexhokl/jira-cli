# WorkflowCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | EXPERIMENTAL. The configuration of the transition rule. | [optional] 
**NodeType** | **string** |  | 
**Type** | **string** | The type of the transition rule. | 
**Conditions** | [**[]WorkflowCondition**](WorkflowCondition.md) | The list of workflow conditions. | 
**Operator** | **string** | The compound condition operator. | 

## Methods

### NewWorkflowCondition

`func NewWorkflowCondition(nodeType string, type_ string, conditions []WorkflowCondition, operator string, ) *WorkflowCondition`

NewWorkflowCondition instantiates a new WorkflowCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowConditionWithDefaults

`func NewWorkflowConditionWithDefaults() *WorkflowCondition`

NewWorkflowConditionWithDefaults instantiates a new WorkflowCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *WorkflowCondition) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *WorkflowCondition) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *WorkflowCondition) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *WorkflowCondition) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetNodeType

`func (o *WorkflowCondition) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *WorkflowCondition) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *WorkflowCondition) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetType

`func (o *WorkflowCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowCondition) SetType(v string)`

SetType sets Type field to given value.


### GetConditions

`func (o *WorkflowCondition) GetConditions() []WorkflowCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowCondition) GetConditionsOk() (*[]WorkflowCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowCondition) SetConditions(v []WorkflowCondition)`

SetConditions sets Conditions field to given value.


### GetOperator

`func (o *WorkflowCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *WorkflowCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *WorkflowCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


