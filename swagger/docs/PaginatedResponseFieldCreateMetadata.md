# PaginatedResponseFieldCreateMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]FieldCreateMetadata**](FieldCreateMetadata.md) |  | [optional] 
**StartAt** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewPaginatedResponseFieldCreateMetadata

`func NewPaginatedResponseFieldCreateMetadata() *PaginatedResponseFieldCreateMetadata`

NewPaginatedResponseFieldCreateMetadata instantiates a new PaginatedResponseFieldCreateMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseFieldCreateMetadataWithDefaults

`func NewPaginatedResponseFieldCreateMetadataWithDefaults() *PaginatedResponseFieldCreateMetadata`

NewPaginatedResponseFieldCreateMetadataWithDefaults instantiates a new PaginatedResponseFieldCreateMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *PaginatedResponseFieldCreateMetadata) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PaginatedResponseFieldCreateMetadata) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PaginatedResponseFieldCreateMetadata) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PaginatedResponseFieldCreateMetadata) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseFieldCreateMetadata) GetResults() []FieldCreateMetadata`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseFieldCreateMetadata) GetResultsOk() (*[]FieldCreateMetadata, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseFieldCreateMetadata) SetResults(v []FieldCreateMetadata)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseFieldCreateMetadata) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PaginatedResponseFieldCreateMetadata) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PaginatedResponseFieldCreateMetadata) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PaginatedResponseFieldCreateMetadata) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PaginatedResponseFieldCreateMetadata) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PaginatedResponseFieldCreateMetadata) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedResponseFieldCreateMetadata) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedResponseFieldCreateMetadata) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginatedResponseFieldCreateMetadata) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


