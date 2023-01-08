# \ProjectVersionsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRelatedWork**](ProjectVersionsAPI.md#CreateRelatedWork) | **Post** /rest/api/3/version/{id}/relatedwork | Create related work
[**CreateVersion**](ProjectVersionsAPI.md#CreateVersion) | **Post** /rest/api/3/version | Create version
[**DeleteAndReplaceVersion**](ProjectVersionsAPI.md#DeleteAndReplaceVersion) | **Post** /rest/api/3/version/{id}/removeAndSwap | Delete and replace version
[**DeleteRelatedWork**](ProjectVersionsAPI.md#DeleteRelatedWork) | **Delete** /rest/api/3/version/{versionId}/relatedwork/{relatedWorkId} | Delete related work
[**DeleteVersion**](ProjectVersionsAPI.md#DeleteVersion) | **Delete** /rest/api/3/version/{id} | Delete version
[**GetProjectVersions**](ProjectVersionsAPI.md#GetProjectVersions) | **Get** /rest/api/3/project/{projectIdOrKey}/versions | Get project versions
[**GetProjectVersionsPaginated**](ProjectVersionsAPI.md#GetProjectVersionsPaginated) | **Get** /rest/api/3/project/{projectIdOrKey}/version | Get project versions paginated
[**GetRelatedWork**](ProjectVersionsAPI.md#GetRelatedWork) | **Get** /rest/api/3/version/{id}/relatedwork | Get related work
[**GetVersion**](ProjectVersionsAPI.md#GetVersion) | **Get** /rest/api/3/version/{id} | Get version
[**GetVersionRelatedIssues**](ProjectVersionsAPI.md#GetVersionRelatedIssues) | **Get** /rest/api/3/version/{id}/relatedIssueCounts | Get version&#39;s related issues count
[**GetVersionUnresolvedIssues**](ProjectVersionsAPI.md#GetVersionUnresolvedIssues) | **Get** /rest/api/3/version/{id}/unresolvedIssueCount | Get version&#39;s unresolved issues count
[**MergeVersions**](ProjectVersionsAPI.md#MergeVersions) | **Put** /rest/api/3/version/{id}/mergeto/{moveIssuesTo} | Merge versions
[**MoveVersion**](ProjectVersionsAPI.md#MoveVersion) | **Post** /rest/api/3/version/{id}/move | Move version
[**UpdateRelatedWork**](ProjectVersionsAPI.md#UpdateRelatedWork) | **Put** /rest/api/3/version/{id}/relatedwork | Update related work
[**UpdateVersion**](ProjectVersionsAPI.md#UpdateVersion) | **Put** /rest/api/3/version/{id} | Update version



## CreateRelatedWork

> VersionRelatedWork CreateRelatedWork(ctx, id).VersionRelatedWork(versionRelatedWork).Execute()

Create related work



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
	id := "id_example" // string | 
	versionRelatedWork := *openapiclient.NewVersionRelatedWork("Category_example") // VersionRelatedWork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.CreateRelatedWork(context.Background(), id).VersionRelatedWork(versionRelatedWork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.CreateRelatedWork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateRelatedWork`: VersionRelatedWork
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.CreateRelatedWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRelatedWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionRelatedWork** | [**VersionRelatedWork**](VersionRelatedWork.md) |  | 

### Return type

[**VersionRelatedWork**](VersionRelatedWork.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateVersion

> Version CreateVersion(ctx).Version(version).Execute()

Create version



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
	version := *openapiclient.NewVersion() // Version | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.CreateVersion(context.Background()).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.CreateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateVersion`: Version
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.CreateVersion`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **version** | [**Version**](Version.md) |  | 

### Return type

[**Version**](Version.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAndReplaceVersion

> interface{} DeleteAndReplaceVersion(ctx, id).DeleteAndReplaceVersionBean(deleteAndReplaceVersionBean).Execute()

Delete and replace version



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
	id := "id_example" // string | The ID of the version.
	deleteAndReplaceVersionBean := *openapiclient.NewDeleteAndReplaceVersionBean() // DeleteAndReplaceVersionBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.DeleteAndReplaceVersion(context.Background(), id).DeleteAndReplaceVersionBean(deleteAndReplaceVersionBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.DeleteAndReplaceVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteAndReplaceVersion`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.DeleteAndReplaceVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAndReplaceVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deleteAndReplaceVersionBean** | [**DeleteAndReplaceVersionBean**](DeleteAndReplaceVersionBean.md) |  | 

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


## DeleteRelatedWork

> DeleteRelatedWork(ctx, versionId, relatedWorkId).Execute()

Delete related work



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
	versionId := "versionId_example" // string | The ID of the version that the target related work belongs to.
	relatedWorkId := "relatedWorkId_example" // string | The ID of the related work to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectVersionsAPI.DeleteRelatedWork(context.Background(), versionId, relatedWorkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.DeleteRelatedWork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**versionId** | **string** | The ID of the version that the target related work belongs to. | 
**relatedWorkId** | **string** | The ID of the related work to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRelatedWorkRequest struct via the builder pattern


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


## DeleteVersion

> DeleteVersion(ctx, id).MoveFixIssuesTo(moveFixIssuesTo).MoveAffectedIssuesTo(moveAffectedIssuesTo).Execute()

Delete version



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
	id := "id_example" // string | The ID of the version.
	moveFixIssuesTo := "moveFixIssuesTo_example" // string | The ID of the version to update `fixVersion` to when the field contains the deleted version. The replacement version must be in the same project as the version being deleted and cannot be the version being deleted. (optional)
	moveAffectedIssuesTo := "moveAffectedIssuesTo_example" // string | The ID of the version to update `affectedVersion` to when the field contains the deleted version. The replacement version must be in the same project as the version being deleted and cannot be the version being deleted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ProjectVersionsAPI.DeleteVersion(context.Background(), id).MoveFixIssuesTo(moveFixIssuesTo).MoveAffectedIssuesTo(moveAffectedIssuesTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.DeleteVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveFixIssuesTo** | **string** | The ID of the version to update &#x60;fixVersion&#x60; to when the field contains the deleted version. The replacement version must be in the same project as the version being deleted and cannot be the version being deleted. | 
 **moveAffectedIssuesTo** | **string** | The ID of the version to update &#x60;affectedVersion&#x60; to when the field contains the deleted version. The replacement version must be in the same project as the version being deleted and cannot be the version being deleted. | 

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


## GetProjectVersions

> []Version GetProjectVersions(ctx, projectIdOrKey).Expand(expand).Execute()

Get project versions



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts `operations`, which returns actions that can be performed on the version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetProjectVersions(context.Background(), projectIdOrKey).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetProjectVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectVersions`: []Version
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetProjectVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;operations&#x60;, which returns actions that can be performed on the version. | 

### Return type

[**[]Version**](Version.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectVersionsPaginated

> PageBeanVersion GetProjectVersionsPaginated(ctx, projectIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Query(query).Status(status).Expand(expand).Execute()

Get project versions paginated



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
	projectIdOrKey := "projectIdOrKey_example" // string | The project ID or project key (case sensitive).
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `description` Sorts by version description.  *  `name` Sorts by version name.  *  `releaseDate` Sorts by release date, starting with the oldest date. Versions with no release date are listed last.  *  `sequence` Sorts by the order of appearance in the user interface.  *  `startDate` Sorts by start date, starting with the oldest date. Versions with no start date are listed last. (optional)
	query := "query_example" // string | Filter the results using a literal string. Versions with matching `name` or `description` are returned (case insensitive). (optional)
	status := "status_example" // string | A list of status values used to filter the results by version status. This parameter accepts a comma-separated list. The status values are `released`, `unreleased`, and `archived`. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `issuesstatus` Returns the number of issues in each status category for each version.  *  `operations` Returns actions that can be performed on the specified version.  *  `driver` Returns the Atlassian account ID of the version driver.  *  `approvers` Returns a list containing the approvers for this version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetProjectVersionsPaginated(context.Background(), projectIdOrKey).StartAt(startAt).MaxResults(maxResults).OrderBy(orderBy).Query(query).Status(status).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetProjectVersionsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectVersionsPaginated`: PageBeanVersion
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetProjectVersionsPaginated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**projectIdOrKey** | **string** | The project ID or project key (case sensitive). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectVersionsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;description&#x60; Sorts by version description.  *  &#x60;name&#x60; Sorts by version name.  *  &#x60;releaseDate&#x60; Sorts by release date, starting with the oldest date. Versions with no release date are listed last.  *  &#x60;sequence&#x60; Sorts by the order of appearance in the user interface.  *  &#x60;startDate&#x60; Sorts by start date, starting with the oldest date. Versions with no start date are listed last. | 
 **query** | **string** | Filter the results using a literal string. Versions with matching &#x60;name&#x60; or &#x60;description&#x60; are returned (case insensitive). | 
 **status** | **string** | A list of status values used to filter the results by version status. This parameter accepts a comma-separated list. The status values are &#x60;released&#x60;, &#x60;unreleased&#x60;, and &#x60;archived&#x60;. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;issuesstatus&#x60; Returns the number of issues in each status category for each version.  *  &#x60;operations&#x60; Returns actions that can be performed on the specified version.  *  &#x60;driver&#x60; Returns the Atlassian account ID of the version driver.  *  &#x60;approvers&#x60; Returns a list containing the approvers for this version. | 

### Return type

[**PageBeanVersion**](PageBeanVersion.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRelatedWork

> []VersionRelatedWork GetRelatedWork(ctx, id).Execute()

Get related work



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
	id := "id_example" // string | The ID of the version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetRelatedWork(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetRelatedWork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRelatedWork`: []VersionRelatedWork
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetRelatedWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRelatedWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]VersionRelatedWork**](VersionRelatedWork.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersion

> Version GetVersion(ctx, id).Expand(expand).Execute()

Get version



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
	id := "id_example" // string | The ID of the version.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about version in the response. This parameter accepts a comma-separated list. Expand options include:   *  `operations` Returns the list of operations available for this version.  *  `issuesstatus` Returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property represents the number of issues with a status other than *to do*, *in progress*, and *done*.  *  `driver` Returns the Atlassian account ID of the version driver.  *  `approvers` Returns a list containing the Atlassian account IDs of approvers for this version. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetVersion(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersion`: Version
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use [expand](#expansion) to include additional information about version in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;operations&#x60; Returns the list of operations available for this version.  *  &#x60;issuesstatus&#x60; Returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property represents the number of issues with a status other than *to do*, *in progress*, and *done*.  *  &#x60;driver&#x60; Returns the Atlassian account ID of the version driver.  *  &#x60;approvers&#x60; Returns a list containing the Atlassian account IDs of approvers for this version. | 

### Return type

[**Version**](Version.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionRelatedIssues

> VersionIssueCounts GetVersionRelatedIssues(ctx, id).Execute()

Get version's related issues count



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
	id := "id_example" // string | The ID of the version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetVersionRelatedIssues(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetVersionRelatedIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionRelatedIssues`: VersionIssueCounts
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetVersionRelatedIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionRelatedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionIssueCounts**](VersionIssueCounts.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVersionUnresolvedIssues

> VersionUnresolvedIssuesCount GetVersionUnresolvedIssues(ctx, id).Execute()

Get version's unresolved issues count



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
	id := "id_example" // string | The ID of the version.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.GetVersionUnresolvedIssues(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.GetVersionUnresolvedIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVersionUnresolvedIssues`: VersionUnresolvedIssuesCount
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.GetVersionUnresolvedIssues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVersionUnresolvedIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**VersionUnresolvedIssuesCount**](VersionUnresolvedIssuesCount.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MergeVersions

> interface{} MergeVersions(ctx, id, moveIssuesTo).Execute()

Merge versions



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
	id := "id_example" // string | The ID of the version to delete.
	moveIssuesTo := "moveIssuesTo_example" // string | The ID of the version to merge into.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.MergeVersions(context.Background(), id, moveIssuesTo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.MergeVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MergeVersions`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.MergeVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version to delete. | 
**moveIssuesTo** | **string** | The ID of the version to merge into. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMergeVersionsRequest struct via the builder pattern


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


## MoveVersion

> Version MoveVersion(ctx, id).VersionMoveBean(versionMoveBean).Execute()

Move version



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
	id := "id_example" // string | The ID of the version to be moved.
	versionMoveBean := *openapiclient.NewVersionMoveBean() // VersionMoveBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.MoveVersion(context.Background(), id).VersionMoveBean(versionMoveBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.MoveVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MoveVersion`: Version
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.MoveVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version to be moved. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionMoveBean** | [**VersionMoveBean**](VersionMoveBean.md) |  | 

### Return type

[**Version**](Version.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRelatedWork

> VersionRelatedWork UpdateRelatedWork(ctx, id).VersionRelatedWork(versionRelatedWork).Execute()

Update related work



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
	id := "id_example" // string | The ID of the version to update the related work on. For the related work id, pass it to the input JSON.
	versionRelatedWork := *openapiclient.NewVersionRelatedWork("Category_example") // VersionRelatedWork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.UpdateRelatedWork(context.Background(), id).VersionRelatedWork(versionRelatedWork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.UpdateRelatedWork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRelatedWork`: VersionRelatedWork
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.UpdateRelatedWork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version to update the related work on. For the related work id, pass it to the input JSON. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRelatedWorkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **versionRelatedWork** | [**VersionRelatedWork**](VersionRelatedWork.md) |  | 

### Return type

[**VersionRelatedWork**](VersionRelatedWork.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateVersion

> Version UpdateVersion(ctx, id).Version(version).Execute()

Update version



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
	id := "id_example" // string | The ID of the version.
	version := *openapiclient.NewVersion() // Version | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ProjectVersionsAPI.UpdateVersion(context.Background(), id).Version(version).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ProjectVersionsAPI.UpdateVersion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateVersion`: Version
	fmt.Fprintf(os.Stdout, "Response from `ProjectVersionsAPI.UpdateVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the version. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **version** | [**Version**](Version.md) |  | 

### Return type

[**Version**](Version.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

