# PropertyKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | Pointer to **string** | The key of the property. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the property. | [optional] [readonly] 

## Methods

### NewPropertyKey

`func NewPropertyKey() *PropertyKey`

NewPropertyKey instantiates a new PropertyKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPropertyKeyWithDefaults

`func NewPropertyKeyWithDefaults() *PropertyKey`

NewPropertyKeyWithDefaults instantiates a new PropertyKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *PropertyKey) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PropertyKey) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PropertyKey) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *PropertyKey) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetSelf

`func (o *PropertyKey) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *PropertyKey) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *PropertyKey) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *PropertyKey) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


