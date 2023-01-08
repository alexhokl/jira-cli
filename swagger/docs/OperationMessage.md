# OperationMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | The human-readable message that describes the result. | 
**StatusCode** | **int32** | The status code of the response. | 

## Methods

### NewOperationMessage

`func NewOperationMessage(message string, statusCode int32, ) *OperationMessage`

NewOperationMessage instantiates a new OperationMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationMessageWithDefaults

`func NewOperationMessageWithDefaults() *OperationMessage`

NewOperationMessageWithDefaults instantiates a new OperationMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *OperationMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *OperationMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *OperationMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetStatusCode

`func (o *OperationMessage) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *OperationMessage) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *OperationMessage) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


