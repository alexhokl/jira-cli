# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectCategory**](ProjectCategoriesApi.md#CreateProjectCategory) | **Post** /rest/api/3/projectCategory | Create project category
[**GetAllProjectCategories**](ProjectCategoriesApi.md#GetAllProjectCategories) | **Get** /rest/api/3/projectCategory | Get all project categories
[**GetProjectCategoryById**](ProjectCategoriesApi.md#GetProjectCategoryById) | **Get** /rest/api/3/projectCategory/{id} | Get project category by ID
[**RemoveProjectCategory**](ProjectCategoriesApi.md#RemoveProjectCategory) | **Delete** /rest/api/3/projectCategory/{id} | Delete project category
[**UpdateProjectCategory**](ProjectCategoriesApi.md#UpdateProjectCategory) | **Put** /rest/api/3/projectCategory/{id} | Update project category

# **CreateProjectCategory**
> ProjectCategory CreateProjectCategory(ctx, body)
Create project category

Creates a project category.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectCategory**](ProjectCategory.md)|  | 

### Return type

[**ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectCategories**
> []ProjectCategory GetAllProjectCategories(ctx, )
Get all project categories

Returns all project categories.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetProjectCategoryById**
> ProjectCategory GetProjectCategoryById(ctx, id)
Get project category by ID

Returns a project category.  **[Permissions](#permissions) required:** Permission to access Jira.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| The ID of the project category. | 

### Return type

[**ProjectCategory**](ProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RemoveProjectCategory**
> RemoveProjectCategory(ctx, id)
Delete project category

Deletes a project category.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **id** | **int64**| ID of the project category to delete. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectCategory**
> UpdatedProjectCategory UpdateProjectCategory(ctx, body, id)
Update project category

Updates a project category.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ProjectCategory**](ProjectCategory.md)|  | 
  **id** | **int64**|  | 

### Return type

[**UpdatedProjectCategory**](UpdatedProjectCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

