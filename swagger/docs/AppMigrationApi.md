# \AppMigrationAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut**](AppMigrationAPI.md#AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut) | **Put** /rest/atlassian-connect/1/migration/field | Bulk update custom field value
[**MigrationResourceUpdateEntityPropertiesValuePut**](AppMigrationAPI.md#MigrationResourceUpdateEntityPropertiesValuePut) | **Put** /rest/atlassian-connect/1/migration/properties/{entityType} | Bulk update entity properties
[**MigrationResourceWorkflowRuleSearchPost**](AppMigrationAPI.md#MigrationResourceWorkflowRuleSearchPost) | **Post** /rest/atlassian-connect/1/migration/workflow/rule/search | Get workflow transition rule configurations



## AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut

> interface{} AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut(ctx).AtlassianTransferId(atlassianTransferId).ConnectCustomFieldValues(connectCustomFieldValues).Execute()

Bulk update custom field value



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
	atlassianTransferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The ID of the transfer.
	connectCustomFieldValues := *openapiclient.NewConnectCustomFieldValues() // ConnectCustomFieldValues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppMigrationAPI.AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut(context.Background()).AtlassianTransferId(atlassianTransferId).ConnectCustomFieldValues(connectCustomFieldValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppMigrationAPI.AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AppMigrationAPI.AppIssueFieldValueUpdateResourceUpdateIssueFieldsPut`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAppIssueFieldValueUpdateResourceUpdateIssueFieldsPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlassianTransferId** | **string** | The ID of the transfer. | 
 **connectCustomFieldValues** | [**ConnectCustomFieldValues**](ConnectCustomFieldValues.md) |  | 

### Return type

**interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationResourceUpdateEntityPropertiesValuePut

> MigrationResourceUpdateEntityPropertiesValuePut(ctx, entityType).AtlassianTransferId(atlassianTransferId).EntityPropertyDetails(entityPropertyDetails).Execute()

Bulk update entity properties



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
	atlassianTransferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app migration transfer ID.
	entityType := "entityType_example" // string | The type indicating the object that contains the entity properties.
	entityPropertyDetails := []openapiclient.EntityPropertyDetails{*openapiclient.NewEntityPropertyDetails(float32(123), "mykey", "newValue")} // []EntityPropertyDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppMigrationAPI.MigrationResourceUpdateEntityPropertiesValuePut(context.Background(), entityType).AtlassianTransferId(atlassianTransferId).EntityPropertyDetails(entityPropertyDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppMigrationAPI.MigrationResourceUpdateEntityPropertiesValuePut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**entityType** | **string** | The type indicating the object that contains the entity properties. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMigrationResourceUpdateEntityPropertiesValuePutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlassianTransferId** | **string** | The app migration transfer ID. | 

 **entityPropertyDetails** | [**[]EntityPropertyDetails**](EntityPropertyDetails.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrationResourceWorkflowRuleSearchPost

> WorkflowRulesSearchDetails MigrationResourceWorkflowRuleSearchPost(ctx).AtlassianTransferId(atlassianTransferId).WorkflowRulesSearch(workflowRulesSearch).Execute()

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
	atlassianTransferId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The app migration transfer ID.
	workflowRulesSearch := *openapiclient.NewWorkflowRulesSearch([]string{"55d44f1d-c859-42e5-9c27-2c5ec3f340b1"}, "a498d711-685d-428d-8c3e-bc03bb450ea7") // WorkflowRulesSearch | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppMigrationAPI.MigrationResourceWorkflowRuleSearchPost(context.Background()).AtlassianTransferId(atlassianTransferId).WorkflowRulesSearch(workflowRulesSearch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppMigrationAPI.MigrationResourceWorkflowRuleSearchPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrationResourceWorkflowRuleSearchPost`: WorkflowRulesSearchDetails
	fmt.Fprintf(os.Stdout, "Response from `AppMigrationAPI.MigrationResourceWorkflowRuleSearchPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrationResourceWorkflowRuleSearchPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **atlassianTransferId** | **string** | The app migration transfer ID. | 
 **workflowRulesSearch** | [**WorkflowRulesSearch**](WorkflowRulesSearch.md) |  | 

### Return type

[**WorkflowRulesSearchDetails**](WorkflowRulesSearchDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

