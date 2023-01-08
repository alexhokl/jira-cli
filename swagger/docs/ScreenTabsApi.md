# \ScreenTabsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScreenTab**](ScreenTabsAPI.md#AddScreenTab) | **Post** /rest/api/3/screens/{screenId}/tabs | Create screen tab
[**DeleteScreenTab**](ScreenTabsAPI.md#DeleteScreenTab) | **Delete** /rest/api/3/screens/{screenId}/tabs/{tabId} | Delete screen tab
[**GetAllScreenTabs**](ScreenTabsAPI.md#GetAllScreenTabs) | **Get** /rest/api/3/screens/{screenId}/tabs | Get all screen tabs
[**GetBulkScreenTabs**](ScreenTabsAPI.md#GetBulkScreenTabs) | **Get** /rest/api/3/screens/tabs | Get bulk screen tabs
[**MoveScreenTab**](ScreenTabsAPI.md#MoveScreenTab) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/move/{pos} | Move screen tab
[**RenameScreenTab**](ScreenTabsAPI.md#RenameScreenTab) | **Put** /rest/api/3/screens/{screenId}/tabs/{tabId} | Update screen tab



## AddScreenTab

> ScreenableTab AddScreenTab(ctx, screenId).ScreenableTab(screenableTab).Execute()

Create screen tab



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
	screenId := int64(789) // int64 | The ID of the screen.
	screenableTab := *openapiclient.NewScreenableTab("Name_example") // ScreenableTab | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabsAPI.AddScreenTab(context.Background(), screenId).ScreenableTab(screenableTab).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.AddScreenTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddScreenTab`: ScreenableTab
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabsAPI.AddScreenTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScreenTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **screenableTab** | [**ScreenableTab**](ScreenableTab.md) |  | 

### Return type

[**ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScreenTab

> DeleteScreenTab(ctx, screenId, tabId).Execute()

Delete screen tab



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScreenTabsAPI.DeleteScreenTab(context.Background(), screenId, tabId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.DeleteScreenTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScreenTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetAllScreenTabs

> []ScreenableTab GetAllScreenTabs(ctx, screenId).ProjectKey(projectKey).Execute()

Get all screen tabs



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
	screenId := int64(789) // int64 | The ID of the screen.
	projectKey := "projectKey_example" // string | The key of the project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabsAPI.GetAllScreenTabs(context.Background(), screenId).ProjectKey(projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.GetAllScreenTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllScreenTabs`: []ScreenableTab
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabsAPI.GetAllScreenTabs`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllScreenTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectKey** | **string** | The key of the project. | 

### Return type

[**[]ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkScreenTabs

> GetBulkScreenTabs(ctx).ScreenId(screenId).TabId(tabId).StartAt(startAt).MaxResult(maxResult).Execute()

Get bulk screen tabs



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
	screenId := []int64{int64(123)} // []int64 | The list of screen IDs. To include multiple screen IDs, provide an ampersand-separated list. For example, `screenId=10000&screenId=10001`. (optional)
	tabId := []int64{int64(123)} // []int64 | The list of tab IDs. To include multiple tab IDs, provide an ampersand-separated list. For example, `tabId=10000&tabId=10001`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResult := int32(56) // int32 | The maximum number of items to return per page. The maximum number is 100, (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScreenTabsAPI.GetBulkScreenTabs(context.Background()).ScreenId(screenId).TabId(tabId).StartAt(startAt).MaxResult(maxResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.GetBulkScreenTabs``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkScreenTabsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **screenId** | **[]int64** | The list of screen IDs. To include multiple screen IDs, provide an ampersand-separated list. For example, &#x60;screenId&#x3D;10000&amp;screenId&#x3D;10001&#x60;. | 
 **tabId** | **[]int64** | The list of tab IDs. To include multiple tab IDs, provide an ampersand-separated list. For example, &#x60;tabId&#x3D;10000&amp;tabId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResult** | **int32** | The maximum number of items to return per page. The maximum number is 100, | [default to 100]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveScreenTab

> interface{} MoveScreenTab(ctx, screenId, tabId, pos).Execute()

Move screen tab



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	pos := int32(56) // int32 | The position of tab. The base index is 0.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabsAPI.MoveScreenTab(context.Background(), screenId, tabId, pos).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.MoveScreenTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveScreenTab`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabsAPI.MoveScreenTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 
**pos** | **int32** | The position of tab. The base index is 0. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveScreenTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




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


## RenameScreenTab

> ScreenableTab RenameScreenTab(ctx, screenId, tabId).ScreenableTab(screenableTab).Execute()

Update screen tab



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	screenableTab := *openapiclient.NewScreenableTab("Name_example") // ScreenableTab | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabsAPI.RenameScreenTab(context.Background(), screenId, tabId).ScreenableTab(screenableTab).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabsAPI.RenameScreenTab``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RenameScreenTab`: ScreenableTab
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabsAPI.RenameScreenTab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRenameScreenTabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **screenableTab** | [**ScreenableTab**](ScreenableTab.md) |  | 

### Return type

[**ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

