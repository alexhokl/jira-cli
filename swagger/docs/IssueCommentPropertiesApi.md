# \IssueCommentPropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCommentProperty**](IssueCommentPropertiesAPI.md#DeleteCommentProperty) | **Delete** /rest/api/3/comment/{commentId}/properties/{propertyKey} | Delete comment property
[**GetCommentProperty**](IssueCommentPropertiesAPI.md#GetCommentProperty) | **Get** /rest/api/3/comment/{commentId}/properties/{propertyKey} | Get comment property
[**GetCommentPropertyKeys**](IssueCommentPropertiesAPI.md#GetCommentPropertyKeys) | **Get** /rest/api/3/comment/{commentId}/properties | Get comment property keys
[**SetCommentProperty**](IssueCommentPropertiesAPI.md#SetCommentProperty) | **Put** /rest/api/3/comment/{commentId}/properties/{propertyKey} | Set comment property



## DeleteCommentProperty

> DeleteCommentProperty(ctx, commentId, propertyKey).Execute()

Delete comment property



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
	commentId := "commentId_example" // string | The ID of the comment.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueCommentPropertiesAPI.DeleteCommentProperty(context.Background(), commentId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentPropertiesAPI.DeleteCommentProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The ID of the comment. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentPropertyRequest struct via the builder pattern


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


## GetCommentProperty

> EntityProperty GetCommentProperty(ctx, commentId, propertyKey).Execute()

Get comment property



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
	commentId := "commentId_example" // string | The ID of the comment.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentPropertiesAPI.GetCommentProperty(context.Background(), commentId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentPropertiesAPI.GetCommentProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentPropertiesAPI.GetCommentProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The ID of the comment. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityProperty**](EntityProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentPropertyKeys

> PropertyKeys GetCommentPropertyKeys(ctx, commentId).Execute()

Get comment property keys



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
	commentId := "commentId_example" // string | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentPropertiesAPI.GetCommentPropertyKeys(context.Background(), commentId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentPropertiesAPI.GetCommentPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentPropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentPropertiesAPI.GetCommentPropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentPropertyKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PropertyKeys**](PropertyKeys.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCommentProperty

> interface{} SetCommentProperty(ctx, commentId, propertyKey).Body(body).Execute()

Set comment property



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
	commentId := "commentId_example" // string | The ID of the comment.
	propertyKey := "propertyKey_example" // string | The key of the property. The maximum length is 255 characters.
	body := interface{}(987) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentPropertiesAPI.SetCommentProperty(context.Background(), commentId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentPropertiesAPI.SetCommentProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetCommentProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentPropertiesAPI.SetCommentProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commentId** | **string** | The ID of the comment. | 
**propertyKey** | **string** | The key of the property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCommentPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **interface{}** | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes. | 

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

