# WorkflowTransitionRulesUpdateErrorDetails

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleUpdateErrors** | [**map[string][]string**](array.md) | A list of transition rule update errors, indexed by the transition rule ID. Any transition rule that appears here wasn&#x27;t updated. | [default to null]
**UpdateErrors** | **[]string** | The list of errors that specify why the workflow update failed. The workflow was not updated if the list contains any entries. | [default to null]
**WorkflowId** | [***WorkflowId**](WorkflowId.md) |  | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

