# \IssueCustomFieldOptionsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomFieldOption**](IssueCustomFieldOptionsAPI.md#CreateCustomFieldOption) | **Post** /rest/api/3/field/{fieldId}/context/{contextId}/option | Create custom field options (context)
[**DeleteCustomFieldOption**](IssueCustomFieldOptionsAPI.md#DeleteCustomFieldOption) | **Delete** /rest/api/3/field/{fieldId}/context/{contextId}/option/{optionId} | Delete custom field options (context)
[**GetCustomFieldOption**](IssueCustomFieldOptionsAPI.md#GetCustomFieldOption) | **Get** /rest/api/3/customFieldOption/{id} | Get custom field option
[**GetOptionsForContext**](IssueCustomFieldOptionsAPI.md#GetOptionsForContext) | **Get** /rest/api/3/field/{fieldId}/context/{contextId}/option | Get custom field options (context)
[**ReorderCustomFieldOptions**](IssueCustomFieldOptionsAPI.md#ReorderCustomFieldOptions) | **Put** /rest/api/3/field/{fieldId}/context/{contextId}/option/move | Reorder custom field options (context)
[**ReplaceCustomFieldOption**](IssueCustomFieldOptionsAPI.md#ReplaceCustomFieldOption) | **Delete** /rest/api/3/field/{fieldId}/context/{contextId}/option/{optionId}/issue | Replace custom field options
[**UpdateCustomFieldOption**](IssueCustomFieldOptionsAPI.md#UpdateCustomFieldOption) | **Put** /rest/api/3/field/{fieldId}/context/{contextId}/option | Update custom field options (context)



## CreateCustomFieldOption

> CustomFieldCreatedContextOptionsList CreateCustomFieldOption(ctx, fieldId, contextId).BulkCustomFieldOptionCreateRequest(bulkCustomFieldOptionCreateRequest).Execute()

Create custom field options (context)



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	contextId := int64(789) // int64 | The ID of the context.
	bulkCustomFieldOptionCreateRequest := *openapiclient.NewBulkCustomFieldOptionCreateRequest() // BulkCustomFieldOptionCreateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAPI.CreateCustomFieldOption(context.Background(), fieldId, contextId).BulkCustomFieldOptionCreateRequest(bulkCustomFieldOptionCreateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.CreateCustomFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomFieldOption`: CustomFieldCreatedContextOptionsList
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAPI.CreateCustomFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkCustomFieldOptionCreateRequest** | [**BulkCustomFieldOptionCreateRequest**](BulkCustomFieldOptionCreateRequest.md) |  | 

### Return type

[**CustomFieldCreatedContextOptionsList**](CustomFieldCreatedContextOptionsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomFieldOption

> DeleteCustomFieldOption(ctx, fieldId, contextId, optionId).Execute()

Delete custom field options (context)



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	contextId := int64(789) // int64 | The ID of the context from which an option should be deleted.
	optionId := int64(789) // int64 | The ID of the option to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueCustomFieldOptionsAPI.DeleteCustomFieldOption(context.Background(), fieldId, contextId, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.DeleteCustomFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context from which an option should be deleted. | 
**optionId** | **int64** | The ID of the option to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldOptionRequest struct via the builder pattern


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


## GetCustomFieldOption

> CustomFieldOption GetCustomFieldOption(ctx, id).Execute()

Get custom field option



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
	id := "id_example" // string | The ID of the custom field option.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAPI.GetCustomFieldOption(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.GetCustomFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFieldOption`: CustomFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAPI.GetCustomFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the custom field option. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CustomFieldOption**](CustomFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOptionsForContext

> PageBeanCustomFieldContextOption GetOptionsForContext(ctx, fieldId, contextId).OptionId(optionId).OnlyOptions(onlyOptions).StartAt(startAt).MaxResults(maxResults).Execute()

Get custom field options (context)



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	contextId := int64(789) // int64 | The ID of the context.
	optionId := int64(789) // int64 | The ID of the option. (optional)
	onlyOptions := true // bool | Whether only options are returned. (optional) (default to false)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAPI.GetOptionsForContext(context.Background(), fieldId, contextId).OptionId(optionId).OnlyOptions(onlyOptions).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.GetOptionsForContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOptionsForContext`: PageBeanCustomFieldContextOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAPI.GetOptionsForContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOptionsForContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **optionId** | **int64** | The ID of the option. | 
 **onlyOptions** | **bool** | Whether only options are returned. | [default to false]
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanCustomFieldContextOption**](PageBeanCustomFieldContextOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReorderCustomFieldOptions

> interface{} ReorderCustomFieldOptions(ctx, fieldId, contextId).OrderOfCustomFieldOptions(orderOfCustomFieldOptions).Execute()

Reorder custom field options (context)



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	contextId := int64(789) // int64 | The ID of the context.
	orderOfCustomFieldOptions := *openapiclient.NewOrderOfCustomFieldOptions([]string{"CustomFieldOptionIds_example"}) // OrderOfCustomFieldOptions | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAPI.ReorderCustomFieldOptions(context.Background(), fieldId, contextId).OrderOfCustomFieldOptions(orderOfCustomFieldOptions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.ReorderCustomFieldOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReorderCustomFieldOptions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAPI.ReorderCustomFieldOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReorderCustomFieldOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **orderOfCustomFieldOptions** | [**OrderOfCustomFieldOptions**](OrderOfCustomFieldOptions.md) |  | 

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


## ReplaceCustomFieldOption

> ReplaceCustomFieldOption(ctx, fieldId, optionId, contextId).ReplaceWith(replaceWith).Jql(jql).Execute()

Replace custom field options



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	optionId := int64(789) // int64 | The ID of the option to be deselected.
	contextId := int64(789) // int64 | The ID of the context.
	replaceWith := int64(789) // int64 | The ID of the option that will replace the currently selected option. (optional)
	jql := "jql_example" // string | A JQL query that specifies the issues to be updated. For example, *project=10000*. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueCustomFieldOptionsAPI.ReplaceCustomFieldOption(context.Background(), fieldId, optionId, contextId).ReplaceWith(replaceWith).Jql(jql).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.ReplaceCustomFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**optionId** | **int64** | The ID of the option to be deselected. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceCustomFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **replaceWith** | **int64** | The ID of the option that will replace the currently selected option. | 
 **jql** | **string** | A JQL query that specifies the issues to be updated. For example, *project&#x3D;10000*. | 

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


## UpdateCustomFieldOption

> CustomFieldUpdatedContextOptionsList UpdateCustomFieldOption(ctx, fieldId, contextId).BulkCustomFieldOptionUpdateRequest(bulkCustomFieldOptionUpdateRequest).Execute()

Update custom field options (context)



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	contextId := int64(789) // int64 | The ID of the context.
	bulkCustomFieldOptionUpdateRequest := *openapiclient.NewBulkCustomFieldOptionUpdateRequest() // BulkCustomFieldOptionUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAPI.UpdateCustomFieldOption(context.Background(), fieldId, contextId).BulkCustomFieldOptionUpdateRequest(bulkCustomFieldOptionUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAPI.UpdateCustomFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomFieldOption`: CustomFieldUpdatedContextOptionsList
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAPI.UpdateCustomFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **bulkCustomFieldOptionUpdateRequest** | [**BulkCustomFieldOptionUpdateRequest**](BulkCustomFieldOptionUpdateRequest.md) |  | 

### Return type

[**CustomFieldUpdatedContextOptionsList**](CustomFieldUpdatedContextOptionsList.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

