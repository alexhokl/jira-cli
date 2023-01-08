# \IssueWatchersAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWatcher**](IssueWatchersAPI.md#AddWatcher) | **Post** /rest/api/3/issue/{issueIdOrKey}/watchers | Add watcher
[**GetIsWatchingIssueBulk**](IssueWatchersAPI.md#GetIsWatchingIssueBulk) | **Post** /rest/api/3/issue/watching | Get is watching issue bulk
[**GetIssueWatchers**](IssueWatchersAPI.md#GetIssueWatchers) | **Get** /rest/api/3/issue/{issueIdOrKey}/watchers | Get issue watchers
[**RemoveWatcher**](IssueWatchersAPI.md#RemoveWatcher) | **Delete** /rest/api/3/issue/{issueIdOrKey}/watchers | Delete watcher



## AddWatcher

> interface{} AddWatcher(ctx, issueIdOrKey).Body(body).Execute()

Add watcher



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	body := "body_example" // string | The account ID of the user. Note that username cannot be used due to privacy changes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWatchersAPI.AddWatcher(context.Background(), issueIdOrKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWatchersAPI.AddWatcher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWatcher`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueWatchersAPI.AddWatcher`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWatcherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **string** | The account ID of the user. Note that username cannot be used due to privacy changes. | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIsWatchingIssueBulk

> BulkIssueIsWatching GetIsWatchingIssueBulk(ctx).IssueList(issueList).Execute()

Get is watching issue bulk



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
	issueList := *openapiclient.NewIssueList([]string{"IssueIds_example"}) // IssueList | A list of issue IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWatchersAPI.GetIsWatchingIssueBulk(context.Background()).IssueList(issueList).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWatchersAPI.GetIsWatchingIssueBulk``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIsWatchingIssueBulk`: BulkIssueIsWatching
	fmt.Fprintf(os.Stdout, "Response from `IssueWatchersAPI.GetIsWatchingIssueBulk`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIsWatchingIssueBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueList** | [**IssueList**](IssueList.md) | A list of issue IDs. | 

### Return type

[**BulkIssueIsWatching**](BulkIssueIsWatching.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueWatchers

> Watchers GetIssueWatchers(ctx, issueIdOrKey).Execute()

Get issue watchers



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWatchersAPI.GetIssueWatchers(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWatchersAPI.GetIssueWatchers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueWatchers`: Watchers
	fmt.Fprintf(os.Stdout, "Response from `IssueWatchersAPI.GetIssueWatchers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueWatchersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Watchers**](Watchers.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveWatcher

> RemoveWatcher(ctx, issueIdOrKey).Username(username).AccountId(accountId).Execute()

Delete watcher



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	accountId := "5b10ac8d82e05b22cc7d4ef5" // string | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueWatchersAPI.RemoveWatcher(context.Background(), issueIdOrKey).Username(username).AccountId(accountId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWatchersAPI.RemoveWatcher``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveWatcherRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **accountId** | **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

