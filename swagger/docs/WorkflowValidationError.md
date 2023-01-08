# WorkflowValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | An error code. | [optional] 
**ElementReference** | Pointer to [**WorkflowElementReference**](WorkflowElementReference.md) |  | [optional] 
**Level** | Pointer to **string** | The validation error level. | [optional] 
**Message** | Pointer to **string** | An error message. | [optional] 
**Type** | Pointer to **string** | The type of element the error or warning references. | [optional] 

## Methods

### NewWorkflowValidationError

`func NewWorkflowValidationError() *WorkflowValidationError`

NewWorkflowValidationError instantiates a new WorkflowValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWorkflowValidationErrorWithDefaults

`func NewWorkflowValidationErrorWithDefaults() *WorkflowValidationError`

NewWorkflowValidationErrorWithDefaults instantiates a new WorkflowValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *WorkflowValidationError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *WorkflowValidationError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *WorkflowValidationError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *WorkflowValidationError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetElementReference

`func (o *WorkflowValidationError) GetElementReference() WorkflowElementReference`

GetElementReference returns the ElementReference field if non-nil, zero value otherwise.

### GetElementReferenceOk

`func (o *WorkflowValidationError) GetElementReferenceOk() (*WorkflowElementReference, bool)`

GetElementReferenceOk returns a tuple with the ElementReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElementReference

`func (o *WorkflowValidationError) SetElementReference(v WorkflowElementReference)`

SetElementReference sets ElementReference field to given value.

### HasElementReference

`func (o *WorkflowValidationError) HasElementReference() bool`

HasElementReference returns a boolean if a field has been set.

### GetLevel

`func (o *WorkflowValidationError) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *WorkflowValidationError) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *WorkflowValidationError) SetLevel(v string)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *WorkflowValidationError) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetMessage

`func (o *WorkflowValidationError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *WorkflowValidationError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *WorkflowValidationError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *WorkflowValidationError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetType

`func (o *WorkflowValidationError) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *WorkflowValidationError) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *WorkflowValidationError) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *WorkflowValidationError) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


