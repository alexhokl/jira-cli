# JiraColorField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Color** | [**JiraColorInput**](JiraColorInput.md) |  | 
**FieldId** | **string** |  | 

## Methods

### NewJiraColorField

`func NewJiraColorField(color JiraColorInput, fieldId string, ) *JiraColorField`

NewJiraColorField instantiates a new JiraColorField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraColorFieldWithDefaults

`func NewJiraColorFieldWithDefaults() *JiraColorField`

NewJiraColorFieldWithDefaults instantiates a new JiraColorField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColor

`func (o *JiraColorField) GetColor() JiraColorInput`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *JiraColorField) GetColorOk() (*JiraColorInput, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *JiraColorField) SetColor(v JiraColorInput)`

SetColor sets Color field to given value.


### GetFieldId

`func (o *JiraColorField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraColorField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraColorField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


