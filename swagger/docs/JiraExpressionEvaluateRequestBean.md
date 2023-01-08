# JiraExpressionEvaluateRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JiraExpressionEvaluateContextBean**](JiraExpressionEvaluateContextBean.md) | The context in which the Jira expression is evaluated. | [optional] 
**Expression** | **string** | The Jira expression to evaluate. | 

## Methods

### NewJiraExpressionEvaluateRequestBean

`func NewJiraExpressionEvaluateRequestBean(expression string, ) *JiraExpressionEvaluateRequestBean`

NewJiraExpressionEvaluateRequestBean instantiates a new JiraExpressionEvaluateRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionEvaluateRequestBeanWithDefaults

`func NewJiraExpressionEvaluateRequestBeanWithDefaults() *JiraExpressionEvaluateRequestBean`

NewJiraExpressionEvaluateRequestBeanWithDefaults instantiates a new JiraExpressionEvaluateRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *JiraExpressionEvaluateRequestBean) GetContext() JiraExpressionEvaluateContextBean`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *JiraExpressionEvaluateRequestBean) GetContextOk() (*JiraExpressionEvaluateContextBean, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *JiraExpressionEvaluateRequestBean) SetContext(v JiraExpressionEvaluateContextBean)`

SetContext sets Context field to given value.

### HasContext

`func (o *JiraExpressionEvaluateRequestBean) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExpression

`func (o *JiraExpressionEvaluateRequestBean) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *JiraExpressionEvaluateRequestBean) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *JiraExpressionEvaluateRequestBean) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


