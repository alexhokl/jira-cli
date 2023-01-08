# JqlQueryOrderByClauseElement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Direction** | Pointer to **string** | The direction in which to order the results. | [optional] 
**Field** | [**JqlQueryField**](JqlQueryField.md) |  | 

## Methods

### NewJqlQueryOrderByClauseElement

`func NewJqlQueryOrderByClauseElement(field JqlQueryField, ) *JqlQueryOrderByClauseElement`

NewJqlQueryOrderByClauseElement instantiates a new JqlQueryOrderByClauseElement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryOrderByClauseElementWithDefaults

`func NewJqlQueryOrderByClauseElementWithDefaults() *JqlQueryOrderByClauseElement`

NewJqlQueryOrderByClauseElementWithDefaults instantiates a new JqlQueryOrderByClauseElement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDirection

`func (o *JqlQueryOrderByClauseElement) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *JqlQueryOrderByClauseElement) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *JqlQueryOrderByClauseElement) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *JqlQueryOrderByClauseElement) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetField

`func (o *JqlQueryOrderByClauseElement) GetField() JqlQueryField`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *JqlQueryOrderByClauseElement) GetFieldOk() (*JqlQueryField, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *JqlQueryOrderByClauseElement) SetField(v JqlQueryField)`

SetField sets Field field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


