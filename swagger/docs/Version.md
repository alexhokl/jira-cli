# Version

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Archived** | **bool** | Indicates that the version is archived. Optional when creating or updating a version. | [optional] [default to null]
**Description** | **string** | The description of the version. Optional when creating or updating a version. | [optional] [default to null]
**Expand** | **string** | Use [expand](em&gt;#expansion) to include additional information about version in the response. This parameter accepts a comma-separated list. Expand options include:   *  &#x60;operations&#x60; Returns the list of operations available for this version.  *  &#x60;issuesstatus&#x60; Returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*.  Optional for create and update. | [optional] [default to null]
**Id** | **string** | The ID of the version. | [optional] [default to null]
**IssuesStatusForFixVersion** | [***AllOfVersionIssuesStatusForFixVersion**](AllOfVersionIssuesStatusForFixVersion.md) | If the expand option &#x60;issuesstatus&#x60; is used, returns the count of issues in this version for each of the status categories *to do*, *in progress*, *done*, and *unmapped*. The *unmapped* property contains a count of issues with a status other than *to do*, *in progress*, and *done*. | [optional] [default to null]
**MoveUnfixedIssuesTo** | **string** | The URL of the self link to the version to which all unfixed issues are moved when a version is released. Not applicable when creating a version. Optional when updating a version. | [optional] [default to null]
**Name** | **string** | The unique name of the version. Required when creating a version. Optional when updating a version. The maximum length is 255 characters. | [optional] [default to null]
**Operations** | [**[]SimpleLink**](SimpleLink.md) | If the expand option &#x60;operations&#x60; is used, returns the list of operations available for this version. | [optional] [default to null]
**Overdue** | **bool** | Indicates that the version is overdue. | [optional] [default to null]
**Project** | **string** | Deprecated. Use &#x60;projectId&#x60;. | [optional] [default to null]
**ProjectId** | **int64** | The ID of the project to which this version is attached. Required when creating a version. Not applicable when updating a version. | [optional] [default to null]
**ReleaseDate** | **string** | The release date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version. | [optional] [default to null]
**Released** | **bool** | Indicates that the version is released. If the version is released a request to release again is ignored. Not applicable when creating a version. Optional when updating a version. | [optional] [default to null]
**Self** | **string** | The URL of the version. | [optional] [default to null]
**StartDate** | **string** | The start date of the version. Expressed in ISO 8601 format (yyyy-mm-dd). Optional when creating or updating a version. | [optional] [default to null]
**UserReleaseDate** | **string** | The date on which work on this version is expected to finish, expressed in the instance&#x27;s *Day/Month/Year Format* date format. | [optional] [default to null]
**UserStartDate** | **string** | The date on which work on this version is expected to start, expressed in the instance&#x27;s *Day/Month/Year Format* date format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

