# \AnnouncementBannerAPI

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBanner**](AnnouncementBannerAPI.md#GetBanner) | **Get** /rest/api/3/announcementBanner | Get announcement banner configuration
[**SetBanner**](AnnouncementBannerAPI.md#SetBanner) | **Put** /rest/api/3/announcementBanner | Update announcement banner configuration



## GetBanner

> AnnouncementBannerConfiguration GetBanner(ctx).Execute()

Get announcement banner configuration



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
	resp, r, err := apiClient.AnnouncementBannerAPI.GetBanner(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementBannerAPI.GetBanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBanner`: AnnouncementBannerConfiguration
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementBannerAPI.GetBanner`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBannerRequest struct via the builder pattern


### Return type

[**AnnouncementBannerConfiguration**](AnnouncementBannerConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetBanner

> interface{} SetBanner(ctx).AnnouncementBannerConfigurationUpdate(announcementBannerConfigurationUpdate).Execute()

Update announcement banner configuration



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
	announcementBannerConfigurationUpdate := *openapiclient.NewAnnouncementBannerConfigurationUpdate() // AnnouncementBannerConfigurationUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.AnnouncementBannerAPI.SetBanner(context.Background()).AnnouncementBannerConfigurationUpdate(announcementBannerConfigurationUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `AnnouncementBannerAPI.SetBanner``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `SetBanner`: interface{}
	fmt.Fprintf(os.Stdout, "Response from `AnnouncementBannerAPI.SetBanner`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSetBannerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **announcementBannerConfigurationUpdate** | [**AnnouncementBannerConfigurationUpdate**](AnnouncementBannerConfigurationUpdate.md) |  | 

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

