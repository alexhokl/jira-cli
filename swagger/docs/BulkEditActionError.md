# BulkEditActionError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | **[]string** | The error messages. | 
**Errors** | **map[string]string** | The errors. | 

## Methods

### NewBulkEditActionError

`func NewBulkEditActionError(errorMessages []string, errors map[string]string, ) *BulkEditActionError`

NewBulkEditActionError instantiates a new BulkEditActionError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditActionErrorWithDefaults

`func NewBulkEditActionErrorWithDefaults() *BulkEditActionError`

NewBulkEditActionErrorWithDefaults instantiates a new BulkEditActionError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *BulkEditActionError) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *BulkEditActionError) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *BulkEditActionError) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.


### GetErrors

`func (o *BulkEditActionError) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *BulkEditActionError) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *BulkEditActionError) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


