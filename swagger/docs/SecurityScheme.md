# SecurityScheme

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultSecurityLevelId** | Pointer to **int64** | The ID of the default security level. | [optional] [readonly] 
**Description** | Pointer to **string** | The description of the issue security scheme. | [optional] [readonly] 
**Id** | Pointer to **int64** | The ID of the issue security scheme. | [optional] [readonly] 
**Levels** | Pointer to [**[]SecurityLevel**](SecurityLevel.md) |  | [optional] 
**Name** | Pointer to **string** | The name of the issue security scheme. | [optional] [readonly] 
**Self** | Pointer to **string** | The URL of the issue security scheme. | [optional] [readonly] 

## Methods

### NewSecurityScheme

`func NewSecurityScheme() *SecurityScheme`

NewSecurityScheme instantiates a new SecurityScheme object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemeWithDefaults

`func NewSecuritySchemeWithDefaults() *SecurityScheme`

NewSecuritySchemeWithDefaults instantiates a new SecurityScheme object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultSecurityLevelId

`func (o *SecurityScheme) GetDefaultSecurityLevelId() int64`

GetDefaultSecurityLevelId returns the DefaultSecurityLevelId field if non-nil, zero value otherwise.

### GetDefaultSecurityLevelIdOk

`func (o *SecurityScheme) GetDefaultSecurityLevelIdOk() (*int64, bool)`

GetDefaultSecurityLevelIdOk returns a tuple with the DefaultSecurityLevelId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSecurityLevelId

`func (o *SecurityScheme) SetDefaultSecurityLevelId(v int64)`

SetDefaultSecurityLevelId sets DefaultSecurityLevelId field to given value.

### HasDefaultSecurityLevelId

`func (o *SecurityScheme) HasDefaultSecurityLevelId() bool`

HasDefaultSecurityLevelId returns a boolean if a field has been set.

### GetDescription

`func (o *SecurityScheme) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecurityScheme) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecurityScheme) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecurityScheme) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *SecurityScheme) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SecurityScheme) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SecurityScheme) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SecurityScheme) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLevels

`func (o *SecurityScheme) GetLevels() []SecurityLevel`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *SecurityScheme) GetLevelsOk() (*[]SecurityLevel, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *SecurityScheme) SetLevels(v []SecurityLevel)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *SecurityScheme) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetName

`func (o *SecurityScheme) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecurityScheme) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecurityScheme) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecurityScheme) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSelf

`func (o *SecurityScheme) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *SecurityScheme) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *SecurityScheme) SetSelf(v string)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *SecurityScheme) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


