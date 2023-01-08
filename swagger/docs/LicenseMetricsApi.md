# \LicenseMetricsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApproximateApplicationLicenseCount**](LicenseMetricsAPI.md#GetApproximateApplicationLicenseCount) | **Get** /rest/api/3/license/approximateLicenseCount/product/{applicationKey} | Get approximate application license count
[**GetApproximateLicenseCount**](LicenseMetricsAPI.md#GetApproximateLicenseCount) | **Get** /rest/api/3/license/approximateLicenseCount | Get approximate license count
[**GetLicense**](LicenseMetricsAPI.md#GetLicense) | **Get** /rest/api/3/instance/license | Get license



## GetApproximateApplicationLicenseCount

> LicenseMetric GetApproximateApplicationLicenseCount(ctx, applicationKey).Execute()

Get approximate application license count



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
	applicationKey := "applicationKey_example" // string | The ID of the application, represents a specific version of Jira.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.LicenseMetricsAPI.GetApproximateApplicationLicenseCount(context.Background(), applicationKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseMetricsAPI.GetApproximateApplicationLicenseCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApproximateApplicationLicenseCount`: LicenseMetric
	fmt.Fprintf(os.Stdout, "Response from `LicenseMetricsAPI.GetApproximateApplicationLicenseCount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationKey** | **string** | The ID of the application, represents a specific version of Jira. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApproximateApplicationLicenseCountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**LicenseMetric**](LicenseMetric.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApproximateLicenseCount

> LicenseMetric GetApproximateLicenseCount(ctx).Execute()

Get approximate license count



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
	resp, r, err := apiClient.LicenseMetricsAPI.GetApproximateLicenseCount(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseMetricsAPI.GetApproximateLicenseCount``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApproximateLicenseCount`: LicenseMetric
	fmt.Fprintf(os.Stdout, "Response from `LicenseMetricsAPI.GetApproximateLicenseCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetApproximateLicenseCountRequest struct via the builder pattern


### Return type

[**LicenseMetric**](LicenseMetric.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLicense

> License GetLicense(ctx).Execute()

Get license



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
	resp, r, err := apiClient.LicenseMetricsAPI.GetLicense(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `LicenseMetricsAPI.GetLicense``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetLicense`: License
	fmt.Fprintf(os.Stdout, "Response from `LicenseMetricsAPI.GetLicense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLicenseRequest struct via the builder pattern


### Return type

[**License**](License.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

