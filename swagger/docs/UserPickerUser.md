# UserPickerUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**AccountType** | Pointer to **string** | The user account type. Can take the following values:   *  &#x60;atlassian&#x60; regular Atlassian user account  *  &#x60;app&#x60; system account used for Connect applications and OAuth to represent external systems  *  &#x60;customer&#x60; Jira Service Desk account representing an external service desk | [optional] 
**AvatarUrl** | Pointer to **string** | The avatar URL of the user. | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user. Depending on the user’s privacy setting, this may be returned as null. | [optional] 
**Html** | Pointer to **string** | The display name, email address, and key of the user with the matched query string highlighted with the HTML bold tag. | [optional] 
**Key** | Pointer to **string** | This property is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Name** | Pointer to **string** | This property is no longer available . See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 

## Methods

### NewUserPickerUser

`func NewUserPickerUser() *UserPickerUser`

NewUserPickerUser instantiates a new UserPickerUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPickerUserWithDefaults

`func NewUserPickerUserWithDefaults() *UserPickerUser`

NewUserPickerUserWithDefaults instantiates a new UserPickerUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserPickerUser) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserPickerUser) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserPickerUser) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserPickerUser) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetAccountType

`func (o *UserPickerUser) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *UserPickerUser) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *UserPickerUser) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *UserPickerUser) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAvatarUrl

`func (o *UserPickerUser) GetAvatarUrl() string`

GetAvatarUrl returns the AvatarUrl field if non-nil, zero value otherwise.

### GetAvatarUrlOk

`func (o *UserPickerUser) GetAvatarUrlOk() (*string, bool)`

GetAvatarUrlOk returns a tuple with the AvatarUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrl

`func (o *UserPickerUser) SetAvatarUrl(v string)`

SetAvatarUrl sets AvatarUrl field to given value.

### HasAvatarUrl

`func (o *UserPickerUser) HasAvatarUrl() bool`

HasAvatarUrl returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserPickerUser) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserPickerUser) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserPickerUser) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserPickerUser) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetHtml

`func (o *UserPickerUser) GetHtml() string`

GetHtml returns the Html field if non-nil, zero value otherwise.

### GetHtmlOk

`func (o *UserPickerUser) GetHtmlOk() (*string, bool)`

GetHtmlOk returns a tuple with the Html field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtml

`func (o *UserPickerUser) SetHtml(v string)`

SetHtml sets Html field to given value.

### HasHtml

`func (o *UserPickerUser) HasHtml() bool`

HasHtml returns a boolean if a field has been set.

### GetKey

`func (o *UserPickerUser) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserPickerUser) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserPickerUser) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserPickerUser) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UserPickerUser) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserPickerUser) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserPickerUser) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserPickerUser) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


