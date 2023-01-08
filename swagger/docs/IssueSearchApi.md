# \IssueSearchAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountIssues**](IssueSearchAPI.md#CountIssues) | **Post** /rest/api/3/search/approximate-count | Count issues using JQL
[**GetIssuePickerResource**](IssueSearchAPI.md#GetIssuePickerResource) | **Get** /rest/api/3/issue/picker | Get issue picker suggestions
[**MatchIssues**](IssueSearchAPI.md#MatchIssues) | **Post** /rest/api/3/jql/match | Check issues against JQL
[**SearchAndReconsileIssuesUsingJql**](IssueSearchAPI.md#SearchAndReconsileIssuesUsingJql) | **Get** /rest/api/3/search/jql | Search for issues using JQL enhanced search (GET)
[**SearchAndReconsileIssuesUsingJqlPost**](IssueSearchAPI.md#SearchAndReconsileIssuesUsingJqlPost) | **Post** /rest/api/3/search/jql | Search for issues using JQL enhanced search (POST)
[**SearchForIssuesUsingJql**](IssueSearchAPI.md#SearchForIssuesUsingJql) | **Get** /rest/api/3/search | Currently being removed. Search for issues using JQL (GET)
[**SearchForIssuesUsingJqlPost**](IssueSearchAPI.md#SearchForIssuesUsingJqlPost) | **Post** /rest/api/3/search | Currently being removed. Search for issues using JQL (POST)



## CountIssues

> JQLCountResultsBean CountIssues(ctx).JQLCountRequestBean(jQLCountRequestBean).Execute()

Count issues using JQL



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
	jQLCountRequestBean := *openapiclient.NewJQLCountRequestBean() // JQLCountRequestBean | A JSON object containing the search request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.CountIssues(context.Background()).JQLCountRequestBean(jQLCountRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.CountIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CountIssues`: JQLCountResultsBean
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.CountIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jQLCountRequestBean** | [**JQLCountRequestBean**](JQLCountRequestBean.md) | A JSON object containing the search request. | 

### Return type

[**JQLCountResultsBean**](JQLCountResultsBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssuePickerResource

> IssuePickerSuggestions GetIssuePickerResource(ctx).Query(query).CurrentJQL(currentJQL).CurrentIssueKey(currentIssueKey).CurrentProjectId(currentProjectId).ShowSubTasks(showSubTasks).ShowSubTaskParent(showSubTaskParent).Execute()

Get issue picker suggestions



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
	query := "query" // string | A string to match against text fields in the issue such as title, description, or comments. (optional)
	currentJQL := "currentJQL_example" // string | A JQL query defining a list of issues to search for the query term. Note that `username` and `userkey` cannot be used as search terms for this parameter, due to privacy reasons. Use `accountId` instead. (optional)
	currentIssueKey := "currentIssueKey_example" // string | The key of an issue to exclude from search results. For example, the issue the user is viewing when they perform this query. (optional)
	currentProjectId := "currentProjectId_example" // string | The ID of a project that suggested issues must belong to. (optional)
	showSubTasks := true // bool | Indicate whether to include subtasks in the suggestions list. (optional)
	showSubTaskParent := true // bool | When `currentIssueKey` is a subtask, whether to include the parent issue in the suggestions if it matches the query. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.GetIssuePickerResource(context.Background()).Query(query).CurrentJQL(currentJQL).CurrentIssueKey(currentIssueKey).CurrentProjectId(currentProjectId).ShowSubTasks(showSubTasks).ShowSubTaskParent(showSubTaskParent).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.GetIssuePickerResource``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssuePickerResource`: IssuePickerSuggestions
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.GetIssuePickerResource`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetIssuePickerResourceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **query** | **string** | A string to match against text fields in the issue such as title, description, or comments. | 
 **currentJQL** | **string** | A JQL query defining a list of issues to search for the query term. Note that &#x60;username&#x60; and &#x60;userkey&#x60; cannot be used as search terms for this parameter, due to privacy reasons. Use &#x60;accountId&#x60; instead. | 
 **currentIssueKey** | **string** | The key of an issue to exclude from search results. For example, the issue the user is viewing when they perform this query. | 
 **currentProjectId** | **string** | The ID of a project that suggested issues must belong to. | 
 **showSubTasks** | **bool** | Indicate whether to include subtasks in the suggestions list. | 
 **showSubTaskParent** | **bool** | When &#x60;currentIssueKey&#x60; is a subtask, whether to include the parent issue in the suggestions if it matches the query. | 

### Return type

[**IssuePickerSuggestions**](IssuePickerSuggestions.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MatchIssues

> IssueMatches MatchIssues(ctx).IssuesAndJQLQueries(issuesAndJQLQueries).Execute()

Check issues against JQL



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
	issuesAndJQLQueries := *openapiclient.NewIssuesAndJQLQueries([]int64{int64(123)}, []string{"Jqls_example"}) // IssuesAndJQLQueries | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.MatchIssues(context.Background()).IssuesAndJQLQueries(issuesAndJQLQueries).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.MatchIssues``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `MatchIssues`: IssueMatches
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.MatchIssues`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMatchIssuesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **issuesAndJQLQueries** | [**IssuesAndJQLQueries**](IssuesAndJQLQueries.md) |  | 

### Return type

[**IssueMatches**](IssueMatches.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAndReconsileIssuesUsingJql

> SearchAndReconcileResults SearchAndReconsileIssuesUsingJql(ctx).Jql(jql).NextPageToken(nextPageToken).MaxResults(maxResults).Fields(fields).Expand(expand).Properties(properties).FieldsByKeys(fieldsByKeys).FailFast(failFast).ReconcileIssues(reconcileIssues).Execute()

Search for issues using JQL enhanced search (GET)



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
	jql := "project = HSP" // string | A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. For performance reasons, this parameter requires a bounded query. A bounded query is a query with a search restriction.   *  Example of an unbounded query: `order by key desc`.  *  Example of a bounded query: `assignee = currentUser() order by key`.  Additionally, `orderBy` clause can contain a maximum of 7 fields. (optional)
	nextPageToken := "<string>" // string | The token for a page to fetch that is not the first page. The first page has a `nextPageToken` of `null`. Use the `nextPageToken` to fetch the next page of issues.  Note: The `nextPageToken` field is **not included** in the response for the last page, indicating there is no next page. (optional)
	maxResults := int32(114) // int32 | The maximum number of items to return per page. To manage page size, API may return fewer items per page where a large number of fields or properties are requested. The greatest number of items returned per page is achieved when requesting `id` or `key` only. It returns max 5000 issues. (optional) (default to 50)
	fields := []string{"Inner_example"} // []string | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  `*all` Returns all fields.  *  `*navigable` Returns navigable fields.  *  `id` Returns only issue IDs.  *  Any issue field, prefixed with a minus to exclude.  The default is `id`.  Examples:   *  `summary,comment` Returns only the summary and comments fields only.  *  `-description` Returns all navigable (default) fields except description.  *  `*all,-comment` Returns all fields except comments.  Multiple `fields` parameters can be included in a request.  Note: By default, this resource returns IDs only. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. (optional)
	expand := "<string>" // string | Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where `expand` is specified, `expand` is defined as a comma-delimited string of values. The expand options are:   *  `renderedFields` Returns field values rendered in HTML format.  *  `names` Returns the display name of each field.  *  `schema` Returns the schema describing a field type.  *  `transitions` Returns all possible transitions for the issue.  *  `operations` Returns all possible operations for the issue.  *  `editmeta` Returns information about how each field can be edited.  *  `changelog` Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  `versionedRepresentations` Instead of `fields`, returns `versionedRepresentations` a JSON array containing each version of a field's value, with the highest numbered item representing the most recent version.  Examples: `\"names,changelog\"` Returns the display name of each field as well as a list of recent updates to an issue. (optional)
	properties := []string{"Inner_example"} // []string | A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list. (optional)
	fieldsByKeys := true // bool | Reference fields by their key (rather than ID). The default is `false`. (optional) (default to false)
	failFast := true // bool | Fail this request early if we can't retrieve all field data. (optional) (default to false)
	reconcileIssues := []int64{int64(123)} // []int64 | Strong consistency issue ids to be reconciled with search results. Accepts max 50 ids. This list of ids should be consistent with each paginated request across different pages. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.SearchAndReconsileIssuesUsingJql(context.Background()).Jql(jql).NextPageToken(nextPageToken).MaxResults(maxResults).Fields(fields).Expand(expand).Properties(properties).FieldsByKeys(fieldsByKeys).FailFast(failFast).ReconcileIssues(reconcileIssues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.SearchAndReconsileIssuesUsingJql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAndReconsileIssuesUsingJql`: SearchAndReconcileResults
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.SearchAndReconsileIssuesUsingJql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAndReconsileIssuesUsingJqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jql** | **string** | A [JQL](https://confluence.atlassian.com/x/egORLQ) expression. For performance reasons, this parameter requires a bounded query. A bounded query is a query with a search restriction.   *  Example of an unbounded query: &#x60;order by key desc&#x60;.  *  Example of a bounded query: &#x60;assignee &#x3D; currentUser() order by key&#x60;.  Additionally, &#x60;orderBy&#x60; clause can contain a maximum of 7 fields. | 
 **nextPageToken** | **string** | The token for a page to fetch that is not the first page. The first page has a &#x60;nextPageToken&#x60; of &#x60;null&#x60;. Use the &#x60;nextPageToken&#x60; to fetch the next page of issues.  Note: The &#x60;nextPageToken&#x60; field is **not included** in the response for the last page, indicating there is no next page. | 
 **maxResults** | **int32** | The maximum number of items to return per page. To manage page size, API may return fewer items per page where a large number of fields or properties are requested. The greatest number of items returned per page is achieved when requesting &#x60;id&#x60; or &#x60;key&#x60; only. It returns max 5000 issues. | [default to 50]
 **fields** | **[]string** | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  &#x60;id&#x60; Returns only issue IDs.  *  Any issue field, prefixed with a minus to exclude.  The default is &#x60;id&#x60;.  Examples:   *  &#x60;summary,comment&#x60; Returns only the summary and comments fields only.  *  &#x60;-description&#x60; Returns all navigable (default) fields except description.  *  &#x60;*all,-comment&#x60; Returns all fields except comments.  Multiple &#x60;fields&#x60; parameters can be included in a request.  Note: By default, this resource returns IDs only. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about issues in the response. Note that, unlike the majority of instances where &#x60;expand&#x60; is specified, &#x60;expand&#x60; is defined as a comma-delimited string of values. The expand options are:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;operations&#x60; Returns all possible operations for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  &#x60;versionedRepresentations&#x60; Instead of &#x60;fields&#x60;, returns &#x60;versionedRepresentations&#x60; a JSON array containing each version of a field&#39;s value, with the highest numbered item representing the most recent version.  Examples: &#x60;\&quot;names,changelog\&quot;&#x60; Returns the display name of each field as well as a list of recent updates to an issue. | 
 **properties** | **[]string** | A list of up to 5 issue properties to include in the results. This parameter accepts a comma-separated list. | 
 **fieldsByKeys** | **bool** | Reference fields by their key (rather than ID). The default is &#x60;false&#x60;. | [default to false]
 **failFast** | **bool** | Fail this request early if we can&#39;t retrieve all field data. | [default to false]
 **reconcileIssues** | **[]int64** | Strong consistency issue ids to be reconciled with search results. Accepts max 50 ids. This list of ids should be consistent with each paginated request across different pages. | 

### Return type

[**SearchAndReconcileResults**](SearchAndReconcileResults.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchAndReconsileIssuesUsingJqlPost

> SearchAndReconcileResults SearchAndReconsileIssuesUsingJqlPost(ctx).SearchAndReconcileRequestBean(searchAndReconcileRequestBean).Execute()

Search for issues using JQL enhanced search (POST)



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
	searchAndReconcileRequestBean := *openapiclient.NewSearchAndReconcileRequestBean() // SearchAndReconcileRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.SearchAndReconsileIssuesUsingJqlPost(context.Background()).SearchAndReconcileRequestBean(searchAndReconcileRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.SearchAndReconsileIssuesUsingJqlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchAndReconsileIssuesUsingJqlPost`: SearchAndReconcileResults
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.SearchAndReconsileIssuesUsingJqlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchAndReconsileIssuesUsingJqlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchAndReconcileRequestBean** | [**SearchAndReconcileRequestBean**](SearchAndReconcileRequestBean.md) |  | 

### Return type

[**SearchAndReconcileResults**](SearchAndReconcileResults.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForIssuesUsingJql

> SearchResults SearchForIssuesUsingJql(ctx).Jql(jql).StartAt(startAt).MaxResults(maxResults).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Properties(properties).FieldsByKeys(fieldsByKeys).FailFast(failFast).Execute()

Currently being removed. Search for issues using JQL (GET)



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
	jql := "project = HSP" // string | The [JQL](https://confluence.atlassian.com/x/egORLQ) that defines the search. Note:   *  If no JQL expression is provided, all issues are returned.  *  `username` and `userkey` cannot be used as search terms due to privacy reasons. Use `accountId` instead.  *  If a user has hidden their email address in their user profile, partial matches of the email address will not find the user. An exact match is required. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. To manage page size, Jira may return fewer items per page where a large number of fields or properties are requested. The greatest number of items returned per page is achieved when requesting `id` or `key` only. (optional) (default to 50)
	validateQuery := "validateQuery_example" // string | Determines how to validate the JQL query and treat the validation results. Supported values are:   *  `strict` Returns a 400 response code if any errors are found, along with a list of all errors (and warnings).  *  `warn` Returns all errors as warnings.  *  `none` No validation is performed.  *  `true` *Deprecated* A legacy synonym for `strict`.  *  `false` *Deprecated* A legacy synonym for `warn`.  Note: If the JQL is not correctly formed a 400 response code is returned, regardless of the `validateQuery` value. (optional) (default to "strict")
	fields := []string{"Inner_example"} // []string | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  `*all` Returns all fields.  *  `*navigable` Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  Examples:   *  `summary,comment` Returns only the summary and comments fields.  *  `-description` Returns all navigable (default) fields except description.  *  `*all,-comment` Returns all fields except comments.  This parameter may be specified multiple times. For example, `fields=field1,field2&fields=field3`.  Note: All navigable fields are returned by default. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. (optional)
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about issues in the response. This parameter accepts a comma-separated list. Expand options include:   *  `renderedFields` Returns field values rendered in HTML format.  *  `names` Returns the display name of each field.  *  `schema` Returns the schema describing a field type.  *  `transitions` Returns all possible transitions for the issue.  *  `operations` Returns all possible operations for the issue.  *  `editmeta` Returns information about how each field can be edited.  *  `changelog` Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  `versionedRepresentations` Instead of `fields`, returns `versionedRepresentations` a JSON array containing each version of a field's value, with the highest numbered item representing the most recent version. (optional)
	properties := []string{"Inner_example"} // []string | A list of issue property keys for issue properties to include in the results. This parameter accepts a comma-separated list. Multiple properties can also be provided using an ampersand separated list. For example, `properties=prop1,prop2&properties=prop3`. A maximum of 5 issue property keys can be specified. (optional)
	fieldsByKeys := true // bool | Reference fields by their key (rather than ID). (optional) (default to false)
	failFast := true // bool | Whether to fail the request quickly in case of an error while loading fields for an issue. For `failFast=true`, if one field fails, the entire operation fails. For `failFast=false`, the operation will continue even if a field fails. It will return a valid response, but without values for the failed field(s). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.SearchForIssuesUsingJql(context.Background()).Jql(jql).StartAt(startAt).MaxResults(maxResults).ValidateQuery(validateQuery).Fields(fields).Expand(expand).Properties(properties).FieldsByKeys(fieldsByKeys).FailFast(failFast).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.SearchForIssuesUsingJql``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForIssuesUsingJql`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.SearchForIssuesUsingJql`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForIssuesUsingJqlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **jql** | **string** | The [JQL](https://confluence.atlassian.com/x/egORLQ) that defines the search. Note:   *  If no JQL expression is provided, all issues are returned.  *  &#x60;username&#x60; and &#x60;userkey&#x60; cannot be used as search terms due to privacy reasons. Use &#x60;accountId&#x60; instead.  *  If a user has hidden their email address in their user profile, partial matches of the email address will not find the user. An exact match is required. | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. To manage page size, Jira may return fewer items per page where a large number of fields or properties are requested. The greatest number of items returned per page is achieved when requesting &#x60;id&#x60; or &#x60;key&#x60; only. | [default to 50]
 **validateQuery** | **string** | Determines how to validate the JQL query and treat the validation results. Supported values are:   *  &#x60;strict&#x60; Returns a 400 response code if any errors are found, along with a list of all errors (and warnings).  *  &#x60;warn&#x60; Returns all errors as warnings.  *  &#x60;none&#x60; No validation is performed.  *  &#x60;true&#x60; *Deprecated* A legacy synonym for &#x60;strict&#x60;.  *  &#x60;false&#x60; *Deprecated* A legacy synonym for &#x60;warn&#x60;.  Note: If the JQL is not correctly formed a 400 response code is returned, regardless of the &#x60;validateQuery&#x60; value. | [default to &quot;strict&quot;]
 **fields** | **[]string** | A list of fields to return for each issue, use it to retrieve a subset of fields. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;*all&#x60; Returns all fields.  *  &#x60;*navigable&#x60; Returns navigable fields.  *  Any issue field, prefixed with a minus to exclude.  Examples:   *  &#x60;summary,comment&#x60; Returns only the summary and comments fields.  *  &#x60;-description&#x60; Returns all navigable (default) fields except description.  *  &#x60;*all,-comment&#x60; Returns all fields except comments.  This parameter may be specified multiple times. For example, &#x60;fields&#x3D;field1,field2&amp;fields&#x3D;field3&#x60;.  Note: All navigable fields are returned by default. This differs from [GET issue](#api-rest-api-3-issue-issueIdOrKey-get) where the default is all fields. | 
 **expand** | **string** | Use [expand](#expansion) to include additional information about issues in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;renderedFields&#x60; Returns field values rendered in HTML format.  *  &#x60;names&#x60; Returns the display name of each field.  *  &#x60;schema&#x60; Returns the schema describing a field type.  *  &#x60;transitions&#x60; Returns all possible transitions for the issue.  *  &#x60;operations&#x60; Returns all possible operations for the issue.  *  &#x60;editmeta&#x60; Returns information about how each field can be edited.  *  &#x60;changelog&#x60; Returns a list of recent updates to an issue, sorted by date, starting from the most recent.  *  &#x60;versionedRepresentations&#x60; Instead of &#x60;fields&#x60;, returns &#x60;versionedRepresentations&#x60; a JSON array containing each version of a field&#39;s value, with the highest numbered item representing the most recent version. | 
 **properties** | **[]string** | A list of issue property keys for issue properties to include in the results. This parameter accepts a comma-separated list. Multiple properties can also be provided using an ampersand separated list. For example, &#x60;properties&#x3D;prop1,prop2&amp;properties&#x3D;prop3&#x60;. A maximum of 5 issue property keys can be specified. | 
 **fieldsByKeys** | **bool** | Reference fields by their key (rather than ID). | [default to false]
 **failFast** | **bool** | Whether to fail the request quickly in case of an error while loading fields for an issue. For &#x60;failFast&#x3D;true&#x60;, if one field fails, the entire operation fails. For &#x60;failFast&#x3D;false&#x60;, the operation will continue even if a field fails. It will return a valid response, but without values for the failed field(s). | [default to false]

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchForIssuesUsingJqlPost

> SearchResults SearchForIssuesUsingJqlPost(ctx).SearchRequestBean(searchRequestBean).Execute()

Currently being removed. Search for issues using JQL (POST)



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
	searchRequestBean := *openapiclient.NewSearchRequestBean() // SearchRequestBean | A JSON object containing the search request.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSearchAPI.SearchForIssuesUsingJqlPost(context.Background()).SearchRequestBean(searchRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSearchAPI.SearchForIssuesUsingJqlPost``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchForIssuesUsingJqlPost`: SearchResults
	fmt.Fprintf(os.Stdout, "Response from `IssueSearchAPI.SearchForIssuesUsingJqlPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchForIssuesUsingJqlPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchRequestBean** | [**SearchRequestBean**](SearchRequestBean.md) | A JSON object containing the search request. | 

### Return type

[**SearchResults**](SearchResults.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

