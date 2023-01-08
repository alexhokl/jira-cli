# \ScreensAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddFieldToDefaultScreen**](ScreensAPI.md#AddFieldToDefaultScreen) | **Post** /rest/api/3/screens/addToDefault/{fieldId} | Add field to default screen
[**CreateScreen**](ScreensAPI.md#CreateScreen) | **Post** /rest/api/3/screens | Create screen
[**DeleteScreen**](ScreensAPI.md#DeleteScreen) | **Delete** /rest/api/3/screens/{screenId} | Delete screen
[**GetAvailableScreenFields**](ScreensAPI.md#GetAvailableScreenFields) | **Get** /rest/api/3/screens/{screenId}/availableFields | Get available screen fields
[**GetScreens**](ScreensAPI.md#GetScreens) | **Get** /rest/api/3/screens | Get screens
[**GetScreensForField**](ScreensAPI.md#GetScreensForField) | **Get** /rest/api/3/field/{fieldId}/screens | Get screens for a field
[**UpdateScreen**](ScreensAPI.md#UpdateScreen) | **Put** /rest/api/3/screens/{screenId} | Update screen



## AddFieldToDefaultScreen

> interface{} AddFieldToDefaultScreen(ctx, fieldId).Execute()

Add field to default screen



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
	fieldId := "fieldId_example" // string | The ID of the field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.AddFieldToDefaultScreen(context.Background(), fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.AddFieldToDefaultScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddFieldToDefaultScreen`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.AddFieldToDefaultScreen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddFieldToDefaultScreenRequest struct via the builder pattern


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


## CreateScreen

> Screen CreateScreen(ctx).ScreenDetails(screenDetails).Execute()

Create screen



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
	screenDetails := *openapiclient.NewScreenDetails("Name_example") // ScreenDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.CreateScreen(context.Background()).ScreenDetails(screenDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.CreateScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateScreen`: Screen
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.CreateScreen`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateScreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **screenDetails** | [**ScreenDetails**](ScreenDetails.md) |  | 

### Return type

[**Screen**](Screen.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteScreen

> DeleteScreen(ctx, screenId).Execute()

Delete screen



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScreensAPI.DeleteScreen(context.Background(), screenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.DeleteScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteScreenRequest struct via the builder pattern


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


## GetAvailableScreenFields

> []ScreenableField GetAvailableScreenFields(ctx, screenId).Execute()

Get available screen fields



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.GetAvailableScreenFields(context.Background(), screenId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.GetAvailableScreenFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableScreenFields`: []ScreenableField
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.GetAvailableScreenFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableScreenFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]ScreenableField**](ScreenableField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScreens

> PageBeanScreen GetScreens(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).QueryString(queryString).Scope(scope).OrderBy(orderBy).Execute()

Get screens



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
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)
	id := []int64{int64(123)} // []int64 | The list of screen IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with screen name. (optional) (default to "")
	scope := []string{"Scope_example"} // []string | The scope filter string. To filter by multiple scope, provide an ampersand-separated list. For example, `scope=GLOBAL&scope=PROJECT`. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `id` Sorts by screen ID.  *  `name` Sorts by screen name. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.GetScreens(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).QueryString(queryString).Scope(scope).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.GetScreens``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScreens`: PageBeanScreen
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.GetScreens`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetScreensRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]
 **id** | **[]int64** | The list of screen IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **queryString** | **string** | String used to perform a case-insensitive partial match with screen name. | [default to &quot;&quot;]
 **scope** | **[]string** | The scope filter string. To filter by multiple scope, provide an ampersand-separated list. For example, &#x60;scope&#x3D;GLOBAL&amp;scope&#x3D;PROJECT&#x60;. | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;id&#x60; Sorts by screen ID.  *  &#x60;name&#x60; Sorts by screen name. | 

### Return type

[**PageBeanScreen**](PageBeanScreen.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetScreensForField

> PageBeanScreenWithTab GetScreensForField(ctx, fieldId).StartAt(startAt).MaxResults(maxResults).Expand(expand).Execute()

Get screens for a field



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
	fieldId := "fieldId_example" // string | The ID of the field to return screens for.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about screens in the response. This parameter accepts `tab` which returns details about the screen tabs the field is used in. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.GetScreensForField(context.Background(), fieldId).StartAt(startAt).MaxResults(maxResults).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.GetScreensForField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetScreensForField`: PageBeanScreenWithTab
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.GetScreensForField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the field to return screens for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetScreensForFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]
 **expand** | **string** | Use [expand](#expansion) to include additional information about screens in the response. This parameter accepts &#x60;tab&#x60; which returns details about the screen tabs the field is used in. | 

### Return type

[**PageBeanScreenWithTab**](PageBeanScreenWithTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateScreen

> Screen UpdateScreen(ctx, screenId).UpdateScreenDetails(updateScreenDetails).Execute()

Update screen



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
	updateScreenDetails := *openapiclient.NewUpdateScreenDetails() // UpdateScreenDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreensAPI.UpdateScreen(context.Background(), screenId).UpdateScreenDetails(updateScreenDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreensAPI.UpdateScreen``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateScreen`: Screen
	fmt.Fprintf(os.Stdout, "Response from `ScreensAPI.UpdateScreen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateScreenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateScreenDetails** | [**UpdateScreenDetails**](UpdateScreenDetails.md) |  | 

### Return type

[**Screen**](Screen.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

