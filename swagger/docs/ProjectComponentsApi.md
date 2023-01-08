# \ProjectComponentsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateComponent**](ProjectComponentsAPI.md#CreateComponent) | **Post** /rest/api/3/component | Create component
[**DeleteComponent**](ProjectComponentsAPI.md#DeleteComponent) | **Delete** /rest/api/3/component/{id} | Delete component
[**FindComponentsForProjects**](ProjectComponentsAPI.md#FindComponentsForProjects) | **Get** /rest/api/3/component | Find components for projects
[**GetComponent**](ProjectComponentsAPI.md#GetComponent) | **Get** /rest/api/3/component/{id} | Get component
[**GetComponentRelatedIssues**](ProjectComponentsAPI.md#GetComponentRelatedIssues) | **Get** /rest/api/3/component/{id}/relatedIssueCounts | Get component issues count
[**GetProjectComponents**](ProjectComponentsAPI.md#GetProjectComponents) | **Get** /rest/api/3/project/{projectIdOrKey}/components | Get project components
[**GetProjectComponentsPaginated**](ProjectComponentsAPI.md#GetProjectComponentsPaginated) | **Get** /rest/api/3/project/{projectIdOrKey}/component | Get project components paginated
[**UpdateComponent**](ProjectComponentsAPI.md#UpdateComponent) | **Put** /rest/api/3/component/{id} | Update component



## CreateComponent

> ProjectComponent CreateComponent(ctx).ProjectComponent(projectComponent).Execute()

Create component



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
	projectComponent := *openapiclient.NewProjectComponent() // ProjectComponent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.CreateComponent(context.Background()).ProjectComponent(projectComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.CreateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateComponent`: ProjectComponent
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.CreateComponent`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectComponent** | [**ProjectComponent**](ProjectComponent.md) |  | 

### Return type

[**ProjectComponent**](ProjectComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComponent

> DeleteComponent(ctx, id).MoveIssuesTo(moveIssuesTo).Execute()

Delete component



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
	id := "id_example" // string | The ID of the component.
	moveIssuesTo := "moveIssuesTo_example" // string | The ID of the component to replace the deleted component. If this value is null no replacement is made. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectComponentsAPI.DeleteComponent(context.Background(), id).MoveIssuesTo(moveIssuesTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.DeleteComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveIssuesTo** | **string** | The ID of the component to replace the deleted component. If this value is null no replacement is made. | 

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


## FindComponentsForProjects

> PageBean2ComponentJsonBean FindComponentsForProjects(ctx).ProjectIdsOrKeys(projectIdsOrKeys).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Query(query).Execute()

Find components for projects



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
	projectIdsOrKeys := []string{"Inner_example"} // []string | The project IDs and/or project keys (case sensitive). (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `description` Sorts by the component description.  *  `name` Sorts by component name. (optional)
	query := "query_example" // string | Filter the results using a literal string. Components with a matching `name` or `description` are returned (case insensitive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.FindComponentsForProjects(context.Background()).ProjectIdsOrKeys(projectIdsOrKeys).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.FindComponentsForProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindComponentsForProjects`: PageBean2ComponentJsonBean
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.FindComponentsForProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindComponentsForProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectIdsOrKeys** | **[]string** | The project IDs and/or project keys (case sensitive). | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;description&#x60; Sorts by the component description.  *  &#x60;name&#x60; Sorts by component name. | 
 **query** | **string** | Filter the results using a literal string. Components with a matching &#x60;name&#x60; or &#x60;description&#x60; are returned (case insensitive). | 

### Return type

[**PageBean2ComponentJsonBean**](PageBean2ComponentJsonBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponent

> ProjectComponent GetComponent(ctx, id).Execute()

Get component



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
	id := "id_example" // string | The ID of the component.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.GetComponent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.GetComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponent`: ProjectComponent
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.GetComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectComponent**](ProjectComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentRelatedIssues

> ComponentIssuesCount GetComponentRelatedIssues(ctx, id).Execute()

Get component issues count



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
	id := "id_example" // string | The ID of the component.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.GetComponentRelatedIssues(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.GetComponentRelatedIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentRelatedIssues`: ComponentIssuesCount
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.GetComponentRelatedIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentRelatedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ComponentIssuesCount**](ComponentIssuesCount.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectComponents

> []ProjectComponent GetProjectComponents(ctx, projectIdOrKey).ComponentSource(componentSource).Execute()

Get project components



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	componentSource := "componentSource_example" // string | The source of the components to return. Can be `jira` (default), `compass` or `auto`. When `auto` is specified, the API will return connected Compass components if the project is opted into Compass, otherwise it will return Jira components. Defaults to `jira`. (optional) (default to "jira")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.GetProjectComponents(context.Background(), projectIdOrKey).ComponentSource(componentSource).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.GetProjectComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectComponents`: []ProjectComponent
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.GetProjectComponents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **componentSource** | **string** | The source of the components to return. Can be &#x60;jira&#x60; (default), &#x60;compass&#x60; or &#x60;auto&#x60;. When &#x60;auto&#x60; is specified, the API will return connected Compass components if the project is opted into Compass, otherwise it will return Jira components. Defaults to &#x60;jira&#x60;. | [default to &quot;jira&quot;]

### Return type

[**[]ProjectComponent**](ProjectComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectComponentsPaginated

> PageBeanComponentWithIssueCount GetProjectComponentsPaginated(ctx, projectIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).ComponentSource(componentSource).Query(query).Execute()

Get project components paginated



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `description` Sorts by the component description.  *  `issueCount` Sorts by the count of issues associated with the component.  *  `lead` Sorts by the user key of the component's project lead.  *  `name` Sorts by component name. (optional)
	componentSource := "componentSource_example" // string | The source of the components to return. Can be `jira` (default), `compass` or `auto`. When `auto` is specified, the API will return connected Compass components if the project is opted into Compass, otherwise it will return Jira components. Defaults to `jira`. (optional) (default to "jira")
	query := "query_example" // string | Filter the results using a literal string. Components with a matching `name` or `description` are returned (case insensitive). (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.GetProjectComponentsPaginated(context.Background(), projectIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).ComponentSource(componentSource).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.GetProjectComponentsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectComponentsPaginated`: PageBeanComponentWithIssueCount
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.GetProjectComponentsPaginated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectComponentsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;description&#x60; Sorts by the component description.  *  &#x60;issueCount&#x60; Sorts by the count of issues associated with the component.  *  &#x60;lead&#x60; Sorts by the user key of the component&#39;s project lead.  *  &#x60;name&#x60; Sorts by component name. | 
 **componentSource** | **string** | The source of the components to return. Can be &#x60;jira&#x60; (default), &#x60;compass&#x60; or &#x60;auto&#x60;. When &#x60;auto&#x60; is specified, the API will return connected Compass components if the project is opted into Compass, otherwise it will return Jira components. Defaults to &#x60;jira&#x60;. | [default to &quot;jira&quot;]
 **query** | **string** | Filter the results using a literal string. Components with a matching &#x60;name&#x60; or &#x60;description&#x60; are returned (case insensitive). | 

### Return type

[**PageBeanComponentWithIssueCount**](PageBeanComponentWithIssueCount.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComponent

> ProjectComponent UpdateComponent(ctx, id).ProjectComponent(projectComponent).Execute()

Update component



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
	id := "id_example" // string | The ID of the component.
	projectComponent := *openapiclient.NewProjectComponent() // ProjectComponent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectComponentsAPI.UpdateComponent(context.Background(), id).ProjectComponent(projectComponent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectComponentsAPI.UpdateComponent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComponent`: ProjectComponent
	fmt.Fprintf(os.Stdout, "Response from `ProjectComponentsAPI.UpdateComponent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the component. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateComponentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectComponent** | [**ProjectComponent**](ProjectComponent.md) |  | 

### Return type

[**ProjectComponent**](ProjectComponent.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

