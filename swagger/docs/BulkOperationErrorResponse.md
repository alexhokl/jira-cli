# BulkOperationErrorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) |  | [optional] 

## Methods

### NewBulkOperationErrorResponse

`func NewBulkOperationErrorResponse() *BulkOperationErrorResponse`

NewBulkOperationErrorResponse instantiates a new BulkOperationErrorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkOperationErrorResponseWithDefaults

`func NewBulkOperationErrorResponseWithDefaults() *BulkOperationErrorResponse`

NewBulkOperationErrorResponseWithDefaults instantiates a new BulkOperationErrorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *BulkOperationErrorResponse) GetErrors() []ErrorMessage`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkOperationErrorResponse) GetErrorsOk() (*[]ErrorMessage, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkOperationErrorResponse) SetErrors(v []ErrorMessage)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *BulkOperationErrorResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


