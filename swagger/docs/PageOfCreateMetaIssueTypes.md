# PageOfCreateMetaIssueTypes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreateMetaIssueType** | Pointer to [**[]IssueTypeIssueCreateMetadata**](IssueTypeIssueCreateMetadata.md) |  | [optional] 
**IssueTypes** | Pointer to [**[]IssueTypeIssueCreateMetadata**](IssueTypeIssueCreateMetadata.md) | The list of CreateMetaIssueType. | [optional] [readonly] 
**MaxResults** | Pointer to **int32** | The maximum number of items to return per page. | [optional] [readonly] 
**StartAt** | Pointer to **int64** | The index of the first item returned. | [optional] [readonly] 
**Total** | Pointer to **int64** | The total number of items in all pages. | [optional] [readonly] 

## Methods

### NewPageOfCreateMetaIssueTypes

`func NewPageOfCreateMetaIssueTypes() *PageOfCreateMetaIssueTypes`

NewPageOfCreateMetaIssueTypes instantiates a new PageOfCreateMetaIssueTypes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageOfCreateMetaIssueTypesWithDefaults

`func NewPageOfCreateMetaIssueTypesWithDefaults() *PageOfCreateMetaIssueTypes`

NewPageOfCreateMetaIssueTypesWithDefaults instantiates a new PageOfCreateMetaIssueTypes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreateMetaIssueType

`func (o *PageOfCreateMetaIssueTypes) GetCreateMetaIssueType() []IssueTypeIssueCreateMetadata`

GetCreateMetaIssueType returns the CreateMetaIssueType field if non-nil, zero value otherwise.

### GetCreateMetaIssueTypeOk

`func (o *PageOfCreateMetaIssueTypes) GetCreateMetaIssueTypeOk() (*[]IssueTypeIssueCreateMetadata, bool)`

GetCreateMetaIssueTypeOk returns a tuple with the CreateMetaIssueType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateMetaIssueType

`func (o *PageOfCreateMetaIssueTypes) SetCreateMetaIssueType(v []IssueTypeIssueCreateMetadata)`

SetCreateMetaIssueType sets CreateMetaIssueType field to given value.

### HasCreateMetaIssueType

`func (o *PageOfCreateMetaIssueTypes) HasCreateMetaIssueType() bool`

HasCreateMetaIssueType returns a boolean if a field has been set.

### GetIssueTypes

`func (o *PageOfCreateMetaIssueTypes) GetIssueTypes() []IssueTypeIssueCreateMetadata`

GetIssueTypes returns the IssueTypes field if non-nil, zero value otherwise.

### GetIssueTypesOk

`func (o *PageOfCreateMetaIssueTypes) GetIssueTypesOk() (*[]IssueTypeIssueCreateMetadata, bool)`

GetIssueTypesOk returns a tuple with the IssueTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTypes

`func (o *PageOfCreateMetaIssueTypes) SetIssueTypes(v []IssueTypeIssueCreateMetadata)`

SetIssueTypes sets IssueTypes field to given value.

### HasIssueTypes

`func (o *PageOfCreateMetaIssueTypes) HasIssueTypes() bool`

HasIssueTypes returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageOfCreateMetaIssueTypes) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageOfCreateMetaIssueTypes) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageOfCreateMetaIssueTypes) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageOfCreateMetaIssueTypes) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PageOfCreateMetaIssueTypes) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageOfCreateMetaIssueTypes) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageOfCreateMetaIssueTypes) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageOfCreateMetaIssueTypes) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageOfCreateMetaIssueTypes) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageOfCreateMetaIssueTypes) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageOfCreateMetaIssueTypes) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageOfCreateMetaIssueTypes) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


