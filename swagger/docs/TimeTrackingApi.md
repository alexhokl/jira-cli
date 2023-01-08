# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAvailableTimeTrackingImplementations**](TimeTrackingApi.md#GetAvailableTimeTrackingImplementations) | **Get** /rest/api/3/configuration/timetracking/list | Get all time tracking providers
[**GetSelectedTimeTrackingImplementation**](TimeTrackingApi.md#GetSelectedTimeTrackingImplementation) | **Get** /rest/api/3/configuration/timetracking | Get selected time tracking provider
[**GetSharedTimeTrackingConfiguration**](TimeTrackingApi.md#GetSharedTimeTrackingConfiguration) | **Get** /rest/api/3/configuration/timetracking/options | Get time tracking settings
[**SelectTimeTrackingImplementation**](TimeTrackingApi.md#SelectTimeTrackingImplementation) | **Put** /rest/api/3/configuration/timetracking | Select time tracking provider
[**SetSharedTimeTrackingConfiguration**](TimeTrackingApi.md#SetSharedTimeTrackingConfiguration) | **Put** /rest/api/3/configuration/timetracking/options | Set time tracking settings

# **GetAvailableTimeTrackingImplementations**
> []TimeTrackingProvider GetAvailableTimeTrackingImplementations(ctx, )
Get all time tracking providers

Returns all time tracking providers. By default, Jira only has one time tracking provider: *JIRA provided time tracking*. However, you can install other time tracking providers via apps from the Atlassian Marketplace. For more information on time tracking providers, see the documentation for the [ Time Tracking Provider](https://developer.atlassian.com/cloud/jira/platform/modules/time-tracking-provider/) module.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]TimeTrackingProvider**](TimeTrackingProvider.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSelectedTimeTrackingImplementation**
> TimeTrackingProvider GetSelectedTimeTrackingImplementation(ctx, )
Get selected time tracking provider

Returns the time tracking provider that is currently selected. Note that if time tracking is disabled, then a successful but empty response is returned.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TimeTrackingProvider**](TimeTrackingProvider.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSharedTimeTrackingConfiguration**
> TimeTrackingConfiguration GetSharedTimeTrackingConfiguration(ctx, )
Get time tracking settings

Returns the time tracking settings. This includes settings such as the time format, default time unit, and others. For more information, see [Configuring time tracking](https://confluence.atlassian.com/x/qoXKM).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**TimeTrackingConfiguration**](TimeTrackingConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SelectTimeTrackingImplementation**
> Object SelectTimeTrackingImplementation(ctx, body)
Select time tracking provider

Selects a time tracking provider.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimeTrackingProvider**](TimeTrackingProvider.md)|  | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetSharedTimeTrackingConfiguration**
> TimeTrackingConfiguration SetSharedTimeTrackingConfiguration(ctx, body)
Set time tracking settings

Sets the time tracking settings.  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TimeTrackingConfiguration**](TimeTrackingConfiguration.md)|  | 

### Return type

[**TimeTrackingConfiguration**](TimeTrackingConfiguration.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

