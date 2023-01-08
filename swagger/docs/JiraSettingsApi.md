# \JiraSettingsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAdvancedSettings**](JiraSettingsAPI.md#GetAdvancedSettings) | **Get** /rest/api/3/application-properties/advanced-settings | Get advanced settings
[**GetApplicationProperty**](JiraSettingsAPI.md#GetApplicationProperty) | **Get** /rest/api/3/application-properties | Get application property
[**GetConfiguration**](JiraSettingsAPI.md#GetConfiguration) | **Get** /rest/api/3/configuration | Get global settings
[**SetApplicationProperty**](JiraSettingsAPI.md#SetApplicationProperty) | **Put** /rest/api/3/application-properties/{id} | Set application property



## GetAdvancedSettings

> []ApplicationProperty GetAdvancedSettings(ctx).Execute()

Get advanced settings



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
	resp, r, err := apiClient.JiraSettingsAPI.GetAdvancedSettings(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraSettingsAPI.GetAdvancedSettings``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAdvancedSettings`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `JiraSettingsAPI.GetAdvancedSettings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdvancedSettingsRequest struct via the builder pattern


### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationProperty

> []ApplicationProperty GetApplicationProperty(ctx).Key(key).PermissionLevel(permissionLevel).KeyFilter(keyFilter).Execute()

Get application property



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
	key := "key_example" // string | The key of the application property. (optional)
	permissionLevel := "permissionLevel_example" // string | The permission level of all items being returned in the list. (optional)
	keyFilter := "keyFilter_example" // string | When a `key` isn't provided, this filters the list of results by the application property `key` using a regular expression. For example, using `jira.lf.*` will return all application properties with keys that start with *jira.lf.*. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraSettingsAPI.GetApplicationProperty(context.Background()).Key(key).PermissionLevel(permissionLevel).KeyFilter(keyFilter).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraSettingsAPI.GetApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationProperty`: []ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `JiraSettingsAPI.GetApplicationProperty`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **key** | **string** | The key of the application property. | 
 **permissionLevel** | **string** | The permission level of all items being returned in the list. | 
 **keyFilter** | **string** | When a &#x60;key&#x60; isn&#39;t provided, this filters the list of results by the application property &#x60;key&#x60; using a regular expression. For example, using &#x60;jira.lf.*&#x60; will return all application properties with keys that start with *jira.lf.*. | 

### Return type

[**[]ApplicationProperty**](ApplicationProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConfiguration

> JiraConfiguration GetConfiguration(ctx).Execute()

Get global settings



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
	resp, r, err := apiClient.JiraSettingsAPI.GetConfiguration(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraSettingsAPI.GetConfiguration``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetConfiguration`: JiraConfiguration
	fmt.Fprintf(os.Stdout, "Response from `JiraSettingsAPI.GetConfiguration`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigurationRequest struct via the builder pattern


### Return type

[**JiraConfiguration**](JiraConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApplicationProperty

> ApplicationProperty SetApplicationProperty(ctx, id).SimpleApplicationPropertyBean(simpleApplicationPropertyBean).Execute()

Set application property



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
	id := "id_example" // string | The key of the application property to update.
	simpleApplicationPropertyBean := *openapiclient.NewSimpleApplicationPropertyBean() // SimpleApplicationPropertyBean | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.JiraSettingsAPI.SetApplicationProperty(context.Background(), id).SimpleApplicationPropertyBean(simpleApplicationPropertyBean).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `JiraSettingsAPI.SetApplicationProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetApplicationProperty`: ApplicationProperty
	fmt.Fprintf(os.Stdout, "Response from `JiraSettingsAPI.SetApplicationProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The key of the application property to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApplicationPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **simpleApplicationPropertyBean** | [**SimpleApplicationPropertyBean**](SimpleApplicationPropertyBean.md) |  | 

### Return type

[**ApplicationProperty**](ApplicationProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

