# {{classname}}

All URIs are relative to *https://your-domain.atlassian.net*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflowTransitionProperty**](WorkflowTransitionPropertiesApi.md#CreateWorkflowTransitionProperty) | **Post** /rest/api/3/workflow/transitions/{transitionId}/properties | Create workflow transition property
[**DeleteWorkflowTransitionProperty**](WorkflowTransitionPropertiesApi.md#DeleteWorkflowTransitionProperty) | **Delete** /rest/api/3/workflow/transitions/{transitionId}/properties | Delete workflow transition property
[**GetWorkflowTransitionProperties**](WorkflowTransitionPropertiesApi.md#GetWorkflowTransitionProperties) | **Get** /rest/api/3/workflow/transitions/{transitionId}/properties | Get workflow transition properties
[**UpdateWorkflowTransitionProperty**](WorkflowTransitionPropertiesApi.md#UpdateWorkflowTransitionProperty) | **Put** /rest/api/3/workflow/transitions/{transitionId}/properties | Update workflow transition property

# **CreateWorkflowTransitionProperty**
> ModelMap CreateWorkflowTransitionProperty(ctx, body, transitionId, key, workflowName, optional)
Create workflow transition property

Adds a property to a workflow transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/zIhKLg#Advancedworkflowconfiguration-transitionproperties) and [Workflow properties](https://confluence.atlassian.com/x/JYlKLg).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
  **transitionId** | **int64**| The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 
  **key** | **string**| The key of the property being added, also known as the name of the property. Set this to the same value as the &#x60;key&#x60; defined in the request body. | 
  **workflowName** | **string**| The name of the workflow that the transition belongs to. | 
 **optional** | ***WorkflowTransitionPropertiesApiCreateWorkflowTransitionPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowTransitionPropertiesApiCreateWorkflowTransitionPropertyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **workflowMode** | **optional.**| The workflow status. Set to *live* for inactive workflows or *draft* for draft workflows. Active workflows cannot be edited. | [default to live]

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteWorkflowTransitionProperty**
> DeleteWorkflowTransitionProperty(ctx, transitionId, key, workflowName, optional)
Delete workflow transition property

Deletes a property from a workflow transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/zIhKLg#Advancedworkflowconfiguration-transitionproperties) and [Workflow properties](https://confluence.atlassian.com/x/JYlKLg).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transitionId** | **int64**| The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 
  **key** | **string**| The name of the transition property to delete, also known as the name of the property. | 
  **workflowName** | **string**| The name of the workflow that the transition belongs to. | 
 **optional** | ***WorkflowTransitionPropertiesApiDeleteWorkflowTransitionPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowTransitionPropertiesApiDeleteWorkflowTransitionPropertyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **workflowMode** | **optional.String**| The workflow status. Set to &#x60;live&#x60; for inactive workflows or &#x60;draft&#x60; for draft workflows. Active workflows cannot be edited. | 

### Return type

 (empty response body)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetWorkflowTransitionProperties**
> ModelMap GetWorkflowTransitionProperties(ctx, transitionId, workflowName, optional)
Get workflow transition properties

Returns the properties on a workflow transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/zIhKLg#Advancedworkflowconfiguration-transitionproperties) and [Workflow properties](https://confluence.atlassian.com/x/JYlKLg).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **transitionId** | **int64**| The ID of the transition. To get the ID, view the workflow in text mode in the Jira administration console. The ID is shown next to the transition. | 
  **workflowName** | **string**| The name of the workflow that the transition belongs to. | 
 **optional** | ***WorkflowTransitionPropertiesApiGetWorkflowTransitionPropertiesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowTransitionPropertiesApiGetWorkflowTransitionPropertiesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **includeReservedKeys** | **optional.Bool**| Some properties with keys that have the *jira.* prefix are reserved, which means they are not editable. To include these properties in the results, set this parameter to *true*. | [default to false]
 **key** | **optional.String**| The key of the property being returned, also known as the name of the property. If this parameter is not specified, all properties on the transition are returned. | 
 **workflowMode** | **optional.String**| The workflow status. Set to *live* for active and inactive workflows, or *draft* for draft workflows. | [default to live]

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateWorkflowTransitionProperty**
> ModelMap UpdateWorkflowTransitionProperty(ctx, body, transitionId, key, workflowName, optional)
Update workflow transition property

Updates a workflow transition by changing the property value. Trying to update a property that does not exist results in a new property being added to the transition. Transition properties are used to change the behavior of a transition. For more information, see [Transition properties](https://confluence.atlassian.com/x/zIhKLg#Advancedworkflowconfiguration-transitionproperties) and [Workflow properties](https://confluence.atlassian.com/x/JYlKLg).  **[Permissions](#permissions) required:** *Administer Jira* [global permission](https://confluence.atlassian.com/x/x4dKLg).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ModelMap**](map.md)|  | 
  **transitionId** | **int64**| The ID of the transition. To get the ID, view the workflow in text mode in the Jira admin settings. The ID is shown next to the transition. | 
  **key** | **string**| The key of the property being updated, also known as the name of the property. Set this to the same value as the &#x60;key&#x60; defined in the request body. | 
  **workflowName** | **string**| The name of the workflow that the transition belongs to. | 
 **optional** | ***WorkflowTransitionPropertiesApiUpdateWorkflowTransitionPropertyOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a WorkflowTransitionPropertiesApiUpdateWorkflowTransitionPropertyOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **workflowMode** | **optional.**| The workflow status. Set to &#x60;live&#x60; for inactive workflows or &#x60;draft&#x60; for draft workflows. Active workflows cannot be edited. | 

### Return type

[**ModelMap**](map.md)

### Authorization

[OAuth2](../README.md#OAuth2), [basicAuth](../README.md#basicAuth)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

