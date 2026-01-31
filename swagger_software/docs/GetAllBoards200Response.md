# GetAllBoards200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsLast** | Pointer to **bool** |  | [optional] 
**MaxResults** | Pointer to **int32** |  | [optional] 
**StartAt** | Pointer to **int64** |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Values** | Pointer to [**[]GetAllBoards200ResponseValuesInner**](GetAllBoards200ResponseValuesInner.md) |  | [optional] 

## Methods

### NewGetAllBoards200Response

`func NewGetAllBoards200Response() *GetAllBoards200Response`

NewGetAllBoards200Response instantiates a new GetAllBoards200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAllBoards200ResponseWithDefaults

`func NewGetAllBoards200ResponseWithDefaults() *GetAllBoards200Response`

NewGetAllBoards200ResponseWithDefaults instantiates a new GetAllBoards200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsLast

`func (o *GetAllBoards200Response) GetIsLast() bool`

GetIsLast returns the IsLast field if non-nil, zero value otherwise.

### GetIsLastOk

`func (o *GetAllBoards200Response) GetIsLastOk() (*bool, bool)`

GetIsLastOk returns a tuple with the IsLast field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLast

`func (o *GetAllBoards200Response) SetIsLast(v bool)`

SetIsLast sets IsLast field to given value.

### HasIsLast

`func (o *GetAllBoards200Response) HasIsLast() bool`

HasIsLast returns a boolean if a field has been set.

### GetMaxResults

`func (o *GetAllBoards200Response) GetMaxResults() int32`

GetMaxResults returns the MaxResults field if non-nil, zero value otherwise.

### GetMaxResultsOk

`func (o *GetAllBoards200Response) GetMaxResultsOk() (*int32, bool)`

GetMaxResultsOk returns a tuple with the MaxResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxResults

`func (o *GetAllBoards200Response) SetMaxResults(v int32)`

SetMaxResults sets MaxResults field to given value.

### HasMaxResults

`func (o *GetAllBoards200Response) HasMaxResults() bool`

HasMaxResults returns a boolean if a field has been set.

### GetStartAt

`func (o *GetAllBoards200Response) GetStartAt() int64`

GetStartAt returns the StartAt field if non-nil, zero value otherwise.

### GetStartAtOk

`func (o *GetAllBoards200Response) GetStartAtOk() (*int64, bool)`

GetStartAtOk returns a tuple with the StartAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartAt

`func (o *GetAllBoards200Response) SetStartAt(v int64)`

SetStartAt sets StartAt field to given value.

### HasStartAt

`func (o *GetAllBoards200Response) HasStartAt() bool`

HasStartAt returns a boolean if a field has been set.

### GetTotal

`func (o *GetAllBoards200Response) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetAllBoards200Response) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetAllBoards200Response) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetAllBoards200Response) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetValues

`func (o *GetAllBoards200Response) GetValues() []GetAllBoards200ResponseValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *GetAllBoards200Response) GetValuesOk() (*[]GetAllBoards200ResponseValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *GetAllBoards200Response) SetValues(v []GetAllBoards200ResponseValuesInner)`

SetValues sets Values field to given value.

### HasValues

`func (o *GetAllBoards200Response) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


