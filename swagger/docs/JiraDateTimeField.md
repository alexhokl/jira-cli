# JiraDateTimeField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DateTime** | [**JiraDateTimeInput**](JiraDateTimeInput.md) |  | 
**FieldId** | **string** |  | 

## Methods

### NewJiraDateTimeField

`func NewJiraDateTimeField(dateTime JiraDateTimeInput, fieldId string, ) *JiraDateTimeField`

NewJiraDateTimeField instantiates a new JiraDateTimeField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraDateTimeFieldWithDefaults

`func NewJiraDateTimeFieldWithDefaults() *JiraDateTimeField`

NewJiraDateTimeFieldWithDefaults instantiates a new JiraDateTimeField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDateTime

`func (o *JiraDateTimeField) GetDateTime() JiraDateTimeInput`

GetDateTime returns the DateTime field if non-nil, zero value otherwise.

### GetDateTimeOk

`func (o *JiraDateTimeField) GetDateTimeOk() (*JiraDateTimeInput, bool)`

GetDateTimeOk returns a tuple with the DateTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateTime

`func (o *JiraDateTimeField) SetDateTime(v JiraDateTimeInput)`

SetDateTime sets DateTime field to given value.


### GetFieldId

`func (o *JiraDateTimeField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraDateTimeField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraDateTimeField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


