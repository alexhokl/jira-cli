# ProjectFeature

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Feature** | **string** | The key of the feature. | [optional] [default to null]
**ImageUri** | **string** | URI for the image representing the feature. | [optional] [default to null]
**LocalisedDescription** | **string** | Localized display description for the feature. | [optional] [default to null]
**LocalisedName** | **string** | Localized display name for the feature. | [optional] [default to null]
**Prerequisites** | **[]string** | List of keys of the features required to enable the feature. | [optional] [default to null]
**ProjectId** | **int64** | The ID of the project. | [optional] [default to null]
**State** | **string** | The state of the feature. When updating the state of a feature, only ENABLED and DISABLED are supported. Responses can contain all values | [optional] [default to null]
**ToggleLocked** | **bool** | Whether the state of the feature can be updated. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

