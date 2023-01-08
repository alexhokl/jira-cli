# JiraSingleGroupPickerField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Group** | [**JiraGroupInput**](JiraGroupInput.md) |  | 

## Methods

### NewJiraSingleGroupPickerField

`func NewJiraSingleGroupPickerField(fieldId string, group JiraGroupInput, ) *JiraSingleGroupPickerField`

NewJiraSingleGroupPickerField instantiates a new JiraSingleGroupPickerField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraSingleGroupPickerFieldWithDefaults

`func NewJiraSingleGroupPickerFieldWithDefaults() *JiraSingleGroupPickerField`

NewJiraSingleGroupPickerFieldWithDefaults instantiates a new JiraSingleGroupPickerField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraSingleGroupPickerField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraSingleGroupPickerField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraSingleGroupPickerField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetGroup

`func (o *JiraSingleGroupPickerField) GetGroup() JiraGroupInput`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *JiraSingleGroupPickerField) GetGroupOk() (*JiraGroupInput, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *JiraSingleGroupPickerField) SetGroup(v JiraGroupInput)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


