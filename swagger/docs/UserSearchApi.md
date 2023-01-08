# \UserSearchAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAssignableUsers**](UserSearchAPI.md#FindAssignableUsers) | **Get** /rest/api/3/user/assignable/search | Find users assignable to issues
[**FindBulkAssignableUsers**](UserSearchAPI.md#FindBulkAssignableUsers) | **Get** /rest/api/3/user/assignable/multiProjectSearch | Find users assignable to projects
[**FindUserKeysByQuery**](UserSearchAPI.md#FindUserKeysByQuery) | **Get** /rest/api/3/user/search/query/key | Find user keys by query
[**FindUsers**](UserSearchAPI.md#FindUsers) | **Get** /rest/api/3/user/search | Find users
[**FindUsersByQuery**](UserSearchAPI.md#FindUsersByQuery) | **Get** /rest/api/3/user/search/query | Find users by query
[**FindUsersForPicker**](UserSearchAPI.md#FindUsersForPicker) | **Get** /rest/api/3/user/picker | Find users for picker
[**FindUsersWithAllPermissions**](UserSearchAPI.md#FindUsersWithAllPermissions) | **Get** /rest/api/3/user/permission/search | Find users with permissions
[**FindUsersWithBrowsePermission**](UserSearchAPI.md#FindUsersWithBrowsePermission) | **Get** /rest/api/3/user/viewissue/search | Find users with browse permission



## FindAssignableUsers

> []User FindAssignableUsers(ctx).Query(query).SessionId(sessionId).Username(username).AccountId(accountId).Project(project).IssueKey(issueKey).IssueId(issueId).StartAt(startAt).MaxResults(maxResults).ActionDescriptorId(actionDescriptorId).Recommend(recommend).AccountType(accountType).AppType(appType).Execute()

Find users assignable to issues



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
	query := "query" // string | A query string that is matched against user attributes, such as `displayName`, and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `username` or `accountId` is specified. (optional)
	sessionId := "sessionId_example" // string | The sessionId of this request. SessionId is the same until the assignee is set. (optional)
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	accountId := "accountId_example" // string | A query string that is matched exactly against user `accountId`. Required, unless `query` is specified. (optional)
	project := "project_example" // string | The project ID or project key (case sensitive). Required, unless `issueKey` or `issueId` is specified. (optional)
	issueKey := "issueKey_example" // string | The key of the issue. Required, unless `issueId` or `project` is specified. (optional)
	issueId := "issueId_example" // string | The ID of the issue. Required, unless `issueKey` or `project` is specified. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return. This operation may return less than the maximum number of items even if more are available. The operation fetches users up to the maximum and then, from the fetched users, returns only the users that can be assigned to the issue. (optional) (default to 50)
	actionDescriptorId := int32(56) // int32 | The ID of the transition. (optional)
	recommend := true // bool |  (optional) (default to false)
	accountType := []string{"Inner_example"} // []string |  (optional)
	appType := []string{"Inner_example"} // []string |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindAssignableUsers(context.Background()).Query(query).SessionId(sessionId).Username(username).AccountId(accountId).Project(project).IssueKey(issueKey).IssueId(issueId).StartAt(startAt).MaxResults(maxResults).ActionDescriptorId(actionDescriptorId).Recommend(recommend).AccountType(accountType).AppType(appType).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindAssignableUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindAssignableUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindAssignableUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindAssignableUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A query string that is matched against user attributes, such as &#x60;displayName&#x60;, and &#x60;emailAddress&#x60;, to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. Required, unless &#x60;username&#x60; or &#x60;accountId&#x60; is specified. | 
 **sessionId** | **string** | The sessionId of this request. SessionId is the same until the assignee is set. | 
 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **accountId** | **string** | A query string that is matched exactly against user &#x60;accountId&#x60;. Required, unless &#x60;query&#x60; is specified. | 
 **project** | **string** | The project ID or project key (case sensitive). Required, unless &#x60;issueKey&#x60; or &#x60;issueId&#x60; is specified. | 
 **issueKey** | **string** | The key of the issue. Required, unless &#x60;issueId&#x60; or &#x60;project&#x60; is specified. | 
 **issueId** | **string** | The ID of the issue. Required, unless &#x60;issueKey&#x60; or &#x60;project&#x60; is specified. | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return. This operation may return less than the maximum number of items even if more are available. The operation fetches users up to the maximum and then, from the fetched users, returns only the users that can be assigned to the issue. | [default to 50]
 **actionDescriptorId** | **int32** | The ID of the transition. | 
 **recommend** | **bool** |  | [default to false]
 **accountType** | **[]string** |  | 
 **appType** | **[]string** |  | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindBulkAssignableUsers

> []User FindBulkAssignableUsers(ctx).ProjectKeys(projectKeys).Query(query).Username(username).AccountId(accountId).StartAt(startAt).MaxResults(maxResults).Execute()

Find users assignable to projects



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
	projectKeys := "projectKeys_example" // string | A list of project keys (case sensitive). This parameter accepts a comma-separated list.
	query := "query" // string | A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `accountId` is specified. (optional)
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	accountId := "accountId_example" // string | A query string that is matched exactly against user `accountId`. Required, unless `query` is specified. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindBulkAssignableUsers(context.Background()).ProjectKeys(projectKeys).Query(query).Username(username).AccountId(accountId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindBulkAssignableUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindBulkAssignableUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindBulkAssignableUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindBulkAssignableUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectKeys** | **string** | A list of project keys (case sensitive). This parameter accepts a comma-separated list. | 
 **query** | **string** | A query string that is matched against user attributes, such as &#x60;displayName&#x60; and &#x60;emailAddress&#x60;, to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. Required, unless &#x60;accountId&#x60; is specified. | 
 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **accountId** | **string** | A query string that is matched exactly against user &#x60;accountId&#x60;. Required, unless &#x60;query&#x60; is specified. | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUserKeysByQuery

> PageBeanUserKey FindUserKeysByQuery(ctx).Query(query).StartAt(startAt).MaxResult(maxResult).Execute()

Find user keys by query



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
	query := "query_example" // string | The search query.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResult := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUserKeysByQuery(context.Background()).Query(query).StartAt(startAt).MaxResult(maxResult).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUserKeysByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUserKeysByQuery`: PageBeanUserKey
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUserKeysByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUserKeysByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResult** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanUserKey**](PageBeanUserKey.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsers

> []User FindUsers(ctx).Query(query).Username(username).AccountId(accountId).StartAt(startAt).MaxResults(maxResults).Property(property).Execute()

Find users



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
	query := "query" // string | A query string that is matched against user attributes ( `displayName`, and `emailAddress`) to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `accountId` or `property` is specified. (optional)
	username := "username_example" // string |  (optional)
	accountId := "accountId_example" // string | A query string that is matched exactly against a user `accountId`. Required, unless `query` or `property` is specified. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of filtered results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	property := "property_example" // string | A query string used to search properties. Property keys are specified by path, so property keys containing dot (.) or equals (=) characters cannot be used. The query string cannot be specified using a JSON object. Example: To search for the value of `nested` from `{\"something\":{\"nested\":1,\"other\":2}}` use `thepropertykey.something.nested=1`. Required, unless `accountId` or `query` is specified. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUsers(context.Background()).Query(query).Username(username).AccountId(accountId).StartAt(startAt).MaxResults(maxResults).Property(property).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUsers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsers`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUsers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A query string that is matched against user attributes ( &#x60;displayName&#x60;, and &#x60;emailAddress&#x60;) to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. Required, unless &#x60;accountId&#x60; or &#x60;property&#x60; is specified. | 
 **username** | **string** |  | 
 **accountId** | **string** | A query string that is matched exactly against a user &#x60;accountId&#x60;. Required, unless &#x60;query&#x60; or &#x60;property&#x60; is specified. | 
 **startAt** | **int32** | The index of the first item to return in a page of filtered results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **property** | **string** | A query string used to search properties. Property keys are specified by path, so property keys containing dot (.) or equals (&#x3D;) characters cannot be used. The query string cannot be specified using a JSON object. Example: To search for the value of &#x60;nested&#x60; from &#x60;{\&quot;something\&quot;:{\&quot;nested\&quot;:1,\&quot;other\&quot;:2}}&#x60; use &#x60;thepropertykey.something.nested&#x3D;1&#x60;. Required, unless &#x60;accountId&#x60; or &#x60;query&#x60; is specified. | 

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersByQuery

> PageBeanUser FindUsersByQuery(ctx).Query(query).StartAt(startAt).MaxResults(maxResults).Execute()

Find users by query



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
	query := "query_example" // string | The search query.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUsersByQuery(context.Background()).Query(query).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUsersByQuery``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersByQuery`: PageBeanUser
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUsersByQuery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersByQueryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | The search query. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanUser**](PageBeanUser.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersForPicker

> FoundUsers FindUsersForPicker(ctx).Query(query).MaxResults(maxResults).ShowAvatar(showAvatar).Exclude(exclude).ExcludeAccountIds(excludeAccountIds).AvatarSize(avatarSize).ExcludeConnectUsers(excludeConnectUsers).Execute()

Find users for picker



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
	query := "query_example" // string | A query string that is matched against user attributes, such as `displayName`, and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*.
	maxResults := int32(56) // int32 | The maximum number of items to return. The total number of matched users is returned in `total`. (optional) (default to 50)
	showAvatar := true // bool | Include the URI to the user's avatar. (optional) (default to false)
	exclude := []string{"Inner_example"} // []string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	excludeAccountIds := []string{"Inner_example"} // []string | A list of account IDs to exclude from the search results. This parameter accepts a comma-separated list. Multiple account IDs can also be provided using an ampersand-separated list. For example, `excludeAccountIds=5b10a2844c20165700ede21g,5b10a0effa615349cb016cd8&excludeAccountIds=5b10ac8d82e05b22cc7d4ef5`. Cannot be provided with `exclude`. (optional)
	avatarSize := "avatarSize_example" // string |  (optional)
	excludeConnectUsers := true // bool |  (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUsersForPicker(context.Background()).Query(query).MaxResults(maxResults).ShowAvatar(showAvatar).Exclude(exclude).ExcludeAccountIds(excludeAccountIds).AvatarSize(avatarSize).ExcludeConnectUsers(excludeConnectUsers).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUsersForPicker``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersForPicker`: FoundUsers
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUsersForPicker`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersForPickerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A query string that is matched against user attributes, such as &#x60;displayName&#x60;, and &#x60;emailAddress&#x60;, to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. | 
 **maxResults** | **int32** | The maximum number of items to return. The total number of matched users is returned in &#x60;total&#x60;. | [default to 50]
 **showAvatar** | **bool** | Include the URI to the user&#39;s avatar. | [default to false]
 **exclude** | **[]string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **excludeAccountIds** | **[]string** | A list of account IDs to exclude from the search results. This parameter accepts a comma-separated list. Multiple account IDs can also be provided using an ampersand-separated list. For example, &#x60;excludeAccountIds&#x3D;5b10a2844c20165700ede21g,5b10a0effa615349cb016cd8&amp;excludeAccountIds&#x3D;5b10ac8d82e05b22cc7d4ef5&#x60;. Cannot be provided with &#x60;exclude&#x60;. | 
 **avatarSize** | **string** |  | 
 **excludeConnectUsers** | **bool** |  | [default to false]

### Return type

[**FoundUsers**](FoundUsers.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersWithAllPermissions

> []User FindUsersWithAllPermissions(ctx).Permissions(permissions).Query(query).Username(username).AccountId(accountId).IssueKey(issueKey).ProjectKey(projectKey).StartAt(startAt).MaxResults(maxResults).Execute()

Find users with permissions



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
	permissions := "permissions_example" // string | A comma separated list of permissions. Permissions can be specified as any:   *  permission returned by [Get all permissions](#api-rest-api-3-permissions-get).  *  custom project permission added by Connect apps.  *  (deprecated) one of the following:           *  ASSIGNABLE\\_USER      *  ASSIGN\\_ISSUE      *  ATTACHMENT\\_DELETE\\_ALL      *  ATTACHMENT\\_DELETE\\_OWN      *  BROWSE      *  CLOSE\\_ISSUE      *  COMMENT\\_DELETE\\_ALL      *  COMMENT\\_DELETE\\_OWN      *  COMMENT\\_EDIT\\_ALL      *  COMMENT\\_EDIT\\_OWN      *  COMMENT\\_ISSUE      *  CREATE\\_ATTACHMENT      *  CREATE\\_ISSUE      *  DELETE\\_ISSUE      *  EDIT\\_ISSUE      *  LINK\\_ISSUE      *  MANAGE\\_WATCHER\\_LIST      *  MODIFY\\_REPORTER      *  MOVE\\_ISSUE      *  PROJECT\\_ADMIN      *  RESOLVE\\_ISSUE      *  SCHEDULE\\_ISSUE      *  SET\\_ISSUE\\_SECURITY      *  TRANSITION\\_ISSUE      *  VIEW\\_VERSION\\_CONTROL      *  VIEW\\_VOTERS\\_AND\\_WATCHERS      *  VIEW\\_WORKFLOW\\_READONLY      *  WORKLOG\\_DELETE\\_ALL      *  WORKLOG\\_DELETE\\_OWN      *  WORKLOG\\_EDIT\\_ALL      *  WORKLOG\\_EDIT\\_OWN      *  WORK\\_ISSUE
	query := "query" // string | A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `accountId` is specified. (optional)
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	accountId := "accountId_example" // string | A query string that is matched exactly against user `accountId`. Required, unless `query` is specified. (optional)
	issueKey := "issueKey_example" // string | The issue key for the issue. (optional)
	projectKey := "projectKey_example" // string | The project key for the project (case sensitive). (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUsersWithAllPermissions(context.Background()).Permissions(permissions).Query(query).Username(username).AccountId(accountId).IssueKey(issueKey).ProjectKey(projectKey).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUsersWithAllPermissions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersWithAllPermissions`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUsersWithAllPermissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersWithAllPermissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **permissions** | **string** | A comma separated list of permissions. Permissions can be specified as any:   *  permission returned by [Get all permissions](#api-rest-api-3-permissions-get).  *  custom project permission added by Connect apps.  *  (deprecated) one of the following:           *  ASSIGNABLE\\_USER      *  ASSIGN\\_ISSUE      *  ATTACHMENT\\_DELETE\\_ALL      *  ATTACHMENT\\_DELETE\\_OWN      *  BROWSE      *  CLOSE\\_ISSUE      *  COMMENT\\_DELETE\\_ALL      *  COMMENT\\_DELETE\\_OWN      *  COMMENT\\_EDIT\\_ALL      *  COMMENT\\_EDIT\\_OWN      *  COMMENT\\_ISSUE      *  CREATE\\_ATTACHMENT      *  CREATE\\_ISSUE      *  DELETE\\_ISSUE      *  EDIT\\_ISSUE      *  LINK\\_ISSUE      *  MANAGE\\_WATCHER\\_LIST      *  MODIFY\\_REPORTER      *  MOVE\\_ISSUE      *  PROJECT\\_ADMIN      *  RESOLVE\\_ISSUE      *  SCHEDULE\\_ISSUE      *  SET\\_ISSUE\\_SECURITY      *  TRANSITION\\_ISSUE      *  VIEW\\_VERSION\\_CONTROL      *  VIEW\\_VOTERS\\_AND\\_WATCHERS      *  VIEW\\_WORKFLOW\\_READONLY      *  WORKLOG\\_DELETE\\_ALL      *  WORKLOG\\_DELETE\\_OWN      *  WORKLOG\\_EDIT\\_ALL      *  WORKLOG\\_EDIT\\_OWN      *  WORK\\_ISSUE | 
 **query** | **string** | A query string that is matched against user attributes, such as &#x60;displayName&#x60; and &#x60;emailAddress&#x60;, to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. Required, unless &#x60;accountId&#x60; is specified. | 
 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **accountId** | **string** | A query string that is matched exactly against user &#x60;accountId&#x60;. Required, unless &#x60;query&#x60; is specified. | 
 **issueKey** | **string** | The issue key for the issue. | 
 **projectKey** | **string** | The project key for the project (case sensitive). | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindUsersWithBrowsePermission

> []User FindUsersWithBrowsePermission(ctx).Query(query).Username(username).AccountId(accountId).IssueKey(issueKey).ProjectKey(projectKey).StartAt(startAt).MaxResults(maxResults).Execute()

Find users with browse permission



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
	query := "query" // string | A query string that is matched against user attributes, such as `displayName` and `emailAddress`, to find relevant users. The string can match the prefix of the attribute's value. For example, *query=john* matches a user with a `displayName` of *John Smith* and a user with an `emailAddress` of *johnson@example.com*. Required, unless `accountId` is specified. (optional)
	username := "username_example" // string | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. (optional)
	accountId := "accountId_example" // string | A query string that is matched exactly against user `accountId`. Required, unless `query` is specified. (optional)
	issueKey := "issueKey_example" // string | The issue key for the issue. Required, unless `projectKey` is specified. (optional)
	projectKey := "projectKey_example" // string | The project key for the project (case sensitive). Required, unless `issueKey` is specified. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UserSearchAPI.FindUsersWithBrowsePermission(context.Background()).Query(query).Username(username).AccountId(accountId).IssueKey(issueKey).ProjectKey(projectKey).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UserSearchAPI.FindUsersWithBrowsePermission``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `FindUsersWithBrowsePermission`: []User
	fmt.Fprintf(os.Stdout, "Response from `UserSearchAPI.FindUsersWithBrowsePermission`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiFindUsersWithBrowsePermissionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A query string that is matched against user attributes, such as &#x60;displayName&#x60; and &#x60;emailAddress&#x60;, to find relevant users. The string can match the prefix of the attribute&#39;s value. For example, *query&#x3D;john* matches a user with a &#x60;displayName&#x60; of *John Smith* and a user with an &#x60;emailAddress&#x60; of *johnson@example.com*. Required, unless &#x60;accountId&#x60; is specified. | 
 **username** | **string** | This parameter is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | 
 **accountId** | **string** | A query string that is matched exactly against user &#x60;accountId&#x60;. Required, unless &#x60;query&#x60; is specified. | 
 **issueKey** | **string** | The issue key for the issue. Required, unless &#x60;projectKey&#x60; is specified. | 
 **projectKey** | **string** | The project key for the project (case sensitive). Required, unless &#x60;issueKey&#x60; is specified. | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**[]User**](User.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

