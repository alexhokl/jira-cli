# \IssueLinkTypesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssueLinkType**](IssueLinkTypesAPI.md#CreateIssueLinkType) | **Post** /rest/api/3/issueLinkType | Create issue link type
[**DeleteIssueLinkType**](IssueLinkTypesAPI.md#DeleteIssueLinkType) | **Delete** /rest/api/3/issueLinkType/{issueLinkTypeId} | Delete issue link type
[**GetIssueLinkType**](IssueLinkTypesAPI.md#GetIssueLinkType) | **Get** /rest/api/3/issueLinkType/{issueLinkTypeId} | Get issue link type
[**GetIssueLinkTypes**](IssueLinkTypesAPI.md#GetIssueLinkTypes) | **Get** /rest/api/3/issueLinkType | Get issue link types
[**UpdateIssueLinkType**](IssueLinkTypesAPI.md#UpdateIssueLinkType) | **Put** /rest/api/3/issueLinkType/{issueLinkTypeId} | Update issue link type



## CreateIssueLinkType

> IssueLinkType CreateIssueLinkType(ctx).IssueLinkType(issueLinkType).Execute()

Create issue link type



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
	issueLinkType := *openapiclient.NewIssueLinkType() // IssueLinkType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueLinkTypesAPI.CreateIssueLinkType(context.Background()).IssueLinkType(issueLinkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinkTypesAPI.CreateIssueLinkType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueLinkType`: IssueLinkType
	fmt.Fprintf(os.Stdout, "Response from `IssueLinkTypesAPI.CreateIssueLinkType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueLinkType** | [**IssueLinkType**](IssueLinkType.md) |  | 

### Return type

[**IssueLinkType**](IssueLinkType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueLinkType

> DeleteIssueLinkType(ctx, issueLinkTypeId).Execute()

Delete issue link type



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
	issueLinkTypeId := "issueLinkTypeId_example" // string | The ID of the issue link type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueLinkTypesAPI.DeleteIssueLinkType(context.Background(), issueLinkTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinkTypesAPI.DeleteIssueLinkType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueLinkTypeId** | **string** | The ID of the issue link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueLinkTypeRequest struct via the builder pattern


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


## GetIssueLinkType

> IssueLinkType GetIssueLinkType(ctx, issueLinkTypeId).Execute()

Get issue link type



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
	issueLinkTypeId := "issueLinkTypeId_example" // string | The ID of the issue link type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueLinkTypesAPI.GetIssueLinkType(context.Background(), issueLinkTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinkTypesAPI.GetIssueLinkType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueLinkType`: IssueLinkType
	fmt.Fprintf(os.Stdout, "Response from `IssueLinkTypesAPI.GetIssueLinkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueLinkTypeId** | **string** | The ID of the issue link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueLinkType**](IssueLinkType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueLinkTypes

> IssueLinkTypes GetIssueLinkTypes(ctx).Execute()

Get issue link types



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
	resp, r, err := apiClient.IssueLinkTypesAPI.GetIssueLinkTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinkTypesAPI.GetIssueLinkTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueLinkTypes`: IssueLinkTypes
	fmt.Fprintf(os.Stdout, "Response from `IssueLinkTypesAPI.GetIssueLinkTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueLinkTypesRequest struct via the builder pattern


### Return type

[**IssueLinkTypes**](IssueLinkTypes.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssueLinkType

> IssueLinkType UpdateIssueLinkType(ctx, issueLinkTypeId).IssueLinkType(issueLinkType).Execute()

Update issue link type



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
	issueLinkTypeId := "issueLinkTypeId_example" // string | The ID of the issue link type.
	issueLinkType := *openapiclient.NewIssueLinkType() // IssueLinkType | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueLinkTypesAPI.UpdateIssueLinkType(context.Background(), issueLinkTypeId).IssueLinkType(issueLinkType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinkTypesAPI.UpdateIssueLinkType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueLinkType`: IssueLinkType
	fmt.Fprintf(os.Stdout, "Response from `IssueLinkTypesAPI.UpdateIssueLinkType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueLinkTypeId** | **string** | The ID of the issue link type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueLinkTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueLinkType** | [**IssueLinkType**](IssueLinkType.md) |  | 

### Return type

[**IssueLinkType**](IssueLinkType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

