# \TimeTrackingAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableTimeTrackingImplementations**](TimeTrackingAPI.md#GetAvailableTimeTrackingImplementations) | **Get** /rest/api/3/configuration/timetracking/list | Get all time tracking providers
[**GetSelectedTimeTrackingImplementation**](TimeTrackingAPI.md#GetSelectedTimeTrackingImplementation) | **Get** /rest/api/3/configuration/timetracking | Get selected time tracking provider
[**GetSharedTimeTrackingConfiguration**](TimeTrackingAPI.md#GetSharedTimeTrackingConfiguration) | **Get** /rest/api/3/configuration/timetracking/options | Get time tracking settings
[**SelectTimeTrackingImplementation**](TimeTrackingAPI.md#SelectTimeTrackingImplementation) | **Put** /rest/api/3/configuration/timetracking | Select time tracking provider
[**SetSharedTimeTrackingConfiguration**](TimeTrackingAPI.md#SetSharedTimeTrackingConfiguration) | **Put** /rest/api/3/configuration/timetracking/options | Set time tracking settings



## GetAvailableTimeTrackingImplementations

> []TimeTrackingProvider GetAvailableTimeTrackingImplementations(ctx).Execute()

Get all time tracking providers



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
	resp, r, err := apiClient.TimeTrackingAPI.GetAvailableTimeTrackingImplementations(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingAPI.GetAvailableTimeTrackingImplementations``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAvailableTimeTrackingImplementations`: []TimeTrackingProvider
	fmt.Fprintf(os.Stdout, "Response from `TimeTrackingAPI.GetAvailableTimeTrackingImplementations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAvailableTimeTrackingImplementationsRequest struct via the builder pattern


### Return type

[**[]TimeTrackingProvider**](TimeTrackingProvider.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSelectedTimeTrackingImplementation

> TimeTrackingProvider GetSelectedTimeTrackingImplementation(ctx).Execute()

Get selected time tracking provider



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
	resp, r, err := apiClient.TimeTrackingAPI.GetSelectedTimeTrackingImplementation(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingAPI.GetSelectedTimeTrackingImplementation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSelectedTimeTrackingImplementation`: TimeTrackingProvider
	fmt.Fprintf(os.Stdout, "Response from `TimeTrackingAPI.GetSelectedTimeTrackingImplementation`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSelectedTimeTrackingImplementationRequest struct via the builder pattern


### Return type

[**TimeTrackingProvider**](TimeTrackingProvider.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSharedTimeTrackingConfiguration

> TimeTrackingConfiguration GetSharedTimeTrackingConfiguration(ctx).Execute()

Get time tracking settings



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
	resp, r, err := apiClient.TimeTrackingAPI.GetSharedTimeTrackingConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingAPI.GetSharedTimeTrackingConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetSharedTimeTrackingConfiguration`: TimeTrackingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `TimeTrackingAPI.GetSharedTimeTrackingConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSharedTimeTrackingConfigurationRequest struct via the builder pattern


### Return type

[**TimeTrackingConfiguration**](TimeTrackingConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SelectTimeTrackingImplementation

> interface{} SelectTimeTrackingImplementation(ctx).TimeTrackingProvider(timeTrackingProvider).Execute()

Select time tracking provider



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
	timeTrackingProvider := *openapiclient.NewTimeTrackingProvider("Key_example") // TimeTrackingProvider | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeTrackingAPI.SelectTimeTrackingImplementation(context.Background()).TimeTrackingProvider(timeTrackingProvider).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingAPI.SelectTimeTrackingImplementation``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SelectTimeTrackingImplementation`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `TimeTrackingAPI.SelectTimeTrackingImplementation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSelectTimeTrackingImplementationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeTrackingProvider** | [**TimeTrackingProvider**](TimeTrackingProvider.md) |  | 

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


## SetSharedTimeTrackingConfiguration

> TimeTrackingConfiguration SetSharedTimeTrackingConfiguration(ctx).TimeTrackingConfiguration(timeTrackingConfiguration).Execute()

Set time tracking settings



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
	timeTrackingConfiguration := *openapiclient.NewTimeTrackingConfiguration("DefaultUnit_example", "TimeFormat_example", float64(123), float64(123)) // TimeTrackingConfiguration | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TimeTrackingAPI.SetSharedTimeTrackingConfiguration(context.Background()).TimeTrackingConfiguration(timeTrackingConfiguration).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TimeTrackingAPI.SetSharedTimeTrackingConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetSharedTimeTrackingConfiguration`: TimeTrackingConfiguration
	fmt.Fprintf(os.Stdout, "Response from `TimeTrackingAPI.SetSharedTimeTrackingConfiguration`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetSharedTimeTrackingConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **timeTrackingConfiguration** | [**TimeTrackingConfiguration**](TimeTrackingConfiguration.md) |  | 

### Return type

[**TimeTrackingConfiguration**](TimeTrackingConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

