# \ClassificationLevelsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAllUserDataClassificationLevels**](ClassificationLevelsAPI.md#GetAllUserDataClassificationLevels) | **Get** /rest/api/3/classification-levels | Get all classification levels



## GetAllUserDataClassificationLevels

> DataClassificationLevelsBean GetAllUserDataClassificationLevels(ctx).Status(status).OrderBy(orderBy).Execute()

Get all classification levels



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
	status := []string{"Status_example"} // []string | Optional set of statuses to filter by. (optional)
	orderBy := "orderBy_example" // string | Ordering of the results by a given field. If not provided, values will not be sorted. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ClassificationLevelsAPI.GetAllUserDataClassificationLevels(context.Background()).Status(status).OrderBy(orderBy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ClassificationLevelsAPI.GetAllUserDataClassificationLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllUserDataClassificationLevels`: DataClassificationLevelsBean
	fmt.Fprintf(os.Stdout, "Response from `ClassificationLevelsAPI.GetAllUserDataClassificationLevels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllUserDataClassificationLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **[]string** | Optional set of statuses to filter by. | 
 **orderBy** | **string** | Ordering of the results by a given field. If not provided, values will not be sorted. | 

### Return type

[**DataClassificationLevelsBean**](DataClassificationLevelsBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

