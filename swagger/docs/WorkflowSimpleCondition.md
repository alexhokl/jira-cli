# WorkflowSimpleCondition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configuration** | Pointer to **map[string]interface{}** | EXPERIMENTAL. The configuration of the transition rule. | [optional] 
**NodeType** | **string** |  | 
**Type** | **string** | The type of the transition rule. | 

## Methods

### NewWorkflowSimpleCondition

`func NewWorkflowSimpleCondition(nodeType string, type_ string, ) *WorkflowSimpleCondition`

NewWorkflowSimpleCondition instantiates a new WorkflowSimpleCondition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowSimpleConditionWithDefaults

`func NewWorkflowSimpleConditionWithDefaults() *WorkflowSimpleCondition`

NewWorkflowSimpleConditionWithDefaults instantiates a new WorkflowSimpleCondition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfiguration

`func (o *WorkflowSimpleCondition) GetConfiguration() map[string]interface{}`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *WorkflowSimpleCondition) GetConfigurationOk() (*map[string]interface{}, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *WorkflowSimpleCondition) SetConfiguration(v map[string]interface{})`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *WorkflowSimpleCondition) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetNodeType

`func (o *WorkflowSimpleCondition) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *WorkflowSimpleCondition) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *WorkflowSimpleCondition) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.


### GetType

`func (o *WorkflowSimpleCondition) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowSimpleCondition) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowSimpleCondition) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


