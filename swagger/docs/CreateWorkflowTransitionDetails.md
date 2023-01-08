# CreateWorkflowTransitionDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | **string** | The description of the transition. The maximum length is 1000 characters. | [optional] [default to null]
**From** | **[]string** | The statuses the transition can start from. | [optional] [default to null]
**Name** | **string** | The name of the transition. The maximum length is 60 characters. | [default to null]
**Properties** | **map[string]string** | The properties of the transition. | [optional] [default to null]
**Rules** | [***AllOfCreateWorkflowTransitionDetailsRules**](AllOfCreateWorkflowTransitionDetailsRules.md) | The rules of the transition. | [optional] [default to null]
**Screen** | [***AllOfCreateWorkflowTransitionDetailsScreen**](AllOfCreateWorkflowTransitionDetailsScreen.md) | The screen of the transition. | [optional] [default to null]
**To** | **string** | The status the transition goes to. | [default to null]
**Type_** | **string** | The type of the transition. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

