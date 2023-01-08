# \ScreenSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateScreenScheme**](ScreenSchemesAPI.md#CreateScreenScheme) | **Post** /rest/api/3/screenscheme | Create screen scheme
[**DeleteScreenScheme**](ScreenSchemesAPI.md#DeleteScreenScheme) | **Delete** /rest/api/3/screenscheme/{screenSchemeId} | Delete screen scheme
[**GetScreenSchemes**](ScreenSchemesAPI.md#GetScreenSchemes) | **Get** /rest/api/3/screenscheme | Get screen schemes
[**UpdateScreenScheme**](ScreenSchemesAPI.md#UpdateScreenScheme) | **Put** /rest/api/3/screenscheme/{screenSchemeId} | Update screen scheme



## CreateScreenScheme

> ScreenSchemeId CreateScreenScheme(ctx).ScreenSchemeDetails(screenSchemeDetails).Execute()

Create screen scheme



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
	screenSchemeDetails := *openapiclient.NewScreenSchemeDetails("Name_example", *openapiclient.NewScreenTypes(int64(123))) // ScreenSchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenSchemesAPI.CreateScreenScheme(context.Background()).ScreenSchemeDetails(screenSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenSchemesAPI.CreateScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScreenScheme`: ScreenSchemeId
	fmt.Fprintf(os.Stdout, "Response from `ScreenSchemesAPI.CreateScreenScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **screenSchemeDetails** | [**ScreenSchemeDetails**](ScreenSchemeDetails.md) |  | 

### Return type

[**ScreenSchemeId**](ScreenSchemeId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScreenScheme

> DeleteScreenScheme(ctx, screenSchemeId).Execute()

Delete screen scheme



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
	screenSchemeId := "screenSchemeId_example" // string | The ID of the screen scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScreenSchemesAPI.DeleteScreenScheme(context.Background(), screenSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenSchemesAPI.DeleteScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenSchemeId** | **string** | The ID of the screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetScreenSchemes

> PageBeanScreenScheme GetScreenSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).Expand(expand).QueryString(queryString).OrderBy(orderBy).Execute()

Get screen schemes



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 25)
	id := []int64{int64(123)} // []int64 | The list of screen scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) include additional information in the response. This parameter accepts `issueTypeScreenSchemes` that, for each screen schemes, returns information about the issue type screen scheme the screen scheme is assigned to. (optional) (default to "")
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with screen scheme name. (optional) (default to "")
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `id` Sorts by screen scheme ID.  *  `name` Sorts by screen scheme name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenSchemesAPI.GetScreenSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).Expand(expand).QueryString(queryString).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenSchemesAPI.GetScreenSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScreenSchemes`: PageBeanScreenScheme
	fmt.Fprintf(os.Stdout, "Response from `ScreenSchemesAPI.GetScreenSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScreenSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 25]
 **id** | **[]int64** | The list of screen scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) include additional information in the response. This parameter accepts &#x60;issueTypeScreenSchemes&#x60; that, for each screen schemes, returns information about the issue type screen scheme the screen scheme is assigned to. | [default to &quot;&quot;]
 **queryString** | **string** | String used to perform a case-insensitive partial match with screen scheme name. | [default to &quot;&quot;]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;id&#x60; Sorts by screen scheme ID.  *  &#x60;name&#x60; Sorts by screen scheme name. | 

### Return type

[**PageBeanScreenScheme**](PageBeanScreenScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScreenScheme

> interface{} UpdateScreenScheme(ctx, screenSchemeId).UpdateScreenSchemeDetails(updateScreenSchemeDetails).Execute()

Update screen scheme



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
	screenSchemeId := "screenSchemeId_example" // string | The ID of the screen scheme.
	updateScreenSchemeDetails := *openapiclient.NewUpdateScreenSchemeDetails() // UpdateScreenSchemeDetails | The screen scheme update details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenSchemesAPI.UpdateScreenScheme(context.Background(), screenSchemeId).UpdateScreenSchemeDetails(updateScreenSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenSchemesAPI.UpdateScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScreenSchemesAPI.UpdateScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenSchemeId** | **string** | The ID of the screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScreenSchemeDetails** | [**UpdateScreenSchemeDetails**](UpdateScreenSchemeDetails.md) | The screen scheme update details. | 

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

