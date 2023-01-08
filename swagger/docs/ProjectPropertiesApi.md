# \ProjectPropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteProjectProperty**](ProjectPropertiesAPI.md#DeleteProjectProperty) | **Delete** /rest/api/3/project/{projectIdOrKey}/properties/{propertyKey} | Delete project property
[**GetProjectProperty**](ProjectPropertiesAPI.md#GetProjectProperty) | **Get** /rest/api/3/project/{projectIdOrKey}/properties/{propertyKey} | Get project property
[**GetProjectPropertyKeys**](ProjectPropertiesAPI.md#GetProjectPropertyKeys) | **Get** /rest/api/3/project/{projectIdOrKey}/properties | Get project property keys
[**SetProjectProperty**](ProjectPropertiesAPI.md#SetProjectProperty) | **Put** /rest/api/3/project/{projectIdOrKey}/properties/{propertyKey} | Set project property



## DeleteProjectProperty

> DeleteProjectProperty(ctx, projectIdOrKey, propertyKey).Execute()

Delete project property



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	propertyKey := "propertyKey_example" // string | The project property key. Use [Get project property keys](#api-rest-api-3-project-projectIdOrKey-properties-get) to get a list of all project property keys.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectPropertiesAPI.DeleteProjectProperty(context.Background(), projectIdOrKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPropertiesAPI.DeleteProjectProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**propertyKey** | **string** | The project property key. Use [Get project property keys](#api-rest-api-3-project-projectIdOrKey-properties-get) to get a list of all project property keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProjectPropertyRequest struct via the builder pattern


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


## GetProjectProperty

> EntityProperty GetProjectProperty(ctx, projectIdOrKey, propertyKey).Execute()

Get project property



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	propertyKey := "propertyKey_example" // string | The project property key. Use [Get project property keys](#api-rest-api-3-project-projectIdOrKey-properties-get) to get a list of all project property keys.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPropertiesAPI.GetProjectProperty(context.Background(), projectIdOrKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPropertiesAPI.GetProjectProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `ProjectPropertiesAPI.GetProjectProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**propertyKey** | **string** | The project property key. Use [Get project property keys](#api-rest-api-3-project-projectIdOrKey-properties-get) to get a list of all project property keys. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectPropertyRequest struct via the builder pattern


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


## GetProjectPropertyKeys

> PropertyKeys GetProjectPropertyKeys(ctx, projectIdOrKey).Execute()

Get project property keys



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPropertiesAPI.GetProjectPropertyKeys(context.Background(), projectIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPropertiesAPI.GetProjectPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectPropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `ProjectPropertiesAPI.GetProjectPropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectPropertyKeysRequest struct via the builder pattern


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


## SetProjectProperty

> interface{} SetProjectProperty(ctx, projectIdOrKey, propertyKey).Body(body).Execute()

Set project property



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	propertyKey := "propertyKey_example" // string | The key of the project property. The maximum length is 255 characters.
	body := interface{}({"number":5,"string":"string-value"}) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectPropertiesAPI.SetProjectProperty(context.Background(), projectIdOrKey, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectPropertiesAPI.SetProjectProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProjectProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectPropertiesAPI.SetProjectProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 
**propertyKey** | **string** | The key of the project property. The maximum length is 255 characters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetProjectPropertyRequest struct via the builder pattern


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

