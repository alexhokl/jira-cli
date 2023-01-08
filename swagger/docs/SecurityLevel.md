# SecurityLevel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the issue level security item. | [optional] [readonly] 
**Id** | Pointer to **string** | The ID of the issue level security item. | [optional] [readonly] 
**IsDefault** | Pointer to **bool** | Whether the issue level security item is the default. | [optional] [readonly] 
**IssueSecuritySchemeId** | Pointer to **string** | The ID of the issue level security scheme. | [optional] [readonly] 
**Name** | Pointer to **string** | The name of the issue level security item. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the issue level security item. | [optional] [readonly] 

## Methods

### NewSecurityLevel

`func NewSecurityLevel() *SecurityLevel`

NewSecurityLevel instantiates a new SecurityLevel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecurityLevelWithDefaults

`func NewSecurityLevelWithDefaults() *SecurityLevel`

NewSecurityLevelWithDefaults instantiates a new SecurityLevel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecurityLevel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityLevel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityLevel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityLevel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SecurityLevel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityLevel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityLevel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityLevel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsDefault

`func (o *SecurityLevel) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *SecurityLevel) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *SecurityLevel) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *SecurityLevel) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIssueSecuritySchemeId

`func (o *SecurityLevel) GetIssueSecuritySchemeId() string`

GetIssueSecuritySchemeId returns the IssueSecuritySchemeId field if non-nil, zero value otherwise.

### GetIssueSecuritySchemeIdOk

`func (o *SecurityLevel) GetIssueSecuritySchemeIdOk() (*string, bool)`

GetIssueSecuritySchemeIdOk returns a tuple with the IssueSecuritySchemeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueSecuritySchemeId

`func (o *SecurityLevel) SetIssueSecuritySchemeId(v string)`

SetIssueSecuritySchemeId sets IssueSecuritySchemeId field to given value.

### HasIssueSecuritySchemeId

`func (o *SecurityLevel) HasIssueSecuritySchemeId() bool`

HasIssueSecuritySchemeId returns a boolean if a field has been set.

### GetName

`func (o *SecurityLevel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityLevel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityLevel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityLevel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *SecurityLevel) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SecurityLevel) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SecurityLevel) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *SecurityLevel) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


