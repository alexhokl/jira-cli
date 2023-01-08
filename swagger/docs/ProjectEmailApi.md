# \ProjectEmailAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetProjectEmail**](ProjectEmailAPI.md#GetProjectEmail) | **Get** /rest/api/3/project/{projectId}/email | Get project&#39;s sender email
[**UpdateProjectEmail**](ProjectEmailAPI.md#UpdateProjectEmail) | **Put** /rest/api/3/project/{projectId}/email | Set project&#39;s sender email



## GetProjectEmail

> ProjectEmailAddress GetProjectEmail(ctx, projectId).Execute()

Get project's sender email



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
	projectId := int64(789) // int64 | The project ID.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectEmailAPI.GetProjectEmail(context.Background(), projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectEmailAPI.GetProjectEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectEmail`: ProjectEmailAddress
	fmt.Fprintf(os.Stdout, "Response from `ProjectEmailAPI.GetProjectEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectEmailAddress**](ProjectEmailAddress.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectEmail

> interface{} UpdateProjectEmail(ctx, projectId).ProjectEmailAddress(projectEmailAddress).Execute()

Set project's sender email



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
	projectId := int64(789) // int64 | The project ID.
	projectEmailAddress := *openapiclient.NewProjectEmailAddress() // ProjectEmailAddress | The project's sender email address to be set.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectEmailAPI.UpdateProjectEmail(context.Background(), projectId).ProjectEmailAddress(projectEmailAddress).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectEmailAPI.UpdateProjectEmail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectEmail`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectEmailAPI.UpdateProjectEmail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectId** | **int64** | The project ID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectEmailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectEmailAddress** | [**ProjectEmailAddress**](ProjectEmailAddress.md) | The project&#39;s sender email address to be set. | 

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

