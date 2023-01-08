# \ProjectRoleActorsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddActorUsers**](ProjectRoleActorsAPI.md#AddActorUsers) | **Post** /rest/api/3/project/{projectIdOrKey}/role/{id} | Add actors to project role
[**AddProjectRoleActorsToRole**](ProjectRoleActorsAPI.md#AddProjectRoleActorsToRole) | **Post** /rest/api/3/role/{id}/actors | Add default actors to project role
[**DeleteActor**](ProjectRoleActorsAPI.md#DeleteActor) | **Delete** /rest/api/3/project/{projectIdOrKey}/role/{id} | Delete actors from project role
[**DeleteProjectRoleActorsFromRole**](ProjectRoleActorsAPI.md#DeleteProjectRoleActorsFromRole) | **Delete** /rest/api/3/role/{id}/actors | Delete default actors from project role
[**GetProjectRoleActorsForRole**](ProjectRoleActorsAPI.md#GetProjectRoleActorsForRole) | **Get** /rest/api/3/role/{id}/actors | Get default actors for project role
[**SetActors**](ProjectRoleActorsAPI.md#SetActors) | **Put** /rest/api/3/project/{projectIdOrKey}/role/{id} | Set actors for project role



## AddActorUsers

> ProjectRole AddActorUsers(ctx, projectIdOrKey, id).ActorsMap(actorsMap).Execute()

Add actors to project role



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
	actorsMap := *openapiclient.NewActorsMap() // ActorsMap | The groups or users to associate with the project role for this project. Provide the user account ID, group name, or group ID. As a group's name can change, use of group ID is recommended.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRoleActorsAPI.AddActorUsers(context.Background(), projectIdOrKey, id).ActorsMap(actorsMap).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.AddActorUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddActorUsers`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRoleActorsAPI.AddActorUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddActorUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **actorsMap** | [**ActorsMap**](ActorsMap.md) | The groups or users to associate with the project role for this project. Provide the user account ID, group name, or group ID. As a group&#39;s name can change, use of group ID is recommended. | 

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


## AddProjectRoleActorsToRole

> ProjectRole AddProjectRoleActorsToRole(ctx, id).ActorInputBean(actorInputBean).Execute()

Add default actors to project role



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
	actorInputBean := *openapiclient.NewActorInputBean() // ActorInputBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRoleActorsAPI.AddProjectRoleActorsToRole(context.Background(), id).ActorInputBean(actorInputBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.AddProjectRoleActorsToRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddProjectRoleActorsToRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRoleActorsAPI.AddProjectRoleActorsToRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddProjectRoleActorsToRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **actorInputBean** | [**ActorInputBean**](ActorInputBean.md) |  | 

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


## DeleteActor

> DeleteActor(ctx, projectIdOrKey, id).User(user).Group(group).GroupId(groupId).Execute()

Delete actors from project role



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
	user := "5b10ac8d82e05b22cc7d4ef5" // string | The user account ID of the user to remove from the project role. (optional)
	group := "group_example" // string | The name of the group to remove from the project role. This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended. (optional)
	groupId := "groupId_example" // string | The ID of the group to remove from the project role. This parameter cannot be used with the `group` parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectRoleActorsAPI.DeleteActor(context.Background(), projectIdOrKey, id).User(user).Group(group).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.DeleteActor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteActorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **user** | **string** | The user account ID of the user to remove from the project role. | 
 **group** | **string** | The name of the group to remove from the project role. This parameter cannot be used with the &#x60;groupId&#x60; parameter. As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended. | 
 **groupId** | **string** | The ID of the group to remove from the project role. This parameter cannot be used with the &#x60;group&#x60; parameter. | 

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


## DeleteProjectRoleActorsFromRole

> ProjectRole DeleteProjectRoleActorsFromRole(ctx, id).User(user).GroupId(groupId).Group(group).Execute()

Delete default actors from project role



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
	user := "5b10ac8d82e05b22cc7d4ef5" // string | The user account ID of the user to remove as a default actor. (optional)
	groupId := "groupId_example" // string | The group ID of the group to be removed as a default actor. This parameter cannot be used with the `group` parameter. (optional)
	group := "group_example" // string | The group name of the group to be removed as a default actor.This parameter cannot be used with the `groupId` parameter. As a group's name can change, use of `groupId` is recommended. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRoleActorsAPI.DeleteProjectRoleActorsFromRole(context.Background(), id).User(user).GroupId(groupId).Group(group).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.DeleteProjectRoleActorsFromRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteProjectRoleActorsFromRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRoleActorsAPI.DeleteProjectRoleActorsFromRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectRoleActorsFromRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **user** | **string** | The user account ID of the user to remove as a default actor. | 
 **groupId** | **string** | The group ID of the group to be removed as a default actor. This parameter cannot be used with the &#x60;group&#x60; parameter. | 
 **group** | **string** | The group name of the group to be removed as a default actor.This parameter cannot be used with the &#x60;groupId&#x60; parameter. As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended. | 

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


## GetProjectRoleActorsForRole

> ProjectRole GetProjectRoleActorsForRole(ctx, id).Execute()

Get default actors for project role



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
	resp, r, err := apiClient.ProjectRoleActorsAPI.GetProjectRoleActorsForRole(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.GetProjectRoleActorsForRole``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectRoleActorsForRole`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRoleActorsAPI.GetProjectRoleActorsForRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectRoleActorsForRoleRequest struct via the builder pattern


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


## SetActors

> ProjectRole SetActors(ctx, projectIdOrKey, id).ProjectRoleActorsUpdateBean(projectRoleActorsUpdateBean).Execute()

Set actors for project role



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
	projectRoleActorsUpdateBean := *openapiclient.NewProjectRoleActorsUpdateBean() // ProjectRoleActorsUpdateBean | The groups or users to associate with the project role for this project. Provide the user account ID, group name, or group ID. As a group's name can change, use of group ID is recommended.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectRoleActorsAPI.SetActors(context.Background(), projectIdOrKey, id).ProjectRoleActorsUpdateBean(projectRoleActorsUpdateBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectRoleActorsAPI.SetActors``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetActors`: ProjectRole
	fmt.Fprintf(os.Stdout, "Response from `ProjectRoleActorsAPI.SetActors`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**id** | **int64** | The ID of the project role. Use [Get all project roles](#api-rest-api-3-role-get) to get a list of project role IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetActorsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectRoleActorsUpdateBean** | [**ProjectRoleActorsUpdateBean**](ProjectRoleActorsUpdateBean.md) | The groups or users to associate with the project role for this project. Provide the user account ID, group name, or group ID. As a group&#39;s name can change, use of group ID is recommended. | 

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

