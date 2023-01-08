# \IssueLinksAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteIssueLink**](IssueLinksAPI.md#DeleteIssueLink) | **Delete** /rest/api/3/issueLink/{linkId} | Delete issue link
[**GetIssueLink**](IssueLinksAPI.md#GetIssueLink) | **Get** /rest/api/3/issueLink/{linkId} | Get issue link
[**LinkIssues**](IssueLinksAPI.md#LinkIssues) | **Post** /rest/api/3/issueLink | Create issue link



## DeleteIssueLink

> DeleteIssueLink(ctx, linkId).Execute()

Delete issue link



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
	linkId := "linkId_example" // string | The ID of the issue link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueLinksAPI.DeleteIssueLink(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinksAPI.DeleteIssueLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** | The ID of the issue link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueLinkRequest struct via the builder pattern


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


## GetIssueLink

> IssueLink GetIssueLink(ctx, linkId).Execute()

Get issue link



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
	linkId := "linkId_example" // string | The ID of the issue link.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueLinksAPI.GetIssueLink(context.Background(), linkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinksAPI.GetIssueLink``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueLink`: IssueLink
	fmt.Fprintf(os.Stdout, "Response from `IssueLinksAPI.GetIssueLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**linkId** | **string** | The ID of the issue link. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueLink**](IssueLink.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkIssues

> interface{} LinkIssues(ctx).LinkIssueRequestJsonBean(linkIssueRequestJsonBean).Execute()

Create issue link



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
	linkIssueRequestJsonBean := *openapiclient.NewLinkIssueRequestJsonBean(*openapiclient.NewLinkedIssue(), *openapiclient.NewLinkedIssue(), *openapiclient.NewIssueLinkType()) // LinkIssueRequestJsonBean | The issue link request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueLinksAPI.LinkIssues(context.Background()).LinkIssueRequestJsonBean(linkIssueRequestJsonBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueLinksAPI.LinkIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `LinkIssues`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueLinksAPI.LinkIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkIssueRequestJsonBean** | [**LinkIssueRequestJsonBean**](LinkIssueRequestJsonBean.md) | The issue link request. | 

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

