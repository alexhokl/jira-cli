# \FeatureFlagsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteFeatureFlagById**](FeatureFlagsAPI.md#DeleteFeatureFlagById) | **Delete** /rest/featureflags/0.1/flag/{featureFlagId} | Delete a Feature Flag by ID
[**DeleteFeatureFlagsByProperty**](FeatureFlagsAPI.md#DeleteFeatureFlagsByProperty) | **Delete** /rest/featureflags/0.1/bulkByProperties | Delete Feature Flags by Property
[**GetFeatureFlagById**](FeatureFlagsAPI.md#GetFeatureFlagById) | **Get** /rest/featureflags/0.1/flag/{featureFlagId} | Get a Feature Flag by ID
[**SubmitFeatureFlags**](FeatureFlagsAPI.md#SubmitFeatureFlags) | **Post** /rest/featureflags/0.1/bulk | Submit Feature Flag data



## DeleteFeatureFlagById

> DeleteFeatureFlagById(ctx, featureFlagId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Delete a Feature Flag by ID



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
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	featureFlagId := "featureFlagId_example" // string | The ID of the Feature Flag to delete. 
	updateSequenceId := int64(789) // int64 | This parameter usage is no longer supported.  An optional `_updateSequenceId` to use to control deletion.  Only stored data with an `updateSequenceId` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeatureFlagsAPI.DeleteFeatureFlagById(context.Background(), featureFlagId).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.DeleteFeatureFlagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureFlagId** | **string** | The ID of the Feature Flag to delete.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 

 **updateSequenceId** | **int64** | This parameter usage is no longer supported.  An optional &#x60;_updateSequenceId&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceId&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  | 

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


## DeleteFeatureFlagsByProperty

> DeleteFeatureFlagsByProperty(ctx).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()

Delete Feature Flags by Property



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
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	updateSequenceId := int64(789) // int64 | This parameter usage is no longer supported.  An optional `_updateSequenceId` to use to control deletion.  Only stored data with an `updateSequenceId` less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.FeatureFlagsAPI.DeleteFeatureFlagsByProperty(context.Background()).Authorization(authorization).UpdateSequenceId(updateSequenceId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.DeleteFeatureFlagsByProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFeatureFlagsByPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 
 **updateSequenceId** | **int64** | This parameter usage is no longer supported.  An optional &#x60;_updateSequenceId&#x60; to use to control deletion.  Only stored data with an &#x60;updateSequenceId&#x60; less than or equal to that provided will be deleted. This can be used help ensure submit/delete requests are applied correctly if issued close together.  If not provided, all stored data that matches the request will be deleted.  | 

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


## GetFeatureFlagById

> FeatureFlagData GetFeatureFlagById(ctx, featureFlagId).Authorization(authorization).Execute()

Get a Feature Flag by ID



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
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	featureFlagId := "featureFlagId_example" // string | The ID of the Feature Flag to fetch. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsAPI.GetFeatureFlagById(context.Background(), featureFlagId).Authorization(authorization).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.GetFeatureFlagById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFeatureFlagById`: FeatureFlagData
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsAPI.GetFeatureFlagById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**featureFlagId** | **string** | The ID of the Feature Flag to fetch.  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFeatureFlagByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 


### Return type

[**FeatureFlagData**](FeatureFlagData.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitFeatureFlags

> SubmitFeatureFlagsResponse SubmitFeatureFlags(ctx).Authorization(authorization).SubmitFeatureFlagRequest(submitFeatureFlagRequest).Execute()

Submit Feature Flag data



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
	authorization := "authorization_example" // string | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details. 
	submitFeatureFlagRequest := *openapiclient.NewSubmitFeatureFlagRequest([]openapiclient.FeatureFlagData{*openapiclient.NewFeatureFlagData("111-222-333", "my-awesome-feature", int64(1523494301448), *openapiclient.NewFeatureFlagSummary(*openapiclient.NewFeatureFlagStatus(false), time.Now()), []openapiclient.FeatureFlagDetails{*openapiclient.NewFeatureFlagDetails("https://example.com/project/feature-123/production", time.Now(), *openapiclient.NewEnvironmentDetails("Name_example"), *openapiclient.NewFeatureFlagStatus(false))})}) // SubmitFeatureFlagRequest | Feature Flag data to submit. 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FeatureFlagsAPI.SubmitFeatureFlags(context.Background()).Authorization(authorization).SubmitFeatureFlagRequest(submitFeatureFlagRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FeatureFlagsAPI.SubmitFeatureFlags``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SubmitFeatureFlags`: SubmitFeatureFlagsResponse
	fmt.Fprintf(os.Stdout, "Response from `FeatureFlagsAPI.SubmitFeatureFlags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSubmitFeatureFlagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | All requests must be signed with a Connect JWT token that corresponds to the Provider app installed in Jira.  If the JWT token corresponds to an app that does not define Feature Flags module it will be rejected with a 403.  See https://developer.atlassian.com/blog/2015/01/understanding-jwt/ for more details.  | 
 **submitFeatureFlagRequest** | [**SubmitFeatureFlagRequest**](SubmitFeatureFlagRequest.md) | Feature Flag data to submit.  | 

### Return type

[**SubmitFeatureFlagsResponse**](SubmitFeatureFlagsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

