# JiraExpressionEvaluationMetaDataBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complexity** | Pointer to [**JiraExpressionsComplexityBean**](JiraExpressionsComplexityBean.md) | Contains information about the expression complexity. For example, the number of steps it took to evaluate the expression. | [optional] 
**Issues** | Pointer to [**IssuesMetaBean**](IssuesMetaBean.md) | Contains information about the &#x60;issues&#x60; variable in the context. For example, is the issues were loaded with JQL, information about the page will be included here. | [optional] 

## Methods

### NewJiraExpressionEvaluationMetaDataBean

`func NewJiraExpressionEvaluationMetaDataBean() *JiraExpressionEvaluationMetaDataBean`

NewJiraExpressionEvaluationMetaDataBean instantiates a new JiraExpressionEvaluationMetaDataBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionEvaluationMetaDataBeanWithDefaults

`func NewJiraExpressionEvaluationMetaDataBeanWithDefaults() *JiraExpressionEvaluationMetaDataBean`

NewJiraExpressionEvaluationMetaDataBeanWithDefaults instantiates a new JiraExpressionEvaluationMetaDataBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplexity

`func (o *JiraExpressionEvaluationMetaDataBean) GetComplexity() JiraExpressionsComplexityBean`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *JiraExpressionEvaluationMetaDataBean) GetComplexityOk() (*JiraExpressionsComplexityBean, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *JiraExpressionEvaluationMetaDataBean) SetComplexity(v JiraExpressionsComplexityBean)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *JiraExpressionEvaluationMetaDataBean) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetIssues

`func (o *JiraExpressionEvaluationMetaDataBean) GetIssues() IssuesMetaBean`

GetIssues returns the Issues field if non-nil, zero value otherwise.

### GetIssuesOk

`func (o *JiraExpressionEvaluationMetaDataBean) GetIssuesOk() (*IssuesMetaBean, bool)`

GetIssuesOk returns a tuple with the Issues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssues

`func (o *JiraExpressionEvaluationMetaDataBean) SetIssues(v IssuesMetaBean)`

SetIssues sets Issues field to given value.

### HasIssues

`func (o *JiraExpressionEvaluationMetaDataBean) HasIssues() bool`

HasIssues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


