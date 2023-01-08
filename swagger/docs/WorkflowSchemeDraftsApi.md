# \WorkflowSchemeDraftsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowSchemeDraftFromParent**](WorkflowSchemeDraftsAPI.md#CreateWorkflowSchemeDraftFromParent) | **Post** /rest/api/3/workflowscheme/{id}/createdraft | Create draft workflow scheme
[**DeleteDraftDefaultWorkflow**](WorkflowSchemeDraftsAPI.md#DeleteDraftDefaultWorkflow) | **Delete** /rest/api/3/workflowscheme/{id}/draft/default | Delete draft default workflow
[**DeleteDraftWorkflowMapping**](WorkflowSchemeDraftsAPI.md#DeleteDraftWorkflowMapping) | **Delete** /rest/api/3/workflowscheme/{id}/draft/workflow | Delete issue types for workflow in draft workflow scheme
[**DeleteWorkflowSchemeDraft**](WorkflowSchemeDraftsAPI.md#DeleteWorkflowSchemeDraft) | **Delete** /rest/api/3/workflowscheme/{id}/draft | Delete draft workflow scheme
[**DeleteWorkflowSchemeDraftIssueType**](WorkflowSchemeDraftsAPI.md#DeleteWorkflowSchemeDraftIssueType) | **Delete** /rest/api/3/workflowscheme/{id}/draft/issuetype/{issueType} | Delete workflow for issue type in draft workflow scheme
[**GetDraftDefaultWorkflow**](WorkflowSchemeDraftsAPI.md#GetDraftDefaultWorkflow) | **Get** /rest/api/3/workflowscheme/{id}/draft/default | Get draft default workflow
[**GetDraftWorkflow**](WorkflowSchemeDraftsAPI.md#GetDraftWorkflow) | **Get** /rest/api/3/workflowscheme/{id}/draft/workflow | Get issue types for workflows in draft workflow scheme
[**GetWorkflowSchemeDraft**](WorkflowSchemeDraftsAPI.md#GetWorkflowSchemeDraft) | **Get** /rest/api/3/workflowscheme/{id}/draft | Get draft workflow scheme
[**GetWorkflowSchemeDraftIssueType**](WorkflowSchemeDraftsAPI.md#GetWorkflowSchemeDraftIssueType) | **Get** /rest/api/3/workflowscheme/{id}/draft/issuetype/{issueType} | Get workflow for issue type in draft workflow scheme
[**PublishDraftWorkflowScheme**](WorkflowSchemeDraftsAPI.md#PublishDraftWorkflowScheme) | **Post** /rest/api/3/workflowscheme/{id}/draft/publish | Publish draft workflow scheme
[**SetWorkflowSchemeDraftIssueType**](WorkflowSchemeDraftsAPI.md#SetWorkflowSchemeDraftIssueType) | **Put** /rest/api/3/workflowscheme/{id}/draft/issuetype/{issueType} | Set workflow for issue type in draft workflow scheme
[**UpdateDraftDefaultWorkflow**](WorkflowSchemeDraftsAPI.md#UpdateDraftDefaultWorkflow) | **Put** /rest/api/3/workflowscheme/{id}/draft/default | Update draft default workflow
[**UpdateDraftWorkflowMapping**](WorkflowSchemeDraftsAPI.md#UpdateDraftWorkflowMapping) | **Put** /rest/api/3/workflowscheme/{id}/draft/workflow | Set issue types for workflow in workflow scheme
[**UpdateWorkflowSchemeDraft**](WorkflowSchemeDraftsAPI.md#UpdateWorkflowSchemeDraft) | **Put** /rest/api/3/workflowscheme/{id}/draft | Update draft workflow scheme



## CreateWorkflowSchemeDraftFromParent

> WorkflowScheme CreateWorkflowSchemeDraftFromParent(ctx, id).Execute()

Create draft workflow scheme



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
	id := int64(789) // int64 | The ID of the active workflow scheme that the draft is created from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.CreateWorkflowSchemeDraftFromParent(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.CreateWorkflowSchemeDraftFromParent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowSchemeDraftFromParent`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.CreateWorkflowSchemeDraftFromParent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the active workflow scheme that the draft is created from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowSchemeDraftFromParentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteDraftDefaultWorkflow

> WorkflowScheme DeleteDraftDefaultWorkflow(ctx, id).Execute()

Delete draft default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.DeleteDraftDefaultWorkflow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.DeleteDraftDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDraftDefaultWorkflow`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.DeleteDraftDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDraftDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## DeleteDraftWorkflowMapping

> DeleteDraftWorkflowMapping(ctx, id).WorkflowName(workflowName).Execute()

Delete issue types for workflow in draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	workflowName := "workflowName_example" // string | The name of the workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowSchemeDraftsAPI.DeleteDraftWorkflowMapping(context.Background(), id).WorkflowName(workflowName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.DeleteDraftWorkflowMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDraftWorkflowMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of the workflow. | 

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


## DeleteWorkflowSchemeDraft

> DeleteWorkflowSchemeDraft(ctx, id).Execute()

Delete draft workflow scheme



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
	id := int64(789) // int64 | The ID of the active workflow scheme that the draft was created from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowSchemeDraftsAPI.DeleteWorkflowSchemeDraft(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.DeleteWorkflowSchemeDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the active workflow scheme that the draft was created from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowSchemeDraftRequest struct via the builder pattern


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


## DeleteWorkflowSchemeDraftIssueType

> WorkflowScheme DeleteWorkflowSchemeDraftIssueType(ctx, id, issueType).Execute()

Delete workflow for issue type in draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	issueType := "issueType_example" // string | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.DeleteWorkflowSchemeDraftIssueType(context.Background(), id, issueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.DeleteWorkflowSchemeDraftIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkflowSchemeDraftIssueType`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.DeleteWorkflowSchemeDraftIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowSchemeDraftIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## GetDraftDefaultWorkflow

> DefaultWorkflow GetDraftDefaultWorkflow(ctx, id).Execute()

Get draft default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.GetDraftDefaultWorkflow(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.GetDraftDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftDefaultWorkflow`: DefaultWorkflow
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.GetDraftDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetDraftWorkflow

> IssueTypesWorkflowMapping GetDraftWorkflow(ctx, id).WorkflowName(workflowName).Execute()

Get issue types for workflows in draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	workflowName := "workflowName_example" // string | The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.GetDraftWorkflow(context.Background(), id).WorkflowName(workflowName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.GetDraftWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDraftWorkflow`: IssueTypesWorkflowMapping
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.GetDraftWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDraftWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of a workflow in the scheme. Limits the results to the workflow-issue type mapping for the specified workflow. | 

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


## GetWorkflowSchemeDraft

> WorkflowScheme GetWorkflowSchemeDraft(ctx, id).Execute()

Get draft workflow scheme



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
	id := int64(789) // int64 | The ID of the active workflow scheme that the draft was created from.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraft(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowSchemeDraft`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the active workflow scheme that the draft was created from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetWorkflowSchemeDraftIssueType

> IssueTypeWorkflowMapping GetWorkflowSchemeDraftIssueType(ctx, id, issueType).Execute()

Get workflow for issue type in draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	issueType := "issueType_example" // string | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraftIssueType(context.Background(), id, issueType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraftIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowSchemeDraftIssueType`: IssueTypeWorkflowMapping
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.GetWorkflowSchemeDraftIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeDraftIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## PublishDraftWorkflowScheme

> PublishDraftWorkflowScheme(ctx, id).PublishDraftWorkflowScheme(publishDraftWorkflowScheme).ValidateOnly(validateOnly).Execute()

Publish draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	publishDraftWorkflowScheme := *openapiclient.NewPublishDraftWorkflowScheme() // PublishDraftWorkflowScheme | Details of the status mappings.
	validateOnly := true // bool | Whether the request only performs a validation. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowSchemeDraftsAPI.PublishDraftWorkflowScheme(context.Background(), id).PublishDraftWorkflowScheme(publishDraftWorkflowScheme).ValidateOnly(validateOnly).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.PublishDraftWorkflowScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishDraftWorkflowSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **publishDraftWorkflowScheme** | [**PublishDraftWorkflowScheme**](PublishDraftWorkflowScheme.md) | Details of the status mappings. | 
 **validateOnly** | **bool** | Whether the request only performs a validation. | [default to false]

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


## SetWorkflowSchemeDraftIssueType

> WorkflowScheme SetWorkflowSchemeDraftIssueType(ctx, id, issueType).IssueTypeWorkflowMapping(issueTypeWorkflowMapping).Execute()

Set workflow for issue type in draft workflow scheme



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	issueType := "issueType_example" // string | The ID of the issue type.
	issueTypeWorkflowMapping := *openapiclient.NewIssueTypeWorkflowMapping() // IssueTypeWorkflowMapping | The issue type-project mapping.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.SetWorkflowSchemeDraftIssueType(context.Background(), id, issueType).IssueTypeWorkflowMapping(issueTypeWorkflowMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.SetWorkflowSchemeDraftIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorkflowSchemeDraftIssueType`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.SetWorkflowSchemeDraftIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 
**issueType** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorkflowSchemeDraftIssueTypeRequest struct via the builder pattern


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


## UpdateDraftDefaultWorkflow

> WorkflowScheme UpdateDraftDefaultWorkflow(ctx, id).DefaultWorkflow(defaultWorkflow).Execute()

Update draft default workflow



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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	defaultWorkflow := *openapiclient.NewDefaultWorkflow("Workflow_example") // DefaultWorkflow | The object for the new default workflow.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.UpdateDraftDefaultWorkflow(context.Background(), id).DefaultWorkflow(defaultWorkflow).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.UpdateDraftDefaultWorkflow``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraftDefaultWorkflow`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.UpdateDraftDefaultWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftDefaultWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **defaultWorkflow** | [**DefaultWorkflow**](DefaultWorkflow.md) | The object for the new default workflow. | 

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


## UpdateDraftWorkflowMapping

> WorkflowScheme UpdateDraftWorkflowMapping(ctx, id).WorkflowName(workflowName).IssueTypesWorkflowMapping(issueTypesWorkflowMapping).Execute()

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
	id := int64(789) // int64 | The ID of the workflow scheme that the draft belongs to.
	workflowName := "workflowName_example" // string | The name of the workflow.
	issueTypesWorkflowMapping := *openapiclient.NewIssueTypesWorkflowMapping() // IssueTypesWorkflowMapping | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.UpdateDraftWorkflowMapping(context.Background(), id).WorkflowName(workflowName).IssueTypesWorkflowMapping(issueTypesWorkflowMapping).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.UpdateDraftWorkflowMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDraftWorkflowMapping`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.UpdateDraftWorkflowMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the workflow scheme that the draft belongs to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDraftWorkflowMappingRequest struct via the builder pattern


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


## UpdateWorkflowSchemeDraft

> WorkflowScheme UpdateWorkflowSchemeDraft(ctx, id).WorkflowScheme(workflowScheme).Execute()

Update draft workflow scheme



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
	id := int64(789) // int64 | The ID of the active workflow scheme that the draft was created from.
	workflowScheme := *openapiclient.NewWorkflowScheme() // WorkflowScheme | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeDraftsAPI.UpdateWorkflowSchemeDraft(context.Background(), id).WorkflowScheme(workflowScheme).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeDraftsAPI.UpdateWorkflowSchemeDraft``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowSchemeDraft`: WorkflowScheme
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeDraftsAPI.UpdateWorkflowSchemeDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the active workflow scheme that the draft was created from. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowSchemeDraftRequest struct via the builder pattern


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

