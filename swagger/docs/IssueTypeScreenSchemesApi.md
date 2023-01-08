# \IssueTypeScreenSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppendMappingsForIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#AppendMappingsForIssueTypeScreenScheme) | **Put** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping | Append mappings to issue type screen scheme
[**AssignIssueTypeScreenSchemeToProject**](IssueTypeScreenSchemesAPI.md#AssignIssueTypeScreenSchemeToProject) | **Put** /rest/api/3/issuetypescreenscheme/project | Assign issue type screen scheme to project
[**CreateIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#CreateIssueTypeScreenScheme) | **Post** /rest/api/3/issuetypescreenscheme | Create issue type screen scheme
[**DeleteIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#DeleteIssueTypeScreenScheme) | **Delete** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId} | Delete issue type screen scheme
[**GetIssueTypeScreenSchemeMappings**](IssueTypeScreenSchemesAPI.md#GetIssueTypeScreenSchemeMappings) | **Get** /rest/api/3/issuetypescreenscheme/mapping | Get issue type screen scheme items
[**GetIssueTypeScreenSchemeProjectAssociations**](IssueTypeScreenSchemesAPI.md#GetIssueTypeScreenSchemeProjectAssociations) | **Get** /rest/api/3/issuetypescreenscheme/project | Get issue type screen schemes for projects
[**GetIssueTypeScreenSchemes**](IssueTypeScreenSchemesAPI.md#GetIssueTypeScreenSchemes) | **Get** /rest/api/3/issuetypescreenscheme | Get issue type screen schemes
[**GetProjectsForIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#GetProjectsForIssueTypeScreenScheme) | **Get** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/project | Get issue type screen scheme projects
[**RemoveMappingsFromIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#RemoveMappingsFromIssueTypeScreenScheme) | **Post** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/remove | Remove mappings from issue type screen scheme
[**UpdateDefaultScreenScheme**](IssueTypeScreenSchemesAPI.md#UpdateDefaultScreenScheme) | **Put** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId}/mapping/default | Update issue type screen scheme default screen scheme
[**UpdateIssueTypeScreenScheme**](IssueTypeScreenSchemesAPI.md#UpdateIssueTypeScreenScheme) | **Put** /rest/api/3/issuetypescreenscheme/{issueTypeScreenSchemeId} | Update issue type screen scheme



## AppendMappingsForIssueTypeScreenScheme

> interface{} AppendMappingsForIssueTypeScreenScheme(ctx, issueTypeScreenSchemeId).IssueTypeScreenSchemeMappingDetails(issueTypeScreenSchemeMappingDetails).Execute()

Append mappings to issue type screen scheme



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
	issueTypeScreenSchemeId := "issueTypeScreenSchemeId_example" // string | The ID of the issue type screen scheme.
	issueTypeScreenSchemeMappingDetails := *openapiclient.NewIssueTypeScreenSchemeMappingDetails([]openapiclient.IssueTypeScreenSchemeMapping{*openapiclient.NewIssueTypeScreenSchemeMapping("IssueTypeId_example", "ScreenSchemeId_example")}) // IssueTypeScreenSchemeMappingDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.AppendMappingsForIssueTypeScreenScheme(context.Background(), issueTypeScreenSchemeId).IssueTypeScreenSchemeMappingDetails(issueTypeScreenSchemeMappingDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.AppendMappingsForIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppendMappingsForIssueTypeScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.AppendMappingsForIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **string** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAppendMappingsForIssueTypeScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeScreenSchemeMappingDetails** | [**IssueTypeScreenSchemeMappingDetails**](IssueTypeScreenSchemeMappingDetails.md) |  | 

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


## AssignIssueTypeScreenSchemeToProject

> interface{} AssignIssueTypeScreenSchemeToProject(ctx).IssueTypeScreenSchemeProjectAssociation(issueTypeScreenSchemeProjectAssociation).Execute()

Assign issue type screen scheme to project



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
	issueTypeScreenSchemeProjectAssociation := *openapiclient.NewIssueTypeScreenSchemeProjectAssociation() // IssueTypeScreenSchemeProjectAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.AssignIssueTypeScreenSchemeToProject(context.Background()).IssueTypeScreenSchemeProjectAssociation(issueTypeScreenSchemeProjectAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.AssignIssueTypeScreenSchemeToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignIssueTypeScreenSchemeToProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.AssignIssueTypeScreenSchemeToProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignIssueTypeScreenSchemeToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTypeScreenSchemeProjectAssociation** | [**IssueTypeScreenSchemeProjectAssociation**](IssueTypeScreenSchemeProjectAssociation.md) |  | 

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


## CreateIssueTypeScreenScheme

> IssueTypeScreenSchemeId CreateIssueTypeScreenScheme(ctx).IssueTypeScreenSchemeDetails(issueTypeScreenSchemeDetails).Execute()

Create issue type screen scheme



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
	issueTypeScreenSchemeDetails := *openapiclient.NewIssueTypeScreenSchemeDetails([]openapiclient.IssueTypeScreenSchemeMapping{*openapiclient.NewIssueTypeScreenSchemeMapping("IssueTypeId_example", "ScreenSchemeId_example")}, "Name_example") // IssueTypeScreenSchemeDetails | An issue type screen scheme bean.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.CreateIssueTypeScreenScheme(context.Background()).IssueTypeScreenSchemeDetails(issueTypeScreenSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.CreateIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueTypeScreenScheme`: IssueTypeScreenSchemeId
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.CreateIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueTypeScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTypeScreenSchemeDetails** | [**IssueTypeScreenSchemeDetails**](IssueTypeScreenSchemeDetails.md) | An issue type screen scheme bean. | 

### Return type

[**IssueTypeScreenSchemeId**](IssueTypeScreenSchemeId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueTypeScreenScheme

> interface{} DeleteIssueTypeScreenScheme(ctx, issueTypeScreenSchemeId).Execute()

Delete issue type screen scheme



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
	issueTypeScreenSchemeId := "issueTypeScreenSchemeId_example" // string | The ID of the issue type screen scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.DeleteIssueTypeScreenScheme(context.Background(), issueTypeScreenSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.DeleteIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueTypeScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.DeleteIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **string** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueTypeScreenSchemeRequest struct via the builder pattern


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


## GetIssueTypeScreenSchemeMappings

> PageBeanIssueTypeScreenSchemeItem GetIssueTypeScreenSchemeMappings(ctx).StartAt(startAt).MaxResults(maxResults).IssueTypeScreenSchemeId(issueTypeScreenSchemeId).Execute()

Get issue type screen scheme items



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
	issueTypeScreenSchemeId := []int64{int64(123)} // []int64 | The list of issue type screen scheme IDs. To include multiple issue type screen schemes, separate IDs with ampersand: `issueTypeScreenSchemeId=10000&issueTypeScreenSchemeId=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeMappings(context.Background()).StartAt(startAt).MaxResults(maxResults).IssueTypeScreenSchemeId(issueTypeScreenSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeScreenSchemeMappings`: PageBeanIssueTypeScreenSchemeItem
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeScreenSchemeMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **issueTypeScreenSchemeId** | **[]int64** | The list of issue type screen scheme IDs. To include multiple issue type screen schemes, separate IDs with ampersand: &#x60;issueTypeScreenSchemeId&#x3D;10000&amp;issueTypeScreenSchemeId&#x3D;10001&#x60;. | 

### Return type

[**PageBeanIssueTypeScreenSchemeItem**](PageBeanIssueTypeScreenSchemeItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypeScreenSchemeProjectAssociations

> PageBeanIssueTypeScreenSchemesProjects GetIssueTypeScreenSchemeProjectAssociations(ctx).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()

Get issue type screen schemes for projects



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
	projectId := []int64{int64(123)} // []int64 | The list of project IDs. To include multiple projects, separate IDs with ampersand: `projectId=10000&projectId=10001`.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeProjectAssociations(context.Background()).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeProjectAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeScreenSchemeProjectAssociations`: PageBeanIssueTypeScreenSchemesProjects
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemeProjectAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeScreenSchemeProjectAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | The list of project IDs. To include multiple projects, separate IDs with ampersand: &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanIssueTypeScreenSchemesProjects**](PageBeanIssueTypeScreenSchemesProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypeScreenSchemes

> PageBeanIssueTypeScreenScheme GetIssueTypeScreenSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).QueryString(queryString).OrderBy(orderBy).Expand(expand).Execute()

Get issue type screen schemes



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
	id := []int64{int64(123)} // []int64 | The list of issue type screen scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with issue type screen scheme name. (optional) (default to "")
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `name` Sorts by issue type screen scheme name.  *  `id` Sorts by issue type screen scheme ID. (optional) (default to "id")
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts `projects` that, for each issue type screen schemes, returns information about the projects the issue type screen scheme is assigned to. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).QueryString(queryString).OrderBy(orderBy).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeScreenSchemes`: PageBeanIssueTypeScreenScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.GetIssueTypeScreenSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeScreenSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **id** | **[]int64** | The list of issue type screen scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **queryString** | **string** | String used to perform a case-insensitive partial match with issue type screen scheme name. | [default to &quot;&quot;]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;name&#x60; Sorts by issue type screen scheme name.  *  &#x60;id&#x60; Sorts by issue type screen scheme ID. | [default to &quot;id&quot;]
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;projects&#x60; that, for each issue type screen schemes, returns information about the projects the issue type screen scheme is assigned to. | [default to &quot;&quot;]

### Return type

[**PageBeanIssueTypeScreenScheme**](PageBeanIssueTypeScreenScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsForIssueTypeScreenScheme

> PageBeanProjectDetails GetProjectsForIssueTypeScreenScheme(ctx, issueTypeScreenSchemeId).StartAt(startAt).MaxResults(maxResults).Query(query).Execute()

Get issue type screen scheme projects



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
	issueTypeScreenSchemeId := int64(789) // int64 | The ID of the issue type screen scheme.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	query := "query_example" // string |  (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.GetProjectsForIssueTypeScreenScheme(context.Background(), issueTypeScreenSchemeId).StartAt(startAt).MaxResults(maxResults).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.GetProjectsForIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsForIssueTypeScreenScheme`: PageBeanProjectDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.GetProjectsForIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **int64** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsForIssueTypeScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **query** | **string** |  | [default to &quot;&quot;]

### Return type

[**PageBeanProjectDetails**](PageBeanProjectDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveMappingsFromIssueTypeScreenScheme

> interface{} RemoveMappingsFromIssueTypeScreenScheme(ctx, issueTypeScreenSchemeId).IssueTypeIds(issueTypeIds).Execute()

Remove mappings from issue type screen scheme



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
	issueTypeScreenSchemeId := "issueTypeScreenSchemeId_example" // string | The ID of the issue type screen scheme.
	issueTypeIds := *openapiclient.NewIssueTypeIds([]string{"IssueTypeIds_example"}) // IssueTypeIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.RemoveMappingsFromIssueTypeScreenScheme(context.Background(), issueTypeScreenSchemeId).IssueTypeIds(issueTypeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.RemoveMappingsFromIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMappingsFromIssueTypeScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.RemoveMappingsFromIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **string** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMappingsFromIssueTypeScreenSchemeRequest struct via the builder pattern


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


## UpdateDefaultScreenScheme

> interface{} UpdateDefaultScreenScheme(ctx, issueTypeScreenSchemeId).UpdateDefaultScreenScheme(updateDefaultScreenScheme).Execute()

Update issue type screen scheme default screen scheme



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
	issueTypeScreenSchemeId := "issueTypeScreenSchemeId_example" // string | The ID of the issue type screen scheme.
	updateDefaultScreenScheme := *openapiclient.NewUpdateDefaultScreenScheme("ScreenSchemeId_example") // UpdateDefaultScreenScheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.UpdateDefaultScreenScheme(context.Background(), issueTypeScreenSchemeId).UpdateDefaultScreenScheme(updateDefaultScreenScheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.UpdateDefaultScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.UpdateDefaultScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **string** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDefaultScreenScheme** | [**UpdateDefaultScreenScheme**](UpdateDefaultScreenScheme.md) |  | 

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


## UpdateIssueTypeScreenScheme

> interface{} UpdateIssueTypeScreenScheme(ctx, issueTypeScreenSchemeId).IssueTypeScreenSchemeUpdateDetails(issueTypeScreenSchemeUpdateDetails).Execute()

Update issue type screen scheme



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
	issueTypeScreenSchemeId := "issueTypeScreenSchemeId_example" // string | The ID of the issue type screen scheme.
	issueTypeScreenSchemeUpdateDetails := *openapiclient.NewIssueTypeScreenSchemeUpdateDetails() // IssueTypeScreenSchemeUpdateDetails | The issue type screen scheme update details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypeScreenSchemesAPI.UpdateIssueTypeScreenScheme(context.Background(), issueTypeScreenSchemeId).IssueTypeScreenSchemeUpdateDetails(issueTypeScreenSchemeUpdateDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypeScreenSchemesAPI.UpdateIssueTypeScreenScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueTypeScreenScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypeScreenSchemesAPI.UpdateIssueTypeScreenScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeScreenSchemeId** | **string** | The ID of the issue type screen scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueTypeScreenSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeScreenSchemeUpdateDetails** | [**IssueTypeScreenSchemeUpdateDetails**](IssueTypeScreenSchemeUpdateDetails.md) | The issue type screen scheme update details. | 

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

