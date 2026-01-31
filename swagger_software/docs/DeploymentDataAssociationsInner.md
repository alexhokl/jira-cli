# DeploymentDataAssociationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationType** | **string** | Defines the association type. Currently supported entities can be found in this field&#39;s value enums list.  | 
**Values** | [**[]EntityAssociationValuesInner**](EntityAssociationValuesInner.md) | The entity keys that represent the entities to be associated. The number of values counted across all associationTypes must not exceed a limit of 500.  | 

## Methods

### NewDeploymentDataAssociationsInner

`func NewDeploymentDataAssociationsInner(associationType string, values []EntityAssociationValuesInner, ) *DeploymentDataAssociationsInner`

NewDeploymentDataAssociationsInner instantiates a new DeploymentDataAssociationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentDataAssociationsInnerWithDefaults

`func NewDeploymentDataAssociationsInnerWithDefaults() *DeploymentDataAssociationsInner`

NewDeploymentDataAssociationsInnerWithDefaults instantiates a new DeploymentDataAssociationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationType

`func (o *DeploymentDataAssociationsInner) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *DeploymentDataAssociationsInner) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *DeploymentDataAssociationsInner) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.


### GetValues

`func (o *DeploymentDataAssociationsInner) GetValues() []EntityAssociationValuesInner`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *DeploymentDataAssociationsInner) GetValuesOk() (*[]EntityAssociationValuesInner, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *DeploymentDataAssociationsInner) SetValues(v []EntityAssociationValuesInner)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


