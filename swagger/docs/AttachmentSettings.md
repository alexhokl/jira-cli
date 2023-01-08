# AttachmentSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | Whether the ability to add attachments is enabled. | [optional] [readonly] 
**UploadLimit** | Pointer to **int64** | The maximum size of attachments permitted, in bytes. | [optional] [readonly] 

## Methods

### NewAttachmentSettings

`func NewAttachmentSettings() *AttachmentSettings`

NewAttachmentSettings instantiates a new AttachmentSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAttachmentSettingsWithDefaults

`func NewAttachmentSettingsWithDefaults() *AttachmentSettings`

NewAttachmentSettingsWithDefaults instantiates a new AttachmentSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *AttachmentSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AttachmentSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AttachmentSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AttachmentSettings) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetUploadLimit

`func (o *AttachmentSettings) GetUploadLimit() int64`

GetUploadLimit returns the UploadLimit field if non-nil, zero value otherwise.

### GetUploadLimitOk

`func (o *AttachmentSettings) GetUploadLimitOk() (*int64, bool)`

GetUploadLimitOk returns a tuple with the UploadLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploadLimit

`func (o *AttachmentSettings) SetUploadLimit(v int64)`

SetUploadLimit sets UploadLimit field to given value.

### HasUploadLimit

`func (o *AttachmentSettings) HasUploadLimit() bool`

HasUploadLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


