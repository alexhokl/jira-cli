# TargetStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | **map[string][]string** | An object with the key as the ID of the target status and value with the list of the IDs of the current source statuses. | 

## Methods

### NewTargetStatus

`func NewTargetStatus(statuses map[string][]string, ) *TargetStatus`

NewTargetStatus instantiates a new TargetStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetStatusWithDefaults

`func NewTargetStatusWithDefaults() *TargetStatus`

NewTargetStatusWithDefaults instantiates a new TargetStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *TargetStatus) GetStatuses() map[string][]string`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *TargetStatus) GetStatusesOk() (*map[string][]string, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *TargetStatus) SetStatuses(v map[string][]string)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


