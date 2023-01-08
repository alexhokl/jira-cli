# \IssueTypeSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIssueTypesToIssueTypeScheme**](IssueTypeSchemesAPI.md#AddIssueTypesToIssueTypeScheme) | **Put** /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype | Add issue types to issue type scheme
[**AssignIssueTypeSchemeToProject**](IssueTypeSchemesAPI.md#AssignIssueTypeSchemeToProject) | **Put** /rest/api/3/issuetypescheme/project | Assign issue type scheme to project
[**CreateIssueTypeScheme**](IssueTypeSchemesAPI.md#CreateIssueTypeScheme) | **Post** /rest/api/3/issuetypescheme | Create issue type scheme
[**DeleteIssueTypeScheme**](IssueTypeSchemesAPI.md#DeleteIssueTypeScheme) | **Delete** /rest/api/3/issuetypescheme/{issueTypeSchemeId} | Delete issue type scheme
[**GetAllIssueTypeSchemes**](IssueTypeSchemesAPI.md#GetAllIssueTypeSchemes) | **Get** /rest/api/3/issuetypescheme | Get all issue type schemes
[**GetIssueTypeSchemeForProjects**](IssueTypeSchemesAPI.md#GetIssueTypeSchemeForProjects) | **Get** /rest/api/3/issuetypescheme/project | Get issue type schemes for projects
[**GetIssueTypeSchemesMapping**](IssueTypeSchemesAPI.md#GetIssueTypeSchemesMapping) | **Get** /rest/api/3/issuetypescheme/mapping | Get issue type scheme items
[**RemoveIssueTypeFromIssueTypeScheme**](IssueTypeSchemesAPI.md#RemoveIssueTypeFromIssueTypeScheme) | **Delete** /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype/{issueTypeId} | Remove issue type from issue type scheme
[**ReorderIssueTypesInIssueTypeScheme**](IssueTypeSchemesAPI.md#ReorderIssueTypesInIssueTypeScheme) | **Put** /rest/api/3/issuetypescheme/{issueTypeSchemeId}/issuetype/move | Change order of issue types
[**UpdateIssueTypeScheme**](IssueTypeSchemesAPI.md#UpdateIssueTypeScheme) | **Put** /rest/api/3/issuetypescheme/{issueTypeSchemeId} | Update issue type scheme



## AddIssueTypesToIssueTypeScheme

> interface{} AddIssueTypesToIssueTypeScheme(ctx, issueTypeSchemeId).IssueTypeIds(issueTypeIds).Execute()

Add issue types to issue type scheme



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
	issueTypeSchemeId := int64(789) // int64 | The ID of the issue type scheme.
	issueTypeIds := *openapiclient.NewIssueTypeIds([]string{"IssueTypeIds_example"}) // IssueTypeIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.AddIssueTypesToIssueTypeScheme(context.Background(), issueTypeSchemeId).IssueTypeIds(issueTypeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.AddIssueTypesToIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddIssueTypesToIssueTypeScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.AddIssueTypesToIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeSchemeId** | **int64** | The ID of the issue type scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIssueTypesToIssueTypeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeIds** | [**IssueTypeIds**](IssueTypeIds.md) |  | 

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


## AssignIssueTypeSchemeToProject

> interface{} AssignIssueTypeSchemeToProject(ctx).IssueTypeSchemeProjectAssociation(issueTypeSchemeProjectAssociation).Execute()

Assign issue type scheme to project



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
	issueTypeSchemeProjectAssociation := *openapiclient.NewIssueTypeSchemeProjectAssociation("IssueTypeSchemeId_example", "ProjectId_example") // IssueTypeSchemeProjectAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.AssignIssueTypeSchemeToProject(context.Background()).IssueTypeSchemeProjectAssociation(issueTypeSchemeProjectAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.AssignIssueTypeSchemeToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignIssueTypeSchemeToProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.AssignIssueTypeSchemeToProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignIssueTypeSchemeToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTypeSchemeProjectAssociation** | [**IssueTypeSchemeProjectAssociation**](IssueTypeSchemeProjectAssociation.md) |  | 

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


## CreateIssueTypeScheme

> IssueTypeSchemeID CreateIssueTypeScheme(ctx).IssueTypeSchemeDetails(issueTypeSchemeDetails).Execute()

Create issue type scheme



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
	issueTypeSchemeDetails := *openapiclient.NewIssueTypeSchemeDetails([]string{"IssueTypeIds_example"}, "Name_example") // IssueTypeSchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.CreateIssueTypeScheme(context.Background()).IssueTypeSchemeDetails(issueTypeSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.CreateIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueTypeScheme`: IssueTypeSchemeID
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.CreateIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueTypeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTypeSchemeDetails** | [**IssueTypeSchemeDetails**](IssueTypeSchemeDetails.md) |  | 

### Return type

[**IssueTypeSchemeID**](IssueTypeSchemeID.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueTypeScheme

> interface{} DeleteIssueTypeScheme(ctx, issueTypeSchemeId).Execute()

Delete issue type scheme



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
	issueTypeSchemeId := int64(789) // int64 | The ID of the issue type scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.DeleteIssueTypeScheme(context.Background(), issueTypeSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.DeleteIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueTypeScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.DeleteIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeSchemeId** | **int64** | The ID of the issue type scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueTypeSchemeRequest struct via the builder pattern


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


## GetAllIssueTypeSchemes

> PageBeanIssueTypeScheme GetAllIssueTypeSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).OrderBy(orderBy).Expand(expand).QueryString(queryString).Execute()

Get all issue type schemes



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
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	id := []int64{int64(123)} // []int64 | The list of issue type schemes IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `name` Sorts by issue type scheme name.  *  `id` Sorts by issue type scheme ID. (optional) (default to "id")
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `projects` For each issue type schemes, returns information about the projects the issue type scheme is assigned to.  *  `issueTypes` For each issue type schemes, returns information about the issueTypes the issue type scheme have. (optional) (default to "")
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with issue type scheme name. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.GetAllIssueTypeSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).OrderBy(orderBy).Expand(expand).QueryString(queryString).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.GetAllIssueTypeSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIssueTypeSchemes`: PageBeanIssueTypeScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.GetAllIssueTypeSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIssueTypeSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **id** | **[]int64** | The list of issue type schemes IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;name&#x60; Sorts by issue type scheme name.  *  &#x60;id&#x60; Sorts by issue type scheme ID. | [default to &quot;id&quot;]
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;projects&#x60; For each issue type schemes, returns information about the projects the issue type scheme is assigned to.  *  &#x60;issueTypes&#x60; For each issue type schemes, returns information about the issueTypes the issue type scheme have. | [default to &quot;&quot;]
 **queryString** | **string** | String used to perform a case-insensitive partial match with issue type scheme name. | [default to &quot;&quot;]

### Return type

[**PageBeanIssueTypeScheme**](PageBeanIssueTypeScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypeSchemeForProjects

> PageBeanIssueTypeSchemeProjects GetIssueTypeSchemeForProjects(ctx).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()

Get issue type schemes for projects



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
	projectId := []int64{int64(123)} // []int64 | The list of project IDs. To include multiple project IDs, provide an ampersand-separated list. For example, `projectId=10000&projectId=10001`.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.GetIssueTypeSchemeForProjects(context.Background()).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.GetIssueTypeSchemeForProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeSchemeForProjects`: PageBeanIssueTypeSchemeProjects
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.GetIssueTypeSchemeForProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeSchemeForProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | The list of project IDs. To include multiple project IDs, provide an ampersand-separated list. For example, &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanIssueTypeSchemeProjects**](PageBeanIssueTypeSchemeProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypeSchemesMapping

> PageBeanIssueTypeSchemeMapping GetIssueTypeSchemesMapping(ctx).StartAt(startAt).MaxResults(maxResults).IssueTypeSchemeId(issueTypeSchemeId).Execute()

Get issue type scheme items



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
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	issueTypeSchemeId := []int64{int64(123)} // []int64 | The list of issue type scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `issueTypeSchemeId=10000&issueTypeSchemeId=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.GetIssueTypeSchemesMapping(context.Background()).StartAt(startAt).MaxResults(maxResults).IssueTypeSchemeId(issueTypeSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.GetIssueTypeSchemesMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeSchemesMapping`: PageBeanIssueTypeSchemeMapping
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.GetIssueTypeSchemesMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeSchemesMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **issueTypeSchemeId** | **[]int64** | The list of issue type scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;issueTypeSchemeId&#x3D;10000&amp;issueTypeSchemeId&#x3D;10001&#x60;. | 

### Return type

[**PageBeanIssueTypeSchemeMapping**](PageBeanIssueTypeSchemeMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIssueTypeFromIssueTypeScheme

> interface{} RemoveIssueTypeFromIssueTypeScheme(ctx, issueTypeSchemeId, issueTypeId).Execute()

Remove issue type from issue type scheme



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
	issueTypeSchemeId := int64(789) // int64 | The ID of the issue type scheme.
	issueTypeId := int64(789) // int64 | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.RemoveIssueTypeFromIssueTypeScheme(context.Background(), issueTypeSchemeId, issueTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.RemoveIssueTypeFromIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIssueTypeFromIssueTypeScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.RemoveIssueTypeFromIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeSchemeId** | **int64** | The ID of the issue type scheme. | 
**issueTypeId** | **int64** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIssueTypeFromIssueTypeSchemeRequest struct via the builder pattern


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


## ReorderIssueTypesInIssueTypeScheme

> interface{} ReorderIssueTypesInIssueTypeScheme(ctx, issueTypeSchemeId).OrderOfIssueTypes(orderOfIssueTypes).Execute()

Change order of issue types



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
	issueTypeSchemeId := int64(789) // int64 | The ID of the issue type scheme.
	orderOfIssueTypes := *openapiclient.NewOrderOfIssueTypes([]string{"IssueTypeIds_example"}) // OrderOfIssueTypes | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.ReorderIssueTypesInIssueTypeScheme(context.Background(), issueTypeSchemeId).OrderOfIssueTypes(orderOfIssueTypes).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.ReorderIssueTypesInIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderIssueTypesInIssueTypeScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.ReorderIssueTypesInIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeSchemeId** | **int64** | The ID of the issue type scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderIssueTypesInIssueTypeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **orderOfIssueTypes** | [**OrderOfIssueTypes**](OrderOfIssueTypes.md) |  | 

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


## UpdateIssueTypeScheme

> interface{} UpdateIssueTypeScheme(ctx, issueTypeSchemeId).IssueTypeSchemeUpdateDetails(issueTypeSchemeUpdateDetails).Execute()

Update issue type scheme



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
	issueTypeSchemeId := int64(789) // int64 | The ID of the issue type scheme.
	issueTypeSchemeUpdateDetails := *openapiclient.NewIssueTypeSchemeUpdateDetails() // IssueTypeSchemeUpdateDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeSchemesAPI.UpdateIssueTypeScheme(context.Background(), issueTypeSchemeId).IssueTypeSchemeUpdateDetails(issueTypeSchemeUpdateDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeSchemesAPI.UpdateIssueTypeScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueTypeScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeSchemesAPI.UpdateIssueTypeScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeSchemeId** | **int64** | The ID of the issue type scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueTypeSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeSchemeUpdateDetails** | [**IssueTypeSchemeUpdateDetails**](IssueTypeSchemeUpdateDetails.md) |  | 

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

