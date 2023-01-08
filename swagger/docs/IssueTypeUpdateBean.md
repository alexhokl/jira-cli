# IssueTypeUpdateBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AvatarId** | Pointer to **int64** | The ID of an issue type avatar. This can be obtained be obtained from the following endpoints:   *  [System issue type avatar IDs only](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-avatars/#api-rest-api-3-avatar-type-system-get)  *  [System and custom issue type avatar IDs](https://developer.atlassian.com/cloud/jira/platform/rest/v3/api-group-avatars/#api-rest-api-3-universal-avatar-type-type-owner-entityid-get) | [optional] 
**Description** | Pointer to **string** | The description of the issue type. | [optional] 
**Name** | Pointer to **string** | The unique name for the issue type. The maximum length is 60 characters. | [optional] 

## Methods

### NewIssueTypeUpdateBean

`func NewIssueTypeUpdateBean() *IssueTypeUpdateBean`

NewIssueTypeUpdateBean instantiates a new IssueTypeUpdateBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueTypeUpdateBeanWithDefaults

`func NewIssueTypeUpdateBeanWithDefaults() *IssueTypeUpdateBean`

NewIssueTypeUpdateBeanWithDefaults instantiates a new IssueTypeUpdateBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvatarId

`func (o *IssueTypeUpdateBean) GetAvatarId() int64`

GetAvatarId returns the AvatarId field if non-nil, zero value otherwise.

### GetAvatarIdOk

`func (o *IssueTypeUpdateBean) GetAvatarIdOk() (*int64, bool)`

GetAvatarIdOk returns a tuple with the AvatarId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatarId

`func (o *IssueTypeUpdateBean) SetAvatarId(v int64)`

SetAvatarId sets AvatarId field to given value.

### HasAvatarId

`func (o *IssueTypeUpdateBean) HasAvatarId() bool`

HasAvatarId returns a boolean if a field has been set.

### GetDescription

`func (o *IssueTypeUpdateBean) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueTypeUpdateBean) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueTypeUpdateBean) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueTypeUpdateBean) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *IssueTypeUpdateBean) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueTypeUpdateBean) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueTypeUpdateBean) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueTypeUpdateBean) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


