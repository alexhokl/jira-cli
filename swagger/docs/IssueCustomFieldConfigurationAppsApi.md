# \IssueCustomFieldConfigurationAppsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCustomFieldConfiguration**](IssueCustomFieldConfigurationAppsAPI.md#GetCustomFieldConfiguration) | **Get** /rest/api/3/app/field/{fieldIdOrKey}/context/configuration | Get custom field configurations
[**GetCustomFieldsConfigurations**](IssueCustomFieldConfigurationAppsAPI.md#GetCustomFieldsConfigurations) | **Post** /rest/api/3/app/field/context/configuration/list | Bulk get custom field configurations
[**UpdateCustomFieldConfiguration**](IssueCustomFieldConfigurationAppsAPI.md#UpdateCustomFieldConfiguration) | **Put** /rest/api/3/app/field/{fieldIdOrKey}/context/configuration | Update custom field configurations



## GetCustomFieldConfiguration

> PageBeanContextualConfiguration GetCustomFieldConfiguration(ctx, fieldIdOrKey).Id(id).FieldContextId(fieldContextId).IssueId(issueId).ProjectKeyOrId(projectKeyOrId).IssueTypeId(issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()

Get custom field configurations



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
	fieldIdOrKey := "fieldIdOrKey_example" // string | The ID or key of the custom field, for example `customfield_10000`.
	id := []int64{int64(123)} // []int64 | The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: `id=10000&id=10001`. Can't be provided with `fieldContextId`, `issueId`, `projectKeyOrId`, or `issueTypeId`. (optional)
	fieldContextId := []int64{int64(123)} // []int64 | The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: `fieldContextId=10000&fieldContextId=10001`. Can't be provided with `id`, `issueId`, `projectKeyOrId`, or `issueTypeId`. (optional)
	issueId := int64(789) // int64 | The ID of the issue to filter results by. If the issue doesn't exist, an empty list is returned. Can't be provided with `projectKeyOrId`, or `issueTypeId`. (optional)
	projectKeyOrId := "projectKeyOrId_example" // string | The ID or key of the project to filter results by. Must be provided with `issueTypeId`. Can't be provided with `issueId`. (optional)
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type to filter results by. Must be provided with `projectKeyOrId`. Can't be provided with `issueId`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldConfigurationAppsAPI.GetCustomFieldConfiguration(context.Background(), fieldIdOrKey).Id(id).FieldContextId(fieldContextId).IssueId(issueId).ProjectKeyOrId(projectKeyOrId).IssueTypeId(issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldConfigurationAppsAPI.GetCustomFieldConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFieldConfiguration`: PageBeanContextualConfiguration
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldConfigurationAppsAPI.GetCustomFieldConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldIdOrKey** | **string** | The ID or key of the custom field, for example &#x60;customfield_10000&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **id** | **[]int64** | The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. Can&#39;t be provided with &#x60;fieldContextId&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **fieldContextId** | **[]int64** | The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: &#x60;fieldContextId&#x3D;10000&amp;fieldContextId&#x3D;10001&#x60;. Can&#39;t be provided with &#x60;id&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **issueId** | **int64** | The ID of the issue to filter results by. If the issue doesn&#39;t exist, an empty list is returned. Can&#39;t be provided with &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **projectKeyOrId** | **string** | The ID or key of the project to filter results by. Must be provided with &#x60;issueTypeId&#x60;. Can&#39;t be provided with &#x60;issueId&#x60;. | 
 **issueTypeId** | **string** | The ID of the issue type to filter results by. Must be provided with &#x60;projectKeyOrId&#x60;. Can&#39;t be provided with &#x60;issueId&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanContextualConfiguration**](PageBeanContextualConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFieldsConfigurations

> PageBeanBulkContextualConfiguration GetCustomFieldsConfigurations(ctx).ConfigurationsListParameters(configurationsListParameters).Id(id).FieldContextId(fieldContextId).IssueId(issueId).ProjectKeyOrId(projectKeyOrId).IssueTypeId(issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()

Bulk get custom field configurations



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
	configurationsListParameters := *openapiclient.NewConfigurationsListParameters([]string{"FieldIdsOrKeys_example"}) // ConfigurationsListParameters | 
	id := []int64{int64(123)} // []int64 | The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: `id=10000&id=10001`. Can't be provided with `fieldContextId`, `issueId`, `projectKeyOrId`, or `issueTypeId`. (optional)
	fieldContextId := []int64{int64(123)} // []int64 | The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: `fieldContextId=10000&fieldContextId=10001`. Can't be provided with `id`, `issueId`, `projectKeyOrId`, or `issueTypeId`. (optional)
	issueId := int64(789) // int64 | The ID of the issue to filter results by. If the issue doesn't exist, an empty list is returned. Can't be provided with `projectKeyOrId`, or `issueTypeId`. (optional)
	projectKeyOrId := "projectKeyOrId_example" // string | The ID or key of the project to filter results by. Must be provided with `issueTypeId`. Can't be provided with `issueId`. (optional)
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type to filter results by. Must be provided with `projectKeyOrId`. Can't be provided with `issueId`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldConfigurationAppsAPI.GetCustomFieldsConfigurations(context.Background()).ConfigurationsListParameters(configurationsListParameters).Id(id).FieldContextId(fieldContextId).IssueId(issueId).ProjectKeyOrId(projectKeyOrId).IssueTypeId(issueTypeId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldConfigurationAppsAPI.GetCustomFieldsConfigurations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFieldsConfigurations`: PageBeanBulkContextualConfiguration
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldConfigurationAppsAPI.GetCustomFieldsConfigurations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldsConfigurationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **configurationsListParameters** | [**ConfigurationsListParameters**](ConfigurationsListParameters.md) |  | 
 **id** | **[]int64** | The list of configuration IDs. To include multiple configurations, separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. Can&#39;t be provided with &#x60;fieldContextId&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **fieldContextId** | **[]int64** | The list of field context IDs. To include multiple field contexts, separate IDs with an ampersand: &#x60;fieldContextId&#x3D;10000&amp;fieldContextId&#x3D;10001&#x60;. Can&#39;t be provided with &#x60;id&#x60;, &#x60;issueId&#x60;, &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **issueId** | **int64** | The ID of the issue to filter results by. If the issue doesn&#39;t exist, an empty list is returned. Can&#39;t be provided with &#x60;projectKeyOrId&#x60;, or &#x60;issueTypeId&#x60;. | 
 **projectKeyOrId** | **string** | The ID or key of the project to filter results by. Must be provided with &#x60;issueTypeId&#x60;. Can&#39;t be provided with &#x60;issueId&#x60;. | 
 **issueTypeId** | **string** | The ID of the issue type to filter results by. Must be provided with &#x60;projectKeyOrId&#x60;. Can&#39;t be provided with &#x60;issueId&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanBulkContextualConfiguration**](PageBeanBulkContextualConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCustomFieldConfiguration

> interface{} UpdateCustomFieldConfiguration(ctx, fieldIdOrKey).CustomFieldConfigurations(customFieldConfigurations).Execute()

Update custom field configurations



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
	fieldIdOrKey := "fieldIdOrKey_example" // string | The ID or key of the custom field, for example `customfield_10000`.
	customFieldConfigurations := *openapiclient.NewCustomFieldConfigurations([]openapiclient.ContextualConfiguration{*openapiclient.NewContextualConfiguration("FieldContextId_example", "Id_example")}) // CustomFieldConfigurations | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldConfigurationAppsAPI.UpdateCustomFieldConfiguration(context.Background(), fieldIdOrKey).CustomFieldConfigurations(customFieldConfigurations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldConfigurationAppsAPI.UpdateCustomFieldConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomFieldConfiguration`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldConfigurationAppsAPI.UpdateCustomFieldConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldIdOrKey** | **string** | The ID or key of the custom field, for example &#x60;customfield_10000&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFieldConfigurations** | [**CustomFieldConfigurations**](CustomFieldConfigurations.md) |  | 

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

