# RemoteIssueLink

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Application**](Application.md) | Details of the remote application the linked item is in. | [optional] 
**GlobalId** | Pointer to **string** | The global ID of the link, such as the ID of the item on the remote system. | [optional] 
**Id** | Pointer to **int64** | The ID of the link. | [optional] 
**Object** | Pointer to [**RemoteObject**](RemoteObject.md) | Details of the item linked to. | [optional] 
**Relationship** | Pointer to **string** | Description of the relationship between the issue and the linked item. | [optional] 
**Self** | Pointer to **string** | The URL of the link. | [optional] 

## Methods

### NewRemoteIssueLink

`func NewRemoteIssueLink() *RemoteIssueLink`

NewRemoteIssueLink instantiates a new RemoteIssueLink object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteIssueLinkWithDefaults

`func NewRemoteIssueLinkWithDefaults() *RemoteIssueLink`

NewRemoteIssueLinkWithDefaults instantiates a new RemoteIssueLink object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *RemoteIssueLink) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RemoteIssueLink) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RemoteIssueLink) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *RemoteIssueLink) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetGlobalId

`func (o *RemoteIssueLink) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *RemoteIssueLink) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *RemoteIssueLink) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *RemoteIssueLink) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetId

`func (o *RemoteIssueLink) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteIssueLink) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteIssueLink) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteIssueLink) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *RemoteIssueLink) GetObject() RemoteObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RemoteIssueLink) GetObjectOk() (*RemoteObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RemoteIssueLink) SetObject(v RemoteObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *RemoteIssueLink) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelationship

`func (o *RemoteIssueLink) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RemoteIssueLink) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RemoteIssueLink) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *RemoteIssueLink) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetSelf

`func (o *RemoteIssueLink) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RemoteIssueLink) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RemoteIssueLink) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *RemoteIssueLink) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


