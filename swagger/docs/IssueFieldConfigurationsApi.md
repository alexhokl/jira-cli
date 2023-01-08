# \IssueFieldConfigurationsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignFieldConfigurationSchemeToProject**](IssueFieldConfigurationsAPI.md#AssignFieldConfigurationSchemeToProject) | **Put** /rest/api/3/fieldconfigurationscheme/project | Assign field configuration scheme to project
[**CreateFieldConfiguration**](IssueFieldConfigurationsAPI.md#CreateFieldConfiguration) | **Post** /rest/api/3/fieldconfiguration | Create field configuration
[**CreateFieldConfigurationScheme**](IssueFieldConfigurationsAPI.md#CreateFieldConfigurationScheme) | **Post** /rest/api/3/fieldconfigurationscheme | Create field configuration scheme
[**DeleteFieldConfiguration**](IssueFieldConfigurationsAPI.md#DeleteFieldConfiguration) | **Delete** /rest/api/3/fieldconfiguration/{id} | Delete field configuration
[**DeleteFieldConfigurationScheme**](IssueFieldConfigurationsAPI.md#DeleteFieldConfigurationScheme) | **Delete** /rest/api/3/fieldconfigurationscheme/{id} | Delete field configuration scheme
[**GetAllFieldConfigurationSchemes**](IssueFieldConfigurationsAPI.md#GetAllFieldConfigurationSchemes) | **Get** /rest/api/3/fieldconfigurationscheme | Get all field configuration schemes
[**GetAllFieldConfigurations**](IssueFieldConfigurationsAPI.md#GetAllFieldConfigurations) | **Get** /rest/api/3/fieldconfiguration | Get all field configurations
[**GetFieldConfigurationItems**](IssueFieldConfigurationsAPI.md#GetFieldConfigurationItems) | **Get** /rest/api/3/fieldconfiguration/{id}/fields | Get field configuration items
[**GetFieldConfigurationSchemeMappings**](IssueFieldConfigurationsAPI.md#GetFieldConfigurationSchemeMappings) | **Get** /rest/api/3/fieldconfigurationscheme/mapping | Get field configuration issue type items
[**GetFieldConfigurationSchemeProjectMapping**](IssueFieldConfigurationsAPI.md#GetFieldConfigurationSchemeProjectMapping) | **Get** /rest/api/3/fieldconfigurationscheme/project | Get field configuration schemes for projects
[**RemoveIssueTypesFromGlobalFieldConfigurationScheme**](IssueFieldConfigurationsAPI.md#RemoveIssueTypesFromGlobalFieldConfigurationScheme) | **Post** /rest/api/3/fieldconfigurationscheme/{id}/mapping/delete | Remove issue types from field configuration scheme
[**SetFieldConfigurationSchemeMapping**](IssueFieldConfigurationsAPI.md#SetFieldConfigurationSchemeMapping) | **Put** /rest/api/3/fieldconfigurationscheme/{id}/mapping | Assign issue types to field configurations
[**UpdateFieldConfiguration**](IssueFieldConfigurationsAPI.md#UpdateFieldConfiguration) | **Put** /rest/api/3/fieldconfiguration/{id} | Update field configuration
[**UpdateFieldConfigurationItems**](IssueFieldConfigurationsAPI.md#UpdateFieldConfigurationItems) | **Put** /rest/api/3/fieldconfiguration/{id}/fields | Update field configuration items
[**UpdateFieldConfigurationScheme**](IssueFieldConfigurationsAPI.md#UpdateFieldConfigurationScheme) | **Put** /rest/api/3/fieldconfigurationscheme/{id} | Update field configuration scheme



## AssignFieldConfigurationSchemeToProject

> interface{} AssignFieldConfigurationSchemeToProject(ctx).FieldConfigurationSchemeProjectAssociation(fieldConfigurationSchemeProjectAssociation).Execute()

Assign field configuration scheme to project



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
	fieldConfigurationSchemeProjectAssociation := *openapiclient.NewFieldConfigurationSchemeProjectAssociation("ProjectId_example") // FieldConfigurationSchemeProjectAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.AssignFieldConfigurationSchemeToProject(context.Background()).FieldConfigurationSchemeProjectAssociation(fieldConfigurationSchemeProjectAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.AssignFieldConfigurationSchemeToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignFieldConfigurationSchemeToProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.AssignFieldConfigurationSchemeToProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignFieldConfigurationSchemeToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldConfigurationSchemeProjectAssociation** | [**FieldConfigurationSchemeProjectAssociation**](FieldConfigurationSchemeProjectAssociation.md) |  | 

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


## CreateFieldConfiguration

> FieldConfiguration CreateFieldConfiguration(ctx).FieldConfigurationDetails(fieldConfigurationDetails).Execute()

Create field configuration



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
	fieldConfigurationDetails := *openapiclient.NewFieldConfigurationDetails("Name_example") // FieldConfigurationDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.CreateFieldConfiguration(context.Background()).FieldConfigurationDetails(fieldConfigurationDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.CreateFieldConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFieldConfiguration`: FieldConfiguration
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.CreateFieldConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldConfigurationDetails** | [**FieldConfigurationDetails**](FieldConfigurationDetails.md) |  | 

### Return type

[**FieldConfiguration**](FieldConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldConfigurationScheme

> FieldConfigurationScheme CreateFieldConfigurationScheme(ctx).UpdateFieldConfigurationSchemeDetails(updateFieldConfigurationSchemeDetails).Execute()

Create field configuration scheme



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
	updateFieldConfigurationSchemeDetails := *openapiclient.NewUpdateFieldConfigurationSchemeDetails("Name_example") // UpdateFieldConfigurationSchemeDetails | The details of the field configuration scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.CreateFieldConfigurationScheme(context.Background()).UpdateFieldConfigurationSchemeDetails(updateFieldConfigurationSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.CreateFieldConfigurationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFieldConfigurationScheme`: FieldConfigurationScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.CreateFieldConfigurationScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldConfigurationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateFieldConfigurationSchemeDetails** | [**UpdateFieldConfigurationSchemeDetails**](UpdateFieldConfigurationSchemeDetails.md) | The details of the field configuration scheme. | 

### Return type

[**FieldConfigurationScheme**](FieldConfigurationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldConfiguration

> interface{} DeleteFieldConfiguration(ctx, id).Execute()

Delete field configuration



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
	id := int64(789) // int64 | The ID of the field configuration.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.DeleteFieldConfiguration(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.DeleteFieldConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFieldConfiguration`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.DeleteFieldConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldConfigurationRequest struct via the builder pattern


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


## DeleteFieldConfigurationScheme

> interface{} DeleteFieldConfigurationScheme(ctx, id).Execute()

Delete field configuration scheme



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
	id := int64(789) // int64 | The ID of the field configuration scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.DeleteFieldConfigurationScheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.DeleteFieldConfigurationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFieldConfigurationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.DeleteFieldConfigurationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldConfigurationSchemeRequest struct via the builder pattern


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


## GetAllFieldConfigurationSchemes

> PageBeanFieldConfigurationScheme GetAllFieldConfigurationSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).Execute()

Get all field configuration schemes



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
	id := []int64{int64(123)} // []int64 | The list of field configuration scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.GetAllFieldConfigurationSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.GetAllFieldConfigurationSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFieldConfigurationSchemes`: PageBeanFieldConfigurationScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.GetAllFieldConfigurationSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFieldConfigurationSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **id** | **[]int64** | The list of field configuration scheme IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 

### Return type

[**PageBeanFieldConfigurationScheme**](PageBeanFieldConfigurationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllFieldConfigurations

> PageBeanFieldConfigurationDetails GetAllFieldConfigurations(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).IsDefault(isDefault).Query(query).Execute()

Get all field configurations



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
	id := []int64{int64(123)} // []int64 | The list of field configuration IDs. To include multiple IDs, provide an ampersand-separated list. For example, `id=10000&id=10001`. (optional)
	isDefault := true // bool | If *true* returns default field configurations only. (optional) (default to false)
	query := "query_example" // string | The query string used to match against field configuration names and descriptions. (optional) (default to "")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.GetAllFieldConfigurations(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).IsDefault(isDefault).Query(query).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.GetAllFieldConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllFieldConfigurations`: PageBeanFieldConfigurationDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.GetAllFieldConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllFieldConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **id** | **[]int64** | The list of field configuration IDs. To include multiple IDs, provide an ampersand-separated list. For example, &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **isDefault** | **bool** | If *true* returns default field configurations only. | [default to false]
 **query** | **string** | The query string used to match against field configuration names and descriptions. | [default to &quot;&quot;]

### Return type

[**PageBeanFieldConfigurationDetails**](PageBeanFieldConfigurationDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldConfigurationItems

> PageBeanFieldConfigurationItem GetFieldConfigurationItems(ctx, id).StartAt(startAt).MaxResults(maxResults).Execute()

Get field configuration items



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
	id := int64(789) // int64 | The ID of the field configuration.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.GetFieldConfigurationItems(context.Background(), id).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.GetFieldConfigurationItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldConfigurationItems`: PageBeanFieldConfigurationItem
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.GetFieldConfigurationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanFieldConfigurationItem**](PageBeanFieldConfigurationItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldConfigurationSchemeMappings

> PageBeanFieldConfigurationIssueTypeItem GetFieldConfigurationSchemeMappings(ctx).StartAt(startAt).MaxResults(maxResults).FieldConfigurationSchemeId(fieldConfigurationSchemeId).Execute()

Get field configuration issue type items



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
	fieldConfigurationSchemeId := []int64{int64(10020)} // []int64 | The list of field configuration scheme IDs. To include multiple field configuration schemes separate IDs with ampersand: `fieldConfigurationSchemeId=10000&fieldConfigurationSchemeId=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeMappings(context.Background()).StartAt(startAt).MaxResults(maxResults).FieldConfigurationSchemeId(fieldConfigurationSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldConfigurationSchemeMappings`: PageBeanFieldConfigurationIssueTypeItem
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldConfigurationSchemeMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **fieldConfigurationSchemeId** | **[]int64** | The list of field configuration scheme IDs. To include multiple field configuration schemes separate IDs with ampersand: &#x60;fieldConfigurationSchemeId&#x3D;10000&amp;fieldConfigurationSchemeId&#x3D;10001&#x60;. | 

### Return type

[**PageBeanFieldConfigurationIssueTypeItem**](PageBeanFieldConfigurationIssueTypeItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldConfigurationSchemeProjectMapping

> PageBeanFieldConfigurationSchemeProjects GetFieldConfigurationSchemeProjectMapping(ctx).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()

Get field configuration schemes for projects



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
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeProjectMapping(context.Background()).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeProjectMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldConfigurationSchemeProjectMapping`: PageBeanFieldConfigurationSchemeProjects
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.GetFieldConfigurationSchemeProjectMapping`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldConfigurationSchemeProjectMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | The list of project IDs. To include multiple projects, separate IDs with ampersand: &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanFieldConfigurationSchemeProjects**](PageBeanFieldConfigurationSchemeProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveIssueTypesFromGlobalFieldConfigurationScheme

> interface{} RemoveIssueTypesFromGlobalFieldConfigurationScheme(ctx, id).IssueTypeIdsToRemove(issueTypeIdsToRemove).Execute()

Remove issue types from field configuration scheme



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
	id := int64(789) // int64 | The ID of the field configuration scheme.
	issueTypeIdsToRemove := *openapiclient.NewIssueTypeIdsToRemove([]string{"IssueTypeIds_example"}) // IssueTypeIdsToRemove | The issue type IDs to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.RemoveIssueTypesFromGlobalFieldConfigurationScheme(context.Background(), id).IssueTypeIdsToRemove(issueTypeIdsToRemove).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.RemoveIssueTypesFromGlobalFieldConfigurationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIssueTypesFromGlobalFieldConfigurationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.RemoveIssueTypesFromGlobalFieldConfigurationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIssueTypesFromGlobalFieldConfigurationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeIdsToRemove** | [**IssueTypeIdsToRemove**](IssueTypeIdsToRemove.md) | The issue type IDs to remove. | 

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


## SetFieldConfigurationSchemeMapping

> interface{} SetFieldConfigurationSchemeMapping(ctx, id).AssociateFieldConfigurationsWithIssueTypesRequest(associateFieldConfigurationsWithIssueTypesRequest).Execute()

Assign issue types to field configurations



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
	id := int64(789) // int64 | The ID of the field configuration scheme.
	associateFieldConfigurationsWithIssueTypesRequest := *openapiclient.NewAssociateFieldConfigurationsWithIssueTypesRequest([]openapiclient.FieldConfigurationToIssueTypeMapping{*openapiclient.NewFieldConfigurationToIssueTypeMapping("FieldConfigurationId_example", "IssueTypeId_example")}) // AssociateFieldConfigurationsWithIssueTypesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.SetFieldConfigurationSchemeMapping(context.Background(), id).AssociateFieldConfigurationsWithIssueTypesRequest(associateFieldConfigurationsWithIssueTypesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.SetFieldConfigurationSchemeMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetFieldConfigurationSchemeMapping`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.SetFieldConfigurationSchemeMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetFieldConfigurationSchemeMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **associateFieldConfigurationsWithIssueTypesRequest** | [**AssociateFieldConfigurationsWithIssueTypesRequest**](AssociateFieldConfigurationsWithIssueTypesRequest.md) |  | 

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


## UpdateFieldConfiguration

> interface{} UpdateFieldConfiguration(ctx, id).FieldConfigurationDetails(fieldConfigurationDetails).Execute()

Update field configuration



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
	id := int64(789) // int64 | The ID of the field configuration.
	fieldConfigurationDetails := *openapiclient.NewFieldConfigurationDetails("Name_example") // FieldConfigurationDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.UpdateFieldConfiguration(context.Background(), id).FieldConfigurationDetails(fieldConfigurationDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.UpdateFieldConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldConfiguration`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.UpdateFieldConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldConfigurationDetails** | [**FieldConfigurationDetails**](FieldConfigurationDetails.md) |  | 

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


## UpdateFieldConfigurationItems

> interface{} UpdateFieldConfigurationItems(ctx, id).FieldConfigurationItemsDetails(fieldConfigurationItemsDetails).Execute()

Update field configuration items



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
	id := int64(789) // int64 | The ID of the field configuration.
	fieldConfigurationItemsDetails := *openapiclient.NewFieldConfigurationItemsDetails([]openapiclient.FieldConfigurationItem{*openapiclient.NewFieldConfigurationItem("Id_example")}) // FieldConfigurationItemsDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.UpdateFieldConfigurationItems(context.Background(), id).FieldConfigurationItemsDetails(fieldConfigurationItemsDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.UpdateFieldConfigurationItems``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldConfigurationItems`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.UpdateFieldConfigurationItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldConfigurationItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fieldConfigurationItemsDetails** | [**FieldConfigurationItemsDetails**](FieldConfigurationItemsDetails.md) |  | 

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


## UpdateFieldConfigurationScheme

> interface{} UpdateFieldConfigurationScheme(ctx, id).UpdateFieldConfigurationSchemeDetails(updateFieldConfigurationSchemeDetails).Execute()

Update field configuration scheme



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
	id := int64(789) // int64 | The ID of the field configuration scheme.
	updateFieldConfigurationSchemeDetails := *openapiclient.NewUpdateFieldConfigurationSchemeDetails("Name_example") // UpdateFieldConfigurationSchemeDetails | The details of the field configuration scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldConfigurationsAPI.UpdateFieldConfigurationScheme(context.Background(), id).UpdateFieldConfigurationSchemeDetails(updateFieldConfigurationSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldConfigurationsAPI.UpdateFieldConfigurationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldConfigurationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldConfigurationsAPI.UpdateFieldConfigurationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field configuration scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldConfigurationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldConfigurationSchemeDetails** | [**UpdateFieldConfigurationSchemeDetails**](UpdateFieldConfigurationSchemeDetails.md) | The details of the field configuration scheme. | 

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

