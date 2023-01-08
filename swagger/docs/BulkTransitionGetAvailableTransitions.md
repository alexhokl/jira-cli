# BulkTransitionGetAvailableTransitions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvailableTransitions** | Pointer to [**[]IssueBulkTransitionForWorkflow**](IssueBulkTransitionForWorkflow.md) | List of available transitions for bulk transition operation for requested issues grouped by workflow | [optional] [readonly] 
**EndingBefore** | Pointer to **string** | The end cursor for use in pagination. | [optional] [readonly] 
**StartingAfter** | Pointer to **string** | The start cursor for use in pagination. | [optional] [readonly] 

## Methods

### NewBulkTransitionGetAvailableTransitions

`func NewBulkTransitionGetAvailableTransitions() *BulkTransitionGetAvailableTransitions`

NewBulkTransitionGetAvailableTransitions instantiates a new BulkTransitionGetAvailableTransitions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkTransitionGetAvailableTransitionsWithDefaults

`func NewBulkTransitionGetAvailableTransitionsWithDefaults() *BulkTransitionGetAvailableTransitions`

NewBulkTransitionGetAvailableTransitionsWithDefaults instantiates a new BulkTransitionGetAvailableTransitions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailableTransitions

`func (o *BulkTransitionGetAvailableTransitions) GetAvailableTransitions() []IssueBulkTransitionForWorkflow`

GetAvailableTransitions returns the AvailableTransitions field if non-nil, zero value otherwise.

### GetAvailableTransitionsOk

`func (o *BulkTransitionGetAvailableTransitions) GetAvailableTransitionsOk() (*[]IssueBulkTransitionForWorkflow, bool)`

GetAvailableTransitionsOk returns a tuple with the AvailableTransitions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableTransitions

`func (o *BulkTransitionGetAvailableTransitions) SetAvailableTransitions(v []IssueBulkTransitionForWorkflow)`

SetAvailableTransitions sets AvailableTransitions field to given value.

### HasAvailableTransitions

`func (o *BulkTransitionGetAvailableTransitions) HasAvailableTransitions() bool`

HasAvailableTransitions returns a boolean if a field has been set.

### GetEndingBefore

`func (o *BulkTransitionGetAvailableTransitions) GetEndingBefore() string`

GetEndingBefore returns the EndingBefore field if non-nil, zero value otherwise.

### GetEndingBeforeOk

`func (o *BulkTransitionGetAvailableTransitions) GetEndingBeforeOk() (*string, bool)`

GetEndingBeforeOk returns a tuple with the EndingBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndingBefore

`func (o *BulkTransitionGetAvailableTransitions) SetEndingBefore(v string)`

SetEndingBefore sets EndingBefore field to given value.

### HasEndingBefore

`func (o *BulkTransitionGetAvailableTransitions) HasEndingBefore() bool`

HasEndingBefore returns a boolean if a field has been set.

### GetStartingAfter

`func (o *BulkTransitionGetAvailableTransitions) GetStartingAfter() string`

GetStartingAfter returns the StartingAfter field if non-nil, zero value otherwise.

### GetStartingAfterOk

`func (o *BulkTransitionGetAvailableTransitions) GetStartingAfterOk() (*string, bool)`

GetStartingAfterOk returns a tuple with the StartingAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartingAfter

`func (o *BulkTransitionGetAvailableTransitions) SetStartingAfter(v string)`

SetStartingAfter sets StartingAfter field to given value.

### HasStartingAfter

`func (o *BulkTransitionGetAvailableTransitions) HasStartingAfter() bool`

HasStartingAfter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


