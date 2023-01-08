# PaginatedResponseComment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxResults** | Pointer to **int32** |  | [optional] 
**Results** | Pointer to [**[]Comment**](Comment.md) |  | [optional] 
**StartAt** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 

## Methods

### NewPaginatedResponseComment

`func NewPaginatedResponseComment() *PaginatedResponseComment`

NewPaginatedResponseComment instantiates a new PaginatedResponseComment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaginatedResponseCommentWithDefaults

`func NewPaginatedResponseCommentWithDefaults() *PaginatedResponseComment`

NewPaginatedResponseCommentWithDefaults instantiates a new PaginatedResponseComment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxResults

`func (o *PaginatedResponseComment) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PaginatedResponseComment) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PaginatedResponseComment) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PaginatedResponseComment) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetResults

`func (o *PaginatedResponseComment) GetResults() []Comment`

GetResults returns the Results field if non-nil, zero value otherwise.

### GetResultsOk

`func (o *PaginatedResponseComment) GetResultsOk() (*[]Comment, bool)`

GetResultsOk returns a tuple with the Results field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResults

`func (o *PaginatedResponseComment) SetResults(v []Comment)`

SetResults sets Results field to given value.

### HasResults

`func (o *PaginatedResponseComment) HasResults() bool`

HasResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PaginatedResponseComment) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PaginatedResponseComment) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PaginatedResponseComment) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PaginatedResponseComment) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PaginatedResponseComment) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PaginatedResponseComment) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PaginatedResponseComment) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PaginatedResponseComment) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


