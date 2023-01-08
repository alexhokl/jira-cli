# \ProjectKeyAndNameValidationAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetValidProjectKey**](ProjectKeyAndNameValidationAPI.md#GetValidProjectKey) | **Get** /rest/api/3/projectvalidate/validProjectKey | Get valid project key
[**GetValidProjectName**](ProjectKeyAndNameValidationAPI.md#GetValidProjectName) | **Get** /rest/api/3/projectvalidate/validProjectName | Get valid project name
[**ValidateProjectKey**](ProjectKeyAndNameValidationAPI.md#ValidateProjectKey) | **Get** /rest/api/3/projectvalidate/key | Validate project key



## GetValidProjectKey

> string GetValidProjectKey(ctx).Key(key).Execute()

Get valid project key



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
	key := "HSP" // string | The project key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectKeyAndNameValidationAPI.GetValidProjectKey(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectKeyAndNameValidationAPI.GetValidProjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidProjectKey`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectKeyAndNameValidationAPI.GetValidProjectKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The project key. | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetValidProjectName

> string GetValidProjectName(ctx).Name(name).Execute()

Get valid project name



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
	name := "name_example" // string | The project name.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectKeyAndNameValidationAPI.GetValidProjectName(context.Background()).Name(name).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectKeyAndNameValidationAPI.GetValidProjectName``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetValidProjectName`: string
	fmt.Fprintf(os.Stdout, "Response from `ProjectKeyAndNameValidationAPI.GetValidProjectName`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetValidProjectNameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | The project name. | 

### Return type

**string**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateProjectKey

> ErrorCollection ValidateProjectKey(ctx).Key(key).Execute()

Validate project key



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
	key := "HSP" // string | The project key. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectKeyAndNameValidationAPI.ValidateProjectKey(context.Background()).Key(key).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectKeyAndNameValidationAPI.ValidateProjectKey``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ValidateProjectKey`: ErrorCollection
	fmt.Fprintf(os.Stdout, "Response from `ProjectKeyAndNameValidationAPI.ValidateProjectKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateProjectKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The project key. | 

### Return type

[**ErrorCollection**](ErrorCollection.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

