# BulkEditGetFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EndingBefore** | Pointer to **string** | The end cursor for use in pagination. | [optional] [readonly] 
**Fields** | Pointer to [**[]IssueBulkEditField**](IssueBulkEditField.md) | List of all the fields | [optional] [readonly] 
**StartingAfter** | Pointer to **string** | The start cursor for use in pagination. | [optional] [readonly] 

## Methods

### NewBulkEditGetFields

`func NewBulkEditGetFields() *BulkEditGetFields`

NewBulkEditGetFields instantiates a new BulkEditGetFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditGetFieldsWithDefaults

`func NewBulkEditGetFieldsWithDefaults() *BulkEditGetFields`

NewBulkEditGetFieldsWithDefaults instantiates a new BulkEditGetFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndingBefore

`func (o *BulkEditGetFields) GetEndingBefore() string`

GetEndingBefore returns the EndingBefore field if non-nil, zero value otherwise.

### GetEndingBeforeOk

`func (o *BulkEditGetFields) GetEndingBeforeOk() (*string, bool)`

GetEndingBeforeOk returns a tuple with the EndingBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingBefore

`func (o *BulkEditGetFields) SetEndingBefore(v string)`

SetEndingBefore sets EndingBefore field to given value.

### HasEndingBefore

`func (o *BulkEditGetFields) HasEndingBefore() bool`

HasEndingBefore returns a boolean if a field has been set.

### GetFields

`func (o *BulkEditGetFields) GetFields() []IssueBulkEditField`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *BulkEditGetFields) GetFieldsOk() (*[]IssueBulkEditField, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *BulkEditGetFields) SetFields(v []IssueBulkEditField)`

SetFields sets Fields field to given value.

### HasFields

`func (o *BulkEditGetFields) HasFields() bool`

HasFields returns a boolean if a field has been set.

### GetStartingAfter

`func (o *BulkEditGetFields) GetStartingAfter() string`

GetStartingAfter returns the StartingAfter field if non-nil, zero value otherwise.

### GetStartingAfterOk

`func (o *BulkEditGetFields) GetStartingAfterOk() (*string, bool)`

GetStartingAfterOk returns a tuple with the StartingAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAfter

`func (o *BulkEditGetFields) SetStartingAfter(v string)`

SetStartingAfter sets StartingAfter field to given value.

### HasStartingAfter

`func (o *BulkEditGetFields) HasStartingAfter() bool`

HasStartingAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


