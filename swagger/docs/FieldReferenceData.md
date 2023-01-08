# FieldReferenceData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auto** | Pointer to **string** | Whether the field provide auto-complete suggestions. | [optional] 
**Cfid** | Pointer to **string** | If the item is a custom field, the ID of the custom field. | [optional] 
**Deprecated** | Pointer to **string** | Whether this field has been deprecated. | [optional] 
**DeprecatedSearcherKey** | Pointer to **string** | The searcher key of the field, only passed when the field is deprecated. | [optional] 
**DisplayName** | Pointer to **string** | The display name contains the following:   *  for system fields, the field name. For example, &#x60;Summary&#x60;.  *  for collapsed custom fields, the field name followed by a hyphen and then the field name and field type. For example, &#x60;Component - Component[Dropdown]&#x60;.  *  for other custom fields, the field name followed by a hyphen and then the custom field ID. For example, &#x60;Component - cf[10061]&#x60;. | [optional] 
**Operators** | Pointer to **[]string** | The valid search operators for the field. | [optional] 
**Orderable** | Pointer to **string** | Whether the field can be used in a query&#39;s &#x60;ORDER BY&#x60; clause. | [optional] 
**Searchable** | Pointer to **string** | Whether the content of this field can be searched. | [optional] 
**Types** | Pointer to **[]string** | The data types of items in the field. | [optional] 
**Value** | Pointer to **string** | The field identifier. | [optional] 

## Methods

### NewFieldReferenceData

`func NewFieldReferenceData() *FieldReferenceData`

NewFieldReferenceData instantiates a new FieldReferenceData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldReferenceDataWithDefaults

`func NewFieldReferenceDataWithDefaults() *FieldReferenceData`

NewFieldReferenceDataWithDefaults instantiates a new FieldReferenceData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuto

`func (o *FieldReferenceData) GetAuto() string`

GetAuto returns the Auto field if non-nil, zero value otherwise.

### GetAutoOk

`func (o *FieldReferenceData) GetAutoOk() (*string, bool)`

GetAutoOk returns a tuple with the Auto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuto

`func (o *FieldReferenceData) SetAuto(v string)`

SetAuto sets Auto field to given value.

### HasAuto

`func (o *FieldReferenceData) HasAuto() bool`

HasAuto returns a boolean if a field has been set.

### GetCfid

`func (o *FieldReferenceData) GetCfid() string`

GetCfid returns the Cfid field if non-nil, zero value otherwise.

### GetCfidOk

`func (o *FieldReferenceData) GetCfidOk() (*string, bool)`

GetCfidOk returns a tuple with the Cfid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCfid

`func (o *FieldReferenceData) SetCfid(v string)`

SetCfid sets Cfid field to given value.

### HasCfid

`func (o *FieldReferenceData) HasCfid() bool`

HasCfid returns a boolean if a field has been set.

### GetDeprecated

`func (o *FieldReferenceData) GetDeprecated() string`

GetDeprecated returns the Deprecated field if non-nil, zero value otherwise.

### GetDeprecatedOk

`func (o *FieldReferenceData) GetDeprecatedOk() (*string, bool)`

GetDeprecatedOk returns a tuple with the Deprecated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecated

`func (o *FieldReferenceData) SetDeprecated(v string)`

SetDeprecated sets Deprecated field to given value.

### HasDeprecated

`func (o *FieldReferenceData) HasDeprecated() bool`

HasDeprecated returns a boolean if a field has been set.

### GetDeprecatedSearcherKey

`func (o *FieldReferenceData) GetDeprecatedSearcherKey() string`

GetDeprecatedSearcherKey returns the DeprecatedSearcherKey field if non-nil, zero value otherwise.

### GetDeprecatedSearcherKeyOk

`func (o *FieldReferenceData) GetDeprecatedSearcherKeyOk() (*string, bool)`

GetDeprecatedSearcherKeyOk returns a tuple with the DeprecatedSearcherKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeprecatedSearcherKey

`func (o *FieldReferenceData) SetDeprecatedSearcherKey(v string)`

SetDeprecatedSearcherKey sets DeprecatedSearcherKey field to given value.

### HasDeprecatedSearcherKey

`func (o *FieldReferenceData) HasDeprecatedSearcherKey() bool`

HasDeprecatedSearcherKey returns a boolean if a field has been set.

### GetDisplayName

`func (o *FieldReferenceData) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *FieldReferenceData) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *FieldReferenceData) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *FieldReferenceData) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetOperators

`func (o *FieldReferenceData) GetOperators() []string`

GetOperators returns the Operators field if non-nil, zero value otherwise.

### GetOperatorsOk

`func (o *FieldReferenceData) GetOperatorsOk() (*[]string, bool)`

GetOperatorsOk returns a tuple with the Operators field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperators

`func (o *FieldReferenceData) SetOperators(v []string)`

SetOperators sets Operators field to given value.

### HasOperators

`func (o *FieldReferenceData) HasOperators() bool`

HasOperators returns a boolean if a field has been set.

### GetOrderable

`func (o *FieldReferenceData) GetOrderable() string`

GetOrderable returns the Orderable field if non-nil, zero value otherwise.

### GetOrderableOk

`func (o *FieldReferenceData) GetOrderableOk() (*string, bool)`

GetOrderableOk returns a tuple with the Orderable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderable

`func (o *FieldReferenceData) SetOrderable(v string)`

SetOrderable sets Orderable field to given value.

### HasOrderable

`func (o *FieldReferenceData) HasOrderable() bool`

HasOrderable returns a boolean if a field has been set.

### GetSearchable

`func (o *FieldReferenceData) GetSearchable() string`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *FieldReferenceData) GetSearchableOk() (*string, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *FieldReferenceData) SetSearchable(v string)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *FieldReferenceData) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.

### GetTypes

`func (o *FieldReferenceData) GetTypes() []string`

GetTypes returns the Types field if non-nil, zero value otherwise.

### GetTypesOk

`func (o *FieldReferenceData) GetTypesOk() (*[]string, bool)`

GetTypesOk returns a tuple with the Types field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypes

`func (o *FieldReferenceData) SetTypes(v []string)`

SetTypes sets Types field to given value.

### HasTypes

`func (o *FieldReferenceData) HasTypes() bool`

HasTypes returns a boolean if a field has been set.

### GetValue

`func (o *FieldReferenceData) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *FieldReferenceData) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *FieldReferenceData) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *FieldReferenceData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


