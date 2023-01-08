# SwimlanesPayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomSwimlanes** | Pointer to [**[]SwimlanePayload**](SwimlanePayload.md) | The custom swimlane definitions. | [optional] 
**DefaultCustomSwimlaneName** | Pointer to **string** | The name of the custom swimlane to use for work items that don&#39;t match any other swimlanes. | [optional] 
**SwimlaneStrategy** | Pointer to **string** | The swimlane strategy for the board. | [optional] 

## Methods

### NewSwimlanesPayload

`func NewSwimlanesPayload() *SwimlanesPayload`

NewSwimlanesPayload instantiates a new SwimlanesPayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwimlanesPayloadWithDefaults

`func NewSwimlanesPayloadWithDefaults() *SwimlanesPayload`

NewSwimlanesPayloadWithDefaults instantiates a new SwimlanesPayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomSwimlanes

`func (o *SwimlanesPayload) GetCustomSwimlanes() []SwimlanePayload`

GetCustomSwimlanes returns the CustomSwimlanes field if non-nil, zero value otherwise.

### GetCustomSwimlanesOk

`func (o *SwimlanesPayload) GetCustomSwimlanesOk() (*[]SwimlanePayload, bool)`

GetCustomSwimlanesOk returns a tuple with the CustomSwimlanes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomSwimlanes

`func (o *SwimlanesPayload) SetCustomSwimlanes(v []SwimlanePayload)`

SetCustomSwimlanes sets CustomSwimlanes field to given value.

### HasCustomSwimlanes

`func (o *SwimlanesPayload) HasCustomSwimlanes() bool`

HasCustomSwimlanes returns a boolean if a field has been set.

### GetDefaultCustomSwimlaneName

`func (o *SwimlanesPayload) GetDefaultCustomSwimlaneName() string`

GetDefaultCustomSwimlaneName returns the DefaultCustomSwimlaneName field if non-nil, zero value otherwise.

### GetDefaultCustomSwimlaneNameOk

`func (o *SwimlanesPayload) GetDefaultCustomSwimlaneNameOk() (*string, bool)`

GetDefaultCustomSwimlaneNameOk returns a tuple with the DefaultCustomSwimlaneName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCustomSwimlaneName

`func (o *SwimlanesPayload) SetDefaultCustomSwimlaneName(v string)`

SetDefaultCustomSwimlaneName sets DefaultCustomSwimlaneName field to given value.

### HasDefaultCustomSwimlaneName

`func (o *SwimlanesPayload) HasDefaultCustomSwimlaneName() bool`

HasDefaultCustomSwimlaneName returns a boolean if a field has been set.

### GetSwimlaneStrategy

`func (o *SwimlanesPayload) GetSwimlaneStrategy() string`

GetSwimlaneStrategy returns the SwimlaneStrategy field if non-nil, zero value otherwise.

### GetSwimlaneStrategyOk

`func (o *SwimlanesPayload) GetSwimlaneStrategyOk() (*string, bool)`

GetSwimlaneStrategyOk returns a tuple with the SwimlaneStrategy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwimlaneStrategy

`func (o *SwimlanesPayload) SetSwimlaneStrategy(v string)`

SetSwimlaneStrategy sets SwimlaneStrategy field to given value.

### HasSwimlaneStrategy

`func (o *SwimlanesPayload) HasSwimlaneStrategy() bool`

HasSwimlaneStrategy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


