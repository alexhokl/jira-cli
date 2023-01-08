# JiraSingleSelectUserPickerField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**User** | Pointer to [**JiraUserField**](JiraUserField.md) |  | [optional] 

## Methods

### NewJiraSingleSelectUserPickerField

`func NewJiraSingleSelectUserPickerField(fieldId string, ) *JiraSingleSelectUserPickerField`

NewJiraSingleSelectUserPickerField instantiates a new JiraSingleSelectUserPickerField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraSingleSelectUserPickerFieldWithDefaults

`func NewJiraSingleSelectUserPickerFieldWithDefaults() *JiraSingleSelectUserPickerField`

NewJiraSingleSelectUserPickerFieldWithDefaults instantiates a new JiraSingleSelectUserPickerField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraSingleSelectUserPickerField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraSingleSelectUserPickerField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraSingleSelectUserPickerField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetUser

`func (o *JiraSingleSelectUserPickerField) GetUser() JiraUserField`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *JiraSingleSelectUserPickerField) GetUserOk() (*JiraUserField, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *JiraSingleSelectUserPickerField) SetUser(v JiraUserField)`

SetUser sets User field to given value.

### HasUser

`func (o *JiraSingleSelectUserPickerField) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


