# AuditRecordBean

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociatedItems** | [**[]AssociatedItemBean**](AssociatedItemBean.md) | The list of items associated with the changed record. | [optional] [default to null]
**AuthorKey** | **string** | Deprecated, use &#x60;authorAccountId&#x60; instead. The key of the user who created the audit record. | [optional] [default to null]
**Category** | **string** | The category of the audit record. For a list of these categories, see the help article [Auditing in Jira applications](https://confluence.atlassian.com/x/noXKM). | [optional] [default to null]
**ChangedValues** | [**[]ChangedValueBean**](ChangedValueBean.md) | The list of values changed in the record event. | [optional] [default to null]
**Created** | [**time.Time**](time.Time.md) | The date and time on which the audit record was created. | [optional] [default to null]
**Description** | **string** | The description of the audit record. | [optional] [default to null]
**EventSource** | **string** | The event the audit record originated from. | [optional] [default to null]
**Id** | **int64** | The ID of the audit record. | [optional] [default to null]
**ObjectItem** | [***AssociatedItemBean**](AssociatedItemBean.md) |  | [optional] [default to null]
**RemoteAddress** | **string** | The URL of the computer where the creation of the audit record was initiated. | [optional] [default to null]
**Summary** | **string** | The summary of the audit record. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

