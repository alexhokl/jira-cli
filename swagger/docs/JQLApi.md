# \JQLAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAutoComplete**](JQLAPI.md#GetAutoComplete) | **Get** /rest/api/3/jql/autocompletedata | Get field reference data (GET)
[**GetAutoCompletePost**](JQLAPI.md#GetAutoCompletePost) | **Post** /rest/api/3/jql/autocompletedata | Get field reference data (POST)
[**GetFieldAutoCompleteForQueryString**](JQLAPI.md#GetFieldAutoCompleteForQueryString) | **Get** /rest/api/3/jql/autocompletedata/suggestions | Get field auto complete suggestions
[**MigrateQueries**](JQLAPI.md#MigrateQueries) | **Post** /rest/api/3/jql/pdcleaner | Convert user identifiers to account IDs in JQL queries
[**ParseJqlQueries**](JQLAPI.md#ParseJqlQueries) | **Post** /rest/api/3/jql/parse | Parse JQL query
[**SanitiseJqlQueries**](JQLAPI.md#SanitiseJqlQueries) | **Post** /rest/api/3/jql/sanitize | Sanitize JQL queries



## GetAutoComplete

> JQLReferenceData GetAutoComplete(ctx).Execute()

Get field reference data (GET)



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
	resp, r, err := apiClient.JQLAPI.GetAutoComplete(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.GetAutoComplete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoComplete`: JQLReferenceData
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.GetAutoComplete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoCompleteRequest struct via the builder pattern


### Return type

[**JQLReferenceData**](JQLReferenceData.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAutoCompletePost

> JQLReferenceData GetAutoCompletePost(ctx).SearchAutoCompleteFilter(searchAutoCompleteFilter).Execute()

Get field reference data (POST)



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
	searchAutoCompleteFilter := *openapiclient.NewSearchAutoCompleteFilter() // SearchAutoCompleteFilter | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLAPI.GetAutoCompletePost(context.Background()).SearchAutoCompleteFilter(searchAutoCompleteFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.GetAutoCompletePost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAutoCompletePost`: JQLReferenceData
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.GetAutoCompletePost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAutoCompletePostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchAutoCompleteFilter** | [**SearchAutoCompleteFilter**](SearchAutoCompleteFilter.md) |  | 

### Return type

[**JQLReferenceData**](JQLReferenceData.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFieldAutoCompleteForQueryString

> AutoCompleteSuggestions GetFieldAutoCompleteForQueryString(ctx).FieldName(fieldName).FieldValue(fieldValue).PredicateName(predicateName).PredicateValue(predicateValue).Execute()

Get field auto complete suggestions



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
	fieldName := "reporter" // string | The name of the field. (optional)
	fieldValue := "fieldValue_example" // string | The partial field item name entered by the user. (optional)
	predicateName := "predicateName_example" // string | The name of the [ CHANGED operator predicate](https://confluence.atlassian.com/x/hQORLQ#Advancedsearching-operatorsreference-CHANGEDCHANGED) for which the suggestions are generated. The valid predicate operators are *by*, *from*, and *to*. (optional)
	predicateValue := "predicateValue_example" // string | The partial predicate item name entered by the user. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLAPI.GetFieldAutoCompleteForQueryString(context.Background()).FieldName(fieldName).FieldValue(fieldValue).PredicateName(predicateName).PredicateValue(predicateValue).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.GetFieldAutoCompleteForQueryString``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFieldAutoCompleteForQueryString`: AutoCompleteSuggestions
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.GetFieldAutoCompleteForQueryString`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldAutoCompleteForQueryStringRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fieldName** | **string** | The name of the field. | 
 **fieldValue** | **string** | The partial field item name entered by the user. | 
 **predicateName** | **string** | The name of the [ CHANGED operator predicate](https://confluence.atlassian.com/x/hQORLQ#Advancedsearching-operatorsreference-CHANGEDCHANGED) for which the suggestions are generated. The valid predicate operators are *by*, *from*, and *to*. | 
 **predicateValue** | **string** | The partial predicate item name entered by the user. | 

### Return type

[**AutoCompleteSuggestions**](AutoCompleteSuggestions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MigrateQueries

> ConvertedJQLQueries MigrateQueries(ctx).JQLPersonalDataMigrationRequest(jQLPersonalDataMigrationRequest).Execute()

Convert user identifiers to account IDs in JQL queries



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
	jQLPersonalDataMigrationRequest := *openapiclient.NewJQLPersonalDataMigrationRequest() // JQLPersonalDataMigrationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLAPI.MigrateQueries(context.Background()).JQLPersonalDataMigrationRequest(jQLPersonalDataMigrationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.MigrateQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MigrateQueries`: ConvertedJQLQueries
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.MigrateQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMigrateQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jQLPersonalDataMigrationRequest** | [**JQLPersonalDataMigrationRequest**](JQLPersonalDataMigrationRequest.md) |  | 

### Return type

[**ConvertedJQLQueries**](ConvertedJQLQueries.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ParseJqlQueries

> ParsedJqlQueries ParseJqlQueries(ctx).Validation(validation).JqlQueriesToParse(jqlQueriesToParse).Execute()

Parse JQL query



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
	validation := "validation_example" // string | How to validate the JQL query and treat the validation results. Validation options include:   *  `strict` Returns all errors. If validation fails, the query structure is not returned.  *  `warn` Returns all errors. If validation fails but the JQL query is correctly formed, the query structure is returned.  *  `none` No validation is performed. If JQL query is correctly formed, the query structure is returned. (default to "strict")
	jqlQueriesToParse := *openapiclient.NewJqlQueriesToParse([]string{"Queries_example"}) // JqlQueriesToParse | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLAPI.ParseJqlQueries(context.Background()).Validation(validation).JqlQueriesToParse(jqlQueriesToParse).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.ParseJqlQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ParseJqlQueries`: ParsedJqlQueries
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.ParseJqlQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiParseJqlQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **validation** | **string** | How to validate the JQL query and treat the validation results. Validation options include:   *  &#x60;strict&#x60; Returns all errors. If validation fails, the query structure is not returned.  *  &#x60;warn&#x60; Returns all errors. If validation fails but the JQL query is correctly formed, the query structure is returned.  *  &#x60;none&#x60; No validation is performed. If JQL query is correctly formed, the query structure is returned. | [default to &quot;strict&quot;]
 **jqlQueriesToParse** | [**JqlQueriesToParse**](JqlQueriesToParse.md) |  | 

### Return type

[**ParsedJqlQueries**](ParsedJqlQueries.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SanitiseJqlQueries

> SanitizedJqlQueries SanitiseJqlQueries(ctx).JqlQueriesToSanitize(jqlQueriesToSanitize).Execute()

Sanitize JQL queries



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
	jqlQueriesToSanitize := *openapiclient.NewJqlQueriesToSanitize([]openapiclient.JqlQueryToSanitize{*openapiclient.NewJqlQueryToSanitize("Query_example")}) // JqlQueriesToSanitize | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JQLAPI.SanitiseJqlQueries(context.Background()).JqlQueriesToSanitize(jqlQueriesToSanitize).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JQLAPI.SanitiseJqlQueries``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SanitiseJqlQueries`: SanitizedJqlQueries
	fmt.Fprintf(os.Stdout, "Response from `JQLAPI.SanitiseJqlQueries`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSanitiseJqlQueriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jqlQueriesToSanitize** | [**JqlQueriesToSanitize**](JqlQueriesToSanitize.md) |  | 

### Return type

[**SanitizedJqlQueries**](SanitizedJqlQueries.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

