# ToLayoutPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Port** | Pointer to **int32** | Defines where the transition line will be connected to a status. Port 0 to 7 are acceptable values. | [optional] 
**Status** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 

## Methods

### NewToLayoutPayload

`func NewToLayoutPayload() *ToLayoutPayload`

NewToLayoutPayload instantiates a new ToLayoutPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewToLayoutPayloadWithDefaults

`func NewToLayoutPayloadWithDefaults() *ToLayoutPayload`

NewToLayoutPayloadWithDefaults instantiates a new ToLayoutPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPort

`func (o *ToLayoutPayload) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ToLayoutPayload) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ToLayoutPayload) SetPort(v int32)`

SetPort sets Port field to given value.

### HasPort

`func (o *ToLayoutPayload) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetStatus

`func (o *ToLayoutPayload) GetStatus() ProjectCreateResourceIdentifier`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ToLayoutPayload) GetStatusOk() (*ProjectCreateResourceIdentifier, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ToLayoutPayload) SetStatus(v ProjectCreateResourceIdentifier)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ToLayoutPayload) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


