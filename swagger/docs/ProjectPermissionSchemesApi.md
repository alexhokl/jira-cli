# \ProjectPermissionSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignPermissionScheme**](ProjectPermissionSchemesAPI.md#AssignPermissionScheme) | **Put** /rest/api/3/project/{projectKeyOrId}/permissionscheme | Assign permission scheme
[**GetAssignedPermissionScheme**](ProjectPermissionSchemesAPI.md#GetAssignedPermissionScheme) | **Get** /rest/api/3/project/{projectKeyOrId}/permissionscheme | Get assigned permission scheme
[**GetProjectIssueSecurityScheme**](ProjectPermissionSchemesAPI.md#GetProjectIssueSecurityScheme) | **Get** /rest/api/3/project/{projectKeyOrId}/issuesecuritylevelscheme | Get project issue security scheme
[**GetSecurityLevelsForProject**](ProjectPermissionSchemesAPI.md#GetSecurityLevelsForProject) | **Get** /rest/api/3/project/{projectKeyOrId}/securitylevel | Get project issue security levels



## AssignPermissionScheme

> PermissionScheme AssignPermissionScheme(ctx, projectKeyOrId).IdBean(idBean).Expand(expand).Execute()

Assign permission scheme



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
	projectKeyOrId := "projectKeyOrId_example" // string | The project ID or project key (case sensitive).
	idBean := *openapiclient.NewIdBean(int64(123)) // IdBean | 
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPermissionSchemesAPI.AssignPermissionScheme(context.Background(), projectKeyOrId).IdBean(idBean).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPermissionSchemesAPI.AssignPermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignPermissionScheme`: PermissionScheme
	fmt.Fprintf(os.Stdout, "Response from `ProjectPermissionSchemesAPI.AssignPermissionScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKeyOrId** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignPermissionSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **idBean** | [**IdBean**](IdBean.md) |  | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

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


## GetAssignedPermissionScheme

> PermissionScheme GetAssignedPermissionScheme(ctx, projectKeyOrId).Expand(expand).Execute()

Get assigned permission scheme



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
	projectKeyOrId := "projectKeyOrId_example" // string | The project ID or project key (case sensitive).
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `permissions` Returns all permission grants for each permission scheme.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPermissionSchemesAPI.GetAssignedPermissionScheme(context.Background(), projectKeyOrId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPermissionSchemesAPI.GetAssignedPermissionScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAssignedPermissionScheme`: PermissionScheme
	fmt.Fprintf(os.Stdout, "Response from `ProjectPermissionSchemesAPI.GetAssignedPermissionScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKeyOrId** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssignedPermissionSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Note that permissions are included when you specify any value. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;permissions&#x60; Returns all permission grants for each permission scheme.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

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


## GetProjectIssueSecurityScheme

> SecurityScheme GetProjectIssueSecurityScheme(ctx, projectKeyOrId).Execute()

Get project issue security scheme



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
	projectKeyOrId := "projectKeyOrId_example" // string | The project ID or project key (case sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPermissionSchemesAPI.GetProjectIssueSecurityScheme(context.Background(), projectKeyOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPermissionSchemesAPI.GetProjectIssueSecurityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectIssueSecurityScheme`: SecurityScheme
	fmt.Fprintf(os.Stdout, "Response from `ProjectPermissionSchemesAPI.GetProjectIssueSecurityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKeyOrId** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectIssueSecuritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityScheme**](SecurityScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityLevelsForProject

> ProjectIssueSecurityLevels GetSecurityLevelsForProject(ctx, projectKeyOrId).Execute()

Get project issue security levels



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
	projectKeyOrId := "projectKeyOrId_example" // string | The project ID or project key (case sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPermissionSchemesAPI.GetSecurityLevelsForProject(context.Background(), projectKeyOrId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPermissionSchemesAPI.GetSecurityLevelsForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLevelsForProject`: ProjectIssueSecurityLevels
	fmt.Fprintf(os.Stdout, "Response from `ProjectPermissionSchemesAPI.GetSecurityLevelsForProject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectKeyOrId** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLevelsForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectIssueSecurityLevels**](ProjectIssueSecurityLevels.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

