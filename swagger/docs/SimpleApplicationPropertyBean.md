# SimpleApplicationPropertyBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the application property. | [optional] 
**Value** | Pointer to **string** | The new value. | [optional] 

## Methods

### NewSimpleApplicationPropertyBean

`func NewSimpleApplicationPropertyBean() *SimpleApplicationPropertyBean`

NewSimpleApplicationPropertyBean instantiates a new SimpleApplicationPropertyBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSimpleApplicationPropertyBeanWithDefaults

`func NewSimpleApplicationPropertyBeanWithDefaults() *SimpleApplicationPropertyBean`

NewSimpleApplicationPropertyBeanWithDefaults instantiates a new SimpleApplicationPropertyBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SimpleApplicationPropertyBean) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SimpleApplicationPropertyBean) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SimpleApplicationPropertyBean) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SimpleApplicationPropertyBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetValue

`func (o *SimpleApplicationPropertyBean) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SimpleApplicationPropertyBean) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SimpleApplicationPropertyBean) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *SimpleApplicationPropertyBean) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


