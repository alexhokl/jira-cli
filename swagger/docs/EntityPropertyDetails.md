# EntityPropertyDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EntityId** | **float32** | The entity property ID. | 
**Key** | **string** | The entity property key. | 
**Value** | **string** | The new value of the entity property. | 

## Methods

### NewEntityPropertyDetails

`func NewEntityPropertyDetails(entityId float32, key string, value string, ) *EntityPropertyDetails`

NewEntityPropertyDetails instantiates a new EntityPropertyDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntityPropertyDetailsWithDefaults

`func NewEntityPropertyDetailsWithDefaults() *EntityPropertyDetails`

NewEntityPropertyDetailsWithDefaults instantiates a new EntityPropertyDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntityId

`func (o *EntityPropertyDetails) GetEntityId() float32`

GetEntityId returns the EntityId field if non-nil, zero value otherwise.

### GetEntityIdOk

`func (o *EntityPropertyDetails) GetEntityIdOk() (*float32, bool)`

GetEntityIdOk returns a tuple with the EntityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityId

`func (o *EntityPropertyDetails) SetEntityId(v float32)`

SetEntityId sets EntityId field to given value.


### GetKey

`func (o *EntityPropertyDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *EntityPropertyDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *EntityPropertyDetails) SetKey(v string)`

SetKey sets Key field to given value.


### GetValue

`func (o *EntityPropertyDetails) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *EntityPropertyDetails) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *EntityPropertyDetails) SetValue(v string)`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


