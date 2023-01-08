# StatusCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scope** | [**StatusScope**](StatusScope.md) |  | 
**Statuses** | [**[]StatusCreate**](StatusCreate.md) | Details of the statuses being created. | 

## Methods

### NewStatusCreateRequest

`func NewStatusCreateRequest(scope StatusScope, statuses []StatusCreate, ) *StatusCreateRequest`

NewStatusCreateRequest instantiates a new StatusCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusCreateRequestWithDefaults

`func NewStatusCreateRequestWithDefaults() *StatusCreateRequest`

NewStatusCreateRequestWithDefaults instantiates a new StatusCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScope

`func (o *StatusCreateRequest) GetScope() StatusScope`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *StatusCreateRequest) GetScopeOk() (*StatusScope, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *StatusCreateRequest) SetScope(v StatusScope)`

SetScope sets Scope field to given value.


### GetStatuses

`func (o *StatusCreateRequest) GetStatuses() []StatusCreate`

GetStatuses returns the Statuses field if non-nil, zero value otherwise.

### GetStatusesOk

`func (o *StatusCreateRequest) GetStatusesOk() (*[]StatusCreate, bool)`

GetStatusesOk returns a tuple with the Statuses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatuses

`func (o *StatusCreateRequest) SetStatuses(v []StatusCreate)`

SetStatuses sets Statuses field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


