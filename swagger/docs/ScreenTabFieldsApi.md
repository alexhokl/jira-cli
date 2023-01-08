# \ScreenTabFieldsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddScreenTabField**](ScreenTabFieldsAPI.md#AddScreenTabField) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields | Add screen tab field
[**GetAllScreenTabFields**](ScreenTabFieldsAPI.md#GetAllScreenTabFields) | **Get** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields | Get all screen tab fields
[**MoveScreenTabField**](ScreenTabFieldsAPI.md#MoveScreenTabField) | **Post** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields/{id}/move | Move screen tab field
[**RemoveScreenTabField**](ScreenTabFieldsAPI.md#RemoveScreenTabField) | **Delete** /rest/api/3/screens/{screenId}/tabs/{tabId}/fields/{id} | Remove screen tab field



## AddScreenTabField

> ScreenableField AddScreenTabField(ctx, screenId, tabId).AddFieldBean(addFieldBean).Execute()

Add screen tab field



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	addFieldBean := *openapiclient.NewAddFieldBean("FieldId_example") // AddFieldBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabFieldsAPI.AddScreenTabField(context.Background(), screenId, tabId).AddFieldBean(addFieldBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabFieldsAPI.AddScreenTabField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddScreenTabField`: ScreenableField
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabFieldsAPI.AddScreenTabField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddScreenTabFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **addFieldBean** | [**AddFieldBean**](AddFieldBean.md) |  | 

### Return type

[**ScreenableField**](ScreenableField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllScreenTabFields

> []ScreenableField GetAllScreenTabFields(ctx, screenId, tabId).ProjectKey(projectKey).Execute()

Get all screen tab fields



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	projectKey := "projectKey_example" // string | The key of the project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabFieldsAPI.GetAllScreenTabFields(context.Background(), screenId, tabId).ProjectKey(projectKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabFieldsAPI.GetAllScreenTabFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllScreenTabFields`: []ScreenableField
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabFieldsAPI.GetAllScreenTabFields`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllScreenTabFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectKey** | **string** | The key of the project. | 

### Return type

[**[]ScreenableField**](ScreenableField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveScreenTabField

> interface{} MoveScreenTabField(ctx, screenId, tabId, id).MoveFieldBean(moveFieldBean).Execute()

Move screen tab field



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	id := "id_example" // string | The ID of the field.
	moveFieldBean := *openapiclient.NewMoveFieldBean() // MoveFieldBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ScreenTabFieldsAPI.MoveScreenTabField(context.Background(), screenId, tabId, id).MoveFieldBean(moveFieldBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabFieldsAPI.MoveScreenTabField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveScreenTabField`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ScreenTabFieldsAPI.MoveScreenTabField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 
**id** | **string** | The ID of the field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveScreenTabFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **moveFieldBean** | [**MoveFieldBean**](MoveFieldBean.md) |  | 

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


## RemoveScreenTabField

> RemoveScreenTabField(ctx, screenId, tabId, id).Execute()

Remove screen tab field



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
	screenId := int64(789) // int64 | The ID of the screen.
	tabId := int64(789) // int64 | The ID of the screen tab.
	id := "id_example" // string | The ID of the field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ScreenTabFieldsAPI.RemoveScreenTabField(context.Background(), screenId, tabId, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ScreenTabFieldsAPI.RemoveScreenTabField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**screenId** | **int64** | The ID of the screen. | 
**tabId** | **int64** | The ID of the screen tab. | 
**id** | **string** | The ID of the field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveScreenTabFieldRequest struct via the builder pattern


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

