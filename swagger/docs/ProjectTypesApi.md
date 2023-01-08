# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAccessibleProjectTypeByKey**](ProjectTypesApi.md#GetAccessibleProjectTypeByKey) | **Get** /rest/api/3/project/type/{projectTypeKey}/accessible | Get accessible project type by key
[**GetAllAccessibleProjectTypes**](ProjectTypesApi.md#GetAllAccessibleProjectTypes) | **Get** /rest/api/3/project/type/accessible | Get licensed project types
[**GetAllProjectTypes**](ProjectTypesApi.md#GetAllProjectTypes) | **Get** /rest/api/3/project/type | Get all project types
[**GetProjectTypeByKey**](ProjectTypesApi.md#GetProjectTypeByKey) | **Get** /rest/api/3/project/type/{projectTypeKey} | Get project type by key

# **GetAccessibleProjectTypeByKey**
> ProjectType GetAccessibleProjectTypeByKey(ctx, projectTypeKey)
Get accessible project type by key

Returns a [project type](https://confluence.atlassian.com/x/Var1Nw) if it is accessible to the user.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTypeKey** | **string**| The key of the project type. | 

### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllAccessibleProjectTypes**
> []ProjectType GetAllAccessibleProjectTypes(ctx, )
Get licensed project types

Returns all [project types](https://confluence.atlassian.com/x/Var1Nw) with a valid license.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectTypes**
> []ProjectType GetAllProjectTypes(ctx, )
Get all project types

Returns all [project types](https://confluence.atlassian.com/x/Var1Nw), whether or not the instance has a valid license for each type.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:** None.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectTypeByKey**
> ProjectType GetProjectTypeByKey(ctx, projectTypeKey)
Get project type by key

Returns a [project type](https://confluence.atlassian.com/x/Var1Nw).  This operation can be accessed anonymously.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectTypeKey** | **string**| The key of the project type. | 

### Return type

[**ProjectType**](ProjectType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

