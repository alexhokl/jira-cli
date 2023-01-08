# RedactionPosition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdfPointer** | Pointer to **string** | The ADF pointer indicating the position of the text to be redacted. This is only required when redacting from rich text(ADF) fields. For plain text fields, this field can be omitted. | [optional] 
**ExpectedText** | **string** | The text which will be redacted, encoded using SHA256 hash and Base64 digest | 
**From** | **int32** | The start index(inclusive) for the redaction in specified content | 
**To** | **int32** | The ending index(exclusive) for the redaction in specified content | 

## Methods

### NewRedactionPosition

`func NewRedactionPosition(expectedText string, from int32, to int32, ) *RedactionPosition`

NewRedactionPosition instantiates a new RedactionPosition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRedactionPositionWithDefaults

`func NewRedactionPositionWithDefaults() *RedactionPosition`

NewRedactionPositionWithDefaults instantiates a new RedactionPosition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdfPointer

`func (o *RedactionPosition) GetAdfPointer() string`

GetAdfPointer returns the AdfPointer field if non-nil, zero value otherwise.

### GetAdfPointerOk

`func (o *RedactionPosition) GetAdfPointerOk() (*string, bool)`

GetAdfPointerOk returns a tuple with the AdfPointer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdfPointer

`func (o *RedactionPosition) SetAdfPointer(v string)`

SetAdfPointer sets AdfPointer field to given value.

### HasAdfPointer

`func (o *RedactionPosition) HasAdfPointer() bool`

HasAdfPointer returns a boolean if a field has been set.

### GetExpectedText

`func (o *RedactionPosition) GetExpectedText() string`

GetExpectedText returns the ExpectedText field if non-nil, zero value otherwise.

### GetExpectedTextOk

`func (o *RedactionPosition) GetExpectedTextOk() (*string, bool)`

GetExpectedTextOk returns a tuple with the ExpectedText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedText

`func (o *RedactionPosition) SetExpectedText(v string)`

SetExpectedText sets ExpectedText field to given value.


### GetFrom

`func (o *RedactionPosition) GetFrom() int32`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *RedactionPosition) GetFromOk() (*int32, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *RedactionPosition) SetFrom(v int32)`

SetFrom sets From field to given value.


### GetTo

`func (o *RedactionPosition) GetTo() int32`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *RedactionPosition) GetToOk() (*int32, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *RedactionPosition) SetTo(v int32)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


