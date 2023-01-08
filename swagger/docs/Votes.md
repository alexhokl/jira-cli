# Votes

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasVoted** | **bool** | Whether the user making this request has voted on the issue. | [optional] [default to null]
**Self** | **string** | The URL of these issue vote details. | [optional] [default to null]
**Voters** | [**[]User**](User.md) | List of the users who have voted on this issue. An empty list is returned when the calling user doesn&#x27;t have the *View voters and watchers* project permission. | [optional] [default to null]
**Votes** | **int64** | The number of votes on the issue. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

