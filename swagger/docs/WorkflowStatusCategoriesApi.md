# \WorkflowStatusCategoriesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetStatusCategories**](WorkflowStatusCategoriesAPI.md#GetStatusCategories) | **Get** /rest/api/3/statuscategory | Get all status categories
[**GetStatusCategory**](WorkflowStatusCategoriesAPI.md#GetStatusCategory) | **Get** /rest/api/3/statuscategory/{idOrKey} | Get status category



## GetStatusCategories

> []StatusCategory GetStatusCategories(ctx).Execute()

Get all status categories



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
	resp, r, err := apiClient.WorkflowStatusCategoriesAPI.GetStatusCategories(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowStatusCategoriesAPI.GetStatusCategories``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusCategories`: []StatusCategory
	fmt.Fprintf(os.Stdout, "Response from `WorkflowStatusCategoriesAPI.GetStatusCategories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCategoriesRequest struct via the builder pattern


### Return type

[**[]StatusCategory**](StatusCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatusCategory

> StatusCategory GetStatusCategory(ctx, idOrKey).Execute()

Get status category



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
	idOrKey := "idOrKey_example" // string | The ID or key of the status category.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WorkflowStatusCategoriesAPI.GetStatusCategory(context.Background(), idOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WorkflowStatusCategoriesAPI.GetStatusCategory``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatusCategory`: StatusCategory
	fmt.Fprintf(os.Stdout, "Response from `WorkflowStatusCategoriesAPI.GetStatusCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**idOrKey** | **string** | The ID or key of the status category. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StatusCategory**](StatusCategory.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

