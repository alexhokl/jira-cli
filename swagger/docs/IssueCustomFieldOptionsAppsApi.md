# \IssueCustomFieldOptionsAppsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIssueFieldOption**](IssueCustomFieldOptionsAppsAPI.md#CreateIssueFieldOption) | **Post** /rest/api/3/field/{fieldKey}/option | Create issue field option
[**DeleteIssueFieldOption**](IssueCustomFieldOptionsAppsAPI.md#DeleteIssueFieldOption) | **Delete** /rest/api/3/field/{fieldKey}/option/{optionId} | Delete issue field option
[**GetAllIssueFieldOptions**](IssueCustomFieldOptionsAppsAPI.md#GetAllIssueFieldOptions) | **Get** /rest/api/3/field/{fieldKey}/option | Get all issue field options
[**GetIssueFieldOption**](IssueCustomFieldOptionsAppsAPI.md#GetIssueFieldOption) | **Get** /rest/api/3/field/{fieldKey}/option/{optionId} | Get issue field option
[**GetSelectableIssueFieldOptions**](IssueCustomFieldOptionsAppsAPI.md#GetSelectableIssueFieldOptions) | **Get** /rest/api/3/field/{fieldKey}/option/suggestions/edit | Get selectable issue field options
[**GetVisibleIssueFieldOptions**](IssueCustomFieldOptionsAppsAPI.md#GetVisibleIssueFieldOptions) | **Get** /rest/api/3/field/{fieldKey}/option/suggestions/search | Get visible issue field options
[**ReplaceIssueFieldOption**](IssueCustomFieldOptionsAppsAPI.md#ReplaceIssueFieldOption) | **Delete** /rest/api/3/field/{fieldKey}/option/{optionId}/issue | Replace issue field option
[**UpdateIssueFieldOption**](IssueCustomFieldOptionsAppsAPI.md#UpdateIssueFieldOption) | **Put** /rest/api/3/field/{fieldKey}/option/{optionId} | Update issue field option



## CreateIssueFieldOption

> IssueFieldOption CreateIssueFieldOption(ctx, fieldKey).IssueFieldOptionCreateBean(issueFieldOptionCreateBean).Execute()

Create issue field option



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	issueFieldOptionCreateBean := *openapiclient.NewIssueFieldOptionCreateBean("Value_example") // IssueFieldOptionCreateBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.CreateIssueFieldOption(context.Background(), fieldKey).IssueFieldOptionCreateBean(issueFieldOptionCreateBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.CreateIssueFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueFieldOption`: IssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.CreateIssueFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **issueFieldOptionCreateBean** | [**IssueFieldOptionCreateBean**](IssueFieldOptionCreateBean.md) |  | 

### Return type

[**IssueFieldOption**](IssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteIssueFieldOption

> interface{} DeleteIssueFieldOption(ctx, fieldKey, optionId).Execute()

Delete issue field option



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	optionId := int64(789) // int64 | The ID of the option to be deleted.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.DeleteIssueFieldOption(context.Background(), fieldKey, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.DeleteIssueFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIssueFieldOption`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.DeleteIssueFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 
**optionId** | **int64** | The ID of the option to be deleted. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIssueFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**interface{}**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllIssueFieldOptions

> PageBeanIssueFieldOption GetAllIssueFieldOptions(ctx, fieldKey).StartAt(startAt).MaxResults(maxResults).Execute()

Get all issue field options



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.GetAllIssueFieldOptions(context.Background(), fieldKey).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.GetAllIssueFieldOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllIssueFieldOptions`: PageBeanIssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.GetAllIssueFieldOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllIssueFieldOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]

### Return type

[**PageBeanIssueFieldOption**](PageBeanIssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueFieldOption

> IssueFieldOption GetIssueFieldOption(ctx, fieldKey, optionId).Execute()

Get issue field option



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	optionId := int64(789) // int64 | The ID of the option to be returned.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.GetIssueFieldOption(context.Background(), fieldKey, optionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.GetIssueFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueFieldOption`: IssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.GetIssueFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 
**optionId** | **int64** | The ID of the option to be returned. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IssueFieldOption**](IssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectableIssueFieldOptions

> PageBeanIssueFieldOption GetSelectableIssueFieldOptions(ctx, fieldKey).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()

Get selectable issue field options



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	projectId := int64(789) // int64 | Filters the results to options that are only available in the specified project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.GetSelectableIssueFieldOptions(context.Background(), fieldKey).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.GetSelectableIssueFieldOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectableIssueFieldOptions`: PageBeanIssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.GetSelectableIssueFieldOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectableIssueFieldOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **projectId** | **int64** | Filters the results to options that are only available in the specified project. | 

### Return type

[**PageBeanIssueFieldOption**](PageBeanIssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVisibleIssueFieldOptions

> PageBeanIssueFieldOption GetVisibleIssueFieldOptions(ctx, fieldKey).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()

Get visible issue field options



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	projectId := int64(789) // int64 | Filters the results to options that are only available in the specified project. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.GetVisibleIssueFieldOptions(context.Background(), fieldKey).StartAt(startAt).MaxResults(maxResults).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.GetVisibleIssueFieldOptions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVisibleIssueFieldOptions`: PageBeanIssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.GetVisibleIssueFieldOptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVisibleIssueFieldOptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **projectId** | **int64** | Filters the results to options that are only available in the specified project. | 

### Return type

[**PageBeanIssueFieldOption**](PageBeanIssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplaceIssueFieldOption

> ReplaceIssueFieldOption(ctx, fieldKey, optionId).ReplaceWith(replaceWith).Jql(jql).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).Execute()

Replace issue field option



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	optionId := int64(789) // int64 | The ID of the option to be deselected.
	replaceWith := int64(789) // int64 | The ID of the option that will replace the currently selected option. (optional)
	jql := "jql_example" // string | A JQL query that specifies the issues to be updated. For example, *project=10000*. (optional)
	overrideScreenSecurity := true // bool | Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users with admin permission. (optional) (default to false)
	overrideEditableFlag := true // bool | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueCustomFieldOptionsAppsAPI.ReplaceIssueFieldOption(context.Background(), fieldKey, optionId).ReplaceWith(replaceWith).Jql(jql).OverrideScreenSecurity(overrideScreenSecurity).OverrideEditableFlag(overrideEditableFlag).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.ReplaceIssueFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 
**optionId** | **int64** | The ID of the option to be deselected. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplaceIssueFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceWith** | **int64** | The ID of the option that will replace the currently selected option. | 
 **jql** | **string** | A JQL query that specifies the issues to be updated. For example, *project&#x3D;10000*. | 
 **overrideScreenSecurity** | **bool** | Whether screen security is overridden to enable hidden fields to be edited. Available to Connect and Forge app users with admin permission. | [default to false]
 **overrideEditableFlag** | **bool** | Whether screen security is overridden to enable uneditable fields to be edited. Available to Connect and Forge app users with *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg). | [default to false]

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateIssueFieldOption

> IssueFieldOption UpdateIssueFieldOption(ctx, fieldKey, optionId).IssueFieldOption(issueFieldOption).Execute()

Update issue field option



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
	fieldKey := "fieldKey_example" // string | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the `fieldKey` value, do one of the following:   *  open the app's plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the `jiraIssueFields` module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in `key`. For example, `\"key\": \"teams-add-on__team-issue-field\"`
	optionId := int64(789) // int64 | The ID of the option to be updated.
	issueFieldOption := *openapiclient.NewIssueFieldOption(int64(123), "Value_example") // IssueFieldOption | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueCustomFieldOptionsAppsAPI.UpdateIssueFieldOption(context.Background(), fieldKey, optionId).IssueFieldOption(issueFieldOption).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueCustomFieldOptionsAppsAPI.UpdateIssueFieldOption``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueFieldOption`: IssueFieldOption
	fmt.Fprintf(os.Stdout, "Response from `IssueCustomFieldOptionsAppsAPI.UpdateIssueFieldOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fieldKey** | **string** | The field key is specified in the following format: **$(app-key)\\_\\_$(field-key)**. For example, *example-add-on\\_\\_example-issue-field*. To determine the &#x60;fieldKey&#x60; value, do one of the following:   *  open the app&#39;s plugin descriptor, then **app-key** is the key at the top and **field-key** is the key in the &#x60;jiraIssueFields&#x60; module. **app-key** can also be found in the app listing in the Atlassian Universal Plugin Manager.  *  run [Get fields](#api-rest-api-3-field-get) and in the field details the value is returned in &#x60;key&#x60;. For example, &#x60;\&quot;key\&quot;: \&quot;teams-add-on__team-issue-field\&quot;&#x60; | 
**optionId** | **int64** | The ID of the option to be updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueFieldOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **issueFieldOption** | [**IssueFieldOption**](IssueFieldOption.md) |  | 

### Return type

[**IssueFieldOption**](IssueFieldOption.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

