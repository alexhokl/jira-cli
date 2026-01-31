# ErrorMessage1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Message** | **string** | A human-readable message describing the error. | 
**ErrorTraceId** | Pointer to **string** | An optional trace ID that can be used by Jira developers to locate the source of the error. | [optional] 

## Methods

### NewErrorMessage1

`func NewErrorMessage1(message string, ) *ErrorMessage1`

NewErrorMessage1 instantiates a new ErrorMessage1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorMessage1WithDefaults

`func NewErrorMessage1WithDefaults() *ErrorMessage1`

NewErrorMessage1WithDefaults instantiates a new ErrorMessage1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessage

`func (o *ErrorMessage1) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorMessage1) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorMessage1) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetErrorTraceId

`func (o *ErrorMessage1) GetErrorTraceId() string`

GetErrorTraceId returns the ErrorTraceId field if non-nil, zero value otherwise.

### GetErrorTraceIdOk

`func (o *ErrorMessage1) GetErrorTraceIdOk() (*string, bool)`

GetErrorTraceIdOk returns a tuple with the ErrorTraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorTraceId

`func (o *ErrorMessage1) SetErrorTraceId(v string)`

SetErrorTraceId sets ErrorTraceId field to given value.

### HasErrorTraceId

`func (o *ErrorMessage1) HasErrorTraceId() bool`

HasErrorTraceId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


