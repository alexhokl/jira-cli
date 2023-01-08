# BoardFeaturePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FeatureKey** | Pointer to **string** | The key of the feature | [optional] 
**State** | Pointer to **bool** | Whether the feature should be turned on or off | [optional] 

## Methods

### NewBoardFeaturePayload

`func NewBoardFeaturePayload() *BoardFeaturePayload`

NewBoardFeaturePayload instantiates a new BoardFeaturePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBoardFeaturePayloadWithDefaults

`func NewBoardFeaturePayloadWithDefaults() *BoardFeaturePayload`

NewBoardFeaturePayloadWithDefaults instantiates a new BoardFeaturePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFeatureKey

`func (o *BoardFeaturePayload) GetFeatureKey() string`

GetFeatureKey returns the FeatureKey field if non-nil, zero value otherwise.

### GetFeatureKeyOk

`func (o *BoardFeaturePayload) GetFeatureKeyOk() (*string, bool)`

GetFeatureKeyOk returns a tuple with the FeatureKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatureKey

`func (o *BoardFeaturePayload) SetFeatureKey(v string)`

SetFeatureKey sets FeatureKey field to given value.

### HasFeatureKey

`func (o *BoardFeaturePayload) HasFeatureKey() bool`

HasFeatureKey returns a boolean if a field has been set.

### GetState

`func (o *BoardFeaturePayload) GetState() bool`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *BoardFeaturePayload) GetStateOk() (*bool, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *BoardFeaturePayload) SetState(v bool)`

SetState sets State field to given value.

### HasState

`func (o *BoardFeaturePayload) HasState() bool`

HasState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


