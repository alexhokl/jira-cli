# BulkOperationErrorResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ElementErrors** | Pointer to [**ErrorCollection**](ErrorCollection.md) |  | [optional] 
**FailedElementNumber** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewBulkOperationErrorResult

`func NewBulkOperationErrorResult() *BulkOperationErrorResult`

NewBulkOperationErrorResult instantiates a new BulkOperationErrorResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationErrorResultWithDefaults

`func NewBulkOperationErrorResultWithDefaults() *BulkOperationErrorResult`

NewBulkOperationErrorResultWithDefaults instantiates a new BulkOperationErrorResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElementErrors

`func (o *BulkOperationErrorResult) GetElementErrors() ErrorCollection`

GetElementErrors returns the ElementErrors field if non-nil, zero value otherwise.

### GetElementErrorsOk

`func (o *BulkOperationErrorResult) GetElementErrorsOk() (*ErrorCollection, bool)`

GetElementErrorsOk returns a tuple with the ElementErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementErrors

`func (o *BulkOperationErrorResult) SetElementErrors(v ErrorCollection)`

SetElementErrors sets ElementErrors field to given value.

### HasElementErrors

`func (o *BulkOperationErrorResult) HasElementErrors() bool`

HasElementErrors returns a boolean if a field has been set.

### GetFailedElementNumber

`func (o *BulkOperationErrorResult) GetFailedElementNumber() int32`

GetFailedElementNumber returns the FailedElementNumber field if non-nil, zero value otherwise.

### GetFailedElementNumberOk

`func (o *BulkOperationErrorResult) GetFailedElementNumberOk() (*int32, bool)`

GetFailedElementNumberOk returns a tuple with the FailedElementNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailedElementNumber

`func (o *BulkOperationErrorResult) SetFailedElementNumber(v int32)`

SetFailedElementNumber sets FailedElementNumber field to given value.

### HasFailedElementNumber

`func (o *BulkOperationErrorResult) HasFailedElementNumber() bool`

HasFailedElementNumber returns a boolean if a field has been set.

### GetStatus

`func (o *BulkOperationErrorResult) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BulkOperationErrorResult) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BulkOperationErrorResult) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BulkOperationErrorResult) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


