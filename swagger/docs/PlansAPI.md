# \PlansAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ArchivePlan**](PlansAPI.md#ArchivePlan) | **Put** /rest/api/3/plans/plan/{planId}/archive | Archive plan
[**CreatePlan**](PlansAPI.md#CreatePlan) | **Post** /rest/api/3/plans/plan | Create plan
[**DuplicatePlan**](PlansAPI.md#DuplicatePlan) | **Post** /rest/api/3/plans/plan/{planId}/duplicate | Duplicate plan
[**GetPlan**](PlansAPI.md#GetPlan) | **Get** /rest/api/3/plans/plan/{planId} | Get plan
[**GetPlans**](PlansAPI.md#GetPlans) | **Get** /rest/api/3/plans/plan | Get plans paginated
[**TrashPlan**](PlansAPI.md#TrashPlan) | **Put** /rest/api/3/plans/plan/{planId}/trash | Trash plan
[**UpdatePlan**](PlansAPI.md#UpdatePlan) | **Put** /rest/api/3/plans/plan/{planId} | Update plan



## ArchivePlan

> interface{} ArchivePlan(ctx, planId).Execute()

Archive plan



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.ArchivePlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.ArchivePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ArchivePlan`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.ArchivePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiArchivePlanRequest struct via the builder pattern


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


## CreatePlan

> int64 CreatePlan(ctx).CreatePlanRequest(createPlanRequest).UseGroupId(useGroupId).Execute()

Create plan



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
	createPlanRequest := *openapiclient.NewCreatePlanRequest([]openapiclient.CreateIssueSourceRequest{*openapiclient.NewCreateIssueSourceRequest("Type_example", int64(123))}, "Name_example", *openapiclient.NewCreateSchedulingRequest("Estimation_example")) // CreatePlanRequest | 
	useGroupId := true // bool | Whether to accept group IDs instead of group names. Group names are deprecated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.CreatePlan(context.Background()).CreatePlanRequest(createPlanRequest).UseGroupId(useGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.CreatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePlan`: int64
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPlanRequest** | [**CreatePlanRequest**](CreatePlanRequest.md) |  | 
 **useGroupId** | **bool** | Whether to accept group IDs instead of group names. Group names are deprecated. | [default to false]

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


## DuplicatePlan

> int64 DuplicatePlan(ctx, planId).DuplicatePlanRequest(duplicatePlanRequest).Execute()

Duplicate plan



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
	duplicatePlanRequest := *openapiclient.NewDuplicatePlanRequest("Name_example") // DuplicatePlanRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.DuplicatePlan(context.Background(), planId).DuplicatePlanRequest(duplicatePlanRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.DuplicatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DuplicatePlan`: int64
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.DuplicatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDuplicatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **duplicatePlanRequest** | [**DuplicatePlanRequest**](DuplicatePlanRequest.md) |  | 

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


## GetPlan

> GetPlanResponse GetPlan(ctx, planId).UseGroupId(useGroupId).Execute()

Get plan



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
	useGroupId := true // bool | Whether to return group IDs instead of group names. Group names are deprecated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlan(context.Background(), planId).UseGroupId(useGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlan`: GetPlanResponse
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **useGroupId** | **bool** | Whether to return group IDs instead of group names. Group names are deprecated. | [default to false]

### Return type

[**GetPlanResponse**](GetPlanResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPlans

> PageWithCursorGetPlanResponseForPage GetPlans(ctx).IncludeTrashed(includeTrashed).IncludeArchived(includeArchived).Cursor(cursor).MaxResults(maxResults).Execute()

Get plans paginated



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
	includeTrashed := true // bool | Whether to include trashed plans in the results. (optional) (default to false)
	includeArchived := true // bool | Whether to include archived plans in the results. (optional) (default to false)
	cursor := "cursor_example" // string | The cursor to start from. If not provided, the first page will be returned. (optional) (default to "")
	maxResults := int32(56) // int32 | The maximum number of plans to return per page. The maximum value is 50. The default value is 50. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.GetPlans(context.Background()).IncludeTrashed(includeTrashed).IncludeArchived(includeArchived).Cursor(cursor).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.GetPlans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPlans`: PageWithCursorGetPlanResponseForPage
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.GetPlans`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPlansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **includeTrashed** | **bool** | Whether to include trashed plans in the results. | [default to false]
 **includeArchived** | **bool** | Whether to include archived plans in the results. | [default to false]
 **cursor** | **string** | The cursor to start from. If not provided, the first page will be returned. | [default to &quot;&quot;]
 **maxResults** | **int32** | The maximum number of plans to return per page. The maximum value is 50. The default value is 50. | [default to 50]

### Return type

[**PageWithCursorGetPlanResponseForPage**](PageWithCursorGetPlanResponseForPage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TrashPlan

> interface{} TrashPlan(ctx, planId).Execute()

Trash plan



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.TrashPlan(context.Background(), planId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.TrashPlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrashPlan`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.TrashPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrashPlanRequest struct via the builder pattern


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


## UpdatePlan

> interface{} UpdatePlan(ctx, planId).Body(body).UseGroupId(useGroupId).Execute()

Update plan



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
	body := map[string]interface{}{ ... } // map[string]interface{} | 
	useGroupId := true // bool | Whether to accept group IDs instead of group names. Group names are deprecated. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PlansAPI.UpdatePlan(context.Background(), planId).Body(body).UseGroupId(useGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PlansAPI.UpdatePlan``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePlan`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `PlansAPI.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **int64** | The ID of the plan. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **map[string]interface{}** |  | 
 **useGroupId** | **bool** | Whether to accept group IDs instead of group names. Group names are deprecated. | [default to false]

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

