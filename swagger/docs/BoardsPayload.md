# BoardsPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Boards** | Pointer to [**[]BoardPayload**](BoardPayload.md) | The boards to be associated with the project. | [optional] 

## Methods

### NewBoardsPayload

`func NewBoardsPayload() *BoardsPayload`

NewBoardsPayload instantiates a new BoardsPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardsPayloadWithDefaults

`func NewBoardsPayloadWithDefaults() *BoardsPayload`

NewBoardsPayloadWithDefaults instantiates a new BoardsPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoards

`func (o *BoardsPayload) GetBoards() []BoardPayload`

GetBoards returns the Boards field if non-nil, zero value otherwise.

### GetBoardsOk

`func (o *BoardsPayload) GetBoardsOk() (*[]BoardPayload, bool)`

GetBoardsOk returns a tuple with the Boards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoards

`func (o *BoardsPayload) SetBoards(v []BoardPayload)`

SetBoards sets Boards field to given value.

### HasBoards

`func (o *BoardsPayload) HasBoards() bool`

HasBoards returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


