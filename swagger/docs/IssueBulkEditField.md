# IssueBulkEditField

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Description of the field. | [optional] 
**FieldOptions** | Pointer to **[]map[string]interface{}** | A list of options related to the field, applicable in contexts where multiple selections are allowed. | [optional] 
**Id** | Pointer to **string** | The unique ID of the field. | [optional] 
**IsRequired** | Pointer to **bool** | Indicates whether the field is mandatory for the operation. | [optional] 
**MultiSelectFieldOptions** | Pointer to **[]string** | Specifies supported actions (like add, replace, remove) on multi-select fields via an enum. | [optional] 
**Name** | Pointer to **string** | The display name of the field. | [optional] 
**SearchUrl** | Pointer to **string** | A URL to fetch additional data for the field | [optional] 
**Type** | Pointer to **string** | The type of the field. | [optional] 
**UnavailableMessage** | Pointer to **string** | A message indicating why the field is unavailable for editing. | [optional] 

## Methods

### NewIssueBulkEditField

`func NewIssueBulkEditField() *IssueBulkEditField`

NewIssueBulkEditField instantiates a new IssueBulkEditField object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIssueBulkEditFieldWithDefaults

`func NewIssueBulkEditFieldWithDefaults() *IssueBulkEditField`

NewIssueBulkEditFieldWithDefaults instantiates a new IssueBulkEditField object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *IssueBulkEditField) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *IssueBulkEditField) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *IssueBulkEditField) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *IssueBulkEditField) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFieldOptions

`func (o *IssueBulkEditField) GetFieldOptions() []map[string]interface{}`

GetFieldOptions returns the FieldOptions field if non-nil, zero value otherwise.

### GetFieldOptionsOk

`func (o *IssueBulkEditField) GetFieldOptionsOk() (*[]map[string]interface{}, bool)`

GetFieldOptionsOk returns a tuple with the FieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldOptions

`func (o *IssueBulkEditField) SetFieldOptions(v []map[string]interface{})`

SetFieldOptions sets FieldOptions field to given value.

### HasFieldOptions

`func (o *IssueBulkEditField) HasFieldOptions() bool`

HasFieldOptions returns a boolean if a field has been set.

### GetId

`func (o *IssueBulkEditField) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IssueBulkEditField) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IssueBulkEditField) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *IssueBulkEditField) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsRequired

`func (o *IssueBulkEditField) GetIsRequired() bool`

GetIsRequired returns the IsRequired field if non-nil, zero value otherwise.

### GetIsRequiredOk

`func (o *IssueBulkEditField) GetIsRequiredOk() (*bool, bool)`

GetIsRequiredOk returns a tuple with the IsRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRequired

`func (o *IssueBulkEditField) SetIsRequired(v bool)`

SetIsRequired sets IsRequired field to given value.

### HasIsRequired

`func (o *IssueBulkEditField) HasIsRequired() bool`

HasIsRequired returns a boolean if a field has been set.

### GetMultiSelectFieldOptions

`func (o *IssueBulkEditField) GetMultiSelectFieldOptions() []string`

GetMultiSelectFieldOptions returns the MultiSelectFieldOptions field if non-nil, zero value otherwise.

### GetMultiSelectFieldOptionsOk

`func (o *IssueBulkEditField) GetMultiSelectFieldOptionsOk() (*[]string, bool)`

GetMultiSelectFieldOptionsOk returns a tuple with the MultiSelectFieldOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultiSelectFieldOptions

`func (o *IssueBulkEditField) SetMultiSelectFieldOptions(v []string)`

SetMultiSelectFieldOptions sets MultiSelectFieldOptions field to given value.

### HasMultiSelectFieldOptions

`func (o *IssueBulkEditField) HasMultiSelectFieldOptions() bool`

HasMultiSelectFieldOptions returns a boolean if a field has been set.

### GetName

`func (o *IssueBulkEditField) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IssueBulkEditField) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IssueBulkEditField) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IssueBulkEditField) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSearchUrl

`func (o *IssueBulkEditField) GetSearchUrl() string`

GetSearchUrl returns the SearchUrl field if non-nil, zero value otherwise.

### GetSearchUrlOk

`func (o *IssueBulkEditField) GetSearchUrlOk() (*string, bool)`

GetSearchUrlOk returns a tuple with the SearchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchUrl

`func (o *IssueBulkEditField) SetSearchUrl(v string)`

SetSearchUrl sets SearchUrl field to given value.

### HasSearchUrl

`func (o *IssueBulkEditField) HasSearchUrl() bool`

HasSearchUrl returns a boolean if a field has been set.

### GetType

`func (o *IssueBulkEditField) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IssueBulkEditField) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IssueBulkEditField) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IssueBulkEditField) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUnavailableMessage

`func (o *IssueBulkEditField) GetUnavailableMessage() string`

GetUnavailableMessage returns the UnavailableMessage field if non-nil, zero value otherwise.

### GetUnavailableMessageOk

`func (o *IssueBulkEditField) GetUnavailableMessageOk() (*string, bool)`

GetUnavailableMessageOk returns a tuple with the UnavailableMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnavailableMessage

`func (o *IssueBulkEditField) SetUnavailableMessage(v string)`

SetUnavailableMessage sets UnavailableMessage field to given value.

### HasUnavailableMessage

`func (o *IssueBulkEditField) HasUnavailableMessage() bool`

HasUnavailableMessage returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


