# \DevOpsComponentsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteComponentById**](DevOpsComponentsAPI.md#DeleteComponentById) | **Delete** /rest/devopscomponents/1.0/devopscomponents/{componentId} | Delete a Component by ID
[**DeleteComponentsByProperty**](DevOpsComponentsAPI.md#DeleteComponentsByProperty) | **Delete** /rest/devopscomponents/1.0/bulkByProperties | Delete DevOps Components by Property
[**GetComponentById**](DevOpsComponentsAPI.md#GetComponentById) | **Get** /rest/devopscomponents/1.0/devopscomponents/{componentId} | Get a Component by ID
[**SubmitComponents**](DevOpsComponentsAPI.md#SubmitComponents) | **Post** /rest/devopscomponents/1.0/bulk | Submit DevOps Components



## DeleteComponentById

> DeleteComponentById(ctx, componentId).Authorization(authorization).Execute()

Delete a Component by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	componentId := "componentId_example" // string | The ID of the Component to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsComponentsAPI.DeleteComponentById(context.Background(), componentId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsComponentsAPI.DeleteComponentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | The ID of the Component to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComponentsByProperty

> DeleteComponentsByProperty(ctx).Authorization(authorization).Execute()

Delete DevOps Components by Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsComponentsAPI.DeleteComponentsByProperty(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsComponentsAPI.DeleteComponentsByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteComponentsByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComponentById

> GetComponentById200Response GetComponentById(ctx, componentId).Authorization(authorization).Execute()

Get a Component by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	componentId := "componentId_example" // string | The ID of the Component to fetch. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsComponentsAPI.GetComponentById(context.Background(), componentId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsComponentsAPI.GetComponentById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComponentById`: GetComponentById200Response
	fmt.Fprintf(os.Stdout, "Response from `DevOpsComponentsAPI.GetComponentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**componentId** | **string** | The ID of the Component to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetComponentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Operations Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 


### Return type

[**GetComponentById200Response**](GetComponentById200Response.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitComponents

> SubmitDevopsComponentsResponse SubmitComponents(ctx).Authorization(authorization).SubmitDevopsComponentsRequest(submitDevopsComponentsRequest).Execute()

Submit DevOps Components



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the DevOps Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	submitDevopsComponentsRequest := *openapiclient.NewSubmitDevopsComponentsRequest([]openapiclient.Component{*openapiclient.NewComponent("1.0", "111-222-333", int64(1523494301448), "Galactus", "Description_example", "https://example.com/project/MS/galactus/summary", "https://example.com/project/MS/galactus/logo", "Tier 1", "Service", time.Now())}) // SubmitDevopsComponentsRequest | DevOps Component data to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsComponentsAPI.SubmitComponents(context.Background()).Authorization(authorization).SubmitDevopsComponentsRequest(submitDevopsComponentsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsComponentsAPI.SubmitComponents``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitComponents`: SubmitDevopsComponentsResponse
	fmt.Fprintf(os.Stdout, "Response from `DevOpsComponentsAPI.SubmitComponents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitComponentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the DevOps Information module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 
 **submitDevopsComponentsRequest** | [**SubmitDevopsComponentsRequest**](SubmitDevopsComponentsRequest.md) | DevOps Component data to submit.  | 

### Return type

[**SubmitDevopsComponentsResponse**](SubmitDevopsComponentsResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

