# RemoteLinkDataAssociationsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssociationType** | **string** | Defines the association type.  | 
**Values** | **[]string** | The service ID or keys to associate the entity with.  The number of values counted across all associationTypes must not exceed a limit of 500.  | 

## Methods

### NewRemoteLinkDataAssociationsInner

`func NewRemoteLinkDataAssociationsInner(associationType string, values []string, ) *RemoteLinkDataAssociationsInner`

NewRemoteLinkDataAssociationsInner instantiates a new RemoteLinkDataAssociationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoteLinkDataAssociationsInnerWithDefaults

`func NewRemoteLinkDataAssociationsInnerWithDefaults() *RemoteLinkDataAssociationsInner`

NewRemoteLinkDataAssociationsInnerWithDefaults instantiates a new RemoteLinkDataAssociationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssociationType

`func (o *RemoteLinkDataAssociationsInner) GetAssociationType() string`

GetAssociationType returns the AssociationType field if non-nil, zero value otherwise.

### GetAssociationTypeOk

`func (o *RemoteLinkDataAssociationsInner) GetAssociationTypeOk() (*string, bool)`

GetAssociationTypeOk returns a tuple with the AssociationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssociationType

`func (o *RemoteLinkDataAssociationsInner) SetAssociationType(v string)`

SetAssociationType sets AssociationType field to given value.


### GetValues

`func (o *RemoteLinkDataAssociationsInner) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *RemoteLinkDataAssociationsInner) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *RemoteLinkDataAssociationsInner) SetValues(v []string)`

SetValues sets Values field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


