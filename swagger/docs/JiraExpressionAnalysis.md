# JiraExpressionAnalysis

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complexity** | Pointer to [**JiraExpressionComplexity**](JiraExpressionComplexity.md) |  | [optional] 
**Errors** | Pointer to [**[]JiraExpressionValidationError**](JiraExpressionValidationError.md) | A list of validation errors. Not included if the expression is valid. | [optional] 
**Expression** | **string** | The analysed expression. | 
**Type** | Pointer to **string** | EXPERIMENTAL. The inferred type of the expression. | [optional] 
**Valid** | **bool** | Whether the expression is valid and the interpreter will evaluate it. Note that the expression may fail at runtime (for example, if it executes too many expensive operations). | 

## Methods

### NewJiraExpressionAnalysis

`func NewJiraExpressionAnalysis(expression string, valid bool, ) *JiraExpressionAnalysis`

NewJiraExpressionAnalysis instantiates a new JiraExpressionAnalysis object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraExpressionAnalysisWithDefaults

`func NewJiraExpressionAnalysisWithDefaults() *JiraExpressionAnalysis`

NewJiraExpressionAnalysisWithDefaults instantiates a new JiraExpressionAnalysis object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetComplexity

`func (o *JiraExpressionAnalysis) GetComplexity() JiraExpressionComplexity`

GetComplexity returns the Complexity field if non-nil, zero value otherwise.

### GetComplexityOk

`func (o *JiraExpressionAnalysis) GetComplexityOk() (*JiraExpressionComplexity, bool)`

GetComplexityOk returns a tuple with the Complexity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComplexity

`func (o *JiraExpressionAnalysis) SetComplexity(v JiraExpressionComplexity)`

SetComplexity sets Complexity field to given value.

### HasComplexity

`func (o *JiraExpressionAnalysis) HasComplexity() bool`

HasComplexity returns a boolean if a field has been set.

### GetErrors

`func (o *JiraExpressionAnalysis) GetErrors() []JiraExpressionValidationError`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *JiraExpressionAnalysis) GetErrorsOk() (*[]JiraExpressionValidationError, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *JiraExpressionAnalysis) SetErrors(v []JiraExpressionValidationError)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *JiraExpressionAnalysis) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetExpression

`func (o *JiraExpressionAnalysis) GetExpression() string`

GetExpression returns the Expression field if non-nil, zero value otherwise.

### GetExpressionOk

`func (o *JiraExpressionAnalysis) GetExpressionOk() (*string, bool)`

GetExpressionOk returns a tuple with the Expression field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpression

`func (o *JiraExpressionAnalysis) SetExpression(v string)`

SetExpression sets Expression field to given value.


### GetType

`func (o *JiraExpressionAnalysis) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JiraExpressionAnalysis) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JiraExpressionAnalysis) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JiraExpressionAnalysis) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValid

`func (o *JiraExpressionAnalysis) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *JiraExpressionAnalysis) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *JiraExpressionAnalysis) SetValid(v bool)`

SetValid sets Valid field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


