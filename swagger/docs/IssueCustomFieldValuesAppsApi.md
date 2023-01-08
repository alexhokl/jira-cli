# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCustomFieldValue**](IssueCustomFieldValuesAppsApi.md#UpdateCustomFieldValue) | **Put** /rest/api/3/app/field/{fieldIdOrKey}/value | Update custom field value
[**UpdateMultipleCustomFieldValues**](IssueCustomFieldValuesAppsApi.md#UpdateMultipleCustomFieldValues) | **Post** /rest/api/3/app/field/value | Update custom fields

# **UpdateCustomFieldValue**
> Object UpdateCustomFieldValue(ctx, body, fieldIdOrKey, optional)
Update custom field value

Updates the value of a custom field on one or more issues. Custom fields can only be updated by the Forge app that created them.  **[Permissions](#permissions) required:** Only the app that created the custom field can update its values with this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**CustomFieldValueUpdateDetails**](CustomFieldValueUpdateDetails.md)|  | 
  **fieldIdOrKey** | **string**| The ID or key of the custom field. For example, &#x60;customfield_10010&#x60;. | 
 **optional** | ***IssueCustomFieldValuesAppsApiUpdateCustomFieldValueOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueCustomFieldValuesAppsApiUpdateCustomFieldValueOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **generateChangelog** | **optional.**| Whether to generate a changelog for this update. | [default to true]

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateMultipleCustomFieldValues**
> Object UpdateMultipleCustomFieldValues(ctx, body, optional)
Update custom fields

Updates the value of one or more custom fields on one or more issues. Combinations of custom field and issue should be unique within the request. Custom fields can only be updated by the Forge app that created them.  **[Permissions](#permissions) required:** Only the app that created the custom field can update its values with this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**MultipleCustomFieldValuesUpdateDetails**](MultipleCustomFieldValuesUpdateDetails.md)|  | 
 **optional** | ***IssueCustomFieldValuesAppsApiUpdateMultipleCustomFieldValuesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a IssueCustomFieldValuesAppsApiUpdateMultipleCustomFieldValuesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **generateChangelog** | **optional.**| Whether to generate a changelog for this update. | [default to true]

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

