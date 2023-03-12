# DeleteAndReplaceVersionBean

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldReplacementList** | [**[]CustomFieldReplacement**](CustomFieldReplacement.md) | An array of custom field IDs (&#x60;customFieldId&#x60;) and version IDs (&#x60;moveTo&#x60;) to update when the fields contain the deleted version. | [optional] [default to null]
**MoveAffectedIssuesTo** | **int64** | The ID of the version to update &#x60;affectedVersion&#x60; to when the field contains the deleted version. | [optional] [default to null]
**MoveFixIssuesTo** | **int64** | The ID of the version to update &#x60;fixVersion&#x60; to when the field contains the deleted version. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

