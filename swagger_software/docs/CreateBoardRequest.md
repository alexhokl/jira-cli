# CreateBoardRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterId** | Pointer to **int64** |  | [optional] 
**Location** | Pointer to [**CreateBoardRequestLocation**](CreateBoardRequestLocation.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateBoardRequest

`func NewCreateBoardRequest() *CreateBoardRequest`

NewCreateBoardRequest instantiates a new CreateBoardRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBoardRequestWithDefaults

`func NewCreateBoardRequestWithDefaults() *CreateBoardRequest`

NewCreateBoardRequestWithDefaults instantiates a new CreateBoardRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterId

`func (o *CreateBoardRequest) GetFilterId() int64`

GetFilterId returns the FilterId field if non-nil, zero value otherwise.

### GetFilterIdOk

`func (o *CreateBoardRequest) GetFilterIdOk() (*int64, bool)`

GetFilterIdOk returns a tuple with the FilterId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterId

`func (o *CreateBoardRequest) SetFilterId(v int64)`

SetFilterId sets FilterId field to given value.

### HasFilterId

`func (o *CreateBoardRequest) HasFilterId() bool`

HasFilterId returns a boolean if a field has been set.

### GetLocation

`func (o *CreateBoardRequest) GetLocation() CreateBoardRequestLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *CreateBoardRequest) GetLocationOk() (*CreateBoardRequestLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *CreateBoardRequest) SetLocation(v CreateBoardRequestLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *CreateBoardRequest) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *CreateBoardRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateBoardRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateBoardRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateBoardRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *CreateBoardRequest) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CreateBoardRequest) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CreateBoardRequest) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CreateBoardRequest) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


