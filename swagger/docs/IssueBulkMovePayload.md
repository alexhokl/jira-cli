# IssueBulkMovePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SendBulkNotification** | Pointer to **NullableBool** | A boolean value that indicates whether to send a bulk change notification when the issues are being moved.  If &#x60;true&#x60;, dispatches a bulk notification email to users about the updates. | [optional] [default to true]
**TargetToSourcesMapping** | Pointer to [**map[string]TargetToSourcesMapping**](TargetToSourcesMapping.md) | An object representing the mapping of issues and data related to destination entities, like fields and statuses, that are required during a bulk move.  The key is a string that is created by concatenating the following three entities in order, separated by commas. The format is &#x60;&lt;project ID or key&gt;,&lt;issueType ID&gt;,&lt;parent ID or key&gt;&#x60;. It should be unique across mappings provided in the payload. If you provide multiple mappings for the same key, only one will be processed. However, the operation won&#39;t fail, so the error may be hard to track down.   *  ***Destination project*** (Required): ID or key of the project to which the issues are being moved.  *  ***Destination issueType*** (Required): ID of the issueType to which the issues are being moved.  *  ***Destination parent ID or key*** (Optional): ID or key of the issue which will become the parent of the issues being moved. Only required when the destination issueType is a subtask. | [optional] 

## Methods

### NewIssueBulkMovePayload

`func NewIssueBulkMovePayload() *IssueBulkMovePayload`

NewIssueBulkMovePayload instantiates a new IssueBulkMovePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkMovePayloadWithDefaults

`func NewIssueBulkMovePayloadWithDefaults() *IssueBulkMovePayload`

NewIssueBulkMovePayloadWithDefaults instantiates a new IssueBulkMovePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSendBulkNotification

`func (o *IssueBulkMovePayload) GetSendBulkNotification() bool`

GetSendBulkNotification returns the SendBulkNotification field if non-nil, zero value otherwise.

### GetSendBulkNotificationOk

`func (o *IssueBulkMovePayload) GetSendBulkNotificationOk() (*bool, bool)`

GetSendBulkNotificationOk returns a tuple with the SendBulkNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBulkNotification

`func (o *IssueBulkMovePayload) SetSendBulkNotification(v bool)`

SetSendBulkNotification sets SendBulkNotification field to given value.

### HasSendBulkNotification

`func (o *IssueBulkMovePayload) HasSendBulkNotification() bool`

HasSendBulkNotification returns a boolean if a field has been set.

### SetSendBulkNotificationNil

`func (o *IssueBulkMovePayload) SetSendBulkNotificationNil(b bool)`

 SetSendBulkNotificationNil sets the value for SendBulkNotification to be an explicit nil

### UnsetSendBulkNotification
`func (o *IssueBulkMovePayload) UnsetSendBulkNotification()`

UnsetSendBulkNotification ensures that no value is present for SendBulkNotification, not even an explicit nil
### GetTargetToSourcesMapping

`func (o *IssueBulkMovePayload) GetTargetToSourcesMapping() map[string]TargetToSourcesMapping`

GetTargetToSourcesMapping returns the TargetToSourcesMapping field if non-nil, zero value otherwise.

### GetTargetToSourcesMappingOk

`func (o *IssueBulkMovePayload) GetTargetToSourcesMappingOk() (*map[string]TargetToSourcesMapping, bool)`

GetTargetToSourcesMappingOk returns a tuple with the TargetToSourcesMapping field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetToSourcesMapping

`func (o *IssueBulkMovePayload) SetTargetToSourcesMapping(v map[string]TargetToSourcesMapping)`

SetTargetToSourcesMapping sets TargetToSourcesMapping field to given value.

### HasTargetToSourcesMapping

`func (o *IssueBulkMovePayload) HasTargetToSourcesMapping() bool`

HasTargetToSourcesMapping returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


