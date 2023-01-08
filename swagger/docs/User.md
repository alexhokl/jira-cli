# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. Required in requests. | [optional] 
**AccountType** | Pointer to **string** | The user account type. Can take the following values:   *  &#x60;atlassian&#x60; regular Atlassian user account  *  &#x60;app&#x60; system account used for Connect applications and OAuth to represent external systems  *  &#x60;customer&#x60; Jira Service Desk account representing an external service desk | [optional] [readonly] 
**Active** | Pointer to **bool** | Whether the user is active. | [optional] [readonly] 
**AppType** | Pointer to **string** | The app type of the user account when accountType is &#39;app&#39;. Can take the following values:   *  &#x60;service&#x60; Service Account  *  &#x60;agent&#x60; Rovo Agent Account  *  &#x60;unknown&#x60; Unknown app type | [optional] [readonly] 
**ApplicationRoles** | Pointer to [**SimpleListWrapperApplicationRole**](SimpleListWrapperApplicationRole.md) | The application roles the user is assigned to. | [optional] [readonly] 
**AvatarUrls** | Pointer to [**AvatarUrlsBean**](AvatarUrlsBean.md) | The avatars of the user. | [optional] [readonly] 
**DisplayName** | Pointer to **string** | The display name of the user. Depending on the user’s privacy setting, this may return an alternative value. | [optional] [readonly] 
**EmailAddress** | Pointer to **string** | The email address of the user. Depending on the user’s privacy setting, this may be returned as null. | [optional] [readonly] 
**Expand** | Pointer to **string** | Expand options that include additional user details in the response. | [optional] [readonly] 
**Groups** | Pointer to [**SimpleListWrapperGroupName**](SimpleListWrapperGroupName.md) | The groups that the user belongs to. | [optional] [readonly] 
**Key** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Locale** | Pointer to **string** | The locale of the user. Depending on the user’s privacy setting, this may be returned as null. | [optional] [readonly] 
**Name** | Pointer to **string** | This property is no longer available and will be removed from the documentation soon. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Self** | Pointer to **string** | The URL of the user. | [optional] [readonly] 
**TimeZone** | Pointer to **string** | The time zone specified in the user&#39;s profile. If the user&#39;s time zone is not visible to the current user (due to user&#39;s profile setting), or if a time zone has not been set, the instance&#39;s default time zone will be returned. | [optional] [readonly] 

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *User) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *User) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *User) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *User) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *User) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *User) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *User) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *User) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetActive

`func (o *User) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *User) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *User) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *User) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAppType

`func (o *User) GetAppType() string`

GetAppType returns the AppType field if non-nil, zero value otherwise.

### GetAppTypeOk

`func (o *User) GetAppTypeOk() (*string, bool)`

GetAppTypeOk returns a tuple with the AppType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppType

`func (o *User) SetAppType(v string)`

SetAppType sets AppType field to given value.

### HasAppType

`func (o *User) HasAppType() bool`

HasAppType returns a boolean if a field has been set.

### GetApplicationRoles

`func (o *User) GetApplicationRoles() SimpleListWrapperApplicationRole`

GetApplicationRoles returns the ApplicationRoles field if non-nil, zero value otherwise.

### GetApplicationRolesOk

`func (o *User) GetApplicationRolesOk() (*SimpleListWrapperApplicationRole, bool)`

GetApplicationRolesOk returns a tuple with the ApplicationRoles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationRoles

`func (o *User) SetApplicationRoles(v SimpleListWrapperApplicationRole)`

SetApplicationRoles sets ApplicationRoles field to given value.

### HasApplicationRoles

`func (o *User) HasApplicationRoles() bool`

HasApplicationRoles returns a boolean if a field has been set.

### GetAvatarUrls

`func (o *User) GetAvatarUrls() AvatarUrlsBean`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *User) GetAvatarUrlsOk() (*AvatarUrlsBean, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *User) SetAvatarUrls(v AvatarUrlsBean)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *User) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetDisplayName

`func (o *User) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *User) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *User) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *User) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *User) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *User) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *User) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *User) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetExpand

`func (o *User) GetExpand() string`

GetExpand returns the Expand field if non-nil, zero value otherwise.

### GetExpandOk

`func (o *User) GetExpandOk() (*string, bool)`

GetExpandOk returns a tuple with the Expand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpand

`func (o *User) SetExpand(v string)`

SetExpand sets Expand field to given value.

### HasExpand

`func (o *User) HasExpand() bool`

HasExpand returns a boolean if a field has been set.

### GetGroups

`func (o *User) GetGroups() SimpleListWrapperGroupName`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *User) GetGroupsOk() (*SimpleListWrapperGroupName, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *User) SetGroups(v SimpleListWrapperGroupName)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *User) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetKey

`func (o *User) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *User) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *User) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *User) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLocale

`func (o *User) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *User) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *User) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *User) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetName

`func (o *User) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *User) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *User) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *User) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *User) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *User) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *User) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *User) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTimeZone

`func (o *User) GetTimeZone() string`

GetTimeZone returns the TimeZone field if non-nil, zero value otherwise.

### GetTimeZoneOk

`func (o *User) GetTimeZoneOk() (*string, bool)`

GetTimeZoneOk returns a tuple with the TimeZone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeZone

`func (o *User) SetTimeZone(v string)`

SetTimeZone sets TimeZone field to given value.

### HasTimeZone

`func (o *User) HasTimeZone() bool`

HasTimeZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


