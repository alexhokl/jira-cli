# \GroupAndUserPickerAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindUsersAndGroups**](GroupAndUserPickerAPI.md#FindUsersAndGroups) | **Get** /rest/api/3/groupuserpicker | Find users and groups



## FindUsersAndGroups

> FoundUsersAndGroups FindUsersAndGroups(ctx).Query(query).MaxResults(maxResults).ShowAvatar(showAvatar).FieldId(fieldId).ProjectId(projectId).IssueTypeId(issueTypeId).AvatarSize(avatarSize).CaseInsensitive(caseInsensitive).ExcludeConnectAddons(excludeConnectAddons).Execute()

Find users and groups



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
	query := "query_example" // string | The search string.
	maxResults := int32(56) // int32 | The maximum number of items to return in each list. (optional) (default to 50)
	showAvatar := true // bool | Whether the user avatar should be returned. If an invalid value is provided, the default value is used. (optional) (default to false)
	fieldId := "fieldId_example" // string | The custom field ID of the field this request is for. (optional)
	projectId := []string{"Inner_example"} // []string | The ID of a project that returned users and groups must have permission to view. To include multiple projects, provide an ampersand-separated list. For example, `projectId=10000&projectId=10001`. This parameter is only used when `fieldId` is present. (optional)
	issueTypeId := []string{"Inner_example"} // []string | The ID of an issue type that returned users and groups must have permission to view. To include multiple issue types, provide an ampersand-separated list. For example, `issueTypeId=10000&issueTypeId=10001`. Special values, such as `-1` (all standard issue types) and `-2` (all subtask issue types), are supported. This parameter is only used when `fieldId` is present. (optional)
	avatarSize := "avatarSize_example" // string | The size of the avatar to return. If an invalid value is provided, the default value is used. (optional) (default to "xsmall")
	caseInsensitive := true // bool | Whether the search for groups should be case insensitive. (optional) (default to false)
	excludeConnectAddons := true // bool | Whether Connect app users and groups should be excluded from the search results. If an invalid value is provided, the default value is used. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupAndUserPickerAPI.FindUsersAndGroups(context.Background()).Query(query).MaxResults(maxResults).ShowAvatar(showAvatar).FieldId(fieldId).ProjectId(projectId).IssueTypeId(issueTypeId).AvatarSize(avatarSize).CaseInsensitive(caseInsensitive).ExcludeConnectAddons(excludeConnectAddons).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupAndUserPickerAPI.FindUsersAndGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersAndGroups`: FoundUsersAndGroups
	fmt.Fprintf(os.Stdout, "Response from `GroupAndUserPickerAPI.FindUsersAndGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersAndGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search string. | 
 **maxResults** | **int32** | The maximum number of items to return in each list. | [default to 50]
 **showAvatar** | **bool** | Whether the user avatar should be returned. If an invalid value is provided, the default value is used. | [default to false]
 **fieldId** | **string** | The custom field ID of the field this request is for. | 
 **projectId** | **[]string** | The ID of a project that returned users and groups must have permission to view. To include multiple projects, provide an ampersand-separated list. For example, &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. This parameter is only used when &#x60;fieldId&#x60; is present. | 
 **issueTypeId** | **[]string** | The ID of an issue type that returned users and groups must have permission to view. To include multiple issue types, provide an ampersand-separated list. For example, &#x60;issueTypeId&#x3D;10000&amp;issueTypeId&#x3D;10001&#x60;. Special values, such as &#x60;-1&#x60; (all standard issue types) and &#x60;-2&#x60; (all subtask issue types), are supported. This parameter is only used when &#x60;fieldId&#x60; is present. | 
 **avatarSize** | **string** | The size of the avatar to return. If an invalid value is provided, the default value is used. | [default to &quot;xsmall&quot;]
 **caseInsensitive** | **bool** | Whether the search for groups should be case insensitive. | [default to false]
 **excludeConnectAddons** | **bool** | Whether Connect app users and groups should be excluded from the search results. If an invalid value is provided, the default value is used. | [default to false]

### Return type

[**FoundUsersAndGroups**](FoundUsersAndGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

