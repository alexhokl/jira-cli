# \WebhooksAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWebhookById**](WebhooksAPI.md#DeleteWebhookById) | **Delete** /rest/api/3/webhook | Delete webhooks by ID
[**GetDynamicWebhooksForApp**](WebhooksAPI.md#GetDynamicWebhooksForApp) | **Get** /rest/api/3/webhook | Get dynamic webhooks for app
[**GetFailedWebhooks**](WebhooksAPI.md#GetFailedWebhooks) | **Get** /rest/api/3/webhook/failed | Get failed webhooks
[**RefreshWebhooks**](WebhooksAPI.md#RefreshWebhooks) | **Put** /rest/api/3/webhook/refresh | Extend webhook life
[**RegisterDynamicWebhooks**](WebhooksAPI.md#RegisterDynamicWebhooks) | **Post** /rest/api/3/webhook | Register dynamic webhooks



## DeleteWebhookById

> DeleteWebhookById(ctx).ContainerForWebhookIDs(containerForWebhookIDs).Execute()

Delete webhooks by ID



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
	containerForWebhookIDs := *openapiclient.NewContainerForWebhookIDs([]int64{int64(123)}) // ContainerForWebhookIDs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.WebhooksAPI.DeleteWebhookById(context.Background()).ContainerForWebhookIDs(containerForWebhookIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.DeleteWebhookById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerForWebhookIDs** | [**ContainerForWebhookIDs**](ContainerForWebhookIDs.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDynamicWebhooksForApp

> PageBeanWebhook GetDynamicWebhooksForApp(ctx).StartAt(startAt).MaxResults(maxResults).Execute()

Get dynamic webhooks for app



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
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 100)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetDynamicWebhooksForApp(context.Background()).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetDynamicWebhooksForApp``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicWebhooksForApp`: PageBeanWebhook
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetDynamicWebhooksForApp`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicWebhooksForAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 100]

### Return type

[**PageBeanWebhook**](PageBeanWebhook.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFailedWebhooks

> FailedWebhooks GetFailedWebhooks(ctx).MaxResults(maxResults).After(after).Execute()

Get failed webhooks



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
	maxResults := int32(56) // int32 | The maximum number of webhooks to return per page. If obeying the maxResults directive would result in records with the same failure time being split across pages, the directive is ignored and all records with the same failure time included on the page. (optional)
	after := int64(789) // int64 | The time after which any webhook failure must have occurred for the record to be returned, expressed as milliseconds since the UNIX epoch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.GetFailedWebhooks(context.Background()).MaxResults(maxResults).After(after).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.GetFailedWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFailedWebhooks`: FailedWebhooks
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.GetFailedWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFailedWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **maxResults** | **int32** | The maximum number of webhooks to return per page. If obeying the maxResults directive would result in records with the same failure time being split across pages, the directive is ignored and all records with the same failure time included on the page. | 
 **after** | **int64** | The time after which any webhook failure must have occurred for the record to be returned, expressed as milliseconds since the UNIX epoch. | 

### Return type

[**FailedWebhooks**](FailedWebhooks.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RefreshWebhooks

> WebhooksExpirationDate RefreshWebhooks(ctx).ContainerForWebhookIDs(containerForWebhookIDs).Execute()

Extend webhook life



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
	containerForWebhookIDs := *openapiclient.NewContainerForWebhookIDs([]int64{int64(123)}) // ContainerForWebhookIDs | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.RefreshWebhooks(context.Background()).ContainerForWebhookIDs(containerForWebhookIDs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.RefreshWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RefreshWebhooks`: WebhooksExpirationDate
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.RefreshWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRefreshWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **containerForWebhookIDs** | [**ContainerForWebhookIDs**](ContainerForWebhookIDs.md) |  | 

### Return type

[**WebhooksExpirationDate**](WebhooksExpirationDate.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDynamicWebhooks

> ContainerForRegisteredWebhooks RegisterDynamicWebhooks(ctx).WebhookRegistrationDetails(webhookRegistrationDetails).Execute()

Register dynamic webhooks



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
	webhookRegistrationDetails := *openapiclient.NewWebhookRegistrationDetails("Url_example", []openapiclient.WebhookDetails{*openapiclient.NewWebhookDetails([]string{"Events_example"}, "JqlFilter_example")}) // WebhookRegistrationDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WebhooksAPI.RegisterDynamicWebhooks(context.Background()).WebhookRegistrationDetails(webhookRegistrationDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WebhooksAPI.RegisterDynamicWebhooks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RegisterDynamicWebhooks`: ContainerForRegisteredWebhooks
	fmt.Fprintf(os.Stdout, "Response from `WebhooksAPI.RegisterDynamicWebhooks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDynamicWebhooksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookRegistrationDetails** | [**WebhookRegistrationDetails**](WebhookRegistrationDetails.md) |  | 

### Return type

[**ContainerForRegisteredWebhooks**](ContainerForRegisteredWebhooks.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

