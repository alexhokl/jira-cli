# IssueTypeIssueCreateMetadata

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | **int64** | The ID of the issue type&#x27;s avatar. | [optional] [default to null]
**Description** | **string** | The description of the issue type. | [optional] [default to null]
**EntityId** | **string** | Unique ID for next-gen projects. | [optional] [default to null]
**Expand** | **string** | Expand options that include additional issue type metadata details in the response. | [optional] [default to null]
**Fields** | [**map[string]FieldMetadata**](FieldMetadata.md) | List of the fields available when creating an issue for the issue type. | [optional] [default to null]
**HierarchyLevel** | **int32** | Hierarchy level of the issue type. | [optional] [default to null]
**IconUrl** | **string** | The URL of the issue type&#x27;s avatar. | [optional] [default to null]
**Id** | **string** | The ID of the issue type. | [optional] [default to null]
**Name** | **string** | The name of the issue type. | [optional] [default to null]
**Scope** | [***AllOfIssueTypeIssueCreateMetadataScope**](AllOfIssueTypeIssueCreateMetadataScope.md) | Details of the next-gen projects the issue type is available in. | [optional] [default to null]
**Self** | **string** | The URL of these issue type details. | [optional] [default to null]
**Subtask** | **bool** | Whether this issue type is used to create subtasks. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

