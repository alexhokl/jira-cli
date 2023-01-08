# ProjectIdentifierBean

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The ID of the project. | [optional] [readonly] 
**Key** | Pointer to **string** | The key of the project. | [optional] [readonly] 

## Methods

### NewProjectIdentifierBean

`func NewProjectIdentifierBean() *ProjectIdentifierBean`

NewProjectIdentifierBean instantiates a new ProjectIdentifierBean object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectIdentifierBeanWithDefaults

`func NewProjectIdentifierBeanWithDefaults() *ProjectIdentifierBean`

NewProjectIdentifierBeanWithDefaults instantiates a new ProjectIdentifierBean object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ProjectIdentifierBean) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectIdentifierBean) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectIdentifierBean) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectIdentifierBean) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKey

`func (o *ProjectIdentifierBean) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *ProjectIdentifierBean) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *ProjectIdentifierBean) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *ProjectIdentifierBean) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


