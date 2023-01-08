# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScreenTab**](ScreenTabsApi.md#AddScreenTab) | **Post** /rest/api/3/screens/{screenId}/tabs | Create screen tab
[**DeleteScreenTab**](ScreenTabsApi.md#DeleteScreenTab) | **Delete** /rest/api/3/screens/{screenId}/tabs/{tabId} | Delete screen tab
[**GetAllScreenTabs**](ScreenTabsApi.md#GetAllScreenTabs) | **Get** /rest/api/3/screens/{screenId}/tabs | Get all screen tabs
[**MoveScreenTab**](ScreenTabsApi.md#MoveScreenTab) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/move/{pos} | Move screen tab
[**RenameScreenTab**](ScreenTabsApi.md#RenameScreenTab) | **Put** /rest/api/3/screens/{screenId}/tabs/{tabId} | Update screen tab

# **AddScreenTab**
> ScreenableTab AddScreenTab(ctx, body, screenId)
Create screen tab

Creates a tab for a screen.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScreenableTab**](ScreenableTab.md)|  | 
  **screenId** | **int64**| The ID of the screen. | 

### Return type

[**ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteScreenTab**
> DeleteScreenTab(ctx, screenId, tabId)
Delete screen tab

Deletes a screen tab.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllScreenTabs**
> []ScreenableTab GetAllScreenTabs(ctx, screenId, optional)
Get all screen tabs

Returns the list of tabs for a screen.  **[Permissions](#permissions) required:**   *  *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).  *  *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg) when the project key is specified, providing that the screen is associated with the project through a Screen Scheme and Issue Type Screen Scheme.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **screenId** | **int64**| The ID of the screen. | 
 **optional** | ***ScreenTabsApiGetAllScreenTabsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ScreenTabsApiGetAllScreenTabsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectKey** | **optional.String**| The key of the project. | 

### Return type

[**[]ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MoveScreenTab**
> Object MoveScreenTab(ctx, screenId, tabId, pos)
Move screen tab

Moves a screen tab.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 
  **pos** | **int32**| The position of tab. The base index is 0. | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RenameScreenTab**
> ScreenableTab RenameScreenTab(ctx, body, screenId, tabId)
Update screen tab

Updates the name of a screen tab.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ScreenableTab**](ScreenableTab.md)|  | 
  **screenId** | **int64**| The ID of the screen. | 
  **tabId** | **int64**| The ID of the screen tab. | 

### Return type

[**ScreenableTab**](ScreenableTab.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

