# Reviewer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Deprecated. The name of this reviewer. Max length is 255 characters. | [optional] 
**ApprovalStatus** | Pointer to **string** | The approval status of this reviewer, default is UNAPPROVED. | [optional] 
**Url** | Pointer to **string** | Deprecated. The URL of the profile for this reviewer. Max length is 2000 characters. | [optional] 
**Avatar** | Pointer to **string** | Deprecated. The URL of the avatar for this reviewer. Max length is 2000 characters. | [optional] 
**Email** | Pointer to **string** | The email address of this reviewer. Max length is 254 characters. | [optional] 
**AccountId** | Pointer to **string** | The Atlassian Account ID (AAID) of this reviewer. Max length is 128 characters. | [optional] 

## Methods

### NewReviewer

`func NewReviewer() *Reviewer`

NewReviewer instantiates a new Reviewer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReviewerWithDefaults

`func NewReviewerWithDefaults() *Reviewer`

NewReviewerWithDefaults instantiates a new Reviewer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Reviewer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Reviewer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Reviewer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Reviewer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApprovalStatus

`func (o *Reviewer) GetApprovalStatus() string`

GetApprovalStatus returns the ApprovalStatus field if non-nil, zero value otherwise.

### GetApprovalStatusOk

`func (o *Reviewer) GetApprovalStatusOk() (*string, bool)`

GetApprovalStatusOk returns a tuple with the ApprovalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovalStatus

`func (o *Reviewer) SetApprovalStatus(v string)`

SetApprovalStatus sets ApprovalStatus field to given value.

### HasApprovalStatus

`func (o *Reviewer) HasApprovalStatus() bool`

HasApprovalStatus returns a boolean if a field has been set.

### GetUrl

`func (o *Reviewer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *Reviewer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *Reviewer) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *Reviewer) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetAvatar

`func (o *Reviewer) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Reviewer) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Reviewer) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Reviewer) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetEmail

`func (o *Reviewer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Reviewer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Reviewer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Reviewer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetAccountId

`func (o *Reviewer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *Reviewer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *Reviewer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *Reviewer) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


