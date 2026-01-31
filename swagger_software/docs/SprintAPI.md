# \SprintAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSprint**](SprintAPI.md#CreateSprint) | **Post** /rest/agile/1.0/sprint | Create sprint
[**DeleteProperty**](SprintAPI.md#DeleteProperty) | **Delete** /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey} | Delete property
[**DeleteSprint**](SprintAPI.md#DeleteSprint) | **Delete** /rest/agile/1.0/sprint/{sprintId} | Delete sprint
[**GetIssuesForSprint**](SprintAPI.md#GetIssuesForSprint) | **Get** /rest/agile/1.0/sprint/{sprintId}/issue | Get issues for sprint
[**GetPropertiesKeys**](SprintAPI.md#GetPropertiesKeys) | **Get** /rest/agile/1.0/sprint/{sprintId}/properties | Get properties keys
[**GetProperty**](SprintAPI.md#GetProperty) | **Get** /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey} | Get property
[**GetSprint**](SprintAPI.md#GetSprint) | **Get** /rest/agile/1.0/sprint/{sprintId} | Get sprint
[**MoveIssuesToSprintAndRank**](SprintAPI.md#MoveIssuesToSprintAndRank) | **Post** /rest/agile/1.0/sprint/{sprintId}/issue | Move issues to sprint and rank
[**PartiallyUpdateSprint**](SprintAPI.md#PartiallyUpdateSprint) | **Post** /rest/agile/1.0/sprint/{sprintId} | Partially update sprint
[**SetProperty**](SprintAPI.md#SetProperty) | **Put** /rest/agile/1.0/sprint/{sprintId}/properties/{propertyKey} | Set property
[**SwapSprint**](SprintAPI.md#SwapSprint) | **Post** /rest/agile/1.0/sprint/{sprintId}/swap | Swap sprint
[**UpdateSprint**](SprintAPI.md#UpdateSprint) | **Put** /rest/agile/1.0/sprint/{sprintId} | Update sprint



## CreateSprint

> CreateSprint(ctx).CreateSprintRequest(createSprintRequest).Execute()

Create sprint



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
	createSprintRequest := *openapiclient.NewCreateSprintRequest() // CreateSprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.CreateSprint(context.Background()).CreateSprintRequest(createSprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.CreateSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createSprintRequest** | [**CreateSprintRequest**](CreateSprintRequest.md) |  | 

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


## DeleteProperty

> DeleteProperty(ctx, sprintId, propertyKey).Execute()

Delete property



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
	sprintId := "sprintId_example" // string | the ID of the sprint from which the property will be removed.
	propertyKey := "propertyKey_example" // string | the key of the property to remove.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.DeleteProperty(context.Background(), sprintId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.DeleteProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **string** | the ID of the sprint from which the property will be removed. | 
**propertyKey** | **string** | the key of the property to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePropertyRequest struct via the builder pattern


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


## DeleteSprint

> DeleteSprint(ctx, sprintId).Execute()

Delete sprint



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
	sprintId := int64(789) // int64 | The ID of the sprint to delete.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.DeleteSprint(context.Background(), sprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.DeleteSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the sprint to delete. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSprintRequest struct via the builder pattern


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


## GetIssuesForSprint

> GetIssuesForSprint(ctx, sprintId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()

Get issues for sprint



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
	sprintId := int64(789) // int64 | The ID of the sprint that contains the requested issues.
	startAt := int64(789) // int64 | The starting index of the returned issues. Base index: 0. See the 'Pagination' section at the top of this page for more details. (optional)
	maxResults := int32(56) // int32 | The maximum number of issues to return per page. See the 'Pagination' section at the top of this page for more details. Note, the total number of issues returned is limited by the property 'jira.search.views.default.max' in your Jira instance. If you exceed this limit, your results will be truncated. (optional)
	jql := "jql_example" // string | Filters results using a JQL query. If you define an order in your JQL query, it will override the default order of the returned issues.   Note that `username` and `userkey` can't be used as search terms for this parameter due to privacy reasons. Use `accountId` instead. (optional)
	validateQuery := true // bool | Specifies whether to validate the JQL query or not. Default: true. (optional)
	fields := []map[string]interface{}{map[string]interface{}(123)} // []map[string]interface{} | The list of fields to return for each issue. By default, all navigable and Agile fields are returned. (optional)
	expand := "expand_example" // string | A comma-separated list of the parameters to expand. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.GetIssuesForSprint(context.Background(), sprintId).StartAt(startAt).MaxResults(maxResults).Jql(jql).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.GetIssuesForSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the sprint that contains the requested issues. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuesForSprintRequest struct via the builder pattern


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


## GetPropertiesKeys

> GetPropertiesKeys(ctx, sprintId).Execute()

Get properties keys



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
	sprintId := "sprintId_example" // string | the ID of the sprint from which property keys will be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.GetPropertiesKeys(context.Background(), sprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.GetPropertiesKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **string** | the ID of the sprint from which property keys will be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertiesKeysRequest struct via the builder pattern


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


## GetProperty

> GetProperty(ctx, sprintId, propertyKey).Execute()

Get property



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
	sprintId := "sprintId_example" // string | the ID of the sprint from which the property will be returned.
	propertyKey := "propertyKey_example" // string | the key of the property to return.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.GetProperty(context.Background(), sprintId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.GetProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **string** | the ID of the sprint from which the property will be returned. | 
**propertyKey** | **string** | the key of the property to return. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPropertyRequest struct via the builder pattern


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


## GetSprint

> GetSprint(ctx, sprintId).Execute()

Get sprint



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
	sprintId := int64(789) // int64 | The ID of the requested sprint.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.GetSprint(context.Background(), sprintId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.GetSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the requested sprint. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSprintRequest struct via the builder pattern


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


## MoveIssuesToSprintAndRank

> MoveIssuesToSprintAndRank(ctx, sprintId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()

Move issues to sprint and rank



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
	sprintId := int64(789) // int64 | The ID of the sprint that you want to assign issues to.
	moveIssuesToBacklogForBoardRequest := *openapiclient.NewMoveIssuesToBacklogForBoardRequest() // MoveIssuesToBacklogForBoardRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.MoveIssuesToSprintAndRank(context.Background(), sprintId).MoveIssuesToBacklogForBoardRequest(moveIssuesToBacklogForBoardRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.MoveIssuesToSprintAndRank``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the sprint that you want to assign issues to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiMoveIssuesToSprintAndRankRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moveIssuesToBacklogForBoardRequest** | [**MoveIssuesToBacklogForBoardRequest**](MoveIssuesToBacklogForBoardRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PartiallyUpdateSprint

> PartiallyUpdateSprint(ctx, sprintId).UpdateSprintRequest(updateSprintRequest).Execute()

Partially update sprint



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
	sprintId := int64(789) // int64 | The ID of the sprint to update.
	updateSprintRequest := *openapiclient.NewUpdateSprintRequest() // UpdateSprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.PartiallyUpdateSprint(context.Background(), sprintId).UpdateSprintRequest(updateSprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.PartiallyUpdateSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the sprint to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPartiallyUpdateSprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSprintRequest** | [**UpdateSprintRequest**](UpdateSprintRequest.md) |  | 

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


## SetProperty

> interface{} SetProperty(ctx, sprintId, propertyKey).Body(body).Execute()

Set property



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
	sprintId := "sprintId_example" // string | the ID of the sprint on which the property will be set.
	propertyKey := "propertyKey_example" // string | the key of the sprint's property. The maximum length of the key is 255 bytes.
	body := interface{}(987) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SprintAPI.SetProperty(context.Background(), sprintId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.SetProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `SprintAPI.SetProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **string** | the ID of the sprint on which the property will be set. | 
**propertyKey** | **string** | the key of the sprint&#39;s property. The maximum length of the key is 255 bytes. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetPropertyRequest struct via the builder pattern


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


## SwapSprint

> SwapSprint(ctx, sprintId).SwapSprintRequest(swapSprintRequest).Execute()

Swap sprint



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
	sprintId := int64(789) // int64 | The ID of the sprint to swap.
	swapSprintRequest := *openapiclient.NewSwapSprintRequest() // SwapSprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.SwapSprint(context.Background(), sprintId).SwapSprintRequest(swapSprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.SwapSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | The ID of the sprint to swap. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSwapSprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **swapSprintRequest** | [**SwapSprintRequest**](SwapSprintRequest.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateSprint

> UpdateSprint(ctx, sprintId).UpdateSprintRequest(updateSprintRequest).Execute()

Update sprint



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
	sprintId := int64(789) // int64 | the ID of the sprint to update.
	updateSprintRequest := *openapiclient.NewUpdateSprintRequest() // UpdateSprintRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SprintAPI.UpdateSprint(context.Background(), sprintId).UpdateSprintRequest(updateSprintRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SprintAPI.UpdateSprint``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sprintId** | **int64** | the ID of the sprint to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSprintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSprintRequest** | [**UpdateSprintRequest**](UpdateSprintRequest.md) |  | 

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

