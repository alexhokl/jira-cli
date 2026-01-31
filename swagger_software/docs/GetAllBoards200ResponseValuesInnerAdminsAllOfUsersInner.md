# GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner

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

### NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner

`func NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner() *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner`

NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner instantiates a new GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerWithDefaults

`func NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerWithDefaults() *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner`

NewGetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerWithDefaults instantiates a new GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetActive

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAvatarUrls

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetAvatarUrls() GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls`

GetAvatarUrls returns the AvatarUrls field if non-nil, zero value otherwise.

### GetAvatarUrlsOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetAvatarUrlsOk() (*GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls, bool)`

GetAvatarUrlsOk returns a tuple with the AvatarUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarUrls

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetAvatarUrls(v GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInnerAvatarUrls)`

SetAvatarUrls sets AvatarUrls field to given value.

### HasAvatarUrls

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasAvatarUrls() bool`

HasAvatarUrls returns a boolean if a field has been set.

### GetDisplayName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetKey

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


