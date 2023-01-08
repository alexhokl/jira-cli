# JiraDateField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | Pointer to [**JiraDateInput**](JiraDateInput.md) |  | [optional] 
**FieldId** | **string** |  | 

## Methods

### NewJiraDateField

`func NewJiraDateField(fieldId string, ) *JiraDateField`

NewJiraDateField instantiates a new JiraDateField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraDateFieldWithDefaults

`func NewJiraDateFieldWithDefaults() *JiraDateField`

NewJiraDateFieldWithDefaults instantiates a new JiraDateField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *JiraDateField) GetDate() JiraDateInput`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *JiraDateField) GetDateOk() (*JiraDateInput, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *JiraDateField) SetDate(v JiraDateInput)`

SetDate sets Date field to given value.

### HasDate

`func (o *JiraDateField) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetFieldId

`func (o *JiraDateField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraDateField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraDateField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


