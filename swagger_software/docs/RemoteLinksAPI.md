# \RemoteLinksAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteRemoteLinkById**](RemoteLinksAPI.md#DeleteRemoteLinkById) | **Delete** /rest/remotelinks/1.0/remotelink/{remoteLinkId} | Delete a Remote Link by ID
[**DeleteRemoteLinksByProperty**](RemoteLinksAPI.md#DeleteRemoteLinksByProperty) | **Delete** /rest/remotelinks/1.0/bulkByProperties | Delete Remote Links by Property
[**GetRemoteLinkById**](RemoteLinksAPI.md#GetRemoteLinkById) | **Get** /rest/remotelinks/1.0/remotelink/{remoteLinkId} | Get a Remote Link by ID
[**SubmitRemoteLinks**](RemoteLinksAPI.md#SubmitRemoteLinks) | **Post** /rest/remotelinks/1.0/bulk | Submit Remote Link data



## DeleteRemoteLinkById

> DeleteRemoteLinkById(ctx, remoteLinkId).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()

Delete a Remote Link by ID



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraRemoteLinkInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	remoteLinkId := "remoteLinkId_example" // string | The ID of the Remote Link to fetch. 
	updateSequenceNumber := int64(789) // int64 | This parameter usage is no longer supported.  An optional `_updateSequenceNumber` to use to control deletion.  Only stored data with an `updateSequenceNumber` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteLinksAPI.DeleteRemoteLinkById(context.Background(), remoteLinkId).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteLinksAPI.DeleteRemoteLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteLinkId** | **string** | The ID of the Remote Link to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraRemoteLinkInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 

 **updateSequenceNumber** | **int64** | This parameter usage is no longer supported.  An optional &#x60;_updateSequenceNumber&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceNumber&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteRemoteLinksByProperty

> DeleteRemoteLinksByProperty(ctx).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Params(params).Execute()

Delete Remote Links by Property



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraRemoteLinkInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	updateSequenceNumber := int64(789) // int64 | This parameter usage is no longer supported.  An optional `_updateSequenceNumber` to use to control deletion.  Only stored data with an `updateSequenceNumber` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  (optional)
	params := map[string]interface{}{ ... } // map[string]interface{} | Free-form query parameters to specify which properties to delete by. Properties refer to the arbitrary information the provider tagged Remote Links with previously.  For example, if the provider previously tagged a remote link with accountId:   \"properties\": {     \"accountId\": \"account-123\"   }  And now they want to delete Remote Links in bulk by that specific accountId as follows: e.g. DELETE /bulkByProperties?accountId=account-123  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.RemoteLinksAPI.DeleteRemoteLinksByProperty(context.Background()).Authorization(authorization).UpdateSequenceNumber(updateSequenceNumber).Params(params).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteLinksAPI.DeleteRemoteLinksByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRemoteLinksByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraRemoteLinkInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 
 **updateSequenceNumber** | **int64** | This parameter usage is no longer supported.  An optional &#x60;_updateSequenceNumber&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceNumber&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  | 
 **params** | [**map[string]interface{}**](map[string]interface{}.md) | Free-form query parameters to specify which properties to delete by. Properties refer to the arbitrary information the provider tagged Remote Links with previously.  For example, if the provider previously tagged a remote link with accountId:   \&quot;properties\&quot;: {     \&quot;accountId\&quot;: \&quot;account-123\&quot;   }  And now they want to delete Remote Links in bulk by that specific accountId as follows: e.g. DELETE /bulkByProperties?accountId&#x3D;account-123  | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRemoteLinkById

> RemoteLinkData GetRemoteLinkById(ctx, remoteLinkId).Authorization(authorization).Execute()

Get a Remote Link by ID



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraRemoteLinkInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	remoteLinkId := "remoteLinkId_example" // string | The ID of the Remote Link to fetch. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteLinksAPI.GetRemoteLinkById(context.Background(), remoteLinkId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteLinksAPI.GetRemoteLinkById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRemoteLinkById`: RemoteLinkData
	fmt.Fprintf(os.Stdout, "Response from `RemoteLinksAPI.GetRemoteLinkById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**remoteLinkId** | **string** | The ID of the Remote Link to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRemoteLinkByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraRemoteLinkInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 


### Return type

[**RemoteLinkData**](RemoteLinkData.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitRemoteLinks

> SubmitRemoteLinks202Response SubmitRemoteLinks(ctx).Authorization(authorization).SubmitRemoteLinksRequest(submitRemoteLinksRequest).Execute()

Submit Remote Link data



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
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define `jiraRemoteLinkInfoProvider` module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. 
	submitRemoteLinksRequest := *openapiclient.NewSubmitRemoteLinksRequest([]openapiclient.RemoteLinkData{*openapiclient.NewRemoteLinkData("111-222-333", int64(1523494301448), "Remote Link #42", "Url_example", "security", time.Now())}) // SubmitRemoteLinksRequest | Remote Links data to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RemoteLinksAPI.SubmitRemoteLinks(context.Background()).Authorization(authorization).SubmitRemoteLinksRequest(submitRemoteLinksRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RemoteLinksAPI.SubmitRemoteLinks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitRemoteLinks`: SubmitRemoteLinks202Response
	fmt.Fprintf(os.Stdout, "Response from `RemoteLinksAPI.SubmitRemoteLinks`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitRemoteLinksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to an app installed in Jira.  If the Connect JWT token corresponds to an app that does not define &#x60;jiraRemoteLinkInfoProvider&#x60; module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.  | 
 **submitRemoteLinksRequest** | [**SubmitRemoteLinksRequest**](SubmitRemoteLinksRequest.md) | Remote Links data to submit.  | 

### Return type

[**SubmitRemoteLinks202Response**](SubmitRemoteLinks202Response.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

