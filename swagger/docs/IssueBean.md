# IssueBean

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Changelog** | [***AllOfIssueBeanChangelog**](AllOfIssueBeanChangelog.md) | Details of changelogs associated with the issue. | [optional] [default to null]
**Editmeta** | [***AllOfIssueBeanEditmeta**](AllOfIssueBeanEditmeta.md) | The metadata for the fields on the issue that can be amended. | [optional] [default to null]
**Expand** | **string** | Expand options that include additional issue details in the response. | [optional] [default to null]
**Fields** | [**map[string]Object**](.md) |  | [optional] [default to null]
**FieldsToInclude** | [***IncludedFields**](IncludedFields.md) |  | [optional] [default to null]
**Id** | **string** | The ID of the issue. | [optional] [default to null]
**Key** | **string** | The key of the issue. | [optional] [default to null]
**Names** | **map[string]string** | The ID and name of each field present on the issue. | [optional] [default to null]
**Operations** | [***AllOfIssueBeanOperations**](AllOfIssueBeanOperations.md) | The operations that can be performed on the issue. | [optional] [default to null]
**Properties** | [**map[string]Object**](.md) | Details of the issue properties identified in the request. | [optional] [default to null]
**RenderedFields** | [**map[string]Object**](.md) | The rendered value of each field present on the issue. | [optional] [default to null]
**Schema** | [**map[string]JsonTypeBean**](JsonTypeBean.md) | The schema describing each field present on the issue. | [optional] [default to null]
**Self** | **string** | The URL of the issue details. | [optional] [default to null]
**Transitions** | [**[]ModelMap**](map.md) | The transitions that can be performed on the issue. | [optional] [default to null]
**VersionedRepresentations** | [**map[string]map[string]Object**](map.md) | The versions of each field on the issue. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

