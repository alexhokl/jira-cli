# IssueBulkEditPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EditedFieldsInput** | [**JiraIssueFields**](JiraIssueFields.md) | An object that defines the values to be updated in specified fields of an issue. The structure and content of this parameter vary depending on the type of field being edited. Although the order is not significant, ensure that field IDs align with those in selectedActions. | 
**SelectedActions** | **[]string** | List of all the field IDs that are to be bulk edited. Each field ID in this list corresponds to a specific attribute of an issue that is set to be modified in the bulk edit operation. The relevant field ID can be obtained by calling the Bulk Edit Get Fields REST API (documentation available on this page itself). | 
**SelectedIssueIdsOrKeys** | **[]string** | List of issue IDs or keys which are to be bulk edited. These IDs or keys can be from different projects and issue types. | 
**SendBulkNotification** | Pointer to **NullableBool** | A boolean value that indicates whether to send a bulk change notification when the issues are being edited.  If &#x60;true&#x60;, dispatches a bulk notification email to users about the updates. | [optional] [default to true]

## Methods

### NewIssueBulkEditPayload

`func NewIssueBulkEditPayload(editedFieldsInput JiraIssueFields, selectedActions []string, selectedIssueIdsOrKeys []string, ) *IssueBulkEditPayload`

NewIssueBulkEditPayload instantiates a new IssueBulkEditPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkEditPayloadWithDefaults

`func NewIssueBulkEditPayloadWithDefaults() *IssueBulkEditPayload`

NewIssueBulkEditPayloadWithDefaults instantiates a new IssueBulkEditPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEditedFieldsInput

`func (o *IssueBulkEditPayload) GetEditedFieldsInput() JiraIssueFields`

GetEditedFieldsInput returns the EditedFieldsInput field if non-nil, zero value otherwise.

### GetEditedFieldsInputOk

`func (o *IssueBulkEditPayload) GetEditedFieldsInputOk() (*JiraIssueFields, bool)`

GetEditedFieldsInputOk returns a tuple with the EditedFieldsInput field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditedFieldsInput

`func (o *IssueBulkEditPayload) SetEditedFieldsInput(v JiraIssueFields)`

SetEditedFieldsInput sets EditedFieldsInput field to given value.


### GetSelectedActions

`func (o *IssueBulkEditPayload) GetSelectedActions() []string`

GetSelectedActions returns the SelectedActions field if non-nil, zero value otherwise.

### GetSelectedActionsOk

`func (o *IssueBulkEditPayload) GetSelectedActionsOk() (*[]string, bool)`

GetSelectedActionsOk returns a tuple with the SelectedActions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedActions

`func (o *IssueBulkEditPayload) SetSelectedActions(v []string)`

SetSelectedActions sets SelectedActions field to given value.


### GetSelectedIssueIdsOrKeys

`func (o *IssueBulkEditPayload) GetSelectedIssueIdsOrKeys() []string`

GetSelectedIssueIdsOrKeys returns the SelectedIssueIdsOrKeys field if non-nil, zero value otherwise.

### GetSelectedIssueIdsOrKeysOk

`func (o *IssueBulkEditPayload) GetSelectedIssueIdsOrKeysOk() (*[]string, bool)`

GetSelectedIssueIdsOrKeysOk returns a tuple with the SelectedIssueIdsOrKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedIssueIdsOrKeys

`func (o *IssueBulkEditPayload) SetSelectedIssueIdsOrKeys(v []string)`

SetSelectedIssueIdsOrKeys sets SelectedIssueIdsOrKeys field to given value.


### GetSendBulkNotification

`func (o *IssueBulkEditPayload) GetSendBulkNotification() bool`

GetSendBulkNotification returns the SendBulkNotification field if non-nil, zero value otherwise.

### GetSendBulkNotificationOk

`func (o *IssueBulkEditPayload) GetSendBulkNotificationOk() (*bool, bool)`

GetSendBulkNotificationOk returns a tuple with the SendBulkNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBulkNotification

`func (o *IssueBulkEditPayload) SetSendBulkNotification(v bool)`

SetSendBulkNotification sets SendBulkNotification field to given value.

### HasSendBulkNotification

`func (o *IssueBulkEditPayload) HasSendBulkNotification() bool`

HasSendBulkNotification returns a boolean if a field has been set.

### SetSendBulkNotificationNil

`func (o *IssueBulkEditPayload) SetSendBulkNotificationNil(b bool)`

 SetSendBulkNotificationNil sets the value for SendBulkNotification to be an explicit nil

### UnsetSendBulkNotification
`func (o *IssueBulkEditPayload) UnsetSendBulkNotification()`

UnsetSendBulkNotification ensures that no value is present for SendBulkNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


