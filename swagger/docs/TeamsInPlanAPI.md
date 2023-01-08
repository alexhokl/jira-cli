# \TeamsInPlanAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAtlassianTeam**](TeamsInPlanAPI.md#AddAtlassianTeam) | **Post** /rest/api/3/plans/plan/{planId}/team/atlassian | Add Atlassian team to plan
[**CreatePlanOnlyTeam**](TeamsInPlanAPI.md#CreatePlanOnlyTeam) | **Post** /rest/api/3/plans/plan/{planId}/team/planonly | Create plan-only team
[**DeletePlanOnlyTeam**](TeamsInPlanAPI.md#DeletePlanOnlyTeam) | **Delete** /rest/api/3/plans/plan/{planId}/team/planonly/{planOnlyTeamId} | Delete plan-only team
[**GetAtlassianTeam**](TeamsInPlanAPI.md#GetAtlassianTeam) | **Get** /rest/api/3/plans/plan/{planId}/team/atlassian/{atlassianTeamId} | Get Atlassian team in plan
[**GetPlanOnlyTeam**](TeamsInPlanAPI.md#GetPlanOnlyTeam) | **Get** /rest/api/3/plans/plan/{planId}/team/planonly/{planOnlyTeamId} | Get plan-only team
[**GetTeams**](TeamsInPlanAPI.md#GetTeams) | **Get** /rest/api/3/plans/plan/{planId}/team | Get teams in plan paginated
[**RemoveAtlassianTeam**](TeamsInPlanAPI.md#RemoveAtlassianTeam) | **Delete** /rest/api/3/plans/plan/{planId}/team/atlassian/{atlassianTeamId} | Remove Atlassian team from plan
[**UpdateAtlassianTeam**](TeamsInPlanAPI.md#UpdateAtlassianTeam) | **Put** /rest/api/3/plans/plan/{planId}/team/atlassian/{atlassianTeamId} | Update Atlassian team in plan
[**UpdatePlanOnlyTeam**](TeamsInPlanAPI.md#UpdatePlanOnlyTeam) | **Put** /rest/api/3/plans/plan/{planId}/team/planonly/{planOnlyTeamId} | Update plan-only team



## AddAtlassianTeam

> interface{} AddAtlassianTeam(ctx, planId).AddAtlassianTeamRequest(addAtlassianTeamRequest).Execute()

Add Atlassian team to plan



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
	planId := int64(789) // int64 | The ID of the plan.
	addAtlassianTeamRequest := *openapiclient.NewAddAtlassianTeamRequest("Id_example", "PlanningStyle_example") // AddAtlassianTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.AddAtlassianTeam(context.Background(), planId).AddAtlassianTeamRequest(addAtlassianTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.AddAtlassianTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAtlassianTeam`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.AddAtlassianTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAtlassianTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addAtlassianTeamRequest** | [**AddAtlassianTeamRequest**](AddAtlassianTeamRequest.md) |  | 

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


## CreatePlanOnlyTeam

> int64 CreatePlanOnlyTeam(ctx, planId).CreatePlanOnlyTeamRequest(createPlanOnlyTeamRequest).Execute()

Create plan-only team



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
	planId := int64(789) // int64 | The ID of the plan.
	createPlanOnlyTeamRequest := *openapiclient.NewCreatePlanOnlyTeamRequest("Name_example", "PlanningStyle_example") // CreatePlanOnlyTeamRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.CreatePlanOnlyTeam(context.Background(), planId).CreatePlanOnlyTeamRequest(createPlanOnlyTeamRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.CreatePlanOnlyTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlanOnlyTeam`: int64
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.CreatePlanOnlyTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanOnlyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createPlanOnlyTeamRequest** | [**CreatePlanOnlyTeamRequest**](CreatePlanOnlyTeamRequest.md) |  | 

### Return type

**int64**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlanOnlyTeam

> interface{} DeletePlanOnlyTeam(ctx, planId, planOnlyTeamId).Execute()

Delete plan-only team



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
	planId := int64(789) // int64 | The ID of the plan.
	planOnlyTeamId := int64(789) // int64 | The ID of the plan-only team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.DeletePlanOnlyTeam(context.Background(), planId, planOnlyTeamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.DeletePlanOnlyTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePlanOnlyTeam`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.DeletePlanOnlyTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**planOnlyTeamId** | **int64** | The ID of the plan-only team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanOnlyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAtlassianTeam

> GetAtlassianTeamResponse GetAtlassianTeam(ctx, planId, atlassianTeamId).Execute()

Get Atlassian team in plan



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
	planId := int64(789) // int64 | The ID of the plan.
	atlassianTeamId := "atlassianTeamId_example" // string | The ID of the Atlassian team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.GetAtlassianTeam(context.Background(), planId, atlassianTeamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.GetAtlassianTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAtlassianTeam`: GetAtlassianTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.GetAtlassianTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**atlassianTeamId** | **string** | The ID of the Atlassian team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAtlassianTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAtlassianTeamResponse**](GetAtlassianTeamResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlanOnlyTeam

> GetPlanOnlyTeamResponse GetPlanOnlyTeam(ctx, planId, planOnlyTeamId).Execute()

Get plan-only team



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
	planId := int64(789) // int64 | The ID of the plan.
	planOnlyTeamId := int64(789) // int64 | The ID of the plan-only team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.GetPlanOnlyTeam(context.Background(), planId, planOnlyTeamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.GetPlanOnlyTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlanOnlyTeam`: GetPlanOnlyTeamResponse
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.GetPlanOnlyTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**planOnlyTeamId** | **int64** | The ID of the plan-only team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanOnlyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetPlanOnlyTeamResponse**](GetPlanOnlyTeamResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeams

> PageWithCursorGetTeamResponseForPage GetTeams(ctx, planId).Cursor(cursor).MaxResults(maxResults).Execute()

Get teams in plan paginated



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
	planId := int64(789) // int64 | The ID of the plan.
	cursor := "cursor_example" // string | The cursor to start from. If not provided, the first page will be returned. (optional) (default to "")
	maxResults := int32(56) // int32 | The maximum number of plan teams to return per page. The maximum value is 50. The default value is 50. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.GetTeams(context.Background(), planId).Cursor(cursor).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.GetTeams``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTeams`: PageWithCursorGetTeamResponseForPage
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.GetTeams`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cursor** | **string** | The cursor to start from. If not provided, the first page will be returned. | [default to &quot;&quot;]
 **maxResults** | **int32** | The maximum number of plan teams to return per page. The maximum value is 50. The default value is 50. | [default to 50]

### Return type

[**PageWithCursorGetTeamResponseForPage**](PageWithCursorGetTeamResponseForPage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveAtlassianTeam

> interface{} RemoveAtlassianTeam(ctx, planId, atlassianTeamId).Execute()

Remove Atlassian team from plan



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
	planId := int64(789) // int64 | The ID of the plan.
	atlassianTeamId := "atlassianTeamId_example" // string | The ID of the Atlassian team.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.RemoveAtlassianTeam(context.Background(), planId, atlassianTeamId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.RemoveAtlassianTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAtlassianTeam`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.RemoveAtlassianTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**atlassianTeamId** | **string** | The ID of the Atlassian team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAtlassianTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAtlassianTeam

> interface{} UpdateAtlassianTeam(ctx, planId, atlassianTeamId).Body(body).Execute()

Update Atlassian team in plan



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
	planId := int64(789) // int64 | The ID of the plan.
	atlassianTeamId := "atlassianTeamId_example" // string | The ID of the Atlassian team.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.UpdateAtlassianTeam(context.Background(), planId, atlassianTeamId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.UpdateAtlassianTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateAtlassianTeam`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.UpdateAtlassianTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**atlassianTeamId** | **string** | The ID of the Atlassian team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAtlassianTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlanOnlyTeam

> interface{} UpdatePlanOnlyTeam(ctx, planId, planOnlyTeamId).Body(body).Execute()

Update plan-only team



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
	planId := int64(789) // int64 | The ID of the plan.
	planOnlyTeamId := int64(789) // int64 | The ID of the plan-only team.
	body := map[string]interface{}{ ... } // map[string]interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TeamsInPlanAPI.UpdatePlanOnlyTeam(context.Background(), planId, planOnlyTeamId).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TeamsInPlanAPI.UpdatePlanOnlyTeam``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlanOnlyTeam`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TeamsInPlanAPI.UpdatePlanOnlyTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 
**planOnlyTeamId** | **int64** | The ID of the plan-only team. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanOnlyTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **map[string]interface{}** |  | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json-patch+json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

