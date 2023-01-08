# \StatusAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStatuses**](StatusAPI.md#CreateStatuses) | **Post** /rest/api/3/statuses | Bulk create statuses
[**DeleteStatusesById**](StatusAPI.md#DeleteStatusesById) | **Delete** /rest/api/3/statuses | Bulk delete Statuses
[**GetProjectIssueTypeUsagesForStatus**](StatusAPI.md#GetProjectIssueTypeUsagesForStatus) | **Get** /rest/api/3/statuses/{statusId}/project/{projectId}/issueTypeUsages | Get issue type usages by status and project
[**GetProjectUsagesForStatus**](StatusAPI.md#GetProjectUsagesForStatus) | **Get** /rest/api/3/statuses/{statusId}/projectUsages | Get project usages by status
[**GetStatusesById**](StatusAPI.md#GetStatusesById) | **Get** /rest/api/3/statuses | Bulk get statuses
[**GetStatusesByName**](StatusAPI.md#GetStatusesByName) | **Get** /rest/api/3/statuses/byNames | Bulk get statuses by name
[**GetWorkflowUsagesForStatus**](StatusAPI.md#GetWorkflowUsagesForStatus) | **Get** /rest/api/3/statuses/{statusId}/workflowUsages | Get workflow usages by status
[**Search**](StatusAPI.md#Search) | **Get** /rest/api/3/statuses/search | Search statuses paginated
[**UpdateStatuses**](StatusAPI.md#UpdateStatuses) | **Put** /rest/api/3/statuses | Bulk update statuses



## CreateStatuses

> []JiraStatus CreateStatuses(ctx).StatusCreateRequest(statusCreateRequest).Execute()

Bulk create statuses



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
	statusCreateRequest := *openapiclient.NewStatusCreateRequest(*openapiclient.NewStatusScope("Type_example"), []openapiclient.StatusCreate{*openapiclient.NewStatusCreate("Name_example", "StatusCategory_example")}) // StatusCreateRequest | Details of the statuses being created and their scope.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.CreateStatuses(context.Background()).StatusCreateRequest(statusCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.CreateStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStatuses`: []JiraStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.CreateStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusCreateRequest** | [**StatusCreateRequest**](StatusCreateRequest.md) | Details of the statuses being created and their scope. | 

### Return type

[**[]JiraStatus**](JiraStatus.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteStatusesById

> interface{} DeleteStatusesById(ctx).Id(id).Execute()

Bulk delete Statuses



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
	id := []string{"Inner_example"} // []string | The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id=10000&id=10001.  Min items `1`, Max items `50`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.DeleteStatusesById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.DeleteStatusesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStatusesById`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.DeleteStatusesById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStatusesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id&#x3D;10000&amp;id&#x3D;10001.  Min items &#x60;1&#x60;, Max items &#x60;50&#x60; | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectIssueTypeUsagesForStatus

> StatusProjectIssueTypeUsageDTO GetProjectIssueTypeUsagesForStatus(ctx, statusId, projectId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get issue type usages by status and project



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
	statusId := "statusId_example" // string | The statusId to fetch issue type usages for
	projectId := "projectId_example" // string | The projectId to fetch issue type usages for
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetProjectIssueTypeUsagesForStatus(context.Background(), statusId, projectId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetProjectIssueTypeUsagesForStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectIssueTypeUsagesForStatus`: StatusProjectIssueTypeUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetProjectIssueTypeUsagesForStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** | The statusId to fetch issue type usages for | 
**projectId** | **string** | The projectId to fetch issue type usages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectIssueTypeUsagesForStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**StatusProjectIssueTypeUsageDTO**](StatusProjectIssueTypeUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectUsagesForStatus

> StatusProjectUsageDTO GetProjectUsagesForStatus(ctx, statusId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get project usages by status



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
	statusId := "statusId_example" // string | The statusId to fetch project usages for
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetProjectUsagesForStatus(context.Background(), statusId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetProjectUsagesForStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectUsagesForStatus`: StatusProjectUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetProjectUsagesForStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** | The statusId to fetch project usages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectUsagesForStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**StatusProjectUsageDTO**](StatusProjectUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusesById

> []JiraStatus GetStatusesById(ctx).Id(id).Execute()

Bulk get statuses



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
	id := []string{"Inner_example"} // []string | The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id=10000&id=10001.  Min items `1`, Max items `50`

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetStatusesById(context.Background()).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetStatusesById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusesById`: []JiraStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetStatusesById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusesByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **id** | **[]string** | The list of status IDs. To include multiple IDs, provide an ampersand-separated list. For example, id&#x3D;10000&amp;id&#x3D;10001.  Min items &#x60;1&#x60;, Max items &#x60;50&#x60; | 

### Return type

[**[]JiraStatus**](JiraStatus.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusesByName

> []JiraStatus GetStatusesByName(ctx).Name(name).ProjectId(projectId).Execute()

Bulk get statuses by name



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
	name := []string{"Inner_example"} // []string | The list of status names. To include multiple names, provide an ampersand-separated list. For example, name=nameXX&name=nameYY.  Min items `1`, Max items `50`
	projectId := "projectId_example" // string | The project the status is part of or null for global statuses. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetStatusesByName(context.Background()).Name(name).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetStatusesByName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusesByName`: []JiraStatus
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetStatusesByName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusesByNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **[]string** | The list of status names. To include multiple names, provide an ampersand-separated list. For example, name&#x3D;nameXX&amp;name&#x3D;nameYY.  Min items &#x60;1&#x60;, Max items &#x60;50&#x60; | 
 **projectId** | **string** | The project the status is part of or null for global statuses. | 

### Return type

[**[]JiraStatus**](JiraStatus.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowUsagesForStatus

> StatusWorkflowUsageDTO GetWorkflowUsagesForStatus(ctx, statusId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get workflow usages by status



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
	statusId := "statusId_example" // string | The statusId to fetch workflow usages for
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.GetWorkflowUsagesForStatus(context.Background(), statusId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.GetWorkflowUsagesForStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowUsagesForStatus`: StatusWorkflowUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.GetWorkflowUsagesForStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** | The statusId to fetch workflow usages for | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowUsagesForStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**StatusWorkflowUsageDTO**](StatusWorkflowUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Search

> PageOfStatuses Search(ctx).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).SearchString(searchString).StatusCategory(statusCategory).Execute()

Search statuses paginated



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
	projectId := "projectId_example" // string | The project the status is part of or null for global statuses. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 200)
	searchString := "searchString_example" // string | Term to match status names against or null to search for all statuses in the search scope. (optional)
	statusCategory := "statusCategory_example" // string | Category of the status to filter by. The supported values are: `TODO`, `IN_PROGRESS`, and `DONE`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.Search(context.Background()).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).SearchString(searchString).StatusCategory(statusCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.Search``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Search`: PageOfStatuses
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.Search`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **string** | The project the status is part of or null for global statuses. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 200]
 **searchString** | **string** | Term to match status names against or null to search for all statuses in the search scope. | 
 **statusCategory** | **string** | Category of the status to filter by. The supported values are: &#x60;TODO&#x60;, &#x60;IN_PROGRESS&#x60;, and &#x60;DONE&#x60;. | 

### Return type

[**PageOfStatuses**](PageOfStatuses.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateStatuses

> interface{} UpdateStatuses(ctx).StatusUpdateRequest(statusUpdateRequest).Execute()

Bulk update statuses



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
	statusUpdateRequest := *openapiclient.NewStatusUpdateRequest([]openapiclient.StatusUpdate{*openapiclient.NewStatusUpdate("Id_example", "Name_example", "StatusCategory_example")}) // StatusUpdateRequest | The list of statuses that will be updated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StatusAPI.UpdateStatuses(context.Background()).StatusUpdateRequest(statusUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StatusAPI.UpdateStatuses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStatuses`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `StatusAPI.UpdateStatuses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStatusesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **statusUpdateRequest** | [**StatusUpdateRequest**](StatusUpdateRequest.md) | The list of statuses that will be updated. | 

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

