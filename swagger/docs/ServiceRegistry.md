# ServiceRegistry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **NullableString** | service description | [optional] 
**Id** | Pointer to **string** | service ID | [optional] 
**Name** | Pointer to **string** | service name | [optional] 
**OrganizationId** | Pointer to **string** | organization ID | [optional] 
**Revision** | Pointer to **string** | service revision | [optional] 
**ServiceTier** | Pointer to [**ServiceRegistryTier**](ServiceRegistryTier.md) |  | [optional] 

## Methods

### NewServiceRegistry

`func NewServiceRegistry() *ServiceRegistry`

NewServiceRegistry instantiates a new ServiceRegistry object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRegistryWithDefaults

`func NewServiceRegistryWithDefaults() *ServiceRegistry`

NewServiceRegistryWithDefaults instantiates a new ServiceRegistry object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *ServiceRegistry) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServiceRegistry) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServiceRegistry) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServiceRegistry) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ServiceRegistry) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ServiceRegistry) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetId

`func (o *ServiceRegistry) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceRegistry) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceRegistry) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceRegistry) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServiceRegistry) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServiceRegistry) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServiceRegistry) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServiceRegistry) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganizationId

`func (o *ServiceRegistry) GetOrganizationId() string`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *ServiceRegistry) GetOrganizationIdOk() (*string, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *ServiceRegistry) SetOrganizationId(v string)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *ServiceRegistry) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRevision

`func (o *ServiceRegistry) GetRevision() string`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *ServiceRegistry) GetRevisionOk() (*string, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *ServiceRegistry) SetRevision(v string)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *ServiceRegistry) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetServiceTier

`func (o *ServiceRegistry) GetServiceTier() ServiceRegistryTier`

GetServiceTier returns the ServiceTier field if non-nil, zero value otherwise.

### GetServiceTierOk

`func (o *ServiceRegistry) GetServiceTierOk() (*ServiceRegistryTier, bool)`

GetServiceTierOk returns a tuple with the ServiceTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTier

`func (o *ServiceRegistry) SetServiceTier(v ServiceRegistryTier)`

SetServiceTier sets ServiceTier field to given value.

### HasServiceTier

`func (o *ServiceRegistry) HasServiceTier() bool`

HasServiceTier returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


