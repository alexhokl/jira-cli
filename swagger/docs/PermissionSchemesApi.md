# \PermissionSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePermissionGrant**](PermissionSchemesAPI.md#CreatePermissionGrant) | **Post** /rest/api/3/permissionscheme/{schemeId}/permission | Create permission grant
[**CreatePermissionScheme**](PermissionSchemesAPI.md#CreatePermissionScheme) | **Post** /rest/api/3/permissionscheme | Create permission scheme
[**DeletePermissionScheme**](PermissionSchemesAPI.md#DeletePermissionScheme) | **Delete** /rest/api/3/permissionscheme/{schemeId} | Delete permission scheme
[**DeletePermissionSchemeEntity**](PermissionSchemesAPI.md#DeletePermissionSchemeEntity) | **Delete** /rest/api/3/permissionscheme/{schemeId}/permission/{permissionId} | Delete permission scheme grant
[**GetAllPermissionSchemes**](PermissionSchemesAPI.md#GetAllPermissionSchemes) | **Get** /rest/api/3/permissionscheme | Get all permission schemes
[**GetPermissionScheme**](PermissionSchemesAPI.md#GetPermissionScheme) | **Get** /rest/api/3/permissionscheme/{schemeId} | Get permission scheme
[**GetPermissionSchemeGrant**](PermissionSchemesAPI.md#GetPermissionSchemeGrant) | **Get** /rest/api/3/permissionscheme/{schemeId}/permission/{permissionId} | Get permission scheme grant
[**GetPermissionSchemeGrants**](PermissionSchemesAPI.md#GetPermissionSchemeGrants) | **Get** /rest/api/3/permissionscheme/{schemeId}/permission | Get permission scheme grants
[**UpdatePermissionScheme**](PermissionSchemesAPI.md#UpdatePermissionScheme) | **Put** /rest/api/3/permissionscheme/{schemeId} | Update permission scheme



## CreatePermissionGrant

> PermissionGrant CreatePermissionGrant(ctx, schemeId).PermissionGrant(permissionGrant).Expand(expand).Execute()

Create permission grant



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
	schemeId := int64(789) // int64 | The ID of the permission scheme in which to create a new permission grant.
	permissionGrant := *openapiclient.NewPermissionGrant() // PermissionGrant | The permission grant to create.
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  `permissions` Returns all permission grants for each permission scheme.  *  `user` Returns information about the user who is granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `projectRole` Returns information about the project role granted the permission.  *  `field` Returns information about the custom field granted the permission.  *  `all` Returns all expandable information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.CreatePermissionGrant(context.Background(), schemeId).PermissionGrant(permissionGrant).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.CreatePermissionGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePermissionGrant`: PermissionGrant
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.CreatePermissionGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme in which to create a new permission grant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionGrant** | [**PermissionGrant**](PermissionGrant.md) | The permission grant to create. | 
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;user&#x60; Returns information about the user who is granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;all&#x60; Returns all expandable information. | 

### Return type

[**PermissionGrant**](PermissionGrant.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePermissionScheme

> PermissionScheme CreatePermissionScheme(ctx).PermissionScheme(permissionScheme).Expand(expand).Execute()

Create permission scheme



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
	permissionScheme := *openapiclient.NewPermissionScheme("Name_example") // PermissionScheme | The permission scheme to create.
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.CreatePermissionScheme(context.Background()).PermissionScheme(permissionScheme).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.CreatePermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePermissionScheme`: PermissionScheme
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.CreatePermissionScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePermissionSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionScheme** | [**PermissionScheme**](PermissionScheme.md) | The permission scheme to create. | 
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PermissionScheme**](PermissionScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePermissionScheme

> DeletePermissionScheme(ctx, schemeId).Execute()

Delete permission scheme



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
	schemeId := int64(789) // int64 | The ID of the permission scheme being deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionSchemesAPI.DeletePermissionScheme(context.Background(), schemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.DeletePermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme being deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionSchemeRequest struct via the builder pattern


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


## DeletePermissionSchemeEntity

> DeletePermissionSchemeEntity(ctx, schemeId, permissionId).Execute()

Delete permission scheme grant



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
	schemeId := int64(789) // int64 | The ID of the permission scheme to delete the permission grant from.
	permissionId := int64(789) // int64 | The ID of the permission grant to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.PermissionSchemesAPI.DeletePermissionSchemeEntity(context.Background(), schemeId, permissionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.DeletePermissionSchemeEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme to delete the permission grant from. | 
**permissionId** | **int64** | The ID of the permission grant to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePermissionSchemeEntityRequest struct via the builder pattern


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


## GetAllPermissionSchemes

> PermissionSchemes GetAllPermissionSchemes(ctx).Expand(expand).Execute()

Get all permission schemes



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
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.GetAllPermissionSchemes(context.Background()).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.GetAllPermissionSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPermissionSchemes`: PermissionSchemes
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.GetAllPermissionSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPermissionSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PermissionSchemes**](PermissionSchemes.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionScheme

> PermissionScheme GetPermissionScheme(ctx, schemeId).Expand(expand).Execute()

Get permission scheme



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
	schemeId := int64(789) // int64 | The ID of the permission scheme to return.
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.GetPermissionScheme(context.Background(), schemeId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.GetPermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionScheme`: PermissionScheme
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.GetPermissionScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PermissionScheme**](PermissionScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionSchemeGrant

> PermissionGrant GetPermissionSchemeGrant(ctx, schemeId, permissionId).Expand(expand).Execute()

Get permission scheme grant



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
	schemeId := int64(789) // int64 | The ID of the permission scheme.
	permissionId := int64(789) // int64 | The ID of the permission grant.
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.GetPermissionSchemeGrant(context.Background(), schemeId, permissionId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.GetPermissionSchemeGrant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionSchemeGrant`: PermissionGrant
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.GetPermissionSchemeGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme. | 
**permissionId** | **int64** | The ID of the permission grant. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionSchemeGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PermissionGrant**](PermissionGrant.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermissionSchemeGrants

> PermissionGrants GetPermissionSchemeGrants(ctx, schemeId).Expand(expand).Execute()

Get permission scheme grants



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
	schemeId := int64(789) // int64 | The ID of the permission scheme.
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  `permissions` Returns all permission grants for each permission scheme.  *  `user` Returns information about the user who is granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `projectRole` Returns information about the project role granted the permission.  *  `field` Returns information about the custom field granted the permission.  *  `all` Returns all expandable information. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.GetPermissionSchemeGrants(context.Background(), schemeId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.GetPermissionSchemeGrants``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermissionSchemeGrants`: PermissionGrants
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.GetPermissionSchemeGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPermissionSchemeGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;user&#x60; Returns information about the user who is granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;all&#x60; Returns all expandable information. | 

### Return type

[**PermissionGrants**](PermissionGrants.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePermissionScheme

> PermissionScheme UpdatePermissionScheme(ctx, schemeId).PermissionScheme(permissionScheme).Expand(expand).Execute()

Update permission scheme



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
	schemeId := int64(789) // int64 | The ID of the permission scheme to update.
	permissionScheme := *openapiclient.NewPermissionScheme("Name_example") // PermissionScheme | 
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionSchemesAPI.UpdatePermissionScheme(context.Background(), schemeId).PermissionScheme(permissionScheme).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionSchemesAPI.UpdatePermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePermissionScheme`: PermissionScheme
	fmt.Fprintf(os.Stdout, "Response from `PermissionSchemesAPI.UpdatePermissionScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **int64** | The ID of the permission scheme to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePermissionSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **permissionScheme** | [**PermissionScheme**](PermissionScheme.md) |  | 
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are always included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PermissionScheme**](PermissionScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

