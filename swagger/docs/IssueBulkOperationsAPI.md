# \IssueBulkOperationsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableTransitions**](IssueBulkOperationsAPI.md#GetAvailableTransitions) | **Get** /rest/api/3/bulk/issues/transition | Get available transitions
[**GetBulkEditableFields**](IssueBulkOperationsAPI.md#GetBulkEditableFields) | **Get** /rest/api/3/bulk/issues/fields | Get bulk editable fields
[**GetBulkOperationProgress**](IssueBulkOperationsAPI.md#GetBulkOperationProgress) | **Get** /rest/api/3/bulk/queue/{taskId} | Get bulk issue operation progress
[**SubmitBulkDelete**](IssueBulkOperationsAPI.md#SubmitBulkDelete) | **Post** /rest/api/3/bulk/issues/delete | Bulk delete issues
[**SubmitBulkEdit**](IssueBulkOperationsAPI.md#SubmitBulkEdit) | **Post** /rest/api/3/bulk/issues/fields | Bulk edit issues
[**SubmitBulkMove**](IssueBulkOperationsAPI.md#SubmitBulkMove) | **Post** /rest/api/3/bulk/issues/move | Bulk move issues
[**SubmitBulkTransition**](IssueBulkOperationsAPI.md#SubmitBulkTransition) | **Post** /rest/api/3/bulk/issues/transition | Bulk transition issue statuses
[**SubmitBulkUnwatch**](IssueBulkOperationsAPI.md#SubmitBulkUnwatch) | **Post** /rest/api/3/bulk/issues/unwatch | Bulk unwatch issues
[**SubmitBulkWatch**](IssueBulkOperationsAPI.md#SubmitBulkWatch) | **Post** /rest/api/3/bulk/issues/watch | Bulk watch issues



## GetAvailableTransitions

> BulkTransitionGetAvailableTransitions GetAvailableTransitions(ctx).IssueIdsOrKeys(issueIdsOrKeys).EndingBefore(endingBefore).StartingAfter(startingAfter).Execute()

Get available transitions



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
	issueIdsOrKeys := "issueIdsOrKeys_example" // string | Comma (,) separated Ids or keys of the issues to get transitions available for them.
	endingBefore := "endingBefore_example" // string | (Optional)The end cursor for use in pagination. (optional)
	startingAfter := "startingAfter_example" // string | (Optional)The start cursor for use in pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.GetAvailableTransitions(context.Background()).IssueIdsOrKeys(issueIdsOrKeys).EndingBefore(endingBefore).StartingAfter(startingAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.GetAvailableTransitions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableTransitions`: BulkTransitionGetAvailableTransitions
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.GetAvailableTransitions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableTransitionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueIdsOrKeys** | **string** | Comma (,) separated Ids or keys of the issues to get transitions available for them. | 
 **endingBefore** | **string** | (Optional)The end cursor for use in pagination. | 
 **startingAfter** | **string** | (Optional)The start cursor for use in pagination. | 

### Return type

[**BulkTransitionGetAvailableTransitions**](BulkTransitionGetAvailableTransitions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkEditableFields

> BulkEditGetFields GetBulkEditableFields(ctx).IssueIdsOrKeys(issueIdsOrKeys).SearchText(searchText).EndingBefore(endingBefore).StartingAfter(startingAfter).Execute()

Get bulk editable fields



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
	issueIdsOrKeys := "issueIdsOrKeys_example" // string | The IDs or keys of the issues to get editable fields from.
	searchText := "searchText_example" // string | (Optional)The text to search for in the editable fields. (optional)
	endingBefore := "endingBefore_example" // string | (Optional)The end cursor for use in pagination. (optional)
	startingAfter := "startingAfter_example" // string | (Optional)The start cursor for use in pagination. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.GetBulkEditableFields(context.Background()).IssueIdsOrKeys(issueIdsOrKeys).SearchText(searchText).EndingBefore(endingBefore).StartingAfter(startingAfter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.GetBulkEditableFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkEditableFields`: BulkEditGetFields
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.GetBulkEditableFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkEditableFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueIdsOrKeys** | **string** | The IDs or keys of the issues to get editable fields from. | 
 **searchText** | **string** | (Optional)The text to search for in the editable fields. | 
 **endingBefore** | **string** | (Optional)The end cursor for use in pagination. | 
 **startingAfter** | **string** | (Optional)The start cursor for use in pagination. | 

### Return type

[**BulkEditGetFields**](BulkEditGetFields.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBulkOperationProgress

> BulkOperationProgress GetBulkOperationProgress(ctx, taskId).Execute()

Get bulk issue operation progress



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
	taskId := "taskId_example" // string | The ID of the task.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.GetBulkOperationProgress(context.Background(), taskId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.GetBulkOperationProgress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBulkOperationProgress`: BulkOperationProgress
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.GetBulkOperationProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**taskId** | **string** | The ID of the task. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkOperationProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**BulkOperationProgress**](BulkOperationProgress.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkDelete

> SubmittedBulkOperation SubmitBulkDelete(ctx).IssueBulkDeletePayload(issueBulkDeletePayload).Execute()

Bulk delete issues



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
	issueBulkDeletePayload := *openapiclient.NewIssueBulkDeletePayload([]string{"SelectedIssueIdsOrKeys_example"}) // IssueBulkDeletePayload | The request body containing the issues to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkDelete(context.Background()).IssueBulkDeletePayload(issueBulkDeletePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkDelete`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkDelete`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkDeletePayload** | [**IssueBulkDeletePayload**](IssueBulkDeletePayload.md) | The request body containing the issues to be deleted. | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkEdit

> SubmittedBulkOperation SubmitBulkEdit(ctx).IssueBulkEditPayload(issueBulkEditPayload).Execute()

Bulk edit issues



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
	issueBulkEditPayload := *openapiclient.NewIssueBulkEditPayload(*openapiclient.NewJiraIssueFields(), []string{"SelectedActions_example"}, []string{"SelectedIssueIdsOrKeys_example"}) // IssueBulkEditPayload | The request body containing the issues to be edited and the new field values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkEdit(context.Background()).IssueBulkEditPayload(issueBulkEditPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkEdit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkEdit`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkEdit`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkEditPayload** | [**IssueBulkEditPayload**](IssueBulkEditPayload.md) | The request body containing the issues to be edited and the new field values. | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkMove

> SubmittedBulkOperation SubmitBulkMove(ctx).IssueBulkMovePayload(issueBulkMovePayload).Execute()

Bulk move issues



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
	issueBulkMovePayload := *openapiclient.NewIssueBulkMovePayload() // IssueBulkMovePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkMove(context.Background()).IssueBulkMovePayload(issueBulkMovePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkMove``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkMove`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkMove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkMoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkMovePayload** | [**IssueBulkMovePayload**](IssueBulkMovePayload.md) |  | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkTransition

> SubmittedBulkOperation SubmitBulkTransition(ctx).IssueBulkTransitionPayload(issueBulkTransitionPayload).Execute()

Bulk transition issue statuses



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
	issueBulkTransitionPayload := *openapiclient.NewIssueBulkTransitionPayload([]openapiclient.BulkTransitionSubmitInput{*openapiclient.NewBulkTransitionSubmitInput([]string{"SelectedIssueIdsOrKeys_example"}, "TransitionId_example")}) // IssueBulkTransitionPayload | The request body containing the issues to be transitioned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkTransition(context.Background()).IssueBulkTransitionPayload(issueBulkTransitionPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkTransition``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkTransition`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkTransition`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkTransitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkTransitionPayload** | [**IssueBulkTransitionPayload**](IssueBulkTransitionPayload.md) | The request body containing the issues to be transitioned. | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkUnwatch

> SubmittedBulkOperation SubmitBulkUnwatch(ctx).IssueBulkWatchOrUnwatchPayload(issueBulkWatchOrUnwatchPayload).Execute()

Bulk unwatch issues



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
	issueBulkWatchOrUnwatchPayload := *openapiclient.NewIssueBulkWatchOrUnwatchPayload([]string{"SelectedIssueIdsOrKeys_example"}) // IssueBulkWatchOrUnwatchPayload | The request body containing the issues to be unwatched.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkUnwatch(context.Background()).IssueBulkWatchOrUnwatchPayload(issueBulkWatchOrUnwatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkUnwatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkUnwatch`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkUnwatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkUnwatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkWatchOrUnwatchPayload** | [**IssueBulkWatchOrUnwatchPayload**](IssueBulkWatchOrUnwatchPayload.md) | The request body containing the issues to be unwatched. | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBulkWatch

> SubmittedBulkOperation SubmitBulkWatch(ctx).IssueBulkWatchOrUnwatchPayload(issueBulkWatchOrUnwatchPayload).Execute()

Bulk watch issues



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
	issueBulkWatchOrUnwatchPayload := *openapiclient.NewIssueBulkWatchOrUnwatchPayload([]string{"SelectedIssueIdsOrKeys_example"}) // IssueBulkWatchOrUnwatchPayload | The request body containing the issues to be watched.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueBulkOperationsAPI.SubmitBulkWatch(context.Background()).IssueBulkWatchOrUnwatchPayload(issueBulkWatchOrUnwatchPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueBulkOperationsAPI.SubmitBulkWatch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitBulkWatch`: SubmittedBulkOperation
	fmt.Fprintf(os.Stdout, "Response from `IssueBulkOperationsAPI.SubmitBulkWatch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBulkWatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueBulkWatchOrUnwatchPayload** | [**IssueBulkWatchOrUnwatchPayload**](IssueBulkWatchOrUnwatchPayload.md) | The request body containing the issues to be watched. | 

### Return type

[**SubmittedBulkOperation**](SubmittedBulkOperation.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

