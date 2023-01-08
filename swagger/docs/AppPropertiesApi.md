# \AppPropertiesAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddonPropertiesResourceDeleteAddonPropertyDelete**](AppPropertiesAPI.md#AddonPropertiesResourceDeleteAddonPropertyDelete) | **Delete** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Delete app property
[**AddonPropertiesResourceGetAddonPropertiesGet**](AppPropertiesAPI.md#AddonPropertiesResourceGetAddonPropertiesGet) | **Get** /rest/atlassian-connect/1/addons/{addonKey}/properties | Get app properties
[**AddonPropertiesResourceGetAddonPropertyGet**](AppPropertiesAPI.md#AddonPropertiesResourceGetAddonPropertyGet) | **Get** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Get app property
[**AddonPropertiesResourcePutAddonPropertyPut**](AppPropertiesAPI.md#AddonPropertiesResourcePutAddonPropertyPut) | **Put** /rest/atlassian-connect/1/addons/{addonKey}/properties/{propertyKey} | Set app property
[**DeleteForgeAppProperty**](AppPropertiesAPI.md#DeleteForgeAppProperty) | **Delete** /rest/forge/1/app/properties/{propertyKey} | Delete app property (Forge)
[**GetForgeAppProperty**](AppPropertiesAPI.md#GetForgeAppProperty) | **Get** /rest/forge/1/app/properties/{propertyKey} | Get app property (Forge)
[**GetForgeAppPropertyKeys**](AppPropertiesAPI.md#GetForgeAppPropertyKeys) | **Get** /rest/forge/1/app/properties | Get app property keys (Forge)
[**PutForgeAppProperty**](AppPropertiesAPI.md#PutForgeAppProperty) | **Put** /rest/forge/1/app/properties/{propertyKey} | Set app property (Forge)



## AddonPropertiesResourceDeleteAddonPropertyDelete

> AddonPropertiesResourceDeleteAddonPropertyDelete(ctx, addonKey, propertyKey).Execute()

Delete app property



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
	addonKey := "addonKey_example" // string | The key of the app, as defined in its descriptor.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppPropertiesAPI.AddonPropertiesResourceDeleteAddonPropertyDelete(context.Background(), addonKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.AddonPropertiesResourceDeleteAddonPropertyDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonKey** | **string** | The key of the app, as defined in its descriptor. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonPropertiesResourceDeleteAddonPropertyDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



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


## AddonPropertiesResourceGetAddonPropertiesGet

> PropertyKeys AddonPropertiesResourceGetAddonPropertiesGet(ctx, addonKey).Execute()

Get app properties



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
	addonKey := "addonKey_example" // string | The key of the app, as defined in its descriptor.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertiesGet(context.Background(), addonKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertiesGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonPropertiesResourceGetAddonPropertiesGet`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertiesGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonKey** | **string** | The key of the app, as defined in its descriptor. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonPropertiesResourceGetAddonPropertiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**PropertyKeys**](PropertyKeys.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonPropertiesResourceGetAddonPropertyGet

> EntityProperty AddonPropertiesResourceGetAddonPropertyGet(ctx, addonKey, propertyKey).Execute()

Get app property



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
	addonKey := "addonKey_example" // string | The key of the app, as defined in its descriptor.
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertyGet(context.Background(), addonKey, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertyGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonPropertiesResourceGetAddonPropertyGet`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.AddonPropertiesResourceGetAddonPropertyGet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonKey** | **string** | The key of the app, as defined in its descriptor. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonPropertiesResourceGetAddonPropertyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EntityProperty**](EntityProperty.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddonPropertiesResourcePutAddonPropertyPut

> OperationMessage AddonPropertiesResourcePutAddonPropertyPut(ctx, addonKey, propertyKey).Body(body).Execute()

Set app property



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
	addonKey := "addonKey_example" // string | The key of the app, as defined in its descriptor.
	propertyKey := "propertyKey_example" // string | The key of the property.
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.AddonPropertiesResourcePutAddonPropertyPut(context.Background(), addonKey, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.AddonPropertiesResourcePutAddonPropertyPut``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddonPropertiesResourcePutAddonPropertyPut`: OperationMessage
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.AddonPropertiesResourcePutAddonPropertyPut`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**addonKey** | **string** | The key of the app, as defined in its descriptor. | 
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddonPropertiesResourcePutAddonPropertyPutRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | **interface{}** |  | 

### Return type

[**OperationMessage**](OperationMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteForgeAppProperty

> DeleteForgeAppProperty(ctx, propertyKey).Execute()

Delete app property (Forge)



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
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.AppPropertiesAPI.DeleteForgeAppProperty(context.Background(), propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.DeleteForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## GetForgeAppProperty

> GetForgeAppProperty200Response GetForgeAppProperty(ctx, propertyKey).Execute()

Get app property (Forge)



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
	propertyKey := "propertyKey_example" // string | The key of the property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.GetForgeAppProperty(context.Background(), propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.GetForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForgeAppProperty`: GetForgeAppProperty200Response
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.GetForgeAppProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetForgeAppProperty200Response**](GetForgeAppProperty200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetForgeAppPropertyKeys

> GetForgeAppPropertyKeys200Response GetForgeAppPropertyKeys(ctx).Execute()

Get app property keys (Forge)



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
	resp, r, err := apiClient.AppPropertiesAPI.GetForgeAppPropertyKeys(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.GetForgeAppPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetForgeAppPropertyKeys`: GetForgeAppPropertyKeys200Response
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.GetForgeAppPropertyKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetForgeAppPropertyKeysRequest struct via the builder pattern


### Return type

[**GetForgeAppPropertyKeys200Response**](GetForgeAppPropertyKeys200Response.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutForgeAppProperty

> OperationMessage PutForgeAppProperty(ctx, propertyKey).Body(body).Execute()

Set app property (Forge)



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
	propertyKey := "propertyKey_example" // string | The key of the property.
	body := interface{}(987) // interface{} | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AppPropertiesAPI.PutForgeAppProperty(context.Background(), propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AppPropertiesAPI.PutForgeAppProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutForgeAppProperty`: OperationMessage
	fmt.Fprintf(os.Stdout, "Response from `AppPropertiesAPI.PutForgeAppProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**propertyKey** | **string** | The key of the property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutForgeAppPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | **interface{}** |  | 

### Return type

[**OperationMessage**](OperationMessage.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

