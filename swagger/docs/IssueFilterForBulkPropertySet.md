# IssueFilterForBulkPropertySet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentValue** | Pointer to **interface{}** | The value of properties to perform the bulk operation on. | [optional] 
**EntityIds** | Pointer to **[]int64** | List of issues to perform the bulk operation on. | [optional] 
**HasProperty** | Pointer to **bool** | Whether the bulk operation occurs only when the property is present on or absent from an issue. | [optional] 

## Methods

### NewIssueFilterForBulkPropertySet

`func NewIssueFilterForBulkPropertySet() *IssueFilterForBulkPropertySet`

NewIssueFilterForBulkPropertySet instantiates a new IssueFilterForBulkPropertySet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueFilterForBulkPropertySetWithDefaults

`func NewIssueFilterForBulkPropertySetWithDefaults() *IssueFilterForBulkPropertySet`

NewIssueFilterForBulkPropertySetWithDefaults instantiates a new IssueFilterForBulkPropertySet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentValue

`func (o *IssueFilterForBulkPropertySet) GetCurrentValue() interface{}`

GetCurrentValue returns the CurrentValue field if non-nil, zero value otherwise.

### GetCurrentValueOk

`func (o *IssueFilterForBulkPropertySet) GetCurrentValueOk() (*interface{}, bool)`

GetCurrentValueOk returns a tuple with the CurrentValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentValue

`func (o *IssueFilterForBulkPropertySet) SetCurrentValue(v interface{})`

SetCurrentValue sets CurrentValue field to given value.

### HasCurrentValue

`func (o *IssueFilterForBulkPropertySet) HasCurrentValue() bool`

HasCurrentValue returns a boolean if a field has been set.

### SetCurrentValueNil

`func (o *IssueFilterForBulkPropertySet) SetCurrentValueNil(b bool)`

 SetCurrentValueNil sets the value for CurrentValue to be an explicit nil

### UnsetCurrentValue
`func (o *IssueFilterForBulkPropertySet) UnsetCurrentValue()`

UnsetCurrentValue ensures that no value is present for CurrentValue, not even an explicit nil
### GetEntityIds

`func (o *IssueFilterForBulkPropertySet) GetEntityIds() []int64`

GetEntityIds returns the EntityIds field if non-nil, zero value otherwise.

### GetEntityIdsOk

`func (o *IssueFilterForBulkPropertySet) GetEntityIdsOk() (*[]int64, bool)`

GetEntityIdsOk returns a tuple with the EntityIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityIds

`func (o *IssueFilterForBulkPropertySet) SetEntityIds(v []int64)`

SetEntityIds sets EntityIds field to given value.

### HasEntityIds

`func (o *IssueFilterForBulkPropertySet) HasEntityIds() bool`

HasEntityIds returns a boolean if a field has been set.

### GetHasProperty

`func (o *IssueFilterForBulkPropertySet) GetHasProperty() bool`

GetHasProperty returns the HasProperty field if non-nil, zero value otherwise.

### GetHasPropertyOk

`func (o *IssueFilterForBulkPropertySet) GetHasPropertyOk() (*bool, bool)`

GetHasPropertyOk returns a tuple with the HasProperty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProperty

`func (o *IssueFilterForBulkPropertySet) SetHasProperty(v bool)`

SetHasProperty sets HasProperty field to given value.

### HasHasProperty

`func (o *IssueFilterForBulkPropertySet) HasHasProperty() bool`

HasHasProperty returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


