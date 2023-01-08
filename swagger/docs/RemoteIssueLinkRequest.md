# RemoteIssueLinkRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**Application**](Application.md) | Details of the remote application the linked item is in. For example, trello. | [optional] 
**GlobalId** | Pointer to **string** | An identifier for the remote item in the remote system. For example, the global ID for a remote item in Confluence would consist of the app ID and page ID, like this: &#x60;appId&#x3D;456&amp;pageId&#x3D;123&#x60;.  Setting this field enables the remote issue link details to be updated or deleted using remote system and item details as the record identifier, rather than using the record&#39;s Jira ID.  The maximum length is 255 characters. | [optional] 
**Object** | [**RemoteObject**](RemoteObject.md) | Details of the item linked to. | 
**Relationship** | Pointer to **string** | Description of the relationship between the issue and the linked item. If not set, the relationship description \&quot;links to\&quot; is used in Jira. | [optional] 

## Methods

### NewRemoteIssueLinkRequest

`func NewRemoteIssueLinkRequest(object RemoteObject, ) *RemoteIssueLinkRequest`

NewRemoteIssueLinkRequest instantiates a new RemoteIssueLinkRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteIssueLinkRequestWithDefaults

`func NewRemoteIssueLinkRequestWithDefaults() *RemoteIssueLinkRequest`

NewRemoteIssueLinkRequestWithDefaults instantiates a new RemoteIssueLinkRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *RemoteIssueLinkRequest) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *RemoteIssueLinkRequest) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *RemoteIssueLinkRequest) SetApplication(v Application)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *RemoteIssueLinkRequest) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetGlobalId

`func (o *RemoteIssueLinkRequest) GetGlobalId() string`

GetGlobalId returns the GlobalId field if non-nil, zero value otherwise.

### GetGlobalIdOk

`func (o *RemoteIssueLinkRequest) GetGlobalIdOk() (*string, bool)`

GetGlobalIdOk returns a tuple with the GlobalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalId

`func (o *RemoteIssueLinkRequest) SetGlobalId(v string)`

SetGlobalId sets GlobalId field to given value.

### HasGlobalId

`func (o *RemoteIssueLinkRequest) HasGlobalId() bool`

HasGlobalId returns a boolean if a field has been set.

### GetObject

`func (o *RemoteIssueLinkRequest) GetObject() RemoteObject`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *RemoteIssueLinkRequest) GetObjectOk() (*RemoteObject, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *RemoteIssueLinkRequest) SetObject(v RemoteObject)`

SetObject sets Object field to given value.


### GetRelationship

`func (o *RemoteIssueLinkRequest) GetRelationship() string`

GetRelationship returns the Relationship field if non-nil, zero value otherwise.

### GetRelationshipOk

`func (o *RemoteIssueLinkRequest) GetRelationshipOk() (*string, bool)`

GetRelationshipOk returns a tuple with the Relationship field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationship

`func (o *RemoteIssueLinkRequest) SetRelationship(v string)`

SetRelationship sets Relationship field to given value.

### HasRelationship

`func (o *RemoteIssueLinkRequest) HasRelationship() bool`

HasRelationship returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


