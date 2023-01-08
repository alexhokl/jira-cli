# JiraExpressionEvalRequestBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to [**JiraExpressionEvalContextBean**](JiraExpressionEvalContextBean.md) | The context in which the Jira expression is evaluated. | [optional] 
**Expression** | **string** | The Jira expression to evaluate. | 

## Methods

### NewJiraExpressionEvalRequestBean

`func NewJiraExpressionEvalRequestBean(expression string, ) *JiraExpressionEvalRequestBean`

NewJiraExpressionEvalRequestBean instantiates a new JiraExpressionEvalRequestBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionEvalRequestBeanWithDefaults

`func NewJiraExpressionEvalRequestBeanWithDefaults() *JiraExpressionEvalRequestBean`

NewJiraExpressionEvalRequestBeanWithDefaults instantiates a new JiraExpressionEvalRequestBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *JiraExpressionEvalRequestBean) GetContext() JiraExpressionEvalContextBean`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *JiraExpressionEvalRequestBean) GetContextOk() (*JiraExpressionEvalContextBean, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *JiraExpressionEvalRequestBean) SetContext(v JiraExpressionEvalContextBean)`

SetContext sets Context field to given value.

### HasContext

`func (o *JiraExpressionEvalRequestBean) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExpression

`func (o *JiraExpressionEvalRequestBean) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *JiraExpressionEvalRequestBean) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *JiraExpressionEvalRequestBean) SetExpression(v string)`

SetExpression sets Expression field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


