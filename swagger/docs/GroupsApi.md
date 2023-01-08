# \GroupsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddUserToGroup**](GroupsAPI.md#AddUserToGroup) | **Post** /rest/api/3/group/user | Add user to group
[**BulkGetGroups**](GroupsAPI.md#BulkGetGroups) | **Get** /rest/api/3/group/bulk | Bulk get groups
[**CreateGroup**](GroupsAPI.md#CreateGroup) | **Post** /rest/api/3/group | Create group
[**FindGroups**](GroupsAPI.md#FindGroups) | **Get** /rest/api/3/groups/picker | Find groups
[**GetGroup**](GroupsAPI.md#GetGroup) | **Get** /rest/api/3/group | Get group
[**GetUsersFromGroup**](GroupsAPI.md#GetUsersFromGroup) | **Get** /rest/api/3/group/member | Get users from group
[**RemoveGroup**](GroupsAPI.md#RemoveGroup) | **Delete** /rest/api/3/group | Remove group
[**RemoveUserFromGroup**](GroupsAPI.md#RemoveUserFromGroup) | **Delete** /rest/api/3/group/user | Remove user from group



## AddUserToGroup

> Group AddUserToGroup(ctx).UpdateUserToGroupBean(updateUserToGroupBean).Groupname(groupname).GroupId(groupId).Execute()

Add user to group



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
	updateUserToGroupBean := *openapiclient.NewUpdateUserToGroupBean() // UpdateUserToGroupBean | The user to add to the group.
	groupname := "groupname_example" // string | As a group's name can change, use of `groupId` is recommended to identify a group.   The name of the group. This parameter cannot be used with the `groupId` parameter. (optional)
	groupId := "groupId_example" // string | The ID of the group. This parameter cannot be used with the `groupName` parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.AddUserToGroup(context.Background()).UpdateUserToGroupBean(updateUserToGroupBean).Groupname(groupname).GroupId(groupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.AddUserToGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddUserToGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.AddUserToGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddUserToGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateUserToGroupBean** | [**UpdateUserToGroupBean**](UpdateUserToGroupBean.md) | The user to add to the group. | 
 **groupname** | **string** | As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended to identify a group.   The name of the group. This parameter cannot be used with the &#x60;groupId&#x60; parameter. | 
 **groupId** | **string** | The ID of the group. This parameter cannot be used with the &#x60;groupName&#x60; parameter. | 

### Return type

[**Group**](Group.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkGetGroups

> PageBeanGroupDetails BulkGetGroups(ctx).StartAt(startAt).MaxResults(maxResults).GroupId(groupId).GroupName(groupName).AccessType(accessType).ApplicationKey(applicationKey).Execute()

Bulk get groups



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
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	groupId := []string{"3571b9a7-348f-414a-9087-8e1ea03a7df8"} // []string | The ID of a group. To specify multiple IDs, pass multiple `groupId` parameters. For example, `groupId=5b10a2844c20165700ede21g&groupId=5b10ac8d82e05b22cc7d4ef5`. (optional)
	groupName := []string{"Inner_example"} // []string | The name of a group. To specify multiple names, pass multiple `groupName` parameters. For example, `groupName=administrators&groupName=jira-software-users`. (optional)
	accessType := "accessType_example" // string | The access level of a group. Valid values: 'site-admin', 'admin', 'user'. (optional)
	applicationKey := "applicationKey_example" // string | The application key of the product user groups to search for. Valid values: 'jira-servicedesk', 'jira-software', 'jira-product-discovery', 'jira-core'. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.BulkGetGroups(context.Background()).StartAt(startAt).MaxResults(maxResults).GroupId(groupId).GroupName(groupName).AccessType(accessType).ApplicationKey(applicationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.BulkGetGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkGetGroups`: PageBeanGroupDetails
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.BulkGetGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkGetGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **groupId** | **[]string** | The ID of a group. To specify multiple IDs, pass multiple &#x60;groupId&#x60; parameters. For example, &#x60;groupId&#x3D;5b10a2844c20165700ede21g&amp;groupId&#x3D;5b10ac8d82e05b22cc7d4ef5&#x60;. | 
 **groupName** | **[]string** | The name of a group. To specify multiple names, pass multiple &#x60;groupName&#x60; parameters. For example, &#x60;groupName&#x3D;administrators&amp;groupName&#x3D;jira-software-users&#x60;. | 
 **accessType** | **string** | The access level of a group. Valid values: &#39;site-admin&#39;, &#39;admin&#39;, &#39;user&#39;. | 
 **applicationKey** | **string** | The application key of the product user groups to search for. Valid values: &#39;jira-servicedesk&#39;, &#39;jira-software&#39;, &#39;jira-product-discovery&#39;, &#39;jira-core&#39;. | 

### Return type

[**PageBeanGroupDetails**](PageBeanGroupDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateGroup

> Group CreateGroup(ctx).AddGroupBean(addGroupBean).Execute()

Create group



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
	addGroupBean := *openapiclient.NewAddGroupBean("Name_example") // AddGroupBean | The name of the group.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.CreateGroup(context.Background()).AddGroupBean(addGroupBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.CreateGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.CreateGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **addGroupBean** | [**AddGroupBean**](AddGroupBean.md) | The name of the group. | 

### Return type

[**Group**](Group.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindGroups

> FoundGroups FindGroups(ctx).AccountId(accountId).Query(query).Exclude(exclude).ExcludeId(excludeId).MaxResults(maxResults).CaseInsensitive(caseInsensitive).UserName(userName).Execute()

Find groups



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
	accountId := "accountId_example" // string | This parameter is deprecated, setting it does not affect the results. To find groups containing a particular user, use [Get user groups](#api-rest-api-3-user-groups-get). (optional)
	query := "query" // string | The string to find in group names. (optional)
	exclude := []string{"Inner_example"} // []string | As a group's name can change, use of `excludeGroupIds` is recommended to identify a group.   A group to exclude from the result. To exclude multiple groups, provide an ampersand-separated list. For example, `exclude=group1&exclude=group2`. This parameter cannot be used with the `excludeGroupIds` parameter. (optional)
	excludeId := []string{"Inner_example"} // []string | A group ID to exclude from the result. To exclude multiple groups, provide an ampersand-separated list. For example, `excludeId=group1-id&excludeId=group2-id`. This parameter cannot be used with the `excludeGroups` parameter. (optional)
	maxResults := int32(56) // int32 | The maximum number of groups to return. The maximum number of groups that can be returned is limited by the system property `jira.ajax.autocomplete.limit`. (optional)
	caseInsensitive := true // bool | Whether the search for groups should be case insensitive. (optional) (default to false)
	userName := "userName_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.FindGroups(context.Background()).AccountId(accountId).Query(query).Exclude(exclude).ExcludeId(excludeId).MaxResults(maxResults).CaseInsensitive(caseInsensitive).UserName(userName).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.FindGroups``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindGroups`: FoundGroups
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.FindGroups`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindGroupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | This parameter is deprecated, setting it does not affect the results. To find groups containing a particular user, use [Get user groups](#api-rest-api-3-user-groups-get). | 
 **query** | **string** | The string to find in group names. | 
 **exclude** | **[]string** | As a group&#39;s name can change, use of &#x60;excludeGroupIds&#x60; is recommended to identify a group.   A group to exclude from the result. To exclude multiple groups, provide an ampersand-separated list. For example, &#x60;exclude&#x3D;group1&amp;exclude&#x3D;group2&#x60;. This parameter cannot be used with the &#x60;excludeGroupIds&#x60; parameter. | 
 **excludeId** | **[]string** | A group ID to exclude from the result. To exclude multiple groups, provide an ampersand-separated list. For example, &#x60;excludeId&#x3D;group1-id&amp;excludeId&#x3D;group2-id&#x60;. This parameter cannot be used with the &#x60;excludeGroups&#x60; parameter. | 
 **maxResults** | **int32** | The maximum number of groups to return. The maximum number of groups that can be returned is limited by the system property &#x60;jira.ajax.autocomplete.limit&#x60;. | 
 **caseInsensitive** | **bool** | Whether the search for groups should be case insensitive. | [default to false]
 **userName** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 

### Return type

[**FoundGroups**](FoundGroups.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGroup

> Group GetGroup(ctx).Groupname(groupname).GroupId(groupId).Expand(expand).Execute()

Get group



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
	groupname := "groupname_example" // string | As a group's name can change, use of `groupId` is recommended to identify a group.   The name of the group. This parameter cannot be used with the `groupId` parameter. (optional)
	groupId := "groupId_example" // string | The ID of the group. This parameter cannot be used with the `groupName` parameter. (optional)
	expand := "expand_example" // string | List of fields to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetGroup(context.Background()).Groupname(groupname).GroupId(groupId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetGroup`: Group
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupname** | **string** | As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended to identify a group.   The name of the group. This parameter cannot be used with the &#x60;groupId&#x60; parameter. | 
 **groupId** | **string** | The ID of the group. This parameter cannot be used with the &#x60;groupName&#x60; parameter. | 
 **expand** | **string** | List of fields to expand. | 

### Return type

[**Group**](Group.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsersFromGroup

> PageBeanUserDetails GetUsersFromGroup(ctx).Groupname(groupname).GroupId(groupId).IncludeInactiveUsers(includeInactiveUsers).StartAt(startAt).MaxResults(maxResults).Execute()

Get users from group



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
	groupname := "groupname_example" // string | As a group's name can change, use of `groupId` is recommended to identify a group.   The name of the group. This parameter cannot be used with the `groupId` parameter. (optional)
	groupId := "groupId_example" // string | The ID of the group. This parameter cannot be used with the `groupName` parameter. (optional)
	includeInactiveUsers := true // bool | Include inactive users. (optional) (default to false)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page (number should be between 1 and 50). (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GroupsAPI.GetUsersFromGroup(context.Background()).Groupname(groupname).GroupId(groupId).IncludeInactiveUsers(includeInactiveUsers).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.GetUsersFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUsersFromGroup`: PageBeanUserDetails
	fmt.Fprintf(os.Stdout, "Response from `GroupsAPI.GetUsersFromGroup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUsersFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupname** | **string** | As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended to identify a group.   The name of the group. This parameter cannot be used with the &#x60;groupId&#x60; parameter. | 
 **groupId** | **string** | The ID of the group. This parameter cannot be used with the &#x60;groupName&#x60; parameter. | 
 **includeInactiveUsers** | **bool** | Include inactive users. | [default to false]
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page (number should be between 1 and 50). | [default to 50]

### Return type

[**PageBeanUserDetails**](PageBeanUserDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGroup

> RemoveGroup(ctx).Groupname(groupname).GroupId(groupId).SwapGroup(swapGroup).SwapGroupId(swapGroupId).Execute()

Remove group



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
	groupname := "groupname_example" // string |  (optional)
	groupId := "groupId_example" // string | The ID of the group. This parameter cannot be used with the `groupname` parameter. (optional)
	swapGroup := "swapGroup_example" // string | As a group's name can change, use of `swapGroupId` is recommended to identify a group.   The group to transfer restrictions to. Only comments and worklogs are transferred. If restrictions are not transferred, comments and worklogs are inaccessible after the deletion. This parameter cannot be used with the `swapGroupId` parameter. (optional)
	swapGroupId := "swapGroupId_example" // string | The ID of the group to transfer restrictions to. Only comments and worklogs are transferred. If restrictions are not transferred, comments and worklogs are inaccessible after the deletion. This parameter cannot be used with the `swapGroup` parameter. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveGroup(context.Background()).Groupname(groupname).GroupId(groupId).SwapGroup(swapGroup).SwapGroupId(swapGroupId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **groupname** | **string** |  | 
 **groupId** | **string** | The ID of the group. This parameter cannot be used with the &#x60;groupname&#x60; parameter. | 
 **swapGroup** | **string** | As a group&#39;s name can change, use of &#x60;swapGroupId&#x60; is recommended to identify a group.   The group to transfer restrictions to. Only comments and worklogs are transferred. If restrictions are not transferred, comments and worklogs are inaccessible after the deletion. This parameter cannot be used with the &#x60;swapGroupId&#x60; parameter. | 
 **swapGroupId** | **string** | The ID of the group to transfer restrictions to. Only comments and worklogs are transferred. If restrictions are not transferred, comments and worklogs are inaccessible after the deletion. This parameter cannot be used with the &#x60;swapGroup&#x60; parameter. | 

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


## RemoveUserFromGroup

> RemoveUserFromGroup(ctx).AccountId(accountId).Groupname(groupname).GroupId(groupId).Username(username).Execute()

Remove user from group



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
	accountId := "5b10ac8d82e05b22cc7d4ef5" // string | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*.
	groupname := "groupname_example" // string | As a group's name can change, use of `groupId` is recommended to identify a group.   The name of the group. This parameter cannot be used with the `groupId` parameter. (optional)
	groupId := "groupId_example" // string | The ID of the group. This parameter cannot be used with the `groupName` parameter. (optional)
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.GroupsAPI.RemoveUserFromGroup(context.Background()).AccountId(accountId).Groupname(groupname).GroupId(groupId).Username(username).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GroupsAPI.RemoveUserFromGroup``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveUserFromGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountId** | **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | 
 **groupname** | **string** | As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended to identify a group.   The name of the group. This parameter cannot be used with the &#x60;groupId&#x60; parameter. | 
 **groupId** | **string** | The ID of the group. This parameter cannot be used with the &#x60;groupName&#x60; parameter. | 
 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 

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

