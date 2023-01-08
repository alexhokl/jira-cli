# \IssueSecurityLevelAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssueSecurityLevel**](IssueSecurityLevelAPI.md#GetIssueSecurityLevel) | **Get** /rest/api/3/securitylevel/{id} | Get issue security level
[**GetIssueSecurityLevelMembers**](IssueSecurityLevelAPI.md#GetIssueSecurityLevelMembers) | **Get** /rest/api/3/issuesecurityschemes/{issueSecuritySchemeId}/members | Get issue security level members by issue security scheme



## GetIssueSecurityLevel

> SecurityLevel GetIssueSecurityLevel(ctx, id).Execute()

Get issue security level



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
	id := "id_example" // string | The ID of the issue security level.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecurityLevelAPI.GetIssueSecurityLevel(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecurityLevelAPI.GetIssueSecurityLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueSecurityLevel`: SecurityLevel
	fmt.Fprintf(os.Stdout, "Response from `IssueSecurityLevelAPI.GetIssueSecurityLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue security level. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueSecurityLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityLevel**](SecurityLevel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueSecurityLevelMembers

> PageBeanIssueSecurityLevelMember GetIssueSecurityLevelMembers(ctx, issueSecuritySchemeId).StartAt(startAt).MaxResults(maxResults).IssueSecurityLevelId(issueSecurityLevelId).Expand(expand).Execute()

Get issue security level members by issue security scheme



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
	issueSecuritySchemeId := int64(789) // int64 | The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) operation to get a list of issue security scheme IDs.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	issueSecurityLevelId := []string{"Inner_example"} // []string | The list of issue security level IDs. To include multiple issue security levels separate IDs with ampersand: `issueSecurityLevelId=10000&issueSecurityLevelId=10001`. (optional)
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `all` Returns all expandable information.  *  `field` Returns information about the custom field granted the permission.  *  `group` Returns information about the group that is granted the permission.  *  `projectRole` Returns information about the project role granted the permission.  *  `user` Returns information about the user who is granted the permission. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecurityLevelAPI.GetIssueSecurityLevelMembers(context.Background(), issueSecuritySchemeId).StartAt(startAt).MaxResults(maxResults).IssueSecurityLevelId(issueSecurityLevelId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecurityLevelAPI.GetIssueSecurityLevelMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueSecurityLevelMembers`: PageBeanIssueSecurityLevelMember
	fmt.Fprintf(os.Stdout, "Response from `IssueSecurityLevelAPI.GetIssueSecurityLevelMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueSecuritySchemeId** | **int64** | The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) operation to get a list of issue security scheme IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueSecurityLevelMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **issueSecurityLevelId** | **[]string** | The list of issue security level IDs. To include multiple issue security levels separate IDs with ampersand: &#x60;issueSecurityLevelId&#x3D;10000&amp;issueSecurityLevelId&#x3D;10001&#x60;. | 
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information.  *  &#x60;field&#x60; Returns information about the custom field granted the permission.  *  &#x60;group&#x60; Returns information about the group that is granted the permission.  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission.  *  &#x60;user&#x60; Returns information about the user who is granted the permission. | 

### Return type

[**PageBeanIssueSecurityLevelMember**](PageBeanIssueSecurityLevelMember.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

