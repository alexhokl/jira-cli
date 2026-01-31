# ServiceIdOrKeysAssociation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationType** | **string** | Defines the association type.  | 
**Values** | **[]string** | The service ID or keys to associate the entity with.  The number of values counted across all associationTypes must not exceed a limit of 500.  | 

## Methods

### NewServiceIdOrKeysAssociation

`func NewServiceIdOrKeysAssociation(associationType string, values []string, ) *ServiceIdOrKeysAssociation`

NewServiceIdOrKeysAssociation instantiates a new ServiceIdOrKeysAssociation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceIdOrKeysAssociationWithDefaults

`func NewServiceIdOrKeysAssociationWithDefaults() *ServiceIdOrKeysAssociation`

NewServiceIdOrKeysAssociationWithDefaults instantiates a new ServiceIdOrKeysAssociation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationType

`func (o *ServiceIdOrKeysAssociation) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *ServiceIdOrKeysAssociation) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *ServiceIdOrKeysAssociation) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.


### GetValues

`func (o *ServiceIdOrKeysAssociation) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *ServiceIdOrKeysAssociation) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *ServiceIdOrKeysAssociation) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


