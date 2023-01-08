# \IssuePropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**BulkDeleteIssueProperty**](IssuePropertiesAPI.md#BulkDeleteIssueProperty) | **Delete** /rest/api/3/issue/properties/{propertyKey} | Bulk delete issue property
[**BulkSetIssuePropertiesByIssue**](IssuePropertiesAPI.md#BulkSetIssuePropertiesByIssue) | **Post** /rest/api/3/issue/properties/multi | Bulk set issue properties by issue
[**BulkSetIssueProperty**](IssuePropertiesAPI.md#BulkSetIssueProperty) | **Put** /rest/api/3/issue/properties/{propertyKey} | Bulk set issue property
[**BulkSetIssuesPropertiesList**](IssuePropertiesAPI.md#BulkSetIssuesPropertiesList) | **Post** /rest/api/3/issue/properties | Bulk set issues properties by list
[**DeleteIssueProperty**](IssuePropertiesAPI.md#DeleteIssueProperty) | **Delete** /rest/api/3/issue/{issueIdOrKey}/properties/{propertyKey} | Delete issue property
[**GetIssueProperty**](IssuePropertiesAPI.md#GetIssueProperty) | **Get** /rest/api/3/issue/{issueIdOrKey}/properties/{propertyKey} | Get issue property
[**GetIssuePropertyKeys**](IssuePropertiesAPI.md#GetIssuePropertyKeys) | **Get** /rest/api/3/issue/{issueIdOrKey}/properties | Get issue property keys
[**SetIssueProperty**](IssuePropertiesAPI.md#SetIssueProperty) | **Put** /rest/api/3/issue/{issueIdOrKey}/properties/{propertyKey} | Set issue property



## BulkDeleteIssueProperty

> BulkDeleteIssueProperty(ctx, propertyKey).IssueFilterForBulkPropertyDelete(issueFilterForBulkPropertyDelete).Execute()

Bulk delete issue property



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
	propertyKey := "propertyKey_example" // string | The key of the property.
	issueFilterForBulkPropertyDelete := *openapiclient.NewIssueFilterForBulkPropertyDelete() // IssueFilterForBulkPropertyDelete | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePropertiesAPI.BulkDeleteIssueProperty(context.Background(), propertyKey).IssueFilterForBulkPropertyDelete(issueFilterForBulkPropertyDelete).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.BulkDeleteIssueProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkDeleteIssuePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueFilterForBulkPropertyDelete** | [**IssueFilterForBulkPropertyDelete**](IssueFilterForBulkPropertyDelete.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetIssuePropertiesByIssue

> BulkSetIssuePropertiesByIssue(ctx).MultiIssueEntityProperties(multiIssueEntityProperties).Execute()

Bulk set issue properties by issue



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
	multiIssueEntityProperties := *openapiclient.NewMultiIssueEntityProperties() // MultiIssueEntityProperties | Details of the issue properties to be set or updated. Note that if an issue is not found, it is ignored.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePropertiesAPI.BulkSetIssuePropertiesByIssue(context.Background()).MultiIssueEntityProperties(multiIssueEntityProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.BulkSetIssuePropertiesByIssue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetIssuePropertiesByIssueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multiIssueEntityProperties** | [**MultiIssueEntityProperties**](MultiIssueEntityProperties.md) | Details of the issue properties to be set or updated. Note that if an issue is not found, it is ignored. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetIssueProperty

> BulkSetIssueProperty(ctx, propertyKey).BulkIssuePropertyUpdateRequest(bulkIssuePropertyUpdateRequest).Execute()

Bulk set issue property



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
	propertyKey := "propertyKey_example" // string | The key of the property. The maximum length is 255 characters.
	bulkIssuePropertyUpdateRequest := *openapiclient.NewBulkIssuePropertyUpdateRequest() // BulkIssuePropertyUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePropertiesAPI.BulkSetIssueProperty(context.Background(), propertyKey).BulkIssuePropertyUpdateRequest(bulkIssuePropertyUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.BulkSetIssueProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetIssuePropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **bulkIssuePropertyUpdateRequest** | [**BulkIssuePropertyUpdateRequest**](BulkIssuePropertyUpdateRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkSetIssuesPropertiesList

> BulkSetIssuesPropertiesList(ctx).IssueEntityProperties(issueEntityProperties).Execute()

Bulk set issues properties by list



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
	issueEntityProperties := *openapiclient.NewIssueEntityProperties() // IssueEntityProperties | Issue properties to be set or updated with values.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePropertiesAPI.BulkSetIssuesPropertiesList(context.Background()).IssueEntityProperties(issueEntityProperties).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.BulkSetIssuesPropertiesList``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkSetIssuesPropertiesListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueEntityProperties** | [**IssueEntityProperties**](IssueEntityProperties.md) | Issue properties to be set or updated with values. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueProperty

> DeleteIssueProperty(ctx, issueIdOrKey, propertyKey).Execute()

Delete issue property



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
	issueIdOrKey := "issueIdOrKey_example" // string | The key or ID of the issue.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssuePropertiesAPI.DeleteIssueProperty(context.Background(), issueIdOrKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.DeleteIssueProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The key or ID of the issue. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssuePropertyRequest struct via the builder pattern


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


## GetIssueProperty

> EntityProperty GetIssueProperty(ctx, issueIdOrKey, propertyKey).Execute()

Get issue property



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
	issueIdOrKey := "issueIdOrKey_example" // string | The key or ID of the issue.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePropertiesAPI.GetIssueProperty(context.Background(), issueIdOrKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.GetIssueProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `IssuePropertiesAPI.GetIssueProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The key or ID of the issue. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuePropertyRequest struct via the builder pattern


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


## GetIssuePropertyKeys

> PropertyKeys GetIssuePropertyKeys(ctx, issueIdOrKey).Execute()

Get issue property keys



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
	issueIdOrKey := "issueIdOrKey_example" // string | The key or ID of the issue.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePropertiesAPI.GetIssuePropertyKeys(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.GetIssuePropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuePropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `IssuePropertiesAPI.GetIssuePropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The key or ID of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuePropertyKeysRequest struct via the builder pattern


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


## SetIssueProperty

> interface{} SetIssueProperty(ctx, issueIdOrKey, propertyKey).Body(body).Execute()

Set issue property



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
	propertyKey := "propertyKey_example" // string | The key of the issue property. The maximum length is 255 characters.
	body := interface{}(987) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssuePropertiesAPI.SetIssueProperty(context.Background(), issueIdOrKey, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssuePropertiesAPI.SetIssueProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetIssueProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssuePropertiesAPI.SetIssueProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**propertyKey** | **string** | The key of the issue property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetIssuePropertyRequest struct via the builder pattern


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

