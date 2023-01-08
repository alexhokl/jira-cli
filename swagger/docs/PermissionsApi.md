# \PermissionsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllPermissions**](PermissionsAPI.md#GetAllPermissions) | **Get** /rest/api/3/permissions | Get all permissions
[**GetBulkPermissions**](PermissionsAPI.md#GetBulkPermissions) | **Post** /rest/api/3/permissions/check | Get bulk permissions
[**GetMyPermissions**](PermissionsAPI.md#GetMyPermissions) | **Get** /rest/api/3/mypermissions | Get my permissions
[**GetPermittedProjects**](PermissionsAPI.md#GetPermittedProjects) | **Post** /rest/api/3/permissions/project | Get permitted projects



## GetAllPermissions

> Permissions GetAllPermissions(ctx).Execute()

Get all permissions



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
	resp, r, err := apiClient.PermissionsAPI.GetAllPermissions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetAllPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllPermissions`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetAllPermissions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllPermissionsRequest struct via the builder pattern


### Return type

[**Permissions**](Permissions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkPermissions

> BulkPermissionGrants GetBulkPermissions(ctx).BulkPermissionsRequestBean(bulkPermissionsRequestBean).Execute()

Get bulk permissions



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
	bulkPermissionsRequestBean := *openapiclient.NewBulkPermissionsRequestBean() // BulkPermissionsRequestBean | Details of the permissions to check.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.GetBulkPermissions(context.Background()).BulkPermissionsRequestBean(bulkPermissionsRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetBulkPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkPermissions`: BulkPermissionGrants
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetBulkPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkPermissionsRequestBean** | [**BulkPermissionsRequestBean**](BulkPermissionsRequestBean.md) | Details of the permissions to check. | 

### Return type

[**BulkPermissionGrants**](BulkPermissionGrants.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMyPermissions

> Permissions GetMyPermissions(ctx).ProjectKey(projectKey).ProjectId(projectId).IssueKey(issueKey).IssueId(issueId).Permissions(permissions).ProjectUuid(projectUuid).ProjectConfigurationUuid(projectConfigurationUuid).CommentId(commentId).Execute()

Get my permissions



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
	projectKey := "projectKey_example" // string | The key of project. Ignored if `projectId` is provided. (optional)
	projectId := "projectId_example" // string | The ID of project. (optional)
	issueKey := "issueKey_example" // string | The key of the issue. Ignored if `issueId` is provided. (optional)
	issueId := "issueId_example" // string | The ID of the issue. (optional)
	permissions := "BROWSE_PROJECTS,EDIT_ISSUES" // string | A list of permission keys. (Required) This parameter accepts a comma-separated list. To get the list of available permissions, use [Get all permissions](#api-rest-api-3-permissions-get). (optional)
	projectUuid := "projectUuid_example" // string |  (optional)
	projectConfigurationUuid := "projectConfigurationUuid_example" // string |  (optional)
	commentId := "commentId_example" // string | The ID of the comment. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.GetMyPermissions(context.Background()).ProjectKey(projectKey).ProjectId(projectId).IssueKey(issueKey).IssueId(issueId).Permissions(permissions).ProjectUuid(projectUuid).ProjectConfigurationUuid(projectConfigurationUuid).CommentId(commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetMyPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetMyPermissions`: Permissions
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetMyPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMyPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKey** | **string** | The key of project. Ignored if &#x60;projectId&#x60; is provided. | 
 **projectId** | **string** | The ID of project. | 
 **issueKey** | **string** | The key of the issue. Ignored if &#x60;issueId&#x60; is provided. | 
 **issueId** | **string** | The ID of the issue. | 
 **permissions** | **string** | A list of permission keys. (Required) This parameter accepts a comma-separated list. To get the list of available permissions, use [Get all permissions](#api-rest-api-3-permissions-get). | 
 **projectUuid** | **string** |  | 
 **projectConfigurationUuid** | **string** |  | 
 **commentId** | **string** | The ID of the comment. | 

### Return type

[**Permissions**](Permissions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPermittedProjects

> PermittedProjects GetPermittedProjects(ctx).PermissionsKeysBean(permissionsKeysBean).Execute()

Get permitted projects



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
	permissionsKeysBean := *openapiclient.NewPermissionsKeysBean([]string{"Permissions_example"}) // PermissionsKeysBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PermissionsAPI.GetPermittedProjects(context.Background()).PermissionsKeysBean(permissionsKeysBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PermissionsAPI.GetPermittedProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPermittedProjects`: PermittedProjects
	fmt.Fprintf(os.Stdout, "Response from `PermissionsAPI.GetPermittedProjects`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPermittedProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissionsKeysBean** | [**PermissionsKeysBean**](PermissionsKeysBean.md) |  | 

### Return type

[**PermittedProjects**](PermittedProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

