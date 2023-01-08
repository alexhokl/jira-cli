# JiraMultipleSelectUserPickerField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Users** | Pointer to [**[]JiraUserField**](JiraUserField.md) |  | [optional] 

## Methods

### NewJiraMultipleSelectUserPickerField

`func NewJiraMultipleSelectUserPickerField(fieldId string, ) *JiraMultipleSelectUserPickerField`

NewJiraMultipleSelectUserPickerField instantiates a new JiraMultipleSelectUserPickerField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraMultipleSelectUserPickerFieldWithDefaults

`func NewJiraMultipleSelectUserPickerFieldWithDefaults() *JiraMultipleSelectUserPickerField`

NewJiraMultipleSelectUserPickerFieldWithDefaults instantiates a new JiraMultipleSelectUserPickerField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraMultipleSelectUserPickerField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraMultipleSelectUserPickerField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraMultipleSelectUserPickerField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetUsers

`func (o *JiraMultipleSelectUserPickerField) GetUsers() []JiraUserField`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *JiraMultipleSelectUserPickerField) GetUsersOk() (*[]JiraUserField, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *JiraMultipleSelectUserPickerField) SetUsers(v []JiraUserField)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *JiraMultipleSelectUserPickerField) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


