# \ProjectClassificationLevelsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDefaultProjectClassification**](ProjectClassificationLevelsAPI.md#GetDefaultProjectClassification) | **Get** /rest/api/3/project/{projectIdOrKey}/classification-level/default | Get the default data classification level of a project
[**RemoveDefaultProjectClassification**](ProjectClassificationLevelsAPI.md#RemoveDefaultProjectClassification) | **Delete** /rest/api/3/project/{projectIdOrKey}/classification-level/default | Remove the default data classification level from a project
[**UpdateDefaultProjectClassification**](ProjectClassificationLevelsAPI.md#UpdateDefaultProjectClassification) | **Put** /rest/api/3/project/{projectIdOrKey}/classification-level/default | Update the default data classification level of a project



## GetDefaultProjectClassification

> interface{} GetDefaultProjectClassification(ctx, projectIdOrKey).Execute()

Get the default data classification level of a project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case-sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectClassificationLevelsAPI.GetDefaultProjectClassification(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectClassificationLevelsAPI.GetDefaultProjectClassification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultProjectClassification`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectClassificationLevelsAPI.GetDefaultProjectClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case-sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultProjectClassificationRequest struct via the builder pattern


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


## RemoveDefaultProjectClassification

> interface{} RemoveDefaultProjectClassification(ctx, projectIdOrKey).Execute()

Remove the default data classification level from a project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case-sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectClassificationLevelsAPI.RemoveDefaultProjectClassification(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectClassificationLevelsAPI.RemoveDefaultProjectClassification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveDefaultProjectClassification`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectClassificationLevelsAPI.RemoveDefaultProjectClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case-sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveDefaultProjectClassificationRequest struct via the builder pattern


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


## UpdateDefaultProjectClassification

> interface{} UpdateDefaultProjectClassification(ctx, projectIdOrKey).UpdateDefaultProjectClassificationBean(updateDefaultProjectClassificationBean).Execute()

Update the default data classification level of a project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case-sensitive).
	updateDefaultProjectClassificationBean := *openapiclient.NewUpdateDefaultProjectClassificationBean("Id_example") // UpdateDefaultProjectClassificationBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectClassificationLevelsAPI.UpdateDefaultProjectClassification(context.Background(), projectIdOrKey).UpdateDefaultProjectClassificationBean(updateDefaultProjectClassificationBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectClassificationLevelsAPI.UpdateDefaultProjectClassification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultProjectClassification`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectClassificationLevelsAPI.UpdateDefaultProjectClassification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case-sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultProjectClassificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDefaultProjectClassificationBean** | [**UpdateDefaultProjectClassificationBean**](UpdateDefaultProjectClassificationBean.md) |  | 

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

