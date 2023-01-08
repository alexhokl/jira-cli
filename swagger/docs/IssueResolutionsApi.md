# \IssueResolutionsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateResolution**](IssueResolutionsAPI.md#CreateResolution) | **Post** /rest/api/3/resolution | Create resolution
[**DeleteResolution**](IssueResolutionsAPI.md#DeleteResolution) | **Delete** /rest/api/3/resolution/{id} | Delete resolution
[**GetResolution**](IssueResolutionsAPI.md#GetResolution) | **Get** /rest/api/3/resolution/{id} | Get resolution
[**GetResolutions**](IssueResolutionsAPI.md#GetResolutions) | **Get** /rest/api/3/resolution | Get resolutions
[**MoveResolutions**](IssueResolutionsAPI.md#MoveResolutions) | **Put** /rest/api/3/resolution/move | Move resolutions
[**SearchResolutions**](IssueResolutionsAPI.md#SearchResolutions) | **Get** /rest/api/3/resolution/search | Search resolutions
[**SetDefaultResolution**](IssueResolutionsAPI.md#SetDefaultResolution) | **Put** /rest/api/3/resolution/default | Set default resolution
[**UpdateResolution**](IssueResolutionsAPI.md#UpdateResolution) | **Put** /rest/api/3/resolution/{id} | Update resolution



## CreateResolution

> ResolutionId CreateResolution(ctx).CreateResolutionDetails(createResolutionDetails).Execute()

Create resolution



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
	createResolutionDetails := *openapiclient.NewCreateResolutionDetails("Name_example") // CreateResolutionDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.CreateResolution(context.Background()).CreateResolutionDetails(createResolutionDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.CreateResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateResolution`: ResolutionId
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.CreateResolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createResolutionDetails** | [**CreateResolutionDetails**](CreateResolutionDetails.md) |  | 

### Return type

[**ResolutionId**](ResolutionId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResolution

> DeleteResolution(ctx, id).ReplaceWith(replaceWith).Execute()

Delete resolution



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
	id := "id_example" // string | The ID of the issue resolution.
	replaceWith := "replaceWith_example" // string | The ID of the issue resolution that will replace the currently selected resolution. (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueResolutionsAPI.DeleteResolution(context.Background(), id).ReplaceWith(replaceWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.DeleteResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue resolution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **replaceWith** | **string** | The ID of the issue resolution that will replace the currently selected resolution. | [default to &quot;&quot;]

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


## GetResolution

> Resolution GetResolution(ctx, id).Execute()

Get resolution



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
	id := "id_example" // string | The ID of the issue resolution value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.GetResolution(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.GetResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResolution`: Resolution
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.GetResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue resolution value. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Resolution**](Resolution.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResolutions

> []Resolution GetResolutions(ctx).Execute()

Get resolutions



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
	resp, r, err := apiClient.IssueResolutionsAPI.GetResolutions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.GetResolutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetResolutions`: []Resolution
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.GetResolutions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetResolutionsRequest struct via the builder pattern


### Return type

[**[]Resolution**](Resolution.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveResolutions

> interface{} MoveResolutions(ctx).ReorderIssueResolutionsRequest(reorderIssueResolutionsRequest).Execute()

Move resolutions



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
	reorderIssueResolutionsRequest := *openapiclient.NewReorderIssueResolutionsRequest([]string{"Ids_example"}) // ReorderIssueResolutionsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.MoveResolutions(context.Background()).ReorderIssueResolutionsRequest(reorderIssueResolutionsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.MoveResolutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveResolutions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.MoveResolutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMoveResolutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reorderIssueResolutionsRequest** | [**ReorderIssueResolutionsRequest**](ReorderIssueResolutionsRequest.md) |  | 

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


## SearchResolutions

> PageBeanResolutionJsonBean SearchResolutions(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).OnlyDefault(onlyDefault).Execute()

Search resolutions



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	id := []string{"Inner_example"} // []string | The list of resolutions IDs to be filtered out (optional)
	onlyDefault := true // bool | When set to true, return default only, when IDs provided, if none of them is default, return empty page. Default value is false (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.SearchResolutions(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).OnlyDefault(onlyDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.SearchResolutions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchResolutions`: PageBeanResolutionJsonBean
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.SearchResolutions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchResolutionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of resolutions IDs to be filtered out | 
 **onlyDefault** | **bool** | When set to true, return default only, when IDs provided, if none of them is default, return empty page. Default value is false | [default to false]

### Return type

[**PageBeanResolutionJsonBean**](PageBeanResolutionJsonBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultResolution

> interface{} SetDefaultResolution(ctx).SetDefaultResolutionRequest(setDefaultResolutionRequest).Execute()

Set default resolution



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
	setDefaultResolutionRequest := *openapiclient.NewSetDefaultResolutionRequest("Id_example") // SetDefaultResolutionRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.SetDefaultResolution(context.Background()).SetDefaultResolutionRequest(setDefaultResolutionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.SetDefaultResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultResolution`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.SetDefaultResolution`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultResolutionRequest** | [**SetDefaultResolutionRequest**](SetDefaultResolutionRequest.md) |  | 

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


## UpdateResolution

> interface{} UpdateResolution(ctx, id).UpdateResolutionDetails(updateResolutionDetails).Execute()

Update resolution



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
	id := "id_example" // string | The ID of the issue resolution.
	updateResolutionDetails := *openapiclient.NewUpdateResolutionDetails("Name_example") // UpdateResolutionDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueResolutionsAPI.UpdateResolution(context.Background(), id).UpdateResolutionDetails(updateResolutionDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueResolutionsAPI.UpdateResolution``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateResolution`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueResolutionsAPI.UpdateResolution`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue resolution. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResolutionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateResolutionDetails** | [**UpdateResolutionDetails**](UpdateResolutionDetails.md) |  | 

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

