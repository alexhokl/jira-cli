# \FieldSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateProjectsToFieldAssociationSchemes**](FieldSchemesAPI.md#AssociateProjectsToFieldAssociationSchemes) | **Put** /rest/api/3/config/fieldschemes/projects | Associate projects to field schemes
[**CreateFieldAssociationScheme**](FieldSchemesAPI.md#CreateFieldAssociationScheme) | **Post** /rest/api/3/config/fieldschemes | Create field scheme
[**DeleteFieldAssociationScheme**](FieldSchemesAPI.md#DeleteFieldAssociationScheme) | **Delete** /rest/api/3/config/fieldschemes/{id} | Delete a field scheme
[**GetFieldAssociationSchemeById**](FieldSchemesAPI.md#GetFieldAssociationSchemeById) | **Get** /rest/api/3/config/fieldschemes/{id} | Get field scheme
[**GetFieldAssociationSchemeItemParameters**](FieldSchemesAPI.md#GetFieldAssociationSchemeItemParameters) | **Get** /rest/api/3/config/fieldschemes/{id}/fields/{fieldId}/parameters | Get field parameters
[**GetFieldAssociationSchemes**](FieldSchemesAPI.md#GetFieldAssociationSchemes) | **Get** /rest/api/3/config/fieldschemes | Get field schemes
[**GetProjectsWithFieldSchemes**](FieldSchemesAPI.md#GetProjectsWithFieldSchemes) | **Get** /rest/api/3/config/fieldschemes/projects | Get projects with field schemes
[**RemoveFieldAssociationSchemeItemParameters**](FieldSchemesAPI.md#RemoveFieldAssociationSchemeItemParameters) | **Delete** /rest/api/3/config/fieldschemes/fields/parameters | Remove field parameters
[**RemoveFieldsAssociatedWithSchemes**](FieldSchemesAPI.md#RemoveFieldsAssociatedWithSchemes) | **Delete** /rest/api/3/config/fieldschemes/fields | Remove fields associated with field schemes
[**SearchFieldAssociationSchemeFields**](FieldSchemesAPI.md#SearchFieldAssociationSchemeFields) | **Get** /rest/api/3/config/fieldschemes/{id}/fields | Search field scheme fields
[**SearchFieldAssociationSchemeProjects**](FieldSchemesAPI.md#SearchFieldAssociationSchemeProjects) | **Get** /rest/api/3/config/fieldschemes/{id}/projects | Search field scheme projects
[**UpdateFieldAssociationScheme**](FieldSchemesAPI.md#UpdateFieldAssociationScheme) | **Put** /rest/api/3/config/fieldschemes/{id} | Update field scheme
[**UpdateFieldAssociationSchemeItemParameters**](FieldSchemesAPI.md#UpdateFieldAssociationSchemeItemParameters) | **Put** /rest/api/3/config/fieldschemes/fields/parameters | Update field parameters
[**UpdateFieldsAssociatedWithSchemes**](FieldSchemesAPI.md#UpdateFieldsAssociatedWithSchemes) | **Put** /rest/api/3/config/fieldschemes/fields | Update fields associated with field schemes



## AssociateProjectsToFieldAssociationSchemes

> FieldSchemeToProjectsResponse AssociateProjectsToFieldAssociationSchemes(ctx).RequestBody(requestBody).Execute()

Associate projects to field schemes



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
	requestBody := map[string]FieldSchemeToProjectsRequest{"key": *openapiclient.NewFieldSchemeToProjectsRequest([]int64{int64(123)})} // map[string]FieldSchemeToProjectsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.AssociateProjectsToFieldAssociationSchemes(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.AssociateProjectsToFieldAssociationSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssociateProjectsToFieldAssociationSchemes`: FieldSchemeToProjectsResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.AssociateProjectsToFieldAssociationSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociateProjectsToFieldAssociationSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]FieldSchemeToProjectsRequest**](FieldSchemeToProjectsRequest.md) |  | 

### Return type

[**FieldSchemeToProjectsResponse**](FieldSchemeToProjectsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateFieldAssociationScheme

> CreateFieldAssociationSchemeResponse CreateFieldAssociationScheme(ctx).CreateFieldAssociationSchemeRequest(createFieldAssociationSchemeRequest).Execute()

Create field scheme



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
	createFieldAssociationSchemeRequest := *openapiclient.NewCreateFieldAssociationSchemeRequest("Name_example") // CreateFieldAssociationSchemeRequest | The request containing the name and description of the field association scheme

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.CreateFieldAssociationScheme(context.Background()).CreateFieldAssociationSchemeRequest(createFieldAssociationSchemeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.CreateFieldAssociationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateFieldAssociationScheme`: CreateFieldAssociationSchemeResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.CreateFieldAssociationScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldAssociationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createFieldAssociationSchemeRequest** | [**CreateFieldAssociationSchemeRequest**](CreateFieldAssociationSchemeRequest.md) | The request containing the name and description of the field association scheme | 

### Return type

[**CreateFieldAssociationSchemeResponse**](CreateFieldAssociationSchemeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteFieldAssociationScheme

> DeleteFieldAssociationSchemeResponse DeleteFieldAssociationScheme(ctx, id).Execute()

Delete a field scheme



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
	id := int64(789) // int64 | The ID of the field association scheme to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.DeleteFieldAssociationScheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.DeleteFieldAssociationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteFieldAssociationScheme`: DeleteFieldAssociationSchemeResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.DeleteFieldAssociationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the field association scheme to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldAssociationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteFieldAssociationSchemeResponse**](DeleteFieldAssociationSchemeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldAssociationSchemeById

> GetFieldAssociationSchemeByIdResponse GetFieldAssociationSchemeById(ctx, id).Execute()

Get field scheme



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
	id := int64(789) // int64 | The scheme id to fetch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.GetFieldAssociationSchemeById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.GetFieldAssociationSchemeById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldAssociationSchemeById`: GetFieldAssociationSchemeByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.GetFieldAssociationSchemeById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The scheme id to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldAssociationSchemeByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFieldAssociationSchemeByIdResponse**](GetFieldAssociationSchemeByIdResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldAssociationSchemeItemParameters

> GetFieldAssociationParametersResponse GetFieldAssociationSchemeItemParameters(ctx, id, fieldId).Execute()

Get field parameters



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
	id := int64(789) // int64 | the ID of the field association scheme to retrieve parameters for
	fieldId := "fieldId_example" // string | the ID of the field

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.GetFieldAssociationSchemeItemParameters(context.Background(), id, fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.GetFieldAssociationSchemeItemParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldAssociationSchemeItemParameters`: GetFieldAssociationParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.GetFieldAssociationSchemeItemParameters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | the ID of the field association scheme to retrieve parameters for | 
**fieldId** | **string** | the ID of the field | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldAssociationSchemeItemParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetFieldAssociationParametersResponse**](GetFieldAssociationParametersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldAssociationSchemes

> PageBean2GetFieldAssociationSchemeResponse GetFieldAssociationSchemes(ctx).ProjectId(projectId).Query(query).StartAt(startAt).MaxResults(maxResults).Execute()

Get field schemes



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
	projectId := []int64{int64(123)} // []int64 | (optional) List of project IDs to filter schemes by. If not provided, schemes from all projects are returned. (optional)
	query := "query_example" // string | (optional) Text filter for scheme name or description matching (case-insensitive). If not provided, no text filtering is applied. (optional)
	startAt := int64(789) // int64 | Zero-based index of the first item to return (default: 0) (optional) (default to 0)
	maxResults := int32(56) // int32 | Maximum number of items to return per page (default: 50, max: 100) (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.GetFieldAssociationSchemes(context.Background()).ProjectId(projectId).Query(query).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.GetFieldAssociationSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldAssociationSchemes`: PageBean2GetFieldAssociationSchemeResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.GetFieldAssociationSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldAssociationSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | (optional) List of project IDs to filter schemes by. If not provided, schemes from all projects are returned. | 
 **query** | **string** | (optional) Text filter for scheme name or description matching (case-insensitive). If not provided, no text filtering is applied. | 
 **startAt** | **int64** | Zero-based index of the first item to return (default: 0) | [default to 0]
 **maxResults** | **int32** | Maximum number of items to return per page (default: 50, max: 100) | [default to 50]

### Return type

[**PageBean2GetFieldAssociationSchemeResponse**](PageBean2GetFieldAssociationSchemeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsWithFieldSchemes

> PageBean2GetProjectsWithFieldSchemesResponse GetProjectsWithFieldSchemes(ctx).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()

Get projects with field schemes



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
	projectId := []int64{int64(123)} // []int64 | List of project ids to filter the results by.
	startAt := int64(789) // int64 | The starting index of the returned projects. Base index: 0. (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of projects to return per page, maximum allowed value is 100. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.GetProjectsWithFieldSchemes(context.Background()).ProjectId(projectId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.GetProjectsWithFieldSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectsWithFieldSchemes`: PageBean2GetProjectsWithFieldSchemesResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.GetProjectsWithFieldSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsWithFieldSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | List of project ids to filter the results by. | 
 **startAt** | **int64** | The starting index of the returned projects. Base index: 0. | [default to 0]
 **maxResults** | **int32** | The maximum number of projects to return per page, maximum allowed value is 100. | [default to 50]

### Return type

[**PageBean2GetProjectsWithFieldSchemesResponse**](PageBean2GetProjectsWithFieldSchemesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFieldAssociationSchemeItemParameters

> RemoveFieldAssociationSchemeItemParameters(ctx).RequestBody(requestBody).Execute()

Remove field parameters



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
	requestBody := map[string][]openapiclient.ParameterRemovalDetails{"key": []openapiclient.ParameterRemovalDetails{*openapiclient.NewParameterRemovalDetails()}} // map[string][]ParameterRemovalDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FieldSchemesAPI.RemoveFieldAssociationSchemeItemParameters(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.RemoveFieldAssociationSchemeItemParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFieldAssociationSchemeItemParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string][]ParameterRemovalDetails**](array.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFieldsAssociatedWithSchemes

> MinimalFieldSchemeToFieldsResponse RemoveFieldsAssociatedWithSchemes(ctx).RequestBody(requestBody).Execute()

Remove fields associated with field schemes



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
	requestBody := map[string]RemoveFieldAssociationsRequestItem{"key": *openapiclient.NewRemoveFieldAssociationsRequestItem([]int64{int64(123)})} // map[string]RemoveFieldAssociationsRequestItem | The request containing the schemes and fields to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.RemoveFieldsAssociatedWithSchemes(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.RemoveFieldsAssociatedWithSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveFieldsAssociatedWithSchemes`: MinimalFieldSchemeToFieldsResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.RemoveFieldsAssociatedWithSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFieldsAssociatedWithSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string]RemoveFieldAssociationsRequestItem**](RemoveFieldAssociationsRequestItem.md) | The request containing the schemes and fields to be removed. | 

### Return type

[**MinimalFieldSchemeToFieldsResponse**](MinimalFieldSchemeToFieldsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFieldAssociationSchemeFields

> PageBean2FieldAssociationSchemeFieldSearchResult SearchFieldAssociationSchemeFields(ctx, id).StartAt(startAt).MaxResults(maxResults).FieldId(fieldId).Execute()

Search field scheme fields



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
	id := int64(789) // int64 | The scheme ID to search for child fields
	startAt := int64(789) // int64 | The starting index of the returned fields. Base index: 0. (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of fields to return per page, maximum allowed value is 100. (optional) (default to 50)
	fieldId := []string{"Inner_example"} // []string | The field IDs to filter by, if empty then all fields belonging to a field association scheme will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.SearchFieldAssociationSchemeFields(context.Background(), id).StartAt(startAt).MaxResults(maxResults).FieldId(fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.SearchFieldAssociationSchemeFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFieldAssociationSchemeFields`: PageBean2FieldAssociationSchemeFieldSearchResult
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.SearchFieldAssociationSchemeFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The scheme ID to search for child fields | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFieldAssociationSchemeFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned fields. Base index: 0. | [default to 0]
 **maxResults** | **int32** | The maximum number of fields to return per page, maximum allowed value is 100. | [default to 50]
 **fieldId** | **[]string** | The field IDs to filter by, if empty then all fields belonging to a field association scheme will be returned | 

### Return type

[**PageBean2FieldAssociationSchemeFieldSearchResult**](PageBean2FieldAssociationSchemeFieldSearchResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchFieldAssociationSchemeProjects

> PageBean2FieldAssociationSchemeProjectSearchResult SearchFieldAssociationSchemeProjects(ctx, id).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()

Search field scheme projects



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
	id := int64(789) // int64 | The scheme id to search for associated projects
	startAt := int64(789) // int64 | The starting index of the returned projects. Base index: 0. (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of projects to return per page, maximum allowed value is 100. (optional) (default to 50)
	projectId := []int64{int64(123)} // []int64 | The project Ids to filter by, if empty then all projects belonging to a field association scheme will be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.SearchFieldAssociationSchemeProjects(context.Background(), id).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.SearchFieldAssociationSchemeProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchFieldAssociationSchemeProjects`: PageBean2FieldAssociationSchemeProjectSearchResult
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.SearchFieldAssociationSchemeProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The scheme id to search for associated projects | 

### Other Parameters

Other parameters are passed through a pointer to a apiSearchFieldAssociationSchemeProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned projects. Base index: 0. | [default to 0]
 **maxResults** | **int32** | The maximum number of projects to return per page, maximum allowed value is 100. | [default to 50]
 **projectId** | **[]int64** | The project Ids to filter by, if empty then all projects belonging to a field association scheme will be returned | 

### Return type

[**PageBean2FieldAssociationSchemeProjectSearchResult**](PageBean2FieldAssociationSchemeProjectSearchResult.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldAssociationScheme

> UpdateFieldAssociationSchemeResponse UpdateFieldAssociationScheme(ctx, id).UpdateFieldAssociationSchemeRequest(updateFieldAssociationSchemeRequest).Execute()

Update field scheme



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
	id := int64(789) // int64 | 
	updateFieldAssociationSchemeRequest := *openapiclient.NewUpdateFieldAssociationSchemeRequest() // UpdateFieldAssociationSchemeRequest | The request containing the desired updates to the field association scheme

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.UpdateFieldAssociationScheme(context.Background(), id).UpdateFieldAssociationSchemeRequest(updateFieldAssociationSchemeRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.UpdateFieldAssociationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldAssociationScheme`: UpdateFieldAssociationSchemeResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.UpdateFieldAssociationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldAssociationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateFieldAssociationSchemeRequest** | [**UpdateFieldAssociationSchemeRequest**](UpdateFieldAssociationSchemeRequest.md) | The request containing the desired updates to the field association scheme | 

### Return type

[**UpdateFieldAssociationSchemeResponse**](UpdateFieldAssociationSchemeResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldAssociationSchemeItemParameters

> UpdateFieldSchemeParametersResponse UpdateFieldAssociationSchemeItemParameters(ctx).RequestBody(requestBody).Execute()

Update field parameters



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
	requestBody := map[string][]openapiclient.UpdateFieldSchemeParametersRequest{"key": []openapiclient.UpdateFieldSchemeParametersRequest{*openapiclient.NewUpdateFieldSchemeParametersRequest()}} // map[string][]UpdateFieldSchemeParametersRequest | The request containing the field association scheme id and the parameters to update.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.UpdateFieldAssociationSchemeItemParameters(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.UpdateFieldAssociationSchemeItemParameters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldAssociationSchemeItemParameters`: UpdateFieldSchemeParametersResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.UpdateFieldAssociationSchemeItemParameters`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldAssociationSchemeItemParametersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string][]UpdateFieldSchemeParametersRequest**](array.md) | The request containing the field association scheme id and the parameters to update. | 

### Return type

[**UpdateFieldSchemeParametersResponse**](UpdateFieldSchemeParametersResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFieldsAssociatedWithSchemes

> FieldSchemeToFieldsResponse UpdateFieldsAssociatedWithSchemes(ctx).RequestBody(requestBody).Execute()

Update fields associated with field schemes



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
	requestBody := map[string][]openapiclient.UpdateFieldAssociationsRequestItem{"key": []openapiclient.UpdateFieldAssociationsRequestItem{*openapiclient.NewUpdateFieldAssociationsRequestItem([]int64{int64(123)})}} // map[string][]UpdateFieldAssociationsRequestItem | The request containing the schemes and work types to associate each field with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FieldSchemesAPI.UpdateFieldsAssociatedWithSchemes(context.Background()).RequestBody(requestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FieldSchemesAPI.UpdateFieldsAssociatedWithSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFieldsAssociatedWithSchemes`: FieldSchemeToFieldsResponse
	fmt.Fprintf(os.Stdout, "Response from `FieldSchemesAPI.UpdateFieldsAssociatedWithSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldsAssociatedWithSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | [**map[string][]UpdateFieldAssociationsRequestItem**](array.md) | The request containing the schemes and work types to associate each field with. | 

### Return type

[**FieldSchemeToFieldsResponse**](FieldSchemeToFieldsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

