# SubmitSecurityWorkspacesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WorkspaceIds** | **[]string** | The IDs of Security Workspaces to link to this Jira site. These must follow this regex pattern: &#x60;[a-zA-Z0-9\\\\-_.~@:{}&#x3D;]+(/[a-zA-Z0-9\\\\-_.~@:{}&#x3D;]+)*&#x60;  | 

## Methods

### NewSubmitSecurityWorkspacesRequest

`func NewSubmitSecurityWorkspacesRequest(workspaceIds []string, ) *SubmitSecurityWorkspacesRequest`

NewSubmitSecurityWorkspacesRequest instantiates a new SubmitSecurityWorkspacesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubmitSecurityWorkspacesRequestWithDefaults

`func NewSubmitSecurityWorkspacesRequestWithDefaults() *SubmitSecurityWorkspacesRequest`

NewSubmitSecurityWorkspacesRequestWithDefaults instantiates a new SubmitSecurityWorkspacesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWorkspaceIds

`func (o *SubmitSecurityWorkspacesRequest) GetWorkspaceIds() []string`

GetWorkspaceIds returns the WorkspaceIds field if non-nil, zero value otherwise.

### GetWorkspaceIdsOk

`func (o *SubmitSecurityWorkspacesRequest) GetWorkspaceIdsOk() (*[]string, bool)`

GetWorkspaceIdsOk returns a tuple with the WorkspaceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkspaceIds

`func (o *SubmitSecurityWorkspacesRequest) SetWorkspaceIds(v []string)`

SetWorkspaceIds sets WorkspaceIds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


