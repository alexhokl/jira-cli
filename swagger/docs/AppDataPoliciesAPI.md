# \AppDataPoliciesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPolicies**](AppDataPoliciesAPI.md#GetPolicies) | **Get** /rest/api/3/data-policy/project | Get data policy for projects
[**GetPolicy**](AppDataPoliciesAPI.md#GetPolicy) | **Get** /rest/api/3/data-policy | Get data policy for the workspace



## GetPolicies

> ProjectDataPolicies GetPolicies(ctx).Ids(ids).Execute()

Get data policy for projects



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
	ids := "ids_example" // string | A list of project identifiers. This parameter accepts a comma-separated list. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppDataPoliciesAPI.GetPolicies(context.Background()).Ids(ids).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataPoliciesAPI.GetPolicies``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicies`: ProjectDataPolicies
	fmt.Fprintf(os.Stdout, "Response from `AppDataPoliciesAPI.GetPolicies`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPoliciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | A list of project identifiers. This parameter accepts a comma-separated list. | 

### Return type

[**ProjectDataPolicies**](ProjectDataPolicies.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPolicy

> WorkspaceDataPolicy GetPolicy(ctx).Execute()

Get data policy for the workspace



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
	resp, r, err := apiClient.AppDataPoliciesAPI.GetPolicy(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppDataPoliciesAPI.GetPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPolicy`: WorkspaceDataPolicy
	fmt.Fprintf(os.Stdout, "Response from `AppDataPoliciesAPI.GetPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


### Return type

[**WorkspaceDataPolicy**](WorkspaceDataPolicy.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

