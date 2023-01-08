# \ProjectFeaturesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeaturesForProject**](ProjectFeaturesAPI.md#GetFeaturesForProject) | **Get** /rest/api/3/project/{projectIdOrKey}/features | Get project features
[**ToggleFeatureForProject**](ProjectFeaturesAPI.md#ToggleFeatureForProject) | **Put** /rest/api/3/project/{projectIdOrKey}/features/{featureKey} | Set project feature state



## GetFeaturesForProject

> ContainerForProjectFeatures GetFeaturesForProject(ctx, projectIdOrKey).Execute()

Get project features



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
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or (case-sensitive) key of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectFeaturesAPI.GetFeaturesForProject(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFeaturesAPI.GetFeaturesForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturesForProject`: ContainerForProjectFeatures
	fmt.Fprintf(os.Stdout, "Response from `ProjectFeaturesAPI.GetFeaturesForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or (case-sensitive) key of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ContainerForProjectFeatures**](ContainerForProjectFeatures.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFeatureForProject

> ContainerForProjectFeatures ToggleFeatureForProject(ctx, projectIdOrKey, featureKey).ProjectFeatureState(projectFeatureState).Execute()

Set project feature state



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
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or (case-sensitive) key of the project.
	featureKey := "featureKey_example" // string | The key of the feature.
	projectFeatureState := *openapiclient.NewProjectFeatureState() // ProjectFeatureState | Details of the feature state change.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectFeaturesAPI.ToggleFeatureForProject(context.Background(), projectIdOrKey, featureKey).ProjectFeatureState(projectFeatureState).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectFeaturesAPI.ToggleFeatureForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFeatureForProject`: ContainerForProjectFeatures
	fmt.Fprintf(os.Stdout, "Response from `ProjectFeaturesAPI.ToggleFeatureForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or (case-sensitive) key of the project. | 
**featureKey** | **string** | The key of the feature. | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFeatureForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectFeatureState** | [**ProjectFeatureState**](ProjectFeatureState.md) | Details of the feature state change. | 

### Return type

[**ContainerForProjectFeatures**](ContainerForProjectFeatures.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

