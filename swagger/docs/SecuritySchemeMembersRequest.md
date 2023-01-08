# SecuritySchemeMembersRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | Pointer to [**[]SecuritySchemeLevelMemberBean**](SecuritySchemeLevelMemberBean.md) | The list of level members which should be added to the issue security scheme level. | [optional] 

## Methods

### NewSecuritySchemeMembersRequest

`func NewSecuritySchemeMembersRequest() *SecuritySchemeMembersRequest`

NewSecuritySchemeMembersRequest instantiates a new SecuritySchemeMembersRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeMembersRequestWithDefaults

`func NewSecuritySchemeMembersRequestWithDefaults() *SecuritySchemeMembersRequest`

NewSecuritySchemeMembersRequestWithDefaults instantiates a new SecuritySchemeMembersRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *SecuritySchemeMembersRequest) GetMembers() []SecuritySchemeLevelMemberBean`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *SecuritySchemeMembersRequest) GetMembersOk() (*[]SecuritySchemeLevelMemberBean, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *SecuritySchemeMembersRequest) SetMembers(v []SecuritySchemeLevelMemberBean)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *SecuritySchemeMembersRequest) HasMembers() bool`

HasMembers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


