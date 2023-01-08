# AnnouncementBannerConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HashId** | Pointer to **string** | Hash of the banner data. The client detects updates by comparing hash IDs. | [optional] [readonly] 
**IsDismissible** | Pointer to **bool** | Flag indicating if the announcement banner can be dismissed by the user. | [optional] [readonly] 
**IsEnabled** | Pointer to **bool** | Flag indicating if the announcement banner is enabled or not. | [optional] [readonly] 
**Message** | Pointer to **string** | The text on the announcement banner. | [optional] [readonly] 
**Visibility** | Pointer to **string** | Visibility of the announcement banner. | [optional] [readonly] 

## Methods

### NewAnnouncementBannerConfiguration

`func NewAnnouncementBannerConfiguration() *AnnouncementBannerConfiguration`

NewAnnouncementBannerConfiguration instantiates a new AnnouncementBannerConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAnnouncementBannerConfigurationWithDefaults

`func NewAnnouncementBannerConfigurationWithDefaults() *AnnouncementBannerConfiguration`

NewAnnouncementBannerConfigurationWithDefaults instantiates a new AnnouncementBannerConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHashId

`func (o *AnnouncementBannerConfiguration) GetHashId() string`

GetHashId returns the HashId field if non-nil, zero value otherwise.

### GetHashIdOk

`func (o *AnnouncementBannerConfiguration) GetHashIdOk() (*string, bool)`

GetHashIdOk returns a tuple with the HashId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashId

`func (o *AnnouncementBannerConfiguration) SetHashId(v string)`

SetHashId sets HashId field to given value.

### HasHashId

`func (o *AnnouncementBannerConfiguration) HasHashId() bool`

HasHashId returns a boolean if a field has been set.

### GetIsDismissible

`func (o *AnnouncementBannerConfiguration) GetIsDismissible() bool`

GetIsDismissible returns the IsDismissible field if non-nil, zero value otherwise.

### GetIsDismissibleOk

`func (o *AnnouncementBannerConfiguration) GetIsDismissibleOk() (*bool, bool)`

GetIsDismissibleOk returns a tuple with the IsDismissible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDismissible

`func (o *AnnouncementBannerConfiguration) SetIsDismissible(v bool)`

SetIsDismissible sets IsDismissible field to given value.

### HasIsDismissible

`func (o *AnnouncementBannerConfiguration) HasIsDismissible() bool`

HasIsDismissible returns a boolean if a field has been set.

### GetIsEnabled

`func (o *AnnouncementBannerConfiguration) GetIsEnabled() bool`

GetIsEnabled returns the IsEnabled field if non-nil, zero value otherwise.

### GetIsEnabledOk

`func (o *AnnouncementBannerConfiguration) GetIsEnabledOk() (*bool, bool)`

GetIsEnabledOk returns a tuple with the IsEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEnabled

`func (o *AnnouncementBannerConfiguration) SetIsEnabled(v bool)`

SetIsEnabled sets IsEnabled field to given value.

### HasIsEnabled

`func (o *AnnouncementBannerConfiguration) HasIsEnabled() bool`

HasIsEnabled returns a boolean if a field has been set.

### GetMessage

`func (o *AnnouncementBannerConfiguration) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AnnouncementBannerConfiguration) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AnnouncementBannerConfiguration) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AnnouncementBannerConfiguration) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetVisibility

`func (o *AnnouncementBannerConfiguration) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AnnouncementBannerConfiguration) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AnnouncementBannerConfiguration) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AnnouncementBannerConfiguration) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


