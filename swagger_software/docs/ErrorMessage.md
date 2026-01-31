# ErrorMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable message describing the error. | 
**ErrorTraceId** | Pointer to **string** | An optional trace ID that can be used by Jira developers to locate the source of the error. | [optional] 

## Methods

### NewErrorMessage

`func NewErrorMessage(message string, ) *ErrorMessage`

NewErrorMessage instantiates a new ErrorMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessageWithDefaults

`func NewErrorMessageWithDefaults() *ErrorMessage`

NewErrorMessageWithDefaults instantiates a new ErrorMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorMessage) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMessage) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMessage) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorTraceId

`func (o *ErrorMessage) GetErrorTraceId() string`

GetErrorTraceId returns the ErrorTraceId field if non-nil, zero value otherwise.

### GetErrorTraceIdOk

`func (o *ErrorMessage) GetErrorTraceIdOk() (*string, bool)`

GetErrorTraceIdOk returns a tuple with the ErrorTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTraceId

`func (o *ErrorMessage) SetErrorTraceId(v string)`

SetErrorTraceId sets ErrorTraceId field to given value.

### HasErrorTraceId

`func (o *ErrorMessage) HasErrorTraceId() bool`

HasErrorTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


