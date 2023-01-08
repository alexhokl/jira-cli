# \MigrationOfConnectModulesToForgeAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet**](MigrationOfConnectModulesToForgeAPI.md#ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet) | **Get** /rest/atlassian-connect/1/migration/{connectKey}/{jiraIssueFieldsKey}/task | Get Connect issue field migration task



## ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet

> TaskProgress ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet(ctx, connectKey, jiraIssueFieldsKey).Execute()

Get Connect issue field migration task



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
	connectKey := "com.example.app" // string | The key of the Connect app that contains the Jira issue field being migrated.
	jiraIssueFieldsKey := "my-custom-field" // string | The module key of the Connect issue field being migrated.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MigrationOfConnectModulesToForgeAPI.ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet(context.Background(), connectKey, jiraIssueFieldsKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MigrationOfConnectModulesToForgeAPI.ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet`: TaskProgress
	fmt.Fprintf(os.Stdout, "Response from `MigrationOfConnectModulesToForgeAPI.ConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connectKey** | **string** | The key of the Connect app that contains the Jira issue field being migrated. | 
**jiraIssueFieldsKey** | **string** | The module key of the Connect issue field being migrated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiConnectToForgeMigrationFetchTaskResourceFetchMigrationTaskGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**TaskProgress**](TaskProgress.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

