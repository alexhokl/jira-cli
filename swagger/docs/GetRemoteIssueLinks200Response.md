# GetRemoteIssueLinks200Response

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

### NewGetRemoteIssueLinks200Response

`func NewGetRemoteIssueLinks200Response() *GetRemoteIssueLinks200Response`

NewGetRemoteIssueLinks200Response instantiates a new GetRemoteIssueLinks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetRemoteIssueLinks200ResponseWithDefaults

`func NewGetRemoteIssueLinks200ResponseWithDefaults() *GetRemoteIssueLinks200Response`

NewGetRemoteIssueLinks200ResponseWithDefaults instantiates a new GetRemoteIssueLinks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *GetRemoteIssueLinks200Response) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *GetRemoteIssueLinks200Response) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *GetRemoteIssueLinks200Response) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *GetRemoteIssueLinks200Response) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetGlobalId

`func (o *GetRemoteIssueLinks200Response) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *GetRemoteIssueLinks200Response) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *GetRemoteIssueLinks200Response) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *GetRemoteIssueLinks200Response) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetId

`func (o *GetRemoteIssueLinks200Response) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetRemoteIssueLinks200Response) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetRemoteIssueLinks200Response) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetRemoteIssueLinks200Response) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObject

`func (o *GetRemoteIssueLinks200Response) GetObject() RemoteObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GetRemoteIssueLinks200Response) GetObjectOk() (*RemoteObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GetRemoteIssueLinks200Response) SetObject(v RemoteObject)`

SetObject sets Object field to given value.

### HasObject

`func (o *GetRemoteIssueLinks200Response) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetRelationship

`func (o *GetRemoteIssueLinks200Response) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *GetRemoteIssueLinks200Response) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *GetRemoteIssueLinks200Response) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *GetRemoteIssueLinks200Response) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.

### GetSelf

`func (o *GetRemoteIssueLinks200Response) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *GetRemoteIssueLinks200Response) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *GetRemoteIssueLinks200Response) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *GetRemoteIssueLinks200Response) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


