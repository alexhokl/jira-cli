# \WorkflowSchemeProjectAssociationsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssignSchemeToProject**](WorkflowSchemeProjectAssociationsAPI.md#AssignSchemeToProject) | **Put** /rest/api/3/workflowscheme/project | Assign workflow scheme to project
[**GetWorkflowSchemeProjectAssociations**](WorkflowSchemeProjectAssociationsAPI.md#GetWorkflowSchemeProjectAssociations) | **Get** /rest/api/3/workflowscheme/project | Get workflow scheme project associations



## AssignSchemeToProject

> interface{} AssignSchemeToProject(ctx).WorkflowSchemeProjectAssociation(workflowSchemeProjectAssociation).Execute()

Assign workflow scheme to project



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
	workflowSchemeProjectAssociation := *openapiclient.NewWorkflowSchemeProjectAssociation("ProjectId_example") // WorkflowSchemeProjectAssociation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeProjectAssociationsAPI.AssignSchemeToProject(context.Background()).WorkflowSchemeProjectAssociation(workflowSchemeProjectAssociation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeProjectAssociationsAPI.AssignSchemeToProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignSchemeToProject`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeProjectAssociationsAPI.AssignSchemeToProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssignSchemeToProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowSchemeProjectAssociation** | [**WorkflowSchemeProjectAssociation**](WorkflowSchemeProjectAssociation.md) |  | 

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


## GetWorkflowSchemeProjectAssociations

> ContainerOfWorkflowSchemeAssociations GetWorkflowSchemeProjectAssociations(ctx).ProjectId(projectId).Execute()

Get workflow scheme project associations



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
	projectId := []int64{int64(10010)} // []int64 | The ID of a project to return the workflow schemes for. To include multiple projects, provide an ampersand-Jim: oneseparated list. For example, `projectId=10000&projectId=10001`.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowSchemeProjectAssociationsAPI.GetWorkflowSchemeProjectAssociations(context.Background()).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowSchemeProjectAssociationsAPI.GetWorkflowSchemeProjectAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowSchemeProjectAssociations`: ContainerOfWorkflowSchemeAssociations
	fmt.Fprintf(os.Stdout, "Response from `WorkflowSchemeProjectAssociationsAPI.GetWorkflowSchemeProjectAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowSchemeProjectAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | The ID of a project to return the workflow schemes for. To include multiple projects, provide an ampersand-Jim: oneseparated list. For example, &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 

### Return type

[**ContainerOfWorkflowSchemeAssociations**](ContainerOfWorkflowSchemeAssociations.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

