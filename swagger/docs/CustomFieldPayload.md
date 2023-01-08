# CustomFieldPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CfType** | Pointer to **string** | The type of the custom field | [optional] 
**Description** | Pointer to **string** | The description of the custom field | [optional] 
**Name** | Pointer to **string** | The name of the custom field | [optional] 
**OnConflict** | Pointer to **string** | The strategy to use when there is a conflict with an existing custom field. FAIL - Fail execution, this always needs to be unique; USE - Use the existing entity and ignore new entity parameters | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Scope** | Pointer to **string** | Allows an overwrite to declare the new Custom Field to be created as a GLOBAL-scoped field. Leave this as empty or null to use the project&#39;s default scope. | [optional] 
**SearcherKey** | Pointer to **string** | The searcher key of the custom field | [optional] 

## Methods

### NewCustomFieldPayload

`func NewCustomFieldPayload() *CustomFieldPayload`

NewCustomFieldPayload instantiates a new CustomFieldPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomFieldPayloadWithDefaults

`func NewCustomFieldPayloadWithDefaults() *CustomFieldPayload`

NewCustomFieldPayloadWithDefaults instantiates a new CustomFieldPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCfType

`func (o *CustomFieldPayload) GetCfType() string`

GetCfType returns the CfType field if non-nil, zero value otherwise.

### GetCfTypeOk

`func (o *CustomFieldPayload) GetCfTypeOk() (*string, bool)`

GetCfTypeOk returns a tuple with the CfType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfType

`func (o *CustomFieldPayload) SetCfType(v string)`

SetCfType sets CfType field to given value.

### HasCfType

`func (o *CustomFieldPayload) HasCfType() bool`

HasCfType returns a boolean if a field has been set.

### GetDescription

`func (o *CustomFieldPayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CustomFieldPayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CustomFieldPayload) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CustomFieldPayload) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CustomFieldPayload) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CustomFieldPayload) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CustomFieldPayload) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CustomFieldPayload) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOnConflict

`func (o *CustomFieldPayload) GetOnConflict() string`

GetOnConflict returns the OnConflict field if non-nil, zero value otherwise.

### GetOnConflictOk

`func (o *CustomFieldPayload) GetOnConflictOk() (*string, bool)`

GetOnConflictOk returns a tuple with the OnConflict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnConflict

`func (o *CustomFieldPayload) SetOnConflict(v string)`

SetOnConflict sets OnConflict field to given value.

### HasOnConflict

`func (o *CustomFieldPayload) HasOnConflict() bool`

HasOnConflict returns a boolean if a field has been set.

### GetPcri

`func (o *CustomFieldPayload) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *CustomFieldPayload) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *CustomFieldPayload) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *CustomFieldPayload) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetScope

`func (o *CustomFieldPayload) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CustomFieldPayload) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CustomFieldPayload) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CustomFieldPayload) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSearcherKey

`func (o *CustomFieldPayload) GetSearcherKey() string`

GetSearcherKey returns the SearcherKey field if non-nil, zero value otherwise.

### GetSearcherKeyOk

`func (o *CustomFieldPayload) GetSearcherKeyOk() (*string, bool)`

GetSearcherKeyOk returns a tuple with the SearcherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearcherKey

`func (o *CustomFieldPayload) SetSearcherKey(v string)`

SetSearcherKey sets SearcherKey field to given value.

### HasSearcherKey

`func (o *CustomFieldPayload) HasSearcherKey() bool`

HasSearcherKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


