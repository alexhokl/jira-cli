# JiraExpressionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**JiraExpressionEvaluationMetaDataBean**](JiraExpressionEvaluationMetaDataBean.md) | Contains various characteristics of the performed expression evaluation. | [optional] 
**Value** | **interface{}** | The value of the evaluated expression. It may be a primitive JSON value or a Jira REST API object. (Some expressions do not produce any meaningful results—for example, an expression that returns a lambda function—if that&#39;s the case a simple string representation is returned. These string representations should not be relied upon and may change without notice.) | 

## Methods

### NewJiraExpressionResult

`func NewJiraExpressionResult(value interface{}, ) *JiraExpressionResult`

NewJiraExpressionResult instantiates a new JiraExpressionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionResultWithDefaults

`func NewJiraExpressionResultWithDefaults() *JiraExpressionResult`

NewJiraExpressionResultWithDefaults instantiates a new JiraExpressionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *JiraExpressionResult) GetMeta() JiraExpressionEvaluationMetaDataBean`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JiraExpressionResult) GetMetaOk() (*JiraExpressionEvaluationMetaDataBean, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JiraExpressionResult) SetMeta(v JiraExpressionEvaluationMetaDataBean)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JiraExpressionResult) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValue

`func (o *JiraExpressionResult) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JiraExpressionResult) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JiraExpressionResult) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JiraExpressionResult) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JiraExpressionResult) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


