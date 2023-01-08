# \DynamicModulesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DynamicModulesResourceGetModulesGet**](DynamicModulesAPI.md#DynamicModulesResourceGetModulesGet) | **Get** /rest/atlassian-connect/1/app/module/dynamic | Get modules
[**DynamicModulesResourceRegisterModulesPost**](DynamicModulesAPI.md#DynamicModulesResourceRegisterModulesPost) | **Post** /rest/atlassian-connect/1/app/module/dynamic | Register modules
[**DynamicModulesResourceRemoveModulesDelete**](DynamicModulesAPI.md#DynamicModulesResourceRemoveModulesDelete) | **Delete** /rest/atlassian-connect/1/app/module/dynamic | Remove modules



## DynamicModulesResourceGetModulesGet

> ConnectModules DynamicModulesResourceGetModulesGet(ctx).Execute()

Get modules



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
	resp, r, err := apiClient.DynamicModulesAPI.DynamicModulesResourceGetModulesGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicModulesAPI.DynamicModulesResourceGetModulesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DynamicModulesResourceGetModulesGet`: ConnectModules
	fmt.Fprintf(os.Stdout, "Response from `DynamicModulesAPI.DynamicModulesResourceGetModulesGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDynamicModulesResourceGetModulesGetRequest struct via the builder pattern


### Return type

[**ConnectModules**](ConnectModules.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DynamicModulesResourceRegisterModulesPost

> DynamicModulesResourceRegisterModulesPost(ctx).ConnectModules(connectModules).Execute()

Register modules



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
	connectModules := *openapiclient.NewConnectModules([]map[string]interface{}{map[string]interface{}({"description":{"value":"field with team"},"type":"single_select","extractions":[{"path":"category","type":"text","name":"categoryName"}],"name":{"value":"Team"},"key":"team-field"})}) // ConnectModules | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DynamicModulesAPI.DynamicModulesResourceRegisterModulesPost(context.Background()).ConnectModules(connectModules).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicModulesAPI.DynamicModulesResourceRegisterModulesPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDynamicModulesResourceRegisterModulesPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **connectModules** | [**ConnectModules**](ConnectModules.md) |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DynamicModulesResourceRemoveModulesDelete

> DynamicModulesResourceRemoveModulesDelete(ctx).ModuleKey(moduleKey).Execute()

Remove modules



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
	moduleKey := []string{"Inner_example"} // []string | The key of the module to remove. To include multiple module keys, provide multiple copies of this parameter. For example, `moduleKey=dynamic-attachment-entity-property&moduleKey=dynamic-select-field`. Nonexistent keys are ignored. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DynamicModulesAPI.DynamicModulesResourceRemoveModulesDelete(context.Background()).ModuleKey(moduleKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DynamicModulesAPI.DynamicModulesResourceRemoveModulesDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDynamicModulesResourceRemoveModulesDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **moduleKey** | **[]string** | The key of the module to remove. To include multiple module keys, provide multiple copies of this parameter. For example, &#x60;moduleKey&#x3D;dynamic-attachment-entity-property&amp;moduleKey&#x3D;dynamic-select-field&#x60;. Nonexistent keys are ignored. | 

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

