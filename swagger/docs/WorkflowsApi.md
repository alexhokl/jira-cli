# \WorkflowsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflow**](WorkflowsAPI.md#CreateWorkflow) | **Post** /rest/api/3/workflow | Create workflow
[**CreateWorkflows**](WorkflowsAPI.md#CreateWorkflows) | **Post** /rest/api/3/workflows/create | Bulk create workflows
[**DeleteInactiveWorkflow**](WorkflowsAPI.md#DeleteInactiveWorkflow) | **Delete** /rest/api/3/workflow/{entityId} | Delete inactive workflow
[**GetAllWorkflows**](WorkflowsAPI.md#GetAllWorkflows) | **Get** /rest/api/3/workflow | Get all workflows
[**GetDefaultEditor**](WorkflowsAPI.md#GetDefaultEditor) | **Get** /rest/api/3/workflows/defaultEditor | Get the user&#39;s default workflow editor
[**GetProjectUsagesForWorkflow**](WorkflowsAPI.md#GetProjectUsagesForWorkflow) | **Get** /rest/api/3/workflow/{workflowId}/projectUsages | Get projects using a given workflow
[**GetWorkflowProjectIssueTypeUsages**](WorkflowsAPI.md#GetWorkflowProjectIssueTypeUsages) | **Get** /rest/api/3/workflow/{workflowId}/project/{projectId}/issueTypeUsages | Get issue types in a project that are using a given workflow
[**GetWorkflowSchemeUsagesForWorkflow**](WorkflowsAPI.md#GetWorkflowSchemeUsagesForWorkflow) | **Get** /rest/api/3/workflow/{workflowId}/workflowSchemes | Get workflow schemes which are using a given workflow
[**GetWorkflowsPaginated**](WorkflowsAPI.md#GetWorkflowsPaginated) | **Get** /rest/api/3/workflow/search | Get workflows paginated
[**ListWorkflowHistory**](WorkflowsAPI.md#ListWorkflowHistory) | **Post** /rest/api/3/workflow/history/list | List workflow history entries
[**ReadWorkflowFromHistory**](WorkflowsAPI.md#ReadWorkflowFromHistory) | **Post** /rest/api/3/workflow/history | Read workflow version from history
[**ReadWorkflowPreviews**](WorkflowsAPI.md#ReadWorkflowPreviews) | **Post** /rest/api/3/workflows/preview | Preview workflow
[**ReadWorkflows**](WorkflowsAPI.md#ReadWorkflows) | **Post** /rest/api/3/workflows | Bulk get workflows
[**SearchWorkflows**](WorkflowsAPI.md#SearchWorkflows) | **Get** /rest/api/3/workflows/search | Search workflows
[**UpdateWorkflows**](WorkflowsAPI.md#UpdateWorkflows) | **Post** /rest/api/3/workflows/update | Bulk update workflows
[**ValidateCreateWorkflows**](WorkflowsAPI.md#ValidateCreateWorkflows) | **Post** /rest/api/3/workflows/create/validation | Validate create workflows
[**ValidateUpdateWorkflows**](WorkflowsAPI.md#ValidateUpdateWorkflows) | **Post** /rest/api/3/workflows/update/validation | Validate update workflows
[**WorkflowCapabilities**](WorkflowsAPI.md#WorkflowCapabilities) | **Get** /rest/api/3/workflows/capabilities | Get available workflow capabilities



## CreateWorkflow

> WorkflowIDs CreateWorkflow(ctx).CreateWorkflowDetails(createWorkflowDetails).Execute()

Create workflow



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
	createWorkflowDetails := *openapiclient.NewCreateWorkflowDetails("Name_example", []openapiclient.CreateWorkflowStatusDetails{*openapiclient.NewCreateWorkflowStatusDetails("Id_example")}, []openapiclient.CreateWorkflowTransitionDetails{*openapiclient.NewCreateWorkflowTransitionDetails("Name_example", "To_example", "Type_example")}) // CreateWorkflowDetails | The workflow details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.CreateWorkflow(context.Background()).CreateWorkflowDetails(createWorkflowDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflow`: WorkflowIDs
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowDetails** | [**CreateWorkflowDetails**](CreateWorkflowDetails.md) | The workflow details. | 

### Return type

[**WorkflowIDs**](WorkflowIDs.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateWorkflows

> WorkflowCreateResponse CreateWorkflows(ctx).WorkflowCreateRequest(workflowCreateRequest).Execute()

Bulk create workflows



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
	workflowCreateRequest := *openapiclient.NewWorkflowCreateRequest() // WorkflowCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.CreateWorkflows(context.Background()).WorkflowCreateRequest(workflowCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.CreateWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflows`: WorkflowCreateResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.CreateWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowCreateRequest** | [**WorkflowCreateRequest**](WorkflowCreateRequest.md) |  | 

### Return type

[**WorkflowCreateResponse**](WorkflowCreateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInactiveWorkflow

> DeleteInactiveWorkflow(ctx, entityId).Execute()

Delete inactive workflow



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
	entityId := "entityId_example" // string | The entity ID of the workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowsAPI.DeleteInactiveWorkflow(context.Background(), entityId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.DeleteInactiveWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityId** | **string** | The entity ID of the workflow. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInactiveWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWorkflows

> []DeprecatedWorkflow GetAllWorkflows(ctx).WorkflowName(workflowName).Execute()

Get all workflows



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
	workflowName := "workflowName_example" // string | The name of the workflow to be returned. Only one workflow can be specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetAllWorkflows(context.Background()).WorkflowName(workflowName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetAllWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWorkflows`: []DeprecatedWorkflow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetAllWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowName** | **string** | The name of the workflow to be returned. Only one workflow can be specified. | 

### Return type

[**[]DeprecatedWorkflow**](DeprecatedWorkflow.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultEditor

> DefaultWorkflowEditorResponse GetDefaultEditor(ctx).Execute()

Get the user's default workflow editor



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
	resp, r, err := apiClient.WorkflowsAPI.GetDefaultEditor(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetDefaultEditor``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultEditor`: DefaultWorkflowEditorResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetDefaultEditor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultEditorRequest struct via the builder pattern


### Return type

[**DefaultWorkflowEditorResponse**](DefaultWorkflowEditorResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectUsagesForWorkflow

> WorkflowProjectUsageDTO GetProjectUsagesForWorkflow(ctx, workflowId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get projects using a given workflow



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
	workflowId := "workflowId_example" // string | The workflow ID
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetProjectUsagesForWorkflow(context.Background(), workflowId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetProjectUsagesForWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectUsagesForWorkflow`: WorkflowProjectUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetProjectUsagesForWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectUsagesForWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**WorkflowProjectUsageDTO**](WorkflowProjectUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowProjectIssueTypeUsages

> WorkflowProjectIssueTypeUsageDTO GetWorkflowProjectIssueTypeUsages(ctx, workflowId, projectId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get issue types in a project that are using a given workflow



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
	workflowId := "workflowId_example" // string | The workflow ID
	projectId := int64(789) // int64 | The project ID
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowProjectIssueTypeUsages(context.Background(), workflowId, projectId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowProjectIssueTypeUsages``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowProjectIssueTypeUsages`: WorkflowProjectIssueTypeUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowProjectIssueTypeUsages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The workflow ID | 
**projectId** | **int64** | The project ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowProjectIssueTypeUsagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**WorkflowProjectIssueTypeUsageDTO**](WorkflowProjectIssueTypeUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowSchemeUsagesForWorkflow

> WorkflowSchemeUsageDTO GetWorkflowSchemeUsagesForWorkflow(ctx, workflowId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get workflow schemes which are using a given workflow



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
	workflowId := "workflowId_example" // string | The workflow ID
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowSchemeUsagesForWorkflow(context.Background(), workflowId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowSchemeUsagesForWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowSchemeUsagesForWorkflow`: WorkflowSchemeUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowSchemeUsagesForWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowId** | **string** | The workflow ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeUsagesForWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**WorkflowSchemeUsageDTO**](WorkflowSchemeUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowsPaginated

> PageBeanWorkflow GetWorkflowsPaginated(ctx).StartAt(startAt).MaxResults(maxResults).WorkflowName(workflowName).Expand(expand).QueryString(queryString).OrderBy(orderBy).IsActive(isActive).Execute()

Get workflows paginated



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	workflowName := []string{"Inner_example"} // []string | The name of a workflow to return. To include multiple workflows, provide an ampersand-separated list. For example, `workflowName=name1&workflowName=name2`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `transitions` For each workflow, returns information about the transitions inside the workflow.  *  `transitions.rules` For each workflow transition, returns information about its rules. Transitions are included automatically if this expand is requested.  *  `transitions.properties` For each workflow transition, returns information about its properties. Transitions are included automatically if this expand is requested.  *  `statuses` For each workflow, returns information about the statuses inside the workflow.  *  `statuses.properties` For each workflow status, returns information about its properties. Statuses are included automatically if this expand is requested.  *  `default` For each workflow, returns information about whether this is the default workflow.  *  `schemes` For each workflow, returns information about the workflow schemes the workflow is assigned to.  *  `projects` For each workflow, returns information about the projects the workflow is assigned to, through workflow schemes.  *  `hasDraftWorkflow` For each workflow, returns information about whether the workflow has a draft version.  *  `operations` For each workflow, returns information about the actions that can be undertaken on the workflow. (optional)
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with workflow name. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `name` Sorts by workflow name.  *  `created` Sorts by create time.  *  `updated` Sorts by update time. (optional)
	isActive := true // bool | Filters active and inactive workflows. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.GetWorkflowsPaginated(context.Background()).StartAt(startAt).MaxResults(maxResults).WorkflowName(workflowName).Expand(expand).QueryString(queryString).OrderBy(orderBy).IsActive(isActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.GetWorkflowsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowsPaginated`: PageBeanWorkflow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.GetWorkflowsPaginated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **workflowName** | **[]string** | The name of a workflow to return. To include multiple workflows, provide an ampersand-separated list. For example, &#x60;workflowName&#x3D;name1&amp;workflowName&#x3D;name2&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;transitions&#x60; For each workflow, returns information about the transitions inside the workflow.  *  &#x60;transitions.rules&#x60; For each workflow transition, returns information about its rules. Transitions are included automatically if this expand is requested.  *  &#x60;transitions.properties&#x60; For each workflow transition, returns information about its properties. Transitions are included automatically if this expand is requested.  *  &#x60;statuses&#x60; For each workflow, returns information about the statuses inside the workflow.  *  &#x60;statuses.properties&#x60; For each workflow status, returns information about its properties. Statuses are included automatically if this expand is requested.  *  &#x60;default&#x60; For each workflow, returns information about whether this is the default workflow.  *  &#x60;schemes&#x60; For each workflow, returns information about the workflow schemes the workflow is assigned to.  *  &#x60;projects&#x60; For each workflow, returns information about the projects the workflow is assigned to, through workflow schemes.  *  &#x60;hasDraftWorkflow&#x60; For each workflow, returns information about whether the workflow has a draft version.  *  &#x60;operations&#x60; For each workflow, returns information about the actions that can be undertaken on the workflow. | 
 **queryString** | **string** | String used to perform a case-insensitive partial match with workflow name. | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;name&#x60; Sorts by workflow name.  *  &#x60;created&#x60; Sorts by create time.  *  &#x60;updated&#x60; Sorts by update time. | 
 **isActive** | **bool** | Filters active and inactive workflows. | 

### Return type

[**PageBeanWorkflow**](PageBeanWorkflow.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkflowHistory

> WorkflowHistoryListResponseDTO ListWorkflowHistory(ctx).WorkflowHistoryListRequest(workflowHistoryListRequest).Expand(expand).Execute()

List workflow history entries



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
	workflowHistoryListRequest := *openapiclient.NewWorkflowHistoryListRequest() // WorkflowHistoryListRequest | 
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `includeIntermediateWorkflows` Includes intermediate workflow versions that are sometimes created during workflow updates or migrations. By default, these are omitted from the response. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ListWorkflowHistory(context.Background()).WorkflowHistoryListRequest(workflowHistoryListRequest).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ListWorkflowHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListWorkflowHistory`: WorkflowHistoryListResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ListWorkflowHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowHistoryListRequest** | [**WorkflowHistoryListRequest**](WorkflowHistoryListRequest.md) |  | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;includeIntermediateWorkflows&#x60; Includes intermediate workflow versions that are sometimes created during workflow updates or migrations. By default, these are omitted from the response. | 

### Return type

[**WorkflowHistoryListResponseDTO**](WorkflowHistoryListResponseDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadWorkflowFromHistory

> WorkflowHistoryReadResponseDTO ReadWorkflowFromHistory(ctx).WorkflowHistoryReadRequest(workflowHistoryReadRequest).Execute()

Read workflow version from history



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
	workflowHistoryReadRequest := *openapiclient.NewWorkflowHistoryReadRequest() // WorkflowHistoryReadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ReadWorkflowFromHistory(context.Background()).WorkflowHistoryReadRequest(workflowHistoryReadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ReadWorkflowFromHistory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadWorkflowFromHistory`: WorkflowHistoryReadResponseDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ReadWorkflowFromHistory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadWorkflowFromHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowHistoryReadRequest** | [**WorkflowHistoryReadRequest**](WorkflowHistoryReadRequest.md) |  | 

### Return type

[**WorkflowHistoryReadResponseDTO**](WorkflowHistoryReadResponseDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadWorkflowPreviews

> WorkflowPreviewResponse ReadWorkflowPreviews(ctx).WorkflowPreviewRequest(workflowPreviewRequest).Execute()

Preview workflow



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
	workflowPreviewRequest := *openapiclient.NewWorkflowPreviewRequest("ProjectId_example") // WorkflowPreviewRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ReadWorkflowPreviews(context.Background()).WorkflowPreviewRequest(workflowPreviewRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ReadWorkflowPreviews``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadWorkflowPreviews`: WorkflowPreviewResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ReadWorkflowPreviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadWorkflowPreviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowPreviewRequest** | [**WorkflowPreviewRequest**](WorkflowPreviewRequest.md) |  | 

### Return type

[**WorkflowPreviewResponse**](WorkflowPreviewResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadWorkflows

> WorkflowReadResponse ReadWorkflows(ctx).WorkflowReadRequest(workflowReadRequest).Execute()

Bulk get workflows



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
	workflowReadRequest := *openapiclient.NewWorkflowReadRequest() // WorkflowReadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ReadWorkflows(context.Background()).WorkflowReadRequest(workflowReadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ReadWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadWorkflows`: WorkflowReadResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ReadWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowReadRequest** | [**WorkflowReadRequest**](WorkflowReadRequest.md) |  | 

### Return type

[**WorkflowReadResponse**](WorkflowReadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchWorkflows

> WorkflowSearchResponse SearchWorkflows(ctx).StartAt(startAt).MaxResults(maxResults).Expand(expand).QueryString(queryString).OrderBy(orderBy).Scope(scope).IsActive(isActive).Execute()

Search workflows



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `values.transitions` Returns the transitions that each workflow is associated with. (optional)
	queryString := "queryString_example" // string | String used to perform a case-insensitive partial match with workflow name. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `name` Sorts by workflow name.  *  `created` Sorts by create time.  *  `updated` Sorts by update time. (optional)
	scope := "scope_example" // string | The scope of the workflow. Global for company-managed projects and Project for team-managed projects. (optional)
	isActive := true // bool | Filters active and inactive workflows. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.SearchWorkflows(context.Background()).StartAt(startAt).MaxResults(maxResults).Expand(expand).QueryString(queryString).OrderBy(orderBy).Scope(scope).IsActive(isActive).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.SearchWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchWorkflows`: WorkflowSearchResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.SearchWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | 
 **maxResults** | **int32** | The maximum number of items to return per page. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;values.transitions&#x60; Returns the transitions that each workflow is associated with. | 
 **queryString** | **string** | String used to perform a case-insensitive partial match with workflow name. | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;name&#x60; Sorts by workflow name.  *  &#x60;created&#x60; Sorts by create time.  *  &#x60;updated&#x60; Sorts by update time. | 
 **scope** | **string** | The scope of the workflow. Global for company-managed projects and Project for team-managed projects. | 
 **isActive** | **bool** | Filters active and inactive workflows. | 

### Return type

[**WorkflowSearchResponse**](WorkflowSearchResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflows

> WorkflowUpdateResponse UpdateWorkflows(ctx).WorkflowUpdateRequest(workflowUpdateRequest).Execute()

Bulk update workflows



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
	workflowUpdateRequest := *openapiclient.NewWorkflowUpdateRequest() // WorkflowUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.UpdateWorkflows(context.Background()).WorkflowUpdateRequest(workflowUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.UpdateWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflows`: WorkflowUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.UpdateWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowUpdateRequest** | [**WorkflowUpdateRequest**](WorkflowUpdateRequest.md) |  | 

### Return type

[**WorkflowUpdateResponse**](WorkflowUpdateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateCreateWorkflows

> WorkflowValidationErrorList ValidateCreateWorkflows(ctx).WorkflowCreateValidateRequest(workflowCreateValidateRequest).Execute()

Validate create workflows



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
	workflowCreateValidateRequest := *openapiclient.NewWorkflowCreateValidateRequest(*openapiclient.NewWorkflowCreateRequest()) // WorkflowCreateValidateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ValidateCreateWorkflows(context.Background()).WorkflowCreateValidateRequest(workflowCreateValidateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ValidateCreateWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateCreateWorkflows`: WorkflowValidationErrorList
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ValidateCreateWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateCreateWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowCreateValidateRequest** | [**WorkflowCreateValidateRequest**](WorkflowCreateValidateRequest.md) |  | 

### Return type

[**WorkflowValidationErrorList**](WorkflowValidationErrorList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateUpdateWorkflows

> WorkflowValidationErrorList ValidateUpdateWorkflows(ctx).WorkflowUpdateValidateRequestBean(workflowUpdateValidateRequestBean).Execute()

Validate update workflows



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
	workflowUpdateValidateRequestBean := *openapiclient.NewWorkflowUpdateValidateRequestBean(*openapiclient.NewWorkflowUpdateRequest()) // WorkflowUpdateValidateRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.ValidateUpdateWorkflows(context.Background()).WorkflowUpdateValidateRequestBean(workflowUpdateValidateRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.ValidateUpdateWorkflows``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateUpdateWorkflows`: WorkflowValidationErrorList
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.ValidateUpdateWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateUpdateWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowUpdateValidateRequestBean** | [**WorkflowUpdateValidateRequestBean**](WorkflowUpdateValidateRequestBean.md) |  | 

### Return type

[**WorkflowValidationErrorList**](WorkflowValidationErrorList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WorkflowCapabilities

> WorkflowCapabilities WorkflowCapabilities(ctx).WorkflowId(workflowId).ProjectId(projectId).IssueTypeId(issueTypeId).Execute()

Get available workflow capabilities



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
	workflowId := "workflowId_example" // string |  (optional)
	projectId := "projectId_example" // string |  (optional)
	issueTypeId := "issueTypeId_example" // string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowsAPI.WorkflowCapabilities(context.Background()).WorkflowId(workflowId).ProjectId(projectId).IssueTypeId(issueTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsAPI.WorkflowCapabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `WorkflowCapabilities`: WorkflowCapabilities
	fmt.Fprintf(os.Stdout, "Response from `WorkflowsAPI.WorkflowCapabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWorkflowCapabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** |  | 
 **projectId** | **string** |  | 
 **issueTypeId** | **string** |  | 

### Return type

[**WorkflowCapabilities**](WorkflowCapabilities.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

