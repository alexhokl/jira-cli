# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicModulesResourceGetModulesGet**](DynamicModulesApi.md#DynamicModulesResourceGetModulesGet) | **Get** /rest/atlassian-connect/1/app/module/dynamic | Get modules
[**DynamicModulesResourceRegisterModulesPost**](DynamicModulesApi.md#DynamicModulesResourceRegisterModulesPost) | **Post** /rest/atlassian-connect/1/app/module/dynamic | Register modules
[**DynamicModulesResourceRemoveModulesDelete**](DynamicModulesApi.md#DynamicModulesResourceRemoveModulesDelete) | **Delete** /rest/atlassian-connect/1/app/module/dynamic | Remove modules

# **DynamicModulesResourceGetModulesGet**
> ConnectModules DynamicModulesResourceGetModulesGet(ctx, )
Get modules

Returns all modules registered dynamically by the calling app.  **[Permissions](#permissions) required:** Only Connect apps can make this request.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**ConnectModules**](ConnectModules.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicModulesResourceRegisterModulesPost**
> DynamicModulesResourceRegisterModulesPost(ctx, body)
Register modules

Registers a list of modules.  **[Permissions](#permissions) required:** Only Connect apps can make this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectModules**](ConnectModules.md)|  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DynamicModulesResourceRemoveModulesDelete**
> DynamicModulesResourceRemoveModulesDelete(ctx, optional)
Remove modules

Remove all or a list of modules registered by the calling app.  **[Permissions](#permissions) required:** Only Connect apps can make this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DynamicModulesApiDynamicModulesResourceRemoveModulesDeleteOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DynamicModulesApiDynamicModulesResourceRemoveModulesDeleteOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moduleKey** | [**optional.Interface of []string**](string.md)| The key of the module to remove. To include multiple module keys, provide multiple copies of this parameter. For example, &#x60;moduleKey&#x3D;dynamic-attachment-entity-property&amp;moduleKey&#x3D;dynamic-select-field&#x60;. Nonexistent keys are ignored. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

