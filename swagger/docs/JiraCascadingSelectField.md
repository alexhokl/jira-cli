# JiraCascadingSelectField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ChildOptionValue** | Pointer to [**JiraSelectedOptionField**](JiraSelectedOptionField.md) |  | [optional] 
**FieldId** | **string** |  | 
**ParentOptionValue** | [**JiraSelectedOptionField**](JiraSelectedOptionField.md) |  | 

## Methods

### NewJiraCascadingSelectField

`func NewJiraCascadingSelectField(fieldId string, parentOptionValue JiraSelectedOptionField, ) *JiraCascadingSelectField`

NewJiraCascadingSelectField instantiates a new JiraCascadingSelectField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraCascadingSelectFieldWithDefaults

`func NewJiraCascadingSelectFieldWithDefaults() *JiraCascadingSelectField`

NewJiraCascadingSelectFieldWithDefaults instantiates a new JiraCascadingSelectField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChildOptionValue

`func (o *JiraCascadingSelectField) GetChildOptionValue() JiraSelectedOptionField`

GetChildOptionValue returns the ChildOptionValue field if non-nil, zero value otherwise.

### GetChildOptionValueOk

`func (o *JiraCascadingSelectField) GetChildOptionValueOk() (*JiraSelectedOptionField, bool)`

GetChildOptionValueOk returns a tuple with the ChildOptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildOptionValue

`func (o *JiraCascadingSelectField) SetChildOptionValue(v JiraSelectedOptionField)`

SetChildOptionValue sets ChildOptionValue field to given value.

### HasChildOptionValue

`func (o *JiraCascadingSelectField) HasChildOptionValue() bool`

HasChildOptionValue returns a boolean if a field has been set.

### GetFieldId

`func (o *JiraCascadingSelectField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraCascadingSelectField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraCascadingSelectField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetParentOptionValue

`func (o *JiraCascadingSelectField) GetParentOptionValue() JiraSelectedOptionField`

GetParentOptionValue returns the ParentOptionValue field if non-nil, zero value otherwise.

### GetParentOptionValueOk

`func (o *JiraCascadingSelectField) GetParentOptionValueOk() (*JiraSelectedOptionField, bool)`

GetParentOptionValueOk returns a tuple with the ParentOptionValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOptionValue

`func (o *JiraCascadingSelectField) SetParentOptionValue(v JiraSelectedOptionField)`

SetParentOptionValue sets ParentOptionValue field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


