# \DefaultAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorklogsByIssueIdAndWorklogId**](DefaultAPI.md#GetWorklogsByIssueIdAndWorklogId) | **Post** /rest/internal/api/latest/worklog/bulk | Get worklogs by issue id and worklog id



## GetWorklogsByIssueIdAndWorklogId

> BulkWorklogKeyResponseBean GetWorklogsByIssueIdAndWorklogId(ctx).BulkWorklogKeyRequestBean(bulkWorklogKeyRequestBean).Execute()

Get worklogs by issue id and worklog id



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
	bulkWorklogKeyRequestBean := *openapiclient.NewBulkWorklogKeyRequestBean() // BulkWorklogKeyRequestBean | A JSON object containing a list of issue ID and worklog ID pairs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DefaultAPI.GetWorklogsByIssueIdAndWorklogId(context.Background()).BulkWorklogKeyRequestBean(bulkWorklogKeyRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DefaultAPI.GetWorklogsByIssueIdAndWorklogId``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWorklogsByIssueIdAndWorklogId`: BulkWorklogKeyResponseBean
	fmt.Fprintf(os.Stdout, "Response from `DefaultAPI.GetWorklogsByIssueIdAndWorklogId`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetWorklogsByIssueIdAndWorklogIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkWorklogKeyRequestBean** | [**BulkWorklogKeyRequestBean**](BulkWorklogKeyRequestBean.md) | A JSON object containing a list of issue ID and worklog ID pairs. | 

### Return type

[**BulkWorklogKeyResponseBean**](BulkWorklogKeyResponseBean.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

