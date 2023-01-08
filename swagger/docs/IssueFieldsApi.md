# \IssueFieldsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomField**](IssueFieldsAPI.md#CreateCustomField) | **Post** /rest/api/3/field | Create custom field
[**DeleteCustomField**](IssueFieldsAPI.md#DeleteCustomField) | **Delete** /rest/api/3/field/{id} | Delete custom field
[**GetContextsForFieldDeprecated**](IssueFieldsAPI.md#GetContextsForFieldDeprecated) | **Get** /rest/api/3/field/{fieldId}/contexts | Get contexts for a field
[**GetFields**](IssueFieldsAPI.md#GetFields) | **Get** /rest/api/3/field | Get fields
[**GetFieldsPaginated**](IssueFieldsAPI.md#GetFieldsPaginated) | **Get** /rest/api/3/field/search | Get fields paginated
[**GetProjectFields**](IssueFieldsAPI.md#GetProjectFields) | **Get** /rest/api/3/projects/fields | Get fields for projects
[**GetTrashedFieldsPaginated**](IssueFieldsAPI.md#GetTrashedFieldsPaginated) | **Get** /rest/api/3/field/search/trashed | Get fields in trash paginated
[**RestoreCustomField**](IssueFieldsAPI.md#RestoreCustomField) | **Post** /rest/api/3/field/{id}/restore | Restore custom field from trash
[**TrashCustomField**](IssueFieldsAPI.md#TrashCustomField) | **Post** /rest/api/3/field/{id}/trash | Move custom field to trash
[**UpdateCustomField**](IssueFieldsAPI.md#UpdateCustomField) | **Put** /rest/api/3/field/{fieldId} | Update custom field



## CreateCustomField

> FieldDetails CreateCustomField(ctx).CustomFieldDefinitionJsonBean(customFieldDefinitionJsonBean).Execute()

Create custom field



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
	customFieldDefinitionJsonBean := *openapiclient.NewCustomFieldDefinitionJsonBean("Name_example", "Type_example") // CustomFieldDefinitionJsonBean | Definition of the custom field to be created

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.CreateCustomField(context.Background()).CustomFieldDefinitionJsonBean(customFieldDefinitionJsonBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.CreateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomField`: FieldDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.CreateCustomField`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **customFieldDefinitionJsonBean** | [**CustomFieldDefinitionJsonBean**](CustomFieldDefinitionJsonBean.md) | Definition of the custom field to be created | 

### Return type

[**FieldDetails**](FieldDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomField

> DeleteCustomField(ctx, id).Execute()

Delete custom field



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
	id := "id_example" // string | The ID of a custom field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueFieldsAPI.DeleteCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.DeleteCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of a custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetContextsForFieldDeprecated

> PageBeanContext GetContextsForFieldDeprecated(ctx, fieldId).StartAt(startAt).MaxResults(maxResults).Execute()

Get contexts for a field



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
	fieldId := "fieldId_example" // string | The ID of the field to return contexts for.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.GetContextsForFieldDeprecated(context.Background(), fieldId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.GetContextsForFieldDeprecated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextsForFieldDeprecated`: PageBeanContext
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.GetContextsForFieldDeprecated`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the field to return contexts for. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsForFieldDeprecatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 20]

### Return type

[**PageBeanContext**](PageBeanContext.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFields

> []FieldDetails GetFields(ctx).Execute()

Get fields



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.GetFields(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.GetFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFields`: []FieldDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.GetFields`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsRequest struct via the builder pattern


### Return type

[**[]FieldDetails**](FieldDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldsPaginated

> PageBeanField GetFieldsPaginated(ctx).StartAt(startAt).MaxResults(maxResults).Type_(type_).Id(id).Query(query).OrderBy(orderBy).Expand(expand).ProjectIds(projectIds).Execute()

Get fields paginated



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
	type_ := []string{"Type_example"} // []string | The type of fields to search. (optional)
	id := []string{"Inner_example"} // []string | The IDs of the custom fields to return or, where `query` is specified, filter. (optional)
	query := "query_example" // string | String used to perform a case-insensitive partial match with field names or descriptions. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by:   *  `contextsCount` sorts by the number of contexts related to a field  *  `lastUsed` sorts by the date when the value of the field last changed  *  `name` sorts by the field name  *  `screensCount` sorts by the number of screens related to a field (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `key` returns the key for each field  *  `stableId` returns the stableId for each field  *  `lastUsed` returns the date when the value of the field last changed  *  `screensCount` returns the number of screens related to a field  *  `contextsCount` returns the number of contexts related to a field  *  `isLocked` returns information about whether the field is locked  *  `searcherKey` returns the searcher key for each custom field (optional)
	projectIds := []int64{int64(123)} // []int64 | The IDs of the projects to filter the fields by. Fields belonging to project Ids that the user does not have access to will not be returned (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.GetFieldsPaginated(context.Background()).StartAt(startAt).MaxResults(maxResults).Type_(type_).Id(id).Query(query).OrderBy(orderBy).Expand(expand).ProjectIds(projectIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.GetFieldsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldsPaginated`: PageBeanField
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.GetFieldsPaginated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **type_** | **[]string** | The type of fields to search. | 
 **id** | **[]string** | The IDs of the custom fields to return or, where &#x60;query&#x60; is specified, filter. | 
 **query** | **string** | String used to perform a case-insensitive partial match with field names or descriptions. | 
 **orderBy** | **string** | [Order](#ordering) the results by:   *  &#x60;contextsCount&#x60; sorts by the number of contexts related to a field  *  &#x60;lastUsed&#x60; sorts by the date when the value of the field last changed  *  &#x60;name&#x60; sorts by the field name  *  &#x60;screensCount&#x60; sorts by the number of screens related to a field | 
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;key&#x60; returns the key for each field  *  &#x60;stableId&#x60; returns the stableId for each field  *  &#x60;lastUsed&#x60; returns the date when the value of the field last changed  *  &#x60;screensCount&#x60; returns the number of screens related to a field  *  &#x60;contextsCount&#x60; returns the number of contexts related to a field  *  &#x60;isLocked&#x60; returns information about whether the field is locked  *  &#x60;searcherKey&#x60; returns the searcher key for each custom field | 
 **projectIds** | **[]int64** | The IDs of the projects to filter the fields by. Fields belonging to project Ids that the user does not have access to will not be returned | 

### Return type

[**PageBeanField**](PageBeanField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectFields

> PageBean2ProjectFieldBean GetProjectFields(ctx).ProjectId(projectId).WorkTypeId(workTypeId).StartAt(startAt).MaxResults(maxResults).FieldId(fieldId).Execute()

Get fields for projects



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
	projectId := []int64{int64(123)} // []int64 | The IDs of projects to return fields for.
	workTypeId := []int64{int64(123)} // []int64 | The IDs of work types (issue types) to return fields for.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	fieldId := []string{"["summary","description"]"} // []string | The IDs of fields to return. If not provided, all fields are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.GetProjectFields(context.Background()).ProjectId(projectId).WorkTypeId(workTypeId).StartAt(startAt).MaxResults(maxResults).FieldId(fieldId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.GetProjectFields``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectFields`: PageBean2ProjectFieldBean
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.GetProjectFields`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectFieldsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **[]int64** | The IDs of projects to return fields for. | 
 **workTypeId** | **[]int64** | The IDs of work types (issue types) to return fields for. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **fieldId** | **[]string** | The IDs of fields to return. If not provided, all fields are returned. | 

### Return type

[**PageBean2ProjectFieldBean**](PageBean2ProjectFieldBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTrashedFieldsPaginated

> PageBeanField GetTrashedFieldsPaginated(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).Query(query).Expand(expand).OrderBy(orderBy).Execute()

Get fields in trash paginated



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
	id := []string{"Inner_example"} // []string |  (optional)
	query := "query_example" // string | String used to perform a case-insensitive partial match with field names or descriptions. (optional)
	expand := "expand_example" // string |  (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `name` sorts by the field name  *  `trashDate` sorts by the date the field was moved to the trash  *  `plannedDeletionDate` sorts by the planned deletion date (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.GetTrashedFieldsPaginated(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).Query(query).Expand(expand).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.GetTrashedFieldsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetTrashedFieldsPaginated`: PageBeanField
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.GetTrashedFieldsPaginated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTrashedFieldsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **id** | **[]string** |  | 
 **query** | **string** | String used to perform a case-insensitive partial match with field names or descriptions. | 
 **expand** | **string** |  | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;name&#x60; sorts by the field name  *  &#x60;trashDate&#x60; sorts by the date the field was moved to the trash  *  &#x60;plannedDeletionDate&#x60; sorts by the planned deletion date | 

### Return type

[**PageBeanField**](PageBeanField.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RestoreCustomField

> interface{} RestoreCustomField(ctx, id).Execute()

Restore custom field from trash



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
	id := "id_example" // string | The ID of a custom field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.RestoreCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.RestoreCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RestoreCustomField`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.RestoreCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of a custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRestoreCustomFieldRequest struct via the builder pattern


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


## TrashCustomField

> interface{} TrashCustomField(ctx, id).Execute()

Move custom field to trash



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
	id := "id_example" // string | The ID of a custom field.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.TrashCustomField(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.TrashCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TrashCustomField`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.TrashCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of a custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiTrashCustomFieldRequest struct via the builder pattern


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


## UpdateCustomField

> interface{} UpdateCustomField(ctx, fieldId).UpdateCustomFieldDetails(updateCustomFieldDetails).Execute()

Update custom field



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
	fieldId := "fieldId_example" // string | The ID of the custom field.
	updateCustomFieldDetails := *openapiclient.NewUpdateCustomFieldDetails() // UpdateCustomFieldDetails | The custom field update details.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueFieldsAPI.UpdateCustomField(context.Background(), fieldId).UpdateCustomFieldDetails(updateCustomFieldDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueFieldsAPI.UpdateCustomField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomField`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueFieldsAPI.UpdateCustomField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomFieldDetails** | [**UpdateCustomFieldDetails**](UpdateCustomFieldDetails.md) | The custom field update details. | 

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

