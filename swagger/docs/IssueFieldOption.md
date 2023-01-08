# IssueFieldOption

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [***IssueFieldOptionConfiguration**](IssueFieldOptionConfiguration.md) |  | [optional] [default to null]
**Id** | **int64** | The unique identifier for the option. This is only unique within the select field&#x27;s set of options. | [default to null]
**Properties** | [**map[string]Object**](.md) | The properties of the object, as arbitrary key-value pairs. These properties can be searched using JQL, if the extractions (see [Issue Field Option Property Index](https://developer.atlassian.com/cloud/jira/platform/modules/issue-field-option-property-index/)) are defined in the descriptor for the issue field module. | [optional] [default to null]
**Value** | **string** | The option&#x27;s name, which is displayed in Jira. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

