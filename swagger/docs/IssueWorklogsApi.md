# \IssueWorklogsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddWorklog**](IssueWorklogsAPI.md#AddWorklog) | **Post** /rest/api/3/issue/{issueIdOrKey}/worklog | Add worklog
[**BulkDeleteWorklogs**](IssueWorklogsAPI.md#BulkDeleteWorklogs) | **Delete** /rest/api/3/issue/{issueIdOrKey}/worklog | Bulk delete worklogs
[**BulkMoveWorklogs**](IssueWorklogsAPI.md#BulkMoveWorklogs) | **Post** /rest/api/3/issue/{issueIdOrKey}/worklog/move | Bulk move worklogs
[**DeleteWorklog**](IssueWorklogsAPI.md#DeleteWorklog) | **Delete** /rest/api/3/issue/{issueIdOrKey}/worklog/{id} | Delete worklog
[**GetIdsOfWorklogsDeletedSince**](IssueWorklogsAPI.md#GetIdsOfWorklogsDeletedSince) | **Get** /rest/api/3/worklog/deleted | Get IDs of deleted worklogs
[**GetIdsOfWorklogsModifiedSince**](IssueWorklogsAPI.md#GetIdsOfWorklogsModifiedSince) | **Get** /rest/api/3/worklog/updated | Get IDs of updated worklogs
[**GetIssueWorklog**](IssueWorklogsAPI.md#GetIssueWorklog) | **Get** /rest/api/3/issue/{issueIdOrKey}/worklog | Get issue worklogs
[**GetWorklog**](IssueWorklogsAPI.md#GetWorklog) | **Get** /rest/api/3/issue/{issueIdOrKey}/worklog/{id} | Get worklog
[**GetWorklogsForIds**](IssueWorklogsAPI.md#GetWorklogsForIds) | **Post** /rest/api/3/worklog/list | Get worklogs
[**UpdateWorklog**](IssueWorklogsAPI.md#UpdateWorklog) | **Put** /rest/api/3/issue/{issueIdOrKey}/worklog/{id} | Update worklog



## AddWorklog

> Worklog AddWorklog(ctx, issueIdOrKey).Worklog(worklog).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).ReduceBy(reduceBy).Expand(expand).OverrideEditableFlag(overrideEditableFlag).Execute()

Add worklog



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key the issue.
	worklog := *openapiclient.NewWorklog() // Worklog | 
	notifyUsers := true // bool | Whether users watching the issue are notified by email. (optional) (default to true)
	adjustEstimate := "adjustEstimate_example" // string | Defines how to update the issue's time estimate, the options are:   *  `new` Sets the estimate to a specific value, defined in `newEstimate`.  *  `leave` Leaves the estimate unchanged.  *  `manual` Reduces the estimate by amount specified in `reduceBy`.  *  `auto` Reduces the estimate by the value of `timeSpent` in the worklog. (optional) (default to "auto")
	newEstimate := "newEstimate_example" // string | The value to set as the issue's remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when `adjustEstimate` is `new`. (optional)
	reduceBy := "reduceBy_example" // string | The amount to reduce the issue's remaining estimate by, as days (\\#d), hours (\\#h), or minutes (\\#m). For example, *2d*. Required when `adjustEstimate` is `manual`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts `properties`, which returns worklog properties. (optional) (default to "")
	overrideEditableFlag := true // bool | Whether the worklog entry should be added to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can use this flag. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.AddWorklog(context.Background(), issueIdOrKey).Worklog(worklog).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).ReduceBy(reduceBy).Expand(expand).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.AddWorklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddWorklog`: Worklog
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.AddWorklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddWorklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **worklog** | [**Worklog**](Worklog.md) |  | 
 **notifyUsers** | **bool** | Whether users watching the issue are notified by email. | [default to true]
 **adjustEstimate** | **string** | Defines how to update the issue&#39;s time estimate, the options are:   *  &#x60;new&#x60; Sets the estimate to a specific value, defined in &#x60;newEstimate&#x60;.  *  &#x60;leave&#x60; Leaves the estimate unchanged.  *  &#x60;manual&#x60; Reduces the estimate by amount specified in &#x60;reduceBy&#x60;.  *  &#x60;auto&#x60; Reduces the estimate by the value of &#x60;timeSpent&#x60; in the worklog. | [default to &quot;auto&quot;]
 **newEstimate** | **string** | The value to set as the issue&#39;s remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when &#x60;adjustEstimate&#x60; is &#x60;new&#x60;. | 
 **reduceBy** | **string** | The amount to reduce the issue&#39;s remaining estimate by, as days (\\#d), hours (\\#h), or minutes (\\#m). For example, *2d*. Required when &#x60;adjustEstimate&#x60; is &#x60;manual&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts &#x60;properties&#x60;, which returns worklog properties. | [default to &quot;&quot;]
 **overrideEditableFlag** | **bool** | Whether the worklog entry should be added to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can use this flag. | [default to false]

### Return type

[**Worklog**](Worklog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkDeleteWorklogs

> BulkDeleteWorklogs(ctx, issueIdOrKey).WorklogIdsRequestBean(worklogIdsRequestBean).AdjustEstimate(adjustEstimate).OverrideEditableFlag(overrideEditableFlag).Execute()

Bulk delete worklogs



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
	worklogIdsRequestBean := *openapiclient.NewWorklogIdsRequestBean([]int64{int64(123)}) // WorklogIdsRequestBean | A JSON object containing a list of worklog IDs.
	adjustEstimate := "adjustEstimate_example" // string | Defines how to update the issue's time estimate, the options are:   *  `leave` Leaves the estimate unchanged.  *  `auto` Reduces the estimate by the aggregate value of `timeSpent` across all worklogs being deleted. (optional) (default to "auto")
	overrideEditableFlag := true // bool | Whether the work log entries should be removed to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueWorklogsAPI.BulkDeleteWorklogs(context.Background(), issueIdOrKey).WorklogIdsRequestBean(worklogIdsRequestBean).AdjustEstimate(adjustEstimate).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.BulkDeleteWorklogs``: %v\n", err)
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

Other parameters are passed through a pointer to a apiBulkDeleteWorklogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **worklogIdsRequestBean** | [**WorklogIdsRequestBean**](WorklogIdsRequestBean.md) | A JSON object containing a list of worklog IDs. | 
 **adjustEstimate** | **string** | Defines how to update the issue&#39;s time estimate, the options are:   *  &#x60;leave&#x60; Leaves the estimate unchanged.  *  &#x60;auto&#x60; Reduces the estimate by the aggregate value of &#x60;timeSpent&#x60; across all worklogs being deleted. | [default to &quot;auto&quot;]
 **overrideEditableFlag** | **bool** | Whether the work log entries should be removed to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkMoveWorklogs

> BulkMoveWorklogs(ctx, issueIdOrKey).WorklogsMoveRequestBean(worklogsMoveRequestBean).AdjustEstimate(adjustEstimate).OverrideEditableFlag(overrideEditableFlag).Execute()

Bulk move worklogs



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
	issueIdOrKey := "issueIdOrKey_example" // string | 
	worklogsMoveRequestBean := *openapiclient.NewWorklogsMoveRequestBean() // WorklogsMoveRequestBean | A JSON object containing a list of worklog IDs and the ID or key of the destination issue.
	adjustEstimate := "adjustEstimate_example" // string | Defines how to update the issues' time estimate, the options are:   *  `leave` Leaves the estimate unchanged.  *  `auto` Reduces the estimate by the aggregate value of `timeSpent` across all worklogs being moved in the source issue, and increases it in the destination issue. (optional) (default to "auto")
	overrideEditableFlag := true // bool | Whether the work log entry should be moved to and from the issues even if the issues are not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueWorklogsAPI.BulkMoveWorklogs(context.Background(), issueIdOrKey).WorklogsMoveRequestBean(worklogsMoveRequestBean).AdjustEstimate(adjustEstimate).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.BulkMoveWorklogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkMoveWorklogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **worklogsMoveRequestBean** | [**WorklogsMoveRequestBean**](WorklogsMoveRequestBean.md) | A JSON object containing a list of worklog IDs and the ID or key of the destination issue. | 
 **adjustEstimate** | **string** | Defines how to update the issues&#39; time estimate, the options are:   *  &#x60;leave&#x60; Leaves the estimate unchanged.  *  &#x60;auto&#x60; Reduces the estimate by the aggregate value of &#x60;timeSpent&#x60; across all worklogs being moved in the source issue, and increases it in the destination issue. | [default to &quot;auto&quot;]
 **overrideEditableFlag** | **bool** | Whether the work log entry should be moved to and from the issues even if the issues are not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorklog

> DeleteWorklog(ctx, issueIdOrKey, id).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).IncreaseBy(increaseBy).OverrideEditableFlag(overrideEditableFlag).Execute()

Delete worklog



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
	id := "id_example" // string | The ID of the worklog.
	notifyUsers := true // bool | Whether users watching the issue are notified by email. (optional) (default to true)
	adjustEstimate := "adjustEstimate_example" // string | Defines how to update the issue's time estimate, the options are:   *  `new` Sets the estimate to a specific value, defined in `newEstimate`.  *  `leave` Leaves the estimate unchanged.  *  `manual` Increases the estimate by amount specified in `increaseBy`.  *  `auto` Reduces the estimate by the value of `timeSpent` in the worklog. (optional) (default to "auto")
	newEstimate := "newEstimate_example" // string | The value to set as the issue's remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when `adjustEstimate` is `new`. (optional)
	increaseBy := "increaseBy_example" // string | The amount to increase the issue's remaining estimate by, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when `adjustEstimate` is `manual`. (optional)
	overrideEditableFlag := true // bool | Whether the work log entry should be added to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueWorklogsAPI.DeleteWorklog(context.Background(), issueIdOrKey, id).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).IncreaseBy(increaseBy).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.DeleteWorklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**id** | **string** | The ID of the worklog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **notifyUsers** | **bool** | Whether users watching the issue are notified by email. | [default to true]
 **adjustEstimate** | **string** | Defines how to update the issue&#39;s time estimate, the options are:   *  &#x60;new&#x60; Sets the estimate to a specific value, defined in &#x60;newEstimate&#x60;.  *  &#x60;leave&#x60; Leaves the estimate unchanged.  *  &#x60;manual&#x60; Increases the estimate by amount specified in &#x60;increaseBy&#x60;.  *  &#x60;auto&#x60; Reduces the estimate by the value of &#x60;timeSpent&#x60; in the worklog. | [default to &quot;auto&quot;]
 **newEstimate** | **string** | The value to set as the issue&#39;s remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when &#x60;adjustEstimate&#x60; is &#x60;new&#x60;. | 
 **increaseBy** | **string** | The amount to increase the issue&#39;s remaining estimate by, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when &#x60;adjustEstimate&#x60; is &#x60;manual&#x60;. | 
 **overrideEditableFlag** | **bool** | Whether the work log entry should be added to the issue even if the issue is not editable, because jira.issue.editable set to false or missing. For example, the issue is closed. Connect and Forge app users with admin permission can use this flag. | [default to false]

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


## GetIdsOfWorklogsDeletedSince

> ChangedWorklogs GetIdsOfWorklogsDeletedSince(ctx).Since(since).Execute()

Get IDs of deleted worklogs



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
	since := int64(789) // int64 | The date and time, as a UNIX timestamp in milliseconds, after which deleted worklogs are returned. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.GetIdsOfWorklogsDeletedSince(context.Background()).Since(since).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.GetIdsOfWorklogsDeletedSince``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdsOfWorklogsDeletedSince`: ChangedWorklogs
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.GetIdsOfWorklogsDeletedSince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdsOfWorklogsDeletedSinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int64** | The date and time, as a UNIX timestamp in milliseconds, after which deleted worklogs are returned. | [default to 0]

### Return type

[**ChangedWorklogs**](ChangedWorklogs.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIdsOfWorklogsModifiedSince

> ChangedWorklogs GetIdsOfWorklogsModifiedSince(ctx).Since(since).Expand(expand).Execute()

Get IDs of updated worklogs



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
	since := int64(789) // int64 | The date and time, as a UNIX timestamp in milliseconds, after which updated worklogs are returned. (optional) (default to 0)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `properties` that returns the properties of each worklog. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.GetIdsOfWorklogsModifiedSince(context.Background()).Since(since).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.GetIdsOfWorklogsModifiedSince``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIdsOfWorklogsModifiedSince`: ChangedWorklogs
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.GetIdsOfWorklogsModifiedSince`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIdsOfWorklogsModifiedSinceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **since** | **int64** | The date and time, as a UNIX timestamp in milliseconds, after which updated worklogs are returned. | [default to 0]
 **expand** | **string** | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts &#x60;properties&#x60; that returns the properties of each worklog. | [default to &quot;&quot;]

### Return type

[**ChangedWorklogs**](ChangedWorklogs.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueWorklog

> PageOfWorklogs GetIssueWorklog(ctx, issueIdOrKey).StartAt(startAt).MaxResults(maxResults).StartedAfter(startedAfter).StartedBefore(startedBefore).Expand(expand).Execute()

Get issue worklogs



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 5000)
	startedAfter := int64(789) // int64 | The worklog start date and time, as a UNIX timestamp in milliseconds, after which worklogs are returned. (optional)
	startedBefore := int64(789) // int64 | The worklog start date and time, as a UNIX timestamp in milliseconds, before which worklogs are returned. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts`properties`, which returns worklog properties. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.GetIssueWorklog(context.Background(), issueIdOrKey).StartAt(startAt).MaxResults(maxResults).StartedAfter(startedAfter).StartedBefore(startedBefore).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.GetIssueWorklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueWorklog`: PageOfWorklogs
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.GetIssueWorklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueWorklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 5000]
 **startedAfter** | **int64** | The worklog start date and time, as a UNIX timestamp in milliseconds, after which worklogs are returned. | 
 **startedBefore** | **int64** | The worklog start date and time, as a UNIX timestamp in milliseconds, before which worklogs are returned. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts&#x60;properties&#x60;, which returns worklog properties. | [default to &quot;&quot;]

### Return type

[**PageOfWorklogs**](PageOfWorklogs.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorklog

> Worklog GetWorklog(ctx, issueIdOrKey, id).Expand(expand).Execute()

Get worklog



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
	id := "id_example" // string | The ID of the worklog.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts  `properties`, which returns worklog properties. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.GetWorklog(context.Background(), issueIdOrKey, id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.GetWorklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorklog`: Worklog
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.GetWorklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**id** | **string** | The ID of the worklog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Use [expand](#expansion) to include additional information about work logs in the response. This parameter accepts  &#x60;properties&#x60;, which returns worklog properties. | [default to &quot;&quot;]

### Return type

[**Worklog**](Worklog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorklogsForIds

> []Worklog GetWorklogsForIds(ctx).WorklogIdsRequestBean(worklogIdsRequestBean).Expand(expand).Execute()

Get worklogs



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
	worklogIdsRequestBean := *openapiclient.NewWorklogIdsRequestBean([]int64{int64(123)}) // WorklogIdsRequestBean | A JSON object containing a list of worklog IDs.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `properties` that returns the properties of each worklog. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.GetWorklogsForIds(context.Background()).WorklogIdsRequestBean(worklogIdsRequestBean).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.GetWorklogsForIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorklogsForIds`: []Worklog
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.GetWorklogsForIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorklogsForIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **worklogIdsRequestBean** | [**WorklogIdsRequestBean**](WorklogIdsRequestBean.md) | A JSON object containing a list of worklog IDs. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts &#x60;properties&#x60; that returns the properties of each worklog. | [default to &quot;&quot;]

### Return type

[**[]Worklog**](Worklog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorklog

> Worklog UpdateWorklog(ctx, issueIdOrKey, id).Worklog(worklog).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).Expand(expand).OverrideEditableFlag(overrideEditableFlag).Execute()

Update worklog



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key the issue.
	id := "id_example" // string | The ID of the worklog.
	worklog := *openapiclient.NewWorklog() // Worklog | 
	notifyUsers := true // bool | Whether users watching the issue are notified by email. (optional) (default to true)
	adjustEstimate := "adjustEstimate_example" // string | Defines how to update the issue's time estimate, the options are:   *  `new` Sets the estimate to a specific value, defined in `newEstimate`.  *  `leave` Leaves the estimate unchanged.  *  `auto` Updates the estimate by the difference between the original and updated value of `timeSpent` or `timeSpentSeconds`. (optional) (default to "auto")
	newEstimate := "newEstimate_example" // string | The value to set as the issue's remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when `adjustEstimate` is `new`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts `properties`, which returns worklog properties. (optional) (default to "")
	overrideEditableFlag := true // bool | Whether the worklog should be added to the issue even if the issue is not editable. For example, because the issue is closed. Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can use this flag. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogsAPI.UpdateWorklog(context.Background(), issueIdOrKey, id).Worklog(worklog).NotifyUsers(notifyUsers).AdjustEstimate(adjustEstimate).NewEstimate(newEstimate).Expand(expand).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogsAPI.UpdateWorklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorklog`: Worklog
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogsAPI.UpdateWorklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key the issue. | 
**id** | **string** | The ID of the worklog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **worklog** | [**Worklog**](Worklog.md) |  | 
 **notifyUsers** | **bool** | Whether users watching the issue are notified by email. | [default to true]
 **adjustEstimate** | **string** | Defines how to update the issue&#39;s time estimate, the options are:   *  &#x60;new&#x60; Sets the estimate to a specific value, defined in &#x60;newEstimate&#x60;.  *  &#x60;leave&#x60; Leaves the estimate unchanged.  *  &#x60;auto&#x60; Updates the estimate by the difference between the original and updated value of &#x60;timeSpent&#x60; or &#x60;timeSpentSeconds&#x60;. | [default to &quot;auto&quot;]
 **newEstimate** | **string** | The value to set as the issue&#39;s remaining time estimate, as days (\\#d), hours (\\#h), or minutes (\\#m or \\#). For example, *2d*. Required when &#x60;adjustEstimate&#x60; is &#x60;new&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about worklogs in the response. This parameter accepts &#x60;properties&#x60;, which returns worklog properties. | [default to &quot;&quot;]
 **overrideEditableFlag** | **bool** | Whether the worklog should be added to the issue even if the issue is not editable. For example, because the issue is closed. Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) can use this flag. | [default to false]

### Return type

[**Worklog**](Worklog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

