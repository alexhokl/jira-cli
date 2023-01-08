# \WorkflowTransitionPropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowTransitionProperty**](WorkflowTransitionPropertiesAPI.md#CreateWorkflowTransitionProperty) | **Post** /rest/api/3/workflow/transitions/{transitionId}/properties | Create workflow transition property
[**DeleteWorkflowTransitionProperty**](WorkflowTransitionPropertiesAPI.md#DeleteWorkflowTransitionProperty) | **Delete** /rest/api/3/workflow/transitions/{transitionId}/properties | Delete workflow transition property
[**GetWorkflowTransitionProperties**](WorkflowTransitionPropertiesAPI.md#GetWorkflowTransitionProperties) | **Get** /rest/api/3/workflow/transitions/{transitionId}/properties | Get workflow transition properties
[**UpdateWorkflowTransitionProperty**](WorkflowTransitionPropertiesAPI.md#UpdateWorkflowTransitionProperty) | **Put** /rest/api/3/workflow/transitions/{transitionId}/properties | Update workflow transition property



## CreateWorkflowTransitionProperty

> WorkflowTransitionProperty CreateWorkflowTransitionProperty(ctx, transitionId).Key(key).WorkflowName(workflowName).WorkflowTransitionProperty(workflowTransitionProperty).WorkflowMode(workflowMode).Execute()

Create workflow transition property



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
	transitionId := int64(789) // int64 | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition.
	key := "key_example" // string | The key of the property being added, also known as the name of the property. Set this to the same value as the `key` defined in the request body.
	workflowName := "workflowName_example" // string | The name of the workflow that the transition belongs to.
	workflowTransitionProperty := *openapiclient.NewWorkflowTransitionProperty("Value_example") // WorkflowTransitionProperty | 
	workflowMode := "workflowMode_example" // string | The workflow status. Set to *live* for inactive workflows or *draft* for draft workflows. Active workflows cannot be edited. (optional) (default to "live")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionPropertiesAPI.CreateWorkflowTransitionProperty(context.Background(), transitionId).Key(key).WorkflowName(workflowName).WorkflowTransitionProperty(workflowTransitionProperty).WorkflowMode(workflowMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionPropertiesAPI.CreateWorkflowTransitionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWorkflowTransitionProperty`: WorkflowTransitionProperty
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionPropertiesAPI.CreateWorkflowTransitionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **int64** | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTransitionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | The key of the property being added, also known as the name of the property. Set this to the same value as the &#x60;key&#x60; defined in the request body. | 
 **workflowName** | **string** | The name of the workflow that the transition belongs to. | 
 **workflowTransitionProperty** | [**WorkflowTransitionProperty**](WorkflowTransitionProperty.md) |  | 
 **workflowMode** | **string** | The workflow status. Set to *live* for inactive workflows or *draft* for draft workflows. Active workflows cannot be edited. | [default to &quot;live&quot;]

### Return type

[**WorkflowTransitionProperty**](WorkflowTransitionProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkflowTransitionProperty

> DeleteWorkflowTransitionProperty(ctx, transitionId).Key(key).WorkflowName(workflowName).WorkflowMode(workflowMode).Execute()

Delete workflow transition property



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
	transitionId := int64(789) // int64 | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition.
	key := "key_example" // string | The name of the transition property to delete, also known as the name of the property.
	workflowName := "workflowName_example" // string | The name of the workflow that the transition belongs to.
	workflowMode := "workflowMode_example" // string | The workflow status. Set to `live` for inactive workflows or `draft` for draft workflows. Active workflows cannot be edited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WorkflowTransitionPropertiesAPI.DeleteWorkflowTransitionProperty(context.Background(), transitionId).Key(key).WorkflowName(workflowName).WorkflowMode(workflowMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionPropertiesAPI.DeleteWorkflowTransitionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **int64** | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowTransitionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | The name of the transition property to delete, also known as the name of the property. | 
 **workflowName** | **string** | The name of the workflow that the transition belongs to. | 
 **workflowMode** | **string** | The workflow status. Set to &#x60;live&#x60; for inactive workflows or &#x60;draft&#x60; for draft workflows. Active workflows cannot be edited. | 

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


## GetWorkflowTransitionProperties

> WorkflowTransitionProperty GetWorkflowTransitionProperties(ctx, transitionId).WorkflowName(workflowName).IncludeReservedKeys(includeReservedKeys).Key(key).WorkflowMode(workflowMode).Execute()

Get workflow transition properties



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
	transitionId := int64(789) // int64 | The ID of the transition. To get the ID, view the workflow in text mode in the Jira administration console. The ID is shown next to the transition.
	workflowName := "workflowName_example" // string | The name of the workflow that the transition belongs to.
	includeReservedKeys := true // bool | Some properties with keys that have the *jira.* prefix are reserved, which means they are not editable. To include these properties in the results, set this parameter to *true*. (optional) (default to false)
	key := "key_example" // string | The key of the property being returned, also known as the name of the property. If this parameter is not specified, all properties on the transition are returned. (optional)
	workflowMode := "workflowMode_example" // string | The workflow status. Set to *live* for active and inactive workflows, or *draft* for draft workflows. (optional) (default to "live")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionPropertiesAPI.GetWorkflowTransitionProperties(context.Background(), transitionId).WorkflowName(workflowName).IncludeReservedKeys(includeReservedKeys).Key(key).WorkflowMode(workflowMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionPropertiesAPI.GetWorkflowTransitionProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowTransitionProperties`: WorkflowTransitionProperty
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionPropertiesAPI.GetWorkflowTransitionProperties`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **int64** | The ID of the transition. To get the ID, view the workflow in text mode in the Jira administration console. The ID is shown next to the transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTransitionPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **workflowName** | **string** | The name of the workflow that the transition belongs to. | 
 **includeReservedKeys** | **bool** | Some properties with keys that have the *jira.* prefix are reserved, which means they are not editable. To include these properties in the results, set this parameter to *true*. | [default to false]
 **key** | **string** | The key of the property being returned, also known as the name of the property. If this parameter is not specified, all properties on the transition are returned. | 
 **workflowMode** | **string** | The workflow status. Set to *live* for active and inactive workflows, or *draft* for draft workflows. | [default to &quot;live&quot;]

### Return type

[**WorkflowTransitionProperty**](WorkflowTransitionProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTransitionProperty

> WorkflowTransitionProperty UpdateWorkflowTransitionProperty(ctx, transitionId).Key(key).WorkflowName(workflowName).WorkflowTransitionProperty(workflowTransitionProperty).WorkflowMode(workflowMode).Execute()

Update workflow transition property



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
	transitionId := int64(789) // int64 | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition.
	key := "key_example" // string | The key of the property being updated, also known as the name of the property. Set this to the same value as the `key` defined in the request body.
	workflowName := "workflowName_example" // string | The name of the workflow that the transition belongs to.
	workflowTransitionProperty := *openapiclient.NewWorkflowTransitionProperty("Value_example") // WorkflowTransitionProperty | 
	workflowMode := "workflowMode_example" // string | The workflow status. Set to `live` for inactive workflows or `draft` for draft workflows. Active workflows cannot be edited. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionPropertiesAPI.UpdateWorkflowTransitionProperty(context.Background(), transitionId).Key(key).WorkflowName(workflowName).WorkflowTransitionProperty(workflowTransitionProperty).WorkflowMode(workflowMode).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionPropertiesAPI.UpdateWorkflowTransitionProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowTransitionProperty`: WorkflowTransitionProperty
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionPropertiesAPI.UpdateWorkflowTransitionProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**transitionId** | **int64** | The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTransitionPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **key** | **string** | The key of the property being updated, also known as the name of the property. Set this to the same value as the &#x60;key&#x60; defined in the request body. | 
 **workflowName** | **string** | The name of the workflow that the transition belongs to. | 
 **workflowTransitionProperty** | [**WorkflowTransitionProperty**](WorkflowTransitionProperty.md) |  | 
 **workflowMode** | **string** | The workflow status. Set to &#x60;live&#x60; for inactive workflows or &#x60;draft&#x60; for draft workflows. Active workflows cannot be edited. | 

### Return type

[**WorkflowTransitionProperty**](WorkflowTransitionProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

