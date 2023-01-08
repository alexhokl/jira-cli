# \JiraExpressionsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AnalyseExpression**](JiraExpressionsAPI.md#AnalyseExpression) | **Post** /rest/api/3/expression/analyse | Analyse Jira expression
[**EvaluateJSISJiraExpression**](JiraExpressionsAPI.md#EvaluateJSISJiraExpression) | **Post** /rest/api/3/expression/evaluate | Evaluate Jira expression using enhanced search API
[**EvaluateJiraExpression**](JiraExpressionsAPI.md#EvaluateJiraExpression) | **Post** /rest/api/3/expression/eval | Currently being removed. Evaluate Jira expression



## AnalyseExpression

> JiraExpressionsAnalysis AnalyseExpression(ctx).JiraExpressionForAnalysis(jiraExpressionForAnalysis).Check(check).Execute()

Analyse Jira expression



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	jiraExpressionForAnalysis := *openapiclient.NewJiraExpressionForAnalysis([]string{"issues.map(issue => issue.properties['property_key'])"}) // JiraExpressionForAnalysis | The Jira expressions to analyse.
	check := "check_example" // string | The check to perform:   *  `syntax` Each expression's syntax is checked to ensure the expression can be parsed. Also, syntactic limits are validated. For example, the expression's length.  *  `type` EXPERIMENTAL. Each expression is type checked and the final type of the expression inferred. Any type errors that would result in the expression failure at runtime are reported. For example, accessing properties that don't exist or passing the wrong number of arguments to functions. Also performs the syntax check.  *  `complexity` EXPERIMENTAL. Determines the formulae for how many [expensive operations](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#expensive-operations) each expression may execute. (optional) (default to "syntax")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraExpressionsAPI.AnalyseExpression(context.Background()).JiraExpressionForAnalysis(jiraExpressionForAnalysis).Check(check).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraExpressionsAPI.AnalyseExpression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AnalyseExpression`: JiraExpressionsAnalysis
	fmt.Fprintf(os.Stdout, "Response from `JiraExpressionsAPI.AnalyseExpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAnalyseExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jiraExpressionForAnalysis** | [**JiraExpressionForAnalysis**](JiraExpressionForAnalysis.md) | The Jira expressions to analyse. | 
 **check** | **string** | The check to perform:   *  &#x60;syntax&#x60; Each expression&#39;s syntax is checked to ensure the expression can be parsed. Also, syntactic limits are validated. For example, the expression&#39;s length.  *  &#x60;type&#x60; EXPERIMENTAL. Each expression is type checked and the final type of the expression inferred. Any type errors that would result in the expression failure at runtime are reported. For example, accessing properties that don&#39;t exist or passing the wrong number of arguments to functions. Also performs the syntax check.  *  &#x60;complexity&#x60; EXPERIMENTAL. Determines the formulae for how many [expensive operations](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#expensive-operations) each expression may execute. | [default to &quot;syntax&quot;]

### Return type

[**JiraExpressionsAnalysis**](JiraExpressionsAnalysis.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateJSISJiraExpression

> JExpEvaluateJiraExpressionResultBean EvaluateJSISJiraExpression(ctx).JiraExpressionEvaluateRequestBean(jiraExpressionEvaluateRequestBean).Expand(expand).Execute()

Evaluate Jira expression using enhanced search API



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	jiraExpressionEvaluateRequestBean := *openapiclient.NewJiraExpressionEvaluateRequestBean("{ key: issue.key, type: issue.issueType.name, links: issue.links.map(link => link.linkedIssue.id) }") // JiraExpressionEvaluateRequestBean | The Jira expression and the evaluation context.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts `meta.complexity` that returns information about the expression complexity. For example, the number of expensive operations used by the expression and how close the expression is to reaching the [complexity limit](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#restrictions). Useful when designing and debugging your expressions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraExpressionsAPI.EvaluateJSISJiraExpression(context.Background()).JiraExpressionEvaluateRequestBean(jiraExpressionEvaluateRequestBean).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraExpressionsAPI.EvaluateJSISJiraExpression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateJSISJiraExpression`: JExpEvaluateJiraExpressionResultBean
	fmt.Fprintf(os.Stdout, "Response from `JiraExpressionsAPI.EvaluateJSISJiraExpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateJSISJiraExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jiraExpressionEvaluateRequestBean** | [**JiraExpressionEvaluateRequestBean**](JiraExpressionEvaluateRequestBean.md) | The Jira expression and the evaluation context. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;meta.complexity&#x60; that returns information about the expression complexity. For example, the number of expensive operations used by the expression and how close the expression is to reaching the [complexity limit](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#restrictions). Useful when designing and debugging your expressions. | 

### Return type

[**JExpEvaluateJiraExpressionResultBean**](JExpEvaluateJiraExpressionResultBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EvaluateJiraExpression

> JiraExpressionResult EvaluateJiraExpression(ctx).JiraExpressionEvalRequestBean(jiraExpressionEvalRequestBean).Expand(expand).Execute()

Currently being removed. Evaluate Jira expression



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	jiraExpressionEvalRequestBean := *openapiclient.NewJiraExpressionEvalRequestBean("{ key: issue.key, type: issue.issueType.name, links: issue.links.map(link => link.linkedIssue.id) }") // JiraExpressionEvalRequestBean | The Jira expression and the evaluation context.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts `meta.complexity` that returns information about the expression complexity. For example, the number of expensive operations used by the expression and how close the expression is to reaching the [complexity limit](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#restrictions). Useful when designing and debugging your expressions. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraExpressionsAPI.EvaluateJiraExpression(context.Background()).JiraExpressionEvalRequestBean(jiraExpressionEvalRequestBean).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraExpressionsAPI.EvaluateJiraExpression``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EvaluateJiraExpression`: JiraExpressionResult
	fmt.Fprintf(os.Stdout, "Response from `JiraExpressionsAPI.EvaluateJiraExpression`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEvaluateJiraExpressionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jiraExpressionEvalRequestBean** | [**JiraExpressionEvalRequestBean**](JiraExpressionEvalRequestBean.md) | The Jira expression and the evaluation context. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;meta.complexity&#x60; that returns information about the expression complexity. For example, the number of expensive operations used by the expression and how close the expression is to reaching the [complexity limit](https://developer.atlassian.com/cloud/jira/platform/jira-expressions/#restrictions). Useful when designing and debugging your expressions. | 

### Return type

[**JiraExpressionResult**](JiraExpressionResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

