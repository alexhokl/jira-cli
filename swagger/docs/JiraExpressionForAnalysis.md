# JiraExpressionForAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextVariables** | Pointer to **map[string]string** | Context variables and their types. The type checker assumes that [common context variables](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#context-variables), such as &#x60;issue&#x60; or &#x60;project&#x60;, are available in context and sets their type. Use this property to override the default types or provide details of new variables. | [optional] 
**Expressions** | **[]string** | The list of Jira expressions to analyse. | 

## Methods

### NewJiraExpressionForAnalysis

`func NewJiraExpressionForAnalysis(expressions []string, ) *JiraExpressionForAnalysis`

NewJiraExpressionForAnalysis instantiates a new JiraExpressionForAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionForAnalysisWithDefaults

`func NewJiraExpressionForAnalysisWithDefaults() *JiraExpressionForAnalysis`

NewJiraExpressionForAnalysisWithDefaults instantiates a new JiraExpressionForAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextVariables

`func (o *JiraExpressionForAnalysis) GetContextVariables() map[string]string`

GetContextVariables returns the ContextVariables field if non-nil, zero value otherwise.

### GetContextVariablesOk

`func (o *JiraExpressionForAnalysis) GetContextVariablesOk() (*map[string]string, bool)`

GetContextVariablesOk returns a tuple with the ContextVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextVariables

`func (o *JiraExpressionForAnalysis) SetContextVariables(v map[string]string)`

SetContextVariables sets ContextVariables field to given value.

### HasContextVariables

`func (o *JiraExpressionForAnalysis) HasContextVariables() bool`

HasContextVariables returns a boolean if a field has been set.

### GetExpressions

`func (o *JiraExpressionForAnalysis) GetExpressions() []string`

GetExpressions returns the Expressions field if non-nil, zero value otherwise.

### GetExpressionsOk

`func (o *JiraExpressionForAnalysis) GetExpressionsOk() (*[]string, bool)`

GetExpressionsOk returns a tuple with the Expressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpressions

`func (o *JiraExpressionForAnalysis) SetExpressions(v []string)`

SetExpressions sets Expressions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


