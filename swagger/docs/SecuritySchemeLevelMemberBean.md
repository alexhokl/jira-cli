# SecuritySchemeLevelMemberBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameter** | Pointer to **string** | The value corresponding to the specified member type. | [optional] 
**Type** | **string** | The issue security level member type, e.g &#x60;reporter&#x60;, &#x60;group&#x60;, &#x60;user&#x60;, &#x60;projectrole&#x60;, &#x60;applicationRole&#x60;. | 

## Methods

### NewSecuritySchemeLevelMemberBean

`func NewSecuritySchemeLevelMemberBean(type_ string, ) *SecuritySchemeLevelMemberBean`

NewSecuritySchemeLevelMemberBean instantiates a new SecuritySchemeLevelMemberBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeLevelMemberBeanWithDefaults

`func NewSecuritySchemeLevelMemberBeanWithDefaults() *SecuritySchemeLevelMemberBean`

NewSecuritySchemeLevelMemberBeanWithDefaults instantiates a new SecuritySchemeLevelMemberBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameter

`func (o *SecuritySchemeLevelMemberBean) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *SecuritySchemeLevelMemberBean) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *SecuritySchemeLevelMemberBean) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *SecuritySchemeLevelMemberBean) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetType

`func (o *SecuritySchemeLevelMemberBean) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecuritySchemeLevelMemberBean) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecuritySchemeLevelMemberBean) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


