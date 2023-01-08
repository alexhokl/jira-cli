# \ProjectCategoriesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectCategory**](ProjectCategoriesAPI.md#CreateProjectCategory) | **Post** /rest/api/3/projectCategory | Create project category
[**GetAllProjectCategories**](ProjectCategoriesAPI.md#GetAllProjectCategories) | **Get** /rest/api/3/projectCategory | Get all project categories
[**GetProjectCategoryById**](ProjectCategoriesAPI.md#GetProjectCategoryById) | **Get** /rest/api/3/projectCategory/{id} | Get project category by ID
[**RemoveProjectCategory**](ProjectCategoriesAPI.md#RemoveProjectCategory) | **Delete** /rest/api/3/projectCategory/{id} | Delete project category
[**UpdateProjectCategory**](ProjectCategoriesAPI.md#UpdateProjectCategory) | **Put** /rest/api/3/projectCategory/{id} | Update project category



## CreateProjectCategory

> ProjectCategory CreateProjectCategory(ctx).ProjectCategory(projectCategory).Execute()

Create project category



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
	projectCategory := *openapiclient.NewProjectCategory() // ProjectCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectCategoriesAPI.CreateProjectCategory(context.Background()).ProjectCategory(projectCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesAPI.CreateProjectCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectCategory`: ProjectCategory
	fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesAPI.CreateProjectCategory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectCategory** | [**ProjectCategory**](ProjectCategory.md) |  | 

### Return type

[**ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectCategories

> []ProjectCategory GetAllProjectCategories(ctx).Execute()

Get all project categories



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
	resp, r, err := apiClient.ProjectCategoriesAPI.GetAllProjectCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesAPI.GetAllProjectCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectCategories`: []ProjectCategory
	fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesAPI.GetAllProjectCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectCategoriesRequest struct via the builder pattern


### Return type

[**[]ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectCategoryById

> ProjectCategory GetProjectCategoryById(ctx, id).Execute()

Get project category by ID



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
	id := int64(789) // int64 | The ID of the project category.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectCategoriesAPI.GetProjectCategoryById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesAPI.GetProjectCategoryById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectCategoryById`: ProjectCategory
	fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesAPI.GetProjectCategoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectCategoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveProjectCategory

> RemoveProjectCategory(ctx, id).Execute()

Delete project category



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
	id := int64(789) // int64 | ID of the project category to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectCategoriesAPI.RemoveProjectCategory(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesAPI.RemoveProjectCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | ID of the project category to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveProjectCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## UpdateProjectCategory

> UpdatedProjectCategory UpdateProjectCategory(ctx, id).ProjectCategory(projectCategory).Execute()

Update project category



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
	projectCategory := *openapiclient.NewProjectCategory() // ProjectCategory | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectCategoriesAPI.UpdateProjectCategory(context.Background(), id).ProjectCategory(projectCategory).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectCategoriesAPI.UpdateProjectCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectCategory`: UpdatedProjectCategory
	fmt.Fprintf(os.Stdout, "Response from `ProjectCategoriesAPI.UpdateProjectCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectCategory** | [**ProjectCategory**](ProjectCategory.md) |  | 

### Return type

[**UpdatedProjectCategory**](UpdatedProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

