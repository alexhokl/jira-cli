# \ProjectTypesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessibleProjectTypeByKey**](ProjectTypesAPI.md#GetAccessibleProjectTypeByKey) | **Get** /rest/api/3/project/type/{projectTypeKey}/accessible | Get accessible project type by key
[**GetAllAccessibleProjectTypes**](ProjectTypesAPI.md#GetAllAccessibleProjectTypes) | **Get** /rest/api/3/project/type/accessible | Get licensed project types
[**GetAllProjectTypes**](ProjectTypesAPI.md#GetAllProjectTypes) | **Get** /rest/api/3/project/type | Get all project types
[**GetProjectTypeByKey**](ProjectTypesAPI.md#GetProjectTypeByKey) | **Get** /rest/api/3/project/type/{projectTypeKey} | Get project type by key



## GetAccessibleProjectTypeByKey

> ProjectType GetAccessibleProjectTypeByKey(ctx, projectTypeKey).Execute()

Get accessible project type by key



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
	projectTypeKey := "projectTypeKey_example" // string | The key of the project type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTypesAPI.GetAccessibleProjectTypeByKey(context.Background(), projectTypeKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypesAPI.GetAccessibleProjectTypeByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAccessibleProjectTypeByKey`: ProjectType
	fmt.Fprintf(os.Stdout, "Response from `ProjectTypesAPI.GetAccessibleProjectTypeByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectTypeKey** | **string** | The key of the project type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessibleProjectTypeByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAccessibleProjectTypes

> []ProjectType GetAllAccessibleProjectTypes(ctx).Execute()

Get licensed project types



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
	resp, r, err := apiClient.ProjectTypesAPI.GetAllAccessibleProjectTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypesAPI.GetAllAccessibleProjectTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAccessibleProjectTypes`: []ProjectType
	fmt.Fprintf(os.Stdout, "Response from `ProjectTypesAPI.GetAllAccessibleProjectTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAccessibleProjectTypesRequest struct via the builder pattern


### Return type

[**[]ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectTypes

> []ProjectType GetAllProjectTypes(ctx).Execute()

Get all project types



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
	resp, r, err := apiClient.ProjectTypesAPI.GetAllProjectTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypesAPI.GetAllProjectTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectTypes`: []ProjectType
	fmt.Fprintf(os.Stdout, "Response from `ProjectTypesAPI.GetAllProjectTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectTypesRequest struct via the builder pattern


### Return type

[**[]ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectTypeByKey

> ProjectType GetProjectTypeByKey(ctx, projectTypeKey).Execute()

Get project type by key



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
	projectTypeKey := "projectTypeKey_example" // string | The key of the project type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectTypesAPI.GetProjectTypeByKey(context.Background(), projectTypeKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectTypesAPI.GetProjectTypeByKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectTypeByKey`: ProjectType
	fmt.Fprintf(os.Stdout, "Response from `ProjectTypesAPI.GetProjectTypeByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectTypeKey** | **string** | The key of the project type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectTypeByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

