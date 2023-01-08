# IssueBulkDeletePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SelectedIssueIdsOrKeys** | **[]string** | List of issue IDs or keys which are to be bulk deleted. These IDs or keys can be from different projects and issue types. | 
**SendBulkNotification** | Pointer to **NullableBool** | A boolean value that indicates whether to send a bulk change notification when the issues are being deleted.  If &#x60;true&#x60;, dispatches a bulk notification email to users about the updates. | [optional] [default to true]

## Methods

### NewIssueBulkDeletePayload

`func NewIssueBulkDeletePayload(selectedIssueIdsOrKeys []string, ) *IssueBulkDeletePayload`

NewIssueBulkDeletePayload instantiates a new IssueBulkDeletePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkDeletePayloadWithDefaults

`func NewIssueBulkDeletePayloadWithDefaults() *IssueBulkDeletePayload`

NewIssueBulkDeletePayloadWithDefaults instantiates a new IssueBulkDeletePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSelectedIssueIdsOrKeys

`func (o *IssueBulkDeletePayload) GetSelectedIssueIdsOrKeys() []string`

GetSelectedIssueIdsOrKeys returns the SelectedIssueIdsOrKeys field if non-nil, zero value otherwise.

### GetSelectedIssueIdsOrKeysOk

`func (o *IssueBulkDeletePayload) GetSelectedIssueIdsOrKeysOk() (*[]string, bool)`

GetSelectedIssueIdsOrKeysOk returns a tuple with the SelectedIssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedIssueIdsOrKeys

`func (o *IssueBulkDeletePayload) SetSelectedIssueIdsOrKeys(v []string)`

SetSelectedIssueIdsOrKeys sets SelectedIssueIdsOrKeys field to given value.


### GetSendBulkNotification

`func (o *IssueBulkDeletePayload) GetSendBulkNotification() bool`

GetSendBulkNotification returns the SendBulkNotification field if non-nil, zero value otherwise.

### GetSendBulkNotificationOk

`func (o *IssueBulkDeletePayload) GetSendBulkNotificationOk() (*bool, bool)`

GetSendBulkNotificationOk returns a tuple with the SendBulkNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBulkNotification

`func (o *IssueBulkDeletePayload) SetSendBulkNotification(v bool)`

SetSendBulkNotification sets SendBulkNotification field to given value.

### HasSendBulkNotification

`func (o *IssueBulkDeletePayload) HasSendBulkNotification() bool`

HasSendBulkNotification returns a boolean if a field has been set.

### SetSendBulkNotificationNil

`func (o *IssueBulkDeletePayload) SetSendBulkNotificationNil(b bool)`

 SetSendBulkNotificationNil sets the value for SendBulkNotification to be an explicit nil

### UnsetSendBulkNotification
`func (o *IssueBulkDeletePayload) UnsetSendBulkNotification()`

UnsetSendBulkNotification ensures that no value is present for SendBulkNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


