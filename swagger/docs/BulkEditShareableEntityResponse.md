# BulkEditShareableEntityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Allowed action for bulk edit shareable entity | 
**EntityErrors** | Pointer to [**map[string]BulkEditActionError**](BulkEditActionError.md) | The mapping dashboard id to errors if any. | [optional] 

## Methods

### NewBulkEditShareableEntityResponse

`func NewBulkEditShareableEntityResponse(action string, ) *BulkEditShareableEntityResponse`

NewBulkEditShareableEntityResponse instantiates a new BulkEditShareableEntityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBulkEditShareableEntityResponseWithDefaults

`func NewBulkEditShareableEntityResponseWithDefaults() *BulkEditShareableEntityResponse`

NewBulkEditShareableEntityResponseWithDefaults instantiates a new BulkEditShareableEntityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *BulkEditShareableEntityResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *BulkEditShareableEntityResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *BulkEditShareableEntityResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetEntityErrors

`func (o *BulkEditShareableEntityResponse) GetEntityErrors() map[string]BulkEditActionError`

GetEntityErrors returns the EntityErrors field if non-nil, zero value otherwise.

### GetEntityErrorsOk

`func (o *BulkEditShareableEntityResponse) GetEntityErrorsOk() (*map[string]BulkEditActionError, bool)`

GetEntityErrorsOk returns a tuple with the EntityErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntityErrors

`func (o *BulkEditShareableEntityResponse) SetEntityErrors(v map[string]BulkEditActionError)`

SetEntityErrors sets EntityErrors field to given value.

### HasEntityErrors

`func (o *BulkEditShareableEntityResponse) HasEntityErrors() bool`

HasEntityErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


