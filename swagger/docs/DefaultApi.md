# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPrecomputations**](DefaultApi.md#GetPrecomputations) | **Get** /rest/api/3/jql/function/computation | Get precomputation
[**UpdatePrecomputations**](DefaultApi.md#UpdatePrecomputations) | **Post** /rest/api/3/jql/function/computation | Update precomputations

# **GetPrecomputations**
> PageBeanJqlFunctionPrecomputationBean GetPrecomputations(ctx, optional)
Get precomputation

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***DefaultApiGetPrecomputationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a DefaultApiGetPrecomputationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **functionKey** | [**optional.Interface of []string**](string.md)|  | 
 **startAt** | **optional.Int64**|  | [default to 0]
 **maxResults** | **optional.Int32**|  | [default to 5000]
 **orderBy** | **optional.String**|  | 
 **filter** | **optional.String**|  | 

### Return type

[**PageBeanJqlFunctionPrecomputationBean**](PageBeanJqlFunctionPrecomputationBean.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdatePrecomputations**
> Object UpdatePrecomputations(ctx, body)
Update precomputations

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**JqlFunctionPrecomputationUpdateRequestBean**](JqlFunctionPrecomputationUpdateRequestBean.md)|  | 

### Return type

**Object**

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

