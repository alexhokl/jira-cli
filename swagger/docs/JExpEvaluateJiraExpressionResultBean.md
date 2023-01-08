# JExpEvaluateJiraExpressionResultBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Meta** | Pointer to [**JExpEvaluateMetaDataBean**](JExpEvaluateMetaDataBean.md) | Contains various characteristics of the performed expression evaluation. | [optional] 
**Value** | **interface{}** | The value of the evaluated expression. It may be a primitive JSON value or a Jira REST API object. (Some expressions do not produce any meaningful results—for example, an expression that returns a lambda function—if that&#39;s the case a simple string representation is returned. These string representations should not be relied upon and may change without notice.) | 

## Methods

### NewJExpEvaluateJiraExpressionResultBean

`func NewJExpEvaluateJiraExpressionResultBean(value interface{}, ) *JExpEvaluateJiraExpressionResultBean`

NewJExpEvaluateJiraExpressionResultBean instantiates a new JExpEvaluateJiraExpressionResultBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJExpEvaluateJiraExpressionResultBeanWithDefaults

`func NewJExpEvaluateJiraExpressionResultBeanWithDefaults() *JExpEvaluateJiraExpressionResultBean`

NewJExpEvaluateJiraExpressionResultBeanWithDefaults instantiates a new JExpEvaluateJiraExpressionResultBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMeta

`func (o *JExpEvaluateJiraExpressionResultBean) GetMeta() JExpEvaluateMetaDataBean`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *JExpEvaluateJiraExpressionResultBean) GetMetaOk() (*JExpEvaluateMetaDataBean, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *JExpEvaluateJiraExpressionResultBean) SetMeta(v JExpEvaluateMetaDataBean)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *JExpEvaluateJiraExpressionResultBean) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetValue

`func (o *JExpEvaluateJiraExpressionResultBean) GetValue() interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JExpEvaluateJiraExpressionResultBean) GetValueOk() (*interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JExpEvaluateJiraExpressionResultBean) SetValue(v interface{})`

SetValue sets Value field to given value.


### SetValueNil

`func (o *JExpEvaluateJiraExpressionResultBean) SetValueNil(b bool)`

 SetValueNil sets the value for Value to be an explicit nil

### UnsetValue
`func (o *JExpEvaluateJiraExpressionResultBean) UnsetValue()`

UnsetValue ensures that no value is present for Value, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


