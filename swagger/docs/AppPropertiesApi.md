# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonPropertiesResourceDeleteAddonPropertyDelete**](AppPropertiesApi.md#AddonPropertiesResourceDeleteAddonPropertyDelete) | **Delete** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Delete app property
[**AddonPropertiesResourceGetAddonPropertiesGet**](AppPropertiesApi.md#AddonPropertiesResourceGetAddonPropertiesGet) | **Get** /rest/atlassian-connect/1/addons/{addonKey}/properties | Get app properties
[**AddonPropertiesResourceGetAddonPropertyGet**](AppPropertiesApi.md#AddonPropertiesResourceGetAddonPropertyGet) | **Get** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Get app property
[**AddonPropertiesResourcePutAddonPropertyPut**](AppPropertiesApi.md#AddonPropertiesResourcePutAddonPropertyPut) | **Put** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Set app property

# **AddonPropertiesResourceDeleteAddonPropertyDelete**
> AddonPropertiesResourceDeleteAddonPropertyDelete(ctx, addonKey, propertyKey)
Delete app property

Deletes an app's property.  **[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addonKey** | **string**| The key of the app, as defined in its descriptor. | 
  **propertyKey** | **string**| The key of the property. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonPropertiesResourceGetAddonPropertiesGet**
> PropertyKeys AddonPropertiesResourceGetAddonPropertiesGet(ctx, addonKey)
Get app properties

Gets all the properties of an app.  **[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request. Additionally, Forge apps published on the Marketplace can access properties of Connect apps they were [migrated from](https://developer.atlassian.com/platform/forge/build-a-connect-on-forge-app/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addonKey** | **string**| The key of the app, as defined in its descriptor. | 

### Return type

[**PropertyKeys**](PropertyKeys.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonPropertiesResourceGetAddonPropertyGet**
> EntityProperty AddonPropertiesResourceGetAddonPropertyGet(ctx, addonKey, propertyKey)
Get app property

Returns the key and value of an app's property.  **[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request. Additionally, Forge apps published on the Marketplace can access properties of Connect apps they were [migrated from](https://developer.atlassian.com/platform/forge/build-a-connect-on-forge-app/).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **addonKey** | **string**| The key of the app, as defined in its descriptor. | 
  **propertyKey** | **string**| The key of the property. | 

### Return type

[**EntityProperty**](EntityProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **AddonPropertiesResourcePutAddonPropertyPut**
> OperationMessage AddonPropertiesResourcePutAddonPropertyPut(ctx, body, addonKey, propertyKey)
Set app property

Sets the value of an app's property. Use this resource to store custom data for your app.  The value of the request body must be a [valid](http://tools.ietf.org/html/rfc4627), non-empty JSON blob. The maximum length is 32768 characters.  **[Permissions](#permissions) required:** Only a Connect app whose key matches `addonKey` can make this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](.md)|  | 
  **addonKey** | **string**| The key of the app, as defined in its descriptor. | 
  **propertyKey** | **string**| The key of the property. | 

### Return type

[**OperationMessage**](OperationMessage.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

