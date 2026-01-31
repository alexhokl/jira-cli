# Associations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationType** | Pointer to **string** | the type of the association being made | [optional] 
**Values** | Pointer to **[]string** |  | [optional] 

## Methods

### NewAssociations

`func NewAssociations() *Associations`

NewAssociations instantiates a new Associations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssociationsWithDefaults

`func NewAssociationsWithDefaults() *Associations`

NewAssociationsWithDefaults instantiates a new Associations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationType

`func (o *Associations) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *Associations) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *Associations) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.

### HasAssociationType

`func (o *Associations) HasAssociationType() bool`

HasAssociationType returns a boolean if a field has been set.

### GetValues

`func (o *Associations) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *Associations) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *Associations) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *Associations) HasValues() bool`

HasValues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


