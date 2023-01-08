# ProjectIdentifiers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **int64** | The ID of the created project. | [readonly] 
**Key** | **string** | The key of the created project. | [readonly] 
**Self** | **string** | The URL of the created project. | [readonly] 

## Methods

### NewProjectIdentifiers

`func NewProjectIdentifiers(id int64, key string, self string, ) *ProjectIdentifiers`

NewProjectIdentifiers instantiates a new ProjectIdentifiers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIdentifiersWithDefaults

`func NewProjectIdentifiersWithDefaults() *ProjectIdentifiers`

NewProjectIdentifiersWithDefaults instantiates a new ProjectIdentifiers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectIdentifiers) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectIdentifiers) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectIdentifiers) SetId(v int64)`

SetId sets Id field to given value.


### GetKey

`func (o *ProjectIdentifiers) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectIdentifiers) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectIdentifiers) SetKey(v string)`

SetKey sets Key field to given value.


### GetSelf

`func (o *ProjectIdentifiers) GetSelf() string`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *ProjectIdentifiers) GetSelfOk() (*string, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *ProjectIdentifiers) SetSelf(v string)`

SetSelf sets Self field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


