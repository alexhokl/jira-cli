# SecuritySchemePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | The description of the security scheme | [optional] 
**Name** | Pointer to **string** | The name of the security scheme | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**SecurityLevels** | Pointer to [**[]SecurityLevelPayload**](SecurityLevelPayload.md) | The security levels for the security scheme | [optional] 

## Methods

### NewSecuritySchemePayload

`func NewSecuritySchemePayload() *SecuritySchemePayload`

NewSecuritySchemePayload instantiates a new SecuritySchemePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSecuritySchemePayloadWithDefaults

`func NewSecuritySchemePayloadWithDefaults() *SecuritySchemePayload`

NewSecuritySchemePayloadWithDefaults instantiates a new SecuritySchemePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *SecuritySchemePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SecuritySchemePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SecuritySchemePayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SecuritySchemePayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *SecuritySchemePayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SecuritySchemePayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SecuritySchemePayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SecuritySchemePayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPcri

`func (o *SecuritySchemePayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *SecuritySchemePayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *SecuritySchemePayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *SecuritySchemePayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetSecurityLevels

`func (o *SecuritySchemePayload) GetSecurityLevels() []SecurityLevelPayload`

GetSecurityLevels returns the SecurityLevels field if non-nil, zero value otherwise.

### GetSecurityLevelsOk

`func (o *SecuritySchemePayload) GetSecurityLevelsOk() (*[]SecurityLevelPayload, bool)`

GetSecurityLevelsOk returns a tuple with the SecurityLevels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityLevels

`func (o *SecuritySchemePayload) SetSecurityLevels(v []SecurityLevelPayload)`

SetSecurityLevels sets SecurityLevels field to given value.

### HasSecurityLevels

`func (o *SecuritySchemePayload) HasSecurityLevels() bool`

HasSecurityLevels returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


