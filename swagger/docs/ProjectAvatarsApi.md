# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectAvatar**](ProjectAvatarsApi.md#CreateProjectAvatar) | **Post** /rest/api/3/project/{projectIdOrKey}/avatar2 | Load project avatar
[**DeleteProjectAvatar**](ProjectAvatarsApi.md#DeleteProjectAvatar) | **Delete** /rest/api/3/project/{projectIdOrKey}/avatar/{id} | Delete project avatar
[**GetAllProjectAvatars**](ProjectAvatarsApi.md#GetAllProjectAvatars) | **Get** /rest/api/3/project/{projectIdOrKey}/avatars | Get all project avatars
[**UpdateProjectAvatar**](ProjectAvatarsApi.md#UpdateProjectAvatar) | **Put** /rest/api/3/project/{projectIdOrKey}/avatar | Set project avatar

# **CreateProjectAvatar**
> ModelMap CreateProjectAvatar(ctx, body, projectIdOrKey, optional)
Load project avatar

Loads an avatar for a project.  Specify the avatar's local file location in the body of the request. Also, include the following headers:   *  `X-Atlassian-Token: no-check` To prevent XSRF protection blocking the request, for more information see [Special Headers](#special-request-headers).  *  `Content-Type: image/image type` Valid image types are JPEG, GIF, or PNG.  For example:   `curl --request POST `  `--user email@example.com:<api_token> `  `--header 'X-Atlassian-Token: no-check' `  `--header 'Content-Type: image/< image_type>' `  `--data-binary \"<@/path/to/file/with/your/avatar>\" `  `--url 'https://your-domain.atlassian.net/rest/api/3/project/{projectIdOrKey}/avatar2'`  The avatar is cropped to a square. If no crop parameters are specified, the square originates at the top left of the image. The length of the square's sides is set to the smaller of the height or width of the image.  The cropped image is then used to create avatars of 16x16, 24x24, 32x32, and 48x48 in size.  After creating the avatar use [Set project avatar](#api-rest-api-3-project-projectIdOrKey-avatar-put) to set it as the project's displayed avatar.  **[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](.md)|  | 
  **projectIdOrKey** | **string**| The ID or (case-sensitive) key of the project. | 
 **optional** | ***ProjectAvatarsApiCreateProjectAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProjectAvatarsApiCreateProjectAvatarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **x** | **optional.**| The X coordinate of the top-left corner of the crop region. | [default to 0]
 **y** | **optional.**| The Y coordinate of the top-left corner of the crop region. | [default to 0]
 **size** | **optional.**| The length of each side of the crop region. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteProjectAvatar**
> DeleteProjectAvatar(ctx, projectIdOrKey, id)
Delete project avatar

Deletes a custom avatar from a project. Note that system avatars cannot be deleted.  **[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectIdOrKey** | **string**| The project ID or (case-sensitive) key. | 
  **id** | **int64**| The ID of the avatar. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllProjectAvatars**
> ProjectAvatars GetAllProjectAvatars(ctx, projectIdOrKey)
Get all project avatars

Returns all project avatars, grouped by system and custom avatars.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:** *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **projectIdOrKey** | **string**| The ID or (case-sensitive) key of the project. | 

### Return type

[**ProjectAvatars**](ProjectAvatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateProjectAvatar**
> Object UpdateProjectAvatar(ctx, body, projectIdOrKey)
Set project avatar

Sets the avatar displayed for a project.  Use [Load project avatar](#api-rest-api-3-project-projectIdOrKey-avatar2-post) to store avatars against the project, before using this operation to set the displayed avatar.  **[Permissions](#permissions) required:** *Administer projects* [project permission](https://confluence.atlassian.com/x/yodKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
  **projectIdOrKey** | **string**| The ID or (case-sensitive) key of the project. | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

