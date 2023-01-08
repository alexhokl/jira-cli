# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidProjectKey**](ProjectKeyAndNameValidationApi.md#GetValidProjectKey) | **Get** /rest/api/3/projectvalidate/validProjectKey | Get valid project key
[**GetValidProjectName**](ProjectKeyAndNameValidationApi.md#GetValidProjectName) | **Get** /rest/api/3/projectvalidate/validProjectName | Get valid project name
[**ValidateProjectKey**](ProjectKeyAndNameValidationApi.md#ValidateProjectKey) | **Get** /rest/api/3/projectvalidate/key | Validate project key

# **GetValidProjectKey**
> string GetValidProjectKey(ctx, optional)
Get valid project key

Validates a project key and, if the key is invalid or in use, generates a valid random string for the project key.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectKeyAndNameValidationApiGetValidProjectKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectKeyAndNameValidationApiGetValidProjectKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **optional.String**| The project key. | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetValidProjectName**
> string GetValidProjectName(ctx, name)
Get valid project name

Checks that a project name isn't in use. If the name isn't in use, the passed string is returned. If the name is in use, this operation attempts to generate a valid project name based on the one supplied, usually by adding a sequence number. If a valid project name cannot be generated, a 404 response is returned.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **name** | **string**| The project name. | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ValidateProjectKey**
> ErrorCollection ValidateProjectKey(ctx, optional)
Validate project key

Validates a project key by confirming the key is a valid string and not in use.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProjectKeyAndNameValidationApiValidateProjectKeyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectKeyAndNameValidationApiValidateProjectKeyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **optional.String**| The project key. | 

### Return type

[**ErrorCollection**](ErrorCollection.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

