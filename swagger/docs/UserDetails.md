# UserDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**AccountType** | Pointer to **string** | The type of account represented by this user. This will be one of &#39;atlassian&#39; (normal users), &#39;app&#39; (application user) or &#39;customer&#39; (Jira Service Desk customer user) | [optional] [readonly] 
**Active** | Pointer to **bool** | Whether the user is active. | [optional] [readonly] 
**AvatarUrls** | Pointer to [**AvatarUrlsBean**](AvatarUrlsBean.md) | The avatars of the user. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the user. Depending on the user’s privacy settings, this may return an alternative value. | [optional] [readonly] 
**EmailAddress** | Pointer to **string** | The email address of the user. Depending on the user’s privacy settings, this may be returned as null. | [optional] [readonly] 
**Key** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [readonly] 
**Name** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the user. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | The time zone specified in the user&#39;s profile. Depending on the user’s privacy settings, this may be returned as null. | [optional] [readonly] 

## Methods

### NewUserDetails

`func NewUserDetails() *UserDetails`

NewUserDetails instantiates a new UserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserDetailsWithDefaults

`func NewUserDetailsWithDefaults() *UserDetails`

NewUserDetailsWithDefaults instantiates a new UserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserDetails) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserDetails) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserDetails) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserDetails) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *UserDetails) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UserDetails) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UserDetails) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UserDetails) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetActive

`func (o *UserDetails) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserDetails) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserDetails) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserDetails) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvatarUrls

`func (o *UserDetails) GetAvatarUrls() AvatarUrlsBean`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *UserDetails) GetAvatarUrlsOk() (*AvatarUrlsBean, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *UserDetails) SetAvatarUrls(v AvatarUrlsBean)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *UserDetails) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserDetails) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *UserDetails) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserDetails) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserDetails) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserDetails) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetKey

`func (o *UserDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UserDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *UserDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UserDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UserDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *UserDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTimeZone

`func (o *UserDetails) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *UserDetails) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *UserDetails) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *UserDetails) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


