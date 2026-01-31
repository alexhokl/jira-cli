# UserBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**Active** | Pointer to **bool** | Whether the user is active. | [optional] 
**AvatarUrls** | Pointer to [**GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls**](GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls.md) |  | [optional] 
**DisplayName** | Pointer to **string** | The display name of the user. Depending on the user’s privacy setting, this may return an alternative value. | [optional] 
**Key** | Pointer to **string** | This property is deprecated in favor of &#x60;accountId&#x60; because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.   The key of the user. | [optional] 
**Name** | Pointer to **string** | This property is deprecated in favor of &#x60;accountId&#x60; because of privacy changes. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details.   The username of the user. | [optional] 
**Self** | Pointer to **string** | The URL of the user. | [optional] 

## Methods

### NewUserBean

`func NewUserBean() *UserBean`

NewUserBean instantiates a new UserBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserBeanWithDefaults

`func NewUserBeanWithDefaults() *UserBean`

NewUserBeanWithDefaults instantiates a new UserBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UserBean) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UserBean) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UserBean) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UserBean) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *UserBean) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UserBean) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UserBean) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UserBean) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvatarUrls

`func (o *UserBean) GetAvatarUrls() GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *UserBean) GetAvatarUrlsOk() (*GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *UserBean) SetAvatarUrls(v GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *UserBean) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetDisplayName

`func (o *UserBean) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *UserBean) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *UserBean) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *UserBean) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *UserBean) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *UserBean) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *UserBean) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *UserBean) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *UserBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UserBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UserBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UserBean) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *UserBean) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UserBean) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UserBean) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *UserBean) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


