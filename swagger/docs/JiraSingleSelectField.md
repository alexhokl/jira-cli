# JiraSingleSelectField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Option** | [**JiraSelectedOptionField**](JiraSelectedOptionField.md) |  | 

## Methods

### NewJiraSingleSelectField

`func NewJiraSingleSelectField(fieldId string, option JiraSelectedOptionField, ) *JiraSingleSelectField`

NewJiraSingleSelectField instantiates a new JiraSingleSelectField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraSingleSelectFieldWithDefaults

`func NewJiraSingleSelectFieldWithDefaults() *JiraSingleSelectField`

NewJiraSingleSelectFieldWithDefaults instantiates a new JiraSingleSelectField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraSingleSelectField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraSingleSelectField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraSingleSelectField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetOption

`func (o *JiraSingleSelectField) GetOption() JiraSelectedOptionField`

GetOption returns the Option field if non-nil, zero value otherwise.

### GetOptionOk

`func (o *JiraSingleSelectField) GetOptionOk() (*JiraSelectedOptionField, bool)`

GetOptionOk returns a tuple with the Option field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOption

`func (o *JiraSingleSelectField) SetOption(v JiraSelectedOptionField)`

SetOption sets Option field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


