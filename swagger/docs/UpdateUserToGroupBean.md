# UpdateUserToGroupBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The account ID of the user, which uniquely identifies the user across all Atlassian products. For example, *5b10ac8d82e05b22cc7d4ef5*. | [optional] 
**Name** | Pointer to **string** | This property is no longer available. See the [deprecation notice](https://developer.atlassian.com/cloud/jira/platform/deprecation-notice-user-privacy-api-migration-guide/) for details. | [optional] 

## Methods

### NewUpdateUserToGroupBean

`func NewUpdateUserToGroupBean() *UpdateUserToGroupBean`

NewUpdateUserToGroupBean instantiates a new UpdateUserToGroupBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserToGroupBeanWithDefaults

`func NewUpdateUserToGroupBeanWithDefaults() *UpdateUserToGroupBean`

NewUpdateUserToGroupBeanWithDefaults instantiates a new UpdateUserToGroupBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *UpdateUserToGroupBean) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateUserToGroupBean) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateUserToGroupBean) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateUserToGroupBean) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetName

`func (o *UpdateUserToGroupBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateUserToGroupBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateUserToGroupBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateUserToGroupBean) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


