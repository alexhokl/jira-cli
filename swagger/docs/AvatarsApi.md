# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAvatar**](AvatarsApi.md#DeleteAvatar) | **Delete** /rest/api/3/universal_avatar/type/{type}/owner/{owningObjectId}/avatar/{id} | Delete avatar
[**GetAllSystemAvatars**](AvatarsApi.md#GetAllSystemAvatars) | **Get** /rest/api/3/avatar/{type}/system | Get system avatars by type
[**GetAvatarImageByID**](AvatarsApi.md#GetAvatarImageByID) | **Get** /rest/api/3/universal_avatar/view/type/{type}/avatar/{id} | Get avatar image by ID
[**GetAvatarImageByOwner**](AvatarsApi.md#GetAvatarImageByOwner) | **Get** /rest/api/3/universal_avatar/view/type/{type}/owner/{entityId} | Get avatar image by owner
[**GetAvatarImageByType**](AvatarsApi.md#GetAvatarImageByType) | **Get** /rest/api/3/universal_avatar/view/type/{type} | Get avatar image by type
[**GetAvatars**](AvatarsApi.md#GetAvatars) | **Get** /rest/api/3/universal_avatar/type/{type}/owner/{entityId} | Get avatars
[**StoreAvatar**](AvatarsApi.md#StoreAvatar) | **Post** /rest/api/3/universal_avatar/type/{type}/owner/{entityId} | Load avatar

# **DeleteAvatar**
> DeleteAvatar(ctx, type_, owningObjectId, id)
Delete avatar

Deletes an avatar from a project or issue type.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The avatar type. | 
  **owningObjectId** | **string**| The ID of the item the avatar is associated with. | 
  **id** | **int64**| The ID of the avatar. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAllSystemAvatars**
> SystemAvatars GetAllSystemAvatars(ctx, type_)
Get system avatars by type

Returns a list of system avatar details by owner type, where the owner types are issue type, project, or user.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The avatar type. | 

### Return type

[**SystemAvatars**](SystemAvatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatarImageByID**
> GetAvatarImageByID(ctx, type_, id, optional)
Get avatar image by ID

Returns a project or issue type avatar image by ID.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:**   *  For system avatars, none.  *  For custom project avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the avatar belongs to.  *  For custom issue type avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for at least one project the issue type is used in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The icon type of the avatar. | 
  **id** | **int64**| The ID of the avatar. | 
 **optional** | ***AvatarsApiGetAvatarImageByIDOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvatarsApiGetAvatarImageByIDOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.String**| The size of the avatar image. If not provided the default size is returned. | 
 **format** | **optional.String**| The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatarImageByOwner**
> GetAvatarImageByOwner(ctx, type_, entityId, optional)
Get avatar image by owner

Returns the avatar image for a project or issue type.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:**   *  For system avatars, none.  *  For custom project avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the avatar belongs to.  *  For custom issue type avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for at least one project the issue type is used in.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The icon type of the avatar. | 
  **entityId** | **string**| The ID of the project or issue type the avatar belongs to. | 
 **optional** | ***AvatarsApiGetAvatarImageByOwnerOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvatarsApiGetAvatarImageByOwnerOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.String**| The size of the avatar image. If not provided the default size is returned. | 
 **format** | **optional.String**| The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatarImageByType**
> GetAvatarImageByType(ctx, type_, optional)
Get avatar image by type

Returns the default project or issue type avatar image.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:** None.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The icon type of the avatar. | 
 **optional** | ***AvatarsApiGetAvatarImageByTypeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvatarsApiGetAvatarImageByTypeOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.String**| The size of the avatar image. If not provided the default size is returned. | 
 **format** | **optional.String**| The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetAvatars**
> Avatars GetAvatars(ctx, type_, entityId)
Get avatars

Returns the system and custom avatars for a project or issue type.  This operation can be accessed anonymously.  **[Permissions](#permissions) required:**   *  for custom project avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for the project the avatar belongs to.  *  for custom issue type avatars, *Browse projects* [project permission](https://confluence.atlassian.com/x/yodKLg) for at least one project the issue type is used in.  *  for system avatars, none.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **type_** | **string**| The avatar type. | 
  **entityId** | **string**| The ID of the item the avatar is associated with. | 

### Return type

[**Avatars**](Avatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StoreAvatar**
> ModelMap StoreAvatar(ctx, body, type_, entityId, size, optional)
Load avatar

Loads a custom avatar for a project or issue type.  Specify the avatar's local file location in the body of the request. Also, include the following headers:   *  `X-Atlassian-Token: no-check` To prevent XSRF protection blocking the request, for more information see [Special Headers](#special-request-headers).  *  `Content-Type: image/image type` Valid image types are JPEG, GIF, or PNG.  For example:   `curl --request POST `  `--user email@example.com:<api_token> `  `--header 'X-Atlassian-Token: no-check' `  `--header 'Content-Type: image/< image_type>' `  `--data-binary \"<@/path/to/file/with/your/avatar>\" `  `--url 'https://your-domain.atlassian.net/rest/api/3/universal_avatar/type/{type}/owner/{entityId}'`  The avatar is cropped to a square. If no crop parameters are specified, the square originates at the top left of the image. The length of the square's sides is set to the smaller of the height or width of the image.  The cropped image is then used to create avatars of 16x16, 24x24, 32x32, and 48x48 in size.  After creating the avatar use:   *  [Update issue type](#api-rest-api-3-issuetype-id-put) to set it as the issue type's displayed avatar.  *  [Set project avatar](#api-rest-api-3-project-projectIdOrKey-avatar-put) to set it as the project's displayed avatar.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**Object**](.md)|  | 
  **type_** | **string**| The avatar type. | 
  **entityId** | **string**| The ID of the item the avatar is associated with. | 
  **size** | **int32**| The length of each side of the crop region. | 
 **optional** | ***AvatarsApiStoreAvatarOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AvatarsApiStoreAvatarOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **x** | **optional.**| The X coordinate of the top-left corner of the crop region. | [default to 0]
 **y** | **optional.**| The Y coordinate of the top-left corner of the crop region. | [default to 0]

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: */*
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

