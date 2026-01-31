# EntityError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Entity id | 
**ErrorMessages** | Pointer to [**[]ErrorMessage**](ErrorMessage.md) | Error message | [optional] 

## Methods

### NewEntityError

`func NewEntityError(id string, ) *EntityError`

NewEntityError instantiates a new EntityError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityErrorWithDefaults

`func NewEntityErrorWithDefaults() *EntityError`

NewEntityErrorWithDefaults instantiates a new EntityError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntityError) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntityError) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntityError) SetId(v string)`

SetId sets Id field to given value.


### GetErrorMessages

`func (o *EntityError) GetErrorMessages() []ErrorMessage`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *EntityError) GetErrorMessagesOk() (*[]ErrorMessage, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *EntityError) SetErrorMessages(v []ErrorMessage)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *EntityError) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


