# \ServiceRegistryAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ServiceRegistryResourceServicesGet**](ServiceRegistryAPI.md#ServiceRegistryResourceServicesGet) | **Get** /rest/atlassian-connect/1/service-registry | Retrieve the attributes of service registries



## ServiceRegistryResourceServicesGet

> []ServiceRegistry ServiceRegistryResourceServicesGet(ctx).ServiceIds(serviceIds).Execute()

Retrieve the attributes of service registries



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
	serviceIds := []string{"Inner_example"} // []string | The ID of the services (the strings starting with \"b:\" need to be decoded in Base64).

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ServiceRegistryAPI.ServiceRegistryResourceServicesGet(context.Background()).ServiceIds(serviceIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ServiceRegistryAPI.ServiceRegistryResourceServicesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ServiceRegistryResourceServicesGet`: []ServiceRegistry
	fmt.Fprintf(os.Stdout, "Response from `ServiceRegistryAPI.ServiceRegistryResourceServicesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiServiceRegistryResourceServicesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serviceIds** | **[]string** | The ID of the services (the strings starting with \&quot;b:\&quot; need to be decoded in Base64). | 

### Return type

[**[]ServiceRegistry**](ServiceRegistry.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

