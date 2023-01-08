# SecuritySchemeLevelBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the issue security scheme level. | [optional] 
**IsDefault** | Pointer to **bool** | Specifies whether the level is the default level. False by default. | [optional] 
**Members** | Pointer to [**[]SecuritySchemeLevelMemberBean**](SecuritySchemeLevelMemberBean.md) | The list of level members which should be added to the issue security scheme level. | [optional] 
**Name** | **string** | The name of the issue security scheme level. Must be unique. | 

## Methods

### NewSecuritySchemeLevelBean

`func NewSecuritySchemeLevelBean(name string, ) *SecuritySchemeLevelBean`

NewSecuritySchemeLevelBean instantiates a new SecuritySchemeLevelBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeLevelBeanWithDefaults

`func NewSecuritySchemeLevelBeanWithDefaults() *SecuritySchemeLevelBean`

NewSecuritySchemeLevelBeanWithDefaults instantiates a new SecuritySchemeLevelBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecuritySchemeLevelBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySchemeLevelBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySchemeLevelBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecuritySchemeLevelBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetIsDefault

`func (o *SecuritySchemeLevelBean) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SecuritySchemeLevelBean) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SecuritySchemeLevelBean) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SecuritySchemeLevelBean) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMembers

`func (o *SecuritySchemeLevelBean) GetMembers() []SecuritySchemeLevelMemberBean`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SecuritySchemeLevelBean) GetMembersOk() (*[]SecuritySchemeLevelMemberBean, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SecuritySchemeLevelBean) SetMembers(v []SecuritySchemeLevelMemberBean)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SecuritySchemeLevelBean) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetName

`func (o *SecuritySchemeLevelBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySchemeLevelBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySchemeLevelBean) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


