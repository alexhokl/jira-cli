# DetailedErrorCollection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to **map[string]interface{}** | Map of objects representing additional details for an error | [optional] 
**ErrorMessages** | Pointer to **[]string** | The list of error messages produced by this operation. For example, \&quot;input parameter &#39;key&#39; must be provided\&quot; | [optional] 
**Errors** | Pointer to **map[string]string** | The list of errors by parameter returned by the operation. For example,\&quot;projectKey\&quot;: \&quot;Project keys must start with an uppercase letter, followed by one or more uppercase alphanumeric characters.\&quot; | [optional] 

## Methods

### NewDetailedErrorCollection

`func NewDetailedErrorCollection() *DetailedErrorCollection`

NewDetailedErrorCollection instantiates a new DetailedErrorCollection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDetailedErrorCollectionWithDefaults

`func NewDetailedErrorCollectionWithDefaults() *DetailedErrorCollection`

NewDetailedErrorCollectionWithDefaults instantiates a new DetailedErrorCollection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *DetailedErrorCollection) GetDetails() map[string]interface{}`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *DetailedErrorCollection) GetDetailsOk() (*map[string]interface{}, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *DetailedErrorCollection) SetDetails(v map[string]interface{})`

SetDetails sets Details field to given value.

### HasDetails

`func (o *DetailedErrorCollection) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetErrorMessages

`func (o *DetailedErrorCollection) GetErrorMessages() []string`

GetErrorMessages returns the ErrorMessages field if non-nil, zero value otherwise.

### GetErrorMessagesOk

`func (o *DetailedErrorCollection) GetErrorMessagesOk() (*[]string, bool)`

GetErrorMessagesOk returns a tuple with the ErrorMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessages

`func (o *DetailedErrorCollection) SetErrorMessages(v []string)`

SetErrorMessages sets ErrorMessages field to given value.

### HasErrorMessages

`func (o *DetailedErrorCollection) HasErrorMessages() bool`

HasErrorMessages returns a boolean if a field has been set.

### GetErrors

`func (o *DetailedErrorCollection) GetErrors() map[string]string`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DetailedErrorCollection) GetErrorsOk() (*map[string]string, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DetailedErrorCollection) SetErrors(v map[string]string)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DetailedErrorCollection) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


