# \BoardAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBoard**](BoardAPI.md#CreateBoard) | **Post** /rest/agile/1.0/board | Create board
[**DeleteBoard**](BoardAPI.md#DeleteBoard) | **Delete** /rest/agile/1.0/board/{boardId} | Delete board
[**DeleteBoardProperty**](BoardAPI.md#DeleteBoardProperty) | **Delete** /rest/agile/1.0/board/{boardId}/properties/{propertyKey} | Delete board property
[**GetAllBoards**](BoardAPI.md#GetAllBoards) | **Get** /rest/agile/1.0/board | Get all boards
[**GetAllQuickFilters**](BoardAPI.md#GetAllQuickFilters) | **Get** /rest/agile/1.0/board/{boardId}/quickfilter | Get all quick filters
[**GetAllSprints**](BoardAPI.md#GetAllSprints) | **Get** /rest/agile/1.0/board/{boardId}/sprint | Get all sprints
[**GetAllVersions**](BoardAPI.md#GetAllVersions) | **Get** /rest/agile/1.0/board/{boardId}/version | Get all versions
[**GetBoard**](BoardAPI.md#GetBoard) | **Get** /rest/agile/1.0/board/{boardId} | Get board
[**GetBoardByFilterId**](BoardAPI.md#GetBoardByFilterId) | **Get** /rest/agile/1.0/board/filter/{filterId} | Get board by filter id
[**GetBoardIssuesForEpic**](BoardAPI.md#GetBoardIssuesForEpic) | **Get** /rest/agile/1.0/board/{boardId}/epic/{epicId}/issue | Get board issues for epic
[**GetBoardIssuesForSprint**](BoardAPI.md#GetBoardIssuesForSprint) | **Get** /rest/agile/1.0/board/{boardId}/sprint/{sprintId}/issue | Get board issues for sprint
[**GetBoardProperty**](BoardAPI.md#GetBoardProperty) | **Get** /rest/agile/1.0/board/{boardId}/properties/{propertyKey} | Get board property
[**GetBoardPropertyKeys**](BoardAPI.md#GetBoardPropertyKeys) | **Get** /rest/agile/1.0/board/{boardId}/properties | Get board property keys
[**GetConfiguration**](BoardAPI.md#GetConfiguration) | **Get** /rest/agile/1.0/board/{boardId}/configuration | Get configuration
[**GetEpics**](BoardAPI.md#GetEpics) | **Get** /rest/agile/1.0/board/{boardId}/epic | Get epics
[**GetFeaturesForBoard**](BoardAPI.md#GetFeaturesForBoard) | **Get** /rest/agile/1.0/board/{boardId}/features | Get features for board
[**GetIssuesForBacklog**](BoardAPI.md#GetIssuesForBacklog) | **Get** /rest/agile/1.0/board/{boardId}/backlog | Get issues for backlog
[**GetIssuesForBoard**](BoardAPI.md#GetIssuesForBoard) | **Get** /rest/agile/1.0/board/{boardId}/issue | Get issues for board
[**GetIssuesWithoutEpicForBoard**](BoardAPI.md#GetIssuesWithoutEpicForBoard) | **Get** /rest/agile/1.0/board/{boardId}/epic/none/issue | Get issues without epic for board
[**GetProjects**](BoardAPI.md#GetProjects) | **Get** /rest/agile/1.0/board/{boardId}/project | Get projects
[**GetProjectsFull**](BoardAPI.md#GetProjectsFull) | **Get** /rest/agile/1.0/board/{boardId}/project/full | Get projects full
[**GetQuickFilter**](BoardAPI.md#GetQuickFilter) | **Get** /rest/agile/1.0/board/{boardId}/quickfilter/{quickFilterId} | Get quick filter
[**GetReportsForBoard**](BoardAPI.md#GetReportsForBoard) | **Get** /rest/agile/1.0/board/{boardId}/reports | Get reports for board
[**MoveIssuesToBoard**](BoardAPI.md#MoveIssuesToBoard) | **Post** /rest/agile/1.0/board/{boardId}/issue | Move issues to board
[**SetBoardProperty**](BoardAPI.md#SetBoardProperty) | **Put** /rest/agile/1.0/board/{boardId}/properties/{propertyKey} | Set board property
[**ToggleFeatures**](BoardAPI.md#ToggleFeatures) | **Put** /rest/agile/1.0/board/{boardId}/features | Toggle features



## CreateBoard

> GetAllBoards200ResponseValuesInner CreateBoard(ctx).CreateBoardRequest(createBoardRequest).Execute()

Create board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	createBoardRequest := *openapiclient.NewCreateBoardRequest() // CreateBoardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.CreateBoard(context.Background()).CreateBoardRequest(createBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.CreateBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateBoard`: GetAllBoards200ResponseValuesInner
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.CreateBoard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBoardRequest** | [**CreateBoardRequest**](CreateBoardRequest.md) |  | 

### Return type

[**GetAllBoards200ResponseValuesInner**](GetAllBoards200ResponseValuesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBoard

> DeleteBoard(ctx, boardId).Execute()

Delete board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | ID of the board to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.DeleteBoard(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.DeleteBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | ID of the board to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBoardProperty

> DeleteBoardProperty(ctx, boardId, propertyKey).Execute()

Delete board property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := "boardId_example" // string | the id of the board from which the property will be removed.
	propertyKey := "propertyKey_example" // string | the key of the property to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.DeleteBoardProperty(context.Background(), boardId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.DeleteBoardProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | the id of the board from which the property will be removed. | 
**propertyKey** | **string** | the key of the property to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBoardPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllBoards

> GetAllBoards200Response GetAllBoards(ctx).StartAt(startAt).MaxResults(maxResults).Type_(type_).Name(name).ProjectKeyOrId(projectKeyOrId).AccountIdLocation(accountIdLocation).ProjectLocation(projectLocation).IncludePrivate(includePrivate).NegateLocationFiltering(negateLocationFiltering).OrderBy(orderBy).Expand(expand).ProjectTypeLocation(projectTypeLocation).FilterId(filterId).Execute()

Get all boards



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	startAt := int64(789) // int64 | The starting index of the returned boards. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of boards to return per page. See the 'Pagination' section at the top of this page for more details. (optional) (default to 50)
	type_ := map[string]interface{}{ ... } // map[string]interface{} | Filters results to boards of the specified types. Valid values: scrum, kanban, simple. (optional)
	name := "name_example" // string | Filters results to boards that match or partially match the specified name. (optional)
	projectKeyOrId := "projectKeyOrId_example" // string | Filters results to boards that are relevant to a project. Relevance means that the jql filter defined in board contains a reference to a project. (optional)
	accountIdLocation := "accountIdLocation_example" // string |  (optional)
	projectLocation := "projectLocation_example" // string |  (optional)
	includePrivate := true // bool | Appends private boards to the end of the list. The name and type fields are excluded for security reasons. (optional)
	negateLocationFiltering := true // bool | If set to true, negate filters used for querying by location. By default false. (optional)
	orderBy := "orderBy_example" // string | Ordering of the results by a given field. If not provided, values will not be sorted. Valid values: name. (optional)
	expand := "expand_example" // string | List of fields to expand for each board. Valid values: admins, permissions. (optional)
	projectTypeLocation := []string{"Inner_example"} // []string | Filters results to boards that are relevant to a project types. Support Jira Software, Jira Service Management. Valid values: software, service\\_desk. By default software. (optional)
	filterId := int64(789) // int64 | Filters results to boards that are relevant to a filter. Not supported for next-gen boards. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetAllBoards(context.Background()).StartAt(startAt).MaxResults(maxResults).Type_(type_).Name(name).ProjectKeyOrId(projectKeyOrId).AccountIdLocation(accountIdLocation).ProjectLocation(projectLocation).IncludePrivate(includePrivate).NegateLocationFiltering(negateLocationFiltering).OrderBy(orderBy).Expand(expand).ProjectTypeLocation(projectTypeLocation).FilterId(filterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetAllBoards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllBoards`: GetAllBoards200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetAllBoards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllBoardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The starting index of the returned boards. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | [default to 0]
 **maxResults** | **int32** | The maximum number of boards to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | [default to 50]
 **type_** | [**map[string]interface{}**](map[string]interface{}.md) | Filters results to boards of the specified types. Valid values: scrum, kanban, simple. | 
 **name** | **string** | Filters results to boards that match or partially match the specified name. | 
 **projectKeyOrId** | **string** | Filters results to boards that are relevant to a project. Relevance means that the jql filter defined in board contains a reference to a project. | 
 **accountIdLocation** | **string** |  | 
 **projectLocation** | **string** |  | 
 **includePrivate** | **bool** | Appends private boards to the end of the list. The name and type fields are excluded for security reasons. | 
 **negateLocationFiltering** | **bool** | If set to true, negate filters used for querying by location. By default false. | 
 **orderBy** | **string** | Ordering of the results by a given field. If not provided, values will not be sorted. Valid values: name. | 
 **expand** | **string** | List of fields to expand for each board. Valid values: admins, permissions. | 
 **projectTypeLocation** | **[]string** | Filters results to boards that are relevant to a project types. Support Jira Software, Jira Service Management. Valid values: software, service\\_desk. By default software. | 
 **filterId** | **int64** | Filters results to boards that are relevant to a filter. Not supported for next-gen boards. | 

### Return type

[**GetAllBoards200Response**](GetAllBoards200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllQuickFilters

> GetAllQuickFilters200Response GetAllQuickFilters(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Execute()

Get all quick filters



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested quick filters.
	startAt := int64(789) // int64 | The starting index of the returned quick filters. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of sprints to return per page. See the 'Pagination' section at the top of this page for more details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetAllQuickFilters(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetAllQuickFilters``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllQuickFilters`: GetAllQuickFilters200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetAllQuickFilters`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested quick filters. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllQuickFiltersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned quick filters. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of sprints to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | 

### Return type

[**GetAllQuickFilters200Response**](GetAllQuickFilters200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllSprints

> GetAllSprints(ctx, boardId).StartAt(startAt).MaxResults(maxResults).State(state).Execute()

Get all sprints



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested sprints.
	startAt := int64(789) // int64 | The starting index of the returned sprints. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of sprints to return per page. See the 'Pagination' section at the top of this page for more details. (optional)
	state := map[string]interface{}{ ... } // map[string]interface{} | Filters results to sprints in specified states. Valid values: future, active, closed. You can define multiple states separated by commas, e.g. state=active,closed (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetAllSprints(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).State(state).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetAllSprints``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested sprints. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllSprintsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned sprints. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of sprints to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **state** | [**map[string]interface{}**](map[string]interface{}.md) | Filters results to sprints in specified states. Valid values: future, active, closed. You can define multiple states separated by commas, e.g. state&#x3D;active,closed | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllVersions

> GetAllVersions(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Released(released).Execute()

Get all versions



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested versions.
	startAt := int64(789) // int64 | The starting index of the returned versions. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of versions to return per page. See the 'Pagination' section at the top of this page for more details. (optional)
	released := "released_example" // string | Filters results to versions that are either released or unreleased. Valid values: true, false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetAllVersions(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Released(released).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetAllVersions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested versions. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned versions. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of versions to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **released** | **string** | Filters results to versions that are either released or unreleased. Valid values: true, false. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoard

> GetAllBoards200ResponseValuesInner GetBoard(ctx, boardId).Execute()

Get board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the requested board.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetBoard(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoard`: GetAllBoards200ResponseValuesInner
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the requested board. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetAllBoards200ResponseValuesInner**](GetAllBoards200ResponseValuesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardByFilterId

> GetBoardByFilterId200Response GetBoardByFilterId(ctx, filterId).StartAt(startAt).MaxResults(maxResults).Execute()

Get board by filter id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	filterId := int64(789) // int64 | Filters results to boards that are relevant to a filter. Not supported for next-gen boards.
	startAt := int64(789) // int64 | The starting index of the returned boards. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of boards to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetBoardByFilterId(context.Background(), filterId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoardByFilterId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBoardByFilterId`: GetBoardByFilterId200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetBoardByFilterId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**filterId** | **int64** | Filters results to boards that are relevant to a filter. Not supported for next-gen boards. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardByFilterIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned boards. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of boards to return per page. Default: 50. See the &#39;Pagination&#39; section at the top of this page for more details. | 

### Return type

[**GetBoardByFilterId200Response**](GetBoardByFilterId200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardIssuesForEpic

> GetBoardIssuesForEpic(ctx, boardId, epicId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get board issues for epic



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested issues.
	epicId := int64(789) // int64 | The ID of the epic that contains the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetBoardIssuesForEpic(context.Background(), boardId, epicId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoardIssuesForEpic``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested issues. | 
**epicId** | **int64** | The ID of the epic that contains the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardIssuesForEpicRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. Default: 50. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardIssuesForSprint

> GetBoardIssuesForSprint(ctx, boardId, sprintId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get board issues for sprint



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains requested issues.
	sprintId := int64(789) // int64 | The ID of the sprint that contains requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetBoardIssuesForSprint(context.Background(), boardId, sprintId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoardIssuesForSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains requested issues. | 
**sprintId** | **int64** | The ID of the sprint that contains requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardIssuesForSprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that &#x60;username&#x60; and &#x60;userkey&#x60; can&#39;t be used as search terms for this parameter due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardProperty

> GetBoardProperty(ctx, boardId, propertyKey).Execute()

Get board property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := "boardId_example" // string | the ID of the board from which the property will be returned.
	propertyKey := "propertyKey_example" // string | the key of the property to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetBoardProperty(context.Background(), boardId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoardProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | the ID of the board from which the property will be returned. | 
**propertyKey** | **string** | the key of the property to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBoardPropertyKeys

> GetBoardPropertyKeys(ctx, boardId).Execute()

Get board property keys



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := "boardId_example" // string | the ID of the board from which property keys will be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetBoardPropertyKeys(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetBoardPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | the ID of the board from which property keys will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBoardPropertyKeysRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> GetConfiguration200Response GetConfiguration(ctx, boardId).Execute()

Get configuration



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board for which configuration is requested.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetConfiguration(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: GetConfiguration200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board for which configuration is requested. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetConfiguration200Response**](GetConfiguration200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEpics

> GetEpics(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Done(done).Execute()

Get epics



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested epics.
	startAt := int64(789) // int64 | The starting index of the returned epics. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of epics to return per page. See the 'Pagination' section at the top of this page for more details. (optional)
	done := "done_example" // string | Filters results to epics that are either done or not done. Valid values: true, false. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetEpics(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Done(done).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetEpics``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested epics. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEpicsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned epics. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of epics to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **done** | **string** | Filters results to epics that are either done or not done. Valid values: true, false. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFeaturesForBoard

> GetFeaturesForBoard200Response GetFeaturesForBoard(ctx, boardId).Execute()

Get features for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetFeaturesForBoard(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetFeaturesForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeaturesForBoard`: GetFeaturesForBoard200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetFeaturesForBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeaturesForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetFeaturesForBoard200Response**](GetFeaturesForBoard200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesForBacklog

> SearchResults GetIssuesForBacklog(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues for backlog



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that has the backlog containing the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. Default: 50. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | This parameter is currently not used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetIssuesForBacklog(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetIssuesForBacklog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuesForBacklog`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetIssuesForBacklog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that has the backlog containing the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesForBacklogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. Default: 50. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that &#x60;username&#x60; and &#x60;userkey&#x60; can&#39;t be used as search terms for this parameter due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | This parameter is currently not used. | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesForBoard

> SearchResults GetIssuesForBoard(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | This parameter is currently not used. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetIssuesForBoard(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetIssuesForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuesForBoard`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetIssuesForBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that &#x60;username&#x60; and &#x60;userkey&#x60; can&#39;t be used as search terms for this parameter due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | This parameter is currently not used. | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuesWithoutEpicForBoard

> GetIssuesWithoutEpicForBoard(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues without epic for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetIssuesWithoutEpicForBoard(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetIssuesWithoutEpicForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesWithoutEpicForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned issues. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of issues to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. Note, the total number of issues returned is limited by the property &#39;jira.search.views.default.max&#39; in your Jira instance. If you exceed this limit, your results will be truncated. | 
 **jql** | **string** | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that &#x60;username&#x60; and &#x60;userkey&#x60; can&#39;t be used as search terms for this parameter due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **validateQuery** | **bool** | Specifies whether to validate the JQL query or not. Default: true. | 
 **fields** | **[]map[string]interface{}** | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. | 
 **expand** | **string** | A comma-separated list of the parameters to expand. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjects

> GetProjects(ctx, boardId).StartAt(startAt).MaxResults(maxResults).Execute()

Get projects



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains returned projects.
	startAt := int64(789) // int64 | The starting index of the returned projects. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of projects to return per page. See the 'Pagination' section at the top of this page for more details. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetProjects(context.Background(), boardId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains returned projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The starting index of the returned projects. Base index: 0. See the &#39;Pagination&#39; section at the top of this page for more details. | 
 **maxResults** | **int32** | The maximum number of projects to return per page. See the &#39;Pagination&#39; section at the top of this page for more details. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectsFull

> GetProjectsFull(ctx, boardId).Execute()

Get projects full



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | The ID of the board that contains returned projects.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.GetProjectsFull(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetProjectsFull``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** | The ID of the board that contains returned projects. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectsFullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetQuickFilter

> GetAllQuickFilters200ResponseValuesInner GetQuickFilter(ctx, boardId, quickFilterId).Execute()

Get quick filter



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | 
	quickFilterId := int64(789) // int64 | The ID of the requested quick filter.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetQuickFilter(context.Background(), boardId, quickFilterId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetQuickFilter``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetQuickFilter`: GetAllQuickFilters200ResponseValuesInner
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetQuickFilter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 
**quickFilterId** | **int64** | The ID of the requested quick filter. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetQuickFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetAllQuickFilters200ResponseValuesInner**](GetAllQuickFilters200ResponseValuesInner.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetReportsForBoard

> GetReportsForBoard200Response GetReportsForBoard(ctx, boardId).Execute()

Get reports for board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.GetReportsForBoard(context.Background(), boardId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.GetReportsForBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetReportsForBoard`: GetReportsForBoard200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.GetReportsForBoard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReportsForBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetReportsForBoard200Response**](GetReportsForBoard200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MoveIssuesToBoard

> MoveIssuesToBoard(ctx, boardId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()

Move issues to board



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | 
	moveIssuesToBacklogForBoardRequest := *openapiclient.NewMoveIssuesToBacklogForBoardRequest() // MoveIssuesToBacklogForBoardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.BoardAPI.MoveIssuesToBoard(context.Background(), boardId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.MoveIssuesToBoard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveIssuesToBoardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveIssuesToBacklogForBoardRequest** | [**MoveIssuesToBacklogForBoardRequest**](MoveIssuesToBacklogForBoardRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBoardProperty

> interface{} SetBoardProperty(ctx, boardId, propertyKey).Body(body).Execute()

Set board property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := "boardId_example" // string | the ID of the board on which the property will be set.
	propertyKey := "propertyKey_example" // string | the key of the board's property. The maximum length of the key is 255 bytes.
	body := interface{}(987) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.SetBoardProperty(context.Background(), boardId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.SetBoardProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBoardProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.SetBoardProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **string** | the ID of the board on which the property will be set. | 
**propertyKey** | **string** | the key of the board&#39;s property. The maximum length of the key is 255 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetBoardPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **interface{}** | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes. | 

### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ToggleFeatures

> GetFeaturesForBoard200Response ToggleFeatures(ctx, boardId).ToggleFeaturesRequest(toggleFeaturesRequest).Execute()

Toggle features



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	boardId := int64(789) // int64 | 
	toggleFeaturesRequest := *openapiclient.NewToggleFeaturesRequest() // ToggleFeaturesRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.BoardAPI.ToggleFeatures(context.Background(), boardId).ToggleFeaturesRequest(toggleFeaturesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `BoardAPI.ToggleFeatures``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ToggleFeatures`: GetFeaturesForBoard200Response
	fmt.Fprintf(os.Stdout, "Response from `BoardAPI.ToggleFeatures`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boardId** | **int64** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiToggleFeaturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **toggleFeaturesRequest** | [**ToggleFeaturesRequest**](ToggleFeaturesRequest.md) |  | 

### Return type

[**GetFeaturesForBoard200Response**](GetFeaturesForBoard200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

