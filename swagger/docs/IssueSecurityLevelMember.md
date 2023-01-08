# IssueSecurityLevelMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | [**PermissionHolder**](PermissionHolder.md) | The user or group being granted the permission. It consists of a &#x60;type&#x60; and a type-dependent &#x60;parameter&#x60;. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information. | 
**Id** | **int64** | The ID of the issue security level member. | 
**IssueSecurityLevelId** | **int64** | The ID of the issue security level. | 

## Methods

### NewIssueSecurityLevelMember

`func NewIssueSecurityLevelMember(holder PermissionHolder, id int64, issueSecurityLevelId int64, ) *IssueSecurityLevelMember`

NewIssueSecurityLevelMember instantiates a new IssueSecurityLevelMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueSecurityLevelMemberWithDefaults

`func NewIssueSecurityLevelMemberWithDefaults() *IssueSecurityLevelMember`

NewIssueSecurityLevelMemberWithDefaults instantiates a new IssueSecurityLevelMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *IssueSecurityLevelMember) GetHolder() PermissionHolder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *IssueSecurityLevelMember) GetHolderOk() (*PermissionHolder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *IssueSecurityLevelMember) SetHolder(v PermissionHolder)`

SetHolder sets Holder field to given value.


### GetId

`func (o *IssueSecurityLevelMember) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueSecurityLevelMember) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueSecurityLevelMember) SetId(v int64)`

SetId sets Id field to given value.


### GetIssueSecurityLevelId

`func (o *IssueSecurityLevelMember) GetIssueSecurityLevelId() int64`

GetIssueSecurityLevelId returns the IssueSecurityLevelId field if non-nil, zero value otherwise.

### GetIssueSecurityLevelIdOk

`func (o *IssueSecurityLevelMember) GetIssueSecurityLevelIdOk() (*int64, bool)`

GetIssueSecurityLevelIdOk returns a tuple with the IssueSecurityLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecurityLevelId

`func (o *IssueSecurityLevelMember) SetIssueSecurityLevelId(v int64)`

SetIssueSecurityLevelId sets IssueSecurityLevelId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


