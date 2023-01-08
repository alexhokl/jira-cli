# \DashboardsAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddGadget**](DashboardsAPI.md#AddGadget) | **Post** /rest/api/3/dashboard/{dashboardId}/gadget | Add gadget to dashboard
[**BulkEditDashboards**](DashboardsAPI.md#BulkEditDashboards) | **Put** /rest/api/3/dashboard/bulk/edit | Bulk edit dashboards
[**CopyDashboard**](DashboardsAPI.md#CopyDashboard) | **Post** /rest/api/3/dashboard/{id}/copy | Copy dashboard
[**CreateDashboard**](DashboardsAPI.md#CreateDashboard) | **Post** /rest/api/3/dashboard | Create dashboard
[**DeleteDashboard**](DashboardsAPI.md#DeleteDashboard) | **Delete** /rest/api/3/dashboard/{id} | Delete dashboard
[**DeleteDashboardItemProperty**](DashboardsAPI.md#DeleteDashboardItemProperty) | **Delete** /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey} | Delete dashboard item property
[**GetAllAvailableDashboardGadgets**](DashboardsAPI.md#GetAllAvailableDashboardGadgets) | **Get** /rest/api/3/dashboard/gadgets | Get available gadgets
[**GetAllDashboards**](DashboardsAPI.md#GetAllDashboards) | **Get** /rest/api/3/dashboard | Get all dashboards
[**GetAllGadgets**](DashboardsAPI.md#GetAllGadgets) | **Get** /rest/api/3/dashboard/{dashboardId}/gadget | Get gadgets
[**GetDashboard**](DashboardsAPI.md#GetDashboard) | **Get** /rest/api/3/dashboard/{id} | Get dashboard
[**GetDashboardItemProperty**](DashboardsAPI.md#GetDashboardItemProperty) | **Get** /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey} | Get dashboard item property
[**GetDashboardItemPropertyKeys**](DashboardsAPI.md#GetDashboardItemPropertyKeys) | **Get** /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties | Get dashboard item property keys
[**GetDashboardsPaginated**](DashboardsAPI.md#GetDashboardsPaginated) | **Get** /rest/api/3/dashboard/search | Search for dashboards
[**RemoveGadget**](DashboardsAPI.md#RemoveGadget) | **Delete** /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId} | Remove gadget from dashboard
[**SetDashboardItemProperty**](DashboardsAPI.md#SetDashboardItemProperty) | **Put** /rest/api/3/dashboard/{dashboardId}/items/{itemId}/properties/{propertyKey} | Set dashboard item property
[**UpdateDashboard**](DashboardsAPI.md#UpdateDashboard) | **Put** /rest/api/3/dashboard/{id} | Update dashboard
[**UpdateGadget**](DashboardsAPI.md#UpdateGadget) | **Put** /rest/api/3/dashboard/{dashboardId}/gadget/{gadgetId} | Update gadget on dashboard



## AddGadget

> DashboardGadget AddGadget(ctx, dashboardId).DashboardGadgetSettings(dashboardGadgetSettings).Execute()

Add gadget to dashboard



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
	dashboardId := int64(789) // int64 | The ID of the dashboard.
	dashboardGadgetSettings := *openapiclient.NewDashboardGadgetSettings() // DashboardGadgetSettings | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.AddGadget(context.Background(), dashboardId).DashboardGadgetSettings(dashboardGadgetSettings).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.AddGadget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `AddGadget`: DashboardGadget
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.AddGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardGadgetSettings** | [**DashboardGadgetSettings**](DashboardGadgetSettings.md) |  | 

### Return type

[**DashboardGadget**](DashboardGadget.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BulkEditDashboards

> BulkEditShareableEntityResponse BulkEditDashboards(ctx).BulkEditShareableEntityRequest(bulkEditShareableEntityRequest).Execute()

Bulk edit dashboards



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
	bulkEditShareableEntityRequest := *openapiclient.NewBulkEditShareableEntityRequest("Action_example", []int64{int64(123)}) // BulkEditShareableEntityRequest | The details of dashboards being updated in bulk.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.BulkEditDashboards(context.Background()).BulkEditShareableEntityRequest(bulkEditShareableEntityRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.BulkEditDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `BulkEditDashboards`: BulkEditShareableEntityResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.BulkEditDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBulkEditDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bulkEditShareableEntityRequest** | [**BulkEditShareableEntityRequest**](BulkEditShareableEntityRequest.md) | The details of dashboards being updated in bulk. | 

### Return type

[**BulkEditShareableEntityResponse**](BulkEditShareableEntityResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CopyDashboard

> Dashboard CopyDashboard(ctx, id).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()

Copy dashboard



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
	id := "id_example" // string | 
	dashboardDetails := *openapiclient.NewDashboardDetails([]openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}, "Name_example", []openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}) // DashboardDetails | Dashboard details.
	extendAdminPermissions := true // bool | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.CopyDashboard(context.Background(), id).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CopyDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CopyDashboard`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CopyDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCopyDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardDetails** | [**DashboardDetails**](DashboardDetails.md) | Dashboard details. | 
 **extendAdminPermissions** | **bool** | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) | [default to false]

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateDashboard

> Dashboard CreateDashboard(ctx).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()

Create dashboard



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
	dashboardDetails := *openapiclient.NewDashboardDetails([]openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}, "Name_example", []openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}) // DashboardDetails | Dashboard details.
	extendAdminPermissions := true // bool | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.CreateDashboard(context.Background()).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.CreateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDashboard`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.CreateDashboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardDetails** | [**DashboardDetails**](DashboardDetails.md) | Dashboard details. | 
 **extendAdminPermissions** | **bool** | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) | [default to false]

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDashboard

> DeleteDashboard(ctx, id).Execute()

Delete dashboard



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
	id := "id_example" // string | The ID of the dashboard.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DashboardsAPI.DeleteDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardRequest struct via the builder pattern


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


## DeleteDashboardItemProperty

> interface{} DeleteDashboardItemProperty(ctx, dashboardId, itemId, propertyKey).Execute()

Delete dashboard item property



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
	dashboardId := "dashboardId_example" // string | The ID of the dashboard.
	itemId := "itemId_example" // string | The ID of the dashboard item.
	propertyKey := "propertyKey_example" // string | The key of the dashboard item property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.DeleteDashboardItemProperty(context.Background(), dashboardId, itemId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.DeleteDashboardItemProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDashboardItemProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.DeleteDashboardItemProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 
**itemId** | **string** | The ID of the dashboard item. | 
**propertyKey** | **string** | The key of the dashboard item property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDashboardItemPropertyRequest struct via the builder pattern


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


## GetAllAvailableDashboardGadgets

> AvailableDashboardGadgetsResponse GetAllAvailableDashboardGadgets(ctx).Execute()

Get available gadgets



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
	resp, r, err := apiClient.DashboardsAPI.GetAllAvailableDashboardGadgets(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetAllAvailableDashboardGadgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllAvailableDashboardGadgets`: AvailableDashboardGadgetsResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetAllAvailableDashboardGadgets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAvailableDashboardGadgetsRequest struct via the builder pattern


### Return type

[**AvailableDashboardGadgetsResponse**](AvailableDashboardGadgetsResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllDashboards

> PageOfDashboards GetAllDashboards(ctx).Filter(filter).StartAt(startAt).MaxResults(maxResults).Execute()

Get all dashboards



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
	filter := "filter_example" // string | The filter applied to the list of dashboards. Valid values are:   *  `favourite` Returns dashboards the user has marked as favorite.  *  `my` Returns dashboards owned by the user. (optional)
	startAt := int32(56) // int32 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 20)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetAllDashboards(context.Background()).Filter(filter).StartAt(startAt).MaxResults(maxResults).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetAllDashboards``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllDashboards`: PageOfDashboards
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetAllDashboards`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllDashboardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | **string** | The filter applied to the list of dashboards. Valid values are:   *  &#x60;favourite&#x60; Returns dashboards the user has marked as favorite.  *  &#x60;my&#x60; Returns dashboards owned by the user. | 
 **startAt** | **int32** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 20]

### Return type

[**PageOfDashboards**](PageOfDashboards.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllGadgets

> DashboardGadgetResponse GetAllGadgets(ctx, dashboardId).ModuleKey(moduleKey).Uri(uri).GadgetId(gadgetId).Execute()

Get gadgets



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
	dashboardId := int64(789) // int64 | The ID of the dashboard.
	moduleKey := []string{"Inner_example"} // []string | The list of gadgets module keys. To include multiple module keys, separate module keys with ampersand: `moduleKey=key:one&moduleKey=key:two`. (optional)
	uri := []string{"Inner_example"} // []string | The list of gadgets URIs. To include multiple URIs, separate URIs with ampersand: `uri=/rest/example/uri/1&uri=/rest/example/uri/2`. (optional)
	gadgetId := []int64{int64(123)} // []int64 | The list of gadgets IDs. To include multiple IDs, separate IDs with ampersand: `gadgetId=10000&gadgetId=10001`. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetAllGadgets(context.Background(), dashboardId).ModuleKey(moduleKey).Uri(uri).GadgetId(gadgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetAllGadgets``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetAllGadgets`: DashboardGadgetResponse
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetAllGadgets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllGadgetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **moduleKey** | **[]string** | The list of gadgets module keys. To include multiple module keys, separate module keys with ampersand: &#x60;moduleKey&#x3D;key:one&amp;moduleKey&#x3D;key:two&#x60;. | 
 **uri** | **[]string** | The list of gadgets URIs. To include multiple URIs, separate URIs with ampersand: &#x60;uri&#x3D;/rest/example/uri/1&amp;uri&#x3D;/rest/example/uri/2&#x60;. | 
 **gadgetId** | **[]int64** | The list of gadgets IDs. To include multiple IDs, separate IDs with ampersand: &#x60;gadgetId&#x3D;10000&amp;gadgetId&#x3D;10001&#x60;. | 

### Return type

[**DashboardGadgetResponse**](DashboardGadgetResponse.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboard

> Dashboard GetDashboard(ctx, id).Execute()

Get dashboard



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
	id := "id_example" // string | The ID of the dashboard.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetDashboard(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboard`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDashboardItemProperty

> EntityProperty GetDashboardItemProperty(ctx, dashboardId, itemId, propertyKey).Execute()

Get dashboard item property



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
	dashboardId := "dashboardId_example" // string | The ID of the dashboard.
	itemId := "itemId_example" // string | The ID of the dashboard item.
	propertyKey := "propertyKey_example" // string | The key of the dashboard item property.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetDashboardItemProperty(context.Background(), dashboardId, itemId, propertyKey).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardItemProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardItemProperty`: EntityProperty
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardItemProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 
**itemId** | **string** | The ID of the dashboard item. | 
**propertyKey** | **string** | The key of the dashboard item property. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardItemPropertyRequest struct via the builder pattern


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


## GetDashboardItemPropertyKeys

> PropertyKeys GetDashboardItemPropertyKeys(ctx, dashboardId, itemId).Execute()

Get dashboard item property keys



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
	dashboardId := "dashboardId_example" // string | The ID of the dashboard.
	itemId := "itemId_example" // string | The ID of the dashboard item.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetDashboardItemPropertyKeys(context.Background(), dashboardId, itemId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardItemPropertyKeys``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardItemPropertyKeys`: PropertyKeys
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardItemPropertyKeys`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 
**itemId** | **string** | The ID of the dashboard item. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardItemPropertyKeysRequest struct via the builder pattern


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


## GetDashboardsPaginated

> PageBeanDashboard GetDashboardsPaginated(ctx).DashboardName(dashboardName).AccountId(accountId).Owner(owner).Groupname(groupname).GroupId(groupId).ProjectId(projectId).OrderBy(orderBy).StartAt(startAt).MaxResults(maxResults).Status(status).Expand(expand).Execute()

Search for dashboards



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
	dashboardName := "dashboardName_example" // string | String used to perform a case-insensitive partial match with `name`. (optional)
	accountId := "accountId_example" // string | User account ID used to return dashboards with the matching `owner.accountId`. This parameter cannot be used with the `owner` parameter. (optional)
	owner := "owner_example" // string | This parameter is deprecated because of privacy changes. Use `accountId` instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. User name used to return dashboards with the matching `owner.name`. This parameter cannot be used with the `accountId` parameter. (optional)
	groupname := "groupname_example" // string | As a group's name can change, use of `groupId` is recommended. Group name used to return dashboards that are shared with a group that matches `sharePermissions.group.name`. This parameter cannot be used with the `groupId` parameter. (optional)
	groupId := "groupId_example" // string | Group ID used to return dashboards that are shared with a group that matches `sharePermissions.group.groupId`. This parameter cannot be used with the `groupname` parameter. (optional)
	projectId := int64(789) // int64 | Project ID used to returns dashboards that are shared with a project that matches `sharePermissions.project.id`. (optional)
	orderBy := "orderBy_example" // string | [Order](#ordering) the results by a field:   *  `description` Sorts by dashboard description. Note that this sort works independently of whether the expand to display the description field is in use.  *  `favourite_count` Sorts by dashboard popularity.  *  `id` Sorts by dashboard ID.  *  `is_favourite` Sorts by whether the dashboard is marked as a favorite.  *  `name` Sorts by dashboard name.  *  `owner` Sorts by dashboard owner name. (optional) (default to "name")
	startAt := int64(789) // int64 | The index of the first item to return in a page of results (page offset). (optional) (default to 0)
	maxResults := int32(56) // int32 | The maximum number of items to return per page. (optional) (default to 50)
	status := "status_example" // string | The status to filter by. It may be active, archived or deleted. (optional) (default to "active")
	expand := "expand_example" // string | Use [expand](#expansion) to include additional information about dashboard in the response. This parameter accepts a comma-separated list. Expand options include:   *  `description` Returns the description of the dashboard.  *  `owner` Returns the owner of the dashboard.  *  `viewUrl` Returns the URL that is used to view the dashboard.  *  `favourite` Returns `isFavourite`, an indicator of whether the user has set the dashboard as a favorite.  *  `favouritedCount` Returns `popularity`, a count of how many users have set this dashboard as a favorite.  *  `sharePermissions` Returns details of the share permissions defined for the dashboard.  *  `editPermissions` Returns details of the edit permissions defined for the dashboard.  *  `isWritable` Returns whether the current user has permission to edit the dashboard. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.GetDashboardsPaginated(context.Background()).DashboardName(dashboardName).AccountId(accountId).Owner(owner).Groupname(groupname).GroupId(groupId).ProjectId(projectId).OrderBy(orderBy).StartAt(startAt).MaxResults(maxResults).Status(status).Expand(expand).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.GetDashboardsPaginated``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDashboardsPaginated`: PageBeanDashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.GetDashboardsPaginated`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDashboardsPaginatedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardName** | **string** | String used to perform a case-insensitive partial match with &#x60;name&#x60;. | 
 **accountId** | **string** | User account ID used to return dashboards with the matching &#x60;owner.accountId&#x60;. This parameter cannot be used with the &#x60;owner&#x60; parameter. | 
 **owner** | **string** | This parameter is deprecated because of privacy changes. Use &#x60;accountId&#x60; instead. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. User name used to return dashboards with the matching &#x60;owner.name&#x60;. This parameter cannot be used with the &#x60;accountId&#x60; parameter. | 
 **groupname** | **string** | As a group&#39;s name can change, use of &#x60;groupId&#x60; is recommended. Group name used to return dashboards that are shared with a group that matches &#x60;sharePermissions.group.name&#x60;. This parameter cannot be used with the &#x60;groupId&#x60; parameter. | 
 **groupId** | **string** | Group ID used to return dashboards that are shared with a group that matches &#x60;sharePermissions.group.groupId&#x60;. This parameter cannot be used with the &#x60;groupname&#x60; parameter. | 
 **projectId** | **int64** | Project ID used to returns dashboards that are shared with a project that matches &#x60;sharePermissions.project.id&#x60;. | 
 **orderBy** | **string** | [Order](#ordering) the results by a field:   *  &#x60;description&#x60; Sorts by dashboard description. Note that this sort works independently of whether the expand to display the description field is in use.  *  &#x60;favourite_count&#x60; Sorts by dashboard popularity.  *  &#x60;id&#x60; Sorts by dashboard ID.  *  &#x60;is_favourite&#x60; Sorts by whether the dashboard is marked as a favorite.  *  &#x60;name&#x60; Sorts by dashboard name.  *  &#x60;owner&#x60; Sorts by dashboard owner name. | [default to &quot;name&quot;]
 **startAt** | **int64** | The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **int32** | The maximum number of items to return per page. | [default to 50]
 **status** | **string** | The status to filter by. It may be active, archived or deleted. | [default to &quot;active&quot;]
 **expand** | **string** | Use [expand](#expansion) to include additional information about dashboard in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;description&#x60; Returns the description of the dashboard.  *  &#x60;owner&#x60; Returns the owner of the dashboard.  *  &#x60;viewUrl&#x60; Returns the URL that is used to view the dashboard.  *  &#x60;favourite&#x60; Returns &#x60;isFavourite&#x60;, an indicator of whether the user has set the dashboard as a favorite.  *  &#x60;favouritedCount&#x60; Returns &#x60;popularity&#x60;, a count of how many users have set this dashboard as a favorite.  *  &#x60;sharePermissions&#x60; Returns details of the share permissions defined for the dashboard.  *  &#x60;editPermissions&#x60; Returns details of the edit permissions defined for the dashboard.  *  &#x60;isWritable&#x60; Returns whether the current user has permission to edit the dashboard. | 

### Return type

[**PageBeanDashboard**](PageBeanDashboard.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveGadget

> interface{} RemoveGadget(ctx, dashboardId, gadgetId).Execute()

Remove gadget from dashboard



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
	dashboardId := int64(789) // int64 | The ID of the dashboard.
	gadgetId := int64(789) // int64 | The ID of the gadget.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.RemoveGadget(context.Background(), dashboardId, gadgetId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.RemoveGadget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RemoveGadget`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.RemoveGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The ID of the dashboard. | 
**gadgetId** | **int64** | The ID of the gadget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveGadgetRequest struct via the builder pattern


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


## SetDashboardItemProperty

> interface{} SetDashboardItemProperty(ctx, dashboardId, itemId, propertyKey).Body(body).Execute()

Set dashboard item property



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
	dashboardId := "dashboardId_example" // string | The ID of the dashboard.
	itemId := "itemId_example" // string | The ID of the dashboard item.
	propertyKey := "propertyKey_example" // string | The key of the dashboard item property. The maximum length is 255 characters. For dashboard items with a spec URI and no complete module key, if the provided propertyKey is equal to \"config\", the request body's JSON must be an object with all keys and values as strings.
	body := interface{}({"number":5,"string":"string-value"}) // interface{} | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.SetDashboardItemProperty(context.Background(), dashboardId, itemId, propertyKey).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.SetDashboardItemProperty``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetDashboardItemProperty`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.SetDashboardItemProperty`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **string** | The ID of the dashboard. | 
**itemId** | **string** | The ID of the dashboard item. | 
**propertyKey** | **string** | The key of the dashboard item property. The maximum length is 255 characters. For dashboard items with a spec URI and no complete module key, if the provided propertyKey is equal to \&quot;config\&quot;, the request body&#39;s JSON must be an object with all keys and values as strings. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDashboardItemPropertyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | **interface{}** | The value of the property. The value has to be a valid, non-empty [JSON](https://tools.ietf.org/html/rfc4627) value. The maximum length of the property value is 32768 bytes. | 

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


## UpdateDashboard

> Dashboard UpdateDashboard(ctx, id).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()

Update dashboard



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
	id := "id_example" // string | The ID of the dashboard to update.
	dashboardDetails := *openapiclient.NewDashboardDetails([]openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}, "Name_example", []openapiclient.SharePermission{*openapiclient.NewSharePermission("Type_example")}) // DashboardDetails | Replacement dashboard details.
	extendAdminPermissions := true // bool | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) (optional) (default to false)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.UpdateDashboard(context.Background(), id).DashboardDetails(dashboardDetails).ExtendAdminPermissions(extendAdminPermissions).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateDashboard``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDashboard`: Dashboard
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateDashboard`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The ID of the dashboard to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDashboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardDetails** | [**DashboardDetails**](DashboardDetails.md) | Replacement dashboard details. | 
 **extendAdminPermissions** | **bool** | Whether admin level permissions are used. It should only be true if the user has *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg) | [default to false]

### Return type

[**Dashboard**](Dashboard.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGadget

> interface{} UpdateGadget(ctx, dashboardId, gadgetId).DashboardGadgetUpdateRequest(dashboardGadgetUpdateRequest).Execute()

Update gadget on dashboard



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
	dashboardId := int64(789) // int64 | The ID of the dashboard.
	gadgetId := int64(789) // int64 | The ID of the gadget.
	dashboardGadgetUpdateRequest := *openapiclient.NewDashboardGadgetUpdateRequest() // DashboardGadgetUpdateRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DashboardsAPI.UpdateGadget(context.Background(), dashboardId, gadgetId).DashboardGadgetUpdateRequest(dashboardGadgetUpdateRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DashboardsAPI.UpdateGadget``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateGadget`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `DashboardsAPI.UpdateGadget`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dashboardId** | **int64** | The ID of the dashboard. | 
**gadgetId** | **int64** | The ID of the gadget. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGadgetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dashboardGadgetUpdateRequest** | [**DashboardGadgetUpdateRequest**](DashboardGadgetUpdateRequest.md) |  | 

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

