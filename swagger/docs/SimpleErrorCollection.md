# SimpleErrorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessages** | Pointer to **[]string** | The list of error messages produced by this operation. For example, \&quot;input parameter &#39;key&#39; must be provided\&quot; | [optional] 
**Errors** | Pointer to **map[string]string** | The list of errors by parameter returned by the operation. For example,\&quot;projectKey\&quot;: \&quot;Project keys must start with an uppercase letter, followed by one or more uppercase alphanumeric characters.\&quot; | [optional] 
**HttpStatusCode** | Pointer to **int32** |  | [optional] 

## Methods

### NewSimpleErrorCollection

`func NewSimpleErrorCollection() *SimpleErrorCollection`

NewSimpleErrorCollection instantiates a new SimpleErrorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleErrorCollectionWithDefaults

`func NewSimpleErrorCollectionWithDefaults() *SimpleErrorCollection`

NewSimpleErrorCollectionWithDefaults instantiates a new SimpleErrorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessages

`func (o *SimpleErrorCollection) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *SimpleErrorCollection) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *SimpleErrorCollection) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *SimpleErrorCollection) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetErrors

`func (o *SimpleErrorCollection) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *SimpleErrorCollection) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *SimpleErrorCollection) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *SimpleErrorCollection) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetHttpStatusCode

`func (o *SimpleErrorCollection) GetHttpStatusCode() int32`

GetHttpStatusCode returns the HttpStatusCode field if non-nil, zero value otherwise.

### GetHttpStatusCodeOk

`func (o *SimpleErrorCollection) GetHttpStatusCodeOk() (*int32, bool)`

GetHttpStatusCodeOk returns a tuple with the HttpStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatusCode

`func (o *SimpleErrorCollection) SetHttpStatusCode(v int32)`

SetHttpStatusCode sets HttpStatusCode field to given value.

### HasHttpStatusCode

`func (o *SimpleErrorCollection) HasHttpStatusCode() bool`

HasHttpStatusCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


