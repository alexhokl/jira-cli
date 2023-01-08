# \AuditRecordsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAuditRecords**](AuditRecordsAPI.md#GetAuditRecords) | **Get** /rest/api/3/auditing/record | Get audit records



## GetAuditRecords

> AuditRecords GetAuditRecords(ctx).Offset(offset).Limit(limit).Filter(filter).From(from).To(to).Execute()

Get audit records



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
	offset := int32(56) // int32 | The number of records to skip before returning the first result. (optional) (default to 0)
	limit := int32(56) // int32 | The maximum number of results to return. (optional) (default to 1000)
	filter := "filter_example" // string | The strings to match with audit field content, space separated. (optional)
	from := "from_example" // string | The date and time on or after which returned audit records must have been created. If `to` is provided `from` must be before `to` or no audit records are returned. (optional)
	to := "to_example" // string | The date and time on or before which returned audit results must have been created. If `from` is provided `to` must be after `from` or no audit records are returned. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AuditRecordsAPI.GetAuditRecords(context.Background()).Offset(offset).Limit(limit).Filter(filter).From(from).To(to).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AuditRecordsAPI.GetAuditRecords``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAuditRecords`: AuditRecords
	fmt.Fprintf(os.Stdout, "Response from `AuditRecordsAPI.GetAuditRecords`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAuditRecordsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **offset** | **int32** | The number of records to skip before returning the first result. | [default to 0]
 **limit** | **int32** | The maximum number of results to return. | [default to 1000]
 **filter** | **string** | The strings to match with audit field content, space separated. | 
 **from** | **string** | The date and time on or after which returned audit records must have been created. If &#x60;to&#x60; is provided &#x60;from&#x60; must be before &#x60;to&#x60; or no audit records are returned. | 
 **to** | **string** | The date and time on or before which returned audit results must have been created. If &#x60;from&#x60; is provided &#x60;to&#x60; must be after &#x60;from&#x60; or no audit records are returned. | 

### Return type

[**AuditRecords**](AuditRecords.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

