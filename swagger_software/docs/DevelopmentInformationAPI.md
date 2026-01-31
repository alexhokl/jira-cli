# \DevelopmentInformationAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteByProperties**](DevelopmentInformationAPI.md#DeleteByProperties) | **Delete** /rest/devinfo/0.10/bulkByProperties | Delete development information by properties
[**DeleteEntity**](DevelopmentInformationAPI.md#DeleteEntity) | **Delete** /rest/devinfo/0.10/repository/{repositoryId}/{entityType}/{entityId} | Delete development information entity
[**DeleteRepository**](DevelopmentInformationAPI.md#DeleteRepository) | **Delete** /rest/devinfo/0.10/repository/{repositoryId} | Delete repository
[**ExistsByProperties**](DevelopmentInformationAPI.md#ExistsByProperties) | **Get** /rest/devinfo/0.10/existsByProperties | Check if data exists for the supplied properties
[**GetRepository**](DevelopmentInformationAPI.md#GetRepository) | **Get** /rest/devinfo/0.10/repository/{repositoryId} | Get repository
[**StoreDevelopmentInformation**](DevelopmentInformationAPI.md#StoreDevelopmentInformation) | **Post** /rest/devinfo/0.10/bulk | Store development information



## DeleteByProperties

> DeleteByProperties(ctx).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Delete development information by properties



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.
	updateSequenceId := int64(789) // int64 | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevelopmentInformationAPI.DeleteByProperties(context.Background()).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.DeleteByProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteByPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 
 **updateSequenceId** | **int64** | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  | 

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


## DeleteEntity

> DeleteEntity(ctx, repositoryId, entityType, entityId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Delete development information entity



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
	repositoryId := "repositoryId_example" // string | 
	entityType := "entityType_example" // string | 
	entityId := "entityId_example" // string | 
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.
	updateSequenceId := int64(789) // int64 | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevelopmentInformationAPI.DeleteEntity(context.Background(), repositoryId, entityType, entityId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.DeleteEntity``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **string** |  | 
**entityType** | **string** |  | 
**entityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEntityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 
 **updateSequenceId** | **int64** | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  | 

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


## DeleteRepository

> DeleteRepository(ctx, repositoryId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Delete repository



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
	repositoryId := "repositoryId_example" // string | The ID of repository to delete
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.
	updateSequenceId := int64(789) // int64 | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevelopmentInformationAPI.DeleteRepository(context.Background(), repositoryId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.DeleteRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **string** | The ID of repository to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 
 **updateSequenceId** | **int64** | An optional property to use to control deletion. Only stored data with an updateSequenceId less than or equal to that provided will be deleted. This can be used to help ensure submit/delete requests are applied correctly if they are issued close together.  | 

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


## ExistsByProperties

> ExistsForPropertiesResponse ExistsByProperties(ctx).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Check if data exists for the supplied properties



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.
	updateSequenceId := int64(789) // int64 | An optional property. Filters out entities and repositories which have updateSequenceId greater than specified.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentInformationAPI.ExistsByProperties(context.Background()).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.ExistsByProperties``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ExistsByProperties`: ExistsForPropertiesResponse
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentInformationAPI.ExistsByProperties`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiExistsByPropertiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 
 **updateSequenceId** | **int64** | An optional property. Filters out entities and repositories which have updateSequenceId greater than specified.  | 

### Return type

[**ExistsForPropertiesResponse**](ExistsForPropertiesResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRepository

> Repository GetRepository(ctx, repositoryId).Authorization(authorization).Execute()

Get repository



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
	repositoryId := "repositoryId_example" // string | The ID of repository to fetch
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentInformationAPI.GetRepository(context.Background(), repositoryId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.GetRepository``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRepository`: Repository
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentInformationAPI.GetRepository`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**repositoryId** | **string** | The ID of repository to fetch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRepositoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 

### Return type

[**Repository**](Repository.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StoreDevelopmentInformation

> StoreDevinfoResult StoreDevelopmentInformation(ctx).Authorization(authorization).DevInformation(devInformation).Execute()

Store development information



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
	authorization := "authorization_example" // string | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations.
	devInformation := *openapiclient.NewDevInformation([]openapiclient.Repository{*openapiclient.NewRepository("atlassian-connect-jira-example", "https://bitbucket.org/atlassianlabs/atlassian-connect-jira-example", "c6c7c750-cee2-48e2-b920-d7706dfd11f9", int64(1523494301248))}) // DevInformation | Request object, which contains development information

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevelopmentInformationAPI.StoreDevelopmentInformation(context.Background()).Authorization(authorization).DevInformation(devInformation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevelopmentInformationAPI.StoreDevelopmentInformation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StoreDevelopmentInformation`: StoreDevinfoResult
	fmt.Fprintf(os.Stdout, "Response from `DevelopmentInformationAPI.StoreDevelopmentInformation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStoreDevelopmentInformationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with either a Connect JWT token or OAuth token for an on-premise integration that corresponds to an app installed in Jira. If the JWT token corresponds to a Connect app that does not define the jiraDevelopmentTool module it will be rejected with a 403. See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details about Connect JWT tokens. See https://developer.atlassian.com/cloud/jira/software/integrate-jsw-cloud-with-onpremises-tools/ for details about on-premise integrations. | 
 **devInformation** | [**DevInformation**](DevInformation.md) | Request object, which contains development information | 

### Return type

[**StoreDevinfoResult**](StoreDevinfoResult.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

