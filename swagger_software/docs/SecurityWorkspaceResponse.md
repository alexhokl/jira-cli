# SecurityWorkspaceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceId** | **string** | The Security Workspace ID  | 
**UpdatedAt** | **time.Time** | Latest date and time that the Security Workspace was updated in Jira.  | 

## Methods

### NewSecurityWorkspaceResponse

`func NewSecurityWorkspaceResponse(workspaceId string, updatedAt time.Time, ) *SecurityWorkspaceResponse`

NewSecurityWorkspaceResponse instantiates a new SecurityWorkspaceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityWorkspaceResponseWithDefaults

`func NewSecurityWorkspaceResponseWithDefaults() *SecurityWorkspaceResponse`

NewSecurityWorkspaceResponseWithDefaults instantiates a new SecurityWorkspaceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceId

`func (o *SecurityWorkspaceResponse) GetWorkspaceId() string`

GetWorkspaceId returns the WorkspaceId field if non-nil, zero value otherwise.

### GetWorkspaceIdOk

`func (o *SecurityWorkspaceResponse) GetWorkspaceIdOk() (*string, bool)`

GetWorkspaceIdOk returns a tuple with the WorkspaceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceId

`func (o *SecurityWorkspaceResponse) SetWorkspaceId(v string)`

SetWorkspaceId sets WorkspaceId field to given value.


### GetUpdatedAt

`func (o *SecurityWorkspaceResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SecurityWorkspaceResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SecurityWorkspaceResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


