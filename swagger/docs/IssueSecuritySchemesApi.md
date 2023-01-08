# \IssueSecuritySchemesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddSecurityLevel**](IssueSecuritySchemesAPI.md#AddSecurityLevel) | **Put** /rest/api/3/issuesecurityschemes/{schemeId}/level | Add issue security levels
[**AddSecurityLevelMembers**](IssueSecuritySchemesAPI.md#AddSecurityLevelMembers) | **Put** /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId}/member | Add issue security level members
[**AssociateSchemesToProjects**](IssueSecuritySchemesAPI.md#AssociateSchemesToProjects) | **Put** /rest/api/3/issuesecurityschemes/project | Associate security scheme to project
[**CreateIssueSecurityScheme**](IssueSecuritySchemesAPI.md#CreateIssueSecurityScheme) | **Post** /rest/api/3/issuesecurityschemes | Create issue security scheme
[**DeleteSecurityScheme**](IssueSecuritySchemesAPI.md#DeleteSecurityScheme) | **Delete** /rest/api/3/issuesecurityschemes/{schemeId} | Delete issue security scheme
[**GetIssueSecurityScheme**](IssueSecuritySchemesAPI.md#GetIssueSecurityScheme) | **Get** /rest/api/3/issuesecurityschemes/{id} | Get issue security scheme
[**GetIssueSecuritySchemes**](IssueSecuritySchemesAPI.md#GetIssueSecuritySchemes) | **Get** /rest/api/3/issuesecurityschemes | Get issue security schemes
[**GetSecurityLevelMembers**](IssueSecuritySchemesAPI.md#GetSecurityLevelMembers) | **Get** /rest/api/3/issuesecurityschemes/level/member | Get issue security level members
[**GetSecurityLevels**](IssueSecuritySchemesAPI.md#GetSecurityLevels) | **Get** /rest/api/3/issuesecurityschemes/level | Get issue security levels
[**RemoveLevel**](IssueSecuritySchemesAPI.md#RemoveLevel) | **Delete** /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId} | Remove issue security level
[**RemoveMemberFromSecurityLevel**](IssueSecuritySchemesAPI.md#RemoveMemberFromSecurityLevel) | **Delete** /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId}/member/{memberId} | Remove member from issue security level
[**SearchProjectsUsingSecuritySchemes**](IssueSecuritySchemesAPI.md#SearchProjectsUsingSecuritySchemes) | **Get** /rest/api/3/issuesecurityschemes/project | Get projects using issue security schemes
[**SearchSecuritySchemes**](IssueSecuritySchemesAPI.md#SearchSecuritySchemes) | **Get** /rest/api/3/issuesecurityschemes/search | Search issue security schemes
[**SetDefaultLevels**](IssueSecuritySchemesAPI.md#SetDefaultLevels) | **Put** /rest/api/3/issuesecurityschemes/level/default | Set default issue security levels
[**UpdateIssueSecurityScheme**](IssueSecuritySchemesAPI.md#UpdateIssueSecurityScheme) | **Put** /rest/api/3/issuesecurityschemes/{id} | Update issue security scheme
[**UpdateSecurityLevel**](IssueSecuritySchemesAPI.md#UpdateSecurityLevel) | **Put** /rest/api/3/issuesecurityschemes/{schemeId}/level/{levelId} | Update issue security level



## AddSecurityLevel

> interface{} AddSecurityLevel(ctx, schemeId).AddSecuritySchemeLevelsRequestBean(addSecuritySchemeLevelsRequestBean).Execute()

Add issue security levels



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme.
	addSecuritySchemeLevelsRequestBean := *openapiclient.NewAddSecuritySchemeLevelsRequestBean() // AddSecuritySchemeLevelsRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.AddSecurityLevel(context.Background(), schemeId).AddSecuritySchemeLevelsRequestBean(addSecuritySchemeLevelsRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.AddSecurityLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSecurityLevel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.AddSecurityLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **addSecuritySchemeLevelsRequestBean** | [**AddSecuritySchemeLevelsRequestBean**](AddSecuritySchemeLevelsRequestBean.md) |  | 

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


## AddSecurityLevelMembers

> interface{} AddSecurityLevelMembers(ctx, schemeId, levelId).SecuritySchemeMembersRequest(securitySchemeMembersRequest).Execute()

Add issue security level members



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme.
	levelId := "levelId_example" // string | The ID of the issue security level.
	securitySchemeMembersRequest := *openapiclient.NewSecuritySchemeMembersRequest() // SecuritySchemeMembersRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.AddSecurityLevelMembers(context.Background(), schemeId, levelId).SecuritySchemeMembersRequest(securitySchemeMembersRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.AddSecurityLevelMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddSecurityLevelMembers`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.AddSecurityLevelMembers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme. | 
**levelId** | **string** | The ID of the issue security level. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddSecurityLevelMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **securitySchemeMembersRequest** | [**SecuritySchemeMembersRequest**](SecuritySchemeMembersRequest.md) |  | 

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


## AssociateSchemesToProjects

> AssociateSchemesToProjects(ctx).AssociateSecuritySchemeWithProjectDetails(associateSecuritySchemeWithProjectDetails).Execute()

Associate security scheme to project



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
	associateSecuritySchemeWithProjectDetails := *openapiclient.NewAssociateSecuritySchemeWithProjectDetails("ProjectId_example", "SchemeId_example") // AssociateSecuritySchemeWithProjectDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueSecuritySchemesAPI.AssociateSchemesToProjects(context.Background()).AssociateSecuritySchemeWithProjectDetails(associateSecuritySchemeWithProjectDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.AssociateSchemesToProjects``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssociateSchemesToProjectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **associateSecuritySchemeWithProjectDetails** | [**AssociateSecuritySchemeWithProjectDetails**](AssociateSecuritySchemeWithProjectDetails.md) |  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateIssueSecurityScheme

> SecuritySchemeId CreateIssueSecurityScheme(ctx).CreateIssueSecuritySchemeDetails(createIssueSecuritySchemeDetails).Execute()

Create issue security scheme



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
	createIssueSecuritySchemeDetails := *openapiclient.NewCreateIssueSecuritySchemeDetails("Name_example") // CreateIssueSecuritySchemeDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.CreateIssueSecurityScheme(context.Background()).CreateIssueSecuritySchemeDetails(createIssueSecuritySchemeDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.CreateIssueSecurityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIssueSecurityScheme`: SecuritySchemeId
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.CreateIssueSecurityScheme`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateIssueSecuritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createIssueSecuritySchemeDetails** | [**CreateIssueSecuritySchemeDetails**](CreateIssueSecuritySchemeDetails.md) |  | 

### Return type

[**SecuritySchemeId**](SecuritySchemeId.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteSecurityScheme

> interface{} DeleteSecurityScheme(ctx, schemeId).Execute()

Delete issue security scheme



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.DeleteSecurityScheme(context.Background(), schemeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.DeleteSecurityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteSecurityScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.DeleteSecurityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSecuritySchemeRequest struct via the builder pattern


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


## GetIssueSecurityScheme

> SecurityScheme GetIssueSecurityScheme(ctx, id).Execute()

Get issue security scheme



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
	id := int64(789) // int64 | The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) operation to get a list of issue security scheme IDs.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.GetIssueSecurityScheme(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.GetIssueSecurityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueSecurityScheme`: SecurityScheme
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.GetIssueSecurityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **int64** | The ID of the issue security scheme. Use the [Get issue security schemes](#api-rest-api-3-issuesecurityschemes-get) operation to get a list of issue security scheme IDs. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueSecuritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**SecurityScheme**](SecurityScheme.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIssueSecuritySchemes

> SecuritySchemes GetIssueSecuritySchemes(ctx).Execute()

Get issue security schemes



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
	resp, r, err := apiClient.IssueSecuritySchemesAPI.GetIssueSecuritySchemes(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.GetIssueSecuritySchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIssueSecuritySchemes`: SecuritySchemes
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.GetIssueSecuritySchemes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetIssueSecuritySchemesRequest struct via the builder pattern


### Return type

[**SecuritySchemes**](SecuritySchemes.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityLevelMembers

> PageBeanSecurityLevelMember GetSecurityLevelMembers(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).SchemeId(schemeId).LevelId(levelId).Expand(expand).Execute()

Get issue security level members



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	id := []string{"Inner_example"} // []string | The list of issue security level member IDs. To include multiple issue security level members separate IDs with an ampersand: `id=10000&id=10001`. (optional)
	schemeId := []string{"Inner_example"} // []string | The list of issue security scheme IDs. To include multiple issue security schemes separate IDs with an ampersand: `schemeId=10000&schemeId=10001`. (optional)
	levelId := []string{"Inner_example"} // []string | The list of issue security level IDs. To include multiple issue security levels separate IDs with an ampersand: `levelId=10000&levelId=10001`. (optional)
	expand := "expand_example" // string | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  `all` Returns all expandable information  *  `field` Returns information about the custom field granted the permission  *  `group` Returns information about the group that is granted the permission  *  `projectRole` Returns information about the project role granted the permission  *  `user` Returns information about the user who is granted the permission (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.GetSecurityLevelMembers(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).SchemeId(schemeId).LevelId(levelId).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.GetSecurityLevelMembers``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLevelMembers`: PageBeanSecurityLevelMember
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.GetSecurityLevelMembers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLevelMembersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of issue security level member IDs. To include multiple issue security level members separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **schemeId** | **[]string** | The list of issue security scheme IDs. To include multiple issue security schemes separate IDs with an ampersand: &#x60;schemeId&#x3D;10000&amp;schemeId&#x3D;10001&#x60;. | 
 **levelId** | **[]string** | The list of issue security level IDs. To include multiple issue security levels separate IDs with an ampersand: &#x60;levelId&#x3D;10000&amp;levelId&#x3D;10001&#x60;. | 
 **expand** | **string** | Use expand to include additional information in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;all&#x60; Returns all expandable information  *  &#x60;field&#x60; Returns information about the custom field granted the permission  *  &#x60;group&#x60; Returns information about the group that is granted the permission  *  &#x60;projectRole&#x60; Returns information about the project role granted the permission  *  &#x60;user&#x60; Returns information about the user who is granted the permission | 

### Return type

[**PageBeanSecurityLevelMember**](PageBeanSecurityLevelMember.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSecurityLevels

> PageBeanSecurityLevel GetSecurityLevels(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).SchemeId(schemeId).OnlyDefault(onlyDefault).Execute()

Get issue security levels



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	id := []string{"Inner_example"} // []string | The list of issue security scheme level IDs. To include multiple issue security levels, separate IDs with an ampersand: `id=10000&id=10001`. (optional)
	schemeId := []string{"Inner_example"} // []string | The list of issue security scheme IDs. To include multiple issue security schemes, separate IDs with an ampersand: `schemeId=10000&schemeId=10001`. (optional)
	onlyDefault := true // bool | When set to true, returns multiple default levels for each security scheme containing a default. If you provide scheme and level IDs not associated with the default, returns an empty page. The default value is false. (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.GetSecurityLevels(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).SchemeId(schemeId).OnlyDefault(onlyDefault).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.GetSecurityLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSecurityLevels`: PageBeanSecurityLevel
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.GetSecurityLevels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSecurityLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of issue security scheme level IDs. To include multiple issue security levels, separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **schemeId** | **[]string** | The list of issue security scheme IDs. To include multiple issue security schemes, separate IDs with an ampersand: &#x60;schemeId&#x3D;10000&amp;schemeId&#x3D;10001&#x60;. | 
 **onlyDefault** | **bool** | When set to true, returns multiple default levels for each security scheme containing a default. If you provide scheme and level IDs not associated with the default, returns an empty page. The default value is false. | [default to false]

### Return type

[**PageBeanSecurityLevel**](PageBeanSecurityLevel.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveLevel

> RemoveLevel(ctx, schemeId, levelId).ReplaceWith(replaceWith).Execute()

Remove issue security level



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme.
	levelId := "levelId_example" // string | The ID of the issue security level to remove.
	replaceWith := "replaceWith_example" // string | The ID of the issue security level that will replace the currently selected level. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.IssueSecuritySchemesAPI.RemoveLevel(context.Background(), schemeId, levelId).ReplaceWith(replaceWith).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.RemoveLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme. | 
**levelId** | **string** | The ID of the issue security level to remove. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **replaceWith** | **string** | The ID of the issue security level that will replace the currently selected level. | 

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


## RemoveMemberFromSecurityLevel

> interface{} RemoveMemberFromSecurityLevel(ctx, schemeId, levelId, memberId).Execute()

Remove member from issue security level



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme.
	levelId := "levelId_example" // string | The ID of the issue security level.
	memberId := "memberId_example" // string | The ID of the issue security level member to be removed.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.RemoveMemberFromSecurityLevel(context.Background(), schemeId, levelId, memberId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.RemoveMemberFromSecurityLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveMemberFromSecurityLevel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.RemoveMemberFromSecurityLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme. | 
**levelId** | **string** | The ID of the issue security level. | 
**memberId** | **string** | The ID of the issue security level member to be removed. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveMemberFromSecurityLevelRequest struct via the builder pattern


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


## SearchProjectsUsingSecuritySchemes

> PageBeanIssueSecuritySchemeToProjectMapping SearchProjectsUsingSecuritySchemes(ctx).StartAt(startAt).MaxResults(maxResults).IssueSecuritySchemeId(issueSecuritySchemeId).ProjectId(projectId).Execute()

Get projects using issue security schemes



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	issueSecuritySchemeId := []string{"Inner_example"} // []string | The list of security scheme IDs to be filtered out. (optional)
	projectId := []string{"Inner_example"} // []string | The list of project IDs to be filtered out. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.SearchProjectsUsingSecuritySchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).IssueSecuritySchemeId(issueSecuritySchemeId).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.SearchProjectsUsingSecuritySchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchProjectsUsingSecuritySchemes`: PageBeanIssueSecuritySchemeToProjectMapping
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.SearchProjectsUsingSecuritySchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchProjectsUsingSecuritySchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **issueSecuritySchemeId** | **[]string** | The list of security scheme IDs to be filtered out. | 
 **projectId** | **[]string** | The list of project IDs to be filtered out. | 

### Return type

[**PageBeanIssueSecuritySchemeToProjectMapping**](PageBeanIssueSecuritySchemeToProjectMapping.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SearchSecuritySchemes

> PageBeanSecuritySchemeWithProjects SearchSecuritySchemes(ctx).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).Execute()

Search issue security schemes



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
	startAt := "startAt_example" // string | The index of the first item to return in a page of results (page offset). (optional) (default to "0")
	maxResults := "maxResults_example" // string | The maximum number of items to return per page. (optional) (default to "50")
	id := []string{"Inner_example"} // []string | The list of issue security scheme IDs. To include multiple issue security scheme IDs, separate IDs with an ampersand: `id=10000&id=10001`. (optional)
	projectId := []string{"Inner_example"} // []string | The list of project IDs. To include multiple project IDs, separate IDs with an ampersand: `projectId=10000&projectId=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.SearchSecuritySchemes(context.Background()).StartAt(startAt).MaxResults(maxResults).Id(id).ProjectId(projectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.SearchSecuritySchemes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SearchSecuritySchemes`: PageBeanSecuritySchemeWithProjects
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.SearchSecuritySchemes`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchSecuritySchemesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **startAt** | **string** | The index of the first item to return in a page of results (page offset). | [default to &quot;0&quot;]
 **maxResults** | **string** | The maximum number of items to return per page. | [default to &quot;50&quot;]
 **id** | **[]string** | The list of issue security scheme IDs. To include multiple issue security scheme IDs, separate IDs with an ampersand: &#x60;id&#x3D;10000&amp;id&#x3D;10001&#x60;. | 
 **projectId** | **[]string** | The list of project IDs. To include multiple project IDs, separate IDs with an ampersand: &#x60;projectId&#x3D;10000&amp;projectId&#x3D;10001&#x60;. | 

### Return type

[**PageBeanSecuritySchemeWithProjects**](PageBeanSecuritySchemeWithProjects.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDefaultLevels

> interface{} SetDefaultLevels(ctx).SetDefaultLevelsRequest(setDefaultLevelsRequest).Execute()

Set default issue security levels



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
	setDefaultLevelsRequest := *openapiclient.NewSetDefaultLevelsRequest([]openapiclient.DefaultLevelValue{*openapiclient.NewDefaultLevelValue("DefaultLevelId_example", "IssueSecuritySchemeId_example")}) // SetDefaultLevelsRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.SetDefaultLevels(context.Background()).SetDefaultLevelsRequest(setDefaultLevelsRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.SetDefaultLevels``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDefaultLevels`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.SetDefaultLevels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetDefaultLevelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **setDefaultLevelsRequest** | [**SetDefaultLevelsRequest**](SetDefaultLevelsRequest.md) |  | 

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


## UpdateIssueSecurityScheme

> interface{} UpdateIssueSecurityScheme(ctx, id).UpdateIssueSecuritySchemeRequestBean(updateIssueSecuritySchemeRequestBean).Execute()

Update issue security scheme



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
	id := "id_example" // string | The ID of the issue security scheme.
	updateIssueSecuritySchemeRequestBean := *openapiclient.NewUpdateIssueSecuritySchemeRequestBean() // UpdateIssueSecuritySchemeRequestBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.UpdateIssueSecurityScheme(context.Background(), id).UpdateIssueSecuritySchemeRequestBean(updateIssueSecuritySchemeRequestBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.UpdateIssueSecurityScheme``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIssueSecurityScheme`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.UpdateIssueSecurityScheme`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the issue security scheme. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIssueSecuritySchemeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateIssueSecuritySchemeRequestBean** | [**UpdateIssueSecuritySchemeRequestBean**](UpdateIssueSecuritySchemeRequestBean.md) |  | 

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


## UpdateSecurityLevel

> interface{} UpdateSecurityLevel(ctx, schemeId, levelId).UpdateIssueSecurityLevelDetails(updateIssueSecurityLevelDetails).Execute()

Update issue security level



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
	schemeId := "schemeId_example" // string | The ID of the issue security scheme level belongs to.
	levelId := "levelId_example" // string | The ID of the issue security level to update.
	updateIssueSecurityLevelDetails := *openapiclient.NewUpdateIssueSecurityLevelDetails() // UpdateIssueSecurityLevelDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.IssueSecuritySchemesAPI.UpdateSecurityLevel(context.Background(), schemeId, levelId).UpdateIssueSecurityLevelDetails(updateIssueSecurityLevelDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IssueSecuritySchemesAPI.UpdateSecurityLevel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateSecurityLevel`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `IssueSecuritySchemesAPI.UpdateSecurityLevel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**schemeId** | **string** | The ID of the issue security scheme level belongs to. | 
**levelId** | **string** | The ID of the issue security level to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSecurityLevelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIssueSecurityLevelDetails** | [**UpdateIssueSecurityLevelDetails**](UpdateIssueSecurityLevelDetails.md) |  | 

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

