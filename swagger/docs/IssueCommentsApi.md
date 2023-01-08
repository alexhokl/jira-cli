# \IssueCommentsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddComment**](IssueCommentsAPI.md#AddComment) | **Post** /rest/api/3/issue/{issueIdOrKey}/comment | Add comment
[**DeleteComment**](IssueCommentsAPI.md#DeleteComment) | **Delete** /rest/api/3/issue/{issueIdOrKey}/comment/{id} | Delete comment
[**GetComment**](IssueCommentsAPI.md#GetComment) | **Get** /rest/api/3/issue/{issueIdOrKey}/comment/{id} | Get comment
[**GetComments**](IssueCommentsAPI.md#GetComments) | **Get** /rest/api/3/issue/{issueIdOrKey}/comment | Get comments
[**GetCommentsByIds**](IssueCommentsAPI.md#GetCommentsByIds) | **Post** /rest/api/3/comment/list | Get comments by IDs
[**UpdateComment**](IssueCommentsAPI.md#UpdateComment) | **Put** /rest/api/3/issue/{issueIdOrKey}/comment/{id} | Update comment



## AddComment

> Comment AddComment(ctx, issueIdOrKey).Comment(comment).Expand(expand).Execute()

Add comment



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
	comment := *openapiclient.NewComment() // Comment | 
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `renderedBody`, which returns the comment body rendered in HTML. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentsAPI.AddComment(context.Background(), issueIdOrKey).Comment(comment).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.AddComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentsAPI.AddComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **comment** | [**Comment**](Comment.md) |  | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts &#x60;renderedBody&#x60;, which returns the comment body rendered in HTML. | 

### Return type

[**Comment**](Comment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteComment

> DeleteComment(ctx, issueIdOrKey, id).Execute()

Delete comment



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
	id := "id_example" // string | The ID of the comment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueCommentsAPI.DeleteComment(context.Background(), issueIdOrKey, id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.DeleteComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**id** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommentRequest struct via the builder pattern


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


## GetComment

> Comment GetComment(ctx, issueIdOrKey, id).Expand(expand).Execute()

Get comment



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
	id := "id_example" // string | The ID of the comment.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `renderedBody`, which returns the comment body rendered in HTML. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentsAPI.GetComment(context.Background(), issueIdOrKey, id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.GetComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentsAPI.GetComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**id** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **expand** | **string** | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts &#x60;renderedBody&#x60;, which returns the comment body rendered in HTML. | 

### Return type

[**Comment**](Comment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetComments

> PageOfComments GetComments(ctx, issueIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Expand(expand).Execute()

Get comments



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field. Accepts *created* to sort comments by their created date. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `renderedBody`, which returns the comment body rendered in HTML. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentsAPI.GetComments(context.Background(), issueIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.GetComments``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetComments`: PageOfComments
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentsAPI.GetComments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]
 **orderBy** | **string** | [Order](#ordering) the results by a field. Accepts *created* to sort comments by their created date. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts &#x60;renderedBody&#x60;, which returns the comment body rendered in HTML. | 

### Return type

[**PageOfComments**](PageOfComments.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommentsByIds

> PageBeanComment GetCommentsByIds(ctx).IssueCommentListRequestBean(issueCommentListRequestBean).Expand(expand).Execute()

Get comments by IDs



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
	issueCommentListRequestBean := *openapiclient.NewIssueCommentListRequestBean([]int64{int64(123)}) // IssueCommentListRequestBean | The list of comment IDs.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts a comma-separated list. Expand options include:   *  `renderedBody` Returns the comment body rendered in HTML.  *  `properties` Returns the comment's properties. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentsAPI.GetCommentsByIds(context.Background()).IssueCommentListRequestBean(issueCommentListRequestBean).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.GetCommentsByIds``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCommentsByIds`: PageBeanComment
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentsAPI.GetCommentsByIds`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCommentsByIdsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueCommentListRequestBean** | [**IssueCommentListRequestBean**](IssueCommentListRequestBean.md) | The list of comment IDs. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;renderedBody&#x60; Returns the comment body rendered in HTML.  *  &#x60;properties&#x60; Returns the comment&#39;s properties. | 

### Return type

[**PageBeanComment**](PageBeanComment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateComment

> Comment UpdateComment(ctx, issueIdOrKey, id).Comment(comment).NotifyUsers(notifyUsers).OverrideEditableFlag(overrideEditableFlag).Expand(expand).Execute()

Update comment



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
	id := "id_example" // string | The ID of the comment.
	comment := *openapiclient.NewComment() // Comment | 
	notifyUsers := true // bool | Whether users are notified when a comment is updated. (optional) (default to true)
	overrideEditableFlag := true // bool | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect app users with the *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts `renderedBody`, which returns the comment body rendered in HTML. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCommentsAPI.UpdateComment(context.Background(), issueIdOrKey, id).Comment(comment).NotifyUsers(notifyUsers).OverrideEditableFlag(overrideEditableFlag).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCommentsAPI.UpdateComment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateComment`: Comment
	fmt.Fprintf(os.Stdout, "Response from `IssueCommentsAPI.UpdateComment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue. | 
**id** | **string** | The ID of the comment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCommentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **comment** | [**Comment**](Comment.md) |  | 
 **notifyUsers** | **bool** | Whether users are notified when a comment is updated. | [default to true]
 **overrideEditableFlag** | **bool** | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect app users with the *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) and Forge apps acting on behalf of users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]
 **expand** | **string** | Use [expand](#expansion) to include additional information about comments in the response. This parameter accepts &#x60;renderedBody&#x60;, which returns the comment body rendered in HTML. | 

### Return type

[**Comment**](Comment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

