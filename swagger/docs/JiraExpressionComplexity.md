# JiraExpressionComplexity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpensiveOperations** | **string** | Information that can be used to determine how many [expensive operations](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#expensive-operations) the evaluation of the expression will perform. This information may be a formula or number. For example:   *  &#x60;issues.map(i &#x3D;&gt; i.comments)&#x60; performs as many expensive operations as there are issues on the issues list. So this parameter returns &#x60;N&#x60;, where &#x60;N&#x60; is the size of issue list.  *  &#x60;new Issue(10010).comments&#x60; gets comments for one issue, so its complexity is &#x60;2&#x60; (&#x60;1&#x60; to retrieve issue 10010 from the database plus &#x60;1&#x60; to get its comments). | 
**Variables** | Pointer to **map[string]string** | Variables used in the formula, mapped to the parts of the expression they refer to. | [optional] 

## Methods

### NewJiraExpressionComplexity

`func NewJiraExpressionComplexity(expensiveOperations string, ) *JiraExpressionComplexity`

NewJiraExpressionComplexity instantiates a new JiraExpressionComplexity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionComplexityWithDefaults

`func NewJiraExpressionComplexityWithDefaults() *JiraExpressionComplexity`

NewJiraExpressionComplexityWithDefaults instantiates a new JiraExpressionComplexity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpensiveOperations

`func (o *JiraExpressionComplexity) GetExpensiveOperations() string`

GetExpensiveOperations returns the ExpensiveOperations field if non-nil, zero value otherwise.

### GetExpensiveOperationsOk

`func (o *JiraExpressionComplexity) GetExpensiveOperationsOk() (*string, bool)`

GetExpensiveOperationsOk returns a tuple with the ExpensiveOperations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpensiveOperations

`func (o *JiraExpressionComplexity) SetExpensiveOperations(v string)`

SetExpensiveOperations sets ExpensiveOperations field to given value.


### GetVariables

`func (o *JiraExpressionComplexity) GetVariables() map[string]string`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *JiraExpressionComplexity) GetVariablesOk() (*map[string]string, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *JiraExpressionComplexity) SetVariables(v map[string]string)`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *JiraExpressionComplexity) HasVariables() bool`

HasVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


