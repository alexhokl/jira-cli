# MandatoryFieldValueForADF

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retain** | Pointer to **NullableBool** | If &#x60;true&#x60;, will try to retain original non-null issue field values on move. | [optional] [default to true]
**Type** | **string** | Will treat as &#x60;MandatoryFieldValueForADF&#x60; if type is &#x60;adf&#x60; | [default to "raw"]
**Value** | **map[string]interface{}** | Value for each field. Accepts Atlassian Document Format (ADF) for rich text fields like &#x60;description&#x60;, &#x60;environments&#x60;. For ADF format details, refer to: [Atlassian Document Format](https://developer.atlassian.com/cloud/jira/platform/apis/document/structure) | 

## Methods

### NewMandatoryFieldValueForADF

`func NewMandatoryFieldValueForADF(type_ string, value map[string]interface{}, ) *MandatoryFieldValueForADF`

NewMandatoryFieldValueForADF instantiates a new MandatoryFieldValueForADF object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMandatoryFieldValueForADFWithDefaults

`func NewMandatoryFieldValueForADFWithDefaults() *MandatoryFieldValueForADF`

NewMandatoryFieldValueForADFWithDefaults instantiates a new MandatoryFieldValueForADF object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetain

`func (o *MandatoryFieldValueForADF) GetRetain() bool`

GetRetain returns the Retain field if non-nil, zero value otherwise.

### GetRetainOk

`func (o *MandatoryFieldValueForADF) GetRetainOk() (*bool, bool)`

GetRetainOk returns a tuple with the Retain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetain

`func (o *MandatoryFieldValueForADF) SetRetain(v bool)`

SetRetain sets Retain field to given value.

### HasRetain

`func (o *MandatoryFieldValueForADF) HasRetain() bool`

HasRetain returns a boolean if a field has been set.

### SetRetainNil

`func (o *MandatoryFieldValueForADF) SetRetainNil(b bool)`

 SetRetainNil sets the value for Retain to be an explicit nil

### UnsetRetain
`func (o *MandatoryFieldValueForADF) UnsetRetain()`

UnsetRetain ensures that no value is present for Retain, not even an explicit nil
### GetType

`func (o *MandatoryFieldValueForADF) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MandatoryFieldValueForADF) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MandatoryFieldValueForADF) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *MandatoryFieldValueForADF) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MandatoryFieldValueForADF) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MandatoryFieldValueForADF) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


