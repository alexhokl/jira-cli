# JsonContextVariable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type of custom context variable. | 
**Value** | Pointer to **map[string]interface{}** | A JSON object containing custom content. | [optional] 

## Methods

### NewJsonContextVariable

`func NewJsonContextVariable(type_ string, ) *JsonContextVariable`

NewJsonContextVariable instantiates a new JsonContextVariable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJsonContextVariableWithDefaults

`func NewJsonContextVariableWithDefaults() *JsonContextVariable`

NewJsonContextVariableWithDefaults instantiates a new JsonContextVariable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *JsonContextVariable) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JsonContextVariable) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JsonContextVariable) SetType(v string)`

SetType sets Type field to given value.


### GetValue

`func (o *JsonContextVariable) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *JsonContextVariable) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *JsonContextVariable) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *JsonContextVariable) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


