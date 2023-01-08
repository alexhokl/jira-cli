# GetFieldAssociationParametersResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FieldId** | **string** |  | 
**Parameters** | Pointer to [**FieldAssociationParameters**](FieldAssociationParameters.md) |  | [optional] 
**WorkTypeParameters** | Pointer to [**[]WorkTypeParameters**](WorkTypeParameters.md) |  | [optional] 

## Methods

### NewGetFieldAssociationParametersResponse

`func NewGetFieldAssociationParametersResponse(fieldId string, ) *GetFieldAssociationParametersResponse`

NewGetFieldAssociationParametersResponse instantiates a new GetFieldAssociationParametersResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetFieldAssociationParametersResponseWithDefaults

`func NewGetFieldAssociationParametersResponseWithDefaults() *GetFieldAssociationParametersResponse`

NewGetFieldAssociationParametersResponseWithDefaults instantiates a new GetFieldAssociationParametersResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFieldId

`func (o *GetFieldAssociationParametersResponse) GetFieldId() string`

GetFieldId returns the FieldId field if non-nil, zero value otherwise.

### GetFieldIdOk

`func (o *GetFieldAssociationParametersResponse) GetFieldIdOk() (*string, bool)`

GetFieldIdOk returns a tuple with the FieldId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldId

`func (o *GetFieldAssociationParametersResponse) SetFieldId(v string)`

SetFieldId sets FieldId field to given value.


### GetParameters

`func (o *GetFieldAssociationParametersResponse) GetParameters() FieldAssociationParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *GetFieldAssociationParametersResponse) GetParametersOk() (*FieldAssociationParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *GetFieldAssociationParametersResponse) SetParameters(v FieldAssociationParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *GetFieldAssociationParametersResponse) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetWorkTypeParameters

`func (o *GetFieldAssociationParametersResponse) GetWorkTypeParameters() []WorkTypeParameters`

GetWorkTypeParameters returns the WorkTypeParameters field if non-nil, zero value otherwise.

### GetWorkTypeParametersOk

`func (o *GetFieldAssociationParametersResponse) GetWorkTypeParametersOk() (*[]WorkTypeParameters, bool)`

GetWorkTypeParametersOk returns a tuple with the WorkTypeParameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkTypeParameters

`func (o *GetFieldAssociationParametersResponse) SetWorkTypeParameters(v []WorkTypeParameters)`

SetWorkTypeParameters sets WorkTypeParameters field to given value.

### HasWorkTypeParameters

`func (o *GetFieldAssociationParametersResponse) HasWorkTypeParameters() bool`

HasWorkTypeParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


