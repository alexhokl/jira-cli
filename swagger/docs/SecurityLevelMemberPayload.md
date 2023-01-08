# SecurityLevelMemberPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameter** | Pointer to **string** | Defines the value associated with the type. For reporter this would be \\{\&quot;null\&quot;\\}; for users this would be the names of specific users); for group this would be group names like \\{\&quot;administrators\&quot;, \&quot;jira-administrators\&quot;, \&quot;jira-users\&quot;\\} | [optional] 
**Type** | Pointer to **string** | The type of the security level member | [optional] 

## Methods

### NewSecurityLevelMemberPayload

`func NewSecurityLevelMemberPayload() *SecurityLevelMemberPayload`

NewSecurityLevelMemberPayload instantiates a new SecurityLevelMemberPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLevelMemberPayloadWithDefaults

`func NewSecurityLevelMemberPayloadWithDefaults() *SecurityLevelMemberPayload`

NewSecurityLevelMemberPayloadWithDefaults instantiates a new SecurityLevelMemberPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameter

`func (o *SecurityLevelMemberPayload) GetParameter() string`

GetParameter returns the Parameter field if non-nil, zero value otherwise.

### GetParameterOk

`func (o *SecurityLevelMemberPayload) GetParameterOk() (*string, bool)`

GetParameterOk returns a tuple with the Parameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameter

`func (o *SecurityLevelMemberPayload) SetParameter(v string)`

SetParameter sets Parameter field to given value.

### HasParameter

`func (o *SecurityLevelMemberPayload) HasParameter() bool`

HasParameter returns a boolean if a field has been set.

### GetType

`func (o *SecurityLevelMemberPayload) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SecurityLevelMemberPayload) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SecurityLevelMemberPayload) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *SecurityLevelMemberPayload) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


