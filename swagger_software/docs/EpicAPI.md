# \EpicAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetEpic**](EpicAPI.md#GetEpic) | **Get** /rest/agile/1.0/epic/{epicIdOrKey} | Get epic
[**GetIssuesForEpic**](EpicAPI.md#GetIssuesForEpic) | **Get** /rest/agile/1.0/epic/{epicIdOrKey}/issue | Get issues for epic
[**GetIssuesWithoutEpic**](EpicAPI.md#GetIssuesWithoutEpic) | **Get** /rest/agile/1.0/epic/none/issue | Get issues without epic
[**MoveIssuesToEpic**](EpicAPI.md#MoveIssuesToEpic) | **Post** /rest/agile/1.0/epic/{epicIdOrKey}/issue | Move issues to epic
[**PartiallyUpdateEpic**](EpicAPI.md#PartiallyUpdateEpic) | **Post** /rest/agile/1.0/epic/{epicIdOrKey} | Partially update epic
[**RankEpics**](EpicAPI.md#RankEpics) | **Put** /rest/agile/1.0/epic/{epicIdOrKey}/rank | Rank epics
[**RemoveIssuesFromEpic**](EpicAPI.md#RemoveIssuesFromEpic) | **Post** /rest/agile/1.0/epic/none/issue | Remove issues from epic



## GetEpic

> GetEpic(ctx, epicIdOrKey).Execute()

Get epic



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
	epicIdOrKey := "epicIdOrKey_example" // string | The id or key of the requested epic.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.GetEpic(context.Background(), epicIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.GetEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicIdOrKey** | **string** | The id or key of the requested epic. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetIssuesForEpic

> GetIssuesForEpic(ctx, epicIdOrKey).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues for epic



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
	epicIdOrKey := "epicIdOrKey_example" // string | The id or key of the epic that contains the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.GetIssuesForEpic(context.Background(), epicIdOrKey).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.GetIssuesForEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicIdOrKey** | **string** | The id or key of the epic that contains the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesForEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. Default: 50. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that &#x60;username&#x60; and &#x60;userkey&#x60; can&#39;t be used as search terms for this parameter due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 

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


## GetIssuesWithoutEpic

> GetIssuesWithoutEpic(ctx).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues without epic



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
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.GetIssuesWithoutEpic(context.Background()).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.GetIssuesWithoutEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesWithoutEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 

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


## MoveIssuesToEpic

> MoveIssuesToEpic(ctx, epicIdOrKey).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()

Move issues to epic



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
	epicIdOrKey := "epicIdOrKey_example" // string | The id or key of the epic that you want to assign issues to.
	moveIssuesToBacklogRequest := *openapiclient.NewMoveIssuesToBacklogRequest() // MoveIssuesToBacklogRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.MoveIssuesToEpic(context.Background(), epicIdOrKey).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.MoveIssuesToEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicIdOrKey** | **string** | The id or key of the epic that you want to assign issues to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveIssuesToEpicRequest struct via the builder pattern


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


## PartiallyUpdateEpic

> PartiallyUpdateEpic(ctx, epicIdOrKey).PartiallyUpdateEpicRequest(partiallyUpdateEpicRequest).Execute()

Partially update epic



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
	epicIdOrKey := "epicIdOrKey_example" // string | The id or key of the epic to update.
	partiallyUpdateEpicRequest := *openapiclient.NewPartiallyUpdateEpicRequest() // PartiallyUpdateEpicRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.PartiallyUpdateEpic(context.Background(), epicIdOrKey).PartiallyUpdateEpicRequest(partiallyUpdateEpicRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.PartiallyUpdateEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicIdOrKey** | **string** | The id or key of the epic to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartiallyUpdateEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **partiallyUpdateEpicRequest** | [**PartiallyUpdateEpicRequest**](PartiallyUpdateEpicRequest.md) |  | 

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


## RankEpics

> RankEpics(ctx, epicIdOrKey).RankEpicsRequest(rankEpicsRequest).Execute()

Rank epics



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
	epicIdOrKey := "epicIdOrKey_example" // string | The id or key of the epic to rank.
	rankEpicsRequest := *openapiclient.NewRankEpicsRequest() // RankEpicsRequest | bean which contains the information where the given epic should be ranked.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.EpicAPI.RankEpics(context.Background(), epicIdOrKey).RankEpicsRequest(rankEpicsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.RankEpics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**epicIdOrKey** | **string** | The id or key of the epic to rank. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRankEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **rankEpicsRequest** | [**RankEpicsRequest**](RankEpicsRequest.md) | bean which contains the information where the given epic should be ranked. | 

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


## RemoveIssuesFromEpic

> RemoveIssuesFromEpic(ctx).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()

Remove issues from epic



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
	r, err := apiClient.EpicAPI.RemoveIssuesFromEpic(context.Background()).MoveIssuesToBacklogRequest(moveIssuesToBacklogRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EpicAPI.RemoveIssuesFromEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIssuesFromEpicRequest struct via the builder pattern


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

