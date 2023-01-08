# SecurityLevelPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the security level | [optional] 
**IsDefault** | Pointer to **bool** | Whether the security level is default for the security scheme | [optional] 
**Name** | Pointer to **string** | The name of the security level | [optional] 
**SecurityLevelMembers** | Pointer to [**[]SecurityLevelMemberPayload**](SecurityLevelMemberPayload.md) | The members of the security level | [optional] 

## Methods

### NewSecurityLevelPayload

`func NewSecurityLevelPayload() *SecurityLevelPayload`

NewSecurityLevelPayload instantiates a new SecurityLevelPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLevelPayloadWithDefaults

`func NewSecurityLevelPayloadWithDefaults() *SecurityLevelPayload`

NewSecurityLevelPayloadWithDefaults instantiates a new SecurityLevelPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityLevelPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityLevelPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityLevelPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityLevelPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *SecurityLevelPayload) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SecurityLevelPayload) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SecurityLevelPayload) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SecurityLevelPayload) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetName

`func (o *SecurityLevelPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityLevelPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityLevelPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityLevelPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSecurityLevelMembers

`func (o *SecurityLevelPayload) GetSecurityLevelMembers() []SecurityLevelMemberPayload`

GetSecurityLevelMembers returns the SecurityLevelMembers field if non-nil, zero value otherwise.

### GetSecurityLevelMembersOk

`func (o *SecurityLevelPayload) GetSecurityLevelMembersOk() (*[]SecurityLevelMemberPayload, bool)`

GetSecurityLevelMembersOk returns a tuple with the SecurityLevelMembers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevelMembers

`func (o *SecurityLevelPayload) SetSecurityLevelMembers(v []SecurityLevelMemberPayload)`

SetSecurityLevelMembers sets SecurityLevelMembers field to given value.

### HasSecurityLevelMembers

`func (o *SecurityLevelPayload) HasSecurityLevelMembers() bool`

HasSecurityLevelMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


