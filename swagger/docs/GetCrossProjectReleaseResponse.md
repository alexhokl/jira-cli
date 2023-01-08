# GetCrossProjectReleaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The cross-project release name. | [optional] 
**ReleaseIds** | Pointer to **[]int64** | The IDs of the releases included in the cross-project release. | [optional] 

## Methods

### NewGetCrossProjectReleaseResponse

`func NewGetCrossProjectReleaseResponse() *GetCrossProjectReleaseResponse`

NewGetCrossProjectReleaseResponse instantiates a new GetCrossProjectReleaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCrossProjectReleaseResponseWithDefaults

`func NewGetCrossProjectReleaseResponseWithDefaults() *GetCrossProjectReleaseResponse`

NewGetCrossProjectReleaseResponseWithDefaults instantiates a new GetCrossProjectReleaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetCrossProjectReleaseResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCrossProjectReleaseResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCrossProjectReleaseResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCrossProjectReleaseResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetReleaseIds

`func (o *GetCrossProjectReleaseResponse) GetReleaseIds() []int64`

GetReleaseIds returns the ReleaseIds field if non-nil, zero value otherwise.

### GetReleaseIdsOk

`func (o *GetCrossProjectReleaseResponse) GetReleaseIdsOk() (*[]int64, bool)`

GetReleaseIdsOk returns a tuple with the ReleaseIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseIds

`func (o *GetCrossProjectReleaseResponse) SetReleaseIds(v []int64)`

SetReleaseIds sets ReleaseIds field to given value.

### HasReleaseIds

`func (o *GetCrossProjectReleaseResponse) HasReleaseIds() bool`

HasReleaseIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


