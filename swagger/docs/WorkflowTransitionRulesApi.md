# \WorkflowTransitionRulesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesAPI.md#DeleteWorkflowTransitionRuleConfigurations) | **Put** /rest/api/3/workflow/rule/config/delete | Delete workflow transition rule configurations
[**GetWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesAPI.md#GetWorkflowTransitionRuleConfigurations) | **Get** /rest/api/3/workflow/rule/config | Get workflow transition rule configurations
[**UpdateWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesAPI.md#UpdateWorkflowTransitionRuleConfigurations) | **Put** /rest/api/3/workflow/rule/config | Update workflow transition rule configurations



## DeleteWorkflowTransitionRuleConfigurations

> WorkflowTransitionRulesUpdateErrors DeleteWorkflowTransitionRuleConfigurations(ctx).WorkflowsWithTransitionRulesDetails(workflowsWithTransitionRulesDetails).Execute()

Delete workflow transition rule configurations



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
	workflowsWithTransitionRulesDetails := *openapiclient.NewWorkflowsWithTransitionRulesDetails([]openapiclient.WorkflowTransitionRulesDetails{*openapiclient.NewWorkflowTransitionRulesDetails(*openapiclient.NewWorkflowId(false, "Name_example"), []string{"WorkflowRuleIds_example"})}) // WorkflowsWithTransitionRulesDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionRulesAPI.DeleteWorkflowTransitionRuleConfigurations(context.Background()).WorkflowsWithTransitionRulesDetails(workflowsWithTransitionRulesDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionRulesAPI.DeleteWorkflowTransitionRuleConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWorkflowTransitionRuleConfigurations`: WorkflowTransitionRulesUpdateErrors
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionRulesAPI.DeleteWorkflowTransitionRuleConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowTransitionRuleConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowsWithTransitionRulesDetails** | [**WorkflowsWithTransitionRulesDetails**](WorkflowsWithTransitionRulesDetails.md) |  | 

### Return type

[**WorkflowTransitionRulesUpdateErrors**](WorkflowTransitionRulesUpdateErrors.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkflowTransitionRuleConfigurations

> PageBeanWorkflowTransitionRules GetWorkflowTransitionRuleConfigurations(ctx).Types(types).StartAt(startAt).MaxResults(maxResults).Keys(keys).WorkflowNames(workflowNames).WithTags(withTags).Draft(draft).Expand(expand).Execute()

Get workflow transition rule configurations



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
	types := []string{"Types_example"} // []string | The types of the transition rules to return.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 10)
	keys := []string{"Inner_example"} // []string | The transition rule class keys, as defined in the Connect or the Forge app descriptor, of the transition rules to return. (optional)
	workflowNames := []string{"Inner_example"} // []string | The list of workflow names to filter by. (optional)
	withTags := []string{"Inner_example"} // []string | The list of `tags` to filter by. (optional)
	draft := true // bool | Whether draft or published workflows are returned. If not provided, both workflow types are returned. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts `transition`, which, for each rule, returns information about the transition the rule is assigned to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionRulesAPI.GetWorkflowTransitionRuleConfigurations(context.Background()).Types(types).StartAt(startAt).MaxResults(maxResults).Keys(keys).WorkflowNames(workflowNames).WithTags(withTags).Draft(draft).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionRulesAPI.GetWorkflowTransitionRuleConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorkflowTransitionRuleConfigurations`: PageBeanWorkflowTransitionRules
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionRulesAPI.GetWorkflowTransitionRuleConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowTransitionRuleConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **types** | **[]string** | The types of the transition rules to return. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 10]
 **keys** | **[]string** | The transition rule class keys, as defined in the Connect or the Forge app descriptor, of the transition rules to return. | 
 **workflowNames** | **[]string** | The list of workflow names to filter by. | 
 **withTags** | **[]string** | The list of &#x60;tags&#x60; to filter by. | 
 **draft** | **bool** | Whether draft or published workflows are returned. If not provided, both workflow types are returned. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;transition&#x60;, which, for each rule, returns information about the transition the rule is assigned to. | 

### Return type

[**PageBeanWorkflowTransitionRules**](PageBeanWorkflowTransitionRules.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateWorkflowTransitionRuleConfigurations

> WorkflowTransitionRulesUpdateErrors UpdateWorkflowTransitionRuleConfigurations(ctx).WorkflowTransitionRulesUpdate(workflowTransitionRulesUpdate).Execute()

Update workflow transition rule configurations



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
	workflowTransitionRulesUpdate := *openapiclient.NewWorkflowTransitionRulesUpdate([]openapiclient.WorkflowTransitionRules{*openapiclient.NewWorkflowTransitionRules(*openapiclient.NewWorkflowId(false, "Name_example"))}) // WorkflowTransitionRulesUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowTransitionRulesAPI.UpdateWorkflowTransitionRuleConfigurations(context.Background()).WorkflowTransitionRulesUpdate(workflowTransitionRulesUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowTransitionRulesAPI.UpdateWorkflowTransitionRuleConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWorkflowTransitionRuleConfigurations`: WorkflowTransitionRulesUpdateErrors
	fmt.Fprintf(os.Stdout, "Response from `WorkflowTransitionRulesAPI.UpdateWorkflowTransitionRuleConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowTransitionRuleConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowTransitionRulesUpdate** | [**WorkflowTransitionRulesUpdate**](WorkflowTransitionRulesUpdate.md) |  | 

### Return type

[**WorkflowTransitionRulesUpdateErrors**](WorkflowTransitionRulesUpdateErrors.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

