# Workflow

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | [**time.Time**](time.Time.md) | The creation date of the workflow. | [optional] [default to null]
**Description** | **string** | The description of the workflow. | [default to null]
**HasDraftWorkflow** | **bool** | Whether the workflow has a draft version. | [optional] [default to null]
**Id** | [***PublishedWorkflowId**](PublishedWorkflowId.md) |  | [default to null]
**IsDefault** | **bool** | Whether this is the default workflow. | [optional] [default to null]
**Operations** | [***WorkflowOperations**](WorkflowOperations.md) |  | [optional] [default to null]
**Projects** | [**[]ProjectDetails**](ProjectDetails.md) | The projects the workflow is assigned to, through workflow schemes. | [optional] [default to null]
**Schemes** | [**[]WorkflowSchemeIdName**](WorkflowSchemeIdName.md) | The workflow schemes the workflow is assigned to. | [optional] [default to null]
**Statuses** | [**[]WorkflowStatus**](WorkflowStatus.md) | The statuses of the workflow. | [optional] [default to null]
**Transitions** | [**[]Transition**](Transition.md) | The transitions of the workflow. | [optional] [default to null]
**Updated** | [**time.Time**](time.Time.md) | The last edited date of the workflow. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

