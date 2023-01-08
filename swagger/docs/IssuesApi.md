# \IssuesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchiveIssues**](IssuesAPI.md#ArchiveIssues) | **Put** /rest/api/3/issue/archive | Archive issue(s) by issue ID/key
[**ArchiveIssuesAsync**](IssuesAPI.md#ArchiveIssuesAsync) | **Post** /rest/api/3/issue/archive | Archive issue(s) by JQL
[**AssignIssue**](IssuesAPI.md#AssignIssue) | **Put** /rest/api/3/issue/{issueIdOrKey}/assignee | Assign issue
[**BulkFetchIssues**](IssuesAPI.md#BulkFetchIssues) | **Post** /rest/api/3/issue/bulkfetch | Bulk fetch issues
[**CreateIssue**](IssuesAPI.md#CreateIssue) | **Post** /rest/api/3/issue | Create issue
[**CreateIssues**](IssuesAPI.md#CreateIssues) | **Post** /rest/api/3/issue/bulk | Bulk create issue
[**DeleteIssue**](IssuesAPI.md#DeleteIssue) | **Delete** /rest/api/3/issue/{issueIdOrKey} | Delete issue
[**DoTransition**](IssuesAPI.md#DoTransition) | **Post** /rest/api/3/issue/{issueIdOrKey}/transitions | Transition issue
[**EditIssue**](IssuesAPI.md#EditIssue) | **Put** /rest/api/3/issue/{issueIdOrKey} | Edit issue
[**ExportArchivedIssues**](IssuesAPI.md#ExportArchivedIssues) | **Put** /rest/api/3/issues/archive/export | Export archived issue(s)
[**GetBulkChangelogs**](IssuesAPI.md#GetBulkChangelogs) | **Post** /rest/api/3/changelog/bulkfetch | Bulk fetch changelogs
[**GetChangeLogs**](IssuesAPI.md#GetChangeLogs) | **Get** /rest/api/3/issue/{issueIdOrKey}/changelog | Get changelogs
[**GetChangeLogsByIds**](IssuesAPI.md#GetChangeLogsByIds) | **Post** /rest/api/3/issue/{issueIdOrKey}/changelog/list | Get changelogs by IDs
[**GetCreateIssueMeta**](IssuesAPI.md#GetCreateIssueMeta) | **Get** /rest/api/3/issue/createmeta | Get create issue metadata
[**GetCreateIssueMetaIssueTypeId**](IssuesAPI.md#GetCreateIssueMetaIssueTypeId) | **Get** /rest/api/3/issue/createmeta/{projectIdOrKey}/issuetypes/{issueTypeId} | Get create field metadata for a project and issue type id
[**GetCreateIssueMetaIssueTypes**](IssuesAPI.md#GetCreateIssueMetaIssueTypes) | **Get** /rest/api/3/issue/createmeta/{projectIdOrKey}/issuetypes | Get create metadata issue types for a project
[**GetEditIssueMeta**](IssuesAPI.md#GetEditIssueMeta) | **Get** /rest/api/3/issue/{issueIdOrKey}/editmeta | Get edit issue metadata
[**GetEvents**](IssuesAPI.md#GetEvents) | **Get** /rest/api/3/events | Get events
[**GetIssue**](IssuesAPI.md#GetIssue) | **Get** /rest/api/3/issue/{issueIdOrKey} | Get issue
[**GetIssueLimitReport**](IssuesAPI.md#GetIssueLimitReport) | **Get** /rest/api/3/issue/limit/report | Get issue limit report
[**GetTransitions**](IssuesAPI.md#GetTransitions) | **Get** /rest/api/3/issue/{issueIdOrKey}/transitions | Get transitions
[**Notify**](IssuesAPI.md#Notify) | **Post** /rest/api/3/issue/{issueIdOrKey}/notify | Send notification for issue
[**UnarchiveIssues**](IssuesAPI.md#UnarchiveIssues) | **Put** /rest/api/3/issue/unarchive | Unarchive issue(s) by issue keys/ID



## ArchiveIssues

> IssueArchivalSyncResponse ArchiveIssues(ctx).IssueArchivalSyncRequest(issueArchivalSyncRequest).Execute()

Archive issue(s) by issue ID/key



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
	issueArchivalSyncRequest := *openapiclient.NewIssueArchivalSyncRequest() // IssueArchivalSyncRequest | Contains a list of issue keys or IDs to be archived.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.ArchiveIssues(context.Background()).IssueArchivalSyncRequest(issueArchivalSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.ArchiveIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveIssues`: IssueArchivalSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.ArchiveIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueArchivalSyncRequest** | [**IssueArchivalSyncRequest**](IssueArchivalSyncRequest.md) | Contains a list of issue keys or IDs to be archived. | 

### Return type

[**IssueArchivalSyncResponse**](IssueArchivalSyncResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ArchiveIssuesAsync

> string ArchiveIssuesAsync(ctx).ArchiveIssueAsyncRequest(archiveIssueAsyncRequest).Execute()

Archive issue(s) by JQL



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
	archiveIssueAsyncRequest := *openapiclient.NewArchiveIssueAsyncRequest() // ArchiveIssueAsyncRequest | A JQL query specifying the issues to archive. Note that subtasks can only be archived through their parent issues.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.ArchiveIssuesAsync(context.Background()).ArchiveIssueAsyncRequest(archiveIssueAsyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.ArchiveIssuesAsync``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchiveIssuesAsync`: string
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.ArchiveIssuesAsync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiArchiveIssuesAsyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archiveIssueAsyncRequest** | [**ArchiveIssueAsyncRequest**](ArchiveIssueAsyncRequest.md) | A JQL query specifying the issues to archive. Note that subtasks can only be archived through their parent issues. | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssignIssue

> interface{} AssignIssue(ctx, issueIdOrKey).User(user).Execute()

Assign issue



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue to be assigned.
	user := *openapiclient.NewUser() // User | The request object with the user that the issue is assigned to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.AssignIssue(context.Background(), issueIdOrKey).User(user).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.AssignIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignIssue`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.AssignIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue to be assigned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | [**User**](User.md) | The request object with the user that the issue is assigned to. | 

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


## BulkFetchIssues

> BulkIssueResults BulkFetchIssues(ctx).BulkFetchIssueRequestBean(bulkFetchIssueRequestBean).Execute()

Bulk fetch issues



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
	bulkFetchIssueRequestBean := *openapiclient.NewBulkFetchIssueRequestBean([]string{"IssueIdsOrKeys_example"}) // BulkFetchIssueRequestBean | A JSON object containing the information about which issues and fields to fetch.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.BulkFetchIssues(context.Background()).BulkFetchIssueRequestBean(bulkFetchIssueRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.BulkFetchIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkFetchIssues`: BulkIssueResults
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.BulkFetchIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkFetchIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkFetchIssueRequestBean** | [**BulkFetchIssueRequestBean**](BulkFetchIssueRequestBean.md) | A JSON object containing the information about which issues and fields to fetch. | 

### Return type

[**BulkIssueResults**](BulkIssueResults.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssue

> CreatedIssue CreateIssue(ctx).IssueUpdateDetails(issueUpdateDetails).UpdateHistory(updateHistory).Execute()

Create issue



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
	issueUpdateDetails := *openapiclient.NewIssueUpdateDetails() // IssueUpdateDetails | 
	updateHistory := true // bool | Whether the project in which the issue is created is added to the user's **Recently viewed** project list, as shown under **Projects** in Jira. When provided, the issue type and request type are added to the user's history for a project. These values are then used to provide defaults on the issue create screen. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.CreateIssue(context.Background()).IssueUpdateDetails(issueUpdateDetails).UpdateHistory(updateHistory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.CreateIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssue`: CreatedIssue
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.CreateIssue`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueUpdateDetails** | [**IssueUpdateDetails**](IssueUpdateDetails.md) |  | 
 **updateHistory** | **bool** | Whether the project in which the issue is created is added to the user&#39;s **Recently viewed** project list, as shown under **Projects** in Jira. When provided, the issue type and request type are added to the user&#39;s history for a project. These values are then used to provide defaults on the issue create screen. | [default to false]

### Return type

[**CreatedIssue**](CreatedIssue.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssues

> CreatedIssues CreateIssues(ctx).IssuesUpdateBean(issuesUpdateBean).Execute()

Bulk create issue



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
	issuesUpdateBean := *openapiclient.NewIssuesUpdateBean() // IssuesUpdateBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.CreateIssues(context.Background()).IssuesUpdateBean(issuesUpdateBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.CreateIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssues`: CreatedIssues
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.CreateIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issuesUpdateBean** | [**IssuesUpdateBean**](IssuesUpdateBean.md) |  | 

### Return type

[**CreatedIssues**](CreatedIssues.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssue

> DeleteIssue(ctx, issueIdOrKey).DeleteSubtasks(deleteSubtasks).Execute()

Delete issue



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
	deleteSubtasks := "deleteSubtasks_example" // string | Whether the issue's subtasks are deleted when the issue is deleted. (optional) (default to "false")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuesAPI.DeleteIssue(context.Background(), issueIdOrKey).DeleteSubtasks(deleteSubtasks).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.DeleteIssue``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteSubtasks** | **string** | Whether the issue&#39;s subtasks are deleted when the issue is deleted. | [default to &quot;false&quot;]

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


## DoTransition

> interface{} DoTransition(ctx, issueIdOrKey).IssueUpdateDetails(issueUpdateDetails).Execute()

Transition issue



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
	issueUpdateDetails := *openapiclient.NewIssueUpdateDetails() // IssueUpdateDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.DoTransition(context.Background(), issueIdOrKey).IssueUpdateDetails(issueUpdateDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.DoTransition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DoTransition`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.DoTransition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDoTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueUpdateDetails** | [**IssueUpdateDetails**](IssueUpdateDetails.md) |  | 

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


## EditIssue

> interface{} EditIssue(ctx, issueIdOrKey).IssueUpdateDetails(issueUpdateDetails).NotifyUsers(notifyUsers).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).ReturnIssue(returnIssue).Expand(expand).Execute()

Edit issue



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
	issueUpdateDetails := *openapiclient.NewIssueUpdateDetails() // IssueUpdateDetails | 
	notifyUsers := true // bool | Whether a notification email about the issue update is sent to all watchers. To disable the notification, administer Jira or administer project permissions are required. If the user doesn't have the necessary permission the request is ignored. (optional) (default to true)
	overrideScreenSecurity := true // bool | Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)
	overrideEditableFlag := true // bool | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)
	returnIssue := true // bool | Whether the response should contain the issue with fields edited in this request. The returned issue will have the same format as in the [Get issue API](#api-rest-api-3-issue-issueidorkey-get). (optional) (default to false)
	expand := "expand_example" // string | The Get issue API expand parameter to use in the response if the `returnIssue` parameter is `true`. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.EditIssue(context.Background(), issueIdOrKey).IssueUpdateDetails(issueUpdateDetails).NotifyUsers(notifyUsers).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).ReturnIssue(returnIssue).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.EditIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EditIssue`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.EditIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiEditIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueUpdateDetails** | [**IssueUpdateDetails**](IssueUpdateDetails.md) |  | 
 **notifyUsers** | **bool** | Whether a notification email about the issue update is sent to all watchers. To disable the notification, administer Jira or administer project permissions are required. If the user doesn&#39;t have the necessary permission the request is ignored. | [default to true]
 **overrideScreenSecurity** | **bool** | Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]
 **overrideEditableFlag** | **bool** | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]
 **returnIssue** | **bool** | Whether the response should contain the issue with fields edited in this request. The returned issue will have the same format as in the [Get issue API](#api-rest-api-3-issue-issueidorkey-get). | [default to false]
 **expand** | **string** | The Get issue API expand parameter to use in the response if the &#x60;returnIssue&#x60; parameter is &#x60;true&#x60;. | [default to &quot;&quot;]

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


## ExportArchivedIssues

> ExportArchivedIssuesTaskProgressResponse ExportArchivedIssues(ctx).ArchivedIssuesFilterRequest(archivedIssuesFilterRequest).Execute()

Export archived issue(s)



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
	archivedIssuesFilterRequest := *openapiclient.NewArchivedIssuesFilterRequest() // ArchivedIssuesFilterRequest | You can filter the issues in your request by the `projects`, `archivedBy`, `archivedDate`, `issueTypes`, and `reporters` fields. All filters are optional. If you don't provide any filters, you'll get a list of up to one million archived issues.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.ExportArchivedIssues(context.Background()).ArchivedIssuesFilterRequest(archivedIssuesFilterRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.ExportArchivedIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExportArchivedIssues`: ExportArchivedIssuesTaskProgressResponse
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.ExportArchivedIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExportArchivedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **archivedIssuesFilterRequest** | [**ArchivedIssuesFilterRequest**](ArchivedIssuesFilterRequest.md) | You can filter the issues in your request by the &#x60;projects&#x60;, &#x60;archivedBy&#x60;, &#x60;archivedDate&#x60;, &#x60;issueTypes&#x60;, and &#x60;reporters&#x60; fields. All filters are optional. If you don&#39;t provide any filters, you&#39;ll get a list of up to one million archived issues. | 

### Return type

[**ExportArchivedIssuesTaskProgressResponse**](ExportArchivedIssuesTaskProgressResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkChangelogs

> BulkChangelogResponseBean GetBulkChangelogs(ctx).BulkChangelogRequestBean(bulkChangelogRequestBean).Execute()

Bulk fetch changelogs



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
	bulkChangelogRequestBean := *openapiclient.NewBulkChangelogRequestBean([]string{"IssueIdsOrKeys_example"}) // BulkChangelogRequestBean | A JSON object containing the bulk fetch changelog request filters such as issue IDs and field IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetBulkChangelogs(context.Background()).BulkChangelogRequestBean(bulkChangelogRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetBulkChangelogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkChangelogs`: BulkChangelogResponseBean
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetBulkChangelogs`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkChangelogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkChangelogRequestBean** | [**BulkChangelogRequestBean**](BulkChangelogRequestBean.md) | A JSON object containing the bulk fetch changelog request filters such as issue IDs and field IDs. | 

### Return type

[**BulkChangelogResponseBean**](BulkChangelogResponseBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeLogs

> PageBeanChangelog GetChangeLogs(ctx, issueIdOrKey).StartAt(startAt).MaxResults(maxResults).Execute()

Get changelogs



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
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetChangeLogs(context.Background(), issueIdOrKey).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetChangeLogs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangeLogs`: PageBeanChangelog
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetChangeLogs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeLogsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanChangelog**](PageBeanChangelog.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetChangeLogsByIds

> PageOfChangelogs GetChangeLogsByIds(ctx, issueIdOrKey).IssueChangelogIds(issueChangelogIds).Execute()

Get changelogs by IDs



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
	issueChangelogIds := *openapiclient.NewIssueChangelogIds([]int64{int64(123)}) // IssueChangelogIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetChangeLogsByIds(context.Background(), issueIdOrKey).IssueChangelogIds(issueChangelogIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetChangeLogsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetChangeLogsByIds`: PageOfChangelogs
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetChangeLogsByIds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChangeLogsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueChangelogIds** | [**IssueChangelogIds**](IssueChangelogIds.md) |  | 

### Return type

[**PageOfChangelogs**](PageOfChangelogs.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateIssueMeta

> IssueCreateMetadata GetCreateIssueMeta(ctx).ProjectIds(projectIds).ProjectKeys(projectKeys).IssuetypeIds(issuetypeIds).IssuetypeNames(issuetypeNames).Expand(expand).Execute()

Get create issue metadata



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
	projectIds := []string{"Inner_example"} // []string | List of project IDs. This parameter accepts a comma-separated list. Multiple project IDs can also be provided using an ampersand-separated list. For example, `projectIds=10000,10001&projectIds=10020,10021`. This parameter may be provided with `projectKeys`. (optional)
	projectKeys := []string{"Inner_example"} // []string | List of project keys. This parameter accepts a comma-separated list. Multiple project keys can also be provided using an ampersand-separated list. For example, `projectKeys=proj1,proj2&projectKeys=proj3`. This parameter may be provided with `projectIds`. (optional)
	issuetypeIds := []string{"Inner_example"} // []string | List of issue type IDs. This parameter accepts a comma-separated list. Multiple issue type IDs can also be provided using an ampersand-separated list. For example, `issuetypeIds=10000,10001&issuetypeIds=10020,10021`. This parameter may be provided with `issuetypeNames`. (optional)
	issuetypeNames := []string{"Inner_example"} // []string | List of issue type names. This parameter accepts a comma-separated list. Multiple issue type names can also be provided using an ampersand-separated list. For example, `issuetypeNames=name1,name2&issuetypeNames=name3`. This parameter may be provided with `issuetypeIds`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about issue metadata in the response. This parameter accepts `projects.issuetypes.fields`, which returns information about the fields in the issue creation screen for each issue type. Fields hidden from the screen are not returned. Use the information to populate the `fields` and `update` fields in [Create issue](#api-rest-api-3-issue-post) and [Create issues](#api-rest-api-3-issue-bulk-post). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetCreateIssueMeta(context.Background()).ProjectIds(projectIds).ProjectKeys(projectKeys).IssuetypeIds(issuetypeIds).IssuetypeNames(issuetypeNames).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetCreateIssueMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateIssueMeta`: IssueCreateMetadata
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetCreateIssueMeta`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateIssueMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIds** | **[]string** | List of project IDs. This parameter accepts a comma-separated list. Multiple project IDs can also be provided using an ampersand-separated list. For example, &#x60;projectIds&#x3D;10000,10001&amp;projectIds&#x3D;10020,10021&#x60;. This parameter may be provided with &#x60;projectKeys&#x60;. | 
 **projectKeys** | **[]string** | List of project keys. This parameter accepts a comma-separated list. Multiple project keys can also be provided using an ampersand-separated list. For example, &#x60;projectKeys&#x3D;proj1,proj2&amp;projectKeys&#x3D;proj3&#x60;. This parameter may be provided with &#x60;projectIds&#x60;. | 
 **issuetypeIds** | **[]string** | List of issue type IDs. This parameter accepts a comma-separated list. Multiple issue type IDs can also be provided using an ampersand-separated list. For example, &#x60;issuetypeIds&#x3D;10000,10001&amp;issuetypeIds&#x3D;10020,10021&#x60;. This parameter may be provided with &#x60;issuetypeNames&#x60;. | 
 **issuetypeNames** | **[]string** | List of issue type names. This parameter accepts a comma-separated list. Multiple issue type names can also be provided using an ampersand-separated list. For example, &#x60;issuetypeNames&#x3D;name1,name2&amp;issuetypeNames&#x3D;name3&#x60;. This parameter may be provided with &#x60;issuetypeIds&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about issue metadata in the response. This parameter accepts &#x60;projects.issuetypes.fields&#x60;, which returns information about the fields in the issue creation screen for each issue type. Fields hidden from the screen are not returned. Use the information to populate the &#x60;fields&#x60; and &#x60;update&#x60; fields in [Create issue](#api-rest-api-3-issue-post) and [Create issues](#api-rest-api-3-issue-bulk-post). | 

### Return type

[**IssueCreateMetadata**](IssueCreateMetadata.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateIssueMetaIssueTypeId

> PageOfCreateMetaIssueTypeWithField GetCreateIssueMetaIssueTypeId(ctx, projectIdOrKey, issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()

Get create field metadata for a project and issue type id



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
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or key of the project.
	issueTypeId := "issueTypeId_example" // string | The issuetype ID.
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetCreateIssueMetaIssueTypeId(context.Background(), projectIdOrKey, issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetCreateIssueMetaIssueTypeId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateIssueMetaIssueTypeId`: PageOfCreateMetaIssueTypeWithField
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetCreateIssueMetaIssueTypeId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or key of the project. | 
**issueTypeId** | **string** | The issuetype ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateIssueMetaIssueTypeIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageOfCreateMetaIssueTypeWithField**](PageOfCreateMetaIssueTypeWithField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCreateIssueMetaIssueTypes

> PageOfCreateMetaIssueTypes GetCreateIssueMetaIssueTypes(ctx, projectIdOrKey).StartAt(startAt).MaxResults(maxResults).Execute()

Get create metadata issue types for a project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or key of the project.
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetCreateIssueMetaIssueTypes(context.Background(), projectIdOrKey).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetCreateIssueMetaIssueTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCreateIssueMetaIssueTypes`: PageOfCreateMetaIssueTypes
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetCreateIssueMetaIssueTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or key of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreateIssueMetaIssueTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageOfCreateMetaIssueTypes**](PageOfCreateMetaIssueTypes.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEditIssueMeta

> IssueUpdateMetadata GetEditIssueMeta(ctx, issueIdOrKey).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).Execute()

Get edit issue metadata



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
	overrideScreenSecurity := true // bool | Whether hidden fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)
	overrideEditableFlag := true // bool | Whether non-editable fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetEditIssueMeta(context.Background(), issueIdOrKey).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetEditIssueMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEditIssueMeta`: IssueUpdateMetadata
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetEditIssueMeta`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEditIssueMetaRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **overrideScreenSecurity** | **bool** | Whether hidden fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]
 **overrideEditableFlag** | **bool** | Whether non-editable fields are returned. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]

### Return type

[**IssueUpdateMetadata**](IssueUpdateMetadata.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEvents

> []IssueEvent GetEvents(ctx).Execute()

Get events



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetEvents(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetEvents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEvents`: []IssueEvent
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventsRequest struct via the builder pattern


### Return type

[**[]IssueEvent**](IssueEvent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssue

> IssueBean GetIssue(ctx, issueIdOrKey).Fields(fields).FieldsByKeys(fieldsByKeys).Expand(expand).Properties(properties).UpdateHistory(updateHistory).FailFast(failFast).Execute()

Get issue



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
	fields := []string{"Inner_example"} // []string | A list of fields to return for the issue. This parameter accepts a comma-separated list. Use it to retrieve a subset of fields. Allowed values:   *  `*all` Returns all fields.  *  `*navigable` Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  Examples:   *  `summary,comment` Returns only the summary and comments fields.  *  `-description` Returns all (default) fields except description.  *  `*navigable,-comment` Returns all navigable fields except comment.  This parameter may be specified multiple times. For example, `fields=field1,field2& fields=field3`.  Note: All fields are returned by default. This differs from [Search for issues using JQL (GET)](#api-rest-api-3-search-get) and [Search for issues using JQL (POST)](#api-rest-api-3-search-post) where the default is all navigable fields. (optional)
	fieldsByKeys := true // bool | Whether fields in `fields` are referenced by keys rather than IDs. This parameter is useful where fields have been added by a connect app and a field's key may differ from its ID. (optional) (default to false)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about the issues in the response. This parameter accepts a comma-separated list. Expand options include:   *  `renderedFields` Returns field values rendered in HTML format.  *  `names` Returns the display name of each field.  *  `schema` Returns the schema describing a field type.  *  `transitions` Returns all possible transitions for the issue.  *  `editmeta` Returns information about how each field can be edited.  *  `changelog` Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  `versionedRepresentations` Returns a JSON array for each version of a field's value, with the highest number representing the most recent version. Note: When included in the request, the `fields` parameter is ignored. (optional)
	properties := []string{"Inner_example"} // []string | A list of issue properties to return for the issue. This parameter accepts a comma-separated list. Allowed values:   *  `*all` Returns all issue properties.  *  Any issue property key, prefixed with a minus to exclude.  Examples:   *  `*all` Returns all properties.  *  `*all,-prop1` Returns all properties except `prop1`.  *  `prop1,prop2` Returns `prop1` and `prop2` properties.  This parameter may be specified multiple times. For example, `properties=prop1,prop2& properties=prop3`. (optional)
	updateHistory := true // bool | Whether the project in which the issue is created is added to the user's **Recently viewed** project list, as shown under **Projects** in Jira. This also populates the [JQL issues search](#api-rest-api-3-search-get) `lastViewed` field. (optional) (default to false)
	failFast := true // bool | Whether to fail the request quickly in case of an error while loading fields for an issue. For `failFast=true`, if one field fails, the entire operation fails. For `failFast=false`, the operation will continue even if a field fails. It will return a valid response, but without values for the failed field(s). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssue(context.Background(), issueIdOrKey).Fields(fields).FieldsByKeys(fieldsByKeys).Expand(expand).Properties(properties).UpdateHistory(updateHistory).FailFast(failFast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssue`: IssueBean
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fields** | **[]string** | A list of fields to return for the issue. This parameter accepts a comma-separated list. Use it to retrieve a subset of fields. Allowed values:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  Examples:   *  &#x60;summary,comment&#x60; Returns only the summary and comments fields.  *  &#x60;-description&#x60; Returns all (default) fields except description.  *  &#x60;*navigable,-comment&#x60; Returns all navigable fields except comment.  This parameter may be specified multiple times. For example, &#x60;fields&#x3D;field1,field2&amp; fields&#x3D;field3&#x60;.  Note: All fields are returned by default. This differs from [Search for issues using JQL (GET)](#api-rest-api-3-search-get) and [Search for issues using JQL (POST)](#api-rest-api-3-search-post) where the default is all navigable fields. | 
 **fieldsByKeys** | **bool** | Whether fields in &#x60;fields&#x60; are referenced by keys rather than IDs. This parameter is useful where fields have been added by a connect app and a field&#39;s key may differ from its ID. | [default to false]
 **expand** | **string** | Use [expand](#expansion) to include additional information about the issues in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  &#x60;versionedRepresentations&#x60; Returns a JSON array for each version of a field&#39;s value, with the highest number representing the most recent version. Note: When included in the request, the &#x60;fields&#x60; parameter is ignored. | 
 **properties** | **[]string** | A list of issue properties to return for the issue. This parameter accepts a comma-separated list. Allowed values:   *  &#x60;*all&#x60; Returns all issue properties.  *  Any issue property key, prefixed with a minus to exclude.  Examples:   *  &#x60;*all&#x60; Returns all properties.  *  &#x60;*all,-prop1&#x60; Returns all properties except &#x60;prop1&#x60;.  *  &#x60;prop1,prop2&#x60; Returns &#x60;prop1&#x60; and &#x60;prop2&#x60; properties.  This parameter may be specified multiple times. For example, &#x60;properties&#x3D;prop1,prop2&amp; properties&#x3D;prop3&#x60;. | 
 **updateHistory** | **bool** | Whether the project in which the issue is created is added to the user&#39;s **Recently viewed** project list, as shown under **Projects** in Jira. This also populates the [JQL issues search](#api-rest-api-3-search-get) &#x60;lastViewed&#x60; field. | [default to false]
 **failFast** | **bool** | Whether to fail the request quickly in case of an error while loading fields for an issue. For &#x60;failFast&#x3D;true&#x60;, if one field fails, the entire operation fails. For &#x60;failFast&#x3D;false&#x60;, the operation will continue even if a field fails. It will return a valid response, but without values for the failed field(s). | [default to false]

### Return type

[**IssueBean**](IssueBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueLimitReport

> IssueLimitReportResponseBean GetIssueLimitReport(ctx).IsReturningKeys(isReturningKeys).Execute()

Get issue limit report



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
	isReturningKeys := true // bool | Return issue keys instead of issue ids in the response.  Usage: Add `?isReturningKeys=true` to the end of the path to request issue keys. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetIssueLimitReport(context.Background()).IsReturningKeys(isReturningKeys).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetIssueLimitReport``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueLimitReport`: IssueLimitReportResponseBean
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetIssueLimitReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueLimitReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **isReturningKeys** | **bool** | Return issue keys instead of issue ids in the response.  Usage: Add &#x60;?isReturningKeys&#x3D;true&#x60; to the end of the path to request issue keys. | [default to false]

### Return type

[**IssueLimitReportResponseBean**](IssueLimitReportResponseBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTransitions

> Transitions GetTransitions(ctx, issueIdOrKey).Expand(expand).TransitionId(transitionId).SkipRemoteOnlyCondition(skipRemoteOnlyCondition).IncludeUnavailableTransitions(includeUnavailableTransitions).SortByOpsBarAndStatus(sortByOpsBarAndStatus).Execute()

Get transitions



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
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about transitions in the response. This parameter accepts `transitions.fields`, which returns information about the fields in the transition screen for each transition. Fields hidden from the screen are not returned. Use this information to populate the `fields` and `update` fields in [Transition issue](#api-rest-api-3-issue-issueIdOrKey-transitions-post). (optional)
	transitionId := "transitionId_example" // string | The ID of the transition. (optional)
	skipRemoteOnlyCondition := true // bool | Whether transitions with the condition *Hide From User Condition* are included in the response. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)
	includeUnavailableTransitions := true // bool | Whether details of transitions that fail a condition are included in the response (optional) (default to false)
	sortByOpsBarAndStatus := true // bool | Whether the transitions are sorted by ops-bar sequence value first then category order (Todo, In Progress, Done) or only by ops-bar sequence value. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.GetTransitions(context.Background(), issueIdOrKey).Expand(expand).TransitionId(transitionId).SkipRemoteOnlyCondition(skipRemoteOnlyCondition).IncludeUnavailableTransitions(includeUnavailableTransitions).SortByOpsBarAndStatus(sortByOpsBarAndStatus).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.GetTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTransitions`: Transitions
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.GetTransitions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use [expand](#expansion) to include additional information about transitions in the response. This parameter accepts &#x60;transitions.fields&#x60;, which returns information about the fields in the transition screen for each transition. Fields hidden from the screen are not returned. Use this information to populate the &#x60;fields&#x60; and &#x60;update&#x60; fields in [Transition issue](#api-rest-api-3-issue-issueIdOrKey-transitions-post). | 
 **transitionId** | **string** | The ID of the transition. | 
 **skipRemoteOnlyCondition** | **bool** | Whether transitions with the condition *Hide From User Condition* are included in the response. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]
 **includeUnavailableTransitions** | **bool** | Whether details of transitions that fail a condition are included in the response | [default to false]
 **sortByOpsBarAndStatus** | **bool** | Whether the transitions are sorted by ops-bar sequence value first then category order (Todo, In Progress, Done) or only by ops-bar sequence value. | [default to false]

### Return type

[**Transitions**](Transitions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Notify

> interface{} Notify(ctx, issueIdOrKey).Notification(notification).Execute()

Send notification for issue



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
	issueIdOrKey := "issueIdOrKey_example" // string | ID or key of the issue that the notification is sent for.
	notification := *openapiclient.NewNotification() // Notification | The request object for the notification and recipients.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.Notify(context.Background(), issueIdOrKey).Notification(notification).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.Notify``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Notify`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.Notify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | ID or key of the issue that the notification is sent for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiNotifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **notification** | [**Notification**](Notification.md) | The request object for the notification and recipients. | 

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


## UnarchiveIssues

> IssueArchivalSyncResponse UnarchiveIssues(ctx).IssueArchivalSyncRequest(issueArchivalSyncRequest).Execute()

Unarchive issue(s) by issue keys/ID



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
	issueArchivalSyncRequest := *openapiclient.NewIssueArchivalSyncRequest() // IssueArchivalSyncRequest | Contains a list of issue keys or IDs to be unarchived.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuesAPI.UnarchiveIssues(context.Background()).IssueArchivalSyncRequest(issueArchivalSyncRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuesAPI.UnarchiveIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UnarchiveIssues`: IssueArchivalSyncResponse
	fmt.Fprintf(os.Stdout, "Response from `IssuesAPI.UnarchiveIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUnarchiveIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueArchivalSyncRequest** | [**IssueArchivalSyncRequest**](IssueArchivalSyncRequest.md) | Contains a list of issue keys or IDs to be unarchived. | 

### Return type

[**IssueArchivalSyncResponse**](IssueArchivalSyncResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

