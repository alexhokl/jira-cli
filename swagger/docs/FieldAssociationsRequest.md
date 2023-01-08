# FieldAssociationsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationContexts** | [**[]AssociationContextObject**](AssociationContextObject.md) | Contexts to associate/unassociate the fields with. | 
**Fields** | [**[]FieldIdentifierObject**](FieldIdentifierObject.md) | Fields to associate/unassociate with projects. | 

## Methods

### NewFieldAssociationsRequest

`func NewFieldAssociationsRequest(associationContexts []AssociationContextObject, fields []FieldIdentifierObject, ) *FieldAssociationsRequest`

NewFieldAssociationsRequest instantiates a new FieldAssociationsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFieldAssociationsRequestWithDefaults

`func NewFieldAssociationsRequestWithDefaults() *FieldAssociationsRequest`

NewFieldAssociationsRequestWithDefaults instantiates a new FieldAssociationsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationContexts

`func (o *FieldAssociationsRequest) GetAssociationContexts() []AssociationContextObject`

GetAssociationContexts returns the AssociationContexts field if non-nil, zero value otherwise.

### GetAssociationContextsOk

`func (o *FieldAssociationsRequest) GetAssociationContextsOk() (*[]AssociationContextObject, bool)`

GetAssociationContextsOk returns a tuple with the AssociationContexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationContexts

`func (o *FieldAssociationsRequest) SetAssociationContexts(v []AssociationContextObject)`

SetAssociationContexts sets AssociationContexts field to given value.


### GetFields

`func (o *FieldAssociationsRequest) GetFields() []FieldIdentifierObject`

GetFields returns the Fields field if non-nil, zero value otherwise.

### GetFieldsOk

`func (o *FieldAssociationsRequest) GetFieldsOk() (*[]FieldIdentifierObject, bool)`

GetFieldsOk returns a tuple with the Fields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFields

`func (o *FieldAssociationsRequest) SetFields(v []FieldIdentifierObject)`

SetFields sets Fields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


