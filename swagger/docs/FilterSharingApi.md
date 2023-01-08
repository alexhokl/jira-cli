# \FilterSharingAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSharePermission**](FilterSharingAPI.md#AddSharePermission) | **Post** /rest/api/3/filter/{id}/permission | Add share permission
[**DeleteSharePermission**](FilterSharingAPI.md#DeleteSharePermission) | **Delete** /rest/api/3/filter/{id}/permission/{permissionId} | Delete share permission
[**GetDefaultShareScope**](FilterSharingAPI.md#GetDefaultShareScope) | **Get** /rest/api/3/filter/defaultShareScope | Get default share scope
[**GetSharePermission**](FilterSharingAPI.md#GetSharePermission) | **Get** /rest/api/3/filter/{id}/permission/{permissionId} | Get share permission
[**GetSharePermissions**](FilterSharingAPI.md#GetSharePermissions) | **Get** /rest/api/3/filter/{id}/permission | Get share permissions
[**SetDefaultShareScope**](FilterSharingAPI.md#SetDefaultShareScope) | **Put** /rest/api/3/filter/defaultShareScope | Set default share scope



## AddSharePermission

> []SharePermission AddSharePermission(ctx, id).SharePermissionInputBean(sharePermissionInputBean).Execute()

Add share permission



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
	id := int64(789) // int64 | The ID of the filter.
	sharePermissionInputBean := *openapiclient.NewSharePermissionInputBean("Type_example") // SharePermissionInputBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterSharingAPI.AddSharePermission(context.Background(), id).SharePermissionInputBean(sharePermissionInputBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.AddSharePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSharePermission`: []SharePermission
	fmt.Fprintf(os.Stdout, "Response from `FilterSharingAPI.AddSharePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSharePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **sharePermissionInputBean** | [**SharePermissionInputBean**](SharePermissionInputBean.md) |  | 

### Return type

[**[]SharePermission**](SharePermission.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSharePermission

> DeleteSharePermission(ctx, id, permissionId).Execute()

Delete share permission



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
	id := int64(789) // int64 | The ID of the filter.
	permissionId := int64(789) // int64 | The ID of the share permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FilterSharingAPI.DeleteSharePermission(context.Background(), id, permissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.DeleteSharePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the filter. | 
**permissionId** | **int64** | The ID of the share permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSharePermissionRequest struct via the builder pattern


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


## GetDefaultShareScope

> DefaultShareScope GetDefaultShareScope(ctx).Execute()

Get default share scope



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterSharingAPI.GetDefaultShareScope(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.GetDefaultShareScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultShareScope`: DefaultShareScope
	fmt.Fprintf(os.Stdout, "Response from `FilterSharingAPI.GetDefaultShareScope`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultShareScopeRequest struct via the builder pattern


### Return type

[**DefaultShareScope**](DefaultShareScope.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharePermission

> SharePermission GetSharePermission(ctx, id, permissionId).Execute()

Get share permission



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
	id := int64(789) // int64 | The ID of the filter.
	permissionId := int64(789) // int64 | The ID of the share permission.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterSharingAPI.GetSharePermission(context.Background(), id, permissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.GetSharePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharePermission`: SharePermission
	fmt.Fprintf(os.Stdout, "Response from `FilterSharingAPI.GetSharePermission`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the filter. | 
**permissionId** | **int64** | The ID of the share permission. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**SharePermission**](SharePermission.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharePermissions

> []SharePermission GetSharePermissions(ctx, id).Execute()

Get share permissions



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
	id := int64(789) // int64 | The ID of the filter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterSharingAPI.GetSharePermissions(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.GetSharePermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharePermissions`: []SharePermission
	fmt.Fprintf(os.Stdout, "Response from `FilterSharingAPI.GetSharePermissions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharePermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]SharePermission**](SharePermission.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultShareScope

> DefaultShareScope SetDefaultShareScope(ctx).DefaultShareScope(defaultShareScope).Execute()

Set default share scope



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
	defaultShareScope := *openapiclient.NewDefaultShareScope("Scope_example") // DefaultShareScope | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FilterSharingAPI.SetDefaultShareScope(context.Background()).DefaultShareScope(defaultShareScope).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FilterSharingAPI.SetDefaultShareScope``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultShareScope`: DefaultShareScope
	fmt.Fprintf(os.Stdout, "Response from `FilterSharingAPI.SetDefaultShareScope`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultShareScopeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **defaultShareScope** | [**DefaultShareScope**](DefaultShareScope.md) |  | 

### Return type

[**DefaultShareScope**](DefaultShareScope.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

