# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusCategories**](WorkflowStatusCategoriesApi.md#GetStatusCategories) | **Get** /rest/api/3/statuscategory | Get all status categories
[**GetStatusCategory**](WorkflowStatusCategoriesApi.md#GetStatusCategory) | **Get** /rest/api/3/statuscategory/{idOrKey} | Get status category

# **GetStatusCategories**
> []ModelMap GetStatusCategories(ctx, )
Get all status categories

Returns a list of all status categories.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetStatusCategory**
> ModelMap GetStatusCategory(ctx, idOrKey)
Get status category

Returns a status category. Status categories provided a mechanism for categorizing [statuses](#api-rest-api-3-status-idOrName-get).  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **idOrKey** | **string**| The ID or key of the status category. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

