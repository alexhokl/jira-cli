# \ProjectRolesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProjectRole**](ProjectRolesAPI.md#CreateProjectRole) | **Post** /rest/api/3/role | Create project role
[**DeleteProjectRole**](ProjectRolesAPI.md#DeleteProjectRole) | **Delete** /rest/api/3/role/{id} | Delete project role
[**FullyUpdateProjectRole**](ProjectRolesAPI.md#FullyUpdateProjectRole) | **Put** /rest/api/3/role/{id} | Fully update project role
[**GetAllProjectRoles**](ProjectRolesAPI.md#GetAllProjectRoles) | **Get** /rest/api/3/role | Get all project roles
[**GetProjectRole**](ProjectRolesAPI.md#GetProjectRole) | **Get** /rest/api/3/project/{projectIdOrKey}/role/{id} | Get project role for project
[**GetProjectRoleById**](ProjectRolesAPI.md#GetProjectRoleById) | **Get** /rest/api/3/role/{id} | Get project role by ID
[**GetProjectRoleDetails**](ProjectRolesAPI.md#GetProjectRoleDetails) | **Get** /rest/api/3/project/{projectIdOrKey}/roledetails | Get project role details
[**GetProjectRoles**](ProjectRolesAPI.md#GetProjectRoles) | **Get** /rest/api/3/project/{projectIdOrKey}/role | Get project roles for project
[**PartialUpdateProjectRole**](ProjectRolesAPI.md#PartialUpdateProjectRole) | **Post** /rest/api/3/role/{id} | Partial update project role



## CreateProjectRole

> ProjectRole CreateProjectRole(ctx).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()

Create project role



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
	createUpdateRoleRequestBean := *openapiclient.NewCreateUpdateRoleRequestBean() // CreateUpdateRoleRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.CreateProjectRole(context.Background()).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.CreateProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateProjectRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.CreateProjectRole`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUpdateRoleRequestBean** | [**CreateUpdateRoleRequestBean**](CreateUpdateRoleRequestBean.md) |  | 

### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteProjectRole

> DeleteProjectRole(ctx, id).Swap(swap).Execute()

Delete project role



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
	id := int64(789) // int64 | The ID of the project role to delete. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
	swap := int64(789) // int64 | The ID of the project role that will replace the one being deleted. The swap will attempt to swap the role in schemes (notifications, permissions, issue security), workflows, worklogs and comments. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectRolesAPI.DeleteProjectRole(context.Background(), id).Swap(swap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.DeleteProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role to delete. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **swap** | **int64** | The ID of the project role that will replace the one being deleted. The swap will attempt to swap the role in schemes (notifications, permissions, issue security), workflows, worklogs and comments. | 

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


## FullyUpdateProjectRole

> ProjectRole FullyUpdateProjectRole(ctx, id).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()

Fully update project role



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
	id := int64(789) // int64 | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
	createUpdateRoleRequestBean := *openapiclient.NewCreateUpdateRoleRequestBean() // CreateUpdateRoleRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.FullyUpdateProjectRole(context.Background(), id).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.FullyUpdateProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FullyUpdateProjectRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.FullyUpdateProjectRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiFullyUpdateProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateRoleRequestBean** | [**CreateUpdateRoleRequestBean**](CreateUpdateRoleRequestBean.md) |  | 

### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllProjectRoles

> []ProjectRole GetAllProjectRoles(ctx).Execute()

Get all project roles



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
	resp, r, err := apiClient.ProjectRolesAPI.GetAllProjectRoles(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.GetAllProjectRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllProjectRoles`: []ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.GetAllProjectRoles`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllProjectRolesRequest struct via the builder pattern


### Return type

[**[]ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRole

> ProjectRole GetProjectRole(ctx, projectIdOrKey, id).ExcludeInactiveUsers(excludeInactiveUsers).Execute()

Get project role for project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	id := int64(789) // int64 | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
	excludeInactiveUsers := true // bool | Exclude inactive users. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.GetProjectRole(context.Background(), projectIdOrKey, id).ExcludeInactiveUsers(excludeInactiveUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.GetProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.GetProjectRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **excludeInactiveUsers** | **bool** | Exclude inactive users. | [default to false]

### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRoleById

> ProjectRole GetProjectRoleById(ctx, id).Execute()

Get project role by ID



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
	id := int64(789) // int64 | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.GetProjectRoleById(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.GetProjectRoleById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoleById`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.GetProjectRoleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRoleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRoleDetails

> []ProjectRoleDetails GetProjectRoleDetails(ctx, projectIdOrKey).CurrentMember(currentMember).ExcludeConnectAddons(excludeConnectAddons).ExcludeOtherServiceRoles(excludeOtherServiceRoles).Execute()

Get project role details



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	currentMember := true // bool | Whether the roles should be filtered to include only those the user is assigned to. (optional) (default to false)
	excludeConnectAddons := true // bool |  (optional) (default to false)
	excludeOtherServiceRoles := true // bool | Do not return the default JSM company-managed space from CSM spaces, or the default CSM roles from JSM spaces. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.GetProjectRoleDetails(context.Background(), projectIdOrKey).CurrentMember(currentMember).ExcludeConnectAddons(excludeConnectAddons).ExcludeOtherServiceRoles(excludeOtherServiceRoles).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.GetProjectRoleDetails``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoleDetails`: []ProjectRoleDetails
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.GetProjectRoleDetails`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRoleDetailsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **currentMember** | **bool** | Whether the roles should be filtered to include only those the user is assigned to. | [default to false]
 **excludeConnectAddons** | **bool** |  | [default to false]
 **excludeOtherServiceRoles** | **bool** | Do not return the default JSM company-managed space from CSM spaces, or the default CSM roles from JSM spaces. | [default to false]

### Return type

[**[]ProjectRoleDetails**](ProjectRoleDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectRoles

> map[string]string GetProjectRoles(ctx, projectIdOrKey).Execute()

Get project roles for project



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.GetProjectRoles(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.GetProjectRoles``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoles`: map[string]string
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.GetProjectRoles`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRolesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**map[string]string**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartialUpdateProjectRole

> ProjectRole PartialUpdateProjectRole(ctx, id).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()

Partial update project role



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
	id := int64(789) // int64 | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs.
	createUpdateRoleRequestBean := *openapiclient.NewCreateUpdateRoleRequestBean() // CreateUpdateRoleRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRolesAPI.PartialUpdateProjectRole(context.Background(), id).CreateUpdateRoleRequestBean(createUpdateRoleRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRolesAPI.PartialUpdateProjectRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PartialUpdateProjectRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRolesAPI.PartialUpdateProjectRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartialUpdateProjectRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createUpdateRoleRequestBean** | [**CreateUpdateRoleRequestBean**](CreateUpdateRoleRequestBean.md) |  | 

### Return type

[**ProjectRole**](ProjectRole.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

