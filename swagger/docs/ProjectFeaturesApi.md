# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFeaturesForProject**](ProjectFeaturesApi.md#GetFeaturesForProject) | **Get** /rest/api/3/project/{projectIdOrKey}/features | Get project features
[**ToggleFeatureForProject**](ProjectFeaturesApi.md#ToggleFeatureForProject) | **Put** /rest/api/3/project/{projectIdOrKey}/features/{featureKey} | Set project feature state

# **GetFeaturesForProject**
> ContainerForProjectFeatures GetFeaturesForProject(ctx, projectIdOrKey)
Get project features

Returns the list of features for a project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectIdOrKey** | **string**| The ID or (case-sensitive) key of the project. | 

### Return type

[**ContainerForProjectFeatures**](ContainerForProjectFeatures.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ToggleFeatureForProject**
> ContainerForProjectFeatures ToggleFeatureForProject(ctx, body, projectIdOrKey, featureKey)
Set project feature state

Sets the state of a project feature.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectFeatureState**](ProjectFeatureState.md)| Details of the feature state change. | 
  **projectIdOrKey** | **string**| The ID or (case-sensitive) key of the project. | 
  **featureKey** | **string**| The key of the feature. | 

### Return type

[**ContainerForProjectFeatures**](ContainerForProjectFeatures.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

