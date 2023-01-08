# IdOrKeyBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the referenced item. | [optional] 
**Key** | Pointer to **string** | The key of the referenced item. | [optional] 

## Methods

### NewIdOrKeyBean

`func NewIdOrKeyBean() *IdOrKeyBean`

NewIdOrKeyBean instantiates a new IdOrKeyBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIdOrKeyBeanWithDefaults

`func NewIdOrKeyBeanWithDefaults() *IdOrKeyBean`

NewIdOrKeyBeanWithDefaults instantiates a new IdOrKeyBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IdOrKeyBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IdOrKeyBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IdOrKeyBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IdOrKeyBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *IdOrKeyBean) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *IdOrKeyBean) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *IdOrKeyBean) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *IdOrKeyBean) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


