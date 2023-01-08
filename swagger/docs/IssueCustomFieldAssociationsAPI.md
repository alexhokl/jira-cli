# \IssueCustomFieldAssociationsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssociations**](IssueCustomFieldAssociationsAPI.md#CreateAssociations) | **Put** /rest/api/3/field/association | Create associations
[**RemoveAssociations**](IssueCustomFieldAssociationsAPI.md#RemoveAssociations) | **Delete** /rest/api/3/field/association | Remove associations



## CreateAssociations

> interface{} CreateAssociations(ctx).FieldAssociationsRequest(fieldAssociationsRequest).Execute()

Create associations



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
	fieldAssociationsRequest := *openapiclient.NewFieldAssociationsRequest([]openapiclient.AssociationContextObject{*openapiclient.NewAssociationContextObject("Type_example")}, []openapiclient.FieldIdentifierObject{*openapiclient.NewFieldIdentifierObject("Type_example")}) // FieldAssociationsRequest | Payload containing the fields to associate and the projects to associate them to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldAssociationsAPI.CreateAssociations(context.Background()).FieldAssociationsRequest(fieldAssociationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldAssociationsAPI.CreateAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateAssociations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldAssociationsAPI.CreateAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldAssociationsRequest** | [**FieldAssociationsRequest**](FieldAssociationsRequest.md) | Payload containing the fields to associate and the projects to associate them to. | 

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


## RemoveAssociations

> interface{} RemoveAssociations(ctx).FieldAssociationsRequest(fieldAssociationsRequest).Execute()

Remove associations



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
	fieldAssociationsRequest := *openapiclient.NewFieldAssociationsRequest([]openapiclient.AssociationContextObject{*openapiclient.NewAssociationContextObject("Type_example")}, []openapiclient.FieldIdentifierObject{*openapiclient.NewFieldIdentifierObject("Type_example")}) // FieldAssociationsRequest | Payload containing the fields to uassociate and the projects and issue types to unassociate them to.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldAssociationsAPI.RemoveAssociations(context.Background()).FieldAssociationsRequest(fieldAssociationsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldAssociationsAPI.RemoveAssociations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveAssociations`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldAssociationsAPI.RemoveAssociations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveAssociationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldAssociationsRequest** | [**FieldAssociationsRequest**](FieldAssociationsRequest.md) | Payload containing the fields to uassociate and the projects and issue types to unassociate them to. | 

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

