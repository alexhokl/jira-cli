# JiraMultipleGroupPickerField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Groups** | [**[]JiraGroupInput**](JiraGroupInput.md) |  | 

## Methods

### NewJiraMultipleGroupPickerField

`func NewJiraMultipleGroupPickerField(fieldId string, groups []JiraGroupInput, ) *JiraMultipleGroupPickerField`

NewJiraMultipleGroupPickerField instantiates a new JiraMultipleGroupPickerField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraMultipleGroupPickerFieldWithDefaults

`func NewJiraMultipleGroupPickerFieldWithDefaults() *JiraMultipleGroupPickerField`

NewJiraMultipleGroupPickerFieldWithDefaults instantiates a new JiraMultipleGroupPickerField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *JiraMultipleGroupPickerField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraMultipleGroupPickerField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraMultipleGroupPickerField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetGroups

`func (o *JiraMultipleGroupPickerField) GetGroups() []JiraGroupInput`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *JiraMultipleGroupPickerField) GetGroupsOk() (*[]JiraGroupInput, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *JiraMultipleGroupPickerField) SetGroups(v []JiraGroupInput)`

SetGroups sets Groups field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


