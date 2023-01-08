# FromLayoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FromPort** | Pointer to **int32** | The port that the transition can be made from | [optional] 
**Status** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**ToPortOverride** | Pointer to **int32** | The port that the transition goes to | [optional] 

## Methods

### NewFromLayoutPayload

`func NewFromLayoutPayload() *FromLayoutPayload`

NewFromLayoutPayload instantiates a new FromLayoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFromLayoutPayloadWithDefaults

`func NewFromLayoutPayloadWithDefaults() *FromLayoutPayload`

NewFromLayoutPayloadWithDefaults instantiates a new FromLayoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFromPort

`func (o *FromLayoutPayload) GetFromPort() int32`

GetFromPort returns the FromPort field if non-nil, zero value otherwise.

### GetFromPortOk

`func (o *FromLayoutPayload) GetFromPortOk() (*int32, bool)`

GetFromPortOk returns a tuple with the FromPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromPort

`func (o *FromLayoutPayload) SetFromPort(v int32)`

SetFromPort sets FromPort field to given value.

### HasFromPort

`func (o *FromLayoutPayload) HasFromPort() bool`

HasFromPort returns a boolean if a field has been set.

### GetStatus

`func (o *FromLayoutPayload) GetStatus() ProjectCreateResourceIdentifier`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FromLayoutPayload) GetStatusOk() (*ProjectCreateResourceIdentifier, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FromLayoutPayload) SetStatus(v ProjectCreateResourceIdentifier)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FromLayoutPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetToPortOverride

`func (o *FromLayoutPayload) GetToPortOverride() int32`

GetToPortOverride returns the ToPortOverride field if non-nil, zero value otherwise.

### GetToPortOverrideOk

`func (o *FromLayoutPayload) GetToPortOverrideOk() (*int32, bool)`

GetToPortOverrideOk returns a tuple with the ToPortOverride field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToPortOverride

`func (o *FromLayoutPayload) SetToPortOverride(v int32)`

SetToPortOverride sets ToPortOverride field to given value.

### HasToPortOverride

`func (o *FromLayoutPayload) HasToPortOverride() bool`

HasToPortOverride returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


