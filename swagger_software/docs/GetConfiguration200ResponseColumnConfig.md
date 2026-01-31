# GetConfiguration200ResponseColumnConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | Pointer to [**[]GetConfiguration200ResponseColumnConfigColumnsInner**](GetConfiguration200ResponseColumnConfigColumnsInner.md) |  | [optional] 
**ConstraintType** | Pointer to **string** |  | [optional] 

## Methods

### NewGetConfiguration200ResponseColumnConfig

`func NewGetConfiguration200ResponseColumnConfig() *GetConfiguration200ResponseColumnConfig`

NewGetConfiguration200ResponseColumnConfig instantiates a new GetConfiguration200ResponseColumnConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConfiguration200ResponseColumnConfigWithDefaults

`func NewGetConfiguration200ResponseColumnConfigWithDefaults() *GetConfiguration200ResponseColumnConfig`

NewGetConfiguration200ResponseColumnConfigWithDefaults instantiates a new GetConfiguration200ResponseColumnConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *GetConfiguration200ResponseColumnConfig) GetColumns() []GetConfiguration200ResponseColumnConfigColumnsInner`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *GetConfiguration200ResponseColumnConfig) GetColumnsOk() (*[]GetConfiguration200ResponseColumnConfigColumnsInner, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *GetConfiguration200ResponseColumnConfig) SetColumns(v []GetConfiguration200ResponseColumnConfigColumnsInner)`

SetColumns sets Columns field to given value.

### HasColumns

`func (o *GetConfiguration200ResponseColumnConfig) HasColumns() bool`

HasColumns returns a boolean if a field has been set.

### GetConstraintType

`func (o *GetConfiguration200ResponseColumnConfig) GetConstraintType() string`

GetConstraintType returns the ConstraintType field if non-nil, zero value otherwise.

### GetConstraintTypeOk

`func (o *GetConfiguration200ResponseColumnConfig) GetConstraintTypeOk() (*string, bool)`

GetConstraintTypeOk returns a tuple with the ConstraintType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConstraintType

`func (o *GetConfiguration200ResponseColumnConfig) SetConstraintType(v string)`

SetConstraintType sets ConstraintType field to given value.

### HasConstraintType

`func (o *GetConfiguration200ResponseColumnConfig) HasConstraintType() bool`

HasConstraintType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


