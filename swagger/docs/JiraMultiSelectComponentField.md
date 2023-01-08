# JiraMultiSelectComponentField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BulkEditMultiSelectFieldOption** | **string** |  | 
**Components** | [**[]JiraComponentField**](JiraComponentField.md) |  | 
**FieldId** | **string** |  | 

## Methods

### NewJiraMultiSelectComponentField

`func NewJiraMultiSelectComponentField(bulkEditMultiSelectFieldOption string, components []JiraComponentField, fieldId string, ) *JiraMultiSelectComponentField`

NewJiraMultiSelectComponentField instantiates a new JiraMultiSelectComponentField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJiraMultiSelectComponentFieldWithDefaults

`func NewJiraMultiSelectComponentFieldWithDefaults() *JiraMultiSelectComponentField`

NewJiraMultiSelectComponentFieldWithDefaults instantiates a new JiraMultiSelectComponentField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBulkEditMultiSelectFieldOption

`func (o *JiraMultiSelectComponentField) GetBulkEditMultiSelectFieldOption() string`

GetBulkEditMultiSelectFieldOption returns the BulkEditMultiSelectFieldOption field if non-nil, zero value otherwise.

### GetBulkEditMultiSelectFieldOptionOk

`func (o *JiraMultiSelectComponentField) GetBulkEditMultiSelectFieldOptionOk() (*string, bool)`

GetBulkEditMultiSelectFieldOptionOk returns a tuple with the BulkEditMultiSelectFieldOption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkEditMultiSelectFieldOption

`func (o *JiraMultiSelectComponentField) SetBulkEditMultiSelectFieldOption(v string)`

SetBulkEditMultiSelectFieldOption sets BulkEditMultiSelectFieldOption field to given value.


### GetComponents

`func (o *JiraMultiSelectComponentField) GetComponents() []JiraComponentField`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *JiraMultiSelectComponentField) GetComponentsOk() (*[]JiraComponentField, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *JiraMultiSelectComponentField) SetComponents(v []JiraComponentField)`

SetComponents sets Components field to given value.


### GetFieldId

`func (o *JiraMultiSelectComponentField) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *JiraMultiSelectComponentField) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *JiraMultiSelectComponentField) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


