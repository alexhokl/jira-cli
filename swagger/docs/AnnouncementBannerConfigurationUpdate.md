# AnnouncementBannerConfigurationUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsDismissible** | Pointer to **bool** | Flag indicating if the announcement banner can be dismissed by the user. | [optional] 
**IsEnabled** | Pointer to **bool** | Flag indicating if the announcement banner is enabled or not. | [optional] 
**Message** | Pointer to **string** | The text on the announcement banner. | [optional] 
**Visibility** | Pointer to **string** | Visibility of the announcement banner. Can be public or private. | [optional] 

## Methods

### NewAnnouncementBannerConfigurationUpdate

`func NewAnnouncementBannerConfigurationUpdate() *AnnouncementBannerConfigurationUpdate`

NewAnnouncementBannerConfigurationUpdate instantiates a new AnnouncementBannerConfigurationUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementBannerConfigurationUpdateWithDefaults

`func NewAnnouncementBannerConfigurationUpdateWithDefaults() *AnnouncementBannerConfigurationUpdate`

NewAnnouncementBannerConfigurationUpdateWithDefaults instantiates a new AnnouncementBannerConfigurationUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsDismissible

`func (o *AnnouncementBannerConfigurationUpdate) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *AnnouncementBannerConfigurationUpdate) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *AnnouncementBannerConfigurationUpdate) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.

### HasIsDismissible

`func (o *AnnouncementBannerConfigurationUpdate) HasIsDismissible() bool`

HasIsDismissible returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AnnouncementBannerConfigurationUpdate) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AnnouncementBannerConfigurationUpdate) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AnnouncementBannerConfigurationUpdate) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AnnouncementBannerConfigurationUpdate) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMessage

`func (o *AnnouncementBannerConfigurationUpdate) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnnouncementBannerConfigurationUpdate) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnnouncementBannerConfigurationUpdate) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AnnouncementBannerConfigurationUpdate) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *AnnouncementBannerConfigurationUpdate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AnnouncementBannerConfigurationUpdate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AnnouncementBannerConfigurationUpdate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AnnouncementBannerConfigurationUpdate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


