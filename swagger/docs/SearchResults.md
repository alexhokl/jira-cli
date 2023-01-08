# SearchResults

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expand** | **string** | Expand options that include additional search result details in the response. | [optional] [default to null]
**Issues** | [**[]IssueBean**](IssueBean.md) | The list of issues found by the search. | [optional] [default to null]
**MaxResults** | **int32** | The maximum number of results that could be on the page. | [optional] [default to null]
**Names** | **map[string]string** | The ID and name of each field in the search results. | [optional] [default to null]
**Schema** | [**map[string]JsonTypeBean**](JsonTypeBean.md) | The schema describing the field types in the search results. | [optional] [default to null]
**StartAt** | **int32** | The index of the first item returned on the page. | [optional] [default to null]
**Total** | **int32** | The number of results on the page. | [optional] [default to null]
**WarningMessages** | **[]string** | Any warnings related to the JQL query. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

