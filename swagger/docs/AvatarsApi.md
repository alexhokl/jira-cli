# \AvatarsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteAvatar**](AvatarsAPI.md#DeleteAvatar) | **Delete** /rest/api/3/universal_avatar/type/{type}/owner/{owningObjectId}/avatar/{id} | Delete avatar
[**GetAllSystemAvatars**](AvatarsAPI.md#GetAllSystemAvatars) | **Get** /rest/api/3/avatar/{type}/system | Get system avatars by type
[**GetAvatarImageByID**](AvatarsAPI.md#GetAvatarImageByID) | **Get** /rest/api/3/universal_avatar/view/type/{type}/avatar/{id} | Get avatar image by ID
[**GetAvatarImageByOwner**](AvatarsAPI.md#GetAvatarImageByOwner) | **Get** /rest/api/3/universal_avatar/view/type/{type}/owner/{entityId} | Get avatar image by owner
[**GetAvatarImageByType**](AvatarsAPI.md#GetAvatarImageByType) | **Get** /rest/api/3/universal_avatar/view/type/{type} | Get avatar image by type
[**GetAvatars**](AvatarsAPI.md#GetAvatars) | **Get** /rest/api/3/universal_avatar/type/{type}/owner/{entityId} | Get avatars
[**StoreAvatar**](AvatarsAPI.md#StoreAvatar) | **Post** /rest/api/3/universal_avatar/type/{type}/owner/{entityId} | Load avatar



## DeleteAvatar

> DeleteAvatar(ctx, type_, owningObjectId, id).Execute()

Delete avatar



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
	type_ := "type__example" // string | The avatar type.
	owningObjectId := "owningObjectId_example" // string | The ID of the item the avatar is associated with.
	id := int64(789) // int64 | The ID of the avatar.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarsAPI.DeleteAvatar(context.Background(), type_, owningObjectId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.DeleteAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The avatar type. | 
**owningObjectId** | **string** | The ID of the item the avatar is associated with. | 
**id** | **int64** | The ID of the avatar. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAvatarRequest struct via the builder pattern


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


## GetAllSystemAvatars

> SystemAvatars GetAllSystemAvatars(ctx, type_).Execute()

Get system avatars by type



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
	type_ := "project" // string | The avatar type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetAllSystemAvatars(context.Background(), type_).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAllSystemAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllSystemAvatars`: SystemAvatars
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetAllSystemAvatars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The avatar type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSystemAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SystemAvatars**](SystemAvatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatarImageByID

> GetAvatarImageByID(ctx, type_, id).Size(size).Format(format).Execute()

Get avatar image by ID



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
	type_ := "type__example" // string | The icon type of the avatar.
	id := int64(789) // int64 | The ID of the avatar.
	size := "size_example" // string | The size of the avatar image. If not provided the default size is returned. (optional)
	format := "format_example" // string | The format to return the avatar image in. If not provided the original content format is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarsAPI.GetAvatarImageByID(context.Background(), type_, id).Size(size).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatarImageByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The icon type of the avatar. | 
**id** | **int64** | The ID of the avatar. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarImageByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **string** | The size of the avatar image. If not provided the default size is returned. | 
 **format** | **string** | The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatarImageByOwner

> GetAvatarImageByOwner(ctx, type_, entityId).Size(size).Format(format).Execute()

Get avatar image by owner



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
	type_ := "type__example" // string | The icon type of the avatar.
	entityId := "entityId_example" // string | The ID of the project or issue type the avatar belongs to.
	size := "size_example" // string | The size of the avatar image. If not provided the default size is returned. (optional)
	format := "format_example" // string | The format to return the avatar image in. If not provided the original content format is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarsAPI.GetAvatarImageByOwner(context.Background(), type_, entityId).Size(size).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatarImageByOwner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The icon type of the avatar. | 
**entityId** | **string** | The ID of the project or issue type the avatar belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarImageByOwnerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **string** | The size of the avatar image. If not provided the default size is returned. | 
 **format** | **string** | The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatarImageByType

> GetAvatarImageByType(ctx, type_).Size(size).Format(format).Execute()

Get avatar image by type



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
	type_ := "type__example" // string | The icon type of the avatar.
	size := "size_example" // string | The size of the avatar image. If not provided the default size is returned. (optional)
	format := "format_example" // string | The format to return the avatar image in. If not provided the original content format is returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AvatarsAPI.GetAvatarImageByType(context.Background(), type_).Size(size).Format(format).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatarImageByType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The icon type of the avatar. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarImageByTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **string** | The size of the avatar image. If not provided the default size is returned. | 
 **format** | **string** | The format to return the avatar image in. If not provided the original content format is returned. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: */*, application/json, image/png, image/svg+xml

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAvatars

> Avatars GetAvatars(ctx, type_, entityId).Execute()

Get avatars



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
	type_ := "type__example" // string | The avatar type.
	entityId := "entityId_example" // string | The ID of the item the avatar is associated with.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.GetAvatars(context.Background(), type_, entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.GetAvatars``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvatars`: Avatars
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.GetAvatars`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The avatar type. | 
**entityId** | **string** | The ID of the item the avatar is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvatarsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**Avatars**](Avatars.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreAvatar

> Avatar StoreAvatar(ctx, type_, entityId).Size(size).Body(body).X(x).Y(y).Execute()

Load avatar



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
	type_ := "type__example" // string | The avatar type.
	entityId := "entityId_example" // string | The ID of the item the avatar is associated with.
	size := int32(56) // int32 | The length of each side of the crop region. (default to 0)
	body := interface{}(987) // interface{} | 
	x := int32(56) // int32 | The X coordinate of the top-left corner of the crop region. (optional) (default to 0)
	y := int32(56) // int32 | The Y coordinate of the top-left corner of the crop region. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AvatarsAPI.StoreAvatar(context.Background(), type_, entityId).Size(size).Body(body).X(x).Y(y).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AvatarsAPI.StoreAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `AvatarsAPI.StoreAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**type_** | **string** | The avatar type. | 
**entityId** | **string** | The ID of the item the avatar is associated with. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStoreAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **int32** | The length of each side of the crop region. | [default to 0]
 **body** | **interface{}** |  | 
 **x** | **int32** | The X coordinate of the top-left corner of the crop region. | [default to 0]
 **y** | **int32** | The Y coordinate of the top-left corner of the crop region. | [default to 0]

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

