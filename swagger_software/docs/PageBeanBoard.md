# PageBeanBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** |  | [optional] 
**MaxResults** | Pointer to **int32** |  | [optional] 
**StartAt** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Values** | Pointer to [**[]GetAllBoards200ResponseValuesInner**](GetAllBoards200ResponseValuesInner.md) |  | [optional] 

## Methods

### NewPageBeanBoard

`func NewPageBeanBoard() *PageBeanBoard`

NewPageBeanBoard instantiates a new PageBeanBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageBeanBoardWithDefaults

`func NewPageBeanBoardWithDefaults() *PageBeanBoard`

NewPageBeanBoardWithDefaults instantiates a new PageBeanBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *PageBeanBoard) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *PageBeanBoard) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *PageBeanBoard) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *PageBeanBoard) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *PageBeanBoard) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *PageBeanBoard) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *PageBeanBoard) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *PageBeanBoard) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *PageBeanBoard) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *PageBeanBoard) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *PageBeanBoard) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *PageBeanBoard) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *PageBeanBoard) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *PageBeanBoard) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *PageBeanBoard) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *PageBeanBoard) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *PageBeanBoard) GetValues() []GetAllBoards200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *PageBeanBoard) GetValuesOk() (*[]GetAllBoards200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *PageBeanBoard) SetValues(v []GetAllBoards200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *PageBeanBoard) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


