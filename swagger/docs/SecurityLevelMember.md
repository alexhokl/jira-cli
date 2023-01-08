# SecurityLevelMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Holder** | [**PermissionHolder**](PermissionHolder.md) | The user or group being granted the permission. It consists of a &#x60;type&#x60; and a type-dependent &#x60;parameter&#x60;. See [Holder object](../api-group-permission-schemes/#holder-object) in *Get all permission schemes* for more information. | [readonly] 
**Id** | **string** | The ID of the issue security level member. | [readonly] 
**IssueSecurityLevelId** | **string** | The ID of the issue security level. | [readonly] 
**IssueSecuritySchemeId** | **string** | The ID of the issue security scheme. | [readonly] 
**Managed** | Pointer to **bool** |  | [optional] 

## Methods

### NewSecurityLevelMember

`func NewSecurityLevelMember(holder PermissionHolder, id string, issueSecurityLevelId string, issueSecuritySchemeId string, ) *SecurityLevelMember`

NewSecurityLevelMember instantiates a new SecurityLevelMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLevelMemberWithDefaults

`func NewSecurityLevelMemberWithDefaults() *SecurityLevelMember`

NewSecurityLevelMemberWithDefaults instantiates a new SecurityLevelMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHolder

`func (o *SecurityLevelMember) GetHolder() PermissionHolder`

GetHolder returns the Holder field if non-nil, zero value otherwise.

### GetHolderOk

`func (o *SecurityLevelMember) GetHolderOk() (*PermissionHolder, bool)`

GetHolderOk returns a tuple with the Holder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHolder

`func (o *SecurityLevelMember) SetHolder(v PermissionHolder)`

SetHolder sets Holder field to given value.


### GetId

`func (o *SecurityLevelMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityLevelMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityLevelMember) SetId(v string)`

SetId sets Id field to given value.


### GetIssueSecurityLevelId

`func (o *SecurityLevelMember) GetIssueSecurityLevelId() string`

GetIssueSecurityLevelId returns the IssueSecurityLevelId field if non-nil, zero value otherwise.

### GetIssueSecurityLevelIdOk

`func (o *SecurityLevelMember) GetIssueSecurityLevelIdOk() (*string, bool)`

GetIssueSecurityLevelIdOk returns a tuple with the IssueSecurityLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecurityLevelId

`func (o *SecurityLevelMember) SetIssueSecurityLevelId(v string)`

SetIssueSecurityLevelId sets IssueSecurityLevelId field to given value.


### GetIssueSecuritySchemeId

`func (o *SecurityLevelMember) GetIssueSecuritySchemeId() string`

GetIssueSecuritySchemeId returns the IssueSecuritySchemeId field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeIdOk

`func (o *SecurityLevelMember) GetIssueSecuritySchemeIdOk() (*string, bool)`

GetIssueSecuritySchemeIdOk returns a tuple with the IssueSecuritySchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecuritySchemeId

`func (o *SecurityLevelMember) SetIssueSecuritySchemeId(v string)`

SetIssueSecuritySchemeId sets IssueSecuritySchemeId field to given value.


### GetManaged

`func (o *SecurityLevelMember) GetManaged() bool`

GetManaged returns the Managed field if non-nil, zero value otherwise.

### GetManagedOk

`func (o *SecurityLevelMember) GetManagedOk() (*bool, bool)`

GetManagedOk returns a tuple with the Managed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManaged

`func (o *SecurityLevelMember) SetManaged(v bool)`

SetManaged sets Managed field to given value.

### HasManaged

`func (o *SecurityLevelMember) HasManaged() bool`

HasManaged returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


