# \IssueTypesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssueType**](IssueTypesAPI.md#CreateIssueType) | **Post** /rest/api/3/issuetype | Create issue type
[**CreateIssueTypeAvatar**](IssueTypesAPI.md#CreateIssueTypeAvatar) | **Post** /rest/api/3/issuetype/{id}/avatar2 | Load issue type avatar
[**DeleteIssueType**](IssueTypesAPI.md#DeleteIssueType) | **Delete** /rest/api/3/issuetype/{id} | Delete issue type
[**GetAlternativeIssueTypes**](IssueTypesAPI.md#GetAlternativeIssueTypes) | **Get** /rest/api/3/issuetype/{id}/alternatives | Get alternative issue types
[**GetIssueAllTypes**](IssueTypesAPI.md#GetIssueAllTypes) | **Get** /rest/api/3/issuetype | Get all issue types for user
[**GetIssueType**](IssueTypesAPI.md#GetIssueType) | **Get** /rest/api/3/issuetype/{id} | Get issue type
[**GetIssueTypesForProject**](IssueTypesAPI.md#GetIssueTypesForProject) | **Get** /rest/api/3/issuetype/project | Get issue types for project
[**UpdateIssueType**](IssueTypesAPI.md#UpdateIssueType) | **Put** /rest/api/3/issuetype/{id} | Update issue type



## CreateIssueType

> IssueTypeDetails CreateIssueType(ctx).IssueTypeCreateBean(issueTypeCreateBean).Execute()

Create issue type



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
	issueTypeCreateBean := *openapiclient.NewIssueTypeCreateBean("Name_example") // IssueTypeCreateBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.CreateIssueType(context.Background()).IssueTypeCreateBean(issueTypeCreateBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.CreateIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueType`: IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.CreateIssueType`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issueTypeCreateBean** | [**IssueTypeCreateBean**](IssueTypeCreateBean.md) |  | 

### Return type

[**IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueTypeAvatar

> Avatar CreateIssueTypeAvatar(ctx, id).Size(size).Body(body).X(x).Y(y).Execute()

Load issue type avatar



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
	id := "id_example" // string | The ID of the issue type.
	size := int32(56) // int32 | The length of each side of the crop region.
	body := interface{}(987) // interface{} | 
	x := int32(56) // int32 | The X coordinate of the top-left corner of the crop region. (optional) (default to 0)
	y := int32(56) // int32 | The Y coordinate of the top-left corner of the crop region. (optional) (default to 0)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.CreateIssueTypeAvatar(context.Background(), id).Size(size).Body(body).X(x).Y(y).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.CreateIssueTypeAvatar``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueTypeAvatar`: Avatar
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.CreateIssueTypeAvatar`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueTypeAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **int32** | The length of each side of the crop region. | 
 **body** | **interface{}** |  | 
 **x** | **int32** | The X coordinate of the top-left corner of the crop region. | [default to 0]
 **y** | **int32** | The Y coordinate of the top-left corner of the crop region. | [default to 0]

### Return type

[**Avatar**](Avatar.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueType

> DeleteIssueType(ctx, id).AlternativeIssueTypeId(alternativeIssueTypeId).Execute()

Delete issue type



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
	id := "id_example" // string | The ID of the issue type.
	alternativeIssueTypeId := "alternativeIssueTypeId_example" // string | The ID of the replacement issue type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueTypesAPI.DeleteIssueType(context.Background(), id).AlternativeIssueTypeId(alternativeIssueTypeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.DeleteIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **alternativeIssueTypeId** | **string** | The ID of the replacement issue type. | 

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


## GetAlternativeIssueTypes

> []IssueTypeDetails GetAlternativeIssueTypes(ctx, id).Execute()

Get alternative issue types



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
	id := "id_example" // string | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.GetAlternativeIssueTypes(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.GetAlternativeIssueTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAlternativeIssueTypes`: []IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.GetAlternativeIssueTypes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAlternativeIssueTypesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueAllTypes

> []IssueTypeDetails GetIssueAllTypes(ctx).Execute()

Get all issue types for user



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
	resp, r, err := apiClient.IssueTypesAPI.GetIssueAllTypes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.GetIssueAllTypes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueAllTypes`: []IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.GetIssueAllTypes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueAllTypesRequest struct via the builder pattern


### Return type

[**[]IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueType

> IssueTypeDetails GetIssueType(ctx, id).Execute()

Get issue type



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
	id := "id_example" // string | The ID of the issue type.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.GetIssueType(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.GetIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueType`: IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.GetIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueTypesForProject

> []IssueTypeDetails GetIssueTypesForProject(ctx).ProjectId(projectId).Level(level).Execute()

Get issue types for project



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
	projectId := int64(789) // int64 | The ID of the project.
	level := int32(56) // int32 | The level of the issue type to filter by. Use:   *  `-1` for Subtask.  *  `0` for Base.  *  `1` for Epic. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.GetIssueTypesForProject(context.Background()).ProjectId(projectId).Level(level).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.GetIssueTypesForProject``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueTypesForProject`: []IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.GetIssueTypesForProject`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueTypesForProjectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **projectId** | **int64** | The ID of the project. | 
 **level** | **int32** | The level of the issue type to filter by. Use:   *  &#x60;-1&#x60; for Subtask.  *  &#x60;0&#x60; for Base.  *  &#x60;1&#x60; for Epic. | 

### Return type

[**[]IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssueType

> IssueTypeDetails UpdateIssueType(ctx, id).IssueTypeUpdateBean(issueTypeUpdateBean).Execute()

Update issue type



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
	id := "id_example" // string | The ID of the issue type.
	issueTypeUpdateBean := *openapiclient.NewIssueTypeUpdateBean() // IssueTypeUpdateBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueTypesAPI.UpdateIssueType(context.Background(), id).IssueTypeUpdateBean(issueTypeUpdateBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueTypesAPI.UpdateIssueType``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueType`: IssueTypeDetails
	fmt.Fprintf(os.Stdout, "Response from `IssueTypesAPI.UpdateIssueType`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue type. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueTypeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueTypeUpdateBean** | [**IssueTypeUpdateBean**](IssueTypeUpdateBean.md) |  | 

### Return type

[**IssueTypeDetails**](IssueTypeDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

