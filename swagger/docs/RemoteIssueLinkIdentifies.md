# RemoteIssueLinkIdentifies

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the remote issue link, such as the ID of the item on the remote system. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the remote issue link. | [optional] [readonly] 

## Methods

### NewRemoteIssueLinkIdentifies

`func NewRemoteIssueLinkIdentifies() *RemoteIssueLinkIdentifies`

NewRemoteIssueLinkIdentifies instantiates a new RemoteIssueLinkIdentifies object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteIssueLinkIdentifiesWithDefaults

`func NewRemoteIssueLinkIdentifiesWithDefaults() *RemoteIssueLinkIdentifies`

NewRemoteIssueLinkIdentifiesWithDefaults instantiates a new RemoteIssueLinkIdentifies object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *RemoteIssueLinkIdentifies) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *RemoteIssueLinkIdentifies) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *RemoteIssueLinkIdentifies) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *RemoteIssueLinkIdentifies) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSelf

`func (o *RemoteIssueLinkIdentifies) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *RemoteIssueLinkIdentifies) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *RemoteIssueLinkIdentifies) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *RemoteIssueLinkIdentifies) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


