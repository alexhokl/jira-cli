# IssueFilterForBulkPropertyDelete

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **interface{}** | The value of properties to perform the bulk operation on. | [optional] 
**EntityIds** | Pointer to **[]int64** | List of issues to perform the bulk delete operation on. | [optional] 

## Methods

### NewIssueFilterForBulkPropertyDelete

`func NewIssueFilterForBulkPropertyDelete() *IssueFilterForBulkPropertyDelete`

NewIssueFilterForBulkPropertyDelete instantiates a new IssueFilterForBulkPropertyDelete object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFilterForBulkPropertyDeleteWithDefaults

`func NewIssueFilterForBulkPropertyDeleteWithDefaults() *IssueFilterForBulkPropertyDelete`

NewIssueFilterForBulkPropertyDeleteWithDefaults instantiates a new IssueFilterForBulkPropertyDelete object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *IssueFilterForBulkPropertyDelete) GetCurrentValue() interface{}`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *IssueFilterForBulkPropertyDelete) GetCurrentValueOk() (*interface{}, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *IssueFilterForBulkPropertyDelete) SetCurrentValue(v interface{})`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *IssueFilterForBulkPropertyDelete) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *IssueFilterForBulkPropertyDelete) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *IssueFilterForBulkPropertyDelete) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetEntityIds

`func (o *IssueFilterForBulkPropertyDelete) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *IssueFilterForBulkPropertyDelete) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *IssueFilterForBulkPropertyDelete) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *IssueFilterForBulkPropertyDelete) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


