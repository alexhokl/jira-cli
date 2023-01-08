# VersionApprover

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountId** | Pointer to **string** | The Atlassian account ID of the approver. | [optional] [readonly] 
**DeclineReason** | Pointer to **string** | A description of why the user is declining the approval. | [optional] [readonly] 
**Description** | Pointer to **string** | A description of what the user is approving within the specified version. | [optional] [readonly] 
**Status** | Pointer to **string** | The status of the approval, which can be *PENDING*, *APPROVED*, or *DECLINED* | [optional] [readonly] 

## Methods

### NewVersionApprover

`func NewVersionApprover() *VersionApprover`

NewVersionApprover instantiates a new VersionApprover object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVersionApproverWithDefaults

`func NewVersionApproverWithDefaults() *VersionApprover`

NewVersionApproverWithDefaults instantiates a new VersionApprover object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountId

`func (o *VersionApprover) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *VersionApprover) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *VersionApprover) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *VersionApprover) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetDeclineReason

`func (o *VersionApprover) GetDeclineReason() string`

GetDeclineReason returns the DeclineReason field if non-nil, zero value otherwise.

### GetDeclineReasonOk

`func (o *VersionApprover) GetDeclineReasonOk() (*string, bool)`

GetDeclineReasonOk returns a tuple with the DeclineReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclineReason

`func (o *VersionApprover) SetDeclineReason(v string)`

SetDeclineReason sets DeclineReason field to given value.

### HasDeclineReason

`func (o *VersionApprover) HasDeclineReason() bool`

HasDeclineReason returns a boolean if a field has been set.

### GetDescription

`func (o *VersionApprover) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *VersionApprover) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *VersionApprover) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *VersionApprover) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStatus

`func (o *VersionApprover) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *VersionApprover) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *VersionApprover) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *VersionApprover) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


