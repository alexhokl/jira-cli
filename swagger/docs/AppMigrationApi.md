# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut**](AppMigrationApi.md#AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut) | **Put** /rest/atlassian-connect/1/migration/field | Bulk update custom field value
[**MigrationResourceUpdateEntityPropertiesValuePut**](AppMigrationApi.md#MigrationResourceUpdateEntityPropertiesValuePut) | **Put** /rest/atlassian-connect/1/migration/properties/{entityType} | Bulk update entity properties
[**MigrationResourceWorkflowRuleSearchPost**](AppMigrationApi.md#MigrationResourceWorkflowRuleSearchPost) | **Post** /rest/atlassian-connect/1/migration/workflow/rule/search | Get workflow transition rule configurations

# **AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut**
> Object AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut(ctx, body, atlassianTransferId)
Bulk update custom field value

Updates the value of a custom field added by Connect apps on one or more issues. The values of up to 200 custom fields can be updated.  **[Permissions](#permissions) required:** Only Connect apps can make this request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ConnectCustomFieldValues**](ConnectCustomFieldValues.md)|  | 
  **atlassianTransferId** | [**string**](.md)| The ID of the transfer. | 

### Return type

**Object**

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrationResourceUpdateEntityPropertiesValuePut**
> MigrationResourceUpdateEntityPropertiesValuePut(ctx, body, atlassianTransferId, entityType)
Bulk update entity properties

Updates the values of multiple entity properties for an object, up to 50 updates per request. This operation is for use by Connect apps during app migration.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**[]EntityPropertyDetails**](EntityPropertyDetails.md)|  | 
  **atlassianTransferId** | [**string**](.md)| The app migration transfer ID. | 
  **entityType** | **string**| The type indicating the object that contains the entity properties. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **MigrationResourceWorkflowRuleSearchPost**
> WorkflowRulesSearchDetails MigrationResourceWorkflowRuleSearchPost(ctx, body, atlassianTransferId)
Get workflow transition rule configurations

Returns configurations for workflow transition rules migrated from server to cloud and owned by the calling Connect app.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WorkflowRulesSearch**](WorkflowRulesSearch.md)|  | 
  **atlassianTransferId** | [**string**](.md)| The app migration transfer ID. | 

### Return type

[**WorkflowRulesSearchDetails**](WorkflowRulesSearchDetails.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

