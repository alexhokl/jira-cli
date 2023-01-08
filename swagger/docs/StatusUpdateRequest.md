# StatusUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Statuses** | [**[]StatusUpdate**](StatusUpdate.md) | The list of statuses that will be updated. | 

## Methods

### NewStatusUpdateRequest

`func NewStatusUpdateRequest(statuses []StatusUpdate, ) *StatusUpdateRequest`

NewStatusUpdateRequest instantiates a new StatusUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusUpdateRequestWithDefaults

`func NewStatusUpdateRequestWithDefaults() *StatusUpdateRequest`

NewStatusUpdateRequestWithDefaults instantiates a new StatusUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatuses

`func (o *StatusUpdateRequest) GetStatuses() []StatusUpdate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *StatusUpdateRequest) GetStatusesOk() (*[]StatusUpdate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *StatusUpdateRequest) SetStatuses(v []StatusUpdate)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


