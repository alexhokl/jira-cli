# \IssueRemoteLinksAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOrUpdateRemoteIssueLink**](IssueRemoteLinksAPI.md#CreateOrUpdateRemoteIssueLink) | **Post** /rest/api/3/issue/{issueIdOrKey}/remotelink | Create or update remote issue link
[**DeleteRemoteIssueLinkByGlobalId**](IssueRemoteLinksAPI.md#DeleteRemoteIssueLinkByGlobalId) | **Delete** /rest/api/3/issue/{issueIdOrKey}/remotelink | Delete remote issue link by global ID
[**DeleteRemoteIssueLinkById**](IssueRemoteLinksAPI.md#DeleteRemoteIssueLinkById) | **Delete** /rest/api/3/issue/{issueIdOrKey}/remotelink/{linkId} | Delete remote issue link by ID
[**GetRemoteIssueLinkById**](IssueRemoteLinksAPI.md#GetRemoteIssueLinkById) | **Get** /rest/api/3/issue/{issueIdOrKey}/remotelink/{linkId} | Get remote issue link by ID
[**GetRemoteIssueLinks**](IssueRemoteLinksAPI.md#GetRemoteIssueLinks) | **Get** /rest/api/3/issue/{issueIdOrKey}/remotelink | Get remote issue links
[**UpdateRemoteIssueLink**](IssueRemoteLinksAPI.md#UpdateRemoteIssueLink) | **Put** /rest/api/3/issue/{issueIdOrKey}/remotelink/{linkId} | Update remote issue link by ID



## CreateOrUpdateRemoteIssueLink

> RemoteIssueLinkIdentifies CreateOrUpdateRemoteIssueLink(ctx, issueIdOrKey).RemoteIssueLinkRequest(remoteIssueLinkRequest).Execute()

Create or update remote issue link



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
	remoteIssueLinkRequest := *openapiclient.NewRemoteIssueLinkRequest(*openapiclient.NewRemoteObject("Title_example", "Url_example")) // RemoteIssueLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRemoteLinksAPI.CreateOrUpdateRemoteIssueLink(context.Background(), issueIdOrKey).RemoteIssueLinkRequest(remoteIssueLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.CreateOrUpdateRemoteIssueLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOrUpdateRemoteIssueLink`: RemoteIssueLinkIdentifies
	fmt.Fprintf(os.Stdout, "Response from `IssueRemoteLinksAPI.CreateOrUpdateRemoteIssueLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrUpdateRemoteIssueLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **remoteIssueLinkRequest** | [**RemoteIssueLinkRequest**](RemoteIssueLinkRequest.md) |  | 

### Return type

[**RemoteIssueLinkIdentifies**](RemoteIssueLinkIdentifies.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemoteIssueLinkByGlobalId

> DeleteRemoteIssueLinkByGlobalId(ctx, issueIdOrKey).GlobalId(globalId).Execute()

Delete remote issue link by global ID



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
	issueIdOrKey := "10000" // string | The ID or key of the issue.
	globalId := "system=http://www.mycompany.com/support&id=1" // string | The global ID of a remote issue link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueRemoteLinksAPI.DeleteRemoteIssueLinkByGlobalId(context.Background(), issueIdOrKey).GlobalId(globalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.DeleteRemoteIssueLinkByGlobalId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteIssueLinkByGlobalIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalId** | **string** | The global ID of a remote issue link. | 

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


## DeleteRemoteIssueLinkById

> DeleteRemoteIssueLinkById(ctx, issueIdOrKey, linkId).Execute()

Delete remote issue link by ID



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
	issueIdOrKey := "10000" // string | The ID or key of the issue.
	linkId := "10000" // string | The ID of a remote issue link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueRemoteLinksAPI.DeleteRemoteIssueLinkById(context.Background(), issueIdOrKey, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.DeleteRemoteIssueLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**linkId** | **string** | The ID of a remote issue link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteIssueLinkByIdRequest struct via the builder pattern


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


## GetRemoteIssueLinkById

> RemoteIssueLink GetRemoteIssueLinkById(ctx, issueIdOrKey, linkId).Execute()

Get remote issue link by ID



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
	linkId := "linkId_example" // string | The ID of the remote issue link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRemoteLinksAPI.GetRemoteIssueLinkById(context.Background(), issueIdOrKey, linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.GetRemoteIssueLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteIssueLinkById`: RemoteIssueLink
	fmt.Fprintf(os.Stdout, "Response from `IssueRemoteLinksAPI.GetRemoteIssueLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**linkId** | **string** | The ID of the remote issue link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteIssueLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**RemoteIssueLink**](RemoteIssueLink.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteIssueLinks

> GetRemoteIssueLinks200Response GetRemoteIssueLinks(ctx, issueIdOrKey).GlobalId(globalId).Execute()

Get remote issue links



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
	issueIdOrKey := "10000" // string | The ID or key of the issue.
	globalId := "globalId_example" // string | The global ID of the remote issue link. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRemoteLinksAPI.GetRemoteIssueLinks(context.Background(), issueIdOrKey).GlobalId(globalId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.GetRemoteIssueLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteIssueLinks`: GetRemoteIssueLinks200Response
	fmt.Fprintf(os.Stdout, "Response from `IssueRemoteLinksAPI.GetRemoteIssueLinks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteIssueLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **globalId** | **string** | The global ID of the remote issue link. | 

### Return type

[**GetRemoteIssueLinks200Response**](GetRemoteIssueLinks200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRemoteIssueLink

> interface{} UpdateRemoteIssueLink(ctx, issueIdOrKey, linkId).RemoteIssueLinkRequest(remoteIssueLinkRequest).Execute()

Update remote issue link by ID



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
	issueIdOrKey := "10000" // string | The ID or key of the issue.
	linkId := "10000" // string | The ID of the remote issue link.
	remoteIssueLinkRequest := *openapiclient.NewRemoteIssueLinkRequest(*openapiclient.NewRemoteObject("Title_example", "Url_example")) // RemoteIssueLinkRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRemoteLinksAPI.UpdateRemoteIssueLink(context.Background(), issueIdOrKey, linkId).RemoteIssueLinkRequest(remoteIssueLinkRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRemoteLinksAPI.UpdateRemoteIssueLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRemoteIssueLink`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueRemoteLinksAPI.UpdateRemoteIssueLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**linkId** | **string** | The ID of the remote issue link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRemoteIssueLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **remoteIssueLinkRequest** | [**RemoteIssueLinkRequest**](RemoteIssueLinkRequest.md) |  | 

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

