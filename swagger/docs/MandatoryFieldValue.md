# MandatoryFieldValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retain** | Pointer to **NullableBool** | If &#x60;true&#x60;, will try to retain original non-null issue field values on move. | [optional] [default to true]
**Type** | Pointer to **NullableString** | Will treat as &#x60;MandatoryFieldValue&#x60; if type is &#x60;raw&#x60; or &#x60;empty&#x60; | [optional] [default to "raw"]
**Value** | **[]string** | Value for each field. Provide a &#x60;list of strings&#x60; for non-ADF fields. | 

## Methods

### NewMandatoryFieldValue

`func NewMandatoryFieldValue(value []string, ) *MandatoryFieldValue`

NewMandatoryFieldValue instantiates a new MandatoryFieldValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandatoryFieldValueWithDefaults

`func NewMandatoryFieldValueWithDefaults() *MandatoryFieldValue`

NewMandatoryFieldValueWithDefaults instantiates a new MandatoryFieldValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetain

`func (o *MandatoryFieldValue) GetRetain() bool`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *MandatoryFieldValue) GetRetainOk() (*bool, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *MandatoryFieldValue) SetRetain(v bool)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *MandatoryFieldValue) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### SetRetainNil

`func (o *MandatoryFieldValue) SetRetainNil(b bool)`

 SetRetainNil sets the value for Retain to be an explicit nil

### UnsetRetain
`func (o *MandatoryFieldValue) UnsetRetain()`

UnsetRetain ensures that no value is present for Retain, not even an explicit nil
### GetType

`func (o *MandatoryFieldValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MandatoryFieldValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MandatoryFieldValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MandatoryFieldValue) HasType() bool`

HasType returns a boolean if a field has been set.

### SetTypeNil

`func (o *MandatoryFieldValue) SetTypeNil(b bool)`

 SetTypeNil sets the value for Type to be an explicit nil

### UnsetType
`func (o *MandatoryFieldValue) UnsetType()`

UnsetType ensures that no value is present for Type, not even an explicit nil
### GetValue

`func (o *MandatoryFieldValue) GetValue() []string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MandatoryFieldValue) GetValueOk() (*[]string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MandatoryFieldValue) SetValue(v []string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


