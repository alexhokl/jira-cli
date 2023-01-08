# FieldLayoutConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Field** | Pointer to **bool** | Whether to show the field | [optional] 
**Pcri** | Pointer to [**ProjectCreateResourceIdentifier**](ProjectCreateResourceIdentifier.md) |  | [optional] 
**Required** | Pointer to **bool** | Whether the field is required | [optional] 

## Methods

### NewFieldLayoutConfiguration

`func NewFieldLayoutConfiguration() *FieldLayoutConfiguration`

NewFieldLayoutConfiguration instantiates a new FieldLayoutConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldLayoutConfigurationWithDefaults

`func NewFieldLayoutConfigurationWithDefaults() *FieldLayoutConfiguration`

NewFieldLayoutConfigurationWithDefaults instantiates a new FieldLayoutConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetField

`func (o *FieldLayoutConfiguration) GetField() bool`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *FieldLayoutConfiguration) GetFieldOk() (*bool, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *FieldLayoutConfiguration) SetField(v bool)`

SetField sets Field field to given value.

### HasField

`func (o *FieldLayoutConfiguration) HasField() bool`

HasField returns a boolean if a field has been set.

### GetPcri

`func (o *FieldLayoutConfiguration) GetPcri() ProjectCreateResourceIdentifier`

GetPcri returns the Pcri field if non-nil, zero value otherwise.

### GetPcriOk

`func (o *FieldLayoutConfiguration) GetPcriOk() (*ProjectCreateResourceIdentifier, bool)`

GetPcriOk returns a tuple with the Pcri field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPcri

`func (o *FieldLayoutConfiguration) SetPcri(v ProjectCreateResourceIdentifier)`

SetPcri sets Pcri field to given value.

### HasPcri

`func (o *FieldLayoutConfiguration) HasPcri() bool`

HasPcri returns a boolean if a field has been set.

### GetRequired

`func (o *FieldLayoutConfiguration) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *FieldLayoutConfiguration) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *FieldLayoutConfiguration) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *FieldLayoutConfiguration) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


