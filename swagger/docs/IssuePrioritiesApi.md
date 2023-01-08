# \IssuePrioritiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriority**](IssuePrioritiesAPI.md#CreatePriority) | **Post** /rest/api/3/priority | Create priority
[**DeletePriority**](IssuePrioritiesAPI.md#DeletePriority) | **Delete** /rest/api/3/priority/{id} | Delete priority
[**GetPriorities**](IssuePrioritiesAPI.md#GetPriorities) | **Get** /rest/api/3/priority | Get priorities
[**GetPriority**](IssuePrioritiesAPI.md#GetPriority) | **Get** /rest/api/3/priority/{id} | Get priority
[**MovePriorities**](IssuePrioritiesAPI.md#MovePriorities) | **Put** /rest/api/3/priority/move | Move priorities
[**SearchPriorities**](IssuePrioritiesAPI.md#SearchPriorities) | **Get** /rest/api/3/priority/search | Search priorities
[**SetDefaultPriority**](IssuePrioritiesAPI.md#SetDefaultPriority) | **Put** /rest/api/3/priority/default | Set default priority
[**UpdatePriority**](IssuePrioritiesAPI.md#UpdatePriority) | **Put** /rest/api/3/priority/{id} | Update priority



## CreatePriority

> PriorityId CreatePriority(ctx).CreatePriorityDetails(createPriorityDetails).Execute()

Create priority



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
	createPriorityDetails := *openapiclient.NewCreatePriorityDetails("Name_example", "StatusColor_example") // CreatePriorityDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.CreatePriority(context.Background()).CreatePriorityDetails(createPriorityDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.CreatePriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePriority`: PriorityId
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.CreatePriority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPriorityDetails** | [**CreatePriorityDetails**](CreatePriorityDetails.md) |  | 

### Return type

[**PriorityId**](PriorityId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriority

> DeletePriority(ctx, id).Execute()

Delete priority



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
	id := "id_example" // string | The ID of the issue priority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePrioritiesAPI.DeletePriority(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.DeletePriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue priority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriorityRequest struct via the builder pattern


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


## GetPriorities

> []Priority GetPriorities(ctx).Execute()

Get priorities



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
	resp, r, err := apiClient.IssuePrioritiesAPI.GetPriorities(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.GetPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriorities`: []Priority
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.GetPriorities`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritiesRequest struct via the builder pattern


### Return type

[**[]Priority**](Priority.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPriority

> Priority GetPriority(ctx, id).Execute()

Get priority



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
	id := "id_example" // string | The ID of the issue priority.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.GetPriority(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.GetPriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPriority`: Priority
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.GetPriority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue priority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Priority**](Priority.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MovePriorities

> interface{} MovePriorities(ctx).ReorderIssuePriorities(reorderIssuePriorities).Execute()

Move priorities



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
	reorderIssuePriorities := *openapiclient.NewReorderIssuePriorities([]string{"Ids_example"}) // ReorderIssuePriorities | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.MovePriorities(context.Background()).ReorderIssuePriorities(reorderIssuePriorities).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.MovePriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MovePriorities`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.MovePriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMovePrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **reorderIssuePriorities** | [**ReorderIssuePriorities**](ReorderIssuePriorities.md) |  | 

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


## SearchPriorities

> PageBeanPriority SearchPriorities(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).PriorityName(priorityName).OnlyDefault(onlyDefault).Expand(expand).Execute()

Search priorities



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
	id := []string{"Inner_example"} // []string | The list of priority IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=2&id=3`. (optional)
	projectId := []string{"Inner_example"} // []string | The list of projects IDs. To include multiple IDs, provide an ampersand-separated list. For example, `projectId=10010&projectId=10111`. (optional)
	priorityName := "priorityName_example" // string | The name of priority to search for. (optional) (default to "")
	onlyDefault := true // bool | Whether only the default priority is returned. (optional) (default to false)
	expand := "expand_example" // string | Use `schemes` to return the associated priority schemes for each priority. Limited to returning first 15 priority schemes per priority. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.SearchPriorities(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).PriorityName(priorityName).OnlyDefault(onlyDefault).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.SearchPriorities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchPriorities`: PageBeanPriority
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.SearchPriorities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchPrioritiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of priority IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;2&amp;id&#x3D;3&#x60;. | 
 **projectId** | **[]string** | The list of projects IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;projectId&#x3D;10010&amp;projectId&#x3D;10111&#x60;. | 
 **priorityName** | **string** | The name of priority to search for. | [default to &quot;&quot;]
 **onlyDefault** | **bool** | Whether only the default priority is returned. | [default to false]
 **expand** | **string** | Use &#x60;schemes&#x60; to return the associated priority schemes for each priority. Limited to returning first 15 priority schemes per priority. | [default to &quot;&quot;]

### Return type

[**PageBeanPriority**](PageBeanPriority.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultPriority

> interface{} SetDefaultPriority(ctx).SetDefaultPriorityRequest(setDefaultPriorityRequest).Execute()

Set default priority



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
	setDefaultPriorityRequest := *openapiclient.NewSetDefaultPriorityRequest("Id_example") // SetDefaultPriorityRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.SetDefaultPriority(context.Background()).SetDefaultPriorityRequest(setDefaultPriorityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.SetDefaultPriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultPriority`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.SetDefaultPriority`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultPriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultPriorityRequest** | [**SetDefaultPriorityRequest**](SetDefaultPriorityRequest.md) |  | 

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


## UpdatePriority

> interface{} UpdatePriority(ctx, id).UpdatePriorityDetails(updatePriorityDetails).Execute()

Update priority



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
	id := "id_example" // string | The ID of the issue priority.
	updatePriorityDetails := *openapiclient.NewUpdatePriorityDetails() // UpdatePriorityDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePrioritiesAPI.UpdatePriority(context.Background(), id).UpdatePriorityDetails(updatePriorityDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePrioritiesAPI.UpdatePriority``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriority`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuePrioritiesAPI.UpdatePriority`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue priority. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriorityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePriorityDetails** | [**UpdatePriorityDetails**](UpdatePriorityDetails.md) |  | 

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

