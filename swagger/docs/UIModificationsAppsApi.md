# \UIModificationsAppsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUiModification**](UIModificationsAppsAPI.md#CreateUiModification) | **Post** /rest/api/3/uiModifications | Create UI modification
[**DeleteUiModification**](UIModificationsAppsAPI.md#DeleteUiModification) | **Delete** /rest/api/3/uiModifications/{uiModificationId} | Delete UI modification
[**GetUiModifications**](UIModificationsAppsAPI.md#GetUiModifications) | **Get** /rest/api/3/uiModifications | Get UI modifications
[**UpdateUiModification**](UIModificationsAppsAPI.md#UpdateUiModification) | **Put** /rest/api/3/uiModifications/{uiModificationId} | Update UI modification



## CreateUiModification

> UiModificationIdentifiers CreateUiModification(ctx).CreateUiModificationDetails(createUiModificationDetails).Execute()

Create UI modification



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
	createUiModificationDetails := *openapiclient.NewCreateUiModificationDetails("Name_example") // CreateUiModificationDetails | Details of the UI modification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIModificationsAppsAPI.CreateUiModification(context.Background()).CreateUiModificationDetails(createUiModificationDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIModificationsAppsAPI.CreateUiModification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUiModification`: UiModificationIdentifiers
	fmt.Fprintf(os.Stdout, "Response from `UIModificationsAppsAPI.CreateUiModification`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUiModificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createUiModificationDetails** | [**CreateUiModificationDetails**](CreateUiModificationDetails.md) | Details of the UI modification. | 

### Return type

[**UiModificationIdentifiers**](UiModificationIdentifiers.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUiModification

> interface{} DeleteUiModification(ctx, uiModificationId).Execute()

Delete UI modification



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
	uiModificationId := "uiModificationId_example" // string | The ID of the UI modification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIModificationsAppsAPI.DeleteUiModification(context.Background(), uiModificationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIModificationsAppsAPI.DeleteUiModification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteUiModification`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UIModificationsAppsAPI.DeleteUiModification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uiModificationId** | **string** | The ID of the UI modification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUiModificationRequest struct via the builder pattern


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


## GetUiModifications

> PageBeanUiModificationDetails GetUiModifications(ctx).StartAt(startAt).MaxResults(maxResults).Expand(expand).Execute()

Get UI modifications



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
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `data` Returns UI modification data.  *  `contexts` Returns UI modification contexts. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIModificationsAppsAPI.GetUiModifications(context.Background()).StartAt(startAt).MaxResults(maxResults).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIModificationsAppsAPI.GetUiModifications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetUiModifications`: PageBeanUiModificationDetails
	fmt.Fprintf(os.Stdout, "Response from `UIModificationsAppsAPI.GetUiModifications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUiModificationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;data&#x60; Returns UI modification data.  *  &#x60;contexts&#x60; Returns UI modification contexts. | 

### Return type

[**PageBeanUiModificationDetails**](PageBeanUiModificationDetails.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUiModification

> interface{} UpdateUiModification(ctx, uiModificationId).UpdateUiModificationDetails(updateUiModificationDetails).Execute()

Update UI modification



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
	uiModificationId := "uiModificationId_example" // string | The ID of the UI modification.
	updateUiModificationDetails := *openapiclient.NewUpdateUiModificationDetails() // UpdateUiModificationDetails | Details of the UI modification.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.UIModificationsAppsAPI.UpdateUiModification(context.Background(), uiModificationId).UpdateUiModificationDetails(updateUiModificationDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `UIModificationsAppsAPI.UpdateUiModification``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateUiModification`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `UIModificationsAppsAPI.UpdateUiModification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**uiModificationId** | **string** | The ID of the UI modification. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUiModificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateUiModificationDetails** | [**UpdateUiModificationDetails**](UpdateUiModificationDetails.md) | Details of the UI modification. | 

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

