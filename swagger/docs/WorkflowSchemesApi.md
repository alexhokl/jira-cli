# \WorkflowSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowScheme**](WorkflowSchemesAPI.md#CreateWorkflowScheme) | **Post** /rest/api/3/workflowscheme | Create workflow scheme
[**DeleteDefaultWorkflow**](WorkflowSchemesAPI.md#DeleteDefaultWorkflow) | **Delete** /rest/api/3/workflowscheme/{id}/default | Delete default workflow
[**DeleteWorkflowMapping**](WorkflowSchemesAPI.md#DeleteWorkflowMapping) | **Delete** /rest/api/3/workflowscheme/{id}/workflow | Delete issue types for workflow in workflow scheme
[**DeleteWorkflowScheme**](WorkflowSchemesAPI.md#DeleteWorkflowScheme) | **Delete** /rest/api/3/workflowscheme/{id} | Delete workflow scheme
[**DeleteWorkflowSchemeIssueType**](WorkflowSchemesAPI.md#DeleteWorkflowSchemeIssueType) | **Delete** /rest/api/3/workflowscheme/{id}/issuetype/{issueType} | Delete workflow for issue type in workflow scheme
[**GetAllWorkflowSchemes**](WorkflowSchemesAPI.md#GetAllWorkflowSchemes) | **Get** /rest/api/3/workflowscheme | Get all workflow schemes
[**GetDefaultWorkflow**](WorkflowSchemesAPI.md#GetDefaultWorkflow) | **Get** /rest/api/3/workflowscheme/{id}/default | Get default workflow
[**GetProjectUsagesForWorkflowScheme**](WorkflowSchemesAPI.md#GetProjectUsagesForWorkflowScheme) | **Get** /rest/api/3/workflowscheme/{workflowSchemeId}/projectUsages | Get projects which are using a given workflow scheme
[**GetRequiredWorkflowSchemeMappings**](WorkflowSchemesAPI.md#GetRequiredWorkflowSchemeMappings) | **Post** /rest/api/3/workflowscheme/update/mappings | Get required status mappings for workflow scheme update
[**GetWorkflow**](WorkflowSchemesAPI.md#GetWorkflow) | **Get** /rest/api/3/workflowscheme/{id}/workflow | Get issue types for workflows in workflow scheme
[**GetWorkflowScheme**](WorkflowSchemesAPI.md#GetWorkflowScheme) | **Get** /rest/api/3/workflowscheme/{id} | Get workflow scheme
[**GetWorkflowSchemeIssueType**](WorkflowSchemesAPI.md#GetWorkflowSchemeIssueType) | **Get** /rest/api/3/workflowscheme/{id}/issuetype/{issueType} | Get workflow for issue type in workflow scheme
[**ReadWorkflowSchemes**](WorkflowSchemesAPI.md#ReadWorkflowSchemes) | **Post** /rest/api/3/workflowscheme/read | Bulk get workflow schemes
[**SetWorkflowSchemeIssueType**](WorkflowSchemesAPI.md#SetWorkflowSchemeIssueType) | **Put** /rest/api/3/workflowscheme/{id}/issuetype/{issueType} | Set workflow for issue type in workflow scheme
[**SwitchWorkflowSchemeForProject**](WorkflowSchemesAPI.md#SwitchWorkflowSchemeForProject) | **Post** /rest/api/3/workflowscheme/project/switch | Switch workflow scheme for project
[**UpdateDefaultWorkflow**](WorkflowSchemesAPI.md#UpdateDefaultWorkflow) | **Put** /rest/api/3/workflowscheme/{id}/default | Update default workflow
[**UpdateSchemes**](WorkflowSchemesAPI.md#UpdateSchemes) | **Post** /rest/api/3/workflowscheme/update | Update workflow scheme
[**UpdateWorkflowMapping**](WorkflowSchemesAPI.md#UpdateWorkflowMapping) | **Put** /rest/api/3/workflowscheme/{id}/workflow | Set issue types for workflow in workflow scheme
[**UpdateWorkflowScheme**](WorkflowSchemesAPI.md#UpdateWorkflowScheme) | **Put** /rest/api/3/workflowscheme/{id} | Classic update workflow scheme



## CreateWorkflowScheme

> WorkflowScheme CreateWorkflowScheme(ctx).WorkflowScheme(workflowScheme).Execute()

Create workflow scheme



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
	workflowScheme := *openapiclient.NewWorkflowScheme() // WorkflowScheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.CreateWorkflowScheme(context.Background()).WorkflowScheme(workflowScheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.CreateWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowScheme`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.CreateWorkflowScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowScheme** | [**WorkflowScheme**](WorkflowScheme.md) |  | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDefaultWorkflow

> WorkflowScheme DeleteDefaultWorkflow(ctx, id).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()

Delete default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	updateDraftIfNeeded := true // bool | Set to true to create or update the draft of a workflow scheme and delete the mapping from the draft, when the workflow scheme cannot be edited. Defaults to `false`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.DeleteDefaultWorkflow(context.Background(), id).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.DeleteDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDefaultWorkflow`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.DeleteDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateDraftIfNeeded** | **bool** | Set to true to create or update the draft of a workflow scheme and delete the mapping from the draft, when the workflow scheme cannot be edited. Defaults to &#x60;false&#x60;. | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowMapping

> DeleteWorkflowMapping(ctx, id).WorkflowName(workflowName).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()

Delete issue types for workflow in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	workflowName := "workflowName_example" // string | The name of the workflow.
	updateDraftIfNeeded := true // bool | Set to true to create or update the draft of a workflow scheme and delete the mapping from the draft, when the workflow scheme cannot be edited. Defaults to `false`. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowSchemesAPI.DeleteWorkflowMapping(context.Background(), id).WorkflowName(workflowName).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.DeleteWorkflowMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of the workflow. | 
 **updateDraftIfNeeded** | **bool** | Set to true to create or update the draft of a workflow scheme and delete the mapping from the draft, when the workflow scheme cannot be edited. Defaults to &#x60;false&#x60;. | [default to false]

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


## DeleteWorkflowScheme

> interface{} DeleteWorkflowScheme(ctx, id).Execute()

Delete workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as `schemeId`. For example, *schemeId=10301*.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.DeleteWorkflowScheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.DeleteWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkflowScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.DeleteWorkflowScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as &#x60;schemeId&#x60;. For example, *schemeId&#x3D;10301*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowSchemeRequest struct via the builder pattern


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


## DeleteWorkflowSchemeIssueType

> WorkflowScheme DeleteWorkflowSchemeIssueType(ctx, id, issueType).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()

Delete workflow for issue type in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	issueType := "issueType_example" // string | The ID of the issue type.
	updateDraftIfNeeded := true // bool | Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to `false`. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.DeleteWorkflowSchemeIssueType(context.Background(), id, issueType).UpdateDraftIfNeeded(updateDraftIfNeeded).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.DeleteWorkflowSchemeIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkflowSchemeIssueType`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.DeleteWorkflowSchemeIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowSchemeIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateDraftIfNeeded** | **bool** | Set to true to create or update the draft of a workflow scheme and update the mapping in the draft, when the workflow scheme cannot be edited. Defaults to &#x60;false&#x60;. | [default to false]

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllWorkflowSchemes

> PageBeanWorkflowScheme GetAllWorkflowSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Execute()

Get all workflow schemes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetAllWorkflowSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetAllWorkflowSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllWorkflowSchemes`: PageBeanWorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetAllWorkflowSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllWorkflowSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanWorkflowScheme**](PageBeanWorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultWorkflow

> DefaultWorkflow GetDefaultWorkflow(ctx, id).ReturnDraftIfExists(returnDraftIfExists).Execute()

Get default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	returnDraftIfExists := true // bool | Set to `true` to return the default workflow for the workflow scheme's draft rather than scheme itself. If the workflow scheme does not have a draft, then the default workflow for the workflow scheme is returned. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetDefaultWorkflow(context.Background(), id).ReturnDraftIfExists(returnDraftIfExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultWorkflow`: DefaultWorkflow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDraftIfExists** | **bool** | Set to &#x60;true&#x60; to return the default workflow for the workflow scheme&#39;s draft rather than scheme itself. If the workflow scheme does not have a draft, then the default workflow for the workflow scheme is returned. | [default to false]

### Return type

[**DefaultWorkflow**](DefaultWorkflow.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectUsagesForWorkflowScheme

> WorkflowSchemeProjectUsageDTO GetProjectUsagesForWorkflowScheme(ctx, workflowSchemeId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()

Get projects which are using a given workflow scheme



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
	workflowSchemeId := "workflowSchemeId_example" // string | The workflow scheme ID
	nextPageToken := "nextPageToken_example" // string | The cursor for pagination (optional)
	maxResults := int32(56) // int32 | The maximum number of results to return. Must be an integer between 1 and 200. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetProjectUsagesForWorkflowScheme(context.Background(), workflowSchemeId).NextPageToken(nextPageToken).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetProjectUsagesForWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectUsagesForWorkflowScheme`: WorkflowSchemeProjectUsageDTO
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetProjectUsagesForWorkflowScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workflowSchemeId** | **string** | The workflow scheme ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectUsagesForWorkflowSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **nextPageToken** | **string** | The cursor for pagination | 
 **maxResults** | **int32** | The maximum number of results to return. Must be an integer between 1 and 200. | [default to 50]

### Return type

[**WorkflowSchemeProjectUsageDTO**](WorkflowSchemeProjectUsageDTO.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRequiredWorkflowSchemeMappings

> WorkflowSchemeUpdateRequiredMappingsResponse GetRequiredWorkflowSchemeMappings(ctx).WorkflowSchemeUpdateRequiredMappingsRequest(workflowSchemeUpdateRequiredMappingsRequest).Execute()

Get required status mappings for workflow scheme update



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
	workflowSchemeUpdateRequiredMappingsRequest := *openapiclient.NewWorkflowSchemeUpdateRequiredMappingsRequest("Id_example", []openapiclient.WorkflowSchemeAssociation{*openapiclient.NewWorkflowSchemeAssociation([]string{"IssueTypeIds_example"}, "WorkflowId_example")}) // WorkflowSchemeUpdateRequiredMappingsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetRequiredWorkflowSchemeMappings(context.Background()).WorkflowSchemeUpdateRequiredMappingsRequest(workflowSchemeUpdateRequiredMappingsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetRequiredWorkflowSchemeMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRequiredWorkflowSchemeMappings`: WorkflowSchemeUpdateRequiredMappingsResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetRequiredWorkflowSchemeMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetRequiredWorkflowSchemeMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSchemeUpdateRequiredMappingsRequest** | [**WorkflowSchemeUpdateRequiredMappingsRequest**](WorkflowSchemeUpdateRequiredMappingsRequest.md) |  | 

### Return type

[**WorkflowSchemeUpdateRequiredMappingsResponse**](WorkflowSchemeUpdateRequiredMappingsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflow

> IssueTypesWorkflowMapping GetWorkflow(ctx, id).WorkflowName(workflowName).ReturnDraftIfExists(returnDraftIfExists).Execute()

Get issue types for workflows in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	workflowName := "workflowName_example" // string | The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow. (optional)
	returnDraftIfExists := true // bool | Returns the mapping from the workflow scheme's draft rather than the workflow scheme, if set to true. If no draft exists, the mapping from the workflow scheme is returned. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetWorkflow(context.Background(), id).WorkflowName(workflowName).ReturnDraftIfExists(returnDraftIfExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflow`: IssueTypesWorkflowMapping
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow. | 
 **returnDraftIfExists** | **bool** | Returns the mapping from the workflow scheme&#39;s draft rather than the workflow scheme, if set to true. If no draft exists, the mapping from the workflow scheme is returned. | [default to false]

### Return type

[**IssueTypesWorkflowMapping**](IssueTypesWorkflowMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowScheme

> WorkflowScheme GetWorkflowScheme(ctx, id).ReturnDraftIfExists(returnDraftIfExists).Execute()

Get workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as `schemeId`. For example, *schemeId=10301*.
	returnDraftIfExists := true // bool | Returns the workflow scheme's draft rather than scheme itself, if set to true. If the workflow scheme does not have a draft, then the workflow scheme is returned. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetWorkflowScheme(context.Background(), id).ReturnDraftIfExists(returnDraftIfExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowScheme`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetWorkflowScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as &#x60;schemeId&#x60;. For example, *schemeId&#x3D;10301*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **returnDraftIfExists** | **bool** | Returns the workflow scheme&#39;s draft rather than scheme itself, if set to true. If the workflow scheme does not have a draft, then the workflow scheme is returned. | [default to false]

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowSchemeIssueType

> IssueTypeWorkflowMapping GetWorkflowSchemeIssueType(ctx, id, issueType).ReturnDraftIfExists(returnDraftIfExists).Execute()

Get workflow for issue type in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	issueType := "issueType_example" // string | The ID of the issue type.
	returnDraftIfExists := true // bool | Returns the mapping from the workflow scheme's draft rather than the workflow scheme, if set to true. If no draft exists, the mapping from the workflow scheme is returned. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.GetWorkflowSchemeIssueType(context.Background(), id, issueType).ReturnDraftIfExists(returnDraftIfExists).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.GetWorkflowSchemeIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowSchemeIssueType`: IssueTypeWorkflowMapping
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.GetWorkflowSchemeIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **returnDraftIfExists** | **bool** | Returns the mapping from the workflow scheme&#39;s draft rather than the workflow scheme, if set to true. If no draft exists, the mapping from the workflow scheme is returned. | [default to false]

### Return type

[**IssueTypeWorkflowMapping**](IssueTypeWorkflowMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReadWorkflowSchemes

> []WorkflowSchemeReadResponse ReadWorkflowSchemes(ctx).WorkflowSchemeReadRequest(workflowSchemeReadRequest).Execute()

Bulk get workflow schemes



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
	workflowSchemeReadRequest := *openapiclient.NewWorkflowSchemeReadRequest() // WorkflowSchemeReadRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.ReadWorkflowSchemes(context.Background()).WorkflowSchemeReadRequest(workflowSchemeReadRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.ReadWorkflowSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReadWorkflowSchemes`: []WorkflowSchemeReadResponse
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.ReadWorkflowSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiReadWorkflowSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSchemeReadRequest** | [**WorkflowSchemeReadRequest**](WorkflowSchemeReadRequest.md) |  | 

### Return type

[**[]WorkflowSchemeReadResponse**](WorkflowSchemeReadResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetWorkflowSchemeIssueType

> WorkflowScheme SetWorkflowSchemeIssueType(ctx, id, issueType).IssueTypeWorkflowMapping(issueTypeWorkflowMapping).Execute()

Set workflow for issue type in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	issueType := "issueType_example" // string | The ID of the issue type.
	issueTypeWorkflowMapping := *openapiclient.NewIssueTypeWorkflowMapping() // IssueTypeWorkflowMapping | The issue type-project mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.SetWorkflowSchemeIssueType(context.Background(), id, issueType).IssueTypeWorkflowMapping(issueTypeWorkflowMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.SetWorkflowSchemeIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorkflowSchemeIssueType`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.SetWorkflowSchemeIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorkflowSchemeIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueTypeWorkflowMapping** | [**IssueTypeWorkflowMapping**](IssueTypeWorkflowMapping.md) | The issue type-project mapping. | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SwitchWorkflowSchemeForProject

> SwitchWorkflowSchemeForProject(ctx).WorkflowSchemeProjectSwitchBean(workflowSchemeProjectSwitchBean).Execute()

Switch workflow scheme for project



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
	workflowSchemeProjectSwitchBean := *openapiclient.NewWorkflowSchemeProjectSwitchBean() // WorkflowSchemeProjectSwitchBean | The request containing project ID, target scheme ID, and any issue type mappings.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowSchemesAPI.SwitchWorkflowSchemeForProject(context.Background()).WorkflowSchemeProjectSwitchBean(workflowSchemeProjectSwitchBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.SwitchWorkflowSchemeForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSwitchWorkflowSchemeForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSchemeProjectSwitchBean** | [**WorkflowSchemeProjectSwitchBean**](WorkflowSchemeProjectSwitchBean.md) | The request containing project ID, target scheme ID, and any issue type mappings. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDefaultWorkflow

> WorkflowScheme UpdateDefaultWorkflow(ctx, id).DefaultWorkflow(defaultWorkflow).Execute()

Update default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	defaultWorkflow := *openapiclient.NewDefaultWorkflow("Workflow_example") // DefaultWorkflow | The new default workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.UpdateDefaultWorkflow(context.Background(), id).DefaultWorkflow(defaultWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.UpdateDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDefaultWorkflow`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.UpdateDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultWorkflow** | [**DefaultWorkflow**](DefaultWorkflow.md) | The new default workflow. | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSchemes

> interface{} UpdateSchemes(ctx).WorkflowSchemeUpdateRequest(workflowSchemeUpdateRequest).Execute()

Update workflow scheme



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
	workflowSchemeUpdateRequest := *openapiclient.NewWorkflowSchemeUpdateRequest("Description_example", "Id_example", "Name_example", *openapiclient.NewDocumentVersion()) // WorkflowSchemeUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.UpdateSchemes(context.Background()).WorkflowSchemeUpdateRequest(workflowSchemeUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.UpdateSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSchemes`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.UpdateSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSchemeUpdateRequest** | [**WorkflowSchemeUpdateRequest**](WorkflowSchemeUpdateRequest.md) |  | 

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


## UpdateWorkflowMapping

> WorkflowScheme UpdateWorkflowMapping(ctx, id).WorkflowName(workflowName).IssueTypesWorkflowMapping(issueTypesWorkflowMapping).Execute()

Set issue types for workflow in workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme.
	workflowName := "workflowName_example" // string | The name of the workflow.
	issueTypesWorkflowMapping := *openapiclient.NewIssueTypesWorkflowMapping() // IssueTypesWorkflowMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.UpdateWorkflowMapping(context.Background(), id).WorkflowName(workflowName).IssueTypesWorkflowMapping(issueTypesWorkflowMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.UpdateWorkflowMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowMapping`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.UpdateWorkflowMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of the workflow. | 
 **issueTypesWorkflowMapping** | [**IssueTypesWorkflowMapping**](IssueTypesWorkflowMapping.md) |  | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowScheme

> WorkflowScheme UpdateWorkflowScheme(ctx, id).WorkflowScheme(workflowScheme).Execute()

Classic update workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as `schemeId`. For example, *schemeId=10301*.
	workflowScheme := *openapiclient.NewWorkflowScheme() // WorkflowScheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemesAPI.UpdateWorkflowScheme(context.Background(), id).WorkflowScheme(workflowScheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemesAPI.UpdateWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowScheme`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemesAPI.UpdateWorkflowScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme. Find this ID by editing the desired workflow scheme in Jira. The ID is shown in the URL as &#x60;schemeId&#x60;. For example, *schemeId&#x3D;10301*. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowScheme** | [**WorkflowScheme**](WorkflowScheme.md) |  | 

### Return type

[**WorkflowScheme**](WorkflowScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

