# PriorityMapping

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**In** | Pointer to **map[string]int64** | The mapping of priorities for issues being migrated **into** this priority scheme. Key is the old priority ID, value is the new priority ID (must exist in this priority scheme).  E.g. The current priority scheme has priority ID &#x60;10001&#x60;. Issues with priority ID &#x60;10000&#x60; are being migrated into this priority scheme will need mapping to new priorities. The &#x60;in&#x60; mapping would be &#x60;{\&quot;10000\&quot;: 10001}&#x60;. | [optional] 
**Out** | Pointer to **map[string]int64** | The mapping of priorities for issues being migrated **out of** this priority scheme. Key is the old priority ID (must exist in this priority scheme), value is the new priority ID (must exist in the default priority scheme). Required for updating an existing priority scheme. Not used when creating a new priority scheme.  E.g. The current priority scheme has priority ID &#x60;10001&#x60;. Issues with priority ID &#x60;10001&#x60; are being migrated out of this priority scheme will need mapping to new priorities. The &#x60;out&#x60; mapping would be &#x60;{\&quot;10001\&quot;: 10000}&#x60;. | [optional] 

## Methods

### NewPriorityMapping

`func NewPriorityMapping() *PriorityMapping`

NewPriorityMapping instantiates a new PriorityMapping object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriorityMappingWithDefaults

`func NewPriorityMappingWithDefaults() *PriorityMapping`

NewPriorityMappingWithDefaults instantiates a new PriorityMapping object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIn

`func (o *PriorityMapping) GetIn() map[string]int64`

GetIn returns the In field if non-nil, zero value otherwise.

### GetInOk

`func (o *PriorityMapping) GetInOk() (*map[string]int64, bool)`

GetInOk returns a tuple with the In field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIn

`func (o *PriorityMapping) SetIn(v map[string]int64)`

SetIn sets In field to given value.

### HasIn

`func (o *PriorityMapping) HasIn() bool`

HasIn returns a boolean if a field has been set.

### GetOut

`func (o *PriorityMapping) GetOut() map[string]int64`

GetOut returns the Out field if non-nil, zero value otherwise.

### GetOutOk

`func (o *PriorityMapping) GetOutOk() (*map[string]int64, bool)`

GetOutOk returns a tuple with the Out field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOut

`func (o *PriorityMapping) SetOut(v map[string]int64)`

SetOut sets Out field to given value.

### HasOut

`func (o *PriorityMapping) HasOut() bool`

HasOut returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


