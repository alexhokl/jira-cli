# \PrioritySchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriorityScheme**](PrioritySchemesAPI.md#CreatePriorityScheme) | **Post** /rest/api/3/priorityscheme | Create priority scheme
[**DeletePriorityScheme**](PrioritySchemesAPI.md#DeletePriorityScheme) | **Delete** /rest/api/3/priorityscheme/{schemeId} | Delete priority scheme
[**GetAvailablePrioritiesByPriorityScheme**](PrioritySchemesAPI.md#GetAvailablePrioritiesByPriorityScheme) | **Get** /rest/api/3/priorityscheme/priorities/available | Get available priorities by priority scheme
[**GetPrioritiesByPriorityScheme**](PrioritySchemesAPI.md#GetPrioritiesByPriorityScheme) | **Get** /rest/api/3/priorityscheme/{schemeId}/priorities | Get priorities by priority scheme
[**GetPrioritySchemes**](PrioritySchemesAPI.md#GetPrioritySchemes) | **Get** /rest/api/3/priorityscheme | Get priority schemes
[**GetProjectsByPriorityScheme**](PrioritySchemesAPI.md#GetProjectsByPriorityScheme) | **Get** /rest/api/3/priorityscheme/{schemeId}/projects | Get projects by priority scheme
[**SuggestedPrioritiesForMappings**](PrioritySchemesAPI.md#SuggestedPrioritiesForMappings) | **Post** /rest/api/3/priorityscheme/mappings | Suggested priorities for mappings
[**UpdatePriorityScheme**](PrioritySchemesAPI.md#UpdatePriorityScheme) | **Put** /rest/api/3/priorityscheme/{schemeId} | Update priority scheme



## CreatePriorityScheme

> PrioritySchemeId CreatePriorityScheme(ctx).CreatePrioritySchemeDetails(createPrioritySchemeDetails).Execute()

Create priority scheme



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
	createPrioritySchemeDetails := *openapiclient.NewCreatePrioritySchemeDetails(int64(123), "Name_example", []int64{int64(123)}) // CreatePrioritySchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.CreatePriorityScheme(context.Background()).CreatePrioritySchemeDetails(createPrioritySchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.CreatePriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePriorityScheme`: PrioritySchemeId
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.CreatePriorityScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrioritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPrioritySchemeDetails** | [**CreatePrioritySchemeDetails**](CreatePrioritySchemeDetails.md) |  | 

### Return type

[**PrioritySchemeId**](PrioritySchemeId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePriorityScheme

> interface{} DeletePriorityScheme(ctx, schemeId).Execute()

Delete priority scheme



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
	schemeId := int64(789) // int64 | The priority scheme ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.DeletePriorityScheme(context.Background(), schemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.DeletePriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePriorityScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.DeletePriorityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The priority scheme ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePrioritySchemeRequest struct via the builder pattern


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


## GetAvailablePrioritiesByPriorityScheme

> PageBeanPriorityWithSequence GetAvailablePrioritiesByPriorityScheme(ctx).SchemeId(schemeId).StartAt(startAt).MaxResults(maxResults).Query(query).Exclude(exclude).Execute()

Get available priorities by priority scheme



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
	schemeId := "schemeId_example" // string | The priority scheme ID.
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	query := "query_example" // string | The string to query priorities on by name. (optional) (default to "")
	exclude := []string{"Inner_example"} // []string | A list of priority IDs to exclude from the results. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.GetAvailablePrioritiesByPriorityScheme(context.Background()).SchemeId(schemeId).StartAt(startAt).MaxResults(maxResults).Query(query).Exclude(exclude).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.GetAvailablePrioritiesByPriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailablePrioritiesByPriorityScheme`: PageBeanPriorityWithSequence
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.GetAvailablePrioritiesByPriorityScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailablePrioritiesByPrioritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **schemeId** | **string** | The priority scheme ID. | 
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **query** | **string** | The string to query priorities on by name. | [default to &quot;&quot;]
 **exclude** | **[]string** | A list of priority IDs to exclude from the results. | 

### Return type

[**PageBeanPriorityWithSequence**](PageBeanPriorityWithSequence.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrioritiesByPriorityScheme

> PageBeanPriorityWithSequence GetPrioritiesByPriorityScheme(ctx, schemeId).StartAt(startAt).MaxResults(maxResults).Execute()

Get priorities by priority scheme



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
	schemeId := "schemeId_example" // string | The priority scheme ID.
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.GetPrioritiesByPriorityScheme(context.Background(), schemeId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.GetPrioritiesByPriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrioritiesByPriorityScheme`: PageBeanPriorityWithSequence
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.GetPrioritiesByPriorityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The priority scheme ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritiesByPrioritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]

### Return type

[**PageBeanPriorityWithSequence**](PageBeanPriorityWithSequence.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrioritySchemes

> PageBeanPrioritySchemeWithPaginatedPrioritiesAndProjects GetPrioritySchemes(ctx).StartAt(startAt).MaxResults(maxResults).PriorityId(priorityId).SchemeId(schemeId).SchemeName(schemeName).OnlyDefault(onlyDefault).OrderBy(orderBy).Expand(expand).Execute()

Get priority schemes



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
	priorityId := []int64{int64(123)} // []int64 | A set of priority IDs to filter by. To include multiple IDs, provide an ampersand-separated list. For example, `priorityId=10000&priorityId=10001`. (optional)
	schemeId := []int64{int64(123)} // []int64 | A set of priority scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `schemeId=10000&schemeId=10001`. (optional)
	schemeName := "schemeName_example" // string | The name of scheme to search for. (optional) (default to "")
	onlyDefault := true // bool | Whether only the default priority is returned. (optional) (default to false)
	orderBy := "orderBy_example" // string | The ordering to return the priority schemes by. (optional) (default to "+name")
	expand := "expand_example" // string | A comma separated list of additional information to return. \"priorities\" will return priorities associated with the priority scheme. \"projects\" will return projects associated with the priority scheme. `expand=priorities,projects`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.GetPrioritySchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).PriorityId(priorityId).SchemeId(schemeId).SchemeName(schemeName).OnlyDefault(onlyDefault).OrderBy(orderBy).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.GetPrioritySchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrioritySchemes`: PageBeanPrioritySchemeWithPaginatedPrioritiesAndProjects
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.GetPrioritySchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrioritySchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **priorityId** | **[]int64** | A set of priority IDs to filter by. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;priorityId&#x3D;10000&amp;priorityId&#x3D;10001&#x60;. | 
 **schemeId** | **[]int64** | A set of priority scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;schemeId&#x3D;10000&amp;schemeId&#x3D;10001&#x60;. | 
 **schemeName** | **string** | The name of scheme to search for. | [default to &quot;&quot;]
 **onlyDefault** | **bool** | Whether only the default priority is returned. | [default to false]
 **orderBy** | **string** | The ordering to return the priority schemes by. | [default to &quot;+name&quot;]
 **expand** | **string** | A comma separated list of additional information to return. \&quot;priorities\&quot; will return priorities associated with the priority scheme. \&quot;projects\&quot; will return projects associated with the priority scheme. &#x60;expand&#x3D;priorities,projects&#x60;. | 

### Return type

[**PageBeanPrioritySchemeWithPaginatedPrioritiesAndProjects**](PageBeanPrioritySchemeWithPaginatedPrioritiesAndProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsByPriorityScheme

> PageBeanProject GetProjectsByPriorityScheme(ctx, schemeId).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Query(query).Execute()

Get projects by priority scheme



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
	schemeId := "schemeId_example" // string | The priority scheme ID.
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	projectId := []int64{int64(123)} // []int64 | The project IDs to filter by. For example, `projectId=10000&projectId=10001`. (optional)
	query := "query_example" // string | The string to query projects on by name. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.GetProjectsByPriorityScheme(context.Background(), schemeId).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.GetProjectsByPriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsByPriorityScheme`: PageBeanProject
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.GetProjectsByPriorityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The priority scheme ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsByPrioritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **projectId** | **[]int64** | The project IDs to filter by. For example, &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 
 **query** | **string** | The string to query projects on by name. | [default to &quot;&quot;]

### Return type

[**PageBeanProject**](PageBeanProject.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SuggestedPrioritiesForMappings

> PageBeanPriorityWithSequence SuggestedPrioritiesForMappings(ctx).SuggestedMappingsRequestBean(suggestedMappingsRequestBean).Execute()

Suggested priorities for mappings



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
	suggestedMappingsRequestBean := *openapiclient.NewSuggestedMappingsRequestBean() // SuggestedMappingsRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.SuggestedPrioritiesForMappings(context.Background()).SuggestedMappingsRequestBean(suggestedMappingsRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.SuggestedPrioritiesForMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SuggestedPrioritiesForMappings`: PageBeanPriorityWithSequence
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.SuggestedPrioritiesForMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSuggestedPrioritiesForMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **suggestedMappingsRequestBean** | [**SuggestedMappingsRequestBean**](SuggestedMappingsRequestBean.md) |  | 

### Return type

[**PageBeanPriorityWithSequence**](PageBeanPriorityWithSequence.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePriorityScheme

> UpdatePrioritySchemeResponseBean UpdatePriorityScheme(ctx, schemeId).UpdatePrioritySchemeRequestBean(updatePrioritySchemeRequestBean).Execute()

Update priority scheme



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
	schemeId := int64(789) // int64 | The ID of the priority scheme.
	updatePrioritySchemeRequestBean := *openapiclient.NewUpdatePrioritySchemeRequestBean() // UpdatePrioritySchemeRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PrioritySchemesAPI.UpdatePriorityScheme(context.Background(), schemeId).UpdatePrioritySchemeRequestBean(updatePrioritySchemeRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PrioritySchemesAPI.UpdatePriorityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePriorityScheme`: UpdatePrioritySchemeResponseBean
	fmt.Fprintf(os.Stdout, "Response from `PrioritySchemesAPI.UpdatePriorityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the priority scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrioritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePrioritySchemeRequestBean** | [**UpdatePrioritySchemeRequestBean**](UpdatePrioritySchemeRequestBean.md) |  | 

### Return type

[**UpdatePrioritySchemeResponseBean**](UpdatePrioritySchemeResponseBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

