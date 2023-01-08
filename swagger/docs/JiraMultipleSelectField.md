# JiraMultipleSelectField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Options** | [**[]JiraSelectedOptionField**](JiraSelectedOptionField.md) |  | 

## Methods

### NewJiraMultipleSelectField

`func NewJiraMultipleSelectField(fieldId string, options []JiraSelectedOptionField, ) *JiraMultipleSelectField`

NewJiraMultipleSelectField instantiates a new JiraMultipleSelectField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraMultipleSelectFieldWithDefaults

`func NewJiraMultipleSelectFieldWithDefaults() *JiraMultipleSelectField`

NewJiraMultipleSelectFieldWithDefaults instantiates a new JiraMultipleSelectField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraMultipleSelectField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraMultipleSelectField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraMultipleSelectField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetOptions

`func (o *JiraMultipleSelectField) GetOptions() []JiraSelectedOptionField`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *JiraMultipleSelectField) GetOptionsOk() (*[]JiraSelectedOptionField, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *JiraMultipleSelectField) SetOptions(v []JiraSelectedOptionField)`

SetOptions sets Options field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


