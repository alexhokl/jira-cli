# ProjectWithDataPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPolicy** | Pointer to [**ProjectDataPolicy**](ProjectDataPolicy.md) | Data policy. | [optional] [readonly] 
**Id** | Pointer to **int64** | The project ID. | [optional] [readonly] 

## Methods

### NewProjectWithDataPolicy

`func NewProjectWithDataPolicy() *ProjectWithDataPolicy`

NewProjectWithDataPolicy instantiates a new ProjectWithDataPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProjectWithDataPolicyWithDefaults

`func NewProjectWithDataPolicyWithDefaults() *ProjectWithDataPolicy`

NewProjectWithDataPolicyWithDefaults instantiates a new ProjectWithDataPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPolicy

`func (o *ProjectWithDataPolicy) GetDataPolicy() ProjectDataPolicy`

GetDataPolicy returns the DataPolicy field if non-nil, zero value otherwise.

### GetDataPolicyOk

`func (o *ProjectWithDataPolicy) GetDataPolicyOk() (*ProjectDataPolicy, bool)`

GetDataPolicyOk returns a tuple with the DataPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPolicy

`func (o *ProjectWithDataPolicy) SetDataPolicy(v ProjectDataPolicy)`

SetDataPolicy sets DataPolicy field to given value.

### HasDataPolicy

`func (o *ProjectWithDataPolicy) HasDataPolicy() bool`

HasDataPolicy returns a boolean if a field has been set.

### GetId

`func (o *ProjectWithDataPolicy) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProjectWithDataPolicy) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProjectWithDataPolicy) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ProjectWithDataPolicy) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


