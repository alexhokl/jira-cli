# Field

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextsCount** | Pointer to **int64** | Number of contexts where the field is used. | [optional] 
**Description** | Pointer to **string** | The description of the field. | [optional] 
**Id** | **string** | The ID of the field. | 
**IsLocked** | Pointer to **bool** | Whether the field is locked. | [optional] 
**IsUnscreenable** | Pointer to **bool** | Whether the field is shown on screen or not. | [optional] 
**Key** | Pointer to **string** | The key of the field. | [optional] 
**LastUsed** | Pointer to [**FieldLastUsed**](FieldLastUsed.md) |  | [optional] 
**Name** | **string** | The name of the field. | 
**ProjectsCount** | Pointer to **int64** | Number of projects where the field is used. | [optional] 
**Schema** | [**JsonTypeBean**](JsonTypeBean.md) |  | 
**ScreensCount** | Pointer to **int64** | Number of screens where the field is used. | [optional] 
**SearcherKey** | Pointer to **string** | The searcher key of the field. Returned for custom fields. | [optional] 
**StableId** | Pointer to **string** | The stable ID of the field. | [optional] 
**TypeDisplayName** | Pointer to **string** | The display name of the field type | [optional] 

## Methods

### NewField

`func NewField(id string, name string, schema JsonTypeBean, ) *Field`

NewField instantiates a new Field object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldWithDefaults

`func NewFieldWithDefaults() *Field`

NewFieldWithDefaults instantiates a new Field object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextsCount

`func (o *Field) GetContextsCount() int64`

GetContextsCount returns the ContextsCount field if non-nil, zero value otherwise.

### GetContextsCountOk

`func (o *Field) GetContextsCountOk() (*int64, bool)`

GetContextsCountOk returns a tuple with the ContextsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextsCount

`func (o *Field) SetContextsCount(v int64)`

SetContextsCount sets ContextsCount field to given value.

### HasContextsCount

`func (o *Field) HasContextsCount() bool`

HasContextsCount returns a boolean if a field has been set.

### GetDescription

`func (o *Field) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Field) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Field) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Field) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Field) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Field) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Field) SetId(v string)`

SetId sets Id field to given value.


### GetIsLocked

`func (o *Field) GetIsLocked() bool`

GetIsLocked returns the IsLocked field if non-nil, zero value otherwise.

### GetIsLockedOk

`func (o *Field) GetIsLockedOk() (*bool, bool)`

GetIsLockedOk returns a tuple with the IsLocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsLocked

`func (o *Field) SetIsLocked(v bool)`

SetIsLocked sets IsLocked field to given value.

### HasIsLocked

`func (o *Field) HasIsLocked() bool`

HasIsLocked returns a boolean if a field has been set.

### GetIsUnscreenable

`func (o *Field) GetIsUnscreenable() bool`

GetIsUnscreenable returns the IsUnscreenable field if non-nil, zero value otherwise.

### GetIsUnscreenableOk

`func (o *Field) GetIsUnscreenableOk() (*bool, bool)`

GetIsUnscreenableOk returns a tuple with the IsUnscreenable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsUnscreenable

`func (o *Field) SetIsUnscreenable(v bool)`

SetIsUnscreenable sets IsUnscreenable field to given value.

### HasIsUnscreenable

`func (o *Field) HasIsUnscreenable() bool`

HasIsUnscreenable returns a boolean if a field has been set.

### GetKey

`func (o *Field) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Field) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Field) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Field) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetLastUsed

`func (o *Field) GetLastUsed() FieldLastUsed`

GetLastUsed returns the LastUsed field if non-nil, zero value otherwise.

### GetLastUsedOk

`func (o *Field) GetLastUsedOk() (*FieldLastUsed, bool)`

GetLastUsedOk returns a tuple with the LastUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsed

`func (o *Field) SetLastUsed(v FieldLastUsed)`

SetLastUsed sets LastUsed field to given value.

### HasLastUsed

`func (o *Field) HasLastUsed() bool`

HasLastUsed returns a boolean if a field has been set.

### GetName

`func (o *Field) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Field) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Field) SetName(v string)`

SetName sets Name field to given value.


### GetProjectsCount

`func (o *Field) GetProjectsCount() int64`

GetProjectsCount returns the ProjectsCount field if non-nil, zero value otherwise.

### GetProjectsCountOk

`func (o *Field) GetProjectsCountOk() (*int64, bool)`

GetProjectsCountOk returns a tuple with the ProjectsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectsCount

`func (o *Field) SetProjectsCount(v int64)`

SetProjectsCount sets ProjectsCount field to given value.

### HasProjectsCount

`func (o *Field) HasProjectsCount() bool`

HasProjectsCount returns a boolean if a field has been set.

### GetSchema

`func (o *Field) GetSchema() JsonTypeBean`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *Field) GetSchemaOk() (*JsonTypeBean, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *Field) SetSchema(v JsonTypeBean)`

SetSchema sets Schema field to given value.


### GetScreensCount

`func (o *Field) GetScreensCount() int64`

GetScreensCount returns the ScreensCount field if non-nil, zero value otherwise.

### GetScreensCountOk

`func (o *Field) GetScreensCountOk() (*int64, bool)`

GetScreensCountOk returns a tuple with the ScreensCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreensCount

`func (o *Field) SetScreensCount(v int64)`

SetScreensCount sets ScreensCount field to given value.

### HasScreensCount

`func (o *Field) HasScreensCount() bool`

HasScreensCount returns a boolean if a field has been set.

### GetSearcherKey

`func (o *Field) GetSearcherKey() string`

GetSearcherKey returns the SearcherKey field if non-nil, zero value otherwise.

### GetSearcherKeyOk

`func (o *Field) GetSearcherKeyOk() (*string, bool)`

GetSearcherKeyOk returns a tuple with the SearcherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearcherKey

`func (o *Field) SetSearcherKey(v string)`

SetSearcherKey sets SearcherKey field to given value.

### HasSearcherKey

`func (o *Field) HasSearcherKey() bool`

HasSearcherKey returns a boolean if a field has been set.

### GetStableId

`func (o *Field) GetStableId() string`

GetStableId returns the StableId field if non-nil, zero value otherwise.

### GetStableIdOk

`func (o *Field) GetStableIdOk() (*string, bool)`

GetStableIdOk returns a tuple with the StableId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStableId

`func (o *Field) SetStableId(v string)`

SetStableId sets StableId field to given value.

### HasStableId

`func (o *Field) HasStableId() bool`

HasStableId returns a boolean if a field has been set.

### GetTypeDisplayName

`func (o *Field) GetTypeDisplayName() string`

GetTypeDisplayName returns the TypeDisplayName field if non-nil, zero value otherwise.

### GetTypeDisplayNameOk

`func (o *Field) GetTypeDisplayNameOk() (*string, bool)`

GetTypeDisplayNameOk returns a tuple with the TypeDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeDisplayName

`func (o *Field) SetTypeDisplayName(v string)`

SetTypeDisplayName sets TypeDisplayName field to given value.

### HasTypeDisplayName

`func (o *Field) HasTypeDisplayName() bool`

HasTypeDisplayName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


