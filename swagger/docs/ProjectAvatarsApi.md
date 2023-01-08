# \ProjectAvatarsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectAvatar**](ProjectAvatarsAPI.md#CreateProjectAvatar) | **Post** /rest/api/3/project/{projectIdOrKey}/avatar2 | Load project avatar
[**DeleteProjectAvatar**](ProjectAvatarsAPI.md#DeleteProjectAvatar) | **Delete** /rest/api/3/project/{projectIdOrKey}/avatar/{id} | Delete project avatar
[**GetAllProjectAvatars**](ProjectAvatarsAPI.md#GetAllProjectAvatars) | **Get** /rest/api/3/project/{projectIdOrKey}/avatars | Get all project avatars
[**UpdateProjectAvatar**](ProjectAvatarsAPI.md#UpdateProjectAvatar) | **Put** /rest/api/3/project/{projectIdOrKey}/avatar | Set project avatar



## CreateProjectAvatar

> Avatar CreateProjectAvatar(ctx, projectIdOrKey).Body(body).X(x).Y(y).Size(size).Execute()

Load project avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or (case-sensitive) key of the project.
	body := interface{}(987) // interface{} | 
	x := int32(56) // int32 | The X coordinate of the top-left corner of the crop region. (optional) (default to 0)
	y := int32(56) // int32 | The Y coordinate of the top-left corner of the crop region. (optional) (default to 0)
	size := int32(56) // int32 | The length of each side of the crop region. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAvatarsAPI.CreateProjectAvatar(context.Background(), projectIdOrKey).Body(body).X(x).Y(y).Size(size).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAvatarsAPI.CreateProjectAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `ProjectAvatarsAPI.CreateProjectAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or (case-sensitive) key of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 
 **x** | **int32** | The X coordinate of the top-left corner of the crop region. | [default to 0]
 **y** | **int32** | The Y coordinate of the top-left corner of the crop region. | [default to 0]
 **size** | **int32** | The length of each side of the crop region. | [default to 0]

### Return type

[**Avatar**](Avatar.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectAvatar

> DeleteProjectAvatar(ctx, projectIdOrKey, id).Execute()

Delete project avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or (case-sensitive) key.
	id := int64(789) // int64 | The ID of the avatar.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectAvatarsAPI.DeleteProjectAvatar(context.Background(), projectIdOrKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAvatarsAPI.DeleteProjectAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or (case-sensitive) key. | 
**id** | **int64** | The ID of the avatar. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectAvatars

> ProjectAvatars GetAllProjectAvatars(ctx, projectIdOrKey).Execute()

Get all project avatars



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or (case-sensitive) key of the project.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAvatarsAPI.GetAllProjectAvatars(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAvatarsAPI.GetAllProjectAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectAvatars`: ProjectAvatars
	fmt.Fprintf(os.Stdout, "Response from `ProjectAvatarsAPI.GetAllProjectAvatars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or (case-sensitive) key of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectAvatars**](ProjectAvatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateProjectAvatar

> interface{} UpdateProjectAvatar(ctx, projectIdOrKey).Avatar(avatar).Execute()

Set project avatar



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger"
)

func main() {
	projectIdOrKey := "projectIdOrKey_example" // string | The ID or (case-sensitive) key of the project.
	avatar := *openapiclient.NewAvatar("Id_example") // Avatar | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectAvatarsAPI.UpdateProjectAvatar(context.Background(), projectIdOrKey).Avatar(avatar).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectAvatarsAPI.UpdateProjectAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateProjectAvatar`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectAvatarsAPI.UpdateProjectAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The ID or (case-sensitive) key of the project. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProjectAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **avatar** | [**Avatar**](Avatar.md) |  | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

