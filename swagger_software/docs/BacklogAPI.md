# \BacklogAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**MoveIssuesToBacklog**](BacklogAPI.md#MoveIssuesToBacklog) | **Post** /rest/agile/1.0/backlog/issue | Move issues to backlog
[**MoveIssuesToBacklogForBoard**](BacklogAPI.md#MoveIssuesToBacklogForBoard) | **Post** /rest/agile/1.0/backlog/{boardId}/issue | Move issues to backlog for board



## MoveIssuesToBacklog

> MoveIssuesToBacklog(ctx).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()

Move issues to backlog



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
	moveIssuesToBacklogRequest := *openapiclient.NewMoveIssuesToBacklogRequest() // MoveIssuesToBacklogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BacklogAPI.MoveIssuesToBacklog(context.Background()).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BacklogAPI.MoveIssuesToBacklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveIssuesToBacklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moveIssuesToBacklogRequest** | [**MoveIssuesToBacklogRequest**](MoveIssuesToBacklogRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveIssuesToBacklogForBoard

> MoveIssuesToBacklogForBoard(ctx, boardId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()

Move issues to backlog for board



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
	boardId := int64(789) // int64 | 
	moveIssuesToBacklogForBoardRequest := *openapiclient.NewMoveIssuesToBacklogForBoardRequest() // MoveIssuesToBacklogForBoardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BacklogAPI.MoveIssuesToBacklogForBoard(context.Background(), boardId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BacklogAPI.MoveIssuesToBacklogForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveIssuesToBacklogForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveIssuesToBacklogForBoardRequest** | [**MoveIssuesToBacklogForBoardRequest**](MoveIssuesToBacklogForBoardRequest.md) |  | 

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

