# Author

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Deprecated. The name of this user in a format suitable for display. Max length is 255 characters. | [optional] 
**Email** | Pointer to **string** | The email address of the user. Used to associate the user with a Jira user. Max length is 255 characters. | [optional] 
**Username** | Pointer to **string** | Deprecated. The username of the user. Used to associate the user with a Jira user if there are multiple users for a given email. Max length is 255 characters. | [optional] 
**Url** | Pointer to **string** | Deprecated. The URL of the profile for this user. Max length is 2000 characters. | [optional] 
**Avatar** | Pointer to **string** | Deprecated. The URL of the avatar for this user. Max length is 2000 characters. | [optional] 

## Methods

### NewAuthor

`func NewAuthor() *Author`

NewAuthor instantiates a new Author object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorWithDefaults

`func NewAuthorWithDefaults() *Author`

NewAuthorWithDefaults instantiates a new Author object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Author) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Author) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Author) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Author) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEmail

`func (o *Author) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Author) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Author) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Author) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetUsername

`func (o *Author) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *Author) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *Author) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *Author) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUrl

`func (o *Author) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Author) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Author) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Author) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvatar

`func (o *Author) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Author) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Author) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Author) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


