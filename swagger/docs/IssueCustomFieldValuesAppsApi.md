# \IssueCustomFieldValuesAppsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateCustomFieldValue**](IssueCustomFieldValuesAppsAPI.md#UpdateCustomFieldValue) | **Put** /rest/api/3/app/field/{fieldIdOrKey}/value | Update custom field value
[**UpdateMultipleCustomFieldValues**](IssueCustomFieldValuesAppsAPI.md#UpdateMultipleCustomFieldValues) | **Post** /rest/api/3/app/field/value | Update custom fields



## UpdateCustomFieldValue

> interface{} UpdateCustomFieldValue(ctx, fieldIdOrKey).CustomFieldValueUpdateDetails(customFieldValueUpdateDetails).GenerateChangelog(generateChangelog).Execute()

Update custom field value



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
	fieldIdOrKey := "fieldIdOrKey_example" // string | The ID or key of the custom field. For example, `customfield_10010`.
	customFieldValueUpdateDetails := *openapiclient.NewCustomFieldValueUpdateDetails() // CustomFieldValueUpdateDetails | 
	generateChangelog := true // bool | Whether to generate a changelog for this update. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldValuesAppsAPI.UpdateCustomFieldValue(context.Background(), fieldIdOrKey).CustomFieldValueUpdateDetails(customFieldValueUpdateDetails).GenerateChangelog(generateChangelog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldValuesAppsAPI.UpdateCustomFieldValue``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateCustomFieldValue`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldValuesAppsAPI.UpdateCustomFieldValue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldIdOrKey** | **string** | The ID or key of the custom field. For example, &#x60;customfield_10010&#x60;. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomFieldValueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **customFieldValueUpdateDetails** | [**CustomFieldValueUpdateDetails**](CustomFieldValueUpdateDetails.md) |  | 
 **generateChangelog** | **bool** | Whether to generate a changelog for this update. | [default to true]

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


## UpdateMultipleCustomFieldValues

> interface{} UpdateMultipleCustomFieldValues(ctx).MultipleCustomFieldValuesUpdateDetails(multipleCustomFieldValuesUpdateDetails).GenerateChangelog(generateChangelog).Execute()

Update custom fields



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
	multipleCustomFieldValuesUpdateDetails := *openapiclient.NewMultipleCustomFieldValuesUpdateDetails() // MultipleCustomFieldValuesUpdateDetails | 
	generateChangelog := true // bool | Whether to generate a changelog for this update. (optional) (default to true)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldValuesAppsAPI.UpdateMultipleCustomFieldValues(context.Background()).MultipleCustomFieldValuesUpdateDetails(multipleCustomFieldValuesUpdateDetails).GenerateChangelog(generateChangelog).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldValuesAppsAPI.UpdateMultipleCustomFieldValues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMultipleCustomFieldValues`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldValuesAppsAPI.UpdateMultipleCustomFieldValues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMultipleCustomFieldValuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **multipleCustomFieldValuesUpdateDetails** | [**MultipleCustomFieldValuesUpdateDetails**](MultipleCustomFieldValuesUpdateDetails.md) |  | 
 **generateChangelog** | **bool** | Whether to generate a changelog for this update. | [default to true]

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

