# \IssueCustomFieldContextsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddIssueTypesToContext**](IssueCustomFieldContextsAPI.md#AddIssueTypesToContext) | **Put** /rest/api/3/field/{fieldId}/context/{contextId}/issuetype | Add issue types to context
[**AssignProjectsToCustomFieldContext**](IssueCustomFieldContextsAPI.md#AssignProjectsToCustomFieldContext) | **Put** /rest/api/3/field/{fieldId}/context/{contextId}/project | Assign custom field context to projects
[**CreateCustomFieldContext**](IssueCustomFieldContextsAPI.md#CreateCustomFieldContext) | **Post** /rest/api/3/field/{fieldId}/context | Create custom field context
[**DeleteCustomFieldContext**](IssueCustomFieldContextsAPI.md#DeleteCustomFieldContext) | **Delete** /rest/api/3/field/{fieldId}/context/{contextId} | Delete custom field context
[**GetContextsForField**](IssueCustomFieldContextsAPI.md#GetContextsForField) | **Get** /rest/api/3/field/{fieldId}/context | Get custom field contexts
[**GetCustomFieldContextsForProjectsAndIssueTypes**](IssueCustomFieldContextsAPI.md#GetCustomFieldContextsForProjectsAndIssueTypes) | **Post** /rest/api/3/field/{fieldId}/context/mapping | Get custom field contexts for projects and issue types
[**GetDefaultValues**](IssueCustomFieldContextsAPI.md#GetDefaultValues) | **Get** /rest/api/3/field/{fieldId}/context/defaultValue | Get custom field contexts default values
[**GetIssueTypeMappingsForContexts**](IssueCustomFieldContextsAPI.md#GetIssueTypeMappingsForContexts) | **Get** /rest/api/3/field/{fieldId}/context/issuetypemapping | Get issue types for custom field context
[**GetProjectContextMapping**](IssueCustomFieldContextsAPI.md#GetProjectContextMapping) | **Get** /rest/api/3/field/{fieldId}/context/projectmapping | Get project mappings for custom field context
[**RemoveCustomFieldContextFromProjects**](IssueCustomFieldContextsAPI.md#RemoveCustomFieldContextFromProjects) | **Post** /rest/api/3/field/{fieldId}/context/{contextId}/project/remove | Remove custom field context from projects
[**RemoveIssueTypesFromContext**](IssueCustomFieldContextsAPI.md#RemoveIssueTypesFromContext) | **Post** /rest/api/3/field/{fieldId}/context/{contextId}/issuetype/remove | Remove issue types from context
[**SetDefaultValues**](IssueCustomFieldContextsAPI.md#SetDefaultValues) | **Put** /rest/api/3/field/{fieldId}/context/defaultValue | Set custom field contexts default values
[**UpdateCustomFieldContext**](IssueCustomFieldContextsAPI.md#UpdateCustomFieldContext) | **Put** /rest/api/3/field/{fieldId}/context/{contextId} | Update custom field context



## AddIssueTypesToContext

> interface{} AddIssueTypesToContext(ctx, fieldId, contextId).IssueTypeIds(issueTypeIds).Execute()

Add issue types to context



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
	contextId := int64(789) // int64 | The ID of the context.
	issueTypeIds := *openapiclient.NewIssueTypeIds([]string{"IssueTypeIds_example"}) // IssueTypeIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.AddIssueTypesToContext(context.Background(), fieldId, contextId).IssueTypeIds(issueTypeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.AddIssueTypesToContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddIssueTypesToContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.AddIssueTypesToContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddIssueTypesToContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueTypeIds** | [**IssueTypeIds**](IssueTypeIds.md) |  | 

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


## AssignProjectsToCustomFieldContext

> interface{} AssignProjectsToCustomFieldContext(ctx, fieldId, contextId).ProjectIds(projectIds).Execute()

Assign custom field context to projects



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
	contextId := int64(789) // int64 | The ID of the context.
	projectIds := *openapiclient.NewProjectIds([]string{"ProjectIds_example"}) // ProjectIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.AssignProjectsToCustomFieldContext(context.Background(), fieldId, contextId).ProjectIds(projectIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.AssignProjectsToCustomFieldContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AssignProjectsToCustomFieldContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.AssignProjectsToCustomFieldContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAssignProjectsToCustomFieldContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIds** | [**ProjectIds**](ProjectIds.md) |  | 

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


## CreateCustomFieldContext

> CreateCustomFieldContext CreateCustomFieldContext(ctx, fieldId).CreateCustomFieldContext(createCustomFieldContext).Execute()

Create custom field context



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
	createCustomFieldContext := *openapiclient.NewCreateCustomFieldContext("Name_example") // CreateCustomFieldContext | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.CreateCustomFieldContext(context.Background(), fieldId).CreateCustomFieldContext(createCustomFieldContext).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.CreateCustomFieldContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateCustomFieldContext`: CreateCustomFieldContext
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.CreateCustomFieldContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomFieldContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createCustomFieldContext** | [**CreateCustomFieldContext**](CreateCustomFieldContext.md) |  | 

### Return type

[**CreateCustomFieldContext**](CreateCustomFieldContext.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCustomFieldContext

> interface{} DeleteCustomFieldContext(ctx, fieldId, contextId).Execute()

Delete custom field context



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
	contextId := int64(789) // int64 | The ID of the context.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.DeleteCustomFieldContext(context.Background(), fieldId, contextId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.DeleteCustomFieldContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCustomFieldContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.DeleteCustomFieldContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCustomFieldContextRequest struct via the builder pattern


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


## GetContextsForField

> PageBeanCustomFieldContext GetContextsForField(ctx, fieldId).IsAnyIssueType(isAnyIssueType).IsGlobalContext(isGlobalContext).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()

Get custom field contexts



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
	isAnyIssueType := true // bool | Whether to return contexts that apply to all issue types. (optional)
	isGlobalContext := true // bool | Whether to return contexts that apply to all projects. (optional)
	contextId := []int64{int64(123)} // []int64 | The list of context IDs. To include multiple contexts, separate IDs with ampersand: `contextId=10000&contextId=10001`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.GetContextsForField(context.Background(), fieldId).IsAnyIssueType(isAnyIssueType).IsGlobalContext(isGlobalContext).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.GetContextsForField``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetContextsForField`: PageBeanCustomFieldContext
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.GetContextsForField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetContextsForFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **isAnyIssueType** | **bool** | Whether to return contexts that apply to all issue types. | 
 **isGlobalContext** | **bool** | Whether to return contexts that apply to all projects. | 
 **contextId** | **[]int64** | The list of context IDs. To include multiple contexts, separate IDs with ampersand: &#x60;contextId&#x3D;10000&amp;contextId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanCustomFieldContext**](PageBeanCustomFieldContext.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCustomFieldContextsForProjectsAndIssueTypes

> PageBeanContextForProjectAndIssueType GetCustomFieldContextsForProjectsAndIssueTypes(ctx, fieldId).ProjectIssueTypeMappings(projectIssueTypeMappings).StartAt(startAt).MaxResults(maxResults).Execute()

Get custom field contexts for projects and issue types



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
	projectIssueTypeMappings := *openapiclient.NewProjectIssueTypeMappings([]openapiclient.ProjectIssueTypeMapping{*openapiclient.NewProjectIssueTypeMapping("IssueTypeId_example", "ProjectId_example")}) // ProjectIssueTypeMappings | The list of project and issue type mappings.
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.GetCustomFieldContextsForProjectsAndIssueTypes(context.Background(), fieldId).ProjectIssueTypeMappings(projectIssueTypeMappings).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.GetCustomFieldContextsForProjectsAndIssueTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCustomFieldContextsForProjectsAndIssueTypes`: PageBeanContextForProjectAndIssueType
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.GetCustomFieldContextsForProjectsAndIssueTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomFieldContextsForProjectsAndIssueTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **projectIssueTypeMappings** | [**ProjectIssueTypeMappings**](ProjectIssueTypeMappings.md) | The list of project and issue type mappings. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanContextForProjectAndIssueType**](PageBeanContextForProjectAndIssueType.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDefaultValues

> PageBeanCustomFieldContextDefaultValue GetDefaultValues(ctx, fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()

Get custom field contexts default values



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
	fieldId := "fieldId_example" // string | The ID of the custom field, for example `customfield\\_10000`.
	contextId := []int64{int64(123)} // []int64 | The IDs of the contexts. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.GetDefaultValues(context.Background(), fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.GetDefaultValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDefaultValues`: PageBeanCustomFieldContextDefaultValue
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.GetDefaultValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field, for example &#x60;customfield\\_10000&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDefaultValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextId** | **[]int64** | The IDs of the contexts. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanCustomFieldContextDefaultValue**](PageBeanCustomFieldContextDefaultValue.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypeMappingsForContexts

> PageBeanIssueTypeToContextMapping GetIssueTypeMappingsForContexts(ctx, fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()

Get issue types for custom field context



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
	contextId := []int64{int64(123)} // []int64 | The ID of the context. To include multiple contexts, provide an ampersand-separated list. For example, `contextId=10001&contextId=10002`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.GetIssueTypeMappingsForContexts(context.Background(), fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.GetIssueTypeMappingsForContexts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypeMappingsForContexts`: PageBeanIssueTypeToContextMapping
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.GetIssueTypeMappingsForContexts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeMappingsForContextsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextId** | **[]int64** | The ID of the context. To include multiple contexts, provide an ampersand-separated list. For example, &#x60;contextId&#x3D;10001&amp;contextId&#x3D;10002&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanIssueTypeToContextMapping**](PageBeanIssueTypeToContextMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetProjectContextMapping

> PageBeanCustomFieldContextProjectMapping GetProjectContextMapping(ctx, fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()

Get project mappings for custom field context



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
	fieldId := "fieldId_example" // string | The ID of the custom field, for example `customfield\\_10000`.
	contextId := []int64{int64(123)} // []int64 | The list of context IDs. To include multiple context, separate IDs with ampersand: `contextId=10000&contextId=10001`. (optional)
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.GetProjectContextMapping(context.Background(), fieldId).ContextId(contextId).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.GetProjectContextMapping``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetProjectContextMapping`: PageBeanCustomFieldContextProjectMapping
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.GetProjectContextMapping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field, for example &#x60;customfield\\_10000&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProjectContextMappingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **contextId** | **[]int64** | The list of context IDs. To include multiple context, separate IDs with ampersand: &#x60;contextId&#x3D;10000&amp;contextId&#x3D;10001&#x60;. | 
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanCustomFieldContextProjectMapping**](PageBeanCustomFieldContextProjectMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveCustomFieldContextFromProjects

> interface{} RemoveCustomFieldContextFromProjects(ctx, fieldId, contextId).ProjectIds(projectIds).Execute()

Remove custom field context from projects



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
	contextId := int64(789) // int64 | The ID of the context.
	projectIds := *openapiclient.NewProjectIds([]string{"ProjectIds_example"}) // ProjectIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.RemoveCustomFieldContextFromProjects(context.Background(), fieldId, contextId).ProjectIds(projectIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.RemoveCustomFieldContextFromProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveCustomFieldContextFromProjects`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.RemoveCustomFieldContextFromProjects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveCustomFieldContextFromProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **projectIds** | [**ProjectIds**](ProjectIds.md) |  | 

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


## RemoveIssueTypesFromContext

> interface{} RemoveIssueTypesFromContext(ctx, fieldId, contextId).IssueTypeIds(issueTypeIds).Execute()

Remove issue types from context



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
	contextId := int64(789) // int64 | The ID of the context.
	issueTypeIds := *openapiclient.NewIssueTypeIds([]string{"IssueTypeIds_example"}) // IssueTypeIds | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.RemoveIssueTypesFromContext(context.Background(), fieldId, contextId).IssueTypeIds(issueTypeIds).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.RemoveIssueTypesFromContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveIssueTypesFromContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.RemoveIssueTypesFromContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveIssueTypesFromContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueTypeIds** | [**IssueTypeIds**](IssueTypeIds.md) |  | 

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


## SetDefaultValues

> interface{} SetDefaultValues(ctx, fieldId).CustomFieldContextDefaultValueUpdate(customFieldContextDefaultValueUpdate).Execute()

Set custom field contexts default values



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
	customFieldContextDefaultValueUpdate := *openapiclient.NewCustomFieldContextDefaultValueUpdate() // CustomFieldContextDefaultValueUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.SetDefaultValues(context.Background(), fieldId).CustomFieldContextDefaultValueUpdate(customFieldContextDefaultValueUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.SetDefaultValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultValues`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.SetDefaultValues`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFieldContextDefaultValueUpdate** | [**CustomFieldContextDefaultValueUpdate**](CustomFieldContextDefaultValueUpdate.md) |  | 

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


## UpdateCustomFieldContext

> interface{} UpdateCustomFieldContext(ctx, fieldId, contextId).CustomFieldContextUpdateDetails(customFieldContextUpdateDetails).Execute()

Update custom field context



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
	contextId := int64(789) // int64 | The ID of the context.
	customFieldContextUpdateDetails := *openapiclient.NewCustomFieldContextUpdateDetails() // CustomFieldContextUpdateDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldContextsAPI.UpdateCustomFieldContext(context.Background(), fieldId, contextId).CustomFieldContextUpdateDetails(customFieldContextUpdateDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldContextsAPI.UpdateCustomFieldContext``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomFieldContext`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldContextsAPI.UpdateCustomFieldContext`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldId** | **string** | The ID of the custom field. | 
**contextId** | **int64** | The ID of the context. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldContextRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **customFieldContextUpdateDetails** | [**CustomFieldContextUpdateDetails**](CustomFieldContextUpdateDetails.md) |  | 

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

