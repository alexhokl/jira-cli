# BoardAdminsBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Groups** | Pointer to [**[]GetAllBoards200ResponseValuesInnerAdminsAllOfGroupsInner**](GetAllBoards200ResponseValuesInnerAdminsAllOfGroupsInner.md) |  | [optional] 
**Users** | Pointer to [**[]GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner**](GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner.md) |  | [optional] 

## Methods

### NewBoardAdminsBean

`func NewBoardAdminsBean() *BoardAdminsBean`

NewBoardAdminsBean instantiates a new BoardAdminsBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardAdminsBeanWithDefaults

`func NewBoardAdminsBeanWithDefaults() *BoardAdminsBean`

NewBoardAdminsBeanWithDefaults instantiates a new BoardAdminsBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroups

`func (o *BoardAdminsBean) GetGroups() []GetAllBoards200ResponseValuesInnerAdminsAllOfGroupsInner`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *BoardAdminsBean) GetGroupsOk() (*[]GetAllBoards200ResponseValuesInnerAdminsAllOfGroupsInner, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *BoardAdminsBean) SetGroups(v []GetAllBoards200ResponseValuesInnerAdminsAllOfGroupsInner)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *BoardAdminsBean) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetUsers

`func (o *BoardAdminsBean) GetUsers() []GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *BoardAdminsBean) GetUsersOk() (*[]GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *BoardAdminsBean) SetUsers(v []GetAllBoards200ResponseValuesInnerAdminsAllOfUsersInner)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *BoardAdminsBean) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


