# \IssueWorklogPropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorklogProperty**](IssueWorklogPropertiesAPI.md#DeleteWorklogProperty) | **Delete** /rest/api/3/issue/{issueIdOrKey}/worklog/{worklogId}/properties/{propertyKey} | Delete worklog property
[**GetWorklogProperty**](IssueWorklogPropertiesAPI.md#GetWorklogProperty) | **Get** /rest/api/3/issue/{issueIdOrKey}/worklog/{worklogId}/properties/{propertyKey} | Get worklog property
[**GetWorklogPropertyKeys**](IssueWorklogPropertiesAPI.md#GetWorklogPropertyKeys) | **Get** /rest/api/3/issue/{issueIdOrKey}/worklog/{worklogId}/properties | Get worklog property keys
[**SetWorklogProperty**](IssueWorklogPropertiesAPI.md#SetWorklogProperty) | **Put** /rest/api/3/issue/{issueIdOrKey}/worklog/{worklogId}/properties/{propertyKey} | Set worklog property



## DeleteWorklogProperty

> DeleteWorklogProperty(ctx, issueIdOrKey, worklogId, propertyKey).Execute()

Delete worklog property



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	worklogId := "worklogId_example" // string | The ID of the worklog.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueWorklogPropertiesAPI.DeleteWorklogProperty(context.Background(), issueIdOrKey, worklogId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogPropertiesAPI.DeleteWorklogProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**worklogId** | **string** | The ID of the worklog. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorklogPropertyRequest struct via the builder pattern


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


## GetWorklogProperty

> EntityProperty GetWorklogProperty(ctx, issueIdOrKey, worklogId, propertyKey).Execute()

Get worklog property



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	worklogId := "worklogId_example" // string | The ID of the worklog.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogPropertiesAPI.GetWorklogProperty(context.Background(), issueIdOrKey, worklogId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogPropertiesAPI.GetWorklogProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorklogProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogPropertiesAPI.GetWorklogProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**worklogId** | **string** | The ID of the worklog. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorklogPropertyRequest struct via the builder pattern


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


## GetWorklogPropertyKeys

> PropertyKeys GetWorklogPropertyKeys(ctx, issueIdOrKey, worklogId).Execute()

Get worklog property keys



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	worklogId := "worklogId_example" // string | The ID of the worklog.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogPropertiesAPI.GetWorklogPropertyKeys(context.Background(), issueIdOrKey, worklogId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogPropertiesAPI.GetWorklogPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorklogPropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogPropertiesAPI.GetWorklogPropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**worklogId** | **string** | The ID of the worklog. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorklogPropertyKeysRequest struct via the builder pattern


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


## SetWorklogProperty

> interface{} SetWorklogProperty(ctx, issueIdOrKey, worklogId, propertyKey).Body(body).Execute()

Set worklog property



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue.
	worklogId := "worklogId_example" // string | The ID of the worklog.
	propertyKey := "propertyKey_example" // string | The key of the issue property. The maximum length is 255 characters.
	body := interface{}(987) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueWorklogPropertiesAPI.SetWorklogProperty(context.Background(), issueIdOrKey, worklogId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueWorklogPropertiesAPI.SetWorklogProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetWorklogProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueWorklogPropertiesAPI.SetWorklogProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**worklogId** | **string** | The ID of the worklog. | 
**propertyKey** | **string** | The key of the issue property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetWorklogPropertyRequest struct via the builder pattern


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

