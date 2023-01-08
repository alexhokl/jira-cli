# JiraExpressionAnalysis

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Complexity** | [***JiraExpressionComplexity**](JiraExpressionComplexity.md) |  | [optional] [default to null]
**Errors** | [**[]JiraExpressionValidationError**](JiraExpressionValidationError.md) | A list of validation errors. Not included if the expression is valid. | [optional] [default to null]
**Expression** | **string** | The analysed expression. | [default to null]
**Type_** | **string** | EXPERIMENTAL. The inferred type of the expression. | [optional] [default to null]
**Valid** | **bool** | Whether the expression is valid and the interpreter will evaluate it. Note that the expression may fail at runtime (for example, if it executes too many expensive operations). | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

