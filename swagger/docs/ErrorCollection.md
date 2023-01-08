# ErrorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | Pointer to **[]string** | The list of error messages produced by this operation. For example, \&quot;input parameter &#39;key&#39; must be provided\&quot; | [optional] 
**Errors** | Pointer to **map[string]string** | The list of errors by parameter returned by the operation. For example,\&quot;projectKey\&quot;: \&quot;Project keys must start with an uppercase letter, followed by one or more uppercase alphanumeric characters.\&quot; | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewErrorCollection

`func NewErrorCollection() *ErrorCollection`

NewErrorCollection instantiates a new ErrorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorCollectionWithDefaults

`func NewErrorCollectionWithDefaults() *ErrorCollection`

NewErrorCollectionWithDefaults instantiates a new ErrorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *ErrorCollection) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *ErrorCollection) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *ErrorCollection) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *ErrorCollection) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetErrors

`func (o *ErrorCollection) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ErrorCollection) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ErrorCollection) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ErrorCollection) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetStatus

`func (o *ErrorCollection) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ErrorCollection) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ErrorCollection) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ErrorCollection) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


