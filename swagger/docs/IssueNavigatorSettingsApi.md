# \IssueNavigatorSettingsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetIssueNavigatorDefaultColumns**](IssueNavigatorSettingsAPI.md#GetIssueNavigatorDefaultColumns) | **Get** /rest/api/3/settings/columns | Get issue navigator default columns
[**SetIssueNavigatorDefaultColumns**](IssueNavigatorSettingsAPI.md#SetIssueNavigatorDefaultColumns) | **Put** /rest/api/3/settings/columns | Set issue navigator default columns



## GetIssueNavigatorDefaultColumns

> []ColumnItem GetIssueNavigatorDefaultColumns(ctx).Execute()

Get issue navigator default columns



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
	resp, r, err := apiClient.IssueNavigatorSettingsAPI.GetIssueNavigatorDefaultColumns(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNavigatorSettingsAPI.GetIssueNavigatorDefaultColumns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueNavigatorDefaultColumns`: []ColumnItem
	fmt.Fprintf(os.Stdout, "Response from `IssueNavigatorSettingsAPI.GetIssueNavigatorDefaultColumns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueNavigatorDefaultColumnsRequest struct via the builder pattern


### Return type

[**[]ColumnItem**](ColumnItem.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetIssueNavigatorDefaultColumns

> SetIssueNavigatorDefaultColumns(ctx).ColumnRequestBody(columnRequestBody).Execute()

Set issue navigator default columns



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
	columnRequestBody := *openapiclient.NewColumnRequestBody() // ColumnRequestBody | A navigable field value.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueNavigatorSettingsAPI.SetIssueNavigatorDefaultColumns(context.Background()).ColumnRequestBody(columnRequestBody).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueNavigatorSettingsAPI.SetIssueNavigatorDefaultColumns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetIssueNavigatorDefaultColumnsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **columnRequestBody** | [**ColumnRequestBody**](ColumnRequestBody.md) | A navigable field value. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

