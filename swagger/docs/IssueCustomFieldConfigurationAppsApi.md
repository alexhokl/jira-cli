# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomFieldConfiguration**](IssueCustomFieldConfigurationAppsApi.md#GetCustomFieldConfiguration) | **Get** /rest/api/3/app/field/{fieldIdOrKey}/context/configuration | Get custom field configurations
[**UpdateCustomFieldConfiguration**](IssueCustomFieldConfigurationAppsApi.md#UpdateCustomFieldConfiguration) | **Put** /rest/api/3/app/field/{fieldIdOrKey}/context/configuration | Update custom field configurations

# **GetCustomFieldConfiguration**
> PageBeanContextualConfiguration GetCustomFieldConfiguration(ctx, fieldIdOrKey, optional)
Get custom field configurations

Returns a [paginated](#pagination) list of configurations for a custom field created by a [Forge app](https://developer.atlassian.com/platform/forge/).  The result can be filtered by one of these criteria:   *  `id`.  *  `fieldContextId`.  *  `issueId`.  *  `projectKeyOrId` and `issueTypeId`.  Otherwise, all configurations are returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). Jira permissions are not required for the Forge app that created the custom field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **fieldIdOrKey** | **string**| The ID or key of the custom field, for example &#x60;customfield_10000&#x60;. | 
 **optional** | ***IssueCustomFieldConfigurationAppsApiGetCustomFieldConfigurationOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueCustomFieldConfigurationAppsApiGetCustomFieldConfigurationOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | [**optional.Interface of []int64**](int64.md)| The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. Can&#x27;t be provided with &#x60;fieldContextId&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **fieldContextId** | [**optional.Interface of []int64**](int64.md)| The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: &#x60;fieldContextId&#x3D;10000&amp;fieldContextId&#x3D;10001&#x60;. Can&#x27;t be provided with &#x60;id&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **issueId** | **optional.Int64**| The ID of the issue to filter results by. If the issue doesn&#x27;t exist, an empty list is returned. Can&#x27;t be provided with &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **projectKeyOrId** | **optional.String**| The ID or key of the project to filter results by. Must be provided with &#x60;issueTypeId&#x60;. Can&#x27;t be provided with &#x60;issueId&#x60;. | 
 **issueTypeId** | **optional.String**| The ID of the issue type to filter results by. Must be provided with &#x60;projectKeyOrId&#x60;. Can&#x27;t be provided with &#x60;issueId&#x60;. | 
 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanContextualConfiguration**](PageBeanContextualConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateCustomFieldConfiguration**
> Object UpdateCustomFieldConfiguration(ctx, body, fieldIdOrKey)
Update custom field configurations

Update the configuration for contexts of a custom field created by a [Forge app](https://developer.atlassian.com/platform/forge/).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). Jira permissions are not required for the Forge app that created the custom field.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomFieldConfigurations**](CustomFieldConfigurations.md)|  | 
  **fieldIdOrKey** | **string**| The ID or key of the custom field, for example &#x60;customfield_10000&#x60;. | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

