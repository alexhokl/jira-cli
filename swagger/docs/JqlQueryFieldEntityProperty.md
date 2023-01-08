# JqlQueryFieldEntityProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | **string** | The object on which the property is set. | 
**Key** | **string** | The key of the property. | 
**Path** | **string** | The path in the property value to query. | 
**Type** | Pointer to **string** | The type of the property value extraction. Not available if the extraction for the property is not registered on the instance with the [Entity property](https://developer.atlassian.com/cloud/jira/platform/modules/entity-property/) module. | [optional] 

## Methods

### NewJqlQueryFieldEntityProperty

`func NewJqlQueryFieldEntityProperty(entity string, key string, path string, ) *JqlQueryFieldEntityProperty`

NewJqlQueryFieldEntityProperty instantiates a new JqlQueryFieldEntityProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJqlQueryFieldEntityPropertyWithDefaults

`func NewJqlQueryFieldEntityPropertyWithDefaults() *JqlQueryFieldEntityProperty`

NewJqlQueryFieldEntityPropertyWithDefaults instantiates a new JqlQueryFieldEntityProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *JqlQueryFieldEntityProperty) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *JqlQueryFieldEntityProperty) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *JqlQueryFieldEntityProperty) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetKey

`func (o *JqlQueryFieldEntityProperty) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *JqlQueryFieldEntityProperty) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *JqlQueryFieldEntityProperty) SetKey(v string)`

SetKey sets Key field to given value.


### GetPath

`func (o *JqlQueryFieldEntityProperty) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *JqlQueryFieldEntityProperty) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *JqlQueryFieldEntityProperty) SetPath(v string)`

SetPath sets Path field to given value.


### GetType

`func (o *JqlQueryFieldEntityProperty) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *JqlQueryFieldEntityProperty) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *JqlQueryFieldEntityProperty) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *JqlQueryFieldEntityProperty) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


