# Board

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admins** | Pointer to [**GetAllBoards200ResponseValuesInnerAdmins**](GetAllBoards200ResponseValuesInnerAdmins.md) |  | [optional] 
**CanEdit** | Pointer to **bool** | Whether the board can be edited. | [optional] [readonly] 
**Favourite** | Pointer to **bool** | Whether the board is selected as a favorite. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the board. | [optional] 
**IsPrivate** | Pointer to **bool** | Whether the board is private. | [optional] [readonly] 
**Location** | Pointer to [**GetAllBoards200ResponseValuesInnerLocation**](GetAllBoards200ResponseValuesInnerLocation.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the board. | [optional] 
**Self** | Pointer to **string** | The URL of the board. | [optional] [readonly] 
**Type** | Pointer to **string** | The type the board. | [optional] 

## Methods

### NewBoard

`func NewBoard() *Board`

NewBoard instantiates a new Board object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardWithDefaults

`func NewBoardWithDefaults() *Board`

NewBoardWithDefaults instantiates a new Board object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmins

`func (o *Board) GetAdmins() GetAllBoards200ResponseValuesInnerAdmins`

GetAdmins returns the Admins field if non-nil, zero value otherwise.

### GetAdminsOk

`func (o *Board) GetAdminsOk() (*GetAllBoards200ResponseValuesInnerAdmins, bool)`

GetAdminsOk returns a tuple with the Admins field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmins

`func (o *Board) SetAdmins(v GetAllBoards200ResponseValuesInnerAdmins)`

SetAdmins sets Admins field to given value.

### HasAdmins

`func (o *Board) HasAdmins() bool`

HasAdmins returns a boolean if a field has been set.

### GetCanEdit

`func (o *Board) GetCanEdit() bool`

GetCanEdit returns the CanEdit field if non-nil, zero value otherwise.

### GetCanEditOk

`func (o *Board) GetCanEditOk() (*bool, bool)`

GetCanEditOk returns a tuple with the CanEdit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanEdit

`func (o *Board) SetCanEdit(v bool)`

SetCanEdit sets CanEdit field to given value.

### HasCanEdit

`func (o *Board) HasCanEdit() bool`

HasCanEdit returns a boolean if a field has been set.

### GetFavourite

`func (o *Board) GetFavourite() bool`

GetFavourite returns the Favourite field if non-nil, zero value otherwise.

### GetFavouriteOk

`func (o *Board) GetFavouriteOk() (*bool, bool)`

GetFavouriteOk returns a tuple with the Favourite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavourite

`func (o *Board) SetFavourite(v bool)`

SetFavourite sets Favourite field to given value.

### HasFavourite

`func (o *Board) HasFavourite() bool`

HasFavourite returns a boolean if a field has been set.

### GetId

`func (o *Board) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Board) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Board) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Board) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsPrivate

`func (o *Board) GetIsPrivate() bool`

GetIsPrivate returns the IsPrivate field if non-nil, zero value otherwise.

### GetIsPrivateOk

`func (o *Board) GetIsPrivateOk() (*bool, bool)`

GetIsPrivateOk returns a tuple with the IsPrivate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPrivate

`func (o *Board) SetIsPrivate(v bool)`

SetIsPrivate sets IsPrivate field to given value.

### HasIsPrivate

`func (o *Board) HasIsPrivate() bool`

HasIsPrivate returns a boolean if a field has been set.

### GetLocation

`func (o *Board) GetLocation() GetAllBoards200ResponseValuesInnerLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *Board) GetLocationOk() (*GetAllBoards200ResponseValuesInnerLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *Board) SetLocation(v GetAllBoards200ResponseValuesInnerLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *Board) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *Board) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Board) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Board) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Board) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *Board) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *Board) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *Board) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *Board) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetType

`func (o *Board) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Board) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Board) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Board) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


