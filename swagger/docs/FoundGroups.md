# FoundGroups

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]FoundGroup**](FoundGroup.md) |  | [optional] 
**Header** | Pointer to **string** | Header text indicating the number of groups in the response and the total number of groups found in the search. | [optional] 
**Total** | Pointer to **int32** | The total number of groups found in the search. | [optional] 

## Methods

### NewFoundGroups

`func NewFoundGroups() *FoundGroups`

NewFoundGroups instantiates a new FoundGroups object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFoundGroupsWithDefaults

`func NewFoundGroupsWithDefaults() *FoundGroups`

NewFoundGroupsWithDefaults instantiates a new FoundGroups object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *FoundGroups) GetGroups() []FoundGroup`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *FoundGroups) GetGroupsOk() (*[]FoundGroup, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *FoundGroups) SetGroups(v []FoundGroup)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *FoundGroups) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetHeader

`func (o *FoundGroups) GetHeader() string`

GetHeader returns the Header field if non-nil, zero value otherwise.

### GetHeaderOk

`func (o *FoundGroups) GetHeaderOk() (*string, bool)`

GetHeaderOk returns a tuple with the Header field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeader

`func (o *FoundGroups) SetHeader(v string)`

SetHeader sets Header field to given value.

### HasHeader

`func (o *FoundGroups) HasHeader() bool`

HasHeader returns a boolean if a field has been set.

### GetTotal

`func (o *FoundGroups) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *FoundGroups) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *FoundGroups) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *FoundGroups) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


