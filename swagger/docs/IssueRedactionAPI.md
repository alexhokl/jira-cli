# \IssueRedactionAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetRedactionStatus**](IssueRedactionAPI.md#GetRedactionStatus) | **Get** /rest/api/3/redact/status/{jobId} | Get redaction status
[**Redact**](IssueRedactionAPI.md#Redact) | **Post** /rest/api/3/redact | Redact



## GetRedactionStatus

> RedactionJobStatusResponse GetRedactionStatus(ctx, jobId).Execute()

Get redaction status



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
	jobId := "jobId_example" // string | Redaction job id

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRedactionAPI.GetRedactionStatus(context.Background(), jobId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRedactionAPI.GetRedactionStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRedactionStatus`: RedactionJobStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `IssueRedactionAPI.GetRedactionStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** | Redaction job id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRedactionStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**RedactionJobStatusResponse**](RedactionJobStatusResponse.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## Redact

> string Redact(ctx).BulkRedactionRequest(bulkRedactionRequest).Execute()

Redact



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
	bulkRedactionRequest := *openapiclient.NewBulkRedactionRequest() // BulkRedactionRequest | List of redaction requests

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueRedactionAPI.Redact(context.Background()).BulkRedactionRequest(bulkRedactionRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueRedactionAPI.Redact``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `Redact`: string
	fmt.Fprintf(os.Stdout, "Response from `IssueRedactionAPI.Redact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRedactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkRedactionRequest** | [**BulkRedactionRequest**](BulkRedactionRequest.md) | List of redaction requests | 

### Return type

**string**

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

