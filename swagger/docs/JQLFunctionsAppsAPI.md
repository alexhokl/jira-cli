# \JQLFunctionsAppsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrecomputations**](JQLFunctionsAppsAPI.md#GetPrecomputations) | **Get** /rest/api/3/jql/function/computation | Get precomputations (apps)
[**GetPrecomputationsByID**](JQLFunctionsAppsAPI.md#GetPrecomputationsByID) | **Post** /rest/api/3/jql/function/computation/search | Get precomputations by ID (apps)
[**UpdatePrecomputations**](JQLFunctionsAppsAPI.md#UpdatePrecomputations) | **Post** /rest/api/3/jql/function/computation | Update precomputations (apps)



## GetPrecomputations

> PageBean2JqlFunctionPrecomputationBean GetPrecomputations(ctx).FunctionKey(functionKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Execute()

Get precomputations (apps)



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
	functionKey := []string{"Inner_example"} // []string | The function key in format:   *  Forge: `ari:cloud:ecosystem::extension/[App ID]/[Environment ID]/static/[Function key from manifest]`  *  Connect: `[App key]__[Module key]` (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `functionKey` Sorts by the functionKey.  *  `used` Sorts by the used timestamp.  *  `created` Sorts by the created timestamp.  *  `updated` Sorts by the updated timestamp. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLFunctionsAppsAPI.GetPrecomputations(context.Background()).FunctionKey(functionKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLFunctionsAppsAPI.GetPrecomputations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrecomputations`: PageBean2JqlFunctionPrecomputationBean
	fmt.Fprintf(os.Stdout, "Response from `JQLFunctionsAppsAPI.GetPrecomputations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrecomputationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionKey** | **[]string** | The function key in format:   *  Forge: &#x60;ari:cloud:ecosystem::extension/[App ID]/[Environment ID]/static/[Function key from manifest]&#x60;  *  Connect: &#x60;[App key]__[Module key]&#x60; | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;functionKey&#x60; Sorts by the functionKey.  *  &#x60;used&#x60; Sorts by the used timestamp.  *  &#x60;created&#x60; Sorts by the created timestamp.  *  &#x60;updated&#x60; Sorts by the updated timestamp. | 

### Return type

[**PageBean2JqlFunctionPrecomputationBean**](PageBean2JqlFunctionPrecomputationBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPrecomputationsByID

> JqlFunctionPrecomputationGetByIdResponse GetPrecomputationsByID(ctx).JqlFunctionPrecomputationGetByIdRequest(jqlFunctionPrecomputationGetByIdRequest).OrderBy(orderBy).Execute()

Get precomputations by ID (apps)



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
	jqlFunctionPrecomputationGetByIdRequest := *openapiclient.NewJqlFunctionPrecomputationGetByIdRequest() // JqlFunctionPrecomputationGetByIdRequest | 
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `functionKey` Sorts by the functionKey.  *  `used` Sorts by the used timestamp.  *  `created` Sorts by the created timestamp.  *  `updated` Sorts by the updated timestamp. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLFunctionsAppsAPI.GetPrecomputationsByID(context.Background()).JqlFunctionPrecomputationGetByIdRequest(jqlFunctionPrecomputationGetByIdRequest).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLFunctionsAppsAPI.GetPrecomputationsByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPrecomputationsByID`: JqlFunctionPrecomputationGetByIdResponse
	fmt.Fprintf(os.Stdout, "Response from `JQLFunctionsAppsAPI.GetPrecomputationsByID`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPrecomputationsByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jqlFunctionPrecomputationGetByIdRequest** | [**JqlFunctionPrecomputationGetByIdRequest**](JqlFunctionPrecomputationGetByIdRequest.md) |  | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;functionKey&#x60; Sorts by the functionKey.  *  &#x60;used&#x60; Sorts by the used timestamp.  *  &#x60;created&#x60; Sorts by the created timestamp.  *  &#x60;updated&#x60; Sorts by the updated timestamp. | 

### Return type

[**JqlFunctionPrecomputationGetByIdResponse**](JqlFunctionPrecomputationGetByIdResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePrecomputations

> JqlFunctionPrecomputationUpdateResponse UpdatePrecomputations(ctx).JqlFunctionPrecomputationUpdateRequestBean(jqlFunctionPrecomputationUpdateRequestBean).SkipNotFoundPrecomputations(skipNotFoundPrecomputations).Execute()

Update precomputations (apps)



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
	jqlFunctionPrecomputationUpdateRequestBean := *openapiclient.NewJqlFunctionPrecomputationUpdateRequestBean() // JqlFunctionPrecomputationUpdateRequestBean | 
	skipNotFoundPrecomputations := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLFunctionsAppsAPI.UpdatePrecomputations(context.Background()).JqlFunctionPrecomputationUpdateRequestBean(jqlFunctionPrecomputationUpdateRequestBean).SkipNotFoundPrecomputations(skipNotFoundPrecomputations).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLFunctionsAppsAPI.UpdatePrecomputations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdatePrecomputations`: JqlFunctionPrecomputationUpdateResponse
	fmt.Fprintf(os.Stdout, "Response from `JQLFunctionsAppsAPI.UpdatePrecomputations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePrecomputationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jqlFunctionPrecomputationUpdateRequestBean** | [**JqlFunctionPrecomputationUpdateRequestBean**](JqlFunctionPrecomputationUpdateRequestBean.md) |  | 
 **skipNotFoundPrecomputations** | **bool** |  | [default to false]

### Return type

[**JqlFunctionPrecomputationUpdateResponse**](JqlFunctionPrecomputationUpdateResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

