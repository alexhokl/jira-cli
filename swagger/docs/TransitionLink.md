# TransitionLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | Pointer to **int32** | The from port number. | [optional] 
**FromStatusReference** | Pointer to **string** | The from status reference. | [optional] 
**ToPort** | Pointer to **int32** | The to port number. | [optional] 

## Methods

### NewTransitionLink

`func NewTransitionLink() *TransitionLink`

NewTransitionLink instantiates a new TransitionLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransitionLinkWithDefaults

`func NewTransitionLinkWithDefaults() *TransitionLink`

NewTransitionLinkWithDefaults instantiates a new TransitionLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *TransitionLink) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *TransitionLink) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *TransitionLink) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *TransitionLink) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### GetFromStatusReference

`func (o *TransitionLink) GetFromStatusReference() string`

GetFromStatusReference returns the FromStatusReference field if non-nil, zero value otherwise.

### GetFromStatusReferenceOk

`func (o *TransitionLink) GetFromStatusReferenceOk() (*string, bool)`

GetFromStatusReferenceOk returns a tuple with the FromStatusReference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromStatusReference

`func (o *TransitionLink) SetFromStatusReference(v string)`

SetFromStatusReference sets FromStatusReference field to given value.

### HasFromStatusReference

`func (o *TransitionLink) HasFromStatusReference() bool`

HasFromStatusReference returns a boolean if a field has been set.

### GetToPort

`func (o *TransitionLink) GetToPort() int32`

GetToPort returns the ToPort field if non-nil, zero value otherwise.

### GetToPortOk

`func (o *TransitionLink) GetToPortOk() (*int32, bool)`

GetToPortOk returns a tuple with the ToPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPort

`func (o *TransitionLink) SetToPort(v int32)`

SetToPort sets ToPort field to given value.

### HasToPort

`func (o *TransitionLink) HasToPort() bool`

HasToPort returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


