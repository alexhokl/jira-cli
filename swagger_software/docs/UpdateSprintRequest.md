# UpdateSprintRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompleteDate** | Pointer to **string** |  | [optional] 
**CreatedDate** | Pointer to **string** |  | [optional] 
**EndDate** | Pointer to **string** |  | [optional] 
**Goal** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OriginBoardId** | Pointer to **int64** |  | [optional] 
**Self** | Pointer to **string** |  | [optional] 
**StartDate** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 

## Methods

### NewUpdateSprintRequest

`func NewUpdateSprintRequest() *UpdateSprintRequest`

NewUpdateSprintRequest instantiates a new UpdateSprintRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSprintRequestWithDefaults

`func NewUpdateSprintRequestWithDefaults() *UpdateSprintRequest`

NewUpdateSprintRequestWithDefaults instantiates a new UpdateSprintRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompleteDate

`func (o *UpdateSprintRequest) GetCompleteDate() string`

GetCompleteDate returns the CompleteDate field if non-nil, zero value otherwise.

### GetCompleteDateOk

`func (o *UpdateSprintRequest) GetCompleteDateOk() (*string, bool)`

GetCompleteDateOk returns a tuple with the CompleteDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompleteDate

`func (o *UpdateSprintRequest) SetCompleteDate(v string)`

SetCompleteDate sets CompleteDate field to given value.

### HasCompleteDate

`func (o *UpdateSprintRequest) HasCompleteDate() bool`

HasCompleteDate returns a boolean if a field has been set.

### GetCreatedDate

`func (o *UpdateSprintRequest) GetCreatedDate() string`

GetCreatedDate returns the CreatedDate field if non-nil, zero value otherwise.

### GetCreatedDateOk

`func (o *UpdateSprintRequest) GetCreatedDateOk() (*string, bool)`

GetCreatedDateOk returns a tuple with the CreatedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedDate

`func (o *UpdateSprintRequest) SetCreatedDate(v string)`

SetCreatedDate sets CreatedDate field to given value.

### HasCreatedDate

`func (o *UpdateSprintRequest) HasCreatedDate() bool`

HasCreatedDate returns a boolean if a field has been set.

### GetEndDate

`func (o *UpdateSprintRequest) GetEndDate() string`

GetEndDate returns the EndDate field if non-nil, zero value otherwise.

### GetEndDateOk

`func (o *UpdateSprintRequest) GetEndDateOk() (*string, bool)`

GetEndDateOk returns a tuple with the EndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndDate

`func (o *UpdateSprintRequest) SetEndDate(v string)`

SetEndDate sets EndDate field to given value.

### HasEndDate

`func (o *UpdateSprintRequest) HasEndDate() bool`

HasEndDate returns a boolean if a field has been set.

### GetGoal

`func (o *UpdateSprintRequest) GetGoal() string`

GetGoal returns the Goal field if non-nil, zero value otherwise.

### GetGoalOk

`func (o *UpdateSprintRequest) GetGoalOk() (*string, bool)`

GetGoalOk returns a tuple with the Goal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGoal

`func (o *UpdateSprintRequest) SetGoal(v string)`

SetGoal sets Goal field to given value.

### HasGoal

`func (o *UpdateSprintRequest) HasGoal() bool`

HasGoal returns a boolean if a field has been set.

### GetId

`func (o *UpdateSprintRequest) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSprintRequest) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSprintRequest) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSprintRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateSprintRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSprintRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSprintRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSprintRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOriginBoardId

`func (o *UpdateSprintRequest) GetOriginBoardId() int64`

GetOriginBoardId returns the OriginBoardId field if non-nil, zero value otherwise.

### GetOriginBoardIdOk

`func (o *UpdateSprintRequest) GetOriginBoardIdOk() (*int64, bool)`

GetOriginBoardIdOk returns a tuple with the OriginBoardId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginBoardId

`func (o *UpdateSprintRequest) SetOriginBoardId(v int64)`

SetOriginBoardId sets OriginBoardId field to given value.

### HasOriginBoardId

`func (o *UpdateSprintRequest) HasOriginBoardId() bool`

HasOriginBoardId returns a boolean if a field has been set.

### GetSelf

`func (o *UpdateSprintRequest) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *UpdateSprintRequest) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *UpdateSprintRequest) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *UpdateSprintRequest) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetStartDate

`func (o *UpdateSprintRequest) GetStartDate() string`

GetStartDate returns the StartDate field if non-nil, zero value otherwise.

### GetStartDateOk

`func (o *UpdateSprintRequest) GetStartDateOk() (*string, bool)`

GetStartDateOk returns a tuple with the StartDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartDate

`func (o *UpdateSprintRequest) SetStartDate(v string)`

SetStartDate sets StartDate field to given value.

### HasStartDate

`func (o *UpdateSprintRequest) HasStartDate() bool`

HasStartDate returns a boolean if a field has been set.

### GetState

`func (o *UpdateSprintRequest) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UpdateSprintRequest) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UpdateSprintRequest) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *UpdateSprintRequest) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


