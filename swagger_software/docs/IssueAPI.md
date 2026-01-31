# \IssueAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**EstimateIssueForBoard**](IssueAPI.md#EstimateIssueForBoard) | **Put** /rest/agile/1.0/issue/{issueIdOrKey}/estimation | Estimate issue for board
[**GetIssue**](IssueAPI.md#GetIssue) | **Get** /rest/agile/1.0/issue/{issueIdOrKey} | Get issue
[**GetIssueEstimationForBoard**](IssueAPI.md#GetIssueEstimationForBoard) | **Get** /rest/agile/1.0/issue/{issueIdOrKey}/estimation | Get issue estimation for board
[**RankIssues**](IssueAPI.md#RankIssues) | **Put** /rest/agile/1.0/issue/rank | Rank issues



## EstimateIssueForBoard

> EstimateIssueForBoard(ctx, issueIdOrKey).EstimateIssueForBoardRequest(estimateIssueForBoardRequest).BoardId(boardId).Execute()

Estimate issue for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the requested issue.
	estimateIssueForBoardRequest := *openapiclient.NewEstimateIssueForBoardRequest() // EstimateIssueForBoardRequest | bean that contains value of a new estimation.
	boardId := int64(789) // int64 | The ID of the board required to determine which field is used for estimation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.EstimateIssueForBoard(context.Background(), issueIdOrKey).EstimateIssueForBoardRequest(estimateIssueForBoardRequest).BoardId(boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.EstimateIssueForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the requested issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEstimateIssueForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **estimateIssueForBoardRequest** | [**EstimateIssueForBoardRequest**](EstimateIssueForBoardRequest.md) | bean that contains value of a new estimation. | 
 **boardId** | **int64** | The ID of the board required to determine which field is used for estimation. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssue

> GetIssue(ctx, issueIdOrKey).Fields(fields).Expand(expand).UpdateHistory(updateHistory).Execute()

Get issue



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the requested issue.
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)
	updateHistory := true // bool | A boolean indicating whether the issue retrieved by this method should be added to the current user's issue history (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.GetIssue(context.Background(), issueIdOrKey).Fields(fields).Expand(expand).UpdateHistory(updateHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the requested issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 
 **updateHistory** | **bool** | A boolean indicating whether the issue retrieved by this method should be added to the current user&#39;s issue history | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueEstimationForBoard

> GetIssueEstimationForBoard(ctx, issueIdOrKey).BoardId(boardId).Execute()

Get issue estimation for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the requested issue.
	boardId := int64(789) // int64 | The ID of the board required to determine which field is used for estimation. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.GetIssueEstimationForBoard(context.Background(), issueIdOrKey).BoardId(boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.GetIssueEstimationForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the requested issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueEstimationForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **boardId** | **int64** | The ID of the board required to determine which field is used for estimation. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RankIssues

> RankIssues(ctx).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()

Rank issues



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	moveIssuesToBacklogForBoardRequest := *openapiclient.NewMoveIssuesToBacklogForBoardRequest() // MoveIssuesToBacklogForBoardRequest | bean which contains list of issues to rank and information where it should be ranked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAPI.RankIssues(context.Background()).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAPI.RankIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRankIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveIssuesToBacklogForBoardRequest** | [**MoveIssuesToBacklogForBoardRequest**](MoveIssuesToBacklogForBoardRequest.md) | bean which contains list of issues to rank and information where it should be ranked. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

