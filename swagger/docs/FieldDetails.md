# FieldDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClauseNames** | Pointer to **[]string** | The names that can be used to reference the field in an advanced search. For more information, see [Advanced searching - fields reference](https://confluence.atlassian.com/x/gwORLQ). | [optional] 
**Custom** | Pointer to **bool** | Whether the field is a custom field. | [optional] 
**Id** | Pointer to **string** | The ID of the field. | [optional] 
**Key** | Pointer to **string** | The key of the field. | [optional] 
**Name** | Pointer to **string** | The name of the field. | [optional] 
**Navigable** | Pointer to **bool** | Whether the field can be used as a column on the issue navigator. | [optional] 
**Orderable** | Pointer to **bool** | Whether the content of the field can be used to order lists. | [optional] 
**Schema** | Pointer to [**JsonTypeBean**](JsonTypeBean.md) | The data schema for the field. | [optional] 
**Scope** | Pointer to [**Scope**](Scope.md) | The scope of the field. | [optional] 
**Searchable** | Pointer to **bool** | Whether the content of the field can be searched. | [optional] 

## Methods

### NewFieldDetails

`func NewFieldDetails() *FieldDetails`

NewFieldDetails instantiates a new FieldDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldDetailsWithDefaults

`func NewFieldDetailsWithDefaults() *FieldDetails`

NewFieldDetailsWithDefaults instantiates a new FieldDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClauseNames

`func (o *FieldDetails) GetClauseNames() []string`

GetClauseNames returns the ClauseNames field if non-nil, zero value otherwise.

### GetClauseNamesOk

`func (o *FieldDetails) GetClauseNamesOk() (*[]string, bool)`

GetClauseNamesOk returns a tuple with the ClauseNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClauseNames

`func (o *FieldDetails) SetClauseNames(v []string)`

SetClauseNames sets ClauseNames field to given value.

### HasClauseNames

`func (o *FieldDetails) HasClauseNames() bool`

HasClauseNames returns a boolean if a field has been set.

### GetCustom

`func (o *FieldDetails) GetCustom() bool`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *FieldDetails) GetCustomOk() (*bool, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *FieldDetails) SetCustom(v bool)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *FieldDetails) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetId

`func (o *FieldDetails) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FieldDetails) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FieldDetails) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FieldDetails) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *FieldDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *FieldDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *FieldDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *FieldDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetName

`func (o *FieldDetails) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FieldDetails) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FieldDetails) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FieldDetails) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavigable

`func (o *FieldDetails) GetNavigable() bool`

GetNavigable returns the Navigable field if non-nil, zero value otherwise.

### GetNavigableOk

`func (o *FieldDetails) GetNavigableOk() (*bool, bool)`

GetNavigableOk returns a tuple with the Navigable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavigable

`func (o *FieldDetails) SetNavigable(v bool)`

SetNavigable sets Navigable field to given value.

### HasNavigable

`func (o *FieldDetails) HasNavigable() bool`

HasNavigable returns a boolean if a field has been set.

### GetOrderable

`func (o *FieldDetails) GetOrderable() bool`

GetOrderable returns the Orderable field if non-nil, zero value otherwise.

### GetOrderableOk

`func (o *FieldDetails) GetOrderableOk() (*bool, bool)`

GetOrderableOk returns a tuple with the Orderable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderable

`func (o *FieldDetails) SetOrderable(v bool)`

SetOrderable sets Orderable field to given value.

### HasOrderable

`func (o *FieldDetails) HasOrderable() bool`

HasOrderable returns a boolean if a field has been set.

### GetSchema

`func (o *FieldDetails) GetSchema() JsonTypeBean`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *FieldDetails) GetSchemaOk() (*JsonTypeBean, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *FieldDetails) SetSchema(v JsonTypeBean)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *FieldDetails) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### GetScope

`func (o *FieldDetails) GetScope() Scope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *FieldDetails) GetScopeOk() (*Scope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *FieldDetails) SetScope(v Scope)`

SetScope sets Scope field to given value.

### HasScope

`func (o *FieldDetails) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSearchable

`func (o *FieldDetails) GetSearchable() bool`

GetSearchable returns the Searchable field if non-nil, zero value otherwise.

### GetSearchableOk

`func (o *FieldDetails) GetSearchableOk() (*bool, bool)`

GetSearchableOk returns a tuple with the Searchable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchable

`func (o *FieldDetails) SetSearchable(v bool)`

SetSearchable sets Searchable field to given value.

### HasSearchable

`func (o *FieldDetails) HasSearchable() bool`

HasSearchable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


