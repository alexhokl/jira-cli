# JiraNewUserDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApplicationKeys** | Pointer to **[]string** | Deprecated, do not use. | [optional] 
**DisplayName** | Pointer to **string** | This property is no longer available. If the user has an Atlassian account, their display name is not changed. If the user does not have an Atlassian account, they are sent an email asking them set up an account. | [optional] 
**EmailAddress** | **string** | The email address for the user. | 
**Key** | Pointer to **string** | This property is no longer available. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Name** | Pointer to **string** | This property is no longer available. See the [migration guide](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 
**Password** | Pointer to **string** | This property is no longer available. If the user has an Atlassian account, their password is not changed. If the user does not have an Atlassian account, they are sent an email asking them set up an account. | [optional] 
**Products** | **[]string** | Products the new user has access to. Valid products are: jira-core, jira-servicedesk, jira-product-discovery, jira-software. To create a user without product access, set this field to be an empty array. | 
**Self** | Pointer to **string** | The URL of the user. | [optional] [readonly] 

## Methods

### NewJiraNewUserDetails

`func NewJiraNewUserDetails(emailAddress string, products []string, ) *JiraNewUserDetails`

NewJiraNewUserDetails instantiates a new JiraNewUserDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraNewUserDetailsWithDefaults

`func NewJiraNewUserDetailsWithDefaults() *JiraNewUserDetails`

NewJiraNewUserDetailsWithDefaults instantiates a new JiraNewUserDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplicationKeys

`func (o *JiraNewUserDetails) GetApplicationKeys() []string`

GetApplicationKeys returns the ApplicationKeys field if non-nil, zero value otherwise.

### GetApplicationKeysOk

`func (o *JiraNewUserDetails) GetApplicationKeysOk() (*[]string, bool)`

GetApplicationKeysOk returns a tuple with the ApplicationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationKeys

`func (o *JiraNewUserDetails) SetApplicationKeys(v []string)`

SetApplicationKeys sets ApplicationKeys field to given value.

### HasApplicationKeys

`func (o *JiraNewUserDetails) HasApplicationKeys() bool`

HasApplicationKeys returns a boolean if a field has been set.

### GetDisplayName

`func (o *JiraNewUserDetails) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *JiraNewUserDetails) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *JiraNewUserDetails) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *JiraNewUserDetails) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmailAddress

`func (o *JiraNewUserDetails) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *JiraNewUserDetails) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *JiraNewUserDetails) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.


### GetKey

`func (o *JiraNewUserDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JiraNewUserDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JiraNewUserDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *JiraNewUserDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *JiraNewUserDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *JiraNewUserDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *JiraNewUserDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *JiraNewUserDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPassword

`func (o *JiraNewUserDetails) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *JiraNewUserDetails) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *JiraNewUserDetails) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *JiraNewUserDetails) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetProducts

`func (o *JiraNewUserDetails) GetProducts() []string`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *JiraNewUserDetails) GetProductsOk() (*[]string, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *JiraNewUserDetails) SetProducts(v []string)`

SetProducts sets Products field to given value.


### GetSelf

`func (o *JiraNewUserDetails) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *JiraNewUserDetails) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *JiraNewUserDetails) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *JiraNewUserDetails) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


