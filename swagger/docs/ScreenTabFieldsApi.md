# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScreenTabField**](ScreenTabFieldsApi.md#AddScreenTabField) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields | Add screen tab field
[**GetAllScreenTabFields**](ScreenTabFieldsApi.md#GetAllScreenTabFields) | **Get** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields | Get all screen tab fields
[**MoveScreenTabField**](ScreenTabFieldsApi.md#MoveScreenTabField) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields/{id}/move | Move screen tab field
[**RemoveScreenTabField**](ScreenTabFieldsApi.md#RemoveScreenTabField) | **Delete** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields/{id} | Remove screen tab field

# **AddScreenTabField**
> ScreenableField AddScreenTabField(ctx, body, screenId, tabId)
Add screen tab field

Adds a field to a screen tab.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**AddFieldBean**](AddFieldBean.md)|  | 
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 

### Return type

[**ScreenableField**](ScreenableField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllScreenTabFields**
> []ScreenableField GetAllScreenTabFields(ctx, screenId, tabId, optional)
Get all screen tab fields

Returns all fields for a screen tab.  **[Permissions](#permissions) required:**   *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).  *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) when the project key is specified, providing that the screen is associated with the project through a Screen Scheme and Issue Type Screen Scheme.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 
 **optional** | ***ScreenTabFieldsApiGetAllScreenTabFieldsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScreenTabFieldsApiGetAllScreenTabFieldsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectKey** | **optional.String**| The key of the project. | 

### Return type

[**[]ScreenableField**](ScreenableField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveScreenTabField**
> Object MoveScreenTabField(ctx, body, screenId, tabId, id)
Move screen tab field

Moves a screen tab field.  If `after` and `position` are provided in the request, `position` is ignored.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MoveFieldBean**](MoveFieldBean.md)|  | 
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 
  **id** | **string**| The ID of the field. | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveScreenTabField**
> RemoveScreenTabField(ctx, screenId, tabId, id)
Remove screen tab field

Removes a field from a screen tab.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 
  **id** | **string**| The ID of the field. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

