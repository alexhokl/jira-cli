# ActorInputBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to **[]string** | The name of the group to add as a default actor. This parameter cannot be used with the &#x60;groupId&#x60; parameter. As a group&#39;s name can change,use of &#x60;groupId&#x60; is recommended. This parameter accepts a comma-separated list. For example, &#x60;\&quot;group\&quot;:[\&quot;project-admin\&quot;, \&quot;jira-developers\&quot;]&#x60;. | [optional] 
**GroupId** | Pointer to **[]string** | The ID of the group to add as a default actor. This parameter cannot be used with the &#x60;group&#x60; parameter This parameter accepts a comma-separated list. For example, &#x60;\&quot;groupId\&quot;:[\&quot;77f6ab39-e755-4570-a6ae-2d7a8df0bcb8\&quot;, \&quot;0c011f85-69ed-49c4-a801-3b18d0f771bc\&quot;]&#x60;. | [optional] 
**User** | Pointer to **[]string** | The account IDs of the users to add as default actors. This parameter accepts a comma-separated list. For example, &#x60;\&quot;user\&quot;:[\&quot;5b10a2844c20165700ede21g\&quot;, \&quot;5b109f2e9729b51b54dc274d\&quot;]&#x60;. | [optional] 

## Methods

### NewActorInputBean

`func NewActorInputBean() *ActorInputBean`

NewActorInputBean instantiates a new ActorInputBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActorInputBeanWithDefaults

`func NewActorInputBeanWithDefaults() *ActorInputBean`

NewActorInputBeanWithDefaults instantiates a new ActorInputBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *ActorInputBean) GetGroup() []string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *ActorInputBean) GetGroupOk() (*[]string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *ActorInputBean) SetGroup(v []string)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *ActorInputBean) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetGroupId

`func (o *ActorInputBean) GetGroupId() []string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *ActorInputBean) GetGroupIdOk() (*[]string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *ActorInputBean) SetGroupId(v []string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *ActorInputBean) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetUser

`func (o *ActorInputBean) GetUser() []string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ActorInputBean) GetUserOk() (*[]string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ActorInputBean) SetUser(v []string)`

SetUser sets User field to given value.

### HasUser

`func (o *ActorInputBean) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


