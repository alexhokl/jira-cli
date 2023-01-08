# WorkflowCompoundCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conditions** | [**[]WorkflowCondition**](WorkflowCondition.md) | The list of workflow conditions. | 
**NodeType** | **string** |  | 
**Operator** | **string** | The compound condition operator. | 

## Methods

### NewWorkflowCompoundCondition

`func NewWorkflowCompoundCondition(conditions []WorkflowCondition, nodeType string, operator string, ) *WorkflowCompoundCondition`

NewWorkflowCompoundCondition instantiates a new WorkflowCompoundCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowCompoundConditionWithDefaults

`func NewWorkflowCompoundConditionWithDefaults() *WorkflowCompoundCondition`

NewWorkflowCompoundConditionWithDefaults instantiates a new WorkflowCompoundCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConditions

`func (o *WorkflowCompoundCondition) GetConditions() []WorkflowCondition`

GetConditions returns the Conditions field if non-nil, zero value otherwise.

### GetConditionsOk

`func (o *WorkflowCompoundCondition) GetConditionsOk() (*[]WorkflowCondition, bool)`

GetConditionsOk returns a tuple with the Conditions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditions

`func (o *WorkflowCompoundCondition) SetConditions(v []WorkflowCondition)`

SetConditions sets Conditions field to given value.


### GetNodeType

`func (o *WorkflowCompoundCondition) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *WorkflowCompoundCondition) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *WorkflowCompoundCondition) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetOperator

`func (o *WorkflowCompoundCondition) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *WorkflowCompoundCondition) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *WorkflowCompoundCondition) SetOperator(v string)`

SetOperator sets Operator field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


