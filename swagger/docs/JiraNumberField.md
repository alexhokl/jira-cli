# JiraNumberField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewJiraNumberField

`func NewJiraNumberField(fieldId string, ) *JiraNumberField`

NewJiraNumberField instantiates a new JiraNumberField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraNumberFieldWithDefaults

`func NewJiraNumberFieldWithDefaults() *JiraNumberField`

NewJiraNumberFieldWithDefaults instantiates a new JiraNumberField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraNumberField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraNumberField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraNumberField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetValue

`func (o *JiraNumberField) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JiraNumberField) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JiraNumberField) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *JiraNumberField) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


