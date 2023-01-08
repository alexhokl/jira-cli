# \IssueVotesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddVote**](IssueVotesAPI.md#AddVote) | **Post** /rest/api/3/issue/{issueIdOrKey}/votes | Add vote
[**GetVotes**](IssueVotesAPI.md#GetVotes) | **Get** /rest/api/3/issue/{issueIdOrKey}/votes | Get votes
[**RemoveVote**](IssueVotesAPI.md#RemoveVote) | **Delete** /rest/api/3/issue/{issueIdOrKey}/votes | Delete vote



## AddVote

> interface{} AddVote(ctx, issueIdOrKey).Execute()

Add vote



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueVotesAPI.AddVote(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueVotesAPI.AddVote``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddVote`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueVotesAPI.AddVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVotes

> Votes GetVotes(ctx, issueIdOrKey).Execute()

Get votes



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueVotesAPI.GetVotes(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueVotesAPI.GetVotes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVotes`: Votes
	fmt.Fprintf(os.Stdout, "Response from `IssueVotesAPI.GetVotes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVotesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Votes**](Votes.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveVote

> RemoveVote(ctx, issueIdOrKey).Execute()

Delete vote



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueVotesAPI.RemoveVote(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueVotesAPI.RemoveVote``: %v\n", err)
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

Other parameters are passed through a pointer to a apiRemoveVoteRequest struct via the builder pattern


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

