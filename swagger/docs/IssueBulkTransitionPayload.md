# IssueBulkTransitionPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkTransitionInputs** | [**[]BulkTransitionSubmitInput**](BulkTransitionSubmitInput.md) | List of objects and each object has two properties:   *  Issues that will be bulk transitioned.  *  TransitionId that corresponds to a specific transition of issues that share the same workflow. | 
**SendBulkNotification** | Pointer to **NullableBool** | A boolean value that indicates whether to send a bulk change notification when the issues are being transitioned.  If &#x60;true&#x60;, dispatches a bulk notification email to users about the updates. | [optional] [default to true]

## Methods

### NewIssueBulkTransitionPayload

`func NewIssueBulkTransitionPayload(bulkTransitionInputs []BulkTransitionSubmitInput, ) *IssueBulkTransitionPayload`

NewIssueBulkTransitionPayload instantiates a new IssueBulkTransitionPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkTransitionPayloadWithDefaults

`func NewIssueBulkTransitionPayloadWithDefaults() *IssueBulkTransitionPayload`

NewIssueBulkTransitionPayloadWithDefaults instantiates a new IssueBulkTransitionPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkTransitionInputs

`func (o *IssueBulkTransitionPayload) GetBulkTransitionInputs() []BulkTransitionSubmitInput`

GetBulkTransitionInputs returns the BulkTransitionInputs field if non-nil, zero value otherwise.

### GetBulkTransitionInputsOk

`func (o *IssueBulkTransitionPayload) GetBulkTransitionInputsOk() (*[]BulkTransitionSubmitInput, bool)`

GetBulkTransitionInputsOk returns a tuple with the BulkTransitionInputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkTransitionInputs

`func (o *IssueBulkTransitionPayload) SetBulkTransitionInputs(v []BulkTransitionSubmitInput)`

SetBulkTransitionInputs sets BulkTransitionInputs field to given value.


### GetSendBulkNotification

`func (o *IssueBulkTransitionPayload) GetSendBulkNotification() bool`

GetSendBulkNotification returns the SendBulkNotification field if non-nil, zero value otherwise.

### GetSendBulkNotificationOk

`func (o *IssueBulkTransitionPayload) GetSendBulkNotificationOk() (*bool, bool)`

GetSendBulkNotificationOk returns a tuple with the SendBulkNotification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendBulkNotification

`func (o *IssueBulkTransitionPayload) SetSendBulkNotification(v bool)`

SetSendBulkNotification sets SendBulkNotification field to given value.

### HasSendBulkNotification

`func (o *IssueBulkTransitionPayload) HasSendBulkNotification() bool`

HasSendBulkNotification returns a boolean if a field has been set.

### SetSendBulkNotificationNil

`func (o *IssueBulkTransitionPayload) SetSendBulkNotificationNil(b bool)`

 SetSendBulkNotificationNil sets the value for SendBulkNotification to be an explicit nil

### UnsetSendBulkNotification
`func (o *IssueBulkTransitionPayload) UnsetSendBulkNotification()`

UnsetSendBulkNotification ensures that no value is present for SendBulkNotification, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


