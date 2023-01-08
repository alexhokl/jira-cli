# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesApi.md#DeleteWorkflowTransitionRuleConfigurations) | **Put** /rest/api/3/workflow/rule/config/delete | Delete workflow transition rule configurations
[**GetWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesApi.md#GetWorkflowTransitionRuleConfigurations) | **Get** /rest/api/3/workflow/rule/config | Get workflow transition rule configurations
[**UpdateWorkflowTransitionRuleConfigurations**](WorkflowTransitionRulesApi.md#UpdateWorkflowTransitionRuleConfigurations) | **Put** /rest/api/3/workflow/rule/config | Update workflow transition rule configurations

# **DeleteWorkflowTransitionRuleConfigurations**
> WorkflowTransitionRulesUpdateErrors DeleteWorkflowTransitionRuleConfigurations(ctx, body)
Delete workflow transition rule configurations

Deletes workflow transition rules from one or more workflows. These rule types are supported:   *  [post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/)  *  [conditions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-condition/)  *  [validators](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-validator/)  Only rules created by the calling Connect app can be deleted.  **[Permissions](#permissions) required:** Only Connect apps can use this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WorkflowsWithTransitionRulesDetails**](WorkflowsWithTransitionRulesDetails.md)|  | 

### Return type

[**WorkflowTransitionRulesUpdateErrors**](WorkflowTransitionRulesUpdateErrors.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowTransitionRuleConfigurations**
> PageBeanWorkflowTransitionRules GetWorkflowTransitionRuleConfigurations(ctx, types, optional)
Get workflow transition rule configurations

Returns a [paginated](#pagination) list of workflows with transition rules. The workflows can be filtered to return only those containing workflow transition rules:   *  of one or more transition rule types, such as [workflow post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/).  *  matching one or more transition rule keys.  Only workflows containing transition rules created by the calling Connect app are returned.  Due to server-side optimizations, workflows with an empty list of rules may be returned; these workflows can be ignored.  **[Permissions](#permissions) required:** Only Connect apps can use this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **types** | [**[]string**](string.md)| The types of the transition rules to return. | 
 **optional** | ***WorkflowTransitionRulesApiGetWorkflowTransitionRuleConfigurationsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowTransitionRulesApiGetWorkflowTransitionRuleConfigurationsOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startAt** | **optional.Int64**| The index of the first item to return in a page of results (page offset). | [default to 0]
 **maxResults** | **optional.Int32**| The maximum number of items to return per page. | [default to 10]
 **keys** | [**optional.Interface of []string**](string.md)| The transition rule class keys, as defined in the Connect app descriptor, of the transition rules to return. | 
 **workflowNames** | [**optional.Interface of []string**](string.md)| EXPERIMENTAL: The list of workflow names to filter by. | 
 **withTags** | [**optional.Interface of []string**](string.md)| EXPERIMENTAL: The list of &#x60;tags&#x60; to filter by. | 
 **draft** | **optional.Bool**| EXPERIMENTAL: Whether draft or published workflows are returned. If not provided, both workflow types are returned. | 
 **expand** | **optional.String**| Use [expand](#expansion) to include additional information in the response. This parameter accepts &#x60;transition&#x60;, which, for each rule, returns information about the transition the rule is assigned to. | 

### Return type

[**PageBeanWorkflowTransitionRules**](PageBeanWorkflowTransitionRules.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWorkflowTransitionRuleConfigurations**
> WorkflowTransitionRulesUpdateErrors UpdateWorkflowTransitionRuleConfigurations(ctx, body)
Update workflow transition rule configurations

Updates configuration of workflow transition rules. The following rule types are supported:   *  [post functions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-post-function/)  *  [conditions](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-condition/)  *  [validators](https://developer.atlassian.com/cloud/jira/platform/modules/workflow-validator/)  Only rules created by the calling Connect app can be updated.  To assist with app migration, this operation can be used to:   *  Disable a rule.  *  Add a `tag`. Use this to filter rules in the [Get workflow transition rule configurations](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-workflow-transition-rules/#api-rest-api-3-workflow-rule-config-get).  Rules are enabled if the `disabled` parameter is not provided.  **[Permissions](#permissions) required:** Only Connect apps can use this operation.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**WorkflowTransitionRulesUpdate**](WorkflowTransitionRulesUpdate.md)|  | 

### Return type

[**WorkflowTransitionRulesUpdateErrors**](WorkflowTransitionRulesUpdateErrors.md)

### Authorization

[basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

