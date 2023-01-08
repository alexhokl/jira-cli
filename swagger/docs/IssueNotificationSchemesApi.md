# \IssueNotificationSchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNotifications**](IssueNotificationSchemesAPI.md#AddNotifications) | **Put** /rest/api/3/notificationscheme/{id}/notification | Add notifications to notification scheme
[**CreateNotificationScheme**](IssueNotificationSchemesAPI.md#CreateNotificationScheme) | **Post** /rest/api/3/notificationscheme | Create notification scheme
[**DeleteNotificationScheme**](IssueNotificationSchemesAPI.md#DeleteNotificationScheme) | **Delete** /rest/api/3/notificationscheme/{notificationSchemeId} | Delete notification scheme
[**GetNotificationScheme**](IssueNotificationSchemesAPI.md#GetNotificationScheme) | **Get** /rest/api/3/notificationscheme/{id} | Get notification scheme
[**GetNotificationSchemeToProjectMappings**](IssueNotificationSchemesAPI.md#GetNotificationSchemeToProjectMappings) | **Get** /rest/api/3/notificationscheme/project | Get projects using notification schemes paginated
[**GetNotificationSchemes**](IssueNotificationSchemesAPI.md#GetNotificationSchemes) | **Get** /rest/api/3/notificationscheme | Get notification schemes paginated
[**RemoveNotificationFromNotificationScheme**](IssueNotificationSchemesAPI.md#RemoveNotificationFromNotificationScheme) | **Delete** /rest/api/3/notificationscheme/{notificationSchemeId}/notification/{notificationId} | Remove notification from notification scheme
[**UpdateNotificationScheme**](IssueNotificationSchemesAPI.md#UpdateNotificationScheme) | **Put** /rest/api/3/notificationscheme/{id} | Update notification scheme



## AddNotifications

> interface{} AddNotifications(ctx, id).AddNotificationsDetails(addNotificationsDetails).Execute()

Add notifications to notification scheme



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
	id := "id_example" // string | The ID of the notification scheme.
	addNotificationsDetails := *openapiclient.NewAddNotificationsDetails([]openapiclient.NotificationSchemeEventDetails{*openapiclient.NewNotificationSchemeEventDetails(*openapiclient.NewNotificationSchemeEventTypeId("Id_example"), []openapiclient.NotificationSchemeNotificationDetails{*openapiclient.NewNotificationSchemeNotificationDetails("NotificationType_example")})}) // AddNotificationsDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.AddNotifications(context.Background(), id).AddNotificationsDetails(addNotificationsDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.AddNotifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddNotifications`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.AddNotifications`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the notification scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddNotificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addNotificationsDetails** | [**AddNotificationsDetails**](AddNotificationsDetails.md) |  | 

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


## CreateNotificationScheme

> NotificationSchemeId CreateNotificationScheme(ctx).CreateNotificationSchemeDetails(createNotificationSchemeDetails).Execute()

Create notification scheme



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
	createNotificationSchemeDetails := *openapiclient.NewCreateNotificationSchemeDetails("Name_example") // CreateNotificationSchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.CreateNotificationScheme(context.Background()).CreateNotificationSchemeDetails(createNotificationSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.CreateNotificationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateNotificationScheme`: NotificationSchemeId
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.CreateNotificationScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateNotificationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createNotificationSchemeDetails** | [**CreateNotificationSchemeDetails**](CreateNotificationSchemeDetails.md) |  | 

### Return type

[**NotificationSchemeId**](NotificationSchemeId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNotificationScheme

> interface{} DeleteNotificationScheme(ctx, notificationSchemeId).Execute()

Delete notification scheme



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
	notificationSchemeId := "notificationSchemeId_example" // string | The ID of the notification scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.DeleteNotificationScheme(context.Background(), notificationSchemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.DeleteNotificationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteNotificationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.DeleteNotificationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationSchemeId** | **string** | The ID of the notification scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNotificationSchemeRequest struct via the builder pattern


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


## GetNotificationScheme

> NotificationScheme GetNotificationScheme(ctx, id).Expand(expand).Execute()

Get notification scheme



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
	id := int64(789) // int64 | The ID of the notification scheme. Use [Get notification schemes paginated](#api-rest-api-3-notificationscheme-get) to get a list of notification scheme IDs.
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `all` Returns all expandable information  *  `field` Returns information about any custom fields assigned to receive an event  *  `group` Returns information about any groups assigned to receive an event  *  `notificationSchemeEvents` Returns a list of event associations. This list is returned for all expandable information  *  `projectRole` Returns information about any project roles assigned to receive an event  *  `user` Returns information about any users assigned to receive an event (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.GetNotificationScheme(context.Background(), id).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.GetNotificationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationScheme`: NotificationScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.GetNotificationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the notification scheme. Use [Get notification schemes paginated](#api-rest-api-3-notificationscheme-get) to get a list of notification scheme IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information  *  &#x60;field&#x60; Returns information about any custom fields assigned to receive an event  *  &#x60;group&#x60; Returns information about any groups assigned to receive an event  *  &#x60;notificationSchemeEvents&#x60; Returns a list of event associations. This list is returned for all expandable information  *  &#x60;projectRole&#x60; Returns information about any project roles assigned to receive an event  *  &#x60;user&#x60; Returns information about any users assigned to receive an event | 

### Return type

[**NotificationScheme**](NotificationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSchemeToProjectMappings

> PageBeanNotificationSchemeAndProjectMappingJsonBean GetNotificationSchemeToProjectMappings(ctx).StartAt(startAt).MaxResults(maxResults).NotificationSchemeId(notificationSchemeId).ProjectId(projectId).Execute()

Get projects using notification schemes paginated



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	notificationSchemeId := []string{"Inner_example"} // []string | The list of notifications scheme IDs to be filtered out (optional)
	projectId := []string{"Inner_example"} // []string | The list of project IDs to be filtered out (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.GetNotificationSchemeToProjectMappings(context.Background()).StartAt(startAt).MaxResults(maxResults).NotificationSchemeId(notificationSchemeId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.GetNotificationSchemeToProjectMappings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSchemeToProjectMappings`: PageBeanNotificationSchemeAndProjectMappingJsonBean
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.GetNotificationSchemeToProjectMappings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSchemeToProjectMappingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **notificationSchemeId** | **[]string** | The list of notifications scheme IDs to be filtered out | 
 **projectId** | **[]string** | The list of project IDs to be filtered out | 

### Return type

[**PageBeanNotificationSchemeAndProjectMappingJsonBean**](PageBeanNotificationSchemeAndProjectMappingJsonBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNotificationSchemes

> PageBeanNotificationScheme GetNotificationSchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).OnlyDefault(onlyDefault).Expand(expand).Execute()

Get notification schemes paginated



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	id := []string{"Inner_example"} // []string | The list of notification schemes IDs to be filtered by (optional)
	projectId := []string{"Inner_example"} // []string | The list of projects IDs to be filtered by (optional)
	onlyDefault := true // bool | When set to true, returns only the default notification scheme. If you provide project IDs not associated with the default, returns an empty page. The default value is false. (optional) (default to false)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `all` Returns all expandable information  *  `field` Returns information about any custom fields assigned to receive an event  *  `group` Returns information about any groups assigned to receive an event  *  `notificationSchemeEvents` Returns a list of event associations. This list is returned for all expandable information  *  `projectRole` Returns information about any project roles assigned to receive an event  *  `user` Returns information about any users assigned to receive an event (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.GetNotificationSchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).OnlyDefault(onlyDefault).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.GetNotificationSchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNotificationSchemes`: PageBeanNotificationScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.GetNotificationSchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetNotificationSchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of notification schemes IDs to be filtered by | 
 **projectId** | **[]string** | The list of projects IDs to be filtered by | 
 **onlyDefault** | **bool** | When set to true, returns only the default notification scheme. If you provide project IDs not associated with the default, returns an empty page. The default value is false. | [default to false]
 **expand** | **string** | Use [expand](#expansion) to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information  *  &#x60;field&#x60; Returns information about any custom fields assigned to receive an event  *  &#x60;group&#x60; Returns information about any groups assigned to receive an event  *  &#x60;notificationSchemeEvents&#x60; Returns a list of event associations. This list is returned for all expandable information  *  &#x60;projectRole&#x60; Returns information about any project roles assigned to receive an event  *  &#x60;user&#x60; Returns information about any users assigned to receive an event | 

### Return type

[**PageBeanNotificationScheme**](PageBeanNotificationScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveNotificationFromNotificationScheme

> interface{} RemoveNotificationFromNotificationScheme(ctx, notificationSchemeId, notificationId).Execute()

Remove notification from notification scheme



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
	notificationSchemeId := "notificationSchemeId_example" // string | The ID of the notification scheme.
	notificationId := "notificationId_example" // string | The ID of the notification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.RemoveNotificationFromNotificationScheme(context.Background(), notificationSchemeId, notificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.RemoveNotificationFromNotificationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveNotificationFromNotificationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.RemoveNotificationFromNotificationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**notificationSchemeId** | **string** | The ID of the notification scheme. | 
**notificationId** | **string** | The ID of the notification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveNotificationFromNotificationSchemeRequest struct via the builder pattern


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


## UpdateNotificationScheme

> interface{} UpdateNotificationScheme(ctx, id).UpdateNotificationSchemeDetails(updateNotificationSchemeDetails).Execute()

Update notification scheme



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
	id := "id_example" // string | The ID of the notification scheme.
	updateNotificationSchemeDetails := *openapiclient.NewUpdateNotificationSchemeDetails() // UpdateNotificationSchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueNotificationSchemesAPI.UpdateNotificationScheme(context.Background(), id).UpdateNotificationSchemeDetails(updateNotificationSchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNotificationSchemesAPI.UpdateNotificationScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateNotificationScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueNotificationSchemesAPI.UpdateNotificationScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the notification scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNotificationSchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNotificationSchemeDetails** | [**UpdateNotificationSchemeDetails**](UpdateNotificationSchemeDetails.md) |  | 

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

