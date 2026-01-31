# \SecurityInformationAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLinkedWorkspaces**](SecurityInformationAPI.md#DeleteLinkedWorkspaces) | **Delete** /rest/security/1.0/linkedWorkspaces/bulk | Delete linked Security Workspaces
[**DeleteVulnerabilitiesByProperty**](SecurityInformationAPI.md#DeleteVulnerabilitiesByProperty) | **Delete** /rest/security/1.0/bulkByProperties | Delete Vulnerabilities by Property
[**DeleteVulnerabilityById**](SecurityInformationAPI.md#DeleteVulnerabilityById) | **Delete** /rest/security/1.0/vulnerability/{vulnerabilityId} | Delete a Vulnerability by ID
[**GetLinkedWorkspaceById**](SecurityInformationAPI.md#GetLinkedWorkspaceById) | **Get** /rest/security/1.0/linkedWorkspaces/{workspaceId} | Get a linked Security Workspace by ID
[**GetLinkedWorkspaces**](SecurityInformationAPI.md#GetLinkedWorkspaces) | **Get** /rest/security/1.0/linkedWorkspaces | Get linked Security Workspaces
[**GetVulnerabilityById**](SecurityInformationAPI.md#GetVulnerabilityById) | **Get** /rest/security/1.0/vulnerability/{vulnerabilityId} | Get a Vulnerability by ID
[**SubmitVulnerabilities**](SecurityInformationAPI.md#SubmitVulnerabilities) | **Post** /rest/security/1.0/bulk | Submit Vulnerability data
[**SubmitWorkspaces**](SecurityInformationAPI.md#SubmitWorkspaces) | **Post** /rest/security/1.0/linkedWorkspaces/bulk | Submit Security Workspaces to link



## DeleteLinkedWorkspaces

> DeleteLinkedWorkspaces(ctx).Authorization(authorization).Execute()

Delete linked Security Workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read [understanding jwt](https://developer.atlassian.com/blog/2015/01/understanding-jwt/) for more details. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityInformationAPI.DeleteLinkedWorkspaces(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.DeleteLinkedWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLinkedWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read [understanding jwt](https://developer.atlassian.com/blog/2015/01/understanding-jwt/) for more details.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVulnerabilitiesByProperty

> DeleteVulnerabilitiesByProperty(ctx).Authorization(authorization).Execute()

Delete Vulnerabilities by Property



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityInformationAPI.DeleteVulnerabilitiesByProperty(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.DeleteVulnerabilitiesByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVulnerabilitiesByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteVulnerabilityById

> DeleteVulnerabilityById(ctx, vulnerabilityId).Authorization(authorization).Execute()

Delete a Vulnerability by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 
	vulnerabilityId := "vulnerabilityId_example" // string | The ID of the Vulnerability to delete. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityInformationAPI.DeleteVulnerabilityById(context.Background(), vulnerabilityId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.DeleteVulnerabilityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vulnerabilityId** | **string** | The ID of the Vulnerability to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVulnerabilityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedWorkspaceById

> SecurityWorkspaceResponse GetLinkedWorkspaceById(ctx, workspaceId).Authorization(authorization).Execute()

Get a linked Security Workspace by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 
	workspaceId := "workspaceId_example" // string | The ID of the workspace to fetch. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityInformationAPI.GetLinkedWorkspaceById(context.Background(), workspaceId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.GetLinkedWorkspaceById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedWorkspaceById`: SecurityWorkspaceResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityInformationAPI.GetLinkedWorkspaceById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspaceId** | **string** | The ID of the workspace to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedWorkspaceByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 


### Return type

[**SecurityWorkspaceResponse**](SecurityWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLinkedWorkspaces

> SecurityWorkspaceIds GetLinkedWorkspaces(ctx).Authorization(authorization).Execute()

Get linked Security Workspaces



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityInformationAPI.GetLinkedWorkspaces(context.Background()).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.GetLinkedWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLinkedWorkspaces`: SecurityWorkspaceIds
	fmt.Fprintf(os.Stdout, "Response from `SecurityInformationAPI.GetLinkedWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetLinkedWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 

### Return type

[**SecurityWorkspaceIds**](SecurityWorkspaceIds.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetVulnerabilityById

> VulnerabilityDetails GetVulnerabilityById(ctx, vulnerabilityId).Authorization(authorization).Execute()

Get a Vulnerability by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 
	vulnerabilityId := "vulnerabilityId_example" // string | The ID of the Vulnerability to fetch. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityInformationAPI.GetVulnerabilityById(context.Background(), vulnerabilityId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.GetVulnerabilityById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetVulnerabilityById`: VulnerabilityDetails
	fmt.Fprintf(os.Stdout, "Response from `SecurityInformationAPI.GetVulnerabilityById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**vulnerabilityId** | **string** | The ID of the Vulnerability to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVulnerabilityByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 


### Return type

[**VulnerabilityDetails**](VulnerabilityDetails.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitVulnerabilities

> SubmitVulnerabilitiesResponse SubmitVulnerabilities(ctx).Authorization(authorization).SubmitVulnerabilitiesRequest(submitVulnerabilitiesRequest).Execute()

Submit Vulnerability data



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
    "time"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/). 
	submitVulnerabilitiesRequest := *openapiclient.NewSubmitVulnerabilitiesRequest([]openapiclient.VulnerabilityDetails{*openapiclient.NewVulnerabilityDetails("1.0", "111-222-333", int64(1523494301448), "111-222-333", "curl/libcurl3 - Buffer Override", "## Overview


Affected versions of this package are vulnerable to MeltLeak", "https://example.com/project/CWE-123/summary", "sca", time.Now(), time.Now(), *openapiclient.NewVulnerabilitySeverity("critical"), "open")}) // SubmitVulnerabilitiesRequest | Vulnerability data to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.SecurityInformationAPI.SubmitVulnerabilities(context.Background()).Authorization(authorization).SubmitVulnerabilitiesRequest(submitVulnerabilitiesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.SubmitVulnerabilities``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitVulnerabilities`: SubmitVulnerabilitiesResponse
	fmt.Fprintf(os.Stdout, "Response from `SecurityInformationAPI.SubmitVulnerabilities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitVulnerabilitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read more about JWT [here](https://developer.atlassian.com/blog/2015/01/understanding-jwt/).  | 
 **submitVulnerabilitiesRequest** | [**SubmitVulnerabilitiesRequest**](SubmitVulnerabilitiesRequest.md) | Vulnerability data to submit.  | 

### Return type

[**SubmitVulnerabilitiesResponse**](SubmitVulnerabilitiesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitWorkspaces

> SubmitWorkspaces(ctx).Authorization(authorization).SubmitSecurityWorkspacesRequest(submitSecurityWorkspacesRequest).Execute()

Submit Security Workspaces to link



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/alexhokl/jira-cli/swagger_software"
)

func main() {
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read [understanding jwt](https://developer.atlassian.com/blog/2015/01/understanding-jwt/) for more details. 
	submitSecurityWorkspacesRequest := *openapiclient.NewSubmitSecurityWorkspacesRequest([]string{"WorkspaceIds_example"}) // SubmitSecurityWorkspacesRequest | Security Workspace IDs to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.SecurityInformationAPI.SubmitWorkspaces(context.Background()).Authorization(authorization).SubmitSecurityWorkspacesRequest(submitSecurityWorkspacesRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `SecurityInformationAPI.SubmitWorkspaces``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define the Security Information module it will be rejected with a 403.  Read [understanding jwt](https://developer.atlassian.com/blog/2015/01/understanding-jwt/) for more details.  | 
 **submitSecurityWorkspacesRequest** | [**SubmitSecurityWorkspacesRequest**](SubmitSecurityWorkspacesRequest.md) | Security Workspace IDs to submit.  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

