# ProjectEmailAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The email address. | [optional] 
**EmailAddressStatus** | Pointer to **[]string** | When using a custom domain, the status of the email address. | [optional] 

## Methods

### NewProjectEmailAddress

`func NewProjectEmailAddress() *ProjectEmailAddress`

NewProjectEmailAddress instantiates a new ProjectEmailAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectEmailAddressWithDefaults

`func NewProjectEmailAddressWithDefaults() *ProjectEmailAddress`

NewProjectEmailAddressWithDefaults instantiates a new ProjectEmailAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *ProjectEmailAddress) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *ProjectEmailAddress) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *ProjectEmailAddress) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *ProjectEmailAddress) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEmailAddressStatus

`func (o *ProjectEmailAddress) GetEmailAddressStatus() []string`

GetEmailAddressStatus returns the EmailAddressStatus field if non-nil, zero value otherwise.

### GetEmailAddressStatusOk

`func (o *ProjectEmailAddress) GetEmailAddressStatusOk() (*[]string, bool)`

GetEmailAddressStatusOk returns a tuple with the EmailAddressStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddressStatus

`func (o *ProjectEmailAddress) SetEmailAddressStatus(v []string)`

SetEmailAddressStatus sets EmailAddressStatus field to given value.

### HasEmailAddressStatus

`func (o *ProjectEmailAddress) HasEmailAddressStatus() bool`

HasEmailAddressStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


