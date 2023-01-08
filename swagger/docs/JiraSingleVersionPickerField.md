# JiraSingleVersionPickerField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Version** | [**JiraVersionField**](JiraVersionField.md) |  | 

## Methods

### NewJiraSingleVersionPickerField

`func NewJiraSingleVersionPickerField(fieldId string, version JiraVersionField, ) *JiraSingleVersionPickerField`

NewJiraSingleVersionPickerField instantiates a new JiraSingleVersionPickerField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraSingleVersionPickerFieldWithDefaults

`func NewJiraSingleVersionPickerFieldWithDefaults() *JiraSingleVersionPickerField`

NewJiraSingleVersionPickerFieldWithDefaults instantiates a new JiraSingleVersionPickerField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraSingleVersionPickerField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraSingleVersionPickerField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraSingleVersionPickerField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetVersion

`func (o *JiraSingleVersionPickerField) GetVersion() JiraVersionField`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *JiraSingleVersionPickerField) GetVersionOk() (*JiraVersionField, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *JiraSingleVersionPickerField) SetVersion(v JiraVersionField)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


