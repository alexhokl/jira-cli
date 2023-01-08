# \IssueAttachmentsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddAttachment**](IssueAttachmentsAPI.md#AddAttachment) | **Post** /rest/api/3/issue/{issueIdOrKey}/attachments | Add attachment
[**ExpandAttachmentForHumans**](IssueAttachmentsAPI.md#ExpandAttachmentForHumans) | **Get** /rest/api/3/attachment/{id}/expand/human | Get all metadata for an expanded attachment
[**ExpandAttachmentForMachines**](IssueAttachmentsAPI.md#ExpandAttachmentForMachines) | **Get** /rest/api/3/attachment/{id}/expand/raw | Get contents metadata for an expanded attachment
[**GetAttachment**](IssueAttachmentsAPI.md#GetAttachment) | **Get** /rest/api/3/attachment/{id} | Get attachment metadata
[**GetAttachmentContent**](IssueAttachmentsAPI.md#GetAttachmentContent) | **Get** /rest/api/3/attachment/content/{id} | Get attachment content
[**GetAttachmentMeta**](IssueAttachmentsAPI.md#GetAttachmentMeta) | **Get** /rest/api/3/attachment/meta | Get Jira attachment settings
[**GetAttachmentThumbnail**](IssueAttachmentsAPI.md#GetAttachmentThumbnail) | **Get** /rest/api/3/attachment/thumbnail/{id} | Get attachment thumbnail
[**RemoveAttachment**](IssueAttachmentsAPI.md#RemoveAttachment) | **Delete** /rest/api/3/attachment/{id} | Delete attachment



## AddAttachment

> []Attachment AddAttachment(ctx, issueIdOrKey).Execute()

Add attachment



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
	issueIdOrKey := "issueIdOrKey_example" // string | The ID or key of the issue that attachments are added to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAttachmentsAPI.AddAttachment(context.Background(), issueIdOrKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.AddAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddAttachment`: []Attachment
	fmt.Fprintf(os.Stdout, "Response from `IssueAttachmentsAPI.AddAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**issueIdOrKey** | **string** | The ID or key of the issue that attachments are added to. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]Attachment**](Attachment.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandAttachmentForHumans

> AttachmentArchiveMetadataReadable ExpandAttachmentForHumans(ctx, id).Execute()

Get all metadata for an expanded attachment



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
	id := "id_example" // string | The ID of the attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAttachmentsAPI.ExpandAttachmentForHumans(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.ExpandAttachmentForHumans``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpandAttachmentForHumans`: AttachmentArchiveMetadataReadable
	fmt.Fprintf(os.Stdout, "Response from `IssueAttachmentsAPI.ExpandAttachmentForHumans`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpandAttachmentForHumansRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentArchiveMetadataReadable**](AttachmentArchiveMetadataReadable.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ExpandAttachmentForMachines

> AttachmentArchiveImpl ExpandAttachmentForMachines(ctx, id).Execute()

Get contents metadata for an expanded attachment



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
	id := "id_example" // string | The ID of the attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAttachmentsAPI.ExpandAttachmentForMachines(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.ExpandAttachmentForMachines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExpandAttachmentForMachines`: AttachmentArchiveImpl
	fmt.Fprintf(os.Stdout, "Response from `IssueAttachmentsAPI.ExpandAttachmentForMachines`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiExpandAttachmentForMachinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentArchiveImpl**](AttachmentArchiveImpl.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachment

> AttachmentMetadata GetAttachment(ctx, id).Execute()

Get attachment metadata



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
	id := "id_example" // string | The ID of the attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueAttachmentsAPI.GetAttachment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.GetAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachment`: AttachmentMetadata
	fmt.Fprintf(os.Stdout, "Response from `IssueAttachmentsAPI.GetAttachment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AttachmentMetadata**](AttachmentMetadata.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentContent

> GetAttachmentContent(ctx, id).Redirect(redirect).Execute()

Get attachment content



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
	id := "id_example" // string | The ID of the attachment.
	redirect := true // bool | Whether a redirect is provided for the attachment download. Clients that do not automatically follow redirects can set this to `false` to avoid making multiple requests to download the attachment. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAttachmentsAPI.GetAttachmentContent(context.Background(), id).Redirect(redirect).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.GetAttachmentContent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentContentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirect** | **bool** | Whether a redirect is provided for the attachment download. Clients that do not automatically follow redirects can set this to &#x60;false&#x60; to avoid making multiple requests to download the attachment. | [default to true]

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


## GetAttachmentMeta

> AttachmentSettings GetAttachmentMeta(ctx).Execute()

Get Jira attachment settings



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
	resp, r, err := apiClient.IssueAttachmentsAPI.GetAttachmentMeta(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.GetAttachmentMeta``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAttachmentMeta`: AttachmentSettings
	fmt.Fprintf(os.Stdout, "Response from `IssueAttachmentsAPI.GetAttachmentMeta`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentMetaRequest struct via the builder pattern


### Return type

[**AttachmentSettings**](AttachmentSettings.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAttachmentThumbnail

> GetAttachmentThumbnail(ctx, id).Redirect(redirect).FallbackToDefault(fallbackToDefault).Width(width).Height(height).Execute()

Get attachment thumbnail



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
	id := "id_example" // string | The ID of the attachment.
	redirect := true // bool | Whether a redirect is provided for the attachment download. Clients that do not automatically follow redirects can set this to `false` to avoid making multiple requests to download the attachment. (optional) (default to true)
	fallbackToDefault := true // bool | Whether a default thumbnail is returned when the requested thumbnail is not found. (optional) (default to true)
	width := int32(56) // int32 | The maximum width to scale the thumbnail to. (optional)
	height := int32(56) // int32 | The maximum height to scale the thumbnail to. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAttachmentsAPI.GetAttachmentThumbnail(context.Background(), id).Redirect(redirect).FallbackToDefault(fallbackToDefault).Width(width).Height(height).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.GetAttachmentThumbnail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAttachmentThumbnailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **redirect** | **bool** | Whether a redirect is provided for the attachment download. Clients that do not automatically follow redirects can set this to &#x60;false&#x60; to avoid making multiple requests to download the attachment. | [default to true]
 **fallbackToDefault** | **bool** | Whether a default thumbnail is returned when the requested thumbnail is not found. | [default to true]
 **width** | **int32** | The maximum width to scale the thumbnail to. | 
 **height** | **int32** | The maximum height to scale the thumbnail to. | 

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


## RemoveAttachment

> RemoveAttachment(ctx, id).Execute()

Delete attachment



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
	id := "id_example" // string | The ID of the attachment.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueAttachmentsAPI.RemoveAttachment(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueAttachmentsAPI.RemoveAttachment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the attachment. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAttachmentRequest struct via the builder pattern


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

