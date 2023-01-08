# CustomFieldReplacement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomFieldId** | Pointer to **int64** | The ID of the custom field in which to replace the version number. | [optional] 
**MoveTo** | Pointer to **int64** | The version number to use as a replacement for the deleted version. | [optional] 

## Methods

### NewCustomFieldReplacement

`func NewCustomFieldReplacement() *CustomFieldReplacement`

NewCustomFieldReplacement instantiates a new CustomFieldReplacement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldReplacementWithDefaults

`func NewCustomFieldReplacementWithDefaults() *CustomFieldReplacement`

NewCustomFieldReplacementWithDefaults instantiates a new CustomFieldReplacement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomFieldId

`func (o *CustomFieldReplacement) GetCustomFieldId() int64`

GetCustomFieldId returns the CustomFieldId field if non-nil, zero value otherwise.

### GetCustomFieldIdOk

`func (o *CustomFieldReplacement) GetCustomFieldIdOk() (*int64, bool)`

GetCustomFieldIdOk returns a tuple with the CustomFieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomFieldId

`func (o *CustomFieldReplacement) SetCustomFieldId(v int64)`

SetCustomFieldId sets CustomFieldId field to given value.

### HasCustomFieldId

`func (o *CustomFieldReplacement) HasCustomFieldId() bool`

HasCustomFieldId returns a boolean if a field has been set.

### GetMoveTo

`func (o *CustomFieldReplacement) GetMoveTo() int64`

GetMoveTo returns the MoveTo field if non-nil, zero value otherwise.

### GetMoveToOk

`func (o *CustomFieldReplacement) GetMoveToOk() (*int64, bool)`

GetMoveToOk returns a tuple with the MoveTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTo

`func (o *CustomFieldReplacement) SetMoveTo(v int64)`

SetMoveTo sets MoveTo field to given value.

### HasMoveTo

`func (o *CustomFieldReplacement) HasMoveTo() bool`

HasMoveTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


