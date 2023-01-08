# \IssueTypePropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIssueTypeProperty**](IssueTypePropertiesAPI.md#DeleteIssueTypeProperty) | **Delete** /rest/api/3/issuetype/{issueTypeId}/properties/{propertyKey} | Delete issue type property
[**GetIssueTypeProperty**](IssueTypePropertiesAPI.md#GetIssueTypeProperty) | **Get** /rest/api/3/issuetype/{issueTypeId}/properties/{propertyKey} | Get issue type property
[**GetIssueTypePropertyKeys**](IssueTypePropertiesAPI.md#GetIssueTypePropertyKeys) | **Get** /rest/api/3/issuetype/{issueTypeId}/properties | Get issue type property keys
[**SetIssueTypeProperty**](IssueTypePropertiesAPI.md#SetIssueTypeProperty) | **Put** /rest/api/3/issuetype/{issueTypeId}/properties/{propertyKey} | Set issue type property



## DeleteIssueTypeProperty

> DeleteIssueTypeProperty(ctx, issueTypeId, propertyKey).Execute()

Delete issue type property



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
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type.
	propertyKey := "propertyKey_example" // string | The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get a list of all issue type property keys.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueTypePropertiesAPI.DeleteIssueTypeProperty(context.Background(), issueTypeId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypePropertiesAPI.DeleteIssueTypeProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeId** | **string** | The ID of the issue type. | 
**propertyKey** | **string** | The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get a list of all issue type property keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueTypePropertyRequest struct via the builder pattern


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


## GetIssueTypeProperty

> EntityProperty GetIssueTypeProperty(ctx, issueTypeId, propertyKey).Execute()

Get issue type property



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
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type.
	propertyKey := "propertyKey_example" // string | The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get a list of all issue type property keys.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypePropertiesAPI.GetIssueTypeProperty(context.Background(), issueTypeId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypePropertiesAPI.GetIssueTypeProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `IssueTypePropertiesAPI.GetIssueTypeProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeId** | **string** | The ID of the issue type. | 
**propertyKey** | **string** | The key of the property. Use [Get issue type property keys](#api-rest-api-3-issuetype-issueTypeId-properties-get) to get a list of all issue type property keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypePropertyRequest struct via the builder pattern


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


## GetIssueTypePropertyKeys

> PropertyKeys GetIssueTypePropertyKeys(ctx, issueTypeId).Execute()

Get issue type property keys



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
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypePropertiesAPI.GetIssueTypePropertyKeys(context.Background(), issueTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypePropertiesAPI.GetIssueTypePropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypePropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `IssueTypePropertiesAPI.GetIssueTypePropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeId** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypePropertyKeysRequest struct via the builder pattern


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


## SetIssueTypeProperty

> interface{} SetIssueTypeProperty(ctx, issueTypeId, propertyKey).Body(body).Execute()

Set issue type property



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
	issueTypeId := "issueTypeId_example" // string | The ID of the issue type.
	propertyKey := "propertyKey_example" // string | The key of the issue type property. The maximum length is 255 characters.
	body := interface{}({"number":5,"string":"string-value"}) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypePropertiesAPI.SetIssueTypeProperty(context.Background(), issueTypeId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypePropertiesAPI.SetIssueTypeProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetIssueTypeProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueTypePropertiesAPI.SetIssueTypeProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueTypeId** | **string** | The ID of the issue type. | 
**propertyKey** | **string** | The key of the issue type property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIssueTypePropertyRequest struct via the builder pattern


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

